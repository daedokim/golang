package utils

import "time"

// GetCurrentTick 현재 tick을 세팅한다.
func GetCurrentTick() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// GetIntArrayIndexOf 정수형 배열의 해당 엘리먼트의 인덱스를 찾는다.
func GetIntArrayIndexOf(arr []int, element int) int {
	var index int
	for i := 0; i < len(arr); i++ {
		if arr[i] == element {
			index = i
			break
		}
	}
	return index
}
