package leetcode

func isAnagram(s string, t string) bool {
    
    hashMap := make(map[rune]int)

    for _,val:= range s{
        hashMap[val]++
    }

    for _,val:=range t{
        hashMap[val]--
    }

    for _,val := range hashMap{
        if val != 0{
            return false
        }
    }

    return true
}