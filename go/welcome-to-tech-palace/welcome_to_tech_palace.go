package techpalace
import(
    "strings"
)
// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	//panic("Please implement the WelcomeMessage() function")
    return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	//panic("Please implement the AddBorder() function")
    startLine := ""
    for i:=0; i<numStarsPerLine; i++{
        startLine = startLine + "*"
    }
return startLine + "\n" + welcomeMsg + "\n" + startLine
	
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	//panic("Please implement the CleanupMessage() function")
    removedStars := strings.ReplaceAll(oldMsg,"*","")
    removedSpaces := strings.TrimSpace(removedStars)
    return removedSpaces
}
