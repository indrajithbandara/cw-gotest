package main

import (
    "fmt"
    "reflect"

    "github.com/demdxx/gocast"
)

type AccessRegion struct {
    region_id           int64  `tag:"RegionId"`
    provider_id         int64  `tag:"ProviderId"`
    region_name         string `tag2"RegionName"`
    sub_region_names    string
    billing_region_name string
    description         string
}

func main() {
    //数值为0,数组为空,空字符串 等情况返回true
    fmt.Println("IsEmpty:", gocast.IsEmpty(0)) //IsEmpty: true
    //转换日期，需要有时区
    t, _ := gocast.ParseTime("2012-11-01T22:08:41+00:00")
    fmt.Println("ParseTime:", t) //ParseTime: 2012-11-01 22:08:41 +0000 +0000
    //反射获取对象的列名称，如果有Tag，则优先使用其Tag值。
    var ar AccessRegion
    var props = gocast.StructFields(ar, "tag")
    for _, p := range props {
        fmt.Print(p, "\t") //RegionId    ProviderId    region_name    sub_region_names    billing_region_name    description
    }
    fmt.Println()
    //反射获取对象的列名称为Key，列名称为Value的Map集合，如果有Tag，将其相应Key的Value即为Tag中指定的名称。
    var myMap map[string]string = gocast.StructFieldTags(ar, "tag")
    for k, v := range myMap {
        fmt.Print("key:", k, " value:", v, "\t") //key:billing_region_name value:billing_region_name    key:description value:description    key:region_id value:RegionId    key:provider_id value:ProviderId    key:region_name value:region_name    key:sub_region_names value:sub_region_names
    }
    fmt.Println()
    //反射获取对象的列名称和有Tag时的列名的两个数组，如果没有相应tag标签时，将使用列名称。
    var keys, vals = gocast.StructFieldTagsUnsorted(ar, "tag")
    fmt.Print("Keys:")
    for _, p := range keys {
        fmt.Print(p, "\t") //Keys:region_id    provider_id    region_name    sub_region_names    billing_region_name    description
    }
    fmt.Print("\r\nVals:")
    for _, p := range vals {
        fmt.Print(p, "\t") //Vals:RegionId    ProviderId    region_name    sub_region_names    billing_region_name    description
    }
    fmt.Println()
    //将To方法的参数1转换为参数2的类型。Tag在类型为Slice、Map、Struct时使用。
    var obj, _ = gocast.To(0, true, "")
    fmt.Println("To:", obj, " targetType:", reflect.TypeOf(obj).Kind()) //To: false  targetType: bool
    //将指定对象转为bool，其中传入的为字符串类型时，只有"true"或"1"时为true，其它为false。数值类型时，只要不是0，都为true。集合类型时，只要集合长度不为0，都为true。
    fmt.Println(gocast.ToBool(0.00001), gocast.ToBool("True") != gocast.ToBool("true")) //true true
    //将reflect.Value类型对像转为bool类型。
    fmt.Println(gocast.ToBoolByReflect(reflect.ValueOf(0.00001))) //true
    //将指定对象转换为float64类型，bool值时，true转为1，false转为0，集合类型时，返回0.
    fmt.Println(gocast.ToFloat(true)) //return 1
    //将指定对象转换为int64类型后，再转为float32类型。此方法会丢失小数位。
    fmt.Println(gocast.ToFloat32("0.001")) //return 0
    //将给定的Slice对象中的每个元素转换为float64后组合成一个[]float64对象。
    var arr = make([]interface{}, 0)
    arr = append(arr, "0.1")
    arr = append(arr, true)
    arr = append(arr, false)
    arr = append(arr, nil)
    arr = append(arr, 3.1415)
    fmt.Println(gocast.ToFloat64Slice(arr)) //[0.1 1 0 0 3.1415]
    //将指定对像转为int
    fmt.Println(gocast.ToInt("2"), gocast.ToInt(int64(3))) //2 3
    //将给定的Slice对象中的每个元素转换为float64后组合成一个[]int对象
    fmt.Println(gocast.ToIntSlice(arr)) //[0 1 0 0 3]
    //将指定的Slice存放到另一个新的Interface{}类型的Slice中。
    fmt.Println(gocast.ToInterfaceSlice(arr)) //[0.1 true false <nil> 3.1415]
    //将第一参的类型，转为指定Type的类型。Tag在类型为Slice、Map、Struct时使用。
    fmt.Println(gocast.ToT(0, reflect.TypeOf(true), "")) //false <nil>
    //将指定对像转为字符串类型,nil时将转为空字符串
    fmt.Println(gocast.ToString(nil), gocast.ToString(1), gocast.ToString(ar)) // 1 {0 0    }
}