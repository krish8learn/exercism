package cryptosquare

import (
	"fmt"
	"math"
	"regexp"
	"strings"
)

func Encode(pt string) string {
	// panic("Please implement the Encode function")

	// First, the input is normalized: the spaces and punctuation are removed, text is converted to down-cases //
	plainText := removeSpecialCharacters(pt)

	// the plain text is converted to rectangle //
	// find the length of plain text
	textLength := len(plainText)

	// determine column and row from the length of plain text
	// r * c >= length(message)
	// c >= r , c-r <= 1
	rows, cols := determineRowColumn(textLength)
	fmt.Println(rows, cols)

	// divide the plain text by column as element length //
	sliceBasedColumns := sliceMatriceConversion(plainText, textLength, cols)

	// the coded message is divided by the row as element length, transpose of previous matrice form //
	sliceBasedRows := Transpose(sliceBasedColumns)

	//slice to string
	return sliceToString(sliceBasedRows)

}

func removeSpecialCharacters(element string) string {
	input := strings.ToLower(element)
	regex, _ := regexp.Compile(`[^\w]`)
	return regex.ReplaceAllString(input, "")
}

func determineRowColumn(textLength int) (int, int) {
	rows := int(math.Sqrt(float64(textLength)))
	cols := rows
	if rows*cols < textLength {
		cols++
	}

	if rows*cols < textLength {
		rows++
	}

	return rows, cols
}

func sliceMatriceConversion(plainText string, textLength, requireElementLength int) []string {
	var returnSlice []string
	var element string
	var count int
	for index := 0; index < textLength; index++ {
		element += string(plainText[index])
		count++
		if count == requireElementLength {
			returnSlice = append(returnSlice, element)
			element = ""
			count = 0
		}

	}
	//those element who doesn't statisfy the requireElementLength
	if element != "" {
		for i := len(element); i < requireElementLength; i++ {
			element += " "
		}
		returnSlice = append(returnSlice, element)
	}

	return returnSlice
}

func sliceToString(slice []string) (output string) {

	for index, element := range slice {
		if index == len(slice)-1 {
			output += element
			break
		}
		output += element + " "
	}
	return output
}

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
