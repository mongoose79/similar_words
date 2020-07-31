package utils

import (
	"sort"
	"strings"
)

func SortString(str string) string {
	if str == "" {
		return ""
	}
	res := strings.Split(str, "")
	sort.Strings(res)
	return strings.Join(res, "")
}

func IsStringInSlice(sourceStr string, listStr []string) bool {
	if listStr == nil || len(listStr) == 0 {
		return false
	}
	for _, str := range listStr {
		if str == sourceStr {
			return true
		}
	}
	return false
}

func RemoveStrFromSlice(strArr []string, strToRemove string) []string {
	var res []string
	if strToRemove == "" || strArr == nil || len(strArr) == 0 {
		return strArr
	}
	for _, str := range strArr {
		if str != strToRemove {
			res = append(res, str)
		}
	}
	return res
}
