// Qiita API を GET メソッドで叩く
// io/input.in に書き込んだ数値がクエリに設定される

package main

import (
	"api_call/io"
	"api_call/services"
	"log"
)

func main() {
	param, err := io.GetInputParam()
	loggingErr(err)

	resBody, err := services.SendRequest(param)
	loggingErr(err)

	fmtBody, err := services.FormatResponse(resBody)
	loggingErr(err)

	io.Display(fmtBody)
}

func loggingErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
