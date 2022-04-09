package util

import (
	"encoding/json"
	"fmt"
)

func MarshalIndent(v any) string {
	ret, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		fmt.Printf("marshal json failed, err: %v, v: %v", err.Error(), v)
	}
	return string(ret)
}
