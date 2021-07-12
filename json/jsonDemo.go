package main

import (
	"fmt"

	json "github.com/bitly/go-simplejson"
)

func main() {
	val := []byte("{\"code\":0,\"result\":{\"app_wm\":\"123\"}}")

	jsonval, err := json.NewJson(val)
	if err != nil {
		fmt.Printf("Error creating JSON, val: %v, err: %s", val, err.Error())
		return
	}

	app_wm, err := jsonval.Get("result").Get("app_wm").String()
	if err != nil {
		fmt.Printf("App_wm Get error: %v", err.Error())
		return
	}

	fmt.Printf("app_wm = %s\nlen(app_wm) = %d\n", app_wm, len(app_wm))
	return
}
