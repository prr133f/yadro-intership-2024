package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"yadro-intership/internal"
	"yadro-intership/pkg/utils"
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
		keys := make([]int, 0, len(tables.Map))
		for k := range tables.Map {
			keys = append(keys, k)
		}
		slices.Sort(keys)
		for _, k := range keys {
			fmt.Fprintln(out, k, tables.Map[k].Margin, utils.DurationToFormatString(tables.Map[k].TimeInWork))
		}
	}
}
