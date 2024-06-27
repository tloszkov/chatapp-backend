package Tests

import (
	Controllers "ChatApp/Src/Controllers"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetGroupMessagesPing(t *testing.T) {
	tests := []struct {
		name       string
		httpMethod string
		url        string
		expected   string
	}{
		{
			name:       "TestPingSuccess",
			httpMethod: "GET",
			url:        "/api/v1/groupmessage/ping",
			expected:   `{"message": "pong"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := setupRouter("/api/v1/groupmessage/ping", Controllers.GetGroupMessagesPing)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest(tt.httpMethod, tt.url, nil)
			router.ServeHTTP(w, req)

			assert.Equal(t, http.StatusOK, w.Code)
			assert.JSONEq(t, tt.expected, w.Body.String())
		})
	}
}
