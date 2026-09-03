package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
)

func main() {
	// data, err := readConfig("lefthook")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(data))

	// err := validate(-1)
	// fmt.Println(err)

	// err := initApp()
	// if err != nil {
	// 	// initApp: loadConfig(mise.tom): open mise.tom: no such file or directory
	// 	//
	// 	fmt.Println(err)
	// }

	// _, err := os.ReadFile("missing.txt")
	// fmt.Println(err == os.ErrNotExist)
	// if err == os.ErrNotExist {
	// 	log.Fatal("file does not exist")
	// 	fmt.Println("file does not exist")
	// }
	// if errors.Is(err, os.ErrNotExist) {
	// 	fmt.Println("fmt: file does not exisst")
	// 	log.Fatal("log: file does not exist")
	// }

	// Somewhere an error is created and wrapped:
	// err := fmt.Errorf("processing failed: %w", &ValidationError{Field: "email", Value: "bad"})
	// Extract the ValidationError from the chain:
	// var valErr *ValidationError
	// if errors.As(err, &valErr) {
	// 	// valErr is now the extracted *ValidationError — access its fields
	// 	fmt.Printf("validation failed on field %s with value %v\n", valErr.Field, valErr.Value)
	// 	fmt.Printf("field %q had bad value: %v\n", valErr.Field, valErr.Value)
	// }
	// errors.AsType — Go 1.26: the type is a parameter, the value is returned
	// if valueErr, ok := errors.AsType[*ValidationError](err); ok {
	// 	fmt.Println(valueErr.Field)
	// }

	// err := fetch("https://api.ngockhoi96.dev/tasks")
	// var httpErr *HTTPError
	// if errors.As(err, &httpErr) {
	// 	if httpErr.StatusCode == 429 {
	// 		fmt.Println("rate limited")
	// 	}
	// }
	// if valErr, ok := errors.AsType[*HTTPError](err); ok {
	// 	fmt.Println(valErr.StatusCode)
	// }

	err := validate(User{Name: "haha", Age: -23})
	fmt.Println(err)
	fmt.Println(errors.Is(err, ErrEmptyName))
	fmt.Println(errors.Is(err, ErrNegativeAge))
}

type User struct {
	Name string
	Age  int
}

var (
	ErrEmptyName   = errors.New("name is empty")
	ErrNegativeAge = errors.New("age must be non-negative")
)

func validate(u User) error {
	var errs []error
	if u.Name == "" {
		errs = append(errs, ErrEmptyName)
	}
	if u.Age < 0 {
		errs = append(errs, ErrNegativeAge)
	}
	return errors.Join(errs...)
}

type HTTPError struct {
	StatusCode int
	URL        string
	Err        error
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("HTTP %d from %s: %v", e.StatusCode, e.URL, e.Err)
}

func (e *HTTPError) Unwrap() error {
	return e.Err
}

func fetch(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return &HTTPError{StatusCode: 0, URL: url, Err: err}
	}
	if resp.StatusCode >= 400 {
		return &HTTPError{
			StatusCode: resp.StatusCode,
			URL:        url,
			Err:        fmt.Errorf("bad status"),
		}
	}
	return nil
}

type ValidationError struct {
	Field string
	Value any
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation failed on %s", e.Field)
}

var ErrNotExistMine = errors.New("file does not exist mine")

// func validate(age int) error {
// 	if age < 0 {
// 		// return errors.New("age must be non-negative")
// 		return fmt.Errorf("invalid age: %d must be non-negative", age)
// 	}
// 	return nil
// }

func initApp() error {
	_, err := readConfig("mise.tom")
	if err != nil {
		// return err
		return fmt.Errorf("initApp: %w", err)
	}
	return nil
}

func readConfig(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("loadConfig(%s): %w", path, err)
		// return nil, err
	}
	return data, nil
}
