package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

type Config struct {
	FileName string`conf:"file_name"`
	FilePath string`conf:"file_path"`
	MaxSize int64`conf:"max_size"`
}


//从conf文件中读取内容赋值给结构体指针
func parseConf(fileName string, result interface{})(err error){
	//1.前提条件，result必须是一个指针
	t:=reflect.TypeOf(result)
	v:=reflect.ValueOf(result)
	tElem := t.Elem()
	if t.Kind()!=reflect.Ptr{
		err=errors.New("result必须是一个指针")
		return
	}

	//result不是结构体指针也不行
	if t.Elem().Kind()!=reflect.Struct{
		err=errors.New("result必须是一个结构体指针")
		return
	}
	//1.打开文件
	data, err := ioutil.ReadFile(fileName)
	if err!=nil{
		err=fmt.Errorf("打开配置文件%s失败",fileName)
		return err
	}

	//2.将读取的文件按照行分割 得到一个行的切片
	lineSlice:=strings.Split(string(data),"\r\n")
	fmt.Println(lineSlice)

	//逐行解析
	for index,line := range lineSlice {
		//去除字符串首尾的空格
		line=strings.TrimSpace(line)
		//判断当前行上否全是空格或者以#开头
		if len(line)==0||strings.HasPrefix(line,"#"){
			//忽略空行和注释
			continue
		}

		//解析当前行注释是否合法（是否有等号）
		equalIndex := strings.Index(line, "=")
		if equalIndex==-1{
			err=fmt.Errorf("第%d行语法错误",index+1)
			return
		}

		//按照等号分割每一行，左边为key，右边是value
		key:=strings.TrimSpace(line[:equalIndex])
		value:=strings.TrimSpace(line[equalIndex+1:])
		//如果key为空
		if len(key)==0{
			err=fmt.Errorf("第%d行语法错误",equalIndex+1)
			return
		}
		fmt.Println(value)
		//利用反射 给result赋值
		//遍历结构体的每一个字段和key比较，匹配上了酒吧value赋值
		for i:=0;i<tElem.NumField();i++ {
			field := tElem.Field(i)
			tag := field.Tag.Get("conf")
			if key==tag{
				//匹配上了就给value赋值
				fieldType := field.Type
				switch fieldType.Kind() {
				case reflect.String:
					v.Elem().Field(i).SetString(value)
				case reflect.Int64:
					val64, _ := strconv.ParseInt(value, 10, 64)
					v.Elem().Field(i).SetInt(val64)

				}
			}
		}

	}
	return

}
func main() {


	//2.读取内容

	//3.一行一行读取内容 根据tag找结构体中的对应字段

	//4.找到的话进行赋值操作

	var config = &Config{}
	err:=parseConf("E:\\workspace\\go\\awesomeProject\\reflect\\file.conf",config)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println("%#V",config)
}
