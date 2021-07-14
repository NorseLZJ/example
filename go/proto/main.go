package main

import (
	"fmt"
	"log"
	"proto/gameserver/login"

	"google.golang.org/protobuf/proto"
)

func main() {
	msg := &login.Person{
		Name:  "lisa",
		Id:    1,
		Email: "18841685054@163.com",
	}
	msg.Phones = append(msg.Phones, &login.Person_PhoneNumber{
		Number: "12345678",
		Type:   1,
	})
	msg.Phones = append(msg.Phones, &login.Person_PhoneNumber{
		Number: "012345678",
		Type:   2,
	})

	fmt.Printf("src:%v\n", msg)

	out, err := proto.Marshal(msg)
	if err != nil {
		log.Fatal("marshal err:%v", err)
	}

	read := &login.Person{}
	err = proto.Unmarshal(out, read)
	if err != nil {
		log.Fatal("unmarshal err:%v", err)
	}

	fmt.Printf("--------\n")
	fmt.Printf("dest:%v\n", read)

	//fmt.Printf("persion:%+v\n", msg)
}
