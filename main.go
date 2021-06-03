package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonString := `{"name": "battery sensor", "capacity": 40, "time":"2019-01-21T19:07:28Z", "info":{
						"desc": "A sensor reading"
					}}`

	var reading map[string]interface{}
	err := json.Unmarshal([]byte(jsonString), &reading)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", reading)
}
