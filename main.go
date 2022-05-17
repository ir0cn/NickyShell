package main

import (
	"NickyShell/nicky/service"
)

func main() {
	if err := service.Start(); err != nil {
		panic(err)
	}
}
