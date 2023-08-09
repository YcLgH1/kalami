package libs

import (
	"fmt"

	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/encoding/simplifiedchinese"
)

type Charset string

const (
	UTF8      = Charset("UTF-8")
	GB18030   = Charset("GB18030")
	ISO8859_1 = Charset("ISO8859_1")
)

func Byte2String(b []byte, charset Charset) string {
	var str string

	switch charset {
	case GB18030:
		decideBytes, _ := simplifiedchinese.GB18030.NewDecoder().Bytes(b)
		str = string(decideBytes)
	case ISO8859_1:
		encoder := charmap.ISO8859_1.NewEncoder()
		decideBytes, err := encoder.Bytes(b)
		fmt.Println(err)
		str = string(decideBytes)
	case UTF8:
		fallthrough
	default:
		str = string(b)
	}

	return str
}
