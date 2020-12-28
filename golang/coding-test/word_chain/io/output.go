package io

import (
	"fmt"
	"strings"
)

func Output(mostChain, leastRest []string) {
	// しりとりに使用した単語群
	fmt.Println(strings.Join(mostChain, " "))

	// しりとりに使用しなかった単語群
	if len(leastRest) > 0 {
		fmt.Println(strings.Join(leastRest, " "))
	}
}
