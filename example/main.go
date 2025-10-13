package main

import (
	"encoding/json"
	"fmt"
)

type ListEgItem struct {
	Name string
	Desc string
}

func main() {
	item := &ListEgItem{
		Name: "1",
		Desc: "描述1",
	}
	item2 := &ListEgItem{
		Name: "2",
		Desc: "描述2",
	}
	// 创建测试数据
	list := []*ListEgItem{
		item,
		item2,
	}

	listStr, _ := json.Marshal(list)
	fmt.Println(string(listStr))
}
