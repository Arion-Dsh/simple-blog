package utils

import (
	"bytes"
	"log"
	"sync"
	"time"

	"github.com/medivhzhan/weapp/code"
	"github.com/medivhzhan/weapp/token"
)

//WX ...
var WX = new(WXToken)

//WXToken ....
type WXToken struct {
	mutex     sync.Mutex
	Token     string
	AppID     string
	AppSecret string
	ExpAt     time.Time
}

//NewWXToken ...
func NewWXToken(id, secret string) {
	WX = new(WXToken)
	WX.AppID = id
	WX.AppSecret = secret
}

//GetToken ...
func (wx *WXToken) GetToken() string {

	now := time.Now()

	if !wx.ExpAt.IsZero() || now.Before(wx.ExpAt) {
		return wx.Token
	}
	log.Print(2)
	wx.mutex.Lock()
	tok, exp, err := token.AccessToken(wx.AppID, wx.AppSecret)
	if err != nil {
		log.Print(err.Error())
	}
	wx.Token = tok
	wx.ExpAt = now.Add(exp)
	wx.mutex.Unlock()
	return wx.Token
}

//GetQR ...
func (wx *WXToken) getQR1() {

	coder := code.QRCoder{
		Scene:     "sf13",
		Page:      "",
		Width:     400,
		IsHyaline: true,
	}

	// token: 微信 access_token

	res, err := coder.UnlimitedAppCode(wx.GetToken())
	if err != nil {
		log.Print(err.Error())
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)
	log.Print(buf.String())

	defer res.Body.Close()

}

func (wx *WXToken) getQR() {
	coder := code.QRCoder{
		Path:      "pages/index?query=1", // 识别二维码后进入小程序的页面链接
		Width:     430,                   // 图片宽度
		IsHyaline: true,                  // 是否需要透明底色
		AutoColor: true,                  // 自动配置线条颜色, 如果颜色依然是黑色, 则说明不建议配置主色调
		LineColor: code.Color{ //  AutoColor 为 false 时生效, 使用 rgb 设置颜色 十进制表示
			R: "50",
			G: "50",
			B: "50",
		},
	}

	// token: 微信 access_token
	res, err := coder.AppCode(wx.Token)
	if err != nil {
		// handle error
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)
	log.Print(buf.String())

	defer res.Body.Close()
}
