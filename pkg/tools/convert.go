package tools

import (
	"encoding/json"
	"strconv"
	"strings"

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
	return strings.Replace(s, "Z", " ", -1)
}
