package main

import "fmt"

func permutation(s string) []string {
	_res := []string{}
	recur([]byte(s), []byte{}, &[]string{})
	return _res
}

func recur(_b []byte, _path []byte, _res *[]string) {
	if len(_b) == 0 {
		*_res = append(*_res, string(_path))
		return
	}

	_uniq := make(map[byte]bool)
	for _e := range _b {
		if _, ok:=_uniq[_b[_e]];!ok {
			_uniq[_b[_e]] = true
			temp := make([]byte, 0)
			temp = append(temp, _b[:_e]...)
			temp = append(temp, _b[_e+1:]...)

			pathTemp := make([]byte, len(_path))
			copy(pathTemp, _path)
			pathTemp = append(pathTemp, _b[_e])
			recur(temp, pathTemp, _res)
		}
	}
	return
}

func main() {
	fmt.Println(permutation("aab"))
}