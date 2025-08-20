---
layout: doc
title: os
---

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

## func IsPathSeparator
```go{1}
func IsPathSeparator(c uint8) bool
```
IsPathSeparator 报告 参数c 是否是目录分隔符。
在 Windows 上，IsPathSeparator 返回 true 当且仅当 c 是 0x2F 或 0x5C。
在 Unix 上，IsPathSeparator 返回 true 当且仅当 c 是 0x2F。

## func IsPermission
```go{1}
func IsPermission(err error) bool
```
IsPermission 返回一个布尔值，指示其参数是否已知，以报告权限被拒绝。它满足 ErrPermission 以及一些系统调用错误。
:::warning
此函数先于错误。Is。它只支持 os 包返回的错误。新代码应使用错误。Is (err，fs.ErrPermission)。
:::

## func Mkdir
```go{1}
func Mkdir(name string, perm FileMode) error
```
Mkdir 创建一个新目录，其中包含指定的名称和权限位 (在 umask 之前)。如果出现错误，将是类型错误。

#### 使用示例
```go
package main

import (
	"log"
	"os"
)

func main() {
	err := os.Mkdir("testdir", 0750)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}
	err = os.WriteFile("testdir/testfile.txt", []byte("Hello, Gophers!"), 0660)
	if err != nil {
		log.Fatal(err)
	}
}
```

## func MkdirAll
```go{1}
func MkdirAll(path string, perm FileMode) error
```
MkdirAll 创建一个名为 path 的目录，以及任何必要的父目录，并返回 nil, 否则返回错误。权限位 perm (在 umask 之前) 用于 MkdirAll 创建的所有目录。如果 path 已经是一个目录，MkdirAll 什么也不做，并返回 nil。
#### 使用示例
```go
package main

import (
	"log"
	"os"
)

func main() {
	err := os.MkdirAll("test/subdir", 0750)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile("test/subdir/testfile.txt", []byte("Hello, Gophers!"), 0660)
	if err != nil {
		log.Fatal(err)
	}
}
```

## func MkdirTemp
```go{1}
func MkdirTemp(dir, pattern string) (string, error)
```
MkdirTemp 在目录 dir 中创建一个新的临时目录，并返回新目录的路径名。新目录的名称是通过在 pattern 的末尾添加随机字符串来生成的。如果 pattern 包含 “*”, 随机字符串将替换最后一个 “*”。目录是在 0o700 模式下创建的 (在 umask 之前)。如果 dir 是空字符串，MkdirTemp 将使用 TempDir 返回的默认目录来存放临时文件。同时调用 MkdirTemp 的多个程序或 goroutine 不会选择相同的目录。当不再需要该目录时，调用方有责任删除它。

#### 使用示例
```go
package main

import (
	"log"
	"os"
	"path/filepath"
)

func main() {
	dir, err := os.MkdirTemp("", "example")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(dir) // clean up

	file := filepath.Join(dir, "tmpfile")
	if err := os.WriteFile(file, []byte("content"), 0666); err != nil {
		log.Fatal(err)
	}
}
```

#### 使用示例2
```go
package main

import (
	"log"
	"os"
	"path/filepath"
)

func main() {
	logsDir, err := os.MkdirTemp("", "*-logs")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(logsDir) // clean up

	// Logs can be cleaned out earlier if needed by searching
	// for all directories whose suffix ends in *-logs.
	globPattern := filepath.Join(os.TempDir(), "*-logs")
	matches, err := filepath.Glob(globPattern)
	if err != nil {
		log.Fatalf("Failed to match %q: %v", globPattern, err)
	}

	for _, match := range matches {
		if err := os.RemoveAll(match); err != nil {
			log.Printf("Failed to remove %q: %v", match, err)
		}
	}
}

```


## func Pipe
```go{1}
func Pipe() (r *File, w *File, err error)
```
管道返回一对连接的 Files; 从 r 读取返回写入 w 的字节。它返回文件和一个错误 (如果有的话)。


