
// defer类似于Java的finally语句块，在程序或函数的末尾，即将关闭时执行，通常用于清理操作。
package main

import "fmt"
import "os"

// 本示例假定需要创建一个文件，然后写入数据，完成后关闭文件。
// 本示例展示在该场景中如何使用defer。
func main() {

	// 通常在得到文件对象后，立即defer对它的关闭操作；
	// defer操作将在函数执行的最后依次执行。
	f := createFile("godefer.txt")
	defer closeFile(f)

	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	f.Close()
}
