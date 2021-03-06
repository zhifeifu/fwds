package util

import (
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
	"unsafe"
)

// Str
// var Str = newStr
// @Description:
//
var Str = newStr()

type str struct {
}

func newStr() *str {
	return &str{}
}

// StrToLower 转换成小写字母
func (us *str) StrToLower(str string) string {
	runeArr := []rune(str)
	for i, _ := range runeArr {
		if runeArr[i] >= 65 && runeArr[i] <= 90 {
			runeArr[i] += 32
		}
	}
	return string(runeArr)
}

// StrRandom 随机字符串
func (us *str) StrRandom(n int) string {
	var r = rand.New(rand.NewSource(time.Now().UnixNano()))
	const pattern = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcdefghijklmnopqrstuvwxyz"

	salt := make([]byte, 0, n)
	l := len(pattern)

	for i := 0; i < n; i++ {
		p := r.Intn(l)
		salt = append(salt, pattern[p])
	}

	return string(salt)
}

// IsEmpty 是否是空字符串
func (us *str) IsEmpty(s string) bool {
	if s == "" {
		return true
	}
	return strings.TrimSpace(s) == ""
}

// ConcatString 连接字符串
// NOTE: 性能比fmt.Sprintf和+号要好
func (us *str) ConcatString(s ...string) string {
	if len(s) == 0 {
		return ""
	}
	var buffer bytes.Buffer
	for _, i := range s {
		buffer.WriteString(i)
	}
	return buffer.String()
}

// StringToUint64 字符串转uint64
func (us *str) StringToUint64(str string) (uint64, error) {
	if str == "" {
		return 0, nil
	}
	valInt, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}

	return uint64(valInt), nil
}

// StringToInt64 字符串转int64
func (us *str) StringToInt64(str string) (int64, error) {
	if str == "" {
		return 0, nil
	}
	valInt, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}

	return int64(valInt), nil
}

// StringToInt 字符串转int
func (us *str) StringToInt(str string) (int, error) {
	if str == "" {
		return 0, nil
	}
	valInt, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}

	return valInt, nil
}

// --------- 字节切片和字符串转换 ----------
// 性能很高, 原因在于底层无新的内存申请与拷贝

// Bytes2String 字节切片转字符串
func (us *str) Bytes2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// String2Bytes 字符串转字节切片
func (us *str) String2Bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

// SubStr 截取字符串，并返回实际截取的长度和子串
func (us *str) SubStr(str string, start, end int64) (int64, string, error) {
	reader := strings.NewReader(str)

	// Calling NewSectionReader method with its parameters
	r := io.NewSectionReader(reader, start, end)

	// Calling Copy method with its parameters
	var buf bytes.Buffer
	n, err := io.Copy(&buf, r)
	return n, buf.String(), err
}

// SubstrTarget 在字符串中查找指定子串，并返回left或right部分
func (us *str) SubstrTarget(str string, target string, turn string, hasPos bool) (string, error) {
	pos := strings.Index(str, target)

	if pos == -1 {
		return "", nil
	}

	if turn == "left" {
		if hasPos == true {
			pos = pos + 1
		}
		return str[:pos], nil
	} else if turn == "right" {
		if hasPos == false {
			pos = pos + 1
		}
		return str[pos:], nil
	} else {
		return "", errors.New("params 3 error")
	}
}

// GetStringUtf8Len 获得字符串按照uft8编码的长度
func (us *str) GetStringUtf8Len(str string) int {
	return utf8.RuneCountInString(str)
}

// Utf8Index 按照uft8编码匹配子串，返回开头的索引
func (us *str) Utf8Index(str, substr string) int {
	index := strings.Index(str, substr)
	if index < 0 {
		return -1
	}
	return utf8.RuneCountInString(str[:index])
}

// JoinStringAndOther 连接字符串和其他类型
// fmt.Println(JoinStringAndOther("why", 123))
func (us *str) JoinStringAndOther(val ...interface{}) string {
	return fmt.Sprint(val...)
}

// CamelToSnake 驼峰转蛇形
func (us *str) CamelToSnake(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		// or通过ASCII码进行大小写的转化
		// 65-90（A-Z），97-122（a-z）
		//判断如果字母为大写的A-Z就在前面拼接一个_
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	//ToLower把大写字母统一转小写
	return strings.ToLower(string(data[:]))
}

// SnakeToCamel 蛇形转驼峰
func (us *str) SnakeToCamel(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}

// UcFirst 首字母大写
func UcFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

// LcFirst 首字母小写
func LcFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}