## func ReadFile <Badge text="重要" />
```go{1}
func ReadFile(name string) ([]byte, error)
```
ReadFile 读取命名文件并返回其内容。成功的调用返回 err == nil, 而不是 err == EOF。由于 ReadFile 读取整个文件，它不会将来自 Read 的 EOF 视为要报告的错误。

#### 使用示例
```go
package main

import (
	"log"
	"os"
)

func main() {
	data, err := os.ReadFile("testdata/hello")
	if err != nil {
		log.Fatal(err)
	}
	os.Stdout.Write(data)

}
```

```go
Hello, Gophers!
```

## func Readlink
```go{1}
func Readlink(name string) (string, error)
```
Readlink 返回命名符号链接的目的地。如果出现错误，则类型为 * PathError。如果链接目的地是相对的，Readlink 将返回相对路径，而不会将其解析为绝对路径。

#### 使用示例
```go
package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// First, we create a relative symlink to a file.
	d, err := os.MkdirTemp("", "")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(d)
	targetPath := filepath.Join(d, "hello.txt")
	if err := os.WriteFile(targetPath, []byte("Hello, Gophers!"), 0644); err != nil {
		log.Fatal(err)
	}
	linkPath := filepath.Join(d, "hello.link")
	if err := os.Symlink("hello.txt", filepath.Join(d, "hello.link")); err != nil {
		if errors.Is(err, errors.ErrUnsupported) {
			// Allow the example to run on platforms that do not support symbolic links.
			fmt.Printf("%s links to %s\n", filepath.Base(linkPath), "hello.txt")
			return
		}
		log.Fatal(err)
	}

	// Readlink returns the relative path as passed to os.Symlink.
	dst, err := os.Readlink(linkPath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s links to %s\n", filepath.Base(linkPath), dst)

	var dstAbs string
	if filepath.IsAbs(dst) {
		dstAbs = dst
	} else {
		// Symlink targets are relative to the directory containing the link.
		dstAbs = filepath.Join(filepath.Dir(linkPath), dst)
	}

	// Check that the target is correct by comparing it with os.Stat
	// on the original target path.
	dstInfo, err := os.Stat(dstAbs)
	if err != nil {
		log.Fatal(err)
	}
	targetInfo, err := os.Stat(targetPath)
	if err != nil {
		log.Fatal(err)
	}
	if !os.SameFile(dstInfo, targetInfo) {
		log.Fatalf("link destination (%s) is not the same file as %s", dstAbs, targetPath)
	}

}
```

```go
hello.link links to hello.txt
```

## func Remove  <Badge text="重要" />
```go{1}
func Remove(name string) error
```
Remove 会删除命名文件或 (空) 目录。如果出现错误，则类型为 * PathError。

#### 使用示例
```go
package main

import (
	"log"
	"os"
)

func main() {
	err := os.Remove("testfile")
	if err != nil {
		log.Fatal(err)
	}
}
```

## func RemoveAll <Badge text="重要" />
```go{1}
func RemoveAll(path string) error
```
RemoveAll 删除 path 及其所有子目录。它递归地删除所有非符号链接文件。如果出现错误，则类型为 * PathError。
#### 使用示例
```go
package main

import (
	"log"
	"os"
)

func main() {
	err := os.RemoveAll(path)
	if err != nil {
		log.Fatal(err)
	}
}
```

## func Rename
```go{1}
func Rename(oldpath, newpath string) error
```
Rename 将 oldpath 重命名 (移动) 到 newpath。**如果 newpath 已经存在且不是目录，Rename 会替换它。如果 newpath 已经存在且是目录，Rename 会返回一个错误**。当 oldpath 和 newpath 位于不同的目录时，可能会适用特定于操作系统的限制。即使在同一目录内，在非 Unix 平台上，Rename 也不是一个原子操作。如果出现错误，它将是 * LinkError 类型。


