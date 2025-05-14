package main

import (
	"encoding/base64"
	"strings"
	"unicode"
)

func swap(runes []rune, start, end int, change func(rune) rune) {
	if len(runes) > 1 {
		tmp := change(runes[start])
		runes[start] = change(runes[end])
		runes[end] = tmp
	}
}

// swapCaseAndShiftDigit 处理单个字符：
// - 交换字母大小写（大写→小写，小写→大写）
// - 对数字进行循环移位（如 '5'+1 → '6'，'9'+1 → '0'）
// - 其他字符保持不变
func swapCaseAndShiftDigit(r rune, n int) rune {
	// 处理字母大小写转换
	if unicode.IsUpper(r) {
		return unicode.ToLower(r)
	}

	if unicode.IsLower(r) {
		return unicode.ToUpper(r)
	}

	// 处理数字移位
	if unicode.IsDigit(r) {
		// 将数字字符转换为0-9的数值
		digitValue := int(r - '0')
		// 执行移位操作，确保结果在0-9范围内
		newDigit := (digitValue + 10 + n) % 10
		// 转回rune类型的数字字符
		return rune('0' + newDigit)
	}

	// 非字母和数字保持不变
	return r
}

func cipher(s string) string {
	runes := []rune(s)
	change := func(r rune) rune {
		return swapCaseAndShiftDigit(r, -1)
	}
	for i := len(runes) - 5; i >= 0; i-- {
		swap(runes, i+1, i+3, change)
		swap(runes, i, i+2, change)
	}
	return string(runes)
}

func decipher(s string) string {
	runes := []rune(s)
	change := func(r rune) rune {
		return swapCaseAndShiftDigit(r, 1)
	}
	for i := 0; i <= len(runes)-5; i++ {
		swap(runes, i, i+2, change)
		swap(runes, i+1, i+3, change)
	}
	return string(runes)
}

func reverse(s string) string {
	n := []rune(s)
	for i, j := 0, len(n)-1; i < j; i, j = i+1, j-1 {
		n[i], n[j] = n[j], n[i]
	}
	return string(n)
}

func base64Decode(t string) []byte {
	decoded, _ := base64.StdEncoding.DecodeString(t)
	return decoded
}

func decode(encryptedText string) []byte {
	var result string
	if encryptedText != "" {
		result = cipher(encryptedText)
		result = reverse(result)
		l := (len(result) + 1) / 2
		result = result[l:] + result[:l]
		result = strings.Replace(result, "#", "=", 1)
		result = strings.Replace(result, "&", "==", 1)
	}
	return base64Decode(result)
}

func encode(data []byte) []byte {
	// 1. Base64 编码
	encoded := base64.StdEncoding.EncodeToString(data)

	// 2. 替换特殊字符
	encoded = strings.Replace(encoded, "=", "#", 1)
	encoded = strings.Replace(encoded, "==", "&", 1)

	// 3. 交换字符串前后部分
	l := (len(encoded) + 1) / 2
	encoded = encoded[l:] + encoded[:l]

	// 4. 反转字符串
	encoded = reverse(encoded)

	// 5. 应用 decipher 函数
	encoded = decipher(encoded)

	return []byte(encoded)
}
