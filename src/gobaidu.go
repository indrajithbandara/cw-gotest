package main
 
import (
    "fmt"
    "io/ioutil"
    "net"
)
 
func main() {
    var remoteAddress, _ = net.ResolveTCPAddr("tcp4", "www.baidu.com:80") //生成一个net.TcpAddr对像。
    var conn, err = net.DialTCP("tcp4", nil, remoteAddress)               //传入协议，本机地址（传了nil），远程地址，获取连接。
    if err != nil {                                                       //如果连接失败。则返回。
        fmt.Println("连接出错：", err)
        return
    }
    var remoteIpAddress = conn.RemoteAddr()  //获取IP地址的方法。
    fmt.Println("远程IP地址是：", remoteIpAddress) //输出：220.181.111.188:80
    var localIPAddress = conn.LocalAddr()
    fmt.Println("本地IP地址是：", localIPAddress) //输出：192.168.1.9:45712
 
    conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n")) //尝试发送些信息。
    //var reciverBuffer []byte                      //定义一个空切片，用于接收结果。
    //len, err := conn.Read(reciverBuffer) //返回接收到的字节数。
    bys, err := ioutil.ReadAll(conn) //接收消息。
    if err != nil {
        fmt.Println("接收出错：", err)
    }
    //var reciveText = string(reciverBuffer[0:len])
    var reciveText = string(bys)
    fmt.Println(reciveText)
    conn.Close() //关闭连接
    fmt.Println("程序结束")
}