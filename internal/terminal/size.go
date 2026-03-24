package terminal

import (
	"os"
	"strconv"

	"golang.org/x/term"
)

type Size struct {
	Width  int
	Height int
}

func GetSize() (Size, error) {
	fd := int(os.Stdout.Fd())

	width, height, err := term.GetSize(fd)
	if err != nil {
		return Size{Width: 80, Height: 24}, nil
	}

	if width <= 0 {
		width = 80
	}
	if height <= 0 {
		height = 24
	}

	return Size{Width: width, Height: height}, nil
}

func GetSizeFromEnv() Size {
	if w := os.Getenv("COLUMNS"); w != "" {
		if width, err := strconv.Atoi(w); err == nil && width > 0 {
			height := 24
			if h := os.Getenv("LINES"); h != "" {
				if hh, err := strconv.Atoi(h); err == nil && hh > 0 {
					height = hh
				}
			}
			return Size{Width: width, Height: height}
		}
	}

	size, _ := GetSize()
	return size
}
