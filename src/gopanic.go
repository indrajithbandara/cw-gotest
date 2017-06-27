package main
 
import "fmt"
 
func main() {
    fmt.Printf("hello world my name is %s, I'm %d\r\n", "songxingzhu", 26)
    defer func() {
        if err := recover(); err != nil {
            fmt.Println("出了错：", err)
        }
    }()
    myPainc()
    fmt.Printf("这里应该执行不到！")
}
func myPainc() {
    var x = 30
    var y = 0
    //panic("我就是一个大错误！")
    var c = x / y
    fmt.Println(c)
}