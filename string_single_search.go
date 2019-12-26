package merge_search

func contains(tokens []string, input string) bool {
	for _, t := range tokens {
		if t == input {
			return true
		}
	}
	return false
}

func CheckIfConcatenated(tokens []string, input string) bool {
	inputLen := len(input)
	if inputLen == 0 {
		return true
	}

	visited := []int{-1}

	for i := 0; i < inputLen; i++ {
		for j := len(visited) - 1; j >= 0; j-- {
			if contains(tokens, input[visited[j]+1:i+1]) {
				if i == inputLen-1 {
					return true
				}
				visited = append(visited, i)
				break
			}
		}
	}

	return false
}
