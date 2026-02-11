package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type parser struct {
	input string
	pos   int
}

func (p *parser) skipWhitespace() {
	for p.pos < len(p.input) && unicode.IsSpace(rune(p.input[p.pos])) {
		p.pos++
	}
}

func (p *parser) parseExpression() (float64, error) {
	left, err := p.parseTerm()
	if err != nil {
		return 0, err
	}

	for {
		p.skipWhitespace()
		if p.pos >= len(p.input) {
			return left, nil
		}

		op := p.input[p.pos]
		if op != '+' && op != '-' {
			return left, nil
		}
		p.pos++

		right, err := p.parseTerm()
		if err != nil {
			return 0, err
		}

		if op == '+' {
			left += right
		} else {
			left -= right
		}
	}
}

func (p *parser) parseTerm() (float64, error) {
	left, err := p.parseFactor()
	if err != nil {
		return 0, err
	}

	for {
		p.skipWhitespace()
		if p.pos >= len(p.input) {
			return left, nil
		}

		op := p.input[p.pos]
		if op != '*' && op != '/' {
			return left, nil
		}
		p.pos++

		right, err := p.parseFactor()
		if err != nil {
			return 0, err
		}

		if op == '*' {
			left *= right
		} else {
			if right == 0 {
				return 0, errors.New("division by zero")
			}
			left /= right
		}
	}
}

func (p *parser) parseFactor() (float64, error) {
	p.skipWhitespace()
	if p.pos >= len(p.input) {
		return 0, errors.New("unexpected end of expression")
	}

	if p.input[p.pos] == '(' {
		p.pos++
		value, err := p.parseExpression()
		if err != nil {
			return 0, err
		}
		p.skipWhitespace()
		if p.pos >= len(p.input) || p.input[p.pos] != ')' {
			return 0, errors.New("missing closing parenthesis")
		}
		p.pos++
		return value, nil
	}

	if p.input[p.pos] == '+' || p.input[p.pos] == '-' {
		sign := p.input[p.pos]
		p.pos++
		value, err := p.parseFactor()
		if err != nil {
			return 0, err
		}
		if sign == '-' {
			return -value, nil
		}
		return value, nil
	}

	start := p.pos
	hasDecimal := false
	for p.pos < len(p.input) {
		c := p.input[p.pos]
		if c == '.' {
			if hasDecimal {
				break
			}
			hasDecimal = true
			p.pos++
			continue
		}
		if c < '0' || c > '9' {
			break
		}
		p.pos++
	}

	if start == p.pos {
		return 0, fmt.Errorf("unexpected character: %q", p.input[p.pos])
	}

	numberText := p.input[start:p.pos]
	value, err := strconv.ParseFloat(numberText, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid number %q", numberText)
	}
	return value, nil
}

func evaluate(expression string) (float64, error) {
	p := &parser{input: expression}
	value, err := p.parseExpression()
	if err != nil {
		return 0, err
	}
	p.skipWhitespace()
	if p.pos != len(p.input) {
		return 0, fmt.Errorf("unexpected trailing input: %q", p.input[p.pos:])
	}
	return value, nil
}

func repl() {
	fmt.Println("Simple Calculator")
	fmt.Println("Enter an expression (e.g. (2+3)*4/5). Type 'exit' to quit.")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			fmt.Println()
			return
		}

		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		if strings.EqualFold(line, "exit") || strings.EqualFold(line, "quit") {
			return
		}

		result, err := evaluate(line)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}
		fmt.Printf("= %g\n", result)
	}
}

func main() {
	if len(os.Args) > 1 {
		expression := strings.Join(os.Args[1:], " ")
		result, err := evaluate(expression)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%g\n", result)
		return
	}

	repl()
}
