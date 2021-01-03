package main

import "flag"

type Task struct {
	Cmd    string `json:"cmd"`
	Script string `json:"script"`
	Space  string `json:"space"`
}

type Config struct {
	LogDir string `json:"log_dir"`
	Tasks  []Task `json:"tasks"`
}

var (
	cfg = flag.String("c", "./ccrontab.json", "crontab config json")
)

func main() {
	flag.Parse()

}
