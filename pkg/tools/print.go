package t

import (
	"encoding/json"
	"fmt"
	"log"
)

func PrintJSON(v any) {
	result, _ := json.MarshalIndent(v, "", "\t")
	log.Println(string(result))
}

func PrintStruct(v any) {
	fmt.Printf("%#v\n", v)
}
