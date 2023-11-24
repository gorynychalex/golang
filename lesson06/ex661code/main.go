package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	fileSet := token.NewFileSet()

	node, err := parser.ParseFile(fileSet, "lesson06/ex621json/main.go",nil, parser.ParseComments)

	if err != nil {panic(err)}

	for _,c := range node.Imports { fmt.Println(c.Path) }


	ast.Inspect(node, func(n ast.Node) bool {
		ret, ok := n.(*ast.ReturnStmt)
		if ok {fmt.Printf("return statement fount on line %d:\n\t", fileSet.Position(ret.Pos()))}
		return true
	})


	// for _, decl := range node.Decls {
	// 	f,ok := decl.(*ast.FuncDecl)
	// 	if !ok && !f.Name.IsExported() { continue }
	// 	fmt.Println(f.Name.Name)
	// }
}