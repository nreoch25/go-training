package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string
	Last  string `json:"-"`
	Age   int    `json:"wisdom score"`
}

func main() {
	p1 := Person{"Nigel", "Reoch", 39}
	bs, _ := json.Marshal(p1)
	fmt.Println(string(bs))
}
