package sorting

type InputDetailType struct {
	word string // 入力文字列
	head string // 先頭 1 文字目
	tail string // 末尾 1 文字目
}

// しりとり先頭要素の候補とそれ以外の単語群を準備する
func prepareWordChain(inputSlc []string) ([]InputDetailType, []InputDetailType) {
	inputDetailSlc := getInputDetail(inputSlc)
	headSlc, preparedSlc := separateSlc(inputDetailSlc)

	return headSlc, preparedSlc
}

// 入力文字列とその先頭・末尾の 1 文字を構造体に格納する
func getInputDetail(inputSlc []string) []InputDetailType {
	var inputDetailSlc []InputDetailType

	for _, input := range inputSlc {
		headStr := string(input[0])
		tailStr := string(input[len(input)-1])
		inputDetail := InputDetailType{word: input, head: headStr, tail: tailStr}
		inputDetailSlc = append(inputDetailSlc, inputDetail)
	}

	return inputDetailSlc
}

// 先頭に配置すべき単語群とそれ以外の単語群を返却する
func separateSlc(inputDetailSlc []InputDetailType) ([]InputDetailType, []InputDetailType) {
	var headSlc []InputDetailType     // 先頭に配置すべき単語情報
	var preparedSlc []InputDetailType // それ以外の単語情報

	// vi と各要素の関係性を 2 分する
	for i, vi := range inputDetailSlc {
		var headHasTail bool // 先頭文字が前方要素に連結するか
		var tailHasHead bool // 末尾文字が後方要素に連結するか

		for j, vj := range inputDetailSlc {
			if i == j {
				continue
			}
			if vi.head == vj.tail {
				headHasTail = true
			}
			if vi.tail == vj.head {
				tailHasHead = true
			}
		}

		// vi が後方のみに連結要素を持つ場合
		// 先頭に配置すると連鎖数が伸びる
		if !headHasTail && tailHasHead {
			headSlc = append(headSlc, vi)
		} else {
			preparedSlc = append(preparedSlc, vi)
		}
	}

	return headSlc, preparedSlc
}
