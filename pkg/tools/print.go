package tools

import (
	"encoding/json"
	"fmt"
)

// 格式化打印
func PrettyPrint(v interface{}) {
	b, err := json.MarshalIndent(v, "", "    ")
	if err == nil {
		fmt.Println(string(b))
	} else {
		fmt.Println(b)
	}

}
