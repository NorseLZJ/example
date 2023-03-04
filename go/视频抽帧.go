package main

import (
	"bytes"
	"errors"
	"fmt"
	"image/jpeg"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/util/gconv"
	gim "github.com/ozankasikci/go-image-merge"
)

func FFmpegVideoSecTotal(in string) (int64, error) {
	/*
		获取视频总秒数,大小等部分参数,帧数不能在这里边获取
	*/
	args := []string{"-i", in,
		"-select_streams", "v",
		"-show_entries", "format=duration,size",
		"-of", "json"}
	cmd := exec.Command("ffprobe", args...)
	//fmt.Println(cmd.String())
	closer, err := cmd.StdoutPipe()
	if err != nil {
		return 0, err
	}
	defer func() {
		_ = closer.Close()
		_ = cmd.Wait()
	}()
	err = cmd.Start()
	if err != nil {
		return 0, err
	}
	bytes, err := ioutil.ReadAll(closer)
	if err != nil {
		return 0, err
	}
	ss := strings.TrimSpace(string(bytes))
	fmt.Println(ss)
	jsoner, err := gjson.LoadJson(ss)
	if err != nil {
		return 0, err
	}
	duration := gconv.String(jsoner.Get("format.duration"))
	sec, _ := strconv.ParseFloat(duration, 64)
	return int64(sec), nil
}

func FFmpegCmdSecRangeScreenShot(in string) error {
	start := time.Now().Unix()
	sec, err := FFmpegVideoSecTotal(in)
	if sec == 0 {
		return errors.New("换一个链接")
	}
	if err != nil {
		return err
	}
	stdErr := bytes.Buffer{}
	avgSec := int(sec / 20)
	args := []string{"-i", in, "-vf", fmt.Sprintf("fps=1/%d", avgSec), fmt.Sprintf("%s.jpg", "%03d")}
	cmd := exec.Command("ffmpeg", args...)
	cmd.Stderr = &stdErr
	// fmt.Println(cmd.String())
	if err := cmd.Run(); err != nil {
		fmt.Printf("FFmpeg err:%s\n", stdErr.String())
		return err
	}
	MergeJpg()
	end := time.Now().Unix()
	fmt.Printf("总耗时:%d (秒)", end-start)
	return nil
}

func MergeJpg() {
	grids := []*gim.Grid{}
	count := 0
	_ = filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
		if filepath.Ext(d.Name()) == ".jpg" && d.Name() != "merged.jpg" {
			count += 1
			grids = append(grids, &gim.Grid{ImageFilePath: d.Name()})
		}
		return nil
	})
	if len(grids) == 0 {
		return
	}
	rgba, err := gim.New(grids, count, 1).Merge()
	if err != nil {
		log.Panicf(err.Error())
	}
	file, err := os.Create("merged.jpg")
	if err != nil {
		log.Panicf(err.Error())
	}
	err = jpeg.Encode(file, rgba, &jpeg.Options{Quality: 80})
	if err != nil {
		log.Panicf(err.Error())
	}
	_ = filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
		if filepath.Ext(d.Name()) == ".jpg" && d.Name() != "merged.jpg" {
			os.Remove(d.Name())
		}
		return nil
	})
}

func main() {
	downUrl := "http://download-cdn.123pan.cn/123-225/89608dd3/1812752236-0/89608dd324fec988554e4f88e1260ce1?v=2&t=1678003020&s=974aef6b34bfc8b411d2e650e99ccff0&filename=video.mp4&d=4adfd1cd"
	err := FFmpegCmdSecRangeScreenShot(downUrl)
	if err != nil {
		log.Fatal(err)
	}
}
