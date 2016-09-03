package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 读取文件需要检查是否存在异常
func checkX(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fileName := "hello.txt"

	// 一次性读取文件的内容并赋值给变量
	dat, err := ioutil.ReadFile(fileName)
	checkX(err)
	fmt.Print(string(dat))

	// 通过得到文件对象，控制文件的读取
	f, err := os.Open(fileName)
	checkX(err)
	defer f.Close()

	// 将数据保存到字节数组，返回读取的有效个数
	b1 := make([]byte, 6)
	n1, err := f.Read(b1)
	checkX(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))

	// 通过Seek偏移指定节点，然后读取或写入，返回偏移的距离
	o2, err := f.Seek(3, 0)
	checkX(err)
	b2 := make([]byte, 3)
	n2, err := f.Read(b2)
	checkX(err)
	fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))

	// 通过io工具包的方法 ReadAtLeast 进行读取
	o3, err := f.Seek(3, 0)
	checkX(err)
	b3 := make([]byte, 6)
	n3, err := io.ReadAtLeast(f, b3, 2)
	checkX(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	// There is no built-in rewind, but `Seek(0, 0)`
	// accomplishes this.
	_, err = f.Seek(0, 0)
	checkX(err)

	//  bufio 工具包提供可缓存的reader，有更多的方法读取文件
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(6)
	checkX(err)
	fmt.Printf("6 bytes: %s\n", string(b4))

	line,_,_ := r4.ReadLine()
	fmt.Println("line = ", string(line))
}
