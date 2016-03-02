// +build ignore

package main

import (
	"fmt"
	"os"

	"github.com/awoodbeck/acrostic"
	"github.com/shurcooL/vfsgen"
)

func main() {
	err := vfsgen.Generate(acrostic.Assets, vfsgen.Options{
		BuildTags:    "!dev",
		PackageName:  "acrostic",
		VariableName: "Assets",
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
