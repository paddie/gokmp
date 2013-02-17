gokmp
=====

String-matching in Golang using the Knuth–Morris–Pratt algorithm (KMP)

Example:
========
```Go
package main

import (
  "fmt"
	"github.com/paddie/gokmp"
)

const str = "aabaabaaaabbaabaabaaabbaabaabb"
const pattern = "aabb"

func main() {
	kmp, _ := gokmp.NewKMP(pattern)
	ints := kmp.FindAllStringIndex(str)

	fmt.Println(ints)
}
```
