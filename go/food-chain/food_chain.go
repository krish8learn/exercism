package foodchain

import "strings"

func Verse(v int) string {
	// panic("Please implement the Verse function")
	var output string
	if animalsLines[v].animal == "fly" {
		output += strings.ReplaceAll(iknowLine, "animal", animalsLines[v].animal) + "\n"
		output += IdontknowLine
		return output
	} else if animalsLines[v].animal == "horse" {
		output += strings.ReplaceAll(iknowLine, "animal", animalsLines[v].animal) + "\n"
		output += animalsLines[v].animalLine
		return output
	} else if animalsLines[v].animal == "spider" {
		output += strings.ReplaceAll(iknowLine, "animal", animalsLines[v].animal) + "\n"
		output += animalsLines[v].animalLine + "." + "\n"
		swallowedLineWithAnimal := strings.ReplaceAll(swallowedLine, "currentAnimal", animalsLines[v].animal)
		swallowedLineWithAnimal = strings.ReplaceAll(swallowedLineWithAnimal, "previousAnimal", animalsLines[v-1].animal)
		output += swallowedLineWithAnimal + ".\n"
		output += IdontknowLine
		return output
	}
	output += strings.ReplaceAll(iknowLine, "animal", animalsLines[v].animal) + "\n"
	output += animalsLines[v].animalLine + "\n"
	for i := v; i > 1; i-- {
		swallowedLineWithAnimal := strings.ReplaceAll(swallowedLine, "currentAnimal", animalsLines[i].animal)
		swallowedLineWithAnimal = strings.ReplaceAll(swallowedLineWithAnimal, "previousAnimal", animalsLines[i-1].animal)
		if animalsLines[i-1].animal == "spider" {
			output += swallowedLineWithAnimal + strings.ReplaceAll(animalsLines[i-1].animalLine, "It", " that") + ".\n"
			continue
		}
		output += swallowedLineWithAnimal + ".\n"
	}
	output += IdontknowLine

	return output
}

func Verses(start, end int) string {
	// panic("Please implement the Verses function")\

	var output string
	for i := start; i <= end; i++ {
		if i == end {
			output += Verse(i)
			break
		}
		output += Verse(i) + "\n\n"
	}
	return output
}

func Song() string {
	// panic("Please implement the Song function")
	return Verses(1, 8)
}

type access struct {
	animal     string
	animalLine string
}

var animalsLines = map[int]access{
	1: {"fly", ""},
	2: {"spider", "It wriggled and jiggled and tickled inside her"},
	3: {"bird", "How absurd to swallow a bird!"},
	4: {"cat", "Imagine that, to swallow a cat!"},
	5: {"dog", "What a hog, to swallow a dog!"},
	6: {"goat", "Just opened her throat and swallowed a goat!"},
	7: {"cow", "I don't know how she swallowed a cow!"},
	8: {"horse", "She's dead, of course!"},
}

var iknowLine string = "I know an old lady who swallowed a animal."
var swallowedLine string = "She swallowed the currentAnimal to catch the previousAnimal"
var IdontknowLine string = "I don't know why she swallowed the fly. Perhaps she'll die."
