package main

import (
	"bufio"
	"encoding/json"
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

	var srcText, expressionText string
	jsonMode := false

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
			sb.WriteString(line + "\n")
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
			os.Exit(1)
		}
		srcText = sb.String()
	}

	// Try to parse as JSON first
	var jsonData map[string]interface{}
	if err := json.Unmarshal([]byte(srcText), &jsonData); err == nil {
		// If JSON parsing succeeds, look for 'expression' field
		if *verbose {
			fmt.Println("Received JSON: ", srcText)
		}
		if expr, ok := jsonData["expression"].(string); ok {
			jsonMode = true
			expressionText = expr
		} else {
			fmt.Fprintf(os.Stderr, "Missing 'expression' value: %v\n", err)
			os.Exit(1)
		}
	} else {
		// it wasn't JSON, so proceed in non-JSON mode
		expressionText = srcText
	}

	if err := ParseDerived(expressionText, *verbose); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing derived column: %v\n", err)
		os.Exit(1)
	}

	// output JSON meant for the TF external data source, with "expression" as the key name
	if jsonMode {
		result := map[string]string{
			"expression": expressionText,
		}
		jsonResult, err := json.Marshal(result)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating JSON response: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(string(jsonResult))
	}
}

func ParseDerived(src string, verbose bool) error {
	_, err := parser.ANTLRParse(src, verbose)
	return err
}