## func SameFile
```go{1}
func SameFile(fi1, fi2 FileInfo) bool
```
SameFile 报告 fi1 和 fi2 是否描述的是同一个文件。例如，在 Unix 上，这意味着两个底层结构的设备和 inode 字段相同；在其他系统上，则可能基于路径名来判断。SameFile 仅适用于此包的 Stat 返回的结果。在其他情况下，它会返回 false。

## func Setenv
```go{1}
func Setenv(key, value string) error
```
Setenv 设置由键指定的环境变量的值。如果出现错误，则返回错误。

## func Unsetenv
```go{1}
func Unsetenv(key string) error
```
Unsetenv 取消设置单个环境变量。

#### 使用示例
```go
package main

import (
	"os"
)

func main() {
	os.Setenv("TMPDIR", "/my/tmp")
	defer os.Unsetenv("TMPDIR")
}

```

## func Symlink
```go{1}
func Symlink(oldname, newname string) error
```
符号链接会将 newname 创建为指向 oldname 的符号链接。在 Windows 上，指向不存在的 oldname 的符号链接会创建文件符号链接；如果 oldname 之后被创建为目录，则该符号链接将不起作用。如果发生错误，错误类型为 *LinkError。

## func WriteFile <Badge text="重要" />
```go{1}
func WriteFile(name string, data []byte, perm FileMode) error
```
WriteFile 将数据写入指定文件，并在必要时创建该文件。如果文件不存在，WriteFile 将使用 perm 权限（umask 之前）创建该文件；否则，WriteFile 会在写入前截断文件，且不更改权限。由于 WriteFile 需要多个系统调用才能完成，因此操作过程中的故障可能会导致文件处于部分写入状态。

#### 使用示例
```go
package main

import (
	"log"
	"os"
)

func main() {
	err := os.WriteFile("testdata/hello", []byte("Hello, Gophers!"), 0666)
	if err != nil {
		log.Fatal(err)
	}
}
```

## func ReadDir
```go{1}
func ReadDir(name string) ([]DirEntry, error)
```
ReadDir 读取指定目录，并返回按文件名排序的所有目录条目。如果读取目录时发生错误，ReadDir 将返回错误发生前能够读取的条目以及错误本身。
#### 使用示例
```go
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}
```


## func Create
```go{1}
func Create(name string) (*File, error)
```
Create 函数用于创建或截断指定的文件。如果文件已存在，则截断该文件。如果文件不存在，则使用 0o666 模式（umask 之前）创建该文件。如果成功，则返回的 File 上的方法可用于 I/O；关联的文件描述符的模式为 O_RDWR 。包含该文件的目录必须已存在。如果发生错误，错误类型为 *PathError 。

#### 使用示例
```go
package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.Create("testfile")
	if err != nil {
		log.Fatal(err)
	}
}
```

## func CreateTemp <Badge text="added in go1.16" type="tip" />
```go{1}
func CreateTemp(dir, pattern string) (*File, error)
```
CreateTemp 在目录 dir 中创建一个新的临时文件，打开该文件进行读写操作，并返回结果文件。文件名由模式 pattern 和在末尾添加随机字符串生成。如果模式 pattern 包含“*”，则随机字符串将替换最后一个“*”。文件创建时采用模式 0o600（umask 之前）。如果 dir 为空字符串，CreateTemp 将使用 TempDir 返回的默认临时文件目录。多个程序或 goroutine 同时调用 CreateTemp 不会选择同一个文件。调用者可以使用文件的 Name 方法查找文件的路径名。当文件不再需要时，调用者有责任删除该文件。
#### 使用示例
```go
package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.CreateTemp("", "example")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(f.Name()) // clean up

	if _, err := f.Write([]byte("content")); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
```

