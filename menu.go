package main

// **do-this** menu
import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"os"
	//"math/rand"
)

//goland:noinspection ALL
func display_ListingOf_OptionsThese_AllHave_inCommon() {
	fmt.Println("    (using US or Alpha-Numeric input mode):")
	fmt.Println("        Enter 'menu' to return to the the main menu ")
	fmt.Println("        Enter 'dir' to redisplay this menu of available directives")
	fmt.Println("        Enter 'notes' for some background on Romaji conventions")
	fmt.Println("        Enter '?' for context-sensitive help ")
	fmt.Println("        Enter '??' for help on a particular Hiragana char")
	fmt.Println("        Enter 'set' to reset the prompt & \"key\" ")
	fmt.Println("        Enter 'stat' to view what you have done so far in the current session")
	fmt.Println("        Enter 'reset' to reset the hits logs")
	fmt.Println("        Enter 'exit' or 'quit' to terminate the app")
	//goland:noinspection ALL
	fmt.Println("\n")
}

func mainMenuPromptScanSelectAndBeginSelectedExorcise() {
	var mainMenuSelection int
	//seedFile_Maker()
	for {
		fmt.Printf("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
		fmt.Printf("  Main Menu: \n\n")
		fmt.Printf("  Enter 1 to practice recognizing Romaji: and typing Hiragana (simple: quite useful)\n")
		fmt.Printf("  Enter 2 for recognizing Romaji-Katakana pairs: typing Hiragana (somewhat easy: very useful)\n")
		fmt.Printf("  Enter 3 to practice recognizing Katakana, type either Hiragana or Romaji\n")
		fmt.Printf("  Enter 5 to practice typing drill lines\n\n\n")
		_, _ = fmt.Scan(&mainMenuSelection)

		if mainMenuSelection == 1 {
			selectedExorcise := "Romaji_Prompt"                         // 1
			display_Romaji_instructions_BeginExorcise(selectedExorcise) // 1
		} else if mainMenuSelection == 2 {
			selectedExorcise := "Romaji+Kata_Prompt"                              // 2
			display_Romaji_plus_Kata_instructions_BeginExorcise(selectedExorcise) // 2
		} else if mainMenuSelection == 3 || mainMenuSelection == 4 {
			selectedExorcise := "Kata_Prompt-Respond-w-Hira|Romaji"          // 3 or 4
			display_KataExorciseInstructions_BeginExorcise(selectedExorcise) //3 or 4
		} else if mainMenuSelection == 5 {
			drillLines() // 5
		}
	}
}

// 1
//
//goland:noinspection ALL
func display_Romaji_instructions_BeginExorcise(selectedExorcise string) {
	body_of_Romaji_instructions()
	TouchTypingExorcise(selectedExorcise)
}

//goland:noinspection ALL
func body_of_Romaji_instructions() {
	fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
	fmt.Println("Practicing touch-typing (TT) Hiragana in response to Romaji prompts:\n")
	fmt.Println("Using Hiragana-input-mode on your sys, Type the Hiragana corresponding to the Romaji prompt\n")
	display_ListingOf_OptionsThese_AllHave_inCommon() // The func is located at the end of this file
	fmt.Println("Using Hiragana-input-mode, Type the Hiragana corresponding to the Romaji prompt: \n")
}

//goland:noinspection ALL
func reDisplay_Romaji_instructions() {
	body_of_Romaji_instructions()
}

// 2
//
//goland:noinspection ALL
func display_Romaji_plus_Kata_instructions_BeginExorcise(selectedExorcise string) {
	body_of_Romaji_plus_Kata_instructions()
	TouchTypingExorcise(selectedExorcise)
}

//goland:noinspection ALL
func body_of_Romaji_plus_Kata_instructions() {
	fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
	fmt.Println("Practicing touch-typing (TT) Hiragana in response to Romaji-Katakana prompts:\n")
	fmt.Println("Using Hiragana-input-mode on your sys, Type the Hiragana corresponding to the Romaji-Katakana prompt\n")
	display_ListingOf_OptionsThese_AllHave_inCommon() // The func is located at the end of this file
	fmt.Println("Using Hiragana-input-mode, Type the Hiragana corresponding to the Romaji-Katakana prompt: \n")
}

//goland:noinspection ALL
func reDisplay_Romaji_plus_Kata_instructions() {
	body_of_Romaji_plus_Kata_instructions()
}

// 3 and 4
//
//goland:noinspection ALL
func display_KataExorciseInstructions_BeginExorcise(selectedExorcise string) {
	body_of_KataExorciseInstructions()
	TouchTypingExorcise(selectedExorcise)
}

//goland:noinspection ALL
func body_of_KataExorciseInstructions() {
	fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
	fmt.Println("Practicing recognizing Katakana chars: Using Hiragana-input-mode on your system ...")
	fmt.Println("... enter the Hiragana chars that correspond to the Katakana.\n")
	fmt.Println("Or, alternatively, type the Romaji that corresponds to the Katakana\n")
	display_ListingOf_OptionsThese_AllHave_inCommon() // The func is located at the end of this file
	fmt.Println("Type either the Hiragana or Romaji that corresponds to the Katakana prompt: \n")
}

//goland:noinspection ALL
func re_display_KataExorciseInstructions() {
	body_of_KataExorciseInstructions()
}

/*
.
.
*/

const seedFile = "randomSeed.dat"

var seed int64

//goland:noinspection ALL
func seedFile_Maker() {
	// Try to read existing seed
	if data, err := os.ReadFile(seedFile); err == nil {
		seed = int64(binary.LittleEndian.Uint64(data))
	} else {
		// No existing seed, create a new one
		err := binary.Read(rand.Reader, binary.LittleEndian, &seed)
		if err != nil {
			return
		}
		f, _ := os.Create(seedFile)
		defer func(f *os.File) {
			err := f.Close()
			if err != nil {
			}
		}(f)
		err2 := binary.Write(f, binary.LittleEndian, seed)
		if err2 != nil {
			return
		}
	}
}

// Possible other things to do after an activity and before beginning another activity
func betweenMainMenuSelectionsTTE(selectedExorcise string) {
	if selectedExorcise == "Romaji_Prompt" {
		//println("between 1 and another exorcise")
		// Save seed before exiting
		createAndWrite_seedFile()
	} else if selectedExorcise == "Romaji+Kata_Prompt" {
		//println("between 2 and another exorcise")
		// Save seed before exiting
		createAndWrite_seedFile()
	} else if selectedExorcise == "Kata_Prompt-Respond-w-Hira" {
		//println("between 3 and another exorcise")
		// Save seed before exiting
		createAndWrite_seedFile()
	} else if selectedExorcise == "Kata_Prompt-Respond-w-Romaji" {
		//println("between 4 and another exorcise")
		// Save seed before exiting
		createAndWrite_seedFile()
	}
}

func createAndWrite_seedFile() {
	f, _ := os.Create(seedFile)
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
		}
	}(f)
	err := binary.Write(f, binary.LittleEndian, seed)
	if err != nil {
		return
	}
}
