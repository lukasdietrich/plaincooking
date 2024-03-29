package service

import (
	"context"
	"database/sql"
	"encoding"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"log/slog"
	"time"

	"github.com/muesli/smartcrop"
	"github.com/muesli/smartcrop/nfnt"
	"github.com/rs/xid"
	"github.com/spf13/viper"

	"github.com/lukasdietrich/plaincooking/internal/database/models"
)

const (
	mediaTypePng  = "image/png"
	mediaTypeJpeg = "image/jpeg"
)

var (
	_ io.WriteCloser           = &AssetWriter{}
	_ io.Reader                = &AssetReader{}
	_ encoding.TextUnmarshaler = new(ThumbnailSize)
)

func init() {
	viper.SetDefault("asset.chunk.size", "64kb")
}

type AssetService struct {
	transactions *TransactionService
}

func NewAssetService(transactions *TransactionService) *AssetService {
	return &AssetService{
		transactions: transactions,
	}
}

func (s *AssetService) Writer(ctx context.Context, filename, mediaTyp string) (*AssetWriter, error) {
	querier := s.transactions.Querier(ctx)

	createAssetParams := models.CreateAssetParams{
		ID:        xid.New(),
		CreatedAt: time.Now(),
		Filename:  filename,
		MediaType: mediaTyp,
	}

	asset, err := querier.CreateAsset(ctx, createAssetParams)
	if err != nil {
		return nil, err
	}

	w := AssetWriter{
		Asset:   asset,
		ctx:     ctx,
		querier: querier,
		buffer:  make([]byte, viper.GetSizeInBytes("asset.chunk.size")),
	}

	slog.Info("writing asset", slog.Any("id", asset.ID))
	return &w, nil
}

func (s *AssetService) Reader(ctx context.Context, id xid.ID) (*AssetReader, error) {
	querier := s.transactions.Querier(ctx)

	asset, err := querier.ReadAsset(ctx, models.ReadAssetParams{ID: id})
	if err != nil {
		return nil, err
	}

	r := AssetReader{
		ReadAssetRow: asset,
		ctx:          ctx,
		querier:      querier,
	}

	slog.Debug("reading asset", slog.Any("id", asset.ID))
	return &r, nil
}

type AssetWriter struct {
	models.Asset

	ctx     context.Context
	querier models.Querier

	buffer []byte
	n      int
}

func (w *AssetWriter) Close() error {
	return w.flushChunk(true)
}

func (w *AssetWriter) Write(b []byte) (int, error) {
	l := 0

	for l < len(b) {

		n := copy(w.buffer[w.n:], b[l:])
		w.n += n
		l += n

		if err := w.flushChunk(false); err != nil {
			return l, err
		}
	}

	return l, nil
}

func (w *AssetWriter) flushChunk(always bool) error {
	if w.n == 0 {
		return nil
	}

	if !always && w.n < len(w.buffer) {
		return nil
	}

	createAssetChunkParams := models.CreateAssetChunkParams{
		ID:      xid.New(),
		AssetID: w.ID,
		Content: w.buffer[:w.n],
	}

	slog.Debug("writing asset chunk", slog.Any("id", createAssetChunkParams.ID))

	if err := w.querier.CreateAssetChunk(w.ctx, createAssetChunkParams); err != nil {
		return err
	}

	w.n = 0
	return nil
}

type AssetReader struct {
	models.ReadAssetRow

	ctx     context.Context
	querier models.Querier

	offset xid.ID
	buffer []byte
	n      int
}

func (r *AssetReader) Read(b []byte) (int, error) {
	if r.n < len(r.buffer) {
		n := copy(b, r.buffer[r.n:])
		r.n += n

		return n, nil
	}

	if err := r.advanceChunk(); err != nil {
		return 0, err
	}

	return r.Read(b)
}

func (r *AssetReader) advanceChunk() error {
	chunk, err := r.querier.ReadAssetChunk(r.ctx, models.ReadAssetChunkParams{
		AssetID:  r.ID,
		IDOffset: r.offset,
	})

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return io.EOF
		}

		return err
	}

	slog.Debug("reading asset chunk", slog.Any("id", chunk.ID))

	r.offset = chunk.ID
	r.buffer = chunk.Content
	r.n = 0

	return nil
}

