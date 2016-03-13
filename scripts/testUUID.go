package main

import (
	"log"
"github.com/yeldars/crm/utils"
)

func main() {
	// Creating UUID Version 4
	for i:=1;i<100;i++ {
	u1 := utils.Uuid()
	log.Println(u1)
	}
}