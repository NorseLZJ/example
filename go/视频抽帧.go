package main

import (
	"bytes"
	"errors"
	"fmt"
	"image/jpeg"
	"io"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/disintegration/imaging"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/util/gconv"
	gim "github.com/ozankasikci/go-image-merge"
)

func FFmpegVideoFrameTotal(in string) (int64, error) {
	/*
		获取总帧数，这个命令有问题，但是拿到命令行执行又是没有问题的
	*/
	args := []string{"-v", "error",
		"-count_frames",
		"-select_streams", " v:0",
		"-show_entries", "stream=nb_read_frames",
		"-of", "default=nokey=1:noprint_wrappers=1", in}
	cmd := exec.Command("ffprobe", args...)
	fmt.Println(cmd.String())
	closer, err := cmd.StdoutPipe()
	if err != nil {
		return 0, err
	}
	defer func() {
		_ = closer.Close()
	}()
	err = cmd.Start()
	if err != nil {
		return 0, err
	}
	err = cmd.Wait()
	if err != nil {
		return 0, err
	}
	bytes, err := ioutil.ReadAll(closer)
	if err != nil {
		return 0, err
	}
	ss := strings.TrimSpace(string(bytes))
	fmt.Println(ss)
	//	jsoner, err := gjson.LoadJson(ss)
	//	if err != nil {
	//		return 0, err
	//	}
	//	duration := gconv.String(jsoner.Get("format.duration"))
	//	sec, _ := strconv.ParseFloat(duration, 64)
	return int64(0), nil
}

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

func FFmpegCmdFrameScreenShot(in string, sec int64) error {
	/*
		按照帧数截取图片， 不大行
		会截取第一次，后续会断错误，也可能是参数有误
	*/
	subSec := int64(1)
	if sec >= 9 {
		subSec = sec / 9
	}
	fmt.Println(subSec)
	arg := fmt.Sprintf("select=(gte(t\\,%d))*(isnan(prev_selected_t)+gte(t-prev_selected_t\\,%d))", subSec, subSec)
	args := []string{"-i", in, "-vf", arg, "-vsync", "0", "image.jpg"}
	cmd := exec.Command("ffmpeg", args...)
	fmt.Println(cmd.String())
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func FFmpegCmdSecScreenShot(in string) error {
	/*
		按照秒截取图片
		唯一的问题大概是，视频较长，按秒的话，会比较费时
		就目前的规则来说，一个60分钟的视频，截取9张，拼成9宫格
	*/
	start := time.Now().Unix()
	sec, err := FFmpegVideoSecTotal(in)
	if sec == 0 {
		return errors.New("换一个链接")
	}
	if err != nil {
		return err
	}
	avgSec := int(sec / 9)
	count := 1
	for count <= 9 {
		ss := fmt.Sprintf("%d", avgSec*count)
		fmt.Println(ss)
		args := []string{"-i", in, "-f", "image2", "-vcodec", "mjpeg", "-vframes", "1", "-ss", ss, "pipe2"}
		cmd := exec.Command("ffmpeg", args...)
		if err := cmd.Run(); err != nil {
			return err
		}
		reader := ReadFile("pipe2")
		img, err := imaging.Decode(reader)
		if err != nil {
			return err
		}
		imgFile := fmt.Sprintf("image_tmp/%s_out.jpg", ss)
		err = imaging.Save(img, imgFile)
		if err != nil {
			return err
		}
		_ = os.Remove("pipe2")
		count += 1
	}
	end := time.Now().Unix()
	fmt.Printf("总耗时:%d (秒)", end-start)
	return nil
}

func ReadFile(path string) io.Reader {
	f, e := os.Open(path)
	if e != nil {
		log.Fatal(e.Error())
	}
	return f
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
		if d == nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if filepath.Ext(d.Name()) == ".jpg" && d.Name() != "merged.jpg" {
			count += 1
			grids = append(grids, &gim.Grid{
				ImageFilePath: d.Name(),
			})
		}
		return nil
	})
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
}

func main() {
	downUrl := "http://download-cdn.123pan.cn/123-225/89608dd3/1812752236-0/89608dd324fec988554e4f88e1260ce1?v=2&t=1678003020&s=974aef6b34bfc8b411d2e650e99ccff0&filename=video.mp4&d=4adfd1cd"
	//_, err := FFmpegVideoFrameSize(downUrl)
	//if err != nil {
	//	log.Fatal(err)
	//}

	/*
		err := FFmpegCmdSecScreenShot(downUrl)
		if err != nil {
			log.Fatal(err)
		}
	*/
	//MergeJpg()
	err := FFmpegCmdSecRangeScreenShot(downUrl)
	if err != nil {
		log.Fatal(err)
	}
}
