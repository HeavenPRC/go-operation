/**
* 处理charset,导出byte
* @Author: zhangjian@mioji.com
* @Date: 2019/7/22 下午11:52
*/
package fetcher

import (
	"bufio"
	"errors"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("status code " + string(resp.StatusCode))
	}
	//设置字符集
	encode := determineEncodeing(resp.Body)
	utf8Reader := transform.NewReader(resp.Body, encode.NewDecoder())
	//fmt.Println(encode)
	all, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		return nil, err
	}
	return  all, nil
}

//获取网站字符集编码
func determineEncodeing (r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		log.Printf("Fether error: %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}