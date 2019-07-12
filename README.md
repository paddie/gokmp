gokmp
=====

String-matching in Golang using the Knuth–Morris–Pratt algorithm (KMP).

## Disclaimer

This library was written as part of my Master's Thesis and should be used as a helpful implementation reference for people interested in the Knuth-Morris-Pratt algorithm than as a performance string searching library.

I believe the compiler has since caught up to most of the gains that this library bought me back in the day.

See [Documentation](http://godoc.org/github.com/paddie/gokmp) on [GoDoc](http://godoc.org/).

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

Tests and Benchmarks:
=====================
```
go test -v -bench .
```

Output:
=======
```
=== RUN TestFindAllStringIndex
--- PASS: TestFindAllStringIndex (0.00 seconds)
=== RUN TestFindStringIndex
--- PASS: TestFindStringIndex (0.00 seconds)
=== RUN TestContainedIn
--- PASS: TestContainedIn (0.00 seconds)
=== RUN TestOccurrences
--- PASS: TestOccurrences (0.00 seconds)
=== RUN TestOccurrencesFail
--- PASS: TestOccurrencesFail (0.00 seconds)
PASS
BenchmarkKMPIndexComparison	10000000	       178 ns/op
BenchmarkStringsIndexComparison	10000000	       359 ns/op
ok  	github.com/paddie/gokmp	5.854s
```
Comparison:
============
```bash
gokmp.FindStringIndex / strings.Index = 178/359 = 0.4958
```
Almost a 2x improvement over the naive built-in method.
