package main

import "fmt"

func maxNonRepeatSubstr(s string) (string, int) {
	start := 0
	maxLength := 0
	maxNonRepeatStr := ""
	maxNonRepeatStrBak := ""
	lastOccurred := make(map[rune]int)

	for i, ch := range []rune(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start { // 如果 ch 上一次出现了，并且出现在当前不重复子串中
			start = lastI + 1
			if len(maxNonRepeatStr) > len(maxNonRepeatStrBak) { // 记录最长的字串到 maxNonRepeatStrBak
				maxNonRepeatStrBak = maxNonRepeatStr
			}
			maxNonRepeatStr = maxNonRepeatStr[len(maxNonRepeatStr)-(i-start):] //根据起始位置更新 maxNonRepeatStr
		}

		// 如果 maxLength 小于当前字串的长短将更新 maxLength
		if maxLength < i-start+1 {
			maxLength = i - start + 1
		}

		lastOccurred[ch] = i          // 记录 上一次字符的出现位置
		maxNonRepeatStr += string(ch) // 将新的字符添加到 maxNonRepeatStr
	}

	// 获取最长不重复子串
	if len(maxNonRepeatStr) < len(maxNonRepeatStrBak) {
		maxNonRepeatStr = maxNonRepeatStrBak
	}

	return maxNonRepeatStr, maxLength
}

func main() {
	testData := []string{
		"abcabcbb", "bbb", "asdgetdtdd", "ab123", "b", "", "What_are_you弄啥了！",
	}

	for _, val := range testData {
		maxNonRepeatStr, maxLength := maxNonRepeatSubstr(val)
		fmt.Printf("字符串(%s)---->最大不重复字串：%s--->字串长度：%d\n", val, maxNonRepeatStr, maxLength)
	}
}
