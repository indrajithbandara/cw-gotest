package main
 
import (
    "fmt"
    "net/http"
    "net/url"
 
    "github.com/drone/routes"
)
 
func getuser(w http.ResponseWriter, r *http.Request) {
    var params url.Values = r.URL.Query()
    var uid string = params.Get(":uid")
    fmt.Fprintln(w, "get a user ", uid, " success!")
}
func getuserAndAge(w http.ResponseWriter, r *http.Request) {
    var params url.Values = r.URL.Query()
    var uid string = params.Get(":uid")
    var age string = params.Get(":age")
    fmt.Fprintln(w, "get a user ", uid, " success! age is ", age)
}
func edituser(w http.ResponseWriter, r *http.Request) {
    var params url.Values = r.URL.Query()
    var uid string = params.Get(":uid")
    fmt.Fprintln(w, "edit a user ", uid, " success!")
}
func main() {
    fmt.Println("正在启动WEB服务...")
    var mux *routes.RouteMux = routes.New()
    mux.Get("/user/:uid", getuser)
    mux.Get("/user/:uid/:age", getuserAndAge)
    mux.Post("/user/:uid", edituser)
 
    //http.Handle("/", mux)
    http.ListenAndServe(":8088", mux)
    fmt.Println("服务已停止")
}

//export GOPATH=/mnt/pmscodes/cw-gotest/    go build src/gotounte.go
//go run gotounte.go
//http://127.0.0.1:8088/user/cui/32