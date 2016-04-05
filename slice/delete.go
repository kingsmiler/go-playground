package main

import (
    "reflect"
    "fmt"
)

func main() {
    c := "a"
    d := []string{"b", "c", "d", "a"}

    // 由于重置 slice 时要求是可修改的地址，因此传递的参数需要为指针。
    DeleteInSlice(&d, c)
    fmt.Println(d)

    myGreeting := map[int]string{
        0: "Good morning!",
        1: "Bonjour!",
        2: "Buenos dias!",
        3: "Bongiorno!",
    }

    // 删除 map 元素不需要传递指针。
    DeleteInMap(myGreeting, 3)
    fmt.Println(myGreeting)
}

/**
 * 删除 map 中的指定元素。删除 map 元素不需要传递指针。
 *
 */
func DeleteInMap(targetMap interface{}, item interface{}) {

    value := reflect.Indirect(reflect.ValueOf(targetMap))

    if value.MapIndex(reflect.ValueOf(item)).IsValid() {
        value.SetMapIndex(reflect.ValueOf(item), reflect.Value{})
    }
}

/**
 * 删除 slice 中的指定元素。由于重置 slice 时要求是可修改的地址，因此传递的参数需要为指针。
 *
 */
func DeleteInSlice(slice interface{}, item interface{}) {
    index := -1
    ve := reflect.Indirect(reflect.ValueOf(slice))
    size := ve.Len()

    for i := 0; i < size; i++ {
        if reflect.DeepEqual(ve.Index(i).Interface(), item) {
            index = i
            break
        }
    }

    if index >= 0 {
        ve.Set(reflect.AppendSlice(ve.Slice(0, index), ve.Slice(index + 1, size)))
    }
}