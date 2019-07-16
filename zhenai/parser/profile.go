package parser

import (
	"crawler/engine"
	"regexp"
	"strconv"
	"crawler/model"
)

// 预编译
var nameRe = regexp.MustCompile(`<h1 class="nickName" .*>([^<]+)</h1>`)
var ageRe = regexp.MustCompile(`<div class="m-btn purple" .*>([\d]+岁)</div>`)
var xingzuoRe = regexp.MustCompile(`<div class="m-btn purple" .*>(.*\))</div>`)
var heightRe = regexp.MustCompile(`<div class="m-btn purple" .*>([\d]+kg)</div>`)
var weightRe = regexp.MustCompile(`<div class="m-btn purple" .*>([\d]+cm)</div>`)
var introduceRe = regexp.MustCompile(`<div class=\"m-content-box m-des\" .*><span .*>([^<]+)<\/span>`)

func ProfileParse(context []byte) engine.ParseResult {
	user := model.User{}
	age, err := strconv.Atoi(regexContent(context, ageRe))
	if err != nil {
		user.Age = 0
	} else {
		user.Age = age
	}

	user.Name = regexContent(context, nameRe)


}

// change byte to string by regex
func regexContent(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents);
	if len(match) < 2 {
		return ""
	} else {
		return string(match[1])
	}
}
