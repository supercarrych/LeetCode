package main

import (
	"fmt"
)

/*
 * 这是从控制台读取一行的帮助函数
 */
func readLine() (string, error) {
	var (
		cc  []byte
		err error
	)
	for {
		var c byte
		n, err := fmt.Scanf("%c", &c)
		if err == nil && n == 1 && c != 10 {
			cc = append(cc, c)
		} else {
			break
		}
	}
	return string(cc), err
}

/* Write Code Here */
func solution(s string) int32 {
	count := 0
	capsLockOn := false

	for _, char := range s {
		if capsLockOn && 'a' <= char && char <= 'z' {
			// 大写模式下按下小写字母，需要切换大小写模式和按下字母键，共2次按键
			count += 2
		} else if !capsLockOn && 'A' <= char && char <= 'Z' {
			// 小写模式下按下大写字母，需要切换大小写模式和按下字母键，共2次按键
			count += 2
		} else {
			// 其他情况只需要按下字母键一次
			count++
		}

		// 如果当前字符是大写字母，则切换大小写模式
		if 'A' <= char && char <= 'Z' {
			capsLockOn = !capsLockOn
		}
	}

	return int32(count)
}

func main() {
	var res int32

	var T int32
	fmt.Scanf("%d\n", &T)
	for ; T > 0; T-- {
		s, err := readLine()
		if err != nil {
			return
		}

		res = solution(s)
		fmt.Printf("%d\n", res)
	}

}
