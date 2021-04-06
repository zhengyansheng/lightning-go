package tools

import (
	"encoding/json"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/structs"
)

// string to int
func StringToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

// string to uint
func StringToUint(s string) (v uint, err error) {
	var _int int
	_int, err = strconv.Atoi(s)
	v = uint(_int)
	return
}

func StructToMap(s interface{}) (m map[string]interface{}) {
	m = structs.Map(s)
	return
}

func StringToMap(s []byte) (m map[string]interface{}, err error) {
	m = make(map[string]interface{})
	err = json.Unmarshal(s, &m)
	return
}

func ByteToJson(b []byte) (data map[string]interface{}, err error) {
	err = json.Unmarshal(b, &data)
	return
}

func JsonToByte(data map[string]interface{}) (b []byte, err error) {
	b, err = json.Marshal(data)
	return
}

func ReplaceDateTime(dTime string) string {
	// "2020-11-02T15:38Z"
	s := strings.Replace(dTime, "T", " ", -1)
	return strings.Trim(strings.Replace(s, "Z", " ", -1), " ")
}

//利用正则表达式压缩字符串，去除空格或制表符
func CompressStr(str string) string {
	if str == "" {
		return ""
	}
	//匹配一个或多个空白符的正则表达式
	reg := regexp.MustCompile("\\s+")
	return reg.ReplaceAllString(str, "")
}

func FormatNormalDateTime(dateTime time.Time) string {
	return dateTime.Format("2006-01-02 15:04:05")
}
