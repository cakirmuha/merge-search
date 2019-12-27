package merge_search

const SIZE = 26

type Node struct {
	children []*Node
	isToken  bool
}

func CheckMultipleStrings(someTokens, longListOfInputs []string) []bool {
	checkList := make([]bool, len(longListOfInputs))
	validator := preprocess(someTokens)
	for i, input := range longListOfInputs {
		checkList[i] = validator(input)
	}
	return checkList
}

func preprocess(tokens []string) func(input string) bool {
	getNode := func() *Node {
		return &Node{
			make([]*Node, SIZE),
			false,
		}
	}

	// Construct tree
	root := getNode()
	for _, token := range tokens {
		tempNode := root
		for i := 0; i < len(token); i++ {
			index := token[i] - 'a'
			if tempNode.children[index] == nil {
				tempNode.children[index] = getNode()
			}
			tempNode = tempNode.children[index]
		}
		tempNode.isToken = true
	}

	var isValid func(input string) bool
	isValid = func(input string) bool {
		inputLen := len(input)
		if inputLen == 0 {
			return true
		}
		for i := 0; i < inputLen; i++ {
			if search(root, input[0:i+1]) && isValid(input[i+1:inputLen]) {
				return true
			}
		}
		return false
	}
	return isValid
}

func search(root *Node, token string) bool {
	tempNode := root
	for i := 0; i < len(token); i++ {
		index := token[i] - 'a'
		if tempNode.children[index] == nil {
			return false
		}
		tempNode = tempNode.children[index]
	}
	return tempNode != nil && tempNode.isToken
}
