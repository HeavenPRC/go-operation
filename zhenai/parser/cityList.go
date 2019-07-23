/**
* @Author: zhangjian@mioji.com
* @Date: 2019/7/23 下午11:19
*/
package parser

import (
	"github.com/HeavenPRC/go-operation/zhenai/engine"
	"regexp"
)

const CityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	reg := regexp.MustCompile(CityListRe)
	matches := reg.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, string(m[2]))
		result.Request = append(result.Request, engine.Request{
			Url:string(m[1]),
			ParserFunc:engine.NilParser,
		})
	}
	return result
}
