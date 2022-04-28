package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "chenliang"

	fmt.Println(str[:len(str)-1])
}

//题目大意忘了，是找前缀
//["flower","flow","flight"]
func longestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	}
	//选取第一个元素为前缀
	prefix := strs[0]
	//遍历字符串数组
	for _, v := range strs {
		//如果当前前缀不是该字符串的子串，则将前缀去调第一位
		for strings.Index(v, prefix) != 0 {
			if len(prefix) == 0 {
				return ""
			}
			prefix = prefix[:len(prefix)-1]
		}
	}
	return prefix
}
