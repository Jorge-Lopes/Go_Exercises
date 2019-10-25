package main

import (
	"fmt"
	"runtime"
)

const goARCH string = runtime.GOARCH
const goOS string = runtime.GOOS

func main() {
	fmt.Println(goARCH)
	fmt.Println(goOS)
}
