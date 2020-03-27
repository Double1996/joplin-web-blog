package spider

import "net/http"

/**
	https://github.com/lzjun567/zhihu-api
	https://github.com/DeanThompson/zhihu-go
    https://github.com/robertkrimen/otto  A js interpreter in go
	https://github.com/zkqiang/zhihu-login/blob/master/zhihu_login.py python zhihu login
*/

const (
	ZhiHuLoginUrl = ""
	ZHiHuUsername = ""
)

type Auth struct {
	Account  string `json:"account"`
	Password string `json:"password"`

	loginType string // phone_num 或 email
	loginURL  string // 通过 Account 判断
}

type Session struct {
	auth   *Auth
	client *http.Client
}

func NewSession() *Session {
	s := new(Session)
}

func login() {

}

func loadCookies() {

}
