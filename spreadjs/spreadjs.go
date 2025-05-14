package main

import (
	"encoding/base64"
	"fmt"
	"net/url"
	"strings"
	"unicode"
)

func s(t []rune, n, e int, i func(rune) rune) {
	if len(t) != 1 {
		t[e] = i(t[n])
		t[n] = t[e]
	}
}

func d(code int) rune {
	return rune(code)
}

func u(t rune, n int) rune {
	e := int(t)
	if 65 <= e && e <= 90 {
		return unicode.ToLower(t)
	} else if 97 <= e && e <= 122 {
		return unicode.ToUpper(t)
	} else if 48 <= e && e <= 57 {
		return d(48 + (e-48+10+n)%10)
	} else {
		return t
	}
}

func c(t string) string {
	n := []rune(t)
	e := func(t rune) rune {
		return u(t, -1)
	}
	for i := len(n) - 5; i >= 0; i-- {
		s(n, i+1, i+3, e)
		s(n, i, i+2, e)
	}
	return string(n)
}

func f(t string) string {
	n := []rune(t)
	for i, j := 0, len(n)-1; i < j; i, j = i+1, j-1 {
		n[i], n[j] = n[j], n[i]
	}
	return string(n)
}

func m(t string) string {
	decoded, _ := base64.StdEncoding.DecodeString(t)
	return string(decoded)
}

func y(t string) string {
	t = m(t)
	var res strings.Builder
	for _, r := range t {
		s := fmt.Sprintf("%%%02X", r)
		res.WriteString(s)
	}
	decoded, _ := url.QueryUnescape(res.String())
	return decoded
}

func M(t string) string {
	if t != "" {
		t = f(c(t))
		n := len(t) / 2
		t = t[n:] + t[:n]
		t = strings.Replace(t, "#", "=", 1)
		t = strings.Replace(t, "&", "==", 1)
		return y(t)
	}
	return ""
}

type SpreadJSLicense struct {
	R         int    `json:"_r"`
	H         string `json:"H"`
	Signature string `json:"S"`
}
