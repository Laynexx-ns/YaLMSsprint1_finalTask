package test

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"sprint1_finalTask/internal/api/handlers"
	"strings"
	"testing"
)

// p.s. я знаю, что можно было без доп типов, но мне лень уже исправлять((
type ReqBody struct {
	Expression string `json:"expression"`
}

type RespBody struct {
	Result float64 `json:"result"`
}

// йоу тесты
func getTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	router.POST("/api/v1/calculate", handlers.CalcMiddleware(), handlers.CalcHandler)
	return router
}

func TestCalcResult(t *testing.T) {
	router := getTestRouter()

	tests := []struct {
		Req      string  `json:"expression" binding:"required"`
		Expected float64 `json:"result" binding:"required"`
	}{
		{"2+3/2", 3.5},
		{"2/42+6*12", 72.04761904761905},
		{"123*(32+16)", 5904},
	}

	for _, test := range tests {
		testJson, err := json.Marshal(ReqBody{Expression: test.Req})
		if err != nil {
			t.Fatal(err)
		}
		req, _ := http.NewRequest("POST", "/api/v1/calculate", strings.NewReader(string(testJson)))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Fatalf("Wrong code: %v", w.Code)
		}

		var resp RespBody
		err = json.Unmarshal(w.Body.Bytes(), &resp)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, test.Expected, resp.Result)

	}

}

func TestCalcWrongBody(t *testing.T) {
	router := getTestRouter()

	tests := []struct {
		Req string `json:"expression" binding:"required"`
	}{
		{`{"expression":"qwe"`},
		{`{"expression":"3_23+28"}`},
		{`{"expression":"a123409"}`},
		{`{"expression":"@1203"}`},
	}
	for _, test := range tests {
		req := httptest.NewRequest("POST", "/api/v1/calculate", strings.NewReader(string(test.Req)))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		if w.Code == http.StatusOK {
			t.Fatalf("Wrong code: %v", w.Code)
		}
	}
}
