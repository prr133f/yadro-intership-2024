package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"yadro-intership/internal"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var (
		in  = bufio.NewReader(file)
		out = bufio.NewWriter(os.Stdout)
	)
	defer out.Flush()

	if tables, err := internal.Run(in, out); err != nil {
		log.Fatal(err)
	} else {
		for k, v := range tables.Map {
			fmt.Fprintln(out, k, v.Margin, v.TimeInWork)
		}
	}
}
