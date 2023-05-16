package main

import (
	"main/reader"
	"main/api"
)

func main() {
	reader.LoadFile()

	api.HandleRequests()
}