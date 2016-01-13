// Try out Go expression parser.
package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"io"
	"os"
	"strings"

	"github.com/kr/pretty"
)

const (
	callExpr   = "f(x+y, g(z), u...)"
	binaryExpr = "x+y"
	mixExpr    = "100+x+y*2+f(z)"
)

// mustRecover tries to recover a given input string from the parsed
// expression.
func humanString(astE ast.Expr, dump io.Writer) string {
	// TODO: Add specific processing for each expression type
	switch astE := astE.(type) {
	default:
		fmt.Fprintf(dump, "unexpected expression type %T", astE)
		return "<UNK>"
	case *ast.CallExpr:
		var args []string
		for _, arg := range astE.Args { // arg is of type ast.Expr
			args = append(args, humanString(arg, dump))
		}
		argStr := strings.Join(args, ", ")
		if astE.Ellipsis.IsValid() { // if there is variadic argument
			argStr += "..."
		}
		fmt.Fprintf(dump, "Expression type: %T\n", astE)
		return fmt.Sprintf("%v(%v)", astE.Fun, argStr)
	case *ast.BinaryExpr: // astE.X/Y is left/right operand; astE.Op is operator of type go/token.Token
		fmt.Fprintf(dump, "Expression type: %T\n", astE)
		return fmt.Sprintf("%v %v %v", humanString(astE.X, dump), astE.Op, humanString(astE.Y, dump))
	case *ast.Ident:
		fmt.Fprintf(dump, "Expression type: %T\n", astE)
		return fmt.Sprintf("%v", astE.Name)
	case *ast.BasicLit:
		fmt.Fprintf(dump, "Expression type: %T\n", astE)
		return fmt.Sprintf("%v", astE.Value)
	}
}

func haveFun(expr string, debug bool) {
	// Dump logging info to stdout.
	// TODO: Consider better dumping, e.g., to file.
	dump := os.Stdout
	astE, err := parser.ParseExpr(expr)
	if err != nil {
		fmt.Fprintln(dump, err)
	}
	if debug {
		pretty.Printf("%# v\n", astE)
		fmt.Printf("The pos range [%v, %v)\n", astE.Pos(), astE.End())
	}
	println(humanString(astE, dump))
}

func main() {
	haveFun(callExpr, false)
	haveFun(binaryExpr, false)
	haveFun(mixExpr, false)
}
