package Tests

import (
	Controllers "ChatApp/Src/Controllers"
	"ChatApp/Src/DBConnector"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUserPing(t *testing.T) {
	tests := []struct {
		name       string
		httpMethod string
		url        string
		expected   string
	}{
		{
			name:       "TestPingSuccess",
			httpMethod: "GET",
			url:        "/api/v1/user/ping",
			expected:   `{"message": "pong"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := setupRouter("/api/v1/user/ping", Controllers.GetUserPing)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest(tt.httpMethod, tt.url, nil)
			router.ServeHTTP(w, req)

			assert.Equal(t, http.StatusOK, w.Code)
			assert.JSONEq(t, tt.expected, w.Body.String())
		})
	}
}

func TestGetAllUsers(t *testing.T) {
	// Connect to the database
	err := DBConnector.ConnectToMongo()
	if err != nil {
		t.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer DBConnector.DisconnectFromMongo()

	client := DBConnector.GetMongoClient()
	collection := client.Database("ChatApp").Collection("users")
	if collection == nil {
		t.Fatalf("Failed to get 'users' collection")
	}

	testCase := struct {
		name       string
		httpMethod string
		url        string
	}{
		name:       "TestGetAllUsersSuccess",
		httpMethod: "GET",
		url:        "/api/v1/user",
	}

	// Run the test case
	t.Run(testCase.name, func(t *testing.T) {
		router := setupRouter(testCase.url, func(c *gin.Context) {
			Controllers.GetAllUsers(c, collection)
		})

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(testCase.httpMethod, testCase.url, nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var responseBody map[string]interface{}

		assert.NotZero(t, responseBody["users"])
	})
}