## func NewFile
```go{1}
func NewFile(fd uintptr, name string) *File
```
NewFile 返回一个新文件，其中包含给定的文件描述符和名称。如果 fd 不是有效的文件描述符，则返回值为 nil。
NewFile 的行为在某些平台上有所不同：
在 Unix 上，如果 fd 处于非阻塞模式，NewFile 将尝试返回一个可轮询文件。
在 Windows 上，如果 fd 为异步 I/O 打开 (即，在 syscall.CreateFile 调用中指定了 syscall.FILE_FLAG_OVERLAPPED),NewFile 将尝试通过将 fd 与 Go 运行时 I/O 完成端口关联来返回一个可轮询文件。如果关联失败，I/O 操作将同步执行。
只有可轮询文件支持 File.SetDeadline、File.SetReadDeadline 和 File.SetWriteDeadline。
将其传递给 NewFile 后，在 File.Fd 的注释中描述的相同条件下，fd 可能会变得无效，并适用相同的约束。


## func Open
```go{1}
func Open(name string) (*File, error)
```
Open 打开指定文件进行读取。如果成功，则可以使用返回文件上的方法进行读取；关联的文件描述符的模式为 O_RDONLY 。如果发生错误，则返回 *PathError 类型错误。


## func OpenFile
```go{1}
func OpenFile(name string, flag int, perm FileMode) (*File, error)
```
OpenFile 是通用的打开调用；大多数用户会改用 Open 或 Create。它使用指定的标志（例如 O_RDONLY ）打开指定的文件。如果文件不存在，且传递了 O_CREATE 标志，则以 perm 模式（在 umask 之前）创建该文件；文件所在的目录必须存在。如果成功，则返回的 File 上的方法可用于 I/O。如果发生错误，则返回 *PathError 类型的错误。
#### 使用示例1
```go
package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.OpenFile("notes.txt", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
```

```go
package main

import (
	"log"
	"os"
)

func main() {
	// If the file doesn't exist, create it, or append to the file
	f, err := os.OpenFile("access.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write([]byte("appended some data\n")); err != nil {
		f.Close() // ignore error; Write error takes precedence
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
```


## func OpenInRoot <Badge text="added in go1.24.0" type="tip" />
```go{1}
func OpenInRoot(dir, name string) (*File, error)
```
OpenInRoot 在 dir 目录中打开文件名。它等效于 OpenRoot (dir), 后跟在根目录中打开文件。
如果名称的任何组件引用 dir 以外的位置，OpenInRoot 将返回一个错误。
有关详细信息和限制，请参阅根目录。

## Type File


### func (*File) Chdir 
```go{1}
func (f *File) Chdir() error
```
Chdir 将当前工作目录更改为 f 指向的目录。如果 f 是 nil，则 Chdir 将返回一个错误。

### func (*File) Chmod
```go{1}
func (f *File) Chmod(mode FileMode) error
```
chmod 将文件的模式更改为 mode。如果发生错误，则错误类型为 *PathError 。

### func (*File) Chown
```go{1}
func (f *File) Chown(uid, gid int) error
```
Chown 将文件的所有者 (uid) 和组 (gid) 更改为 uid 和 gid。如果发生错误，则错误类型为 *PathError 。

### func (*File) Close
```go{1}
func (f *File) Close() error
```
Close 会关闭 File ，使其无法进行 I/O 操作。对于支持 File.SetDeadline 的文件，任何待处理的 I/O 操作都将被取消并立即返回 ErrClosed 错误。如果 Close 已被调用，则会返回错误。

### func (*File) Fd
```go{1}
func (f *File) Fd() uintptr
```
Fd 返回系统文件描述符或引用打开文件的句柄。如果 f 被关闭，描述符将变得无效。如果 f 被垃圾回收，终结器可能会关闭描述符，使其无效；有关何时可能运行终结器的详细信息，请参阅运行时.SetFinalizer。
不要关闭返回的描述符，这可能会导致 f 稍后关闭一个不相关的描述符。
Fd 的行为在某些平台上有所不同：
在 Unix 和 Windows 上，File.SetDeadline 方法将停止工作。
在 Windows 上，如果文件上没有并发 I/O 操作，文件描述符将与 Go 运行时 I/O 完成端口解除关联。
对于大多数用途，首选 f.SyscallConn 方法。


