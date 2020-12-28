package sorting

import (
	"sort"
)

var mostCount int      // 最長の連鎖数
var mostChain []string // 最長のしりとり単語群
var leastRest []string // 最短の残余単語群

// しりとり本処理
func WordChain(quant int, inputSlc []string) ([]string, []string) {
	headSlc, preparedSlc := prepareWordChain(inputSlc)
	countStopper := quant - len(headSlc)

	var chainCount int
	var wordChain []string
	var wordRest []string

	// 先頭要素の候補が定まらない場合
	if len(headSlc) == 0 {
		i := 0
		vi := preparedSlc[i]
		nextCompare(countStopper, chainCount, i, wordChain, wordRest, vi, preparedSlc)

		// 先頭要素の候補が定まっている場合
	} else {
		for i, vi := range headSlc {
			// 連鎖数・しりとり単語群を初期化する
			wordChain = []string{vi.word}

			// 残余単語群を初期化する
			wordRest = []string{}
			unusedHead := deleteElm(headSlc, i)
			for _, u := range unusedHead {
				wordRest = append(wordRest, u.word)
			}
			compare(countStopper, chainCount, wordChain, wordRest, vi, preparedSlc)
		}
	}

	// 残余単語群をアルファベット順にする
	sort.Strings(leastRest)

	return mostChain, leastRest
}

// しりとり単語群に要素を追加する再帰処理
// 引数 1, 2: 再帰終了条件, 連鎖数
// 引数 3, 4: しりとり単語群, 残余単語群
// 引数 5, 6: しりとり前方の単語, しりとり待機中の単語群
func compare(countStopper, chainCount int, wordChain, wordRest []string, vi InputDetailType, preparedSlc []InputDetailType) {
	if chainCount > countStopper {
		return
	}

	// 先頭・末尾の 1 文字目が同じ場合は優先する (テスト #15)
	for j, vj := range preparedSlc {
		if vi.tail != vj.head {
			continue
		}
		// "bb" などの単語が該当する
		if vj.head == vj.tail {
			nextCompare(countStopper, chainCount, j, wordChain, wordRest, vj, preparedSlc)
			break
		}
	}

	// しりとりの前方 1 文字目と後方 1 文字目を比較する
	for j, vj := range preparedSlc {
		if vi.tail != vj.head {
			continue
		}
		nextCompare(countStopper, chainCount, j, wordChain, wordRest, vj, preparedSlc)
		break
	}

	// しりとり不使用の単語を残余単語群に追加する
	for _, v := range preparedSlc {
		wordRest = append(wordRest, v.word)
	}

	// 最長連鎖の場合は記録更新する
	if chainCount > mostCount {
		mostCount = chainCount
		mostChain = wordChain
		leastRest = wordRest
	}
}

// しりとりを満たす単語発見後の処理
func nextCompare(countStopper, chainCount, j int, wordChain, wordRest []string, vj InputDetailType, preparedSlc []InputDetailType) {
	// しりとり単語群に追加する
	wordChain = append(wordChain, vj.word)

	// 追加した単語を待機列から削除する
	preparedSlc = deleteElm(preparedSlc, j)

	// 連鎖を再帰的に検証する
	chainCount++
	compare(countStopper, chainCount, wordChain, wordRest, vj, preparedSlc)
}

// スライスの要素を削除して新しいスライスを返却する
func deleteElm(slc []InputDetailType, i int) []InputDetailType {
	if i >= len(slc) {
		return slc
	}

	return append(slc[:i], slc[i+1:]...)
}
