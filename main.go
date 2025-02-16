package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/irvingpop/honeycomb-derived-column-validator/pkg/parser"
)

func main() {
	filePath := flag.String("f", "", "Path to input file")
	verbose := flag.Bool("v", false, "Verbose mode prints extra info")
	flag.Parse()

	var srcText string

	if *filePath != "" {
		// Read the entire file content
		data, err := os.ReadFile(*filePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
			os.Exit(1)
		}
		srcText = string(data)
	} else {
		// Read from STDIN using a scanner, echoing each line
		scanner := bufio.NewScanner(os.Stdin)
		var sb strings.Builder
		for scanner.Scan() {
			line := scanner.Text()
			// For example, print each line (could also transform it if needed)
			fmt.Println(line)
			sb.WriteString(line + "\n")
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
			os.Exit(1)
		}
		srcText = sb.String()
	}

	if err := ParseDerived(srcText, *verbose); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing derived column: %v\n", err)
		os.Exit(1)
	}
}

func ParseDerived(src string, verbose bool) error {
	_, err := parser.ANTLRParse(src, verbose)
	return err
}
