package evaluate

import (
	"fmt"
	"professorc/src/ast"
	"professorc/src/environment"
)

func evalStaticVarStatement(node ast.Node, env *environment.Environment, pc *ProgramCounter) {
	stmt := node.(*ast.StaticVarStatement)
	index := stmt.Index
	if index > 6 {
		fmt.Printf("Not exist %dth Professor.", index)
		return
	}
	level := stmt.Level
	value := stmt.Constant.(*ast.ConstantExpression).Value
	resultvalue := value * int64(level)
	env.Professor.Setter(int(resultvalue), index)
}

func evalBuffFromProfessor(node ast.Node, env *environment.Environment, pc *ProgramCounter) {
	stmt := node.(*ast.ProfessorTobuffStatement)
	index := stmt.Index
	if index > 6 {
		fmt.Printf("Not exist %dth Professor.", index)
		return
	}
	jisija := stmt.Type
	n, boolean := env.Professor.Response_Interview(index)
	if !boolean {
		fmt.Printf("Professor don't response your request")
		return
	}

	ppr := &environment.PrintCharacter{Node: n, PrintType: jisija}
	env.Buffer.SetinBuffer(ppr)
}

func evalPrintBuffer(env *environment.Environment, pc *ProgramCounter) {
	for {
		e := env.Buffer.GetinBuffer()
		if e == nil {
			break
		}
		v := e.Node.Value
		typer := e.PrintType
		switch typer {
		case environment.Char:
			fmt.Printf("%c", v)
		case environment.Int:
			fmt.Printf("%d", v)
		}
	}
}

func evalClass(node ast.Node, env *environment.Environment, pc *ProgramCounter) {
	n := node.(*ast.ClassStatement)
	switch n.Type {
	case "보강":
		evalPlusClass(env, pc)
	case "휴강":
		evalMinusClass(env, pc)
	}
}

func evalMinusClass(env *environment.Environment, pc *ProgramCounter) {
	pc.Index++
}

func evalPlusClass(env *environment.Environment, pc *ProgramCounter) {
	pc.ThisisPlused()
	pc.Index--
	pc.Index--
}