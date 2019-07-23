/**
* @Author: zhangjian@mioji.com
* @Date: 2019/7/23 下午11:37
*/
package engine

import (
	"fmt"
	"github.com/HeavenPRC/go-operation/fetcher"
	"log"
)

func Run(seeds ...Request) {
	var requests []Request
	//将请求放入请求
	for _, r := range seeds {
		requests = append(requests, r)
	}
	//开始执行
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		log.Printf("Fether Url: %s", r.Url)
		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("Fether: error fething url %s:%v", r.Url, err)
			continue
		}
		parseResult := r.ParserFunc(body)
		//这样感觉和array_merge有点像
		requests = append(requests, parseResult.Request...)
		//fmt.Println(parseResult.Items)
		for _, r := range parseResult.Items {
			fmt.Println(r)
		}
	}
}
