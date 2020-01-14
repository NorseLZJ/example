package main

import (
	"flag"
)

var (
	cfg = flag.String("conf", "./goGet.json", "go get file")
)

var (
	defaultCmd = "go get -u "
)

func main() {
	//cfgT, err := config.Marshal(*cfg)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//for _, v := range cfgT.Code {
	//	fmt.Println(v)
	//}
	//fmt.Println(cfgT.Proxy)

	//path, err := exec.LookPath("go")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(path)

	//cmd := exec.Command("tr", "a-z", "A-Z")
	//cmd.Stdin = strings.NewReader("some input")

	//var out bytes.Buffer
	//cmd.Stdout = &out
	//err := cmd.Run()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("in all caps : %q\n", out.String())

	//cmd := exec.Command("ls")
	//cmd.Env = append(os.Environ(),
	//	"FOO=duplicate_value",
	//	"FOO=actual_value")

	//if err := cmd.Run(); err != nil { //	log.Fatal(err)
	//}

	//ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	//defer cancel()
	//if err := exec.CommandContext(ctx, "sleep", "5").Run(); err != nil {
	//	log.Fatal(err)
	//}

	//cmd := exec.Command("sh", "-c", "echo stdout; echo 1>&2 stderr")
	//stdoutStderr, err := cmd.CombinedOutput()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("%s\n", stdoutStderr)

	//out, err := exec.Command("date").Output()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("The date is :%s\n", out)

	//cmd := exec.Command("ps", "aux")
	//log.Printf("Running command and waiting for it to finish")
	//err := cmd.Run()
	//log.Printf("Command finished with error : %v", err)

	//cmd := exec.Command("sleep", "5")
	//err := cmd.Start()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Printf("waiting for command to finish ...")
	//err = cmd.Wait()
	//log.Printf("command finished with error : %v", err)
}
