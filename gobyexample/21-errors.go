package main

import "errors"
import "fmt"

// 在Go语言中,错误类型是显式地作为末尾的返回参数定义的。
func f1(arg int) (int, error) {
    if arg == 42 {

        // 通过 errors.New 构造一个错误对象。
        return -1, errors.New("can't work with 42")
    }

    // 如果无错误,则错误对象为nil。
    return arg + 3, nil
}

// 自定义错误类型,需要实现 Error() 方法。
type argError struct {
    arg  int
    prob string
}

// 错误类型的方法实现。
func (e *argError) Error() string {
    return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
    if arg == 42 {
        // &argError 创建对象
        return -1, &argError{arg, "can't work with it"}
    }

    return arg + 3, nil
}

func main() {

    // 默认错误类型使用
    for _, i := range []int{7, 42} {
        if r, e := f1(i); e != nil {
            fmt.Println("f1 failed:", e)
        } else {
            fmt.Println("f1 worked:", r)
        }
    }

    // 自定义错误类型使用
    for _, i := range []int{7, 42} {
        if r, e := f2(i); e != nil {
            fmt.Println("f2 failed:", e)
        } else {
            fmt.Println("f2 worked:", r)
        }
    }

    // 使用自定义错误类型中的数据
    _, e := f2(42)
    if ae, ok := e.(*argError); ok {
        fmt.Println(ae.arg)
        fmt.Println(ae.prob)
    }
}
