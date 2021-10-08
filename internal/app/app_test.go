package app

import (
	"fmt"
	"github.com/Lvas1418/Test_task_change_html-/testTask/internal/httpx"
	"github.com/Lvas1418/Test_task_change_html-/testTask/internal/service"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"regexp"
	"testing"
)

func TestGetArticleID(t *testing.T) {
	tests := []struct {
		name           string
		rec            *httptest.ResponseRecorder
		req            *http.Request
		expectedName   string
		expectedCode   int
		expectedHeader string
	}{
		{
			name:         "OK_Tom",
			rec:          httptest.NewRecorder(),
			req:          httptest.NewRequest("GET", "/Tom", nil),
			expectedName: `Tom`,
			expectedCode: http.StatusOK,
		},
		{
			name:         "OK_Bob",
			rec:          httptest.NewRecorder(),
			req:          httptest.NewRequest("GET", "/Bob", nil),
			expectedName: `Bob`,
			expectedCode: http.StatusOK,
		},
	}
	err := os.Setenv("TEMPLATE_PATH", "../../")
	if err != nil {
		log.Fatal("unable to read template file ")
	}
	serviceTest := service.NewService()
	handler := httpx.NewHandler(serviceTest)
	router := handler.Init()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			router.ServeHTTP(test.rec, test.req)
			re := regexp.MustCompile(fmt.Sprintf("Hello, %s", test.expectedName))
			if test.rec.Code != test.expectedCode {
				t.Errorf("Got: \t\t%v\n\tExpected: \t%v\n", test.rec.Code, test.expectedCode)
			}
			if !re.MatchString(test.rec.Body.String()) {
				t.Errorf("Got: \t\t%s\n\tExpected: \t%s\n", test.rec.Body.String(), test.expectedName)
			}

		})
	}
}
