package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	s := "hello"

	/* 文字数をカウント */
	t := len(s)                    // byte 長を返却
	t := strings.Count(s, "") - 1  // 空文字をカウント
	t := utf8.RuneCountInString(s) // マルチバイト文字列の場合

	/* 大文字に変換 */
	t := strings.ToUpper(s)

	/* 小文字に変換 */
	t := strings.ToLower(s)

	/* 各単語1文字目を大文字化 */
	t := strings.Title(s)

	/* 部分文字列を取得 */
	t := s[M:N] // M ~ (N-1)

	/* 文字列両端の空白を削除 */
	t := strings.TrimSpace(s)

	/* 特定文字列の有無 */
	t := strings.Contains(s, "") // bool

	/* 特定文字列をN回置換 */
	t := strings.Replace(s, "before", "after", N) // N = -1 で全て置換

	/* 文字列をN回羅列 */
	t := strings.Repeat(s, N)

	/* 文字列を結合 */
	slc := []string{"hoge", "fuga", "nyan"}
	t := strings.Join(slc, "/")
	fmt.Println(t) // hoge/fuga/nyan
}
