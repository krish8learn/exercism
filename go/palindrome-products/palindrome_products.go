package palindrome

import (
	"fmt"
	"sort"
)

// Define Product type here.

type Product struct {
	number         int
	// Factorizations []int
	factors        [][2]int
}

func Products(fmin, fmax int) (Product, Product, error) {
	// panic("Please implement the Products function")
	if fmin > fmax {
		return Product{}, Product{}, fmt.Errorf("fmin > fmax")
	}

	var products []Product
	var palindrome []Product

	//prepare products
	for i := fmin; i <= fmax; i++ {
		for j := fmax; j >= fmin; j-- {
			product := Product{
				factors: [][2]int{{i, j}},
				number:  i * j,
			}
			products = append(products, product)
		}
	}

	//find palindromes
	for _, prod := range products {
		if DeterminePalindrome(prod.number) {
			palindrome = append(palindrome, prod)
		}
	}

	if len(palindrome) == 0 {
		return Product{}, Product{}, fmt.Errorf("no palindromes")
	}

	//sort the palindrome
	sort.Slice(palindrome, func(i, j int) bool {
		return palindrome[i].number < palindrome[j].number
	})

	return palindrome[0],palindrome[len(palindrome)-1],  nil

}

func DeterminePalindrome(input int) bool {
	copyInput := input
	remainder := 0
	residue := 0
	for copyInput > 0 {
		remainder = copyInput % 10
		residue = (residue * 10) + remainder
		copyInput /= 10
	}
	if residue == input {
		return true
	}

	return false
}
