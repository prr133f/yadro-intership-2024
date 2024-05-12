package main

import (
	"bufio"
	"log"
	"os"
	"yadro-intership/internal"
)

func main() {
	file, err := os.OpenFile(os.Args[1], os.O_RDONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var (
		in = bufio.NewReader(file)
	)

	if err := internal.Run(in); err != nil {
		log.Fatal(err)
	}
}
