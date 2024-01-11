package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		panic(err)
	}

	first_name := request.PostForm.Get("first_name")
	last_name := request.PostForm.Get("last_name")
	fmt.Fprintf(writer, "Hello %s %s", first_name, last_name)
}

func TestFormPost(t *testing.T) {
	requestBody := strings.NewReader("first_name=Syahid&last_name=Hizbullah")
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080", requestBody)
	recorder := httptest.NewRecorder()
	request.Header.Add("content-type", "application/x-www-form-urlencoded")

	FormPost(recorder, request)

	respones := recorder.Result()
	body, _ := io.ReadAll(respones.Body)
	fmt.Println(string(body))
}
