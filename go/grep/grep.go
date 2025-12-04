package grep

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Search(pattern string, flags, files []string) []string {
	var results []string

	// Handle edge cases
	switch {
	case len(files) == 0:
		return []string{}
	case pattern == "":
		return []string{}
	}

	// Parse flags
	hasN := contains(flags, "-n")
	hasL := contains(flags, "-l")
	hasI := contains(flags, "-i")
	hasV := contains(flags, "-v")
	hasX := contains(flags, "-x")

	// -l flag takes precedence
	if hasL {
		return Lcase(pattern, files)
	}

	patternLower := strings.ToLower(pattern)

	for _, fileName := range files {
		openFile, err := os.Open(fileName)
		if err != nil {
			panic(err)
		}
		defer openFile.Close()

		scanner := bufio.NewScanner(openFile)
		lineNumber := 1

		for scanner.Scan() {
			line := scanner.Text()
			lineLower := strings.ToLower(line)

			// Determine if line matches based on flags
			isMatch := false

			if hasX {
				// Exact line match
				if hasI {
					isMatch = (lineLower == patternLower)
				} else {
					isMatch = (line == pattern)
				}
			} else {
				// Substring match
				if hasI {
					isMatch = strings.Contains(lineLower, patternLower)
				} else {
					isMatch = strings.Contains(line, pattern)
				}
			}

			// Apply inverted flag
			if hasV {
				isMatch = !isMatch
			}

			if isMatch {
				// Format output
				var output string
				if len(files) > 1 {
					if hasN {
						output = fmt.Sprintf("%s:%d:%s", fileName, lineNumber, line)
					} else {
						output = fmt.Sprintf("%s:%s", fileName, line)
					}
				} else {
					if hasN {
						output = fmt.Sprintf("%d:%s", lineNumber, line)
					} else {
						output = line
					}
				}
				results = append(results, output)
			}

			lineNumber++
		}
	}

	return results
}

func contains(arr []string, target string) bool {
	for _, v := range arr {
		if v == target {
			return true
		}
	}
	return false
}

func Ncase(pattern string, files []string) []string {
	var results []string
	for _, fileName := range files {
		openFile, err := os.Open(fileName)
		if err != nil {
			panic(err)
		}
		defer openFile.Close()

		scanner := bufio.NewScanner(openFile)
		lineNumber := 1
		for scanner.Scan() {
			line := scanner.Text()
			if strings.Contains(line, pattern) {
				results = append(results, fmt.Sprintf("%d:%s", lineNumber, line))
			}
			lineNumber++
		}
	}
	return results
}

func Lcase(pattern string, files []string) []string {

	var results []string
	for _, fileName := range files {
		openFile, err := os.Open(fileName)
		if err != nil {
			panic(err)
		}
		defer openFile.Close()

		scanner := bufio.NewScanner(openFile)
		present := false
		for scanner.Scan() {
			line := scanner.Text()
			if strings.Contains(line, pattern) && !present {
				results = append(results, fileName)
				present = true
				break // exit loop after first match
			}
		}
	}
	return results
}

func Icase(pattern string, files []string) []string {
	var results []string
	patternLower := strings.ToLower(pattern)
	for _, fileName := range files {
		openFile, err := os.Open(fileName)
		if err != nil {
			panic(err)
		}
		defer openFile.Close()

		scanner := bufio.NewScanner(openFile)
		for scanner.Scan() {
			line := scanner.Text()
			lineLower := strings.ToLower(line)
			if strings.Contains(lineLower, patternLower) {
				results = append(results, line)
			}
		}
	}
	return results
}

func Vcase(pattern string, files []string) []string {
	var results []string
	for _, fileName := range files {
		openFile, err := os.Open(fileName)
		if err != nil {
			panic(err)
		}
		defer openFile.Close()

		scanner := bufio.NewScanner(openFile)
		for scanner.Scan() {
			line := scanner.Text()
			if !strings.Contains(line, pattern) {
				results = append(results, line)
			}
		}
	}
	return results
}

func Xcase(pattern string, files []string) []string {
	var results []string
	for _, fileName := range files {
		openFile, err := os.Open(fileName)
		if err != nil {
			panic(err)
		}
		defer openFile.Close()

		scanner := bufio.NewScanner(openFile)
		for scanner.Scan() {
			line := scanner.Text()
			if line == pattern {
				results = append(results, line)
			}
		}
	}
	return results
}

func NonCase(pattern string, files []string) []string {
	var results []string
	for _, fileName := range files {
		openFile, err := os.Open(fileName)
		if err != nil {
			panic(err)
		}
		defer openFile.Close()

		scanner := bufio.NewScanner(openFile)
		for scanner.Scan() {
			line := scanner.Text()
			if strings.Contains(line, pattern) {
				results = append(results, line)
			}
		}
	}
	return results
}
