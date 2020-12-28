package io

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func GetInputParam() (string, error) {
	// flagOpt := getFlagOpt()
	argOpts := getArgOpts()
	fileStr, err := getFileStr(argOpts[0]) // 予定：ループ処理
	if err != nil {
		return "", err
	}

	return fileStr, nil
}

/* フラグを指定してオプションを取得する
func getFlagOpt() *string {
	sFlag := flag.String("s", "", "Assign a string argument.") // -s "hoge"
	flag.Parse()
	// flag.NFlag() : count flag options
	return sFlag
}
*/

// フラグ指定せず引数でオプションを取得する
func getArgOpts() []string {
	flag.Parse()
	// flag.NArg() : count arg options
	argSlice := flag.Args()
	return argSlice
}

// ファイルの全行を取得
func getFileStr(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Errorf("Error: while openning file")
		return "", err
	}
	defer file.Close()

	fileByte, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Errorf("Error: while reading file")
		return "", err
	}

	return string(fileByte), nil
}
