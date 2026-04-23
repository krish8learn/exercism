package highscores

import "sort"

type HighScores struct {
	scores []int
}

// NewHighScores returns a new HighScores object.
func NewHighScores(scores []int) *HighScores {
	return &HighScores{scores: scores}
}

// Scores returns all the scores.
func (s *HighScores) Scores() []int {
	return s.scores
}

// Latest returns the latest (last) score.
func (s *HighScores) Latest() int {
	return s.scores[len(s.scores)-1]
}

// PersonalBest returns the best (highest) score.
func (s *HighScores) PersonalBest() int {
	best := s.scores[0]
	for _, score := range s.scores {
		if score > best {
			best = score
		}
	}
	return best
}

// TopThree returns the top three scores.
func (s *HighScores) TopThree() []int {
	// Create a copy of the scores slice to avoid modifying the original
	scores := make([]int, len(s.scores))
	copy(scores, s.scores)

	sort.Ints(scores) // Sort the scores in ascending order
	
	// Reverse the sorted scores to get them in descending order
	for i, j := 0, len(scores)-1; i < j; i, j = i+1, j-1 {
		scores[i], scores[j] = scores[j], scores[i]
	}
	// Return the first three scores (or all scores if there are fewer than three)
	if len(scores) > 3 {
		return scores[:3]
	}
	return scores
}
