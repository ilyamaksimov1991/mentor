package advice

import (
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

type Advice struct {
	advices []string
}

func NewAdvice() *Advice {
	fContent, err := ioutil.ReadFile("advice.txt")
	if err != nil {
		panic(err)
	}

	return &Advice{
		advices: strings.Split(string(fContent), "\n"),
	}
}

func (a *Advice) Get() (string, error) {
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := len(a.advices) - 1
	key := rand.Intn(max-min+1) + min

	return a.advices[key], nil
}
