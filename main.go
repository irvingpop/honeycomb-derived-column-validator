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

type Config struct {
	FilePath string
	Verbose  bool
}

type Processor struct {
	config Config
}

func readInput(filePath string) (string, error) {
	if filePath != "" {
		data, err := os.ReadFile(filePath)
		if err != nil {
			return "", fmt.Errorf("error reading file: %v", err)
		}
		return string(data), nil
	}

	scanner := bufio.NewScanner(os.Stdin)
	var sb strings.Builder
	for scanner.Scan() {
		line := scanner.Text()
		sb.WriteString(line + "\n")
	}
	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading input: %v", err)
	}
	return sb.String(), nil
}

func (p *Processor) processInput(input string) (string, bool, error) {
	// Try to parse as JSON first
	var jsonData map[string]interface{}
	if err := json.Unmarshal([]byte(input), &jsonData); err == nil {
		// If JSON parsing succeeds, look for 'expression' field
		if p.config.Verbose {
			fmt.Println("Received JSON: ", input)
		}
		if expr, ok := jsonData["expression"].(string); ok {
			return expr, true, nil
		}
		return "", false, fmt.Errorf("missing 'expression' value in JSON")
	}
	// it wasn't JSON, so proceed in non-JSON mode
	return input, false, nil
}

func main() {
	filePath := flag.String("f", "", "Path to input file")
	verbose := flag.Bool("v", false, "Verbose mode prints extra info")
	flag.Parse()

	config := Config{
		FilePath: *filePath,
		Verbose:  *verbose,
	}
	processor := &Processor{config: config}

	srcText, err := readInput(config.FilePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	expressionText, jsonMode, err := processor.processInput(srcText)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	if err := ParseDerived(expressionText, config.Verbose); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing derived column: %v\n", err)
		os.Exit(1)
	}

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