### func (*File) Name
```go{1}
func (f *File) Name() string
```
返回打开时显示的文件的名称

### func (*File) Read <Badge text="重要" />
```go{1}
func (f *File) Read(b []byte) (n int, err error)
```
Read 从 File 读取最多 len(b) 个字节并将其存储在 b 中。它返回读取的字节数以及遇到的任何错误。到达文件末尾时，Read 返回 0，即 io.EOF。

### func (*File) ReadAt
```go{1}
func (f *File) ReadAt(b []byte, off int64) (n int, err error)
```
ReadAt 从文件偏移量 off 处开始读取 len(b) 个字节。它返回读取的字节数和错误（如果有）。当 n < len(b) 时，ReadAt 始终返回非零错误。到达文件末尾时，该错误为 io.EOF。

### func (*File) ReadDir <Badge text="added in go1.16" type="tip" />
```go{1}
func (f *File) ReadDir(n int) ([]DirEntry, error)
```
ReadDir 读取与文件 f 关联的目录的内容，并按目录顺序返回 DirEntry 值的切片。后续对同一文件的调用将在目录中产生后续的 DirEntry 记录。
如果 n > 0,ReadDir 最多返回 n 条 DirEntry 记录。在这种情况下，如果 ReadDir 返回一个空切片，它将返回一个错误来解释原因。在目录的末尾，错误为 io.EOF。
如果 n <= 0,ReadDir 将返回目录中剩余的所有 DirEntry 记录。当它成功时，它将返回一个 nil 错误 (而不是 io.EOF)。

### func (*File) ReadFrom <Badge text="added in go1.15" type="tip" />
```go{1}
func (f *File) ReadFrom(r io.Reader) (n int64, err error)
```
ReadFrom 实现了 io.ReaderFrom。


### func (*File) Readdir
```go{1}
func (f *File) Readdir(n int) ([]FileInfo, error)
```
Readdir 读取与 file 关联的目录的内容，并返回一个最多包含 n 个 FileInfo 值的切片，这将由 Lstat 按目录顺序返回。随后对同一文件的调用将产生更多的 FileInfo。
如果 n > 0,Readdir 最多返回 n 个 FileInfo 结构。在这种情况下，如果 Readdir 返回一个空切片，它将返回一个非 nil 错误来解释原因。在目录的末尾，错误为 io.EOF。
如果 n <= 0,Readdir 将在一个切片中返回目录中的所有 FileInfo。在这种情况下，如果 Readir 成功 (读取到目录的末尾), 它将返回切片并返回一个 nil 错误。如果在目录末尾之前遇到错误，Readdir 将返回 FileInfo 读取到那一点并返回一个非 nil 错误。
大多数客户端都使用更高效的 ReadDir 方法提供更好的服务。

### func (*File) Readdirnames
```go{1}
func (f *File) Readdirnames(n int) (names []string, err error)
```
Readdirnames 读取与 file 关联的目录的内容，并按目录顺序返回目录中最多 n 个文件名的切片。后续对同一文件的调用将产生更多的名称。
如果 n > 0,Readdirnames 最多返回 n 个名称。在这种情况下，如果 Readdirnames 返回一个空切片，它将返回一个非 nil 错误来解释原因。在目录的末尾，错误为 io.EOF。
如果 n <= 0,Readdirnames 将在一个切片中返回该目录中的所有名称。在这种情况下，当 Readdirnames 成功 (读取到目录的末尾) 时，它将返回切片并返回一个非 nil 错误。如果在目录末尾之前遇到错误，Readdirnames 将返回读取到该点之前的名称并返回一个非 nil 错误。
















































