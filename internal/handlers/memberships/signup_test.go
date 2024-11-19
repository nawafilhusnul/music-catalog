package memberships

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	membershipsmodel "github.com/nawafilhusnul/music-catalog/internal/models/memberships"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

func Test_handler_SignUp(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockSvc := NewMockservice(ctrlMock)

	tests := []struct {
		name               string
		mockFn             func()
		expectedStatusCode int
		jsonBody           string
	}{
		{
			name: "success",
			mockFn: func() {
				mockSvc.EXPECT().SignUp(gomock.Any(), &membershipsmodel.SignUpRequest{
					Email:    "test@example.com",
					Username: "testuser",
					Password: "testpassword",
				}).Return(nil)
			},
			jsonBody:           `{"email": "test@example.com", "username": "testuser", "password": "testpassword"}`,
			expectedStatusCode: http.StatusCreated,
		},
		{
			name:               "failed when binding json",
			mockFn:             func() {},
			jsonBody:           `{"email": "test@example.com", "username": "testuser"`,
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "failed when sign up",
			mockFn: func() {
				model := &membershipsmodel.SignUpRequest{
					Email:    "test@example.com",
					Username: "testuser",
					Password: "testpassword",
				}
				mockSvc.EXPECT().SignUp(gomock.Any(), model).Return(assert.AnError)
			},
			jsonBody:           `{"email": "test@example.com", "username": "testuser", "password": "testpassword"}`,
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
			endpoint := "/memberships/signup"

			body := bytes.NewReader([]byte(tt.jsonBody))
			req, err := http.NewRequest(http.MethodPost, endpoint, body)
			assert.NoError(t, err)

			h.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
		})
	}
}
