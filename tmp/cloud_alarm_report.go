package main

import (
	//"fmt"
	//"html/template"
	"io/ioutil"
	"net/http"
	"net/smtp"
	//"os"
	"strings"
	"time"
)

func SendToMail(user, password, host, to, subject, body2, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + ";charset=UTF-8"
	} else {
		content_type = "Content-Type: text/" + ";charset=UTF-8"
	}
	msg := []byte("To: " + to + "\r\nFrom" + user + "\r\nSubject: " + subject + "\r\n" +
		content_type + "\r\n" + body2)
	//打印
	//fmt.Println(...)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {

	//生成client，参数默认
	//client := &http.Client{}
	//定义邮件相关参数
	user := "1220495943@qq.com"
	password := "rlchdtvcpdlcfeji"
	host := "smtp.qq.com:25"
	to := "845553422@qq.com"

	subject := "newrocktech云平台出错提醒"

	//body2 := "云平台服务器出错了，\r\n" +

	//生成要访问的url
	url := "http://api.newrocktech.com/v1/services/anylink?user_id=3839953797"

	body2 := `
			<html>
			<body>
			<h3>
			"云平台服务器出错了!"
			<h3>
			</body>
			</html>
			 `

	//无线循环
	for {
		//生成client，参数默认
		client := &http.Client{}
		//返回当前时间
		fmt.Println(time.Now())
		//提交请求
		req, err := http.NewRequest("GET", url, nil)

		if err != nil {
			panic(err)
		}

		//添加额外信息
		req.Header.Set("Authorization", "Bearer MDg4NWFiYjIt.zBiNS00ZTEyLWI5ZmItMzYxM2I4NzBlYjUy")

		//处理返回结果
		resp, _ := client.Do(req)
		defer resp.Body.Close()

		//将结果定位到标准输出
		//var stdout string
		//_, err = io.Copy(stdout, response.Body)
		//fmt.Println(stdout)
		//读取返回
		body, err := ioutil.ReadAll(resp.Body)
		httpBody := string(body)
		fmt.Println(httpBody)
		//httpBodyBak := "I am an Error"
		if strings.Contains(strings.ToLower(httpBody), "error") || resp.StatusCode != 200 {
			//body2 := "云平台服务器出错了，\r\n" + httpBodyBak
			//fmt.Println(body2)
			//t, err := template.New("newt").Parse(body2)
			//checkError(err)
			//data := struct {
			//	msg string
			//}{
			//msg: "服务器出错了",
			//}
			//err = t.Execute(os.Stdout, data)
			//checkError(err)
			//body, err := ioutil.ReadAll()
			//checkError(err)
			//body2 := "服务器出错了"
			//fmt.Println("send email")
			err = SendToMail(user, password, host, to, subject, body2, "html")
			if err != nil {
				fmt.Println("send mail error")
				fmt.Println(err)
			} else {
				fmt.Println("Send mail success!")
			}
		}

		break
		//暂停10秒
		//fmt.Println("\n开始睡60秒")
		time.Sleep(60000000000)
		//fmt.Println("睡醒了!")

		//返回的状态码
		//status := resp.StatusCode

		//fmt.Println(status)

	}
}
