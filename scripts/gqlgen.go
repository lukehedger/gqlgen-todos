// +build ignore

package main

import (
	"fmt"
	"os"

	"github.com/99designs/gqlgen/cmd"
)

func main() {
	// Remove `tmp/resolver.go` file before generating GraphQL resolvers
	err := os.Remove("tmp/resolver.go")
	if err != nil {
		fmt.Println(err)
		return
	}

	cmd.Execute()
}
