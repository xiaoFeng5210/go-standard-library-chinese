# strings

包 strings 实现了用于操作字符串的简单函数。

## func Clone
```go
func Clone(s string) string
```
Clone 返回 s 的全新副本。它保证将 s 的副本复制到新的分配空间中，这在仅保留较大字符串的一小部分子字符串时非常重要。使用 Clone 可以帮助此类程序减少内存消耗。当然，由于使用 Clone 会进行复制，过度使用 Clone 会使程序占用更多内存。Clone 通常很少使用，并且仅在性能分析表明需要时才使用。对于长度为零的字符串，将返回字符串 "" 并且不进行任何分配。

```go
package main

import (
	"fmt"
	"strings"
	"unsafe"
)

func main() {
	s := "abc"
	clone := strings.Clone(s)
	fmt.Println(s == clone)
	fmt.Println(unsafe.StringData(s) == unsafe.StringData(clone))
}

```
```text
true
false
```


## func Compare
```go
func Compare(a, b string) int
```
Compare 比较两个字符串的大小。它返回一个整数，表示 a 和 b 的比较结果。如果 a 小于 b，则返回 -1；如果 a 等于 b，则返回 0；如果 a 大于 b，则返回 +1。
当你需要进行三向比较时使用 Compare（使用 例如， slices.SortFunc ）。使用内置字符串比较运算符 ==、<、> 等通常更清晰，而且速度更快。
```go
package main

import (
	"fmt"
	"strings"
)
```

## func Contains
```go
func Contains(s, substr string) bool
```
包含报告 substr 是否在 s 内。

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("seafood", "foo"))
	fmt.Println(strings.Contains("seafood", "bar"))
	fmt.Println(strings.Contains("seafood", ""))
	fmt.Println(strings.Contains("", ""))
}

```
```text
true
false
true
true
```
