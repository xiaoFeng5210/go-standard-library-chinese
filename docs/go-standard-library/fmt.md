---
layout: doc
title: fmt
---

# fmt

Fmt 包实现了格式化 I/O, 其函数类似于 C 的 printf 和 scanf。格式 “动词” 源于 C, 但更简单。

## 函数
> fmt包的所有函数

### func Append

使用其操作数的默认格式追加格式，将结果追加到字节切片中，并返回更新后的切片。

```go{1}
func Append(b []byte, a ...any) []byte
```

#### 使用示例
```go
package main

import "fmt"

func main() {
  b := []byte("Hello, ")
  b = fmt.Append(b, "World!")
  fmt.Println(string(b))
}
```

### func Appendf
Appendf 根据格式说明符进行格式化，将结果附加到字节切片，并返回更新后的切片。

```go{1}
func Appendf(b []byte, format string, a ...any) []byte
```

#### 使用示例
```go
package main

import "fmt"

func main() {
  b := []byte("Hello, ")
  b = fmt.Appendf(b, "World! %d", 123)
  fmt.Println(string(b))
}
```

### func Appendln

Appendln 使用其操作数的默认格式进行格式化，将结果附加到字节切片，并返回更新后的切片。每次append操作之后，始终添加空格，并附加换行符。

```go{1}
func Appendln(b []byte, a ...any) []byte
```

#### 使用示例
```go
package main

import "fmt"

func main() {
  b := []byte("Hello, ")
  b = fmt.Appendln(b, "World!")
  fmt.Println(string(b))
}
```

### func Errorf
Errorf 根据格式说明符进行格式化，并返回满足错误条件的字符串值。

如果格式说明符包含一个带有错误操作数的 %w 动词，则返回的错误将实现一个 Unwrap 方法，返回该操作数。如果有多个 %w 动词，则返回的错误将实现一个 Unwrap 方法，返回一个 []error，其中包含所有 %w 操作数，这些操作数按它们在参数中出现的顺序排列。为 %w 动词提供未实现错误接口的操作数是无效的。否则，%w 动词与 %v 同义。

```go{1}
func Errorf(format string, a ...any) error
```

#### 使用示例
```go
package main

import (
	"fmt"
)

func main() {
	const name, id = "bueller", 17
	err := fmt.Errorf("user %q (id %d) not found", name, id)
	fmt.Println(err.Error())

}

user "bueller" (id 17) not found
```

### func FormatString
FormatString 返回一个字符串，该字符串表示由 State 捕获的完全限定格式指令，后跟参数动词。（State 本身不包含动词。）结果以百分号为前导，后跟所有标志、宽度和精度。缺失的标志、宽度和精度将被省略。此函数允许 Formatter 重建触发 Format 调用的原始指令。

```go{1}
func FormatString(state State, verb rune) string
```

#### 使用示例
```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println(fmt.FormatString(fmt.State{}, 'v'))
}
```

### func Fprint
Fprint 使用其操作数的默认格式进行格式化，并写入 w。当操作数都不是字符串时，会在操作数之间添加空格。它返回写入的字节数以及遇到的任何写入错误。

```go{1}
func Fprint(w io.Writer, a ...any) (n int, err error)
```

#### 使用示例
```go
import (
	"fmt"
	"os"
)

func main() {
	const name, age = "Kim", 22
	n, err := fmt.Fprint(os.Stdout, name, " is ", age, " years old.\n")

	// The n and err return values from Fprint are
	// those returned by the underlying io.Writer.
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fprint: %v\n", err)
	}
	fmt.Print(n, " bytes written.\n")
}

Kim is 22 years old.
21 bytes written.
```

### func Sprintf
Sprintf 根据格式说明符进行格式化并返回结果字符串。

#### 使用示例
```go
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	const name, age = "Kim", 22
	s := fmt.Sprintf("%s is %d years old.\n", name, age)

	io.WriteString(os.Stdout, s) // Ignoring error for simplicity.

}

Kim is 22 years old.
```


