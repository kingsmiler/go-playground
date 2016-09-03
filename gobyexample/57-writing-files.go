package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func checkX(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fileName := "write_test.txt"

	// 将字节数组写入文件
	d1 := []byte("hello\ngo\n")
	// 如果文件不存在，根据指定权限创建一个新文件，再写入数据
	err := ioutil.WriteFile(fileName, d1, 0644)
	checkX(err)

	// 通过文件对象更好地控制文件的读写
	f, err := os.Create(fileName)
	checkX(err)
	// 应当习惯性地在得到文件对象后，开启它的完成后自动关闭
	defer f.Close()

	// 写入 slices 字节数组
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	checkX(err)
	fmt.Printf("wrote %d bytes\n", n2)

	// 写入字符串
	n3, err := f.WriteString("writes\n")
	fmt.Printf("wrote %d bytes\n", n3)

	// 通过 Sync 保存（同步）数据
	f.Sync()

	// 通过 writer 对象来写数据
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n", n4)

	// writer 通过 flush 来确保全部数据都已写入
	w.Flush()
}
