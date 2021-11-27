package week02

import (
	"fmt"
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	hashMap := map[string]int{}
	for _, cpdomain := range cpdomains {
		strSlice := strings.Split(cpdomain, " ")
		num, _ := strconv.Atoi(strSlice[0])
		strs := strSlice[1]
		subStrs := strings.Split(strs, ".")
		temp := ""
		for i := len(subStrs) - 1; i >= 0; i-- {
			temp = "." + subStrs[i] + temp
			hashMap[temp[1:]] += num
		}
	}
	ans := []string{}
	for key, value := range hashMap {
		ans = append(ans, fmt.Sprintf("%v %v", strconv.Itoa(value), key))
	}
	return ans
}
