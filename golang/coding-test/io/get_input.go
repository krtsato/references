package io

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// In main.go, add the following lines.
//   data := io.GetInput()
//   fmt.Printf("%v\n", data)
// Then, check output by "cat io/input.txt | go run main.go"

// 入力データ : 数字文字列
// 1 列目 		: キー
// 2 列目以降 : 構造体のフィールド値
type dataMapType = map[int]*dataMapValType

type dataMapValType struct {
	Field1 int
	Field2 int
}

// 要件に応じて返却値と型を変更する
func GetInput() dataMapType {
	firstLine, lines := getLines()

	metaData := getMetaDataSlice(firstLine)
	fmt.Printf("メタデータ: (%d, %d)\n", metaData[0], metaData[1])

	dataSlice := getDataSlice(lines)
	fmt.Println(dataSlice)

	dataMap := getDataMap(lines)
	fmt.Println(dataMap)

	return dataMap
}

// 入力データを行単位で取得する
func getLines() (string, []string) {
	scanner := bufio.NewScanner(os.Stdin)

	// 1行目
	scanner.Scan()
	checkScan(scanner)
	firstLine := scanner.Text()

	// 2行目以降
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	checkScan(scanner)

	return firstLine, lines
}

// スキャンエラーを確認する
func checkScan(scanner *bufio.Scanner) {
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

// 1行目のメタデータをスライスで取得する
func getMetaDataSlice(firstLine string) []int {
	metaDataStr := strings.Split(firstLine, " ")
	metaDataSlice := []int{}
	for _, dataStr := range metaDataStr {
		dataInt, _ := strconv.Atoi(dataStr)
		metaDataSlice = append(metaDataSlice, dataInt)
	}
	return metaDataSlice
}

// 全入力データをスライスで取得する
func getDataSlice(lines []string) []int {
	dataSlice := []int{}

	// i 行目
	for _, line := range lines {
		lineElmStr := strings.Split(line, " ")

		// j 列目
		for _, elmStr := range lineElmStr {
			elmInt, _ := strconv.Atoi(elmStr)
			dataSlice = append(dataSlice, elmInt)
		}
	}

	return dataSlice
}

// 全入力データをマップで取得する
// 1 列目 		: マップのキー
// 2 列目以降 : 構造体の各フィールド値
func getDataMap(lines []string) dataMapType {
	dataMap := dataMapType{}

	// i 行目
	for _, line := range lines {
		lineElmStr := strings.Split(line, " ")
		key := -1

		// j 列目
		for j, elmStr := range lineElmStr {
			elmInt, _ := strconv.Atoi(elmStr)

			// 列によってデータの分類を行う
			switch j % 3 {
			case 0:
				key = elmInt // 1 列目をキーに設定する
				dataMap[key] = &dataMapValType{}
			case 1:
				dataMap[key].Field1 = elmInt
			case 2:
				dataMap[key].Field2 = elmInt
			}
		}
	}

	return dataMap
}
