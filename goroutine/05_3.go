package goroutine

import (
	"fmt"
	"time"
)

func MChannelV3_3() {
	randomNumbers := []int{}
	sum := 0
	for i := 1; i <= 1_000_000; i++ {
		randomNumbers = append(randomNumbers, i)
	}

	for i := 0; i < len(randomNumbers); i++ {
		sum += randomNumbers[i] * randomNumbers[i]
	}

	fmt.Println("Sum (MChannelV3_3): ", sum)

	defer timeTrack(time.Now(), "MChannelV3_3")
}
