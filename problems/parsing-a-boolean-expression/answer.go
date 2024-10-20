package main

import (
	"errors"
	"fmt"
)

type Tokenizer struct {
	tokens   []rune
	curIdx   int
	curToken rune
}

func NewTokenizer(s string) *Tokenizer {
	tokens := []rune(s)
	curIdx := 0
	curToken := tokens[0]
	return &Tokenizer{
		tokens:   tokens,
		curIdx:   curIdx,
		curToken: curToken,
	}
}

func (t *Tokenizer) Done() bool {
	return t.curIdx >= len(t.tokens)
}

func (t *Tokenizer) Peek() rune {
	return t.tokens[t.curIdx]
}

func (t *Tokenizer) Advance() rune {
	curToken := t.tokens[t.curIdx]
	t.curIdx += 1
	return curToken
}

func isTrueValue(token rune) bool {
	return token == 't'
}

func isFalseValue(token rune) bool {
	return token == 'f'
}

func isLogicalNot(token rune) bool {
	return token == '!'
}

func isLogicalAnd(token rune) bool {
	return token == '&'
}

func isLogicalOr(token rune) bool {
	return token == '|'
}

func isPrimitive(token rune) bool {
	return isTrueValue(token) || isFalseValue(token)
}

func isUnaryOp(token rune) bool {
	return isLogicalNot(token)
}

func isBinaryOp(token rune) bool {
	return isLogicalAnd(token) || isLogicalOr(token)
}

func isLeftBracket(token rune) bool {
	return token == '('
}

func isRightBracket(token rune) bool {
	return token == ')'
}

func isComma(token rune) bool {
	return token == ','
}

type Expression interface{}

type Primitive struct {
	Val rune
}

type UnaryOp struct {
	Operator rune
	Operand  Expression
}

type BinaryOp struct {
	Operator rune
	Left     Expression
	Right    Expression
}

type Parser struct {
	tokenizer Tokenizer
}

func NewParser(tokenizer Tokenizer) *Parser {
	return &Parser{
		tokenizer: tokenizer,
	}
}

func (p *Parser) Parse() Expression {
	curToken := p.tokenizer.Peek()
	if isPrimitive(curToken) {
		return p.parsePrimitive()
	} else if isUnaryOp(curToken) {
		return p.parseUnaryOp()
	} else {
		return p.parseBinaryOp()
	}
}

func (p *Parser) parsePrimitive() Primitive {
	val := p.tokenizer.Advance()
	return Primitive{
		Val: val,
	}
}

func (p *Parser) parseUnaryOp() UnaryOp {
	operator := p.tokenizer.Advance()
	p.tokenizer.Advance() // '('
	operand := p.Parse()
	p.tokenizer.Advance() // ')'
	return UnaryOp{
		Operator: operator,
		Operand:  operand,
	}
}

func (p *Parser) parseBinaryOp() BinaryOp {
	operator := p.tokenizer.Advance()
	p.tokenizer.Advance() // '('

	left := p.Parse()

	nextToken := p.tokenizer.Advance()
	if isRightBracket(nextToken) {
		return BinaryOp{
			Operator: operator,
			Left:     left,
			Right:    left,
		}
	}

	right := p.Parse()
	nextToken = p.tokenizer.Advance()
	for isComma(nextToken) {
		nextRight := p.Parse()
		right = BinaryOp{
			Left:     right,
			Operator: operator,
			Right:    nextRight,
		}
		nextToken = p.tokenizer.Advance()
	}

	return BinaryOp{
		Operator: operator,
		Left:     left,
		Right:    right,
	}
}

type Evaluator struct{}

func NewEvaluator() *Evaluator {
	return &Evaluator{}
}

func (e *Evaluator) Evaluate(expr Expression) (bool, error) {
	switch v := expr.(type) {
	case Primitive:
		if isTrueValue(v.Val) {
			return true, nil
		} else if isFalseValue(v.Val) {
			return false, nil
		} else {
			return false, errors.New("Unsupported primitive value.")
		}
	case UnaryOp:
		if isLogicalNot(v.Operator) {
			operand, err := e.Evaluate(v.Operand)
			if err != nil {
				return false, err
			}
			return !(operand), nil
		} else {
			return false, errors.New("Unsupported unary operator.")
		}
	case BinaryOp:
		left, err := e.Evaluate(v.Left)
		if err != nil {
			return false, err
		}
		right, err := e.Evaluate(v.Right)
		if err != nil {
			return false, err
		}
		if isLogicalAnd(v.Operator) {
			return left && right, nil
		} else if isLogicalOr(v.Operator) {
			return left || right, nil
		} else {
			return false, errors.New("Unsupported binary operator.")
		}
	default:
		return false, errors.New("Unsupported expression type.")
	}
}

func parseBoolExpr(expression string) bool {
	parser := NewParser(*NewTokenizer(expression))
	expr := parser.Parse()

	evaluator := NewEvaluator()
	res, err := evaluator.Evaluate(expr)
	if err != nil {
		return false
	}
	return res
}

func main() {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"&(|(f))", false},
		{"|(f,f,f,t)", true},
		{"!(&(f,t))", true},
	}

	for _, tc := range testCases {
		result := parseBoolExpr(tc.input)
		if result != tc.expected {
			fmt.Printf("parseBoolExpr(%s) = %t; expected %t\n", tc.input, result, tc.expected)
		}
	}
}
