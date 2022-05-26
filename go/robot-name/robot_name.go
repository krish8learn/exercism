package robotname

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// Define the Robot type here.
type Robot struct {
	name string
	// number int
	// str    string
}

var robots = make(map[string]bool)

//Now if a constant seed is set, it will output the same numbers. So, we need to make it a variable seed which changes after each call. Using time is a way to do it.

func init() {
	rand.Seed(time.Now().UnixNano())
}

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandomStringGenerator() string {
	length := len(alphabet)
	var sb strings.Builder

	for i := 0; i < 2; i++ {
		c := alphabet[rand.Intn(length)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomIntGenerator() string {
	randValue := 100 + int(rand.Int63n(int64(999-100+1)))
	// randValue := rand.Intn(10)
	return strconv.Itoa(randValue)
}

func (r *Robot) Name() (string, error) {
	// panic("Please implement the Name function")
	if r.name == "" {
		// r.name = RandomStringGenerator() + RandomIntGenerator() + RandomIntGenerator() + RandomIntGenerator()
		//genrating the name
		genName := RandomStringGenerator() + RandomIntGenerator()
		//checking whether the name already exits or not
		if _, exists := robots[genName]; exists {
			//exist, create new name so call again
			r.Name()
		}
		//new name , save it and assign to robot
		robots[genName] = true
		r.name = genName
	}

	return r.name, nil
}

func (r *Robot) Reset() {
	// panic("Please implement the Reset function")
	r.name = ""
}
