package core

import (
	"fmt"
	"math/rand"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

var Keys [12]string = [12]string{"富强", "民主", "文明", "和谐", "自由", "平等", "公正", "法治", "爱国", "敬业", "诚信", "友善"}

// 编码
func Encode(str string) string {
	var ret []string
	var re = regexp.MustCompile(`[A-Za-z0-9\-\_\.\!\~\*\'\(\)]`)
	for _, v := range str {
		if ok := re.MatchString(string(v)); ok {
			ret = append(ret, fmt.Sprintf("%x", v))
		} else {
			ret = append(ret, url.QueryEscape(string(v)))
		}
	}
	return hex2duo(strings.ToUpper(strings.Replace(strings.Join(ret, ""), "%", "", -1)))
}

// 转成核心价值观
func hex2duo(str string) string {
	var duo []string
	for _, v := range str {
		n, _ := strconv.ParseInt(string(v), 16, 64)
		if n < 10 {
			duo = append(duo, Keys[int(n)])
		} else {
			// 此处2种随机算法，都是根据相邻的数字来解，一种是10和n-10, 另一种是11和n-6
			if rand := rand.Intn(100); rand < 51 {
				duo = append(duo, Keys[10], Keys[n-10])
			} else {
				duo = append(duo, Keys[11], Keys[n-6])
			}
		}
	}
	return strings.Join(duo, "")
}

func Decode(str string) string {
	//var splited []string
	// 把字符串转成rune
	strRune := []rune(str)
	var hex []int
	for i := 0; i < len(strRune); i = i + 2 {
		val := string(strRune[i]) + string(strRune[i+1])
		for t, v := range Keys {
			if v == val {
				hex = append(hex, t)
			}
		}
	}
	return duo2hex(hex)
}

func duo2hex(hex []int) string {
	var ret []string
	var num int
	for i := 0; i < len(hex); {
		tmp := 0
		if hex[i] > 10 {
			i++
			tmp = hex[i] + 6
		} else if hex[i] == 10 {
			i++
			tmp = hex[i] + 10
		} else {
			tmp = hex[i]
		}
		if num&1 == 0 {
			ret = append(ret, fmt.Sprintf("%%%X", tmp))
		} else {
			ret = append(ret, fmt.Sprintf("%X", tmp))
		}
		i++
		num++
	}
	res, _ := url.QueryUnescape(strings.Join(ret, ""))
	return res
}