type ThumbnailSize uint8

const (
	ThumbnailUnknown ThumbnailSize = iota
	ThumbnailTile
	ThumbnailBanner
)

func (s ThumbnailSize) String() string {
	switch s {
	case ThumbnailTile:
		return "tile"
	case ThumbnailBanner:
		return "banner"
	default:
		return "unknown"
	}
}

func (s *ThumbnailSize) UnmarshalText(text []byte) error {
	switch string(text) {
	case "tile":
		*s = ThumbnailTile
	case "banner":
		*s = ThumbnailBanner
	default:
		*s = ThumbnailUnknown
	}

	return nil
}

func (s *ThumbnailSize) dimensions() image.Point {
	switch *s {
	case ThumbnailTile:
		return image.Point{X: 750, Y: 256}
	case ThumbnailBanner:
		return image.Point{X: 1000, Y: 300}
	default:
		return image.Point{X: 0, Y: 0}
	}
}

func (s *AssetService) ThumbnailReader(ctx context.Context, id xid.ID, size ThumbnailSize) (*AssetReader, error) {
	querier := s.transactions.Querier(ctx)

	readAssetThumbnailParams := models.ReadAssetThumbnailParams{
		AssetID: id,
		Size:    size.String(),
	}

	thumbnail, err := querier.ReadAssetThumbnail(ctx, readAssetThumbnailParams)
	if errors.Is(err, sql.ErrNoRows) {
		newThumbnailId, err := s.persistThumbnail(ctx, id, size)
		if err != nil {
			return nil, err
		}

		return s.Reader(ctx, newThumbnailId)
	}

	if err != nil {
		return nil, err
	}

	return s.Reader(ctx, thumbnail.ThumbnailAssetID)
}

func (s *AssetService) persistThumbnail(ctx context.Context, id xid.ID, size ThumbnailSize) (xid.ID, error) {
	slog.Info("generating thumbnail for asset", slog.Any("id", id), slog.Any("size", size))

	img, err := s.generateThumbnail(ctx, id, size)
	if err != nil {
		return xid.NilID(), err
	}

	filename := fmt.Sprintf("%s.thumbnail-%s.jpeg", id, size)

	w, err := s.Writer(ctx, filename, mediaTypeJpeg)
	if err != nil {
		return xid.NilID(), err
	}

	if err := jpeg.Encode(w, img, &jpeg.Options{Quality: 80}); err != nil {
		return xid.NilID(), err
	}

	if err := w.Close(); err != nil {
		return xid.NilID(), err
	}

	createAssetThumbnailParams := models.CreateAssetThumbnailParams{
		AssetID:          id,
		ThumbnailAssetID: w.ID,
		Size:             size.String(),
	}

	querier := s.transactions.Querier(ctx)
	if err := querier.CreateAssetThumbnail(ctx, createAssetThumbnailParams); err != nil {
		return xid.NilID(), err
	}

	return w.ID, nil
}

func (s *AssetService) generateThumbnail(ctx context.Context, id xid.ID, size ThumbnailSize) (image.Image, error) {
	img, err := s.readImage(ctx, id)
	if err != nil {
		return nil, err
	}

	resizer := nfnt.NewDefaultResizer()
	analyzer := smartcrop.NewAnalyzer(resizer)
	dimensions := size.dimensions()

	crop, err := analyzer.FindBestCrop(img, dimensions.X, dimensions.Y)
	if err != nil {
		return nil, err
	}

	img = img.(interface {
		SubImage(image.Rectangle) image.Image
	}).SubImage(crop)

	if img.Bounds().Dx() != dimensions.X || img.Bounds().Dy() != dimensions.Y {
		img = resizer.Resize(img, uint(dimensions.X), uint(dimensions.Y))
	}

	return img, nil
}

func (s *AssetService) readImage(ctx context.Context, id xid.ID) (image.Image, error) {
	r, err := s.Reader(ctx, id)
	if err != nil {
		return nil, err
	}

	if r.MediaType == mediaTypePng {
		return png.Decode(r)
	}

	if r.MediaType == mediaTypeJpeg {
		return jpeg.Decode(r)
	}

	return nil, image.ErrFormat
}
