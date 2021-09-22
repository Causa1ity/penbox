package main

import (
	"github.com/causality/penbox/pkg/api"
	"os"
)

func main() {
	api.Run(os.Args[1])
}
