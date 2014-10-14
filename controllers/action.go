package controllers

import (
	"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ActionController struct {
	beego.Controller
}

func (this *ActionController) Send() {
	mess := this.GetString("mess")

	api_url := "http://www.tuling123.com/openapi/api?key=ed6c5672803204d96a1e237bef2bc1bc"
	api_url = api_url + "&info=" + url.QueryEscape(mess)
	resp, _ := http.Get(api_url)

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	html := string(body)

	this.Ctx.WriteString(html)
}
