package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := os.Args
	if len(input) == 3 {
		file1, err := os.Open(input[1])
		if err == nil {
			file2, err := os.Create(input[2])
			if err == nil {
				reader := bufio.NewScanner(file1)
				writer := bufio.NewWriter(file2)
				for reader.Scan() {
					text := reader.Text()
					writer.WriteString(text);writer.WriteString("\n")
				}
				if err := reader.Err(); err != nil {
					fmt.Println(err)
				}
				writer.Flush()
			} else {
				fmt.Println(err)
			}
			defer file2.Close()
		} else {
			fmt.Println(err)
		}
		defer file1.Close()
	} else {
		fmt.Println("Expects two input, Run program with the command below\ngo run . <input filename> <output filename>")
	}
}
