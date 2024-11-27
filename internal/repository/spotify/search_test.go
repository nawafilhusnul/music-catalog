package spotify

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/nawafilhusnul/music-catalog/internal/configs"
	"github.com/nawafilhusnul/music-catalog/pkg/httpclient"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_outbound_Search(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockHTTPClient := httpclient.NewMockHTTPClient(mockCtrl)

	type args struct {
		query  string
		limit  int
		offset int
	}
	tests := []struct {
		name    string
		args    args
		want    *SearchResponse
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				query:  "bohemian rhapsody",
				limit:  10,
				offset: 0,
			},
			want:    searchResponseStruct,
			wantErr: false,
			mockFn: func(args args) {
				params := url.Values{}
				params.Add("q", args.query)
				params.Add("type", "track")
				params.Add("limit", strconv.Itoa(args.limit))
				params.Add("offset", strconv.Itoa(args.offset))

				basePath := "https://api.spotify.com/v1/search"
				urlPath := fmt.Sprintf("%s?%s", basePath, params.Encode())

				req, err := http.NewRequest(http.MethodGet, urlPath, nil)
				token := fmt.Sprintf("%s %s", "token_type", "access_token")
				req.Header.Set("Authorization", token)

				assert.NoError(t, err)

				mockHTTPClient.EXPECT().Do(req).Return(&http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(bytes.NewBufferString(searchResponse)),
				}, nil)
			},
		},
		{
			name: "failed",
			args: args{
				query:  "bohemian rhapsody",
				limit:  10,
				offset: 0,
			},
			wantErr: true,
			want:    nil,
			mockFn: func(args args) {
				mockHTTPClient.EXPECT().Do(gomock.Any()).Return(nil, assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		tt.mockFn(tt.args)
		t.Run(tt.name, func(t *testing.T) {
			o := &outbound{
				cfg:         &configs.Config{},
				client:      mockHTTPClient,
				AccessToken: "access_token",
				TokenType:   "token_type",
				ExpiredAt:   time.Now().Add(time.Hour),
			}
			got, err := o.Search(context.Background(), tt.args.query, tt.args.limit, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("outbound.Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("outbound.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
