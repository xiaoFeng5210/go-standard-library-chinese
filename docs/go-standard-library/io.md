# io

包 io 为 I/O 基元提供基本接口。它的主要任务是将这些基元的现有实现 (如包 os 中的实现) 封装到共享的公共接口中，这些接口抽象了功能，以及一些其他相关的基元。
因为这些接口和基元用各种实现封装了底层操作，除非另有通知，否则客户端不应该认为它们对并行执行是安全的。

## func Copy
```go{1}
func Copy(dst Writer, src Reader) (written int64, err error)
```

将 src 复制到 dst, 直到 src 上达到 EOF 或发生错误。它返回复制的字节数以及复制过程中遇到的第一个错误 (如果有)。
成功的复制返回 err == nil, 而不是 err == EOF。由于复制被定义为从 src 读取到 EOF, 因此它不会将 Read 中的 EOF 视为要报告的错误。
如果 src 实现了 WriterTo, 则复制通过调用 src.WriteTo (dst) 实现。否则，如果 dst 实现了 ReaderFrom, 则复制通过调用 dst.ReadFrom (src) 实现。

#### 使用示例
```go
package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")

	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}

}
```


```go
some io.Reader stream to be read
```

## func CopyBuffer
```go{1}
func CopyBuffer(dst Writer, src Reader, buf []byte) (written int64, err error)
```
CopyBuffer 与 Copy 完全相同，只是它会在提供的缓冲区中分段 (如果需要的话), 而不是分配一个临时缓冲区。如果 buf 为 nil, 则分配一个；否则，如果它的长度为 0, 则 CopyBuffer 会发生恐慌。
如果 src 实现 WriterTo 或 dst 实现 ReaderFrom, 则 buf 不会用于执行副本。

#### 使用示例
```go
package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	r1 := strings.NewReader("first reader\n")
	r2 := strings.NewReader("second reader\n")
	buf := make([]byte, 8)

	// buf is used here...
	if _, err := io.CopyBuffer(os.Stdout, r1, buf); err != nil {
		log.Fatal(err)
	}

	// ... reused here also. No need to allocate an extra buffer.
	if _, err := io.CopyBuffer(os.Stdout, r2, buf); err != nil {
		log.Fatal(err)
	}

}
```

```go
first reader
second reader
```

## func CopyN
```go{1}
func CopyN(dst Writer, src Reader, n int64) (written int64, err error)
```

CopyN 从 src 复制 n 个字节到 dst（或直到出现错误）。它返回复制的字节数以及复制过程中遇到的最早错误。当且仅当 err == nil 时，返回值 == n。

```go
package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	r := strings.NewReader("some io.Reader stream to be read")

	if _, err := io.CopyN(os.Stdout, r, 4); err != nil {
		log.Fatal(err)
	}

}
```
```text
some
```

## func Pipe
```go{1}
func Pipe() (*PipeReader, *PipeWriter)
```
Pipe 创建一个同步内存管道。它可以用来连接需要 io.Reader 的代码。 代码需要 io.Writer 。

管道上的读取和写入操作是一对一匹配的，除非需要多个读取操作才能消费单个写入操作。也就是说，每次写入 PipeWriter 的操作都会阻塞，直到它满足了 PipeReader 的一个或多个读取操作，并完全消费了写入的数据。数据会直接从写入操作复制到相应的读取操作（或多个读取操作）；没有内部缓冲。

并行调用 Read 和 Write 函数或使用 Close 函数都是安全的。并行调用 Read 和 Write 函数也是安全的：各个调用将按顺序进行。

```go
package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	r, w := io.Pipe()

	go func() {
		fmt.Fprint(w, "some io.Reader stream to be read\n")
		w.Close()
	}()

	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}

}
```
```text
some io.Reader stream to be read
```

## func ReadAll
```go{1}
func ReadAll(r Reader) ([]byte, error)
```
ReadAll 从 r 读取数据，直到出现错误或 EOF，并返回读取的数据。成功调用返回 err == nil，而不是 err == EOF。由于 ReadAll 定义为从 src 读取直到 EOF，因此它不会将 Read 中的 EOF 视为需要报告的错误。

```go
package main

import (
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	r := strings.NewReader("Go is a general-purpose language designed with systems programming in mind.")

	b, err := io.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", b)
}
```
```text
Go is a general-purpose language designed with systems programming in mind.
```

## func ReadAtLeast
```go{1}
func ReadAtLeast(r Reader, buf []byte, min int) (n int, err error)
```
ReadAtLeast 从 r 读取到 buf 中，直到至少读取 min 个字节。它返回复制的字节数，如果读取的字节数少于 min 个字节，则返回错误。仅当未读取任何字节时，错误才为 EOF。如果在读取少于 min 个字节后发生 EOF，ReadAtLeast 将返回 ErrUnexpectedEOF 。如果 min 大于 buf 的长度，ReadAtLeast 将返回 ErrShortBuffer 。返回时，当且仅当 err == nil 时，n >= min。如果 r 在读取至少 min 个字节后返回错误，则该错误将被丢弃。

```go
package main

import (
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")

	buf := make([]byte, 14)
	if _, err := io.ReadAtLeast(r, buf, 4); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", buf)

	// buffer smaller than minimal read size.
	shortBuf := make([]byte, 3)
	if _, err := io.ReadAtLeast(r, shortBuf, 4); err != nil {
		fmt.Println("error:", err)
	}

	// minimal read size bigger than io.Reader stream
	longBuf := make([]byte, 64)
	if _, err := io.ReadAtLeast(r, longBuf, 64); err != nil {
		fmt.Println("error:", err)
	}
}
```

```text
some io.Reader
error: short buffer
error: unexpected EOF
```

## func ReadFull
```go{1}
func ReadFull(r Reader, buf []byte) (n int, err error)
```
ReadFull 从 r 读取恰好 len(buf) 个字节到 buf 中。它返回复制的字节数，如果读取的字节数少于此值，则返回错误。仅当未读取任何字节时，错误才会是 EOF。如果在读取部分字节（而非全部字节）后发生 EOF，ReadFull 将返回 ErrUnexpectedEOF 。返回时，当且仅当 err == nil 时，n == len(buf)。如果 r 在读取至少 len(buf) 个字节后返回错误，则该错误将被丢弃。

```go
package main

import (
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")

	buf := make([]byte, 4)
	if _, err := io.ReadFull(r, buf); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", buf)

	// minimal read size bigger than io.Reader stream
	longBuf := make([]byte, 64)
	if _, err := io.ReadFull(r, longBuf); err != nil {
		fmt.Println("error:", err)
	}

}
```
```text
some
error: unexpected EOF
```

## func WriteString
```go{1}
func WriteString(w Writer, s string) (n int, err error)
```
WriteString 将 s 写入 w 中。它返回复制的字节数，以及写入过程中遇到的第一个错误。

```go
package main

import (
	"io"
	"log"
	"os"
)

func main() {
	if _, err := io.WriteString(os.Stdout, "Hello World"); err != nil {
		log.Fatal(err)
	}
}
```
```text
Hello World
```
