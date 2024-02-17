package word

import (
	"strings"
	"unicode"
)

// ToUpper 全部转大写
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// ToLower 全部转小写
func ToLower(s string) string {
	return strings.ToLower(s)
}

// UnderscoreToUpperCamelCase 下划线转大写驼峰
func UnderscoreToUpperCamelCase(s string) string {
	//s为要处理的字符串,old为要替换的字符串,new为要替换成的字符串. n为替换的数量个数. n=-1时为替换全部.
	s = strings.Replace(s, "_", "", -1)
	//全部以大小写的单词开头
	s = strings.Title(s)
	return strings.Replace(s, " ", "", -1)
}

// UnderscoreToLowerCamelCase 下划线转小写驼峰
func UnderscoreToLowerCamelCase(s string) string {
	s = UnderscoreToUpperCamelCase(s)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

// CamelCaseToUnderscore 驼峰转下划线
func CamelCaseToUnderscore(s string) string {
	var output []rune
	for i, r := range s {
		if i == 0 {
			output = append(output, unicode.ToLower(r))
			continue
		}
		if unicode.IsUpper(r) {
			output = append(output, '_')
		}
		output = append(output, unicode.ToLower(r))
	}
	return string(output)
}
