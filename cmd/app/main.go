package main

import (
	"fmt"
	"reflect"
	blah "github.com/cometbft/cometbft/types"
)

func main() {
  fmt.Println("hello world from vulnerable application")
  fmt.Println(reflect.TypeOf(blah.Block{}))
}
