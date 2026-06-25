package main

import (
	"bufio"
	"fmt"
	"os"
	m "transformers/transformations"
)

func ReadAndWrite() {
	if len(os.Args) != 3 {
		fmt.Println("Expects 3 arguments. Run program with the following command\ngo run . <input filename> <output filename>")
		return
	}
	inputFile := os.Args[1]
	outputFile := os.Args[2]

	readFile, err := os.Open(inputFile)
	if err == nil {
		writeFile, err := os.Create(outputFile)
		if err == nil {
			reader := bufio.NewScanner(readFile)
			if reader.Err() != nil {
				fmt.Println("An error occurred while reading file")
			}
			writer := bufio.NewWriter(writeFile)
			for reader.Scan() {
				text := reader.Text()
				text = m.HexBin(text)
				text = m.Cases(text)
				text = m.CaseN(text)
				text = m.Puncs(text)
				text = m.Punc(text)
				writer.WriteString(text);writer.WriteString("\n")
			}
			writer.Flush()
		}
		defer writeFile.Close()
	}
	defer readFile.Close()
}

func main() {
	ReadAndWrite()
}
