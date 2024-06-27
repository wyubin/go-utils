package vcf

import (
	"strconv"
	"strings"
)

func EscapeDataStr(str string) string {
	strNew := str
	if strings.ContainsRune(strNew, ';') {
		strNew = strings.Replace(strNew, ";", "\\x3b", -1)
	}
	if strings.ContainsRune(strNew, '=') {
		strNew = strings.Replace(strNew, "=", "\\x3d", -1)
	}
	return strNew
}

func UnEscapeDataStr(str string) string {
	strNew, _ := strconv.Unquote(`"` + str + `"`)
	return strNew
}
