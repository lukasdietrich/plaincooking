package service

import (
	"context"
	"database/sql"
	"errors"
	"io"
	"time"

	"github.com/rs/xid"
	"github.com/spf13/viper"

	"github.com/lukasdietrich/plaincooking/internal/database/models"
)

var (
	_ io.WriteCloser = &assetWriter{}
	_ io.Reader      = &assetReader{}
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
