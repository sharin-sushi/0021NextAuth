package login

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

///https://qiita.com/CST_negi/items/3cbe2bff606a2d3afdfaの解読用コピペ

//SessionStruct.go 　　type無かったが脱字か？
type SessionInfo struct {
	UserID         interface{} //ログインしているユーザのID
	NickName       interface{} //ログインしているユーザの名前
	IsSessionAlive bool        //セッションが生きているかどうか
}

//server.go
func Signin(g *gin.Context) {

	//isExistはユーザーが存在していればtrue?
	//user
	isExist, user := HTTPRequestManager.IsLoginUserExist(g)

	fmt.Println("Login true=Success : ", isExist) //デバッグ用
	if isExist {
		SessionManager.Login(g, user)
	}
	//GetSessionInfoのコードの記載無し（github行けばありそう）
	info := SessionManager.GetSessionInfo(g) //Session情報を取得する

	server.SetHTMLTemplate(templates["index"])
	g.HTML(200, "_base.html", gin.H{
		"SessionInfo": info,
	})
}

//SessionManager.go　　セッション生成
func Login(g *gin.Context, user DBManager.DBUsers) {
	session := sessions.Default(g)
	session.Set("alive", true)
	session.Set("userID", user.ID)
	session.Set("nickName", user.NickName)
	session.Save()
}

//ログアウト(セッションを消す)
func ClearSession(g *gin.Context) {
	session := sessions.Default(g)
	session.Clear()
	session.Save()
	println("Session clear")
}
