package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/golang/freetype"
	//"github.com/golang/freetype/truetype"
	//"golang.org/x/image/font"
)

func main() {
	cmd := exec.Command("/usr/local/bin/wget", "-O", "/Users/zuiyou/Desktop/avatar.png", "http://zyimg101.oss-cn-shanghai-internal.aliyuncs.com/zyimg/e6/28/505e-99d4-11e9-800a-00163e082795?x-oss-process=image/resize,m_mfit,h_208,w_208/format,png/circle,r_104")
	cmd.Run()
	defer func() {
		rm_cmd := exec.Command("/bin/rm", "-f", "/Users/zuiyou/Desktop/avatar.png")
		rm_cmd.Run()
	}()
	favimg, _ := os.Open("/Users/zuiyou/Desktop/avatar.png")
	defer favimg.Close()
	avimg, _, err := image.Decode(favimg)
	if err != nil {
		fmt.Println(err)
	}

	fbgimg, _ := os.Open("/Users/zuiyou/Desktop/helper_answer_success.png")
	defer fbgimg.Close()
	bgimg, _, _ := image.Decode(fbgimg)

	fontBytes, err := ioutil.ReadFile("/Users/zuiyou/Desktop/PingFang-SC-Bold.ttf")
	if err != nil {
		fmt.Println("er font")
		return
	}
	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		fmt.Println("er fb", err)
		return
	}
	background := image.NewRGBA(image.Rect(0, 0, 750, 1334))
	//draw.Draw(background, background.Bounds(), image.NewUniform(color.RGBA{255, 255, 255, 255}), image.ZP, draw.Src)
	draw.Draw(background, background.Bounds(), bgimg, image.ZP, draw.Src)
	draw.Draw(background, background.Bounds(), avimg, image.Pt(-271, -130), draw.Over)

	c := freetype.NewContext()
	c.SetDPI(float64(72))
	c.SetFont(f)
	c.SetFontSize(float64(40.0))
	c.SetClip(background.Bounds())
	c.SetDst(background)
	c.SetSrc(image.NewUniform(color.RGBA{0, 0, 0, 255}))

	name := string("你好你那可是打飞机啊善良的咖啡机alsd")
	if len([]rune(name)) > 14 {
		name = string([]rune(name)[:14])
		name += "..."
	}
	retp, _ := c.StringPixelLength(name)
	pt := freetype.Pt(int(750-retp)/2, 406)
	fmt.Println(retp)
	pixels, err := c.DrawString(name, pt)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("drawn ", pixels, "pixels")
	outfile, err := os.Create("/Users/zuiyou/Desktop/a.png")
	if err != nil {
		fmt.Println(err)
	}
	defer outfile.Close()

	buff := bufio.NewWriter(outfile)
	err = png.Encode(buff, background)
	if err != nil {
		fmt.Println(err)
	}
	buff.Flush()
}
