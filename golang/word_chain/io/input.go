package io

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// 1 行目 		: 整数型にして返却する
// 2 行目以降 : スライスにして返却する
func GetInput() (int, []string) {
	firstLine, lines := getLines()
	quant, _ := strconv.Atoi(firstLine)
	inputSlc := getInputSlc(lines)

	return quant, inputSlc
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

// 全入力データを辞書順スライスで取得する
func getInputSlc(lines []string) []string {
	inputSlc := []string{}

	// i 行目
	for _, line := range lines {
		lineElmStr := strings.Split(line, " ")

		// j 列目
		for _, elmStr := range lineElmStr {
			inputSlc = append(inputSlc, elmStr)
		}
	}

	// 辞書順に並び替えておく
	sort.Strings(inputSlc)
	return inputSlc
}
