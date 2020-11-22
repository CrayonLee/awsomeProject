package main

import (
	"fmt"
	"reflect"
)

func reflectType(x interface{}){
	v := reflect.TypeOf(x)
	fmt.Printf("type:%v\n",v)
}

func reflectValue(x interface{}){
	v := reflect.ValueOf(x)	//获取接口的值信息
	k := v.Kind()	//拿到值对应的种类
	switch k{
	case reflect.Int64:
		//v.Int()从反射中获取整型的原始值，然后通过int64（）强制类型转换
		fmt.Printf("type is int64,value is %d\n",int64(v.Int()))
	case reflect.Float32:
		//v.Int()从反射中获取整型的原始值，然后通过int64（）强制类型转换
		fmt.Printf("type is float32,value is %d\n",float32(v.Float()))
	case reflect.Float64:
		//v.Int()从反射中获取整型的原始值，然后通过int64（）强制类型转换
		fmt.Printf("type is float64,value is %d\n",float64(v.Float()))

	}

}

//通过反射修改变量的值
func modifyValue(x interface{})  {
	v := reflect.ValueOf(x)
	k := v.Kind()
	if k==reflect.Ptr{
		v.Elem().SetInt(100)
	}
}
func main() {
	a:=20
	modifyValue(&a)
	fmt.Printf("%d\n",a)
}
