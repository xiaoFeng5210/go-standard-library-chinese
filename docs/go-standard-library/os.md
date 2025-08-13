# OS

os 软件包提供了一个与平台无关的操作系统功能接口。其设计类似于 Unix，但错误处理类似于 Go；失败的调用返回错误类型值，而不是错误号。错误信息中通常会包含更多信息。例如，如果一个接受文件名的调用（例如 Open 或 Stat）失败，则错误在打印时会包含失败的文件名，并且类型为 *PathError，可以通过解包获取更多信息。

os 接口旨在在所有操作系统上保持一致。一些并非普遍可用的功能包含在特定于系统的 syscall 软件包中。

## func Chdir
```go{1}
func Chdir(dir string) error
```
CHDIR将当前的工作目录更改为指定目录。如果有错误，它将是类型 *Patherror。


#### 使用示例
```go
package main

import "os"

func main() {
  os.Chdir("/tmp")
}
```


## func Chmod
```go{1}
func Chmod(name string, mode FileMode) error
```
CHMOD将命名文件的模式更改为模式。如果文件是符号链接，它将更改链接目标的模式。如果有错误，它将是类型 *Patherror。

根据操作系统，使用了模式位的不同子集。

在UNIX上，使用了模式的权限位，模量，模态和iMesticky。

在Windows上，仅使用0o200位（所有者写）模式。它控制文件的仅读取属性是设置还是清除。其他位目前未使用。为了与GO 1.12及更早的兼容性，请使用非零模式。将模式0O400用于仅读取文件，而0O600用于可读+可写的文件。

在计划9上，使用该模式的权限位，模型ppend，modeexclusive和ModeTemporary。

#### 使用示例
```go
package main

import (
	"log"
	"os"
)

func main() {
	if err := os.Chmod("some-filename", 0644); err != nil {
		log.Fatal(err)
	}
}
```

## func Chtimes
```go{1}
func Chtimes(name string, atime time.Time, mtime time.Time) error
```
CHTIMES更改命名文件的访问和修改时间，类似于Unix UTIME（）或UTIMES（）函数。零时间。时间值将使相应的文件时间保持不变。

基础文件系统可能会截断或将值围绕到较不精确的时间单元。如果有错误，它将是类型 *Patherror。

#### 使用示例
```go
package main

import (
	"log"
	"os"
	"time"
)

func main() {
	mtime := time.Date(2006, time.February, 1, 3, 4, 5, 0, time.UTC)
	atime := time.Date(2007, time.March, 2, 4, 5, 6, 0, time.UTC)
	if err := os.Chtimes("some-filename", atime, mtime); err != nil {
		log.Fatal(err)
	}
}
```

## func Clearenv
```go{1}
func Clearenv()
```
Clearenv删除所有环境变量。


## func CopyFS
```go{1}
func CopyFS(dir string, fsys fs.FS) error
```
COPYFS将文件系统FSYS复制到目录DIR中，如有必要，创建DIR。

使用模式0O666加上来自源的任何执行权限创建文件，并使用模式0O777（在UMASK之前）创建目录。

COPYFS不会覆盖现有文件。如果FSYS中的文件名已经存在于目标中，则COPYFS将返回错误，以便errors.is（err，fs.errexist）为true。

不支持FSY中的符号链接。从符号链接复制时，将返回具有ERR设置为Errinvalid的Patherror。

遵循DIR中的符号链接。

在运行copyFS时，添加到FSYS（包括DIR是FSYS的子目录）中添加的新文件。

复制停止并返回遇到的第一个错误。

#### 使用示例
```go
package main

import (
	"log"
	"os"
)

func main() {
	if err := os.CopyFS("some-dir", os.DirFS("some-other-dir")); err != nil {
		log.Fatal(err)
	}
}
```

## func DirFS
```go{1}
func DirFS(dir string) fs.FS
```
DirFS 返回一个文件系统 (fs.FS), 用于根目录 dir 的文件树。
请注意，DirFS (“/prefix”) 仅保证它对操作系统的 Open 调用将以 “/prefix” 开头：DirFS (“/prefix”)。Open (“file”) 与 os.Open (“/prefix/file”) 相同。因此，如果 /prefix/file 是指向 /prefix 树外部的符号链接，那么使用 DirFS 并不比使用 os.Open 更能阻止访问。此外，为相对路径返回的 fs.FS 的根目录 DirFS (“prefix”) 将受到后续对 Chdir 的调用的影响。因此，当目录树包含任意内容时，DirFS 并不是 chroot 风格安全机制的通用替代品。
使用 Root.FS 获取一个 fs.FS, 它可以防止通过符号链接从树中逃逸。
目录 dir 不能是 “.”
结果实现了 io/fs.StatFS、io/fs.ReadFileFS 和 io/fs.ReadDirFS。

#### 使用示例
```go
package main
```

## func Environ
```go{1}
func Environ() []string
```
Environment 返回一个表示环境的字符串副本，形式为 “key=value”。

