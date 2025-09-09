# time

包 time 提供了时间的显示和测量用的函数。日历的计算采用的是公历。

## func Date
```go
func Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time
```
Date 返回指定年月日时分秒纳秒和时区的时间。
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Printf("Go launched at %s\n", t.Local())
}
```
```text
Go launched at 2009-11-10 15:00:00 -0800 PST
```

## func Now
```go
func Now() Time
```
Now 返回当前本地时间。


## func Parse
```go
func Parse(layout, value string) (Time, error)
```
Parse 解析格式化字符串并返回其所表示的时间值。请参阅名为 Layout 的常量文档，了解如何表示格式。第二个参数必须能够使用作为第一个参数提供的格式字符串（布局）进行解析。

Time.Format 的示例详细演示了布局字符串的工作原理，是一个很好的参考。

解析（仅解析）时，输入可能在秒字段后紧接着包含秒的小数部分，即使布局未指示其存在。在这种情况下，逗号或小数点后跟最大数字序列将被解析为秒的小数部分。秒的小数部分将被截断为纳秒精度。

布局中省略的元素被假定为零，或者当不可能为零时，假定为一，因此解析“3:04pm”将返回对应于 UTC 时间 0 年 1 月 1 日 15:04:00 的时间（请注意，由于年份为 0，因此该时间早于零时间）。年份必须在 0000 到 9999 范围内。系统会检查星期几的语法，但其他情况将被忽略。

对于指定两位数年份 06 的布局，NN >= 69 的值将被视为 19NN，NN < 69 的值将被视为 20NN。

此注释的其余部分描述了时区的处理。

如果没有时区指示符，Parse 将返回 UTC 时间。

解析带有时区偏移量（例如 -0700）的时间时，如果该偏移量对应于当前位置（本地）使用的时区，则 Parse 会在返回的时间中使用该位置和时区。否则，它会将时间记录为位于虚构位置的时间，其时间固定在给定的时区偏移量。

解析带有时区缩写（例如 MST）的时间时，如果该时区缩写在当前位置具有已定义的偏移量，则使用该偏移量。无论位于何处，时区缩写“UTC”都会被识别为 UTC。如果时区缩写未知，Parse 会将时间记录为具有给定时区缩写和零偏移的虚构位置。这意味着此类时间可以使用相同的布局进行解析和重新格式化，且无损，但表示中使用的精确时刻将因实际时区偏移而有所不同。为避免此类问题，请优先使用使用数字时区偏移的时间布局，或使用 ParseInLocation。

## func ParseInLocation
```go
func ParseInLocation(layout, value string, loc *Location) (Time, error)
```
ParseInLocation 与 Parse 类似，但在两个重要方面有所不同。首先，如果没有时区信息，Parse 会将时间解释为 UTC；而 ParseInLocation 会将时间解释为给定位置的时间。其次，当给定时区偏移量或缩写时，Parse 会尝试将其与本地位置进行匹配；而 ParseInLocation 则会使用给定的位置。
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	loc, _ := time.LoadLocation("Europe/Berlin")

	// This will look for the name CEST in the Europe/Berlin time zone.
	const longForm = "Jan 2, 2006 at 3:04pm (MST)"
	t, _ := time.ParseInLocation(longForm, "Jul 9, 2012 at 5:02am (CEST)", loc)
	fmt.Println(t)

	// Note: without explicit zone, returns time in given location.
	const shortForm = "2006-Jan-02"
	t, _ = time.ParseInLocation(shortForm, "2012-Jul-09", loc)
	fmt.Println(t)
}
```
```text
2012-07-09 05:02:00 +0200 CEST
2012-07-09 00:00:00 +0200 CEST
```

## Type Time <Badge text="类型" type="info" />
Time 表示纳秒精度的时间点。

使用时间的程序通常应该将其作为值存储和传递，而不是指针。也就是说，时间变量和结构体字段应该是 time.Time 类型，而不是 *time.Time 类型。

Time 值可以被多个 goroutine 同时使用，但 Time.GobDecode、Time.UnmarshalBinary、Time.UnmarshalJSON 和 Time.UnmarshalText 方法并非并发安全。

可以使用 Time.Before、Time.After 和 Time.Equal 方法比较时间点。Time.Sub 方法将两个时间点相减，得到 Duration。Time.Add 方法将 Time 和 Duration 相加，得到 Time。

Time 类型的零值是 UTC 时间 1 年 1 月 1 日 00:00:00.000000000。由于该时间在实际中不太可能出现，Time.IsZero 方法提供了一种简单的方法来检测未明确初始化的时间。

每个时间都有一个关联的 Location。Time.Local、Time.UTC 和 Time.In 方法返回一个具有特定 Location 的时间值。使用这些方法更改时间值的位置不会改变它所代表的实际时刻，只会改变解释它的时区。

