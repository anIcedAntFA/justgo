package main

import (
	"fmt"
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

	err := initApp()
	if err != nil {
		// initApp: loadConfig(mise.tom): open mise.tom: no such file or directory
		//
		fmt.Println(err)
	}
}

func validate(age int) error {
	if age < 0 {
		// return errors.New("age must be non-negative")
		return fmt.Errorf("invalid age: %d must be non-negative", age)
	}
	return nil
}

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
