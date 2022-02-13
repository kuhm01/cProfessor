package ast

import "professorc/src/token"

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

type YeoseokStatement struct {
	Token token.Token
	Value int64
}

type StaticVarStatement struct {
	Token    token.Token
	Index    int
	Constant Expression
	Level    int
}

type MajorFlagStatement struct {
	Token      token.Token
	Identifier string
}

type JumptoFlagStatement struct {
	Token          token.Token
	JumpIdentifier string
}

type ProfessorTobuffStatement struct {
	Token token.Token
	Index int
	Type  string
}

type PrintBufferStatement struct {
	Token token.Token
}

type ConstantExpression struct {
	Token token.Token
	Value int64
}

type ClassStatement struct {
	Token token.Token
	Type  string
}

type NewProfessorStatement struct {
	Token token.Token
}
