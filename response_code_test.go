package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponesCode(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if name == "" {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(writer, "name is empty")
	} else {
		fmt.Fprintf(writer, "Hello %s", name)
	}
}

func TestResponseError(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	ResponesCode(recorder, request)

	respones := recorder.Result()
	body, _ := io.ReadAll(respones.Body)

	fmt.Println(respones.StatusCode)
	fmt.Println(respones.Status)
	fmt.Println(string(body))
}

func TestResponseValid(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/?name=hisbul", nil)
	recorder := httptest.NewRecorder()

	ResponesCode(recorder, request)

	respones := recorder.Result()
	body, _ := io.ReadAll(respones.Body)

	fmt.Println(respones.Status)
	fmt.Println(string(body))
}
