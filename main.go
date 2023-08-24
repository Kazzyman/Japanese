package main

import "math/rand"

func main() {
	for {
		rand.Seed(seed)
		mainMenuPromptScanSelectAndBeginSelectedExorcise()
	}
}
