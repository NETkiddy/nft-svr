package common

import (
	"fmt"
	"github.com/NETkiddy/common-go/log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func IsPhoneValid(phone string) bool {
	if m, _ := regexp.MatchString(`^(1[3|4|5|8][0-9]\d{4,8})$`, phone); !m {
		return false
	}
	return true
}

func Mkdir(path string) (err error) {
	if len(path) == 0 {
		log.LoggerWrapperWithCaller().Debugf("path[%v] is empty: ", path)
		return
	}
	_, err = os.Stat(path)
	if err == nil {
		log.LoggerWrapperWithCaller().Debugf("path exists: %v", path)
	} else {
		log.LoggerWrapperWithCaller().Debugf("creating path: %v", path)
		err = os.MkdirAll(path, 0755)
		if err != nil {
			log.LoggerWrapperWithCaller().Errorf("creating directory failed: %v", err)
			return
		}
	}

	return
}

// xx.xx元，转为xxxx分
func Yuan2Fen(yuan string) (fen int, err error) {
	// 合法性判定
	if strings.Count(yuan, ".") > 1 {
		err = fmt.Errorf("Yuan2Fen input not valid, %v", yuan)
		return
	}

	// format为 xx.xx
	pos := strings.IndexAny(yuan, ".")
	if pos == -1 {
		yuan += ".00"
	} else if pos+1 == len(yuan)-1 {
		yuan += "0"
	} else if pos+2 < len(yuan)-1 {
		yuan = yuan[:pos+2+1]
	}

	yuan = yuan[:len(yuan)-3] + yuan[len(yuan)-2:]
	yuan = strings.TrimPrefix(strings.Trim(yuan, " "), "0")
	fen, err = strconv.Atoi(yuan)

	return
}
