/**
* @Author: zhangjian@mioji.com
* @Date: 2019/7/20 下午5:34
*/
package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Println("status code ", resp.StatusCode)
		return
	}
    encode := determineEncodeing(resp.Body)
    utf8Reader := trans
	fmt.Println(encode)
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(all))
}

//获取网站字符集编码
func determineEncodeing (r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}