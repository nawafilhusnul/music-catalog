package memberships

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	membershipsmodel "github.com/nawafilhusnul/music-catalog/internal/models/memberships"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

func Test_handler_SignIn(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()
	mockSvc := NewMockservice(ctrlMock)

	tests := []struct {
		name                 string
		mockFn               func()
		expectedStatusCode   int
		expectedBodyResponse membershipsmodel.SignInResponse
		jsonBody             string
		wantErr              bool
	}{
		{
			name: "success",
			mockFn: func() {
				mockSvc.EXPECT().SignIn(gomock.Any(), &membershipsmodel.SignInRequest{
					Email:    "test@example.com",
					Password: "qwerty123",
				}).Return(&membershipsmodel.SignInResponse{
					AccessToken: "access_token",
				}, nil)
			},
			expectedStatusCode: http.StatusOK,
			expectedBodyResponse: membershipsmodel.SignInResponse{
				AccessToken: "access_token",
			},
			jsonBody: `{"email": "test@example.com", "password": "qwerty123"}`,
			wantErr:  false,
		},
		{
			name: "failed to bind json body",
			mockFn: func() {
			},
			jsonBody:           `{"email": "test@example.com", "password": "qwerty1`,
			wantErr:            true,
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "failed to sign in",
			mockFn: func() {
				mockSvc.EXPECT().SignIn(gomock.Any(), &membershipsmodel.SignInRequest{
					Email:    "test@example.com",
					Password: "qwerty123",
				}).Return(nil, assert.AnError)
			},
			jsonBody:           `{"email": "test@example.com", "password": "qwerty123"}`,
			wantErr:            true,
			expectedStatusCode: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn()
			api := gin.New()
			h := &handler{
				Engine: api,
				svc:    mockSvc,
			}

			h.RegisterRoutes()

			w := httptest.NewRecorder()
			endpoint := "/memberships/signin"

			body := bytes.NewReader([]byte(tt.jsonBody))
			req, err := http.NewRequest(http.MethodPost, endpoint, body)
			assert.NoError(t, err)

			h.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)

			res := w.Result()
			defer res.Body.Close()

			response := membershipsmodel.SignInResponse{}
			err = json.Unmarshal(w.Body.Bytes(), &response)
			assert.NoError(t, err)

			assert.Equal(t, tt.expectedBodyResponse, response)
		})
	}
}
