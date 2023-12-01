package main

import (
	pbdef "GmTest/proto"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	any "github.com/golang/protobuf/ptypes/any"
)

func main() {
	v := pbdef.EditScene2Req{
		Point: nil,
		Data: map[string]*any.Any{
			"a": {},
		},
	}
	mapInfo := &pbdef.MapInfo{
		Id:  1,
		Pid: 1,
	}
	any, err := ptypes.MarshalAny(mapInfo)
	if err != nil {
		panic(err)
	}
	v.Data["any"] = any
	data, err := proto.Marshal(&v)
	if err != nil {
		panic(err)
	}

	out := &pbdef.EditScene2Req{}
	err = proto.Unmarshal(data, out)
	if err != nil {
		panic(err)
	}
	outMapInfo := &pbdef.MapInfo{}
	err = ptypes.UnmarshalAny(out.Data["any"], outMapInfo)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", outMapInfo)
}
