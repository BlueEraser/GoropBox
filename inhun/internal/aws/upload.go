package aws

import (
	"bytes"
	"fmt"
)

func AA() {
	b := bytes.NewBuffer([]byte("example object!"))
	fmt.Println(b.Len())
}
