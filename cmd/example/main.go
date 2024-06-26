package main

import (
	"flag"
	"fmt"
	"io"
	"my-project2/postfix"
	"os"
	"strings"
)

type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (h *ComputeHandler) Compute() error {
	inputBytes, err := io.ReadAll(h.Input)
	if err != nil {
		return err
	}

	postfixExpr := strings.TrimSpace(string(inputBytes))
	if postfixExpr == "" {
		return fmt.Errorf("input expression is empty")
	}

	result, err := postfix.PostfixToPrefix(postfixExpr)
	if err != nil {
		return err
	}

	_, err = io.WriteString(h.Output, result)
	return err
}

func main() {
	expr := flag.String("e", "", "Postfix expression")
	inputFile := flag.String("f", "", "Input file")
	outputFile := flag.String("o", "", "Output file")
	flag.Parse()

	var input io.Reader
	var output io.Writer = os.Stdout

	if *expr != "" {
		input = strings.NewReader(*expr)
	} else if *inputFile != "" {
		file, err := os.Open(*inputFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error opening input file:", err)
			os.Exit(1)
		}
		defer file.Close()
		input = file
	} else {
		fmt.Fprintln(os.Stderr, "No input expression provided")
		os.Exit(1)
	}

	if *outputFile != "" {
		file, err := os.Create(*outputFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error creating output file:", err)
			os.Exit(1)
		}
		defer file.Close()
		output = file
	}

	handler := &ComputeHandler{
		Input:  input,
		Output: output,
	}

	if err := handler.Compute(); err != nil {
		fmt.Fprintln(os.Stderr, "Error computing expression:", err)
		os.Exit(1)
	}
}
