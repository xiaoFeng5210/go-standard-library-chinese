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

## func ContainsFunc
```go
func ContainsFunc(s string, f func(rune) bool) bool
```
ContainsFunc 报告 s 内的任何 Unicode 代码点 r 是否满足 f(r)。

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	f := func(r rune) bool {
		return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u'
	}
	fmt.Println(strings.ContainsFunc("hello", f))
	fmt.Println(strings.ContainsFunc("rhythms", f))
}
```
```text
true
false
```


## func Count
```go
func Count(s, substr string) int
```
Count 计算 s 中不重叠的 substr 实例的数量。如果 substr 为空字符串，则 Count 返回 1 + s 中的 Unicode 码位数量。

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Count("cheese", "e"))
	fmt.Println(strings.Count("five", "")) // before & after each rune
}
```
```text
3
5
```


## func Cut
```go
func Cut(s, sep string) (before, after string, found bool)
```
围绕第一个 sep 实例对 s 进行切片，返回 sep 前后的文本。结果报告 sep 是否出现在 s 中。如果 sep 未出现在 s 中，则 cut 返回 s, "", false。

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	show := func(s, sep string) {
		before, after, found := strings.Cut(s, sep)
		fmt.Printf("Cut(%q, %q) = %q, %q, %v\n", s, sep, before, after, found)
	}
	show("Gopher", "Go")
	show("Gopher", "ph")
	show("Gopher", "er")
	show("Gopher", "Badger")
}
```

```text
Cut("Gopher", "Go") = "", "pher", true
Cut("Gopher", "ph") = "Go", "er", true
Cut("Gopher", "er") = "Goph", "", true
Cut("Gopher", "Badger") = "Gopher", "", false
```

## func HasPrefix
```go
func HasPrefix(s, prefix string) bool
```
HasPrefix 报告字符串 s 是否以前缀开头。

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.HasPrefix("Gopher", "Go"))
	fmt.Println(strings.HasPrefix("Gopher", "C"))
	fmt.Println(strings.HasPrefix("Gopher", ""))
}
```

## func Index
```go
func Index(s, substr string) int
```
Index 返回 s 中 substr 的第一个实例的索引，如果 s 中不存在 substr，则返回 -1。

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Index("chicken", "ken"))
	fmt.Println(strings.Index("chicken", "dmr"))
}
```
```text
4
-1
```


## func Join
```go
func Join(elems []string, sep string) string
```
Join 将其第一个参数的元素连接起来以创建一个字符串。分隔符字符串 sep 位于结果字符串的元素之间。


## func Split
```go
func Split(s, sep string) []string
```
- 将切片 s 拆分为由 sep 分隔的所有子字符串，并返回这些分隔符之间的子字符串切片。
- 如果 s 不包含 sep 且 sep 非空，则 Split 返回长度为 1 且唯一元素为 s 的切片。
- 如果 sep 为空，Split 会在每个 UTF-8 序列后进行拆分。如果 s 和 sep 均为空，Split 会返回一个空切片。
- 它相当于计数为 -1 的 SplitN 。