Time.GobEncode、Time.MarshalBinary、Time.AppendBinary、Time.MarshalJSON、Time.MarshalText 和 Time.AppendText 方法保存的时间值表示会存储 Time.Location 的偏移量，但不存储位置名称。因此，它们会丢失有关夏令时的信息。

除了必需的“挂钟”读数外，Time 还可以包含当前进程单调时钟的可选读数，以便为比较或减法提供更高的精度。有关详情，请参阅包文档中的“单调时钟”部分。

请注意，Go 的 == 运算符不仅会比较时间点，还会比较位置和单调时钟读数。因此，不应将时间值用作映射或数据库键，除非首先确保所有值都设置了相同的位置（这可以通过使用 UTC 或 Local 方法实现），并且已通过设置 t = t.Round(0) 去除单调时钟读数。通常，优先使用 t.Equal(u) 而不是 t == u，因为 t.Equal 使用了最精确的比较方法，并且能够正确处理只有一个参数具有单调时钟读数的情况。

### func (Time) Add
```go
func (t Time) Add(d Duration) Time
```
时间加法 t+d。
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Date(2009, 1, 1, 12, 0, 0, 0, time.UTC)
	afterTenSeconds := start.Add(time.Second * 10)
	afterTenMinutes := start.Add(time.Minute * 10)
	afterTenHours := start.Add(time.Hour * 10)
	afterTenDays := start.Add(time.Hour * 24 * 10)

	fmt.Printf("start = %v\n", start)
	fmt.Printf("start.Add(time.Second * 10) = %v\n", afterTenSeconds)
	fmt.Printf("start.Add(time.Minute * 10) = %v\n", afterTenMinutes)
	fmt.Printf("start.Add(time.Hour * 10) = %v\n", afterTenHours)
	fmt.Printf("start.Add(time.Hour * 24 * 10) = %v\n", afterTenDays)

}
```
```text
start = 2009-01-01 12:00:00 +0000 UTC
start.Add(time.Second * 10) = 2009-01-01 12:00:10 +0000 UTC
start.Add(time.Minute * 10) = 2009-01-01 12:10:00 +0000 UTC
start.Add(time.Hour * 10) = 2009-01-01 22:00:00 +0000 UTC
start.Add(time.Hour * 24 * 10) = 2009-01-11 12:00:00 +0000 UTC
```

### func (Time) AddDate
```go
func (t Time) AddDate(years int, months int, days int) Time
```
AddDate 返回将给定的年数、月数和天数添加到 t 后对应的时间。例如，将 AddDate(-1, 2, 3) 应用于 2011 年 1 月 1 日，将返回 2010 年 3 月 4 日。

请注意，日期本质上与时区相关，而日历周期（例如天）没有固定的持续时间。AddDate 使用时间值的位置来确定这些持续时间。这意味着相同的 AddDate 参数可能会根据基准时间值及其位置产生不同的绝对时间偏移。例如，将 AddDate(0, 0, 1) 应用于 3 月 27 日 12:00，将始终返回 3 月 28 日 12:00。在某些地区和某些年份，这是一个 24 小时的偏移。在其他年份，由于夏令时转换，这是一个 23 小时的偏移。

AddDate 以与 Date 相同的方式对其结果进行规范化，例如，在 10 月 31 日加上一个月会得到 12 月 1 日，即 11 月 31 日的规范化形式。
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Date(2023, 03, 25, 12, 0, 0, 0, time.UTC)
	oneDayLater := start.AddDate(0, 0, 1)
	dayDuration := oneDayLater.Sub(start)
	oneMonthLater := start.AddDate(0, 1, 0)
	oneYearLater := start.AddDate(1, 0, 0)

	zurich, err := time.LoadLocation("Europe/Zurich")
	if err != nil {
		panic(err)
	}
	// This was the day before a daylight saving time transition in Zürich.
	startZurich := time.Date(2023, 03, 25, 12, 0, 0, 0, zurich)
	oneDayLaterZurich := startZurich.AddDate(0, 0, 1)
	dayDurationZurich := oneDayLaterZurich.Sub(startZurich)

	fmt.Printf("oneDayLater: start.AddDate(0, 0, 1) = %v\n", oneDayLater)
	fmt.Printf("oneMonthLater: start.AddDate(0, 1, 0) = %v\n", oneMonthLater)
	fmt.Printf("oneYearLater: start.AddDate(1, 0, 0) = %v\n", oneYearLater)
	fmt.Printf("oneDayLaterZurich: startZurich.AddDate(0, 0, 1) = %v\n", oneDayLaterZurich)
	fmt.Printf("Day duration in UTC: %v | Day duration in Zürich: %v\n", dayDuration, dayDurationZurich)

}
```
```text
oneDayLater: start.AddDate(0, 0, 1) = 2023-03-26 12:00:00 +0000 UTC
oneMonthLater: start.AddDate(0, 1, 0) = 2023-04-25 12:00:00 +0000 UTC
oneYearLater: start.AddDate(1, 0, 0) = 2024-03-25 12:00:00 +0000 UTC
oneDayLaterZurich: startZurich.AddDate(0, 0, 1) = 2023-03-26 12:00:00 +0200 CEST
Day duration in UTC: 24h0m0s | Day duration in Zürich: 23h0m0s
```


