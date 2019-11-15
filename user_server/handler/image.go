package handler

import (
	"context"
	"fmt"
	"github.com/afocus/captcha"
	"github.com/astaxie/beego/logs"
	"image/color"
	"io/ioutil"
	img "myproject/user_server/proto/imgpoto"
)
type Img struct {}


func(i *Img) Call(ctx context.Context, req *img.Request, rsp *img.Response) error{
	cap1 := captcha.New()
	err :=cap1.SetFont("./486.ttf")
	if err != nil{
		logs.Error(err)
	}
	cap1.SetSize(91,41)
	cap1.SetDisturbance(captcha.MEDIUM)
	//cap1.SetFrontColor(color.RGBA{255,255,255,255})
	cap1.SetBkgColor(color.RGBA{255, 0, 0, 255}, color.RGBA{0, 0, 255, 255}, color.RGBA{0, 153, 0, 255})
	img_1,str := cap1.Create(4,captcha.NUM)
	fmt.Println(str)
	b := *img_1
	c := *(b.RGBA)
	data := []byte(c.Pix)
	rsp.Pix = data
	rsp.Stride = int64(c.Stride)
	rsp.Min = &img.Response_Point{X:int64(c.Rect.Min.X),Y:int64(c.Rect.Min.Y)}
	rsp.Man = &img.Response_Point{X:int64(c.Rect.Max.X),Y:int64(c.Rect.Max.Y)}
	if ioutil.WriteFile("a.png",data,0644) == nil {
		fmt.Println("成功")
	}
	return nil
}

