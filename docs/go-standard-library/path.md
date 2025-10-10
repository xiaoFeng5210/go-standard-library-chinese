# Path

path 包实现了用于操作和解析文件路径的函数。


## func Base
```go
func Base(path string) string
```

Base 返回路径的最后一个元素。提取最后一个元素之前，会删除路径尾部的斜杠。如果路径为空，Base 返回“.”；如果路径完全由斜杠组成，Base 返回“/”。

```go
package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println(path.Base("/a/b"))
	fmt.Println(path.Base("/"))
	fmt.Println(path.Base(""))
}
```

```text
b
/
.
```
