package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/joerdav/example/foo"
)

func main() {
	fmt.Println(uuid.NewString())
	foo.DoCalculations()
}
