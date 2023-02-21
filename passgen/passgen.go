package passgen

import (
	"math/rand"
	"time"
)

func GetRandomPass() string {
	shuffleArr()

	var newPass string

	for i := 0; i < 20; i++ {
		newPass += signsArray[i]
	}

	return newPass
}

func shuffleArr() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(signsArray),
		func(i, j int) { signsArray[i], signsArray[j] = signsArray[j], signsArray[i] })
}
