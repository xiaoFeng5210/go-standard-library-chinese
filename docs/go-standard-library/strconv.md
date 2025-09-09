## func AppendBool
```go{1}
func AppendBool(dst []byte, b bool) []byte
```
AppendBool 根据 b 的值将“true”或“false”附加到 dst 并返回扩展缓冲区。

#### 使用示例
```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	b := []byte("bool:")
	b = strconv.AppendBool(b, true)
	fmt.Println(string(b))

}

```
```text
bool:true
```

## func AppendFloat
```go{1}
func Atoi(s string) (int, error)
```
Atoi 等价于 ParseInt(s, 10, 0)，转换为 int 类型。

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	v := "10"
	if s, err := strconv.Atoi(v); err == nil {
		fmt.Printf("%T, %v", s, s)
	}

}
```
```text
int, 10
```

## func Atoi

```go{1}
func Atoi(s string) (int, error)
```
Atoi 等价于 ParseInt(s, 10, 0)，转换为 int 类型。字符串转数字

#### 使用示例
```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	v := "10"
	if s, err := strconv.Atoi(v); err == nil {
		fmt.Printf("%T, %v", s, s)
	}

}
```
```text
int, 10
```

## func Itoa
```go{1}
func Itoa(i int) string
```
Itoa 等价于 FormatInt(int64(i), 10)。

#### 使用示例
```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	i := 10
	s := strconv.Itoa(i)
	fmt.Printf("%T, %v\n", s, s)

}

```
```text
string, 10
```
