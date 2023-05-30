package a

import "fmt"

func MyPublicFunction() {
	err := myPrivateFunction()
	if err != nil {
		fmt.Println(err)
	}
}

func myPrivateFunction() error {
	err := doSomething()
	return err
}