#### 使用示例
```go
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	b := os.Environ()
	for _, s := range b {
		parts := strings.Split(s, "=")
		if parts[0] == "USER" {
			fmt.Println(parts[1])
		}
	}
}
```

## func Executable
```go{1}
func Executable() (string, error)
```
Executable 返回启动当前进程的可执行文件的路径名称。无法保证该路径仍然指向正确的可执行文件。如果使用符号链接启动进程，根据操作系统的不同，结果可能是符号链接或其指向的路径。如果需要稳定的结果，path/filepath.EvalSymlinks 可能会有帮助。
Executable 返回绝对路径，除非出现错误。
主要用例是查找相对于可执行文件位置的资源。

## func Exit
```go{1}
func Exit(code int)
```
退出会导致当前程序使用给定的状态码退出。通常，代码 0 表示成功，非零表示错误。程序立即终止；延迟函数不会运行。
为了便携性，状态码应在 [0，125] 范围内。

## func Expand
```go{1}
func Expand(s string, mapping func(string) string) string
```
Expand 根据映射函数替换字符串中的 ${var} 或 $var。例如，os.ExpandEnv (s) 等价于 os.Expand (s，os.Getenv)。

#### 使用示例
```go
package main

import (
	"fmt"
	"os"
)

func main() {
	mapper := func(placeholderName string) string {
		switch placeholderName {
		case "DAY_PART":
			return "morning"
		case "NAME":
			return "Gopher"
		}

		return ""
	}

	fmt.Println(os.Expand("Good ${DAY_PART}, $NAME!", mapper))

	// Output:
	// Good morning, Gopher!
}
```

## func ExpandEnv
```go{1}
func ExpandEnv(s string) string
```
ExpandEnv 将根据当前环境变量的值替换字符串中的 ${var} 或 $var。对未定义变量的引用将替换为空字符串。

#### 使用示例
```go
package main

import (
	"fmt"
	"os"
)

func main() {
	os.Setenv("NAME", "gopher")
	os.Setenv("BURROW", "/usr/gopher")

	fmt.Println(os.ExpandEnv("$NAME lives in ${BURROW}."))

}
// Output:
// gopher lives in /usr/gopher.
```


## func Getegid
```go{1}
func Getegid() int
```
Getegid 返回调用进程的有效组 ID。
在 Windows 上，它返回 - 1。

## func Getenv
```go{1}
func Getenv(key string) string
```
Getenv 检索由键命名的环境变量的值。它返回该值，如果该变量不存在，则该值将为空。要区分空值和未设置的值，请使用 LookupEnv。

#### 使用示例
```go
package main

import (
	"fmt"
	"os"
)

func main() {
	os.Setenv("NAME", "gopher")
	os.Setenv("BURROW", "/usr/gopher")

	fmt.Printf("%s lives in %s.\n", os.Getenv("NAME"), os.Getenv("BURROW"))

}

// Output:
// gopher lives in /usr/gopher.
```

## func Geteuid
```go{1}
func Geteuid() int
```
Geteuid 返回呼叫者的数字有效用户 ID

## func Getgid
```go{1}
func Getgid() int
```
Getgid 返回呼叫者的数字组 ID。
在 Windows 上，它返回 - 1。

## func Getpid
```go{1}
func Getpid() int
```
Getpid 返回调用进程的进程 ID。
在 Windows 上，它返回 - 1。

## func Getppid
```go{1}
func Getppid() int
```
Getppid 返回调用进程的父进程 ID。
在 Windows 上，它返回 - 1。

## func Getuid
```go{1}
func Getuid() int
```
Getuid 返回呼叫者的数字用户 ID。
在 Windows 上，它返回 - 1。


## func Getwd
```go{1}
func Getwd() (dir string, err error)
```
Getwd 返回当前目录对应的绝对路径名。如果当前目录可以通过多条路径（例如符号链接）到达，Getwd 可能会返回其中任意一条路径。
在 Unix 平台上，如果环境变量 PWD 提供了一个绝对名称，并且它是当前目录的名称，则返回该名称。

## func Hostname
```go{1}
func Hostname() (name string, err error)
```
Hostname 返回调用进程的系统主机名。
如果主机名无法确定，则返回错误。
在 Windows 上，如果主机名无法确定，则返回错误。

## func IsExist
```go{1}
func IsExist(err error) bool
```
IsExist 报告是否存在错误，如果错误是类型 *Patherror，则报告错误是否表示文件或目录存在。
其实这里就是专门判断文件或目录是否存在使用。
:::warning
此函数早于 error.Is 。它仅支持 os 包返回的错误。新代码应使用 error.Is(err, fs.ErrExist)。
:::

## func IsNotExist
```go{1}
func IsNotExist(err error) bool
```
IsNotExist 返回一个布尔值，指示其参数是否已知 报告文件或目录不存在。满足条件的是 ErrNotExist 以及一些系统调用错误。
:::warning
此函数早于 error.Is 。它仅支持 os 包返回的错误。新代码应使用 error.Is(err, fs.ErrNotExist)。
:::








