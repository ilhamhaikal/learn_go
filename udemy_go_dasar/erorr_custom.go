package main

import "fmt"

//interface error
type ValidationError struct {
	Msg string
}

func (v *ValidationError) Error() string {
	return v.Msg
}

type notFoundError struct {
	Msg string
}

func (v *notFoundError) Error() string {
	return v.Msg
}

func SaveData(id string, data any) error {
	if id == "" {
		// karna kita menggunakan interface error, kita bisa mengembalikan error dengan menggunakan pointer ke &ValidationError{"Validasi Eror"}
		return &ValidationError{"Validation Erorr"}
	}
	if id != "ilham" {
		return &notFoundError{"Data Not Found"}
	}
	// Save data ok
	return nil
}

func main() {

	/*
		kita bisa menggunakan 2 cara untuk menangani error
		1. menggunakan if else
		2. menggunakan switch case
	*/
	err := SaveData("", nil)
	if err != nil {
		// if ValidationError, ok := err.(*ValidationError); ok {
		// 	fmt.Println("Validation Error", ValidationError.Error())
		// } else if NotFoundError, ok := err.(*notFoundError); ok {
		// 	fmt.Println("Data Not Found", NotFoundError.Error())
		// } else {
		// 	fmt.Println("Unknown Error")
		// }

		switch finalError := err.(type) {
		case *ValidationError:
			fmt.Println("Validation Error", finalError.Error())
		case *notFoundError:
			fmt.Println("Data Not Found", finalError.Error())
		default:
			fmt.Println("Unknown Error", finalError.Error())
		}
	} else {
		fmt.Println("Data Saved")
	}
}
