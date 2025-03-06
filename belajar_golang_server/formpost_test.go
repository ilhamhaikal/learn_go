package belajargolangserver

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(writer http.ResponseWriter, request *http.Request) {
	/*
		ini cara penggunaan secara manual agar tau prosesnya seperti apa
		yang dilakukan oleh ParseForm adalah mengambil data dari request body
		kemudian memparsingnya dan hasilnya disimpan di request.PostForm
		jika terjadi error, maka proses parsing akan dihentikan dan error akan dikembalikan
		ada baiknya untuk mengecek error yang dikembalikan oleh ParseForm
		agar kita tahu apakah proses parsingnya berhasil atau tidak
		jika berhasil, kita bisa melanjutkan proses selanjutnya
		jika gagal, kita bisa memberikan response error ke client
		adapun cara lain agar tidak perlu mengecek error adalah dengan menggunakan ParseFormValues
		yang akan mengembalikan error jika terjadi error dan tidak menghentikan proses parsing
		ini adalah cara otomatis yang lebih mudah digunakan
	*/
	err := request.ParseForm()
	if err != nil {
		panic(err)
	}

	//menggunakan proses parsing otomatis
	// request.ParseFormValues("first_name=Ilham&last_name=Haikal")

	firstName := request.PostForm.Get("first_name")
	lastName := request.PostForm.Get("last_name")

	fmt.Fprintf(writer, "Hello %s %s", firstName, lastName)
}

func TestFormPost(t *testing.T) {
	requesBody := strings.NewReader("first_name=Ilham&last_name=Haikal")
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080", requesBody)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	FormPost(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)
	fmt.Println(bodyString)
}
