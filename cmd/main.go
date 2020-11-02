package main

import (
	"fmt"
	"os"

	"github.com/mvo5/squashfs"
)

func main() {
	fs, err := squashfs.NewFromFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	fmt.Println("bytes used: ", fs.Sb.BytesUsed)
	fmt.Println("bytes used (rounded to 4096)", ((fs.Sb.BytesUsed/4096)+1)*4096)

	st, err := os.Stat(os.Args[1])
	if err != nil {
		panic(err)
	}
	fmt.Println("actual file size", st.Size())
}
