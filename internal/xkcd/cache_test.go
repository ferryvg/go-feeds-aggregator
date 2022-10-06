package xkcd_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/ferryvg/go-feeds-aggregator/internal/xkcd"
	mock_xkcd "github.com/ferryvg/go-feeds-aggregator/internal/xkcd/mock"
)

var CorrectSourceItem = &xkcd.SourceItem{
	Id:          1,
	ContentUrl:  "",
	Day:         24,
	Month:       7,
	Year:        2009,
	Title:       "Woodpecker",
	SafeTitle:   "Woodpecker",
	ImageUrl:    "https://imgs.xkcd.com/comics/woodpecker.png",
	Description: "[[A man with a beret and a woman are standing on a boardwalk]]",
}

func TestCache_Set(t *testing.T) {
	type mockCache struct {
		item *xkcd.SourceItem
		err  error
	}

	tests := []struct {
		name      string
		item      *xkcd.SourceItem
		wantErr   error
		mockCache mockCache
	}{
		{
			name:    "success",
			item:    CorrectSourceItem,
			wantErr: nil,
			mockCache: mockCache{
				item: CorrectSourceItem,
				err:  nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m := mock_xkcd.NewMockCache(ctrl)

			m.EXPECT().Set(tt.mockCache.item).Return(tt.mockCache.err)

			err := m.Set(tt.item)
			assert.EqualValues(t, tt.wantErr, err)
		})
	}
}

func TestCache_Get(t *testing.T) {
	type mockCache struct {
		id   int
		item *xkcd.SourceItem
		err  error
	}

	tests := []struct {
		name      string
		id        int
		want      *xkcd.SourceItem
		wantErr   error
		mockCache mockCache
	}{
		{
			name:    "success",
			id:      CorrectSourceItem.Id,
			want:    CorrectSourceItem,
			wantErr: nil,
			mockCache: mockCache{
				id:   CorrectSourceItem.Id,
				item: CorrectSourceItem,
				err:  nil,
			},
		},
		{
			name:    "err_record_not_found",
			id:      9999,
			want:    nil,
			wantErr: xkcd.NotFoundErr,
			mockCache: mockCache{
				id:   9999,
				item: nil,
				err:  xkcd.NotFoundErr,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m := mock_xkcd.NewMockCache(ctrl)

			m.EXPECT().Get(tt.mockCache.id).Return(tt.mockCache.item, tt.mockCache.err)

			result, err := m.Get(tt.id)
			assert.EqualValues(t, tt.wantErr, err)
			assert.EqualValues(t, tt.want, result)
		})
	}
}
