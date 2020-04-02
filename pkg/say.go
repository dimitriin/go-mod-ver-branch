package pkg

import "golang.org/x/tools/go/ssa/interp/testdata/src/fmt"

func Say(word string) {
	fmt.Printf("[v3] Say %s", word)
}
