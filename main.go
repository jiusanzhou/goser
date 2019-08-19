package main

import (
	"fmt"
	"os"

	"go.zoe.im/goser/cmd"
)

func main() {
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
