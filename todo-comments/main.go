package main

import (
	"flag"
	"fmt"
	"go/parser"
	"go/token"
	"log"
	"strings"
)

func main() {
	flag.Parse()
	fset := token.NewFileSet()
	for _, file := range flag.Args() {
		f, err := parser.ParseFile(fset, file, nil, parser.ParseComments)
		if err != nil {
			log.Fatal(err)
		}
		for _, cmts := range f.Comments {
			for _, cmt := range cmts.List {
				t := strings.TrimSpace(cmt.Text)
				if strings.HasPrefix(t, "// TODO") {
					fmt.Println(fset.Position(cmt.Slash).String(), t)
					// TODO test
				}
			}
		}
	}
	// TODO Надо улучшить
}
