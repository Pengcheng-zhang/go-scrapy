package common

import (
	"github.com/axgle/mahonia"
)

func convertToString(src, srcCode,tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cData, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cData)
	return result
}

func GbKToUTF8(src string) string {
	return convertToString(src, "gbk", "utf-8")
}