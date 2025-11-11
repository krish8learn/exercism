package brackets

func Bracket(input string) bool {

	// empty string case
	if len(input) == 0 {
		return true
	}

	// initialize stack and bracket pairs map
	stack := []rune{}
	bracketPairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range input {
		switch char {
		// if it's an opening bracket, push onto stack
		case '(', '{', '[':
			stack = append(stack, char)
		// if it's a closing bracket, check for matching opening bracket
		case ')', '}', ']':
			if len(stack) == 0 || stack[len(stack)-1] != bracketPairs[char] {
				return false
			}
			// pop the matching opening bracket from stack
			stack = stack[:len(stack)-1]
		}
	}

	// if stack is empty, all brackets were matched
	return len(stack) == 0
}
