/**
* @Author: zhangjian@mioji.com
* @Date: 2019/7/23 下午11:21
*/
package engine

type Request struct {
	Url string
	ParserFunc func([]byte) ParseResult
}

type ParseResult struct {
	Request []Request
	Items []interface{}
}

func NilParser(s []byte) ParseResult {
	return ParseResult{}
}