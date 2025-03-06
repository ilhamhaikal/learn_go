package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

/*
cara menjalankan unit test pada golang dengan perintah
go test ./... (untuk menjalankan semua unit test pada semua folder)
go test ./helper (untuk menjalankan unit test pada folder helper atau folder yang di tuju)
go test ./helper -v (untuk menjalankan unit test pada folder helper atau folder yang di tuju dengan verbose)
go test ./helper -v -cover (untuk menjalankan unit test pada folder helper atau folder yang di tuju dengan verbose dan melihat coverage)
go test ./helper -v -coverprofile=cover.out (untuk menjalankan unit test pada folder helper atau folder yang di tuju dengan verbose dan melihat coverage dan menyimpan hasil coverage ke file cover.out)
go tool cover -func=cover.out (untuk melihat hasil coverage dari file cover.out)
go tool cover -html=cover.out (untuk melihat hasil coverage dari file cover.out dalam bentuk html)

**bisa juga menggunakan testify untuk melakukan unit test pada golang
go get github.com/stretchr/testify
gunakan assertion, skiptest dan require untuk melakukan unit test

bisa juga menggunakan test main untuk menjalankan unit test sebelum dan sesudah unit test dijalankan pada folder tertentu
tidak bisa multiple test main pada package yang berbeda

**benchmark test yaitu test yang digunakan untuk mengukur performa dari suatu function
cara menjalankan benchmark test pada golang dengan perintah
go test -bench=. (untuk menjalankan benchmark test pada semua function)
go test -v -bench=HelloWorld (untuk menjalankan benchmark test pada function HelloWorld)
go test -v -run=TestTidakAda -bench=. (untuk menjalankan benchmark tanpa unit test)
go test -v -run=TestTidakAda -bench=HelloWorld (untuk menjalankan benchmark tanpa unit test pada function HelloWorld)
dan pada function yang ingin di benchmark tambahkan kata Benchmark pada awal nama function contohnya func BenchmarkTest (b *testing.B)
*/

// benchmark test
func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Ilham")
	}
}

// unit test table test menggunakan struct	slice
func TestHelloWorldTableTest(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Ilham",
			request:  "Ilham",
			expected: "Hello, Ilham",
		},
		{
			name:     "Haikal",
			request:  "Haikal",
			expected: "Hello, Haikal",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result, "Result must be "+test.expected)
		})
	}
}

// unit test sub test
func TestSubTest(t *testing.T) {
	t.Run("Ilham", func(t *testing.T) {
		result := HelloWorld("Ilham")
		require.Equal(t, "Hello, Ilham", result, "Result must be Hello, Ilham")
		fmt.Println("TestSubTest Ilham done")
	})

	t.Run("Haikal", func(t *testing.T) {
		result := HelloWorld("Haikal")
		require.Equal(t, "Hello, Haikal", result, "Result must be Hello, Haikal")
		fmt.Println("TestSubTest Haikal done")
	})
}

// unit test beffor and after
func TestMain(m *testing.M) {
	// before unit test
	fmt.Println("BEFORE UNIT TEST")
	// run unit test
	m.Run()
	// after unit test
	fmt.Println("AFTER UNIT TEST")
}

// Unit test menggunakan fail() artinya jika test case gagal maka akan tetap melanjutkan test case berikutnya
func TestHelloWorldIlham(t *testing.T) {
	result := HelloWorld("Ilham")
	if result != "Hello, Ilham" {
		t.Fail()
	}
	fmt.Println("TestHelloWorldIlham done")
}

// Unit test menggunakan failNow() artinya jika test case gagal maka akan langsung berhenti
func TestHelloWorldHaikal(t *testing.T) {
	result := HelloWorld("Haikal")
	if result != "Hello, Haikal" {
		t.FailNow()
	}
	fmt.Println("TestHelloWorldHaikal done")
}

// Unit test menggunakan error() artinya jika test case gagal maka akan menampilkan error message
func TestHelloWorldMahmud(t *testing.T) {
	result := HelloWorld("Mahmud")
	if result != "Hello, Mahmud" {
		t.Error("Result must be Hello, Mahmud")
	}
	fmt.Println("TestHelloWorldMahmud done")
}

// Unit test menggunakan fatal() artinya jika test case gagal maka akan langsung berhenti dan menampilkan error message
func TestHelloWorldIntel(t *testing.T) {
	result := HelloWorld("Intel")
	if result != "Hello, Intel" {
		t.Fatal("Result must be Hello, Intel")
	}
	fmt.Println("TestHelloWorldIntel done")
}

// unit test assertion menggunakan equal artinya jika test case gagal maka akan langsung berhenti dan menampilkan error message
func TestHelloWorldAssertion(t *testing.T) {
	result := HelloWorld("Mantap")
	assert.Equal(t, "Hello, Mantap", result, "Result must be Hello, Mantap")
	fmt.Println("TestHelloWorldAssertion With Assert done")
}

// unit test require menggunakan equal artinya jika test case gagal maka akan langsung berhenti dan menampilkan error message
func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Gurih")
	require.Equal(t, "Hello, Gurih", result, "Result must be Hello, Gurih")
	fmt.Println("TestHelloWorldRequire With Equal done")
}

// unit test skip test
func TestHelloWorldSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Can't run on windows")
	}
	result := HelloWorld("Skip")
	require.Equal(t, "Hello, Skip", result, "Result must be Hello, Skip")
	fmt.Println("TestHelloWorldSkip done")
}
