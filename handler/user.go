package handler

import (
	dblayer "cloud-storage/db"
	"cloud-storage/util"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	pwd_salt = "#*890"
)

//SignupHandler：处理用户注册请求
func SignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		data, err := ioutil.ReadFile("./static/view/signup.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(data)
		return
	}
	r.ParseForm()

	username := r.Form.Get("username")
	passwd := r.Form.Get("password")

	if len(username) < 3 || len(passwd) < 5 {
		w.Write([]byte("Invalid parameter"))
		return
	}

	enc_passwd := util.Sha1([]byte(passwd + pwd_salt))
	suc := dblayer.UserSignup(username, enc_passwd)
	if suc {
		w.Write([]byte("SUCCESS"))
	} else {
		w.Write([]byte("FAILED"))
	}
}

//SiginInHandler：登录接口
func SiginInHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.Redirect(w,r,"/static/view/signin.html",http.StatusFound)
		return
	}

	r.ParseForm()
	username := r.Form.Get("username")
	passwd := r.Form.Get("password")
	encPwd := util.Sha1([]byte(passwd + pwd_salt))

	//1.校验用户名，密码
	pwdCheck := dblayer.UserSignin(username, encPwd)
	if !pwdCheck {
		w.Write([]byte("FAILED"))
		return
	}

	//2.生成访问凭证(token)
	token := GenToken(username)
	upRes := dblayer.UpdateToken(username, token)
	if !upRes {
		w.Write([]byte("FAILED"))
		return
	}

	//3.登录成功后重定向到首页
	w.Write([]byte("http://" + r.Host + "/static/view/home.html"))
}

//GenToken：生成访问凭证（token）
func GenToken(username string) string {
	//40位字符：md5(username + timestamp + token_salt ) + timestamp[:8]
	ts := fmt.Sprintf("%x", time.Now().Unix())
	tokenPrefix := util.MD5([]byte(username + ts + "_tokensalt"))
	return tokenPrefix + ts[:8]
}
