package main

type ToolStruct struct {
	Events    []interface{} `json:"events"`
	FrameRate int64         `json:"frameRate"`
	Frames    []*Frame      `json:"frames"`
}

type Frame struct {
	Res string `json:"res"`
	X   int    `json:"x"`
	Y   int    `json:"y"`
}
