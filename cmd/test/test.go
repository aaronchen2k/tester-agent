package main

import (
	"fmt"
	"log"
)

func main() {
	const maxSize = 5 << 20
	a := maxSize + 1<<20
	log.Println(fmt.Sprintf("%d", a))
}
