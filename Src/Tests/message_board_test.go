package Tests

import (
	Controllers "ChatApp/Src/Controllers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupRouter(route string, handler gin.HandlerFunc) *gin.Engine {
	router := gin.Default()
	router.GET(route, handler)
	return router
}
func TestGetMessageBoardPing(t *testing.T) {
	tests := []struct {
		name       string
		httpMethod string
		url        string
		expected   string
	}{
		{
			name:       "TestPingSuccess",
			httpMethod: "GET",
			url:        "/api/v1/messageboard/ping",
			expected:   `{"message": "pong"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := setupRouter("/api/v1/messageboard/ping", Controllers.GetMessagesBoardPing)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest(tt.httpMethod, tt.url, nil)
			router.ServeHTTP(w, req)

			assert.Equal(t, http.StatusOK, w.Code)
			assert.JSONEq(t, tt.expected, w.Body.String())
		})
	}
}
