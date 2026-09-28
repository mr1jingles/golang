package main

import (
	"errors"
	"log"
)

func returnError(a, b int) error {
	if a == b {
		err := errors.New("error returned!")
		return err
	}

	return nil
}

func main() {
	err := returnError(10, 2)

	if err != nil {
		log.Println(err)
	} else {
		log.Println("ok")
	}

	err = returnError(42, 42)

	if err != nil {
		log.Println(err)
	} else {
		log.Println("ok")
	}

	if err.Error() == "error returned!" {
		log.Println("!!!")
	}
}