### func (Time) After
```go
func (t Time) After(u Time) bool
```
After 报告时间点 t 是否在 u 之后。

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	year2000 := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	year3000 := time.Date(3000, 1, 1, 0, 0, 0, 0, time.UTC)

	isYear3000AfterYear2000 := year3000.After(year2000) // True
	isYear2000AfterYear3000 := year2000.After(year3000) // False

	fmt.Printf("year3000.After(year2000) = %v\n", isYear3000AfterYear2000)
	fmt.Printf("year2000.After(year3000) = %v\n", isYear2000AfterYear3000)

}

```
```text
year3000.After(year2000) = true
year2000.After(year3000) = false
```

### func (Time) AppendFormat
```go
func (t Time) AppendFormat(b []byte, layout string) []byte
```
AppendFormat 与 Time.Format 类似，但将文本表示附加到 b 并返回扩展缓冲区。

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(2017, time.November, 4, 11, 0, 0, 0, time.UTC)
	text := []byte("Time: ")

	text = t.AppendFormat(text, time.Kitchen)
	fmt.Println(string(text))
}
```
```text
Time: 11:00AM
```

### func (Time) Before
```go
func (t Time) Before(u Time) bool
```
Before 报告时间点 t 是否在 u 之前。
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	year2000 := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	year3000 := time.Date(3000, 1, 1, 0, 0, 0, 0, time.UTC)

	isYear2000BeforeYear3000 := year2000.Before(year3000) // True
	isYear3000BeforeYear2000 := year3000.Before(year2000) // False

	fmt.Printf("year2000.Before(year3000) = %v\n", isYear2000BeforeYear3000)
	fmt.Printf("year3000.Before(year2000) = %v\n", isYear3000BeforeYear2000)

}
```
```text
year2000.Before(year3000) = true
year3000.Before(year2000) = false
```

### func (Time) Clock
```go
func (t Time) Clock() (hour, min, sec int)
```
Clock 返回 t 指定日期内的小时、分钟和秒。

### func (Time) Compare
```go
func (t Time) Compare(u Time) int
```
时间比较，Compare 函数将时间点 t 与 u 进行比较。如果 t 在 u 之前，则返回 -1；如果 t 在 u 之后，则返回 +1；如果它们相同，则返回 0。

### func (Time) Date
```go
func (t Time) Date() (year int, month Month, day int)
```
日期返回 t 发生的年、月、日。
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	d := time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC)
	year, month, day := d.Date()

	fmt.Printf("year = %v\n", year)
	fmt.Printf("month = %v\n", month)
	fmt.Printf("day = %v\n", day)

}
```
```text
year = 2000
month = February
day = 1
```

















## func After
```go
func After(d Duration) <-chan Time
```
等待指定时间结束后，将当前时间发送到返回的通道。在 Go 1.23 之前，该文档警告底层 垃圾收集器只有在定时器触发后才会回收 Timer ，如果效率问题，代码应该使用 NewTimer，并在不再需要定时器时调用 Timer.Stop 。从 Go 1.23 开始，垃圾收集器可以回收未引用且未停止的定时器。既然 After 也能做到，那就没有理由优先使用 NewTimer 了。

```go
package main

import (
	"fmt"
	"time"
)

var c chan int

func handle(int) {}

func main() {
	select {
	case m := <-c:
		handle(m)
	case <-time.After(10 * time.Second):
		fmt.Println("timed out")
	}
}
```
```text
timed out
```


## func Sleep ¶
```go
func Sleep(d Duration)
```
Sleep 阻塞当前 goroutine 至少 d 纳秒。它既不消耗 CPU 时间，也不产生任何系统调用。

```go
package main

import (
	"time"
)

func main() {
	time.Sleep(100 * time.Millisecond)
}
```

## func Tick
```go
func Tick(d Duration) <-chan Time
```

Tick 是 NewTicker 的一个便捷包装器，仅提供对 ticking 通道的访问。与 NewTicker 不同，如果 d <= 0，Tick 将返回 nil。

