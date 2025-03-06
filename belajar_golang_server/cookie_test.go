package belajargolangserver

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetCookie(writer http.ResponseWriter, request *http.Request) {
	cookie := new(http.Cookie)
	cookie.Name = "Tournyaka"
	cookie.Value = request.URL.Query().Get("name")
	cookie.Path = "/"

	http.SetCookie(writer, cookie)

	fmt.Fprint(writer, "Succes Cookie created")
}

func GetCookie(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("Tournyaka")
	if err != nil {
		fmt.Fprint(writer, "No cookie found")
	} else {
		name := cookie.Value
		fmt.Fprintf(writer, "Hello %s", name)
	}
}

func DeleteCookie(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("Tournyaka")
	if err != nil {
		fmt.Fprint(writer, "No cookie found")
	} else {
		cookie.MaxAge = -1
		http.SetCookie(writer, cookie)
		fmt.Fprint(writer, "Cookie deleted")
	}
}

func TestCookie(test *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/set-cookie", SetCookie)
	mux.HandleFunc("/get-cookie", GetCookie)
	mux.HandleFunc("/delete-cookie", DeleteCookie)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}

func TestSetCookie(test *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/?name=Ilham", nil)
	recorder := httptest.NewRecorder()

	SetCookie(recorder, request)

	cookies := recorder.Result().Cookies()
	for _, cookie := range cookies {
		fmt.Printf("cookie %s %s \n", cookie.Name, cookie.Value)
	}
}

func TestGetCookie(test *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	cookie := new(http.Cookie)
	cookie.Name = "Tournyaka"
	cookie.Value = "Ilham"
	request.AddCookie(cookie)

	recorder := httptest.NewRecorder()

	GetCookie(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}
