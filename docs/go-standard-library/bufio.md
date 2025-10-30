# bufio

bufio 包实现了缓冲 I/O。它包装了一个 io.Reader 或 io.Writer 对象，并创建了另一个对象（Reader 或 Writer），该对象也实现了该接口，但提供了缓冲功能以及一些文本 I/O 的帮助。


## type Scanner

Scanner 提供了一个便捷的接口，用于读取数据，例如以换行符分隔的文本行文件。连续调用 Scanner.Scan 方法将遍历文件中的“标记”，并跳过标记之间的字节。标记的规范由 SplitFunc 类型的 split 函数定义；默认的 split 函数会将输入拆分为行，并去除行终止符 。Scanner.Split 此软件包中定义了用于将文件扫描到指定位置的函数。 行、字节、UTF-8 编码的符文和空格分隔的单词。 客户也可以提供自定义拆分函数。

扫描会在遇到文件末尾 (EOF)、第一个 I/O 错误或超出 Scanner.Buffer 容量的标记时永久停止。扫描停止时，读取器可能已经前进到超出最后一个标记的任意位置。需要更精细地控制错误处理或处理大型标记，或者必须在读取器上执行顺序扫描的程序，应改用 bufio.Reader 。

### func NewScanner
```go{1}
func NewScanner(r io.Reader) *Scanner
```
NewScanner 返回一个新的 Scanner 来从 r 读取。split 函数默认为 ScanLines 。


### func (*Scanner) Buffer ¶

```go
func (s *Scanner) Buffer(buf []byte, max int)
```
缓冲区控制扫描仪的内存分配。它设置扫描时使用的初始缓冲区以及扫描期间可分配的最大缓冲区大小。缓冲区的内容将被忽略。
最大令牌大小必须小于 max 和 cap(buf) 中的较大者。如果 max <= cap(buf)， Scanner.Scan 将仅使用此缓冲区，不进行任何分配。
默认情况下， Scanner.Scan 使用内部缓冲区并将最大令牌大小设置为 MaxScanTokenSize 。
如果在扫描开始后调用 Buffer，则会导致混乱。

