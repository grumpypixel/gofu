package stringslice

import "strings"

func RemoveElementAt(slice []int, i int) []int {
	return append(slice[:i], slice[i+1:]...)
}

func RemoveElementAtUnordered(s []int, i int) []int {
	l := len(s)
	s[l-1], s[i] = s[i], s[l-1]
	return s[:l-1]
}

func RemoveEmptyElements(s []string) []string {
	var t []string
	for _, str := range s {
		if str == "" {
			continue
		}
		t = append(t, str)
	}
	return t
}

func TrimElements(s []string) []string {
	var t []string
	for _, str := range s {
		t = append(t, strings.TrimSpace(str))
	}
	return t
}

func ElementsToLowercase(s []string) []string {
	var t []string
	for _, str := range s {
		t = append(t, strings.ToLower(str))
	}
	return t
}

func IndexOfElement(value string, s []string) int {
	for i, str := range s {
		if value == str {
			return i
		}
	}
	return -1
}
