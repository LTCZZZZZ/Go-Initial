package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

const jsonStream = `{"Message": "Hello", "Array": [1, 2, 3], "Null": null, "Number": 1.234}`

func main() {
	dec := json.NewDecoder(strings.NewReader(jsonStream))

	//for {
	//	t, err := dec.Token()
	//	fmt.Println("1111: ", t, err)
	//	if err == io.EOF {
	//		break
	//	}
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	//fmt.Printf("%T: %v", t, t)
	//	//if dec.More() {
	//	//	fmt.Printf(" (more)")
	//	//}
	//	//fmt.Printf("\n")
	//}

	// 注意，dec是一个流数据，只能读取一次，如果上面的代码解析了，这里得不到任何东西
	var m map[string]interface{}
	err := dec.Decode(&m)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(m)
}
