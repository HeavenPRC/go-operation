package main

import (
	"fmt"
	"regexp"
)

const text = `
My email is 
zhangjian@mioji.com
aa qq@qq.com
bb aaz@ss.css
`

func main() {
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9.]+)\.([a-zA-Z0-9]+)`)
	//match := re.FindAllString(text, -1)//找出所有
	matxx := re.FindAllStringSubmatch(text, -1)//获取子匹配
	//fmt.Println(match, matxx)
	for _,m := range matxx {
		fmt.Println(m)
	}
}