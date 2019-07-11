/**
 *	此包用来 将网页信息转成[]byte，让engine去调用不同的parse解析相应的数据
 */
package fetcher

import (
	"net/http"
	"golang.org/x/text/transform"
	"io/ioutil"
	"io"
	"bufio"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"github.com/pkg/errors"
	"fmt"
	"golang.org/x/text/encoding/unicode"
)

func Fetch(url string) ([]byte, error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(request)
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("error: status code is %s", resp.StatusCode))
	}

	// 使用 golang.org/text/encoding/simplifiedchinese 文件 来将GBK的字符集转化成UTF8, 如果网页字符集是utf8或者其他格式，则会转码失败
	// 所以我们需要go官方提供的另一个包 golang.org/net/html 来预知 网页的字符集才不会搞错
	// utf8Reader := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())
	// 获取网页的 具体 charset 然后使用 对应的 charset 解码器 读取数据
	e := determineEncoding(resp.Body)
	utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())
	// 不要请求头
	return ioutil.ReadAll(utf8Reader)
}

// 包装charset.DetermineEncoding
func determineEncoding(r io.Reader) encoding.Encoding {
	// 为什么会把resp.body转到从Buff中读取????
	// 唯一想到的原因1. 可以直接读取1024字节, 原因2：速度快
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		return unicode.UTF8
	}
	// 让第三方库去猜网页的编码, 返回编码的解析器
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
