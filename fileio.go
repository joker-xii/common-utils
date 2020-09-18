package common_utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

func OpenFile(pos string) *bufio.Reader {
	file, err := os.Open(pos)
	if err != nil {
		panic(err)
	}
	//defer file.Close()
	reader := bufio.NewReader(file)
	return reader
}

func SaveFile(pos string) *os.File {
	//fmt.Println("open file ", pos)
	sav, err := os.OpenFile(pos,
		os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		panic(err)
	}
	return sav
}

func SaveStructAsJson(file *os.File, a ...interface{}) {
	data, _ := json.Marshal(a)
	fmt.Fprint(file, data)

}

func ReadLine(reader *bufio.Reader) (string, bool) {
	s, err := reader.ReadString('\n')
	if err == io.EOF {
		return "", false
	} else if err != nil {
		panic("fail")
	}
	if s[len(s)-1] == '\n' {
		s = s[:len(s)-1]
	}
	return s, true
}

func SaveLine(f *os.File, a ...interface{}) {
	_, err := fmt.Fprintln(f, a...)
	if err != nil {
		panic("fail")
	}
}

//func SaveLineEx(f *os.File, a Line) {
//	SaveLine(f, a.ToString())
//}

func SaveXsvLine(f *os.File, sep string, a ...interface{}) {
	as := []string{}
	for _, v := range a {
		as = append(as, fmt.Sprint(v))
	}
	SaveLine(f, strings.Join(as, sep))
}

func SaveTsvLine(f *os.File, a ...interface{}) {
	SaveXsvLine(f, "\t", a...)
}

func Use(a ...interface{}) {}

func CloneSlice(a []interface{}) []interface{} {
	if a == nil {
		return nil
	} else {
		b := make([]interface{}, len(a))
		copy(b, a)
		return b
	}
}

func CloneStrSlice(a []string) []string {
	if a == nil {
		return nil
	} else {
		b := make([]string, len(a))
		copy(b, a)
		return b
	}
}

func Fail() {
	panic("fail")
}

const (
	EQ = 1
	GT = 2
	LT = 3
	LE = 4
	GE = 5
)

func ExistsAndCmp(mp map[string]int, str string, val int, mode int, def bool) bool {
	v, ok := mp[str]
	if !ok {
		return def
	}
	if mode == GT {
		return v > val
	} else if mode == LT {
		return v < val
	} else if mode == EQ {
		return v == val
	} else if mode == LE {
		return v <= val
	} else if mode == GE {
		return v >= val
	} else {
		Fail()
		return false
	}
}

func IsDigit(r byte) bool {
	return r <= '9' && r >= '0'
}
func IsAlpha(r byte) bool {
	return IsLowerAlpha(r) || IsUpperAlpha(r)
}
func IsLowerAlpha(r byte) bool {
	return r <= 'z' && r >= 'a'
}
func IsUpperAlpha(r byte) bool {
	return r <= 'Z' && r >= 'A'
}
