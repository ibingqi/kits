package main

import (
	"encoding/csv"
	"os"
	"strconv"
)

type Meta struct {
	Id int64
}

func main() {
	// 创建一个 demo.csv 文件
	csvFile, err := os.Create("demo.csv")
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	// 初始化一个 csv writer，并通过这个 writer 写入数据到 csv 文件
	writer := csv.NewWriter(csvFile)

	for i := 1000000123; i < 1000000123+2000000; i++ {
		line := []string{
			strconv.Itoa(i),
		}
		// 将切片类型行数据写入 csv 文件
		err := writer.Write(line)
		if err != nil {
			panic(err)
		}
	}
	// 将 writer 缓冲中的数据都推送到 csv 文件，至此就完成了数据写入到 csv 文件
	writer.Flush()

	/*
		// 打开这个 csv 文件
		file, err := os.Open("tutorials.csv")
		if err != nil {
			panic(err)
		}
		defer file.Close()

		// 初始化一个 csv reader，并通过这个 reader 从 csv 文件读取数据
		reader := csv.NewReader(file)
		// 设置返回记录中每行数据期望的字段数，-1 表示返回所有字段
		reader.FieldsPerRecord = -1
		// 通过 readAll 方法返回 csv 文件中的所有内容
		record, err := reader.ReadAll()
		if err != nil {
			panic(err)
		}

		// 遍历从 csv 文件中读取的所有内容，并将其追加到 tutorials2 切片中
		var tutorials2 []Tutorial
		for _, item := range record {
			id, _ := strconv.ParseInt(item[0], 0, 0)
			tutorial := Tutorial{Id: int(id), Title: item[1], Summary: item[2], Author: item[3]}
			tutorials2 = append(tutorials, tutorial)
		}

		// 打印 tutorials2 的第一个元素验证 csv 文件写入/读取是否成功
		fmt.Println(tutorials2[0].Id)
		fmt.Println(tutorials2[0].Title)
		fmt.Println(tutorials2[0].Summary)
		fmt.Println(tutorials2[0].Author)
	*/
}
