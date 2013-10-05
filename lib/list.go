package xbmctoollib

import (
	"fmt"
	"path"
)

func NewAbsentNFOList() (fn MovieFunc) {
	fn = func(nfoPath string, hasNfo bool) {
		if !hasNfo {
			fmt.Println("No NFO for:", path.Base(nfoPath))
		}
	}
	return
}
