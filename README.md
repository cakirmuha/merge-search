# merge-search

## Merge N Intervals
### Problem statement
***Input:*** A list of intervals (two integer points)

***Function:*** Merges any N intervals in the list that have common points. 

***Output:*** Merged list (where no two intervals intersect)

***Examples:***

*Input:* {[1, 5], [2, 4], [5, 8]} - 
*Output:* {[1, 8]}

*Input:* {[1, 3], [4, 9], [5, 10]} - 
*Output:* {[1, 3], [4, 10]}
### Test

```go test interval_merge.go interval_merge_test.go -v```

### Computational complexity

Sorting time is ***O(n\*logn)***, and merging time is ***O(n)***. Computational complexity is ***O(n\*logn)***.

## Check If String 
### Single string search
***Input:*** A list of string tokens and a separate text string 

***Function:*** Accepts a list of string tokens and a separate text string as an input and checks if the latter string can be represented as a concatenation of any subset of tokens from the first argument, where each token can be used multiple times.

***Output:*** Boolean value that checks if the string can be represented as a concatenation of any subset of tokens.

***Examples:***

*Input:* {"ab", "bc", "a"} vs "abc" - 
*Output:* true

*Input:* {"ab", "bc"} vs "abc" - 
*Output:* false

#### Computational Complexity

Computational complexity is ***O(m\*n)*** where ***m*** is the length of the input string and ***n*** is the length of the list of string tokens.

### Multiple string search
***Input:*** A list of string tokens and a list of input strings to be searched

***Function:*** Accepts a list of string tokens and a list of strings to be searched as an input and checks if each string input can be represented as a concatenation of any subset of tokens from the first argument, where each token can be used multiple times.

***Output:*** Boolean value list that checks if string list can be represented as a concatenation of any subset of tokens.

***Examples:***

*Input:* {"ab", "bc", "aa"} vs {"ab", "abc", "abaaaabc", "abbcbc"} - 
*Output:* {true, false, true, true}

*Input:* {"b", "bcc", "aba"} vs {"ababccbaba", "bbccb", "ababbbcc", "abaab", "abaababbcc"} - 
*Output:* {true, true, true, false, true}

#### Computational Complexity

Computational complexity is ***O(2^m)*** where ***m*** is the length of the input string.

### Test

```go test string_single_search.go string_multiple_search.go string_sm_search_test.go -v```