### func (*Scanner) Bytes
```go
func (s *Scanner) Bytes() []byte
```
Bytes 返回通过调用 Scanner.Scan 生成的最新令牌。底层数组可能指向将被后续 Scan 调用覆盖的数据。它不进行任何分配。

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(strings.NewReader("gopher"))
	for scanner.Scan() {
		fmt.Println(len(scanner.Bytes()) == 6)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "shouldn't see an error scanning a string")
	}
}
```
```text 
true
```

### func (*Scanner) Err 

Err 返回 Scanner 遇到的第一个非 EOF 错误。

### func (*Scanner) Scan
```go
func (s *Scanner) Scan() bool
```

Scan 将 Scanner 推进到下一个标记，该标记将通过 Scanner.Bytes 或 Scanner.Text 方法可用。当没有更多标记时，它会返回 false, 这可能是因为到达输入末尾或出现错误。在 Scan 返回 false 后，Scanner.Err 方法将返回扫描过程中发生的任何错误，除非是 io.EOF,Scanner.Err 将返回 nil。如果拆分函数返回太多空标记而没有推进输入，Scan 将出现恐慌。这是 Scanner 的一种常见错误模式。

### func (*Scanner) Split
```go
func (s *Scanner) Split(split SplitFunc)
```
Split 设置扫描器的分割函数。默认分割函数为 ScanLines 。
如果在扫描开始后调用 Split，则会引发 panic。

### func (*Scanner) Text
```go
func (s *Scanner) Text() string
```
Text 返回通过调用 Scanner.Scan 生成的最新令牌。底层数组可能指向将被后续 Scan 调用覆盖的数据。它不进行任何分配。


---

## type Reader

Reader 类为 io.Reader 对象实现了缓冲功能。可以通过调用 NewReader 或 NewReaderSize 创建一个新的 Reader 对象；或者，也可以在调用Reset方法后使用 Reader 对象的零值。

### func NewReader
```go
func NewReader(rd io.Reader) *Reader
```
NewReader 返回一个新的 Reader ，其缓冲区大小为默认值。

### func NewReaderSize
```go
func NewReaderSize(rd io.Reader, size int) *Reader
```
NewReaderSize 返回一个新的 Reader ，其缓冲区大小为 size 。

### func (*Reader) Read
```go
func (b *Reader) Read(p []byte) (n int, err error)
```
Read 方法将数据读取到内存单元 p 中，并返回读取到 p 中的字节数。这些字节最多只能从底层 Reader 的一个 Read 操作中读取，因此 n 可能小于 len(p)。要精确读取 len(p) 个字节，请使用 io.ReadFull(b, p)。如果底层 Reader 可以通过 io.EOF 返回非零计数，则此 Read 方法也可以这样做；请参阅 io.Reader 文档。

### func (*Reader) ReadLine
```go
func (b *Reader) ReadLine() (line []byte, isPrefix bool, err error)
```
ReadLine 是一个底层行读取原语。大多数调用者应该使用它。 请改用 Reader.ReadBytes ('\n') 或 Reader.ReadString ('\n')，或者使用 Scanner 。

ReadLine 函数尝试返回单行数据，不包括行尾字节。如果行长度超过缓冲区大小，则 isPrefix 属性会被置位，并返回行的开头部分。行的剩余部分将在后续调用中返回。当返回行的最后一个片段时，isPrefix 属性将为 false。返回的缓冲区仅在下次调用 ReadLine 之前有效。ReadLine 函数要么返回一个非空行，要么返回一个错误，不会同时返回两者。

ReadLine 返回的文本不包含行尾（“\r\n”或“\n”）。如果输入结束时没有行尾，则不会给出任何提示或错误。在 ReadLine 之后调用 Reader.UnreadByte 将始终取消读取​​最后一个字节（可能是行尾的字符），即使该字节不属于 ReadLine 返回的行。

## type Writer
Writer 为 io.Writer 对象实现了缓冲机制。如果向 Writer 写入数据时发生错误，则不会再接受任何数据，并且所有后续的写入操作以及 Writer.Flush 方法都会返回该错误。 所有数据写入完毕后，客户端应该调用 Writer.Flush 方法保证所有数据都已转发到底层 io.Writer 。

```go{13}
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	fmt.Fprint(w, "Hello, ")
	fmt.Fprint(w, "world!")
	w.Flush() // Don't forget to flush!
}
```

### func NewWriter
```go
func NewWriter(w io.Writer) *Writer
```
NewWriter 返回一个新的 Writer， 其缓冲区大小为默认值。如果参数 io.Writer 已经是一个缓冲区大小足够大的 Writer ，则返回底层 Writer 。

### func NewWriterSize
```go
func NewWriterSize(w io.Writer, size int) *Writer
```
NewWriterSize 返回一个新的 Writer， 其缓冲区大小至少为指定值。如果参数 io.Writer 已经是一个大小足够大的 Writer ，则返回底层 Writer 。


### func (*Writer) Flush
```go
func (b *Writer) Flush() error
```
Flush 将任何缓冲的数据写入底层 io.Writer。

### func (*Writer) Reset
```go
func (b *Writer) Reset(w io.Writer)
```
Reset 方法会丢弃所有未刷新的缓冲数据，清除所有错误，并将 b 重置为将其输出写入 w。对 Writer 的零值调用 Reset 会将内部缓冲区初始化为默认大小。调用 w.Reset(w)（即将 Writer 重置为自身）不会执行任何操作。

### func (*Writer) Size
```go
func (b *Writer) Size() int
```
Size 返回底层缓冲区的大小（以字节为单位）。

### func (*Writer) Write
```go
func (b *Writer) Write(p []byte) (nn int, err error)
```
write 函数将 p 的内容写入缓冲区，并返回已写入的字节数。如果 nn < len(p)，则还会返回一个错误，解释写入量不足的原因。

### func (*Writer) WriteString
```go
func (b *Writer) WriteString(s string) (int, error)
```





