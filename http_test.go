package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HelloHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello World")
}

func TestHttp(t *testing.T) {
	request := httptest.NewRequest("GET", "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	HelloHandler(recorder, request)
	respones := recorder.Result()
	body, _ := io.ReadAll(respones.Body)
	bodyString := string(body)
	fmt.Println(bodyString)

}
