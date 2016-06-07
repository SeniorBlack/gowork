package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	//生成client，参数默认
	client := &http.Client{}

	//生成要访问的url
	url := "http://api.newrocktech.com/v1/services/anylink?user_id=3839953797"

	//无线循环
	for {
		//返回当前时间
		fmt.Println(time.Now())
		//提交请求
		request, err := http.NewRequest("GET", url, nil)

		if err != nil {
			panic(err)
		}

		//添加额外信息
		request.Header.Set("Authorization", "Bearer MDg4NWFiYjIt.zBiNS00ZTEyLWI5ZmItMzYxM2I4NzBlYjUy")

		//处理返回结果
		response, _ := client.Do(request)

		//将结果定位到标准输出
		stdout := os.Stdout
		_, err = io.Copy(stdout, response.Body)

		//暂停10秒
		fmt.Println("\n开始睡10秒")
		time.Sleep(10000000000)
		fmt.Println("睡醒了!")

		//返回的状态码
		status := response.StatusCode

		fmt.Println(status)
	}
}
