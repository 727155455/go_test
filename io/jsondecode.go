package main

import (
	"encoding/json"
	"fmt"
)

var jsonStr string = `
{
    "retcode": 2000,
    "msg": "已认证",
    "data": {
        "card_face": "http://hao.shengmo.org:888/Uploads/Picture/2018-09-16/20180916172112913.jpg",
        "status": "0",
		"msg": "已认证"
    },
    "data2": {
        "card_face": "http://hao.shengmo.org:888/Uploads/Picture/2018-09-16/20180916172112913.jpg",
		"msg": "已认证",
        "status": "0"
    },
    "data3": [
		{
	        "card_face": "http://hao.shengmo.org:888/Uploads/Picture/2018-09-16/20180916172112913.jpg",
	        "status": "0",
			"msg": "已认证"
	    },
		{
	        "card_face": "http://hao.shengmo.org:888/Uploads/Picture/2018-09-16/20180916172112913.jpg",
	        "status": "0"
	    }
	]
}
`

type gResp struct {
	RetCode int16 `json:"retcode"`
	Msg string `json:"msg"`
	Data2 struct {
		Status string `json:"status"`
		CardFace string `json:"card_face"`
	} `json:"data2"`
	Data4 []myData `json:"data3"`
	Data myData `json:"data"`
}

type myData struct {
	CardFace string `json:"card_face"`
	Status string `json:"status"`
}

func main(){
	/*url := "http://59.110.28.150:888/Api/Other/getidstate?uid=202904"

	resp, err := http.Get(url)
	fmt.Printf("%+v\n", resp)
	defer resp.Body.Close()

	if err != nil {
		fmt.Println(err)
	}

	var a []interface{}
	err1 := json.NewDecoder(resp.Body).Decode(a)
	if err1 != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", a)*/

	mydecode()
}

func mydecode(){
	var a gResp
	err := json.Unmarshal([]byte(jsonStr), &a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", a)
	fmt.Println("-------------")
	var b map[string]interface{}
	err = json.Unmarshal([]byte(jsonStr), &b)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", b)
}

func myencode() {
	c := make(map[string]interface{})
	c["name"] = "aaa"
	c["age"] = 10
	c["bb"] = map[string]interface{}{
		"home": "beijing",
		"cell": "123456",
	}

	data, err := json.MarshalIndent(c, "", "	")
	if err != nil {
		fmt.Println(err)
	}
	data2, err := json.Marshal(c)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data),string(data2))
}