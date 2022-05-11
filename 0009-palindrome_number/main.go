package main

import "strconv"

func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }
    if x < 10 {
        return true
    }
    ans := 0
    i := x
    var p, left int
    for i > 0 {
        left = i / 10
        p = i - left * 10
        ans = ans * 10 + p
        i = left
    }
    return ans == x
}

func isPalindromeStr(x int) bool {
    input := strconv.Itoa(x)
    i := 0
    j := len(input) - 1
    for i < j {
        if input[i] != input[j] {
            return false
        }
        i ++
        j --
    }
    return true
}

func isPalindromeBetter(x int) bool{
    if x < 0 || x % 10 == 0 {
        return false
    }
    res := 0 
    for x > res {
        res = res * 10 + x % 10
        x /= 10
    }
    return x == res || x == (res / 10)
}



