package nlp

import (
	"regexp"
	"net/url"
	"net/http"
	"strings"
	"io/ioutil"
	"github.com/astaxie/beego"
)

// 获取关键字
func GetKeywords(content string) (words []string) {
	str := curlPost(url.Values{"action": {"keyword"}, "len": {beego.AppConfig.String("keywordsLen")}, "nature": {"false"}, "content": {html2text(content)}})
	words = strings.Split(str, "\t")
	return
}

// 调用NLP程序，分词
func curlPost(params url.Values) (string) {
	resp, err := http.PostForm(beego.AppConfig.String("nlpHost"), params)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		return ""
	}
	return string(body)
}

// 过滤HTML代码
func html2text(str string) (res string) {
	str = regexp.MustCompile("<!DOCTYPE.*?>").ReplaceAllString(str, "")
	str = regexp.MustCompile("<!--.*?-->").ReplaceAllString(str, "")
	str = regexp.MustCompile("<script.*?>.*?<\\/script>").ReplaceAllString(str, "")
	str = regexp.MustCompile("<style.*?>.*?<\\/style>").ReplaceAllString(str, "")
	str = regexp.MustCompile("<textarea.*?>.*?<\\/textarea>").ReplaceAllString(str, "")
	str = regexp.MustCompile("<input[^>]*?>.*?<\\/input>").ReplaceAllString(str, "")
	str = regexp.MustCompile("<button[^>]*?>.*?<\\/button>").ReplaceAllString(str, "")
	str = regexp.MustCompile("<table[^>]*?>.*?<\\/table>").ReplaceAllString(str, "")
	str = regexp.MustCompile("&.{1,5};|&#.{1,5};").ReplaceAllString(str, "")
	str = regexp.MustCompile("<\\/?(.*?)\\>").ReplaceAllString(str, "")
	str = regexp.MustCompile("&amp;nbsp;").ReplaceAllString(str, "")
	str = regexp.MustCompile(` `).ReplaceAllString(str, "")
	str = regexp.MustCompile(`　`).ReplaceAllString(str, "")
	res = str
	return
}
