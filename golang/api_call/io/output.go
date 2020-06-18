package io

import (
	"api_call/services"
	"fmt"
)

func Display(fmtBody services.FmtBodyType) {
	for _, body := range fmtBody {
		fmt.Println(body.Title)
	}
}
