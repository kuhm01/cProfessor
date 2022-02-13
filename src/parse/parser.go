package parse

import (
	"professorc/src/ast"
	"professorc/src/token"
)

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.YEOSEOK:
		return p.parseYeoseokStatement()

	case token.STATICVAR:
		ps := p.curToken.Literal

		switch ps[9] {
		case '!':
			return p.parseInBufferFromStaticVarStatement()
		case '?':
			return p.parseStaticVarOperatingStatement()
		}

	case token.FLAG:
		return p.parseMajorFlag()

	case token.JUMPPOINT:
		return p.parseJumptoMajorFlag()

	case token.PRINTER:
		return &ast.PrintBufferStatement{Token: p.curToken}

	case token.PLUSCLASS:
		return &ast.ClassStatement{Token: p.curToken, Type: "보강"}

	case token.MINUSCLASS:
		return &ast.ClassStatement{Token: p.curToken, Type: "휴강"}
	}

	return nil
}
