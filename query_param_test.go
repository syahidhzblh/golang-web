package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func SayHello(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if name == "" {
		fmt.Fprint(writer, "Hello")
	} else {
		fmt.Fprintf(writer, "Hello %s", name)
	}
}

func TestQueryParam(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=Hisbul", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	respones := recorder.Result()
	body, _ := io.ReadAll(respones.Body)

	fmt.Println(string(body))
}

func MultipleQueryParameter(writer http.ResponseWriter, request *http.Request) {
	first_name := request.URL.Query().Get("first_name")
	last_name := request.URL.Query().Get("last_name")

	fmt.Fprintf(writer, "Hello %s %s", first_name, last_name)
}

func TestMultipleQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?first_name=Syahid&last_name=Hizbullah", nil)
	recorder := httptest.NewRecorder()

	MultipleQueryParameter(recorder, request)

	respones := recorder.Result()
	body, _ := io.ReadAll(respones.Body)

	fmt.Println(string(body))
}

func MultipleParameterValues(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	names := query["name"]

	fmt.Fprint(writer, strings.Join(names, " "))
}

func TestMultipleParameterValues(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=Syahid&name=Hizbullah&name=Ganteng", nil)
	recorder := httptest.NewRecorder()

	MultipleParameterValues(recorder, request)

	respones := recorder.Result()
	body, _ := io.ReadAll(respones.Body)

	fmt.Println(string(body))
}
