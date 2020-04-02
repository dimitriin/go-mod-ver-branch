package pkg

import "golang.org/x/tools/go/ssa/interp/testdata/src/fmt"

func Say(word string) {
	fmt.Printf("[v1] Say %s", word)
}
