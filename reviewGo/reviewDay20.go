package reviewGo

import (
	"encoding/json"
	"fmt"
)

func RunDay20() {
	testJson()
}

func testJson() {
	type Monster struct {
		Name     string
		Age      string
		Birthday string
		Sal      float64
		Skill    string
	}
	laoniu := Monster{
		Name:     "牛魔王",
		Age:      "2300",
		Birthday: "10-09",
		Sal:      800.0,
		Skill:    "还我漂漂拳",
	}
	fmt.Println(laoniu)
	d, _ := json.Marshal(&laoniu)
	fmt.Println(string(d))

	tm := make(map[string]string)
	tm["name"] = "奔波霸"
	tm["age"] = "999"
	tm["sal"] = "999"
	tm["skill"] = "奔波霸之打滚"
	fmt.Println(tm)
	t, _ := json.Marshal(&tm)
	fmt.Println(string(t))

	str := "{\"Name\":\"牛魔王\",\"Age\":\"2300\",\"Birthday\":\"10-09\",\"Sal\":800,\"Skill\":\"还我漂漂拳\"}\n"
	var m Monster
	err := json.Unmarshal([]byte(str), &m)
	if err != nil {
		return
	}
	fmt.Println(m)
}