在 Go 1.23 之前，该文档警告底层 Ticker 永远不会被垃圾收集器回收，如果效率是一个问题，代码应该使用 NewTicker 并在不再需要 Ticker 时调用 Ticker.Stop 。从 Go 1.23 开始，垃圾收集器可以回收未引用的 Ticker，即使它们尚未停止。Stop 方法对于垃圾收集器来说不再是必需的。既然 Tick 已经足够，那么就没有必要再选择 NewTicker 了。

```go
package main

import (
	"fmt"
	"time"
)

func statusUpdate() string { return "" }

func main() {
	c := time.Tick(5 * time.Second)
	for next := range c {
		fmt.Printf("%v %s\n", next, statusUpdate())
	}
}
```

## Duration <Badge text="类型" type="info" />
```go
type Duration int64
```
Duration 表示两个时间点之间经过的时间，纳秒为单位。通俗的讲就是时间计量单位。
```go
package main

import (
	"fmt"
	"time"
)

func expensiveCall() {}

func main() {
	t0 := time.Now()
	expensiveCall()
	t1 := time.Now()
	fmt.Printf("The call took %v to run.\n", t1.Sub(t0))
}
```

### func (Duration) Abs
```go
func (d Duration) Abs() Duration
```
Abs 返回 d 的绝对值。特殊情况下，Duration( math.MinInt64 ) 会转换为 Duration( math.MaxInt64 )，将其量级减少 1 纳秒。
```go
package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	positiveDuration := 5 * time.Second
	negativeDuration := -3 * time.Second
	minInt64CaseDuration := time.Duration(math.MinInt64)

	absPositive := positiveDuration.Abs()
	absNegative := negativeDuration.Abs()
	absSpecial := minInt64CaseDuration.Abs() == time.Duration(math.MaxInt64)

	fmt.Printf("Absolute value of positive duration: %v\n", absPositive)
	fmt.Printf("Absolute value of negative duration: %v\n", absNegative)
	fmt.Printf("Absolute value of MinInt64 equal to MaxInt64: %t\n", absSpecial)

}
```
```text
Absolute value of positive duration: 5s
Absolute value of negative duration: 3s
Absolute value of MinInt64 equal to MaxInt64: true
```

### func (Duration) Hours
```go
func (d Duration) Hours() float64
```
小时以浮点数小时的形式返回持续时间。
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	h, _ := time.ParseDuration("4h30m")
	fmt.Printf("I've got %.1f hours of work left.", h.Hours())
}
```
```text
I've got 4.5 hours of work left.
```

### func (Duration) Minutes
```go
func (d Duration) Minutes() float64
```
Minutes 返回持续时间以分钟为单位的浮点数表示。
```go
```

## func ParseDuration
```go
func ParseDuration(s string) (Duration, error)
```
ParseDuration 用于解析持续时间字符串。持续时间字符串是一个可能带符号的十进制数序列，每个数字带有可选的小数部分和单位后缀，例如“300ms”、“-1.5h”或“2h45m”。有效的时间单位为“ns”、“us”（或“µs”）、“ms”、“s”、“m”、“h”。

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	hours, _ := time.ParseDuration("10h")
	complex, _ := time.ParseDuration("1h10m10s")
	micro, _ := time.ParseDuration("1µs")
	// The package also accepts the incorrect but common prefix u for micro.
	micro2, _ := time.ParseDuration("1us")

	fmt.Println(hours)
	fmt.Println(complex)
	fmt.Printf("There are %.0f seconds in %v.\n", complex.Seconds(), complex)
	fmt.Printf("There are %d nanoseconds in %v.\n", micro.Nanoseconds(), micro)
	fmt.Printf("There are %6.2e seconds in %v.\n", micro2.Seconds(), micro2)
}
```
```text
10h0m0s
1h10m10s
There are 4210 seconds in 1h10m10s.
There are 1000 nanoseconds in 1µs.
There are 1.00e-06 seconds in 1µs.
```

## func Since
```go
func Since(t Time) Duration
```
Since 返回自 t 以来经过的时间。它是 time.Now().Sub(t) 的简写。

```go
package main

import (
	"fmt"
	"time"
)

func expensiveCall() {}

func main() {
	start := time.Now()
	expensiveCall()
	elapsed := time.Since(start)
	fmt.Printf("The call took %v to run.\n", elapsed)
}
```

## func Until
```go
func Until(t Time) Duration
```
Until 返回到 t 为止的持续时间。它是 t.Sub(time.Now()) 的简写。

```go
package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	futureTime := time.Now().Add(5 * time.Second)
	durationUntil := time.Until(futureTime)
	fmt.Printf("Duration until future time: %.0f seconds", math.Ceil(durationUntil.Seconds()))
}
```
```text
Duration until future time: 5 seconds
```




