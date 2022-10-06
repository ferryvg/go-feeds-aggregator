package xkcd_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ferryvg/go-feeds-aggregator/internal/xkcd"
)

func TestService_Load(t *testing.T) {
	tests := []struct {
		name      string
		limit     int
		wantError error
	}{
		{
			name:      "success",
			limit:     2,
			wantError: nil,
		},
		{
			name:      "incorrect_limit",
			limit:     -1,
			wantError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := xkcd.NewService()

			_, err := svc.Load(context.Background(), tt.limit)
			assert.EqualValues(t, tt.wantError, err)
		})
	}
}
