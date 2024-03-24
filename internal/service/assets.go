package service

import (
	"bytes"
	"context"
	"database/sql"
	"encoding"
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"io"
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
	_ io.WriteCloser           = &assetWriter{}
	_ io.Reader                = &assetReader{}
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

func (s *AssetService) Writer(ctx context.Context, filename, mediaTyp string) (*assetWriter, error) {
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

	w := assetWriter{
		Asset:   asset,
		ctx:     ctx,
		querier: querier,
		buffer:  make([]byte, viper.GetSizeInBytes("asset.chunk.size")),
	}

	return &w, nil
}

func (s *AssetService) Reader(ctx context.Context, id xid.ID) (*assetReader, error) {
	querier := s.transactions.Querier(ctx)

	asset, err := querier.ReadAsset(ctx, models.ReadAssetParams{ID: id})
	if err != nil {
		return nil, err
	}

	r := assetReader{
		ReadAssetRow: asset,
		ctx:          ctx,
		querier:      querier,
	}

	return &r, nil
}

type assetWriter struct {
	models.Asset

	ctx     context.Context
	querier models.Querier

	buffer []byte
	n      int
}

func (w *assetWriter) Close() error {
	return w.flushChunk(true)
}

func (w *assetWriter) Write(b []byte) (int, error) {
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

func (w *assetWriter) flushChunk(always bool) error {
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

	if err := w.querier.CreateAssetChunk(w.ctx, createAssetChunkParams); err != nil {
		return err
	}

	w.n = 0
	return nil
}

type assetReader struct {
	models.ReadAssetRow

	ctx     context.Context
	querier models.Querier

	offset xid.ID
	buffer []byte
	n      int
}

func (r *assetReader) Read(b []byte) (int, error) {
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

func (r *assetReader) advanceChunk() error {
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

func (s *AssetService) ThumbnailReader(ctx context.Context, id xid.ID, size ThumbnailSize) (*thumbnailReader, error) {
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

	var buffer bytes.Buffer

	if err := jpeg.Encode(&buffer, img, nil); err != nil {
		return nil, err
	}

	r := thumbnailReader{
		Reader:    &buffer,
		TotalSize: int64(buffer.Len()),
		MediaType: mediaTypeJpeg,
	}

	return &r, nil
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

type thumbnailReader struct {
	io.Reader
	MediaType string
	TotalSize int64
}
