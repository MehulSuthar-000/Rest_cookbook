package main

import (
	"log"

	"github.com/mehulsuthar-000/base62/base62"
)

func main() {
	x := 100
	base62string := base62.ToBase62(x)
	log.Println(base62string)
	normalNumber := base62.ToBase10(base62string)
	log.Println(normalNumber)
}
