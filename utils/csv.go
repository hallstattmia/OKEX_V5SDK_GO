package utils

import (
	"bufio"
	"encoding/csv"
	"log"
	"os"
	"strings"
)

type Csv struct {
	File *os.File
}

func (csv *Csv) OpenCsv(path string, head string) {
	//OpenFile读取文件，不存在时则创建，使用追加模式
	File, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Println("文件打开失败！")
	}
	csv.File = File
	csv.WriterCsv(head)
}

func (csv *Csv) CloseCsv() {
	defer csv.File.Close()
}

func (c *Csv) WriterCsv(str string) {
	//创建写入接口
	WriterCsv := csv.NewWriter(bufio.NewWriter(c.File))
	//写入一条数据，传入数据为切片(追加模式)
	err1 := WriterCsv.Write(strings.Split(str, ","))
	if err1 != nil {
		log.Println("WriterCsv写入文件失败")
	}
	WriterCsv.Flush() //刷新，不刷新是无法写入的
	log.Println("数据写入成功...")
}
