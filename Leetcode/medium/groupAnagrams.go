package leetcode

import (
    "sort"
    "strings"
)

func groupAnagrams(strs []string) [][]string {
    anagrams := make(map[string][]string)
    for _, str := range strs {
        sortedStr := sortString(str)
        anagrams[sortedStr] = append(anagrams[sortedStr], str)
    }
    result := make([][]string, 0, len(anagrams))
    for _, v := range anagrams {
        result = append(result, v)
    }
    return result
}

func sortString(s string) string {
    arr := strings.Split(s, "")
    sort.Strings(arr)
    return strings.Join(arr, "")
}

