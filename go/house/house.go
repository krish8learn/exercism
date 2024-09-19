package house

const testVersion = 1

var things = []string{
	"malt",
	"rat",
	"cat",
	"dog",
	"cow with the crumpled horn",
	"maiden all forlorn",
	"man all tattered and torn",
	"priest all shaven and shorn",
	"rooster that crowed in the morn",
	"farmer sowing his corn",
	"horse and the hound and the horn",
}

var verbs = []string{
	"ate",
	"killed",
	"worried",
	"tossed",
	"milked",
	"kissed",
	"married",
	"woke",
	"kept",
	"belonged to",
}

func Song() (song string) {
	for i := 1; i <= 12; i++ {
		song += Verse(i)

		if i < 12 {
			song += "\n\n"
		}
	}

	return
}

func Verse(n int) (verse string) {
	if n == 1 {
		return "This is the house that Jack built."
	}

	verse += "This is the " + things[n-2] + "\n"

	for i := n - 3; i >= 0; i-- {
		verse += "that " + verbs[i] + " the " + things[i] + "\n"
	}

	verse += "that lay in the house that Jack built."

	return
}
