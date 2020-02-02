package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	/* 確認 : % cat << EOF | go run go-stdio.go */
	scanner := bufio.NewScanner(os.Stdin)

	/* ===== 1行1データ -> 1行1データ ===== */
	scanner.Scan()
	fmt.Println(scanner.Text())
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	/* ===== N行1データ -> N行1データ ===== */
	for scanner.Scan() {
		fmt.println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	/* ===== 行数 + N行1データ -> N行1データ ===== */
	scanner.Scan() // 1行目
	lineCount, _ := strconv.Atoi(scanner.Text())
	for i := 0; scanner.Scan() && i < lineCount; i++ {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	/* ===== 1行Nデータ -> N行1データ ===== */
	scanner.Scan()
	slcData := strings.Split(scanner.Text(), " ")
	for i := 0; i < len(slcData); i++ {
		fmt.Println(slcData[i])
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	// 別解 : スキャン関数を bufio.ScanWords に変更
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	/* ===== 行数 + 1行Nデータ -> N行1データ ===== */
	scanner.Scan() // 1行目
	lineCount, _ := strconv.Atoi(scanner.Text())
	scanner.Scan() // 2行目
	slcData := strings.Split(scanner.Text(), " ")
	for i := 0; i < lineCount; i++ {
		fmt.Println(slcData[i])
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	/* ===== 行数(h) + 列数(w) + N行Nデータ -> 行列成分 + N^2行1データ ===== */
	scanner.Scan() // 1行目
	slcHw := strings.Split(scanner.Text(), " ")
	lineCount, _ := strconv.Atoi(slcHw[0])
	clmnCount, _ := strconv.Atoi(slcHw[1])
	for i := 0; i < lineCount && scanner.Scan(); i++ {
		slcData := strings.Split(scanner.Text(), " ")
		for j := 0; j < clmnCount; j++ {
			fmt.Printf("(行, 列) = (%d, %d) | %s\n", i, j, slcData[j])
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
