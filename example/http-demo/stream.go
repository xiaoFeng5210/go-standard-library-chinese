package http_demo

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
)

// 读取流响应
func streamResponse() {
	req, err := http.NewRequest("GET", "xxx", nil)
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}

// 读流方式2 适合字节流
func streamWithRead(resp *http.Response) {
	buffer := make([]byte, 1024)
	for {
		_, err := resp.Body.Read(buffer)

		if err == io.EOF {
			break
		}
	}
	fmt.Println(string(buffer))

}

// 方式3: 使用 io.CopyBuffer（最简单，适合文件下载等）
func streamWithCopy(resp *http.Response, dst io.Writer) {
	buffer := make([]byte, 32*1024) // 32KB 缓冲区
	io.CopyBuffer(dst, resp.Body, buffer)
}

// 流式发送请求体
func streamRequest() {
	pr, pw := io.Pipe()

	go func() {
		defer pw.Close()
		data := []string{"line1", "line2", "line3"}
		for _, chunk := range data {
			_, err := pw.Write([]byte(chunk))
			if err != nil {
				panic(err)
			}
		}
	}()

	req, err := http.NewRequest("POST", "xxx", pr)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "text/plain")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	fmt.Println(string(body))
}
