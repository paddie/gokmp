package kmp 
import (
    "fmt"
    "os"
)

type Error string

func (e Error) String() string {
    return string(e)
}

type KMP struct {
    pattern string
    prefix []int
}

func (kmp *KMP) String() string {
    return fmt.Sprintf("pattern: %v\nprefix: %v", kmp.pattern, kmp.prefix)
}

func NewKMP(pattern string) (*KMP, os.Error) {
    prefix, err := computePrefix(pattern)
    if err != nil {
        return nil, err
    }
    return &KMP{
        pattern: pattern, 
        prefix:prefix},
        nil
}

func computePrefix(pattern string) ([]int, os.Error) {
    // sanity check
    len_p := len(pattern)
    if len_p < 2 {
        if len_p == 0 {
            return nil, Error("'pattern' must contain at least one character")
        }
        return []int{-1}, nil
    }
    t := make([]int, len_p) 
    t[0], t[1] = -1, 0
    
    pos, count := 2, 0
    for pos < len_p {
        switch {
            default:
                t[pos] = 0
                pos++
            case pattern[pos - 1] == pattern[count]:
                count++
                t[pos] = count
                pos++
            case count > 0:
                count = t[count]
        }
    }
    return t, nil
}

func (kmp *KMP) FindStringIndex(s string) int {
    // sanity check
    len_p := len(kmp.pattern)
    if len(s) < len_p {
        return -1
    }  
    m, i := 0,0 
    for m+i < len(s) {
        switch {
            default:
                m = m + i - kmp.prefix[i]
                if kmp.prefix[i] > -1 {
                    i = kmp.prefix[i]
                } else {
                    i = 0
                }
            case kmp.pattern[i] == s[m+i]:
                if i == len_p-1 {
                    return m
                }
                i++
        }
    }
    return -1
}

func (kmp *KMP) ContainedIn(s string) bool {
    return kmp.FindStringIndex(s) >= 0
}

// as in strings
const startSize = 10

func (kmp *KMP) FindAllStringIndex(s string) []int {
    len_s, len_p := len(s), len(kmp.pattern)
    if len_s < len_p { return nil }
    
    match := make([]int, 0, startSize)
    m, i := 0, 0
    for m + i < len_s {
        if kmp.pattern[i] == s[m+i] {
            if i == len_p-1 {
                // the word was matched
                match = append(match,m)
                // simulate miss, and keep running
                m = m + i - kmp.prefix[i]
                if kmp.prefix[i] > -1 {
                    i = kmp.prefix[i]
                } else {
                    i = 0
                }
            } else {
                i++   
            }
        } else {
            m = m + i - kmp.prefix[i]
            if kmp.prefix[i] > -1 {
                i = kmp.prefix[i]
            } else {
                i = 0
            }  
        }
    }
    return match
}