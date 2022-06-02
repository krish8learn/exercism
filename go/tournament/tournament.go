package tournament

import (
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

// I did not find this problem easy

func Tally(reader io.Reader, writer io.Writer) error {
	// panic("Please implement the Tally function")

	//reader --> line//
	buf := new(strings.Builder)
	//copy reader to buf
	_, err := io.Copy(buf, reader)
	//generate lines from input
	lines := lineGeneration(buf.String())

	//error checking//
	inputElementErr := errorChecking(lines)
	if inputElementErr != nil {
		return inputElementErr
	}

	//generate array of team scores//
	tally := generateScores(lines)

	//sorting the tally//
	tally = sortTally(tally)

	//return the ouput to writer//
	_, _ = fmt.Fprintf(writer, scoreString(tally))

	return err
}

type teamsScores struct {
	teamName    string
	matchPlayed int
	win         int
	draw        int
	loss        int
	points      int
}

func generateScores(lines []string) []teamsScores {

	var tally = make(map[string]teamsScores)
	var firstTeamName, secondTeamName, result string
	var firstTeamData, secondTeamData teamsScores

	for _, line := range lines {
		elements := strings.Split(line, ";")
		firstTeamName = elements[0]
		secondTeamName = elements[1]
		result = elements[2]

		//start recording team Details
		firstTeamData = tally[firstTeamName]
		secondTeamData = tally[secondTeamName]

		if firstTeamData.teamName == "" {
			firstTeamData = teamsScores{
				teamName: firstTeamName,
			}
		}

		if secondTeamData.teamName == "" {
			secondTeamData = teamsScores{
				teamName: secondTeamName,
			}
		}

		firstTeamData.matchPlayed++
		secondTeamData.matchPlayed++

		if result == "win" {
			firstTeamData.win++
			firstTeamData.points += 3
			secondTeamData.loss++
		} else if result == "loss" {
			firstTeamData.loss++
			secondTeamData.win++
			secondTeamData.points += 3
		} else if result == "draw" {
			firstTeamData.draw++
			firstTeamData.points++
			secondTeamData.draw++
			secondTeamData.points++
		}

		tally[firstTeamName] = firstTeamData
		tally[secondTeamName] = secondTeamData

	}
	//map to slice
	var output []teamsScores
	for _, team := range tally {
		output = append(output, team)
	}

	return output
}

func lineGeneration(input string) []string {

	//generate lines based on "\n"
	lines := strings.Split(input, "\n")
	var output []string
	for _, line := range lines {
		//trim extra space from out sides
		line = strings.TrimSpace(line)
		//absence of comments , #
		if !strings.HasPrefix(line, "#") && line != "" {
			output = append(output, line)
		}
	}
	return output
}

func errorChecking(input []string) error {

	for _, line := range input {
		//get individual element
		elements := strings.Split(line, ";")
		//check possible input
		if len(elements) != 3 {
			return errors.New("inputs must be 3")
		} else if elements[2] != "win" && elements[2] != "loss" && elements[2] != "draw" {
			return errors.New("invalid input")
		}
	}

	return nil
}

func sortTally(tally []teamsScores) []teamsScores {
	sort.Slice(tally, func(i, j int) bool {
		if tally[i].points == tally[j].points {
			return tally[i].teamName < tally[j].teamName
		}
		return tally[i].points > tally[j].points
	})
	return tally
}

func scoreString(tally []teamsScores) string {
	var output = "Team                           | MP |  W |  D |  L |  P"

	for _, team := range tally {
		output += fmt.Sprintf("\n%s| %2d | %2d | %2d | %2d | %2d",
			pad(team.teamName), team.matchPlayed, team.win, team.draw, team.loss, team.points)
	}
	return output + string(rune(10))
}

// Append whitespace to the end of the string until it is 31 characters long
func pad(s string) string {
	for len(s) < 31 {
		s += " "
	}
	return s
}
