/**
* @Author: zhangjian@mioji.com
* @Date: 2019/7/20 下午5:34
*/
package main

import (
	"github.com/HeavenPRC/go-operation/zhenai/engine"
	"github.com/HeavenPRC/go-operation/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:"http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
