package main

type ToolStruct struct {
	Events    []interface{} `json:"events"`
	FrameRate int64         `json:"frameRate"`
	Frames    []*Frame      `json:"frames"`
}

type Frame struct {
	Res string  `json:"res"`
	X   float32 `json:"x"`
	Y   float32 `json:"y"`
}
