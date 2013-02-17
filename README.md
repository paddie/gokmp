gokmp
=====

String-matching in Golang using the Knuth–Morris–Pratt algorithm (KMP)

See [Documentation](http://godoc.org/github.com/paddie/gokmp) on [http://godoc.org/](GoDoc).

Example:
========
```Go
package main

import (
  "fmt"
	"github.com/paddie/gokmp"
)

const str = "aabaabaaaabbaabaabaaabbaabaabb"
//          "        _          _      _   "
//                   8          19     26
const pattern = "aabb"

func main() {
	kmp, _ := gokmp.NewKMP(pattern)
	ints := kmp.FindAllStringIndex(str)

	fmt.Println(ints)
}
```
Output:
=======
```Go
[8 19 26]
```
