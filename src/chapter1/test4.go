package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	var str1 string = "my name is chenliang"
	fmt.Println("第一个长度", len(str1))
	var str2 = "陈亮"
	//go语言用Utf-8， 一个中文站三个字节
	fmt.Println("第二个长度", len(str2))
	//统计unicode字符数
	fmt.Println("还是第二个长度,", utf8.RuneCountInString(str2))

	//ASCII遍历用下标，Uncode遍历用range
	str := "我在学习go语言"
	for i := 0; i < len(str); i++ {
		//输出的是数字
		fmt.Println(str[i])
	}
	name := "c罗纳尔多"
	for _, s := range name {
		//输出的是数字
		//fmt.Println(s)
		//得格式化
		fmt.Printf("%c", s)
	}

	//字符串截取
	apple := "你今天怎么了？ 没事啊。"
	index1 := strings.Index(apple, "？")
	index2 := strings.Index(apple, "今天")
	fmt.Println("index1:", index1)
	fmt.Println("index2:", index2)
	fmt.Println("直接截取:", apple[index1:])

	//修改字符串
	var alterStr string = "hello this world"
	//类型装换
	arr := []byte(alterStr)
	for i := 6; i <= 10; i++ {
		arr[i] = ' '
	}
	fmt.Println("修改后的结果:", string(arr))

	//字符串拼接
	var stringBuild bytes.Buffer
	stringBuild.WriteString("又是你")
	stringBuild.WriteString("吃我一电炮")
	fmt.Println(stringBuild.String())

	//格式化字符串
	var process = 2
	var target = 8
	title := fmt.Sprintf("数字1：%d,数字2：%d", process, target)
	fmt.Println(title)
	//按数值本身格式输出
	pi := 3.1415926
	variant := fmt.Sprintf("%v %v %v", "月球基地", pi, true)
	fmt.Println(variant)
	//匿名结构体声明并赋初值
	profile := &struct {
		NAME string
		HP   int
	}{
		NAME: "chenliang",
		HP:   18,
	}
	fmt.Printf("%+v\n", profile)
	fmt.Printf("%#v\n", profile)
	fmt.Printf("%T\n", profile)
}
