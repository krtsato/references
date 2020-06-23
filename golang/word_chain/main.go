package main

import (
	"word_chain/io"
	"word_chain/sorting"
)

func main() {
	quant, inputSlc := io.GetInput()
	mostChain, leastRest := sorting.WordChain(quant, inputSlc)
	io.Output(mostChain, leastRest)
}
