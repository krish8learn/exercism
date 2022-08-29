package transpose

func Transpose(input []string) []string {
	// panic("Please implement the Transpose function")

	//empty input string
	if len(input) == 0 {
		return []string{}
	}

	//row count
	rowCount := len(input)
	//determine the lenght output string slice
	lengthOutputSlice := 0
	for _, value := range input {
		if len(value) > lengthOutputSlice {
			lengthOutputSlice = len(value)
		}
	}

	output := make([]string, lengthOutputSlice)

	for i := 0; i < rowCount; i++ {
		row := input[i]
		output[0] += string(row[0])
		for j := 1; j < len(row); j++ {
			for len(output[j]) != len(output[j-1])-1 {
				output[j] += " "
			}
			output[j] += string(row[j])
		}
	}
	return output
}
