package main

// **do-this** menu
import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"os"
	"time"
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
	var mainMenuSelection string
	//seedFile_Maker()
	for {
		fmt.Printf("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
		fmt.Printf("  Main Menu: \n\n")
		fmt.Printf("  Enter '1' to practice recognizing Romaji: and typing Hiragana (simple, quite useful)\n")
		fmt.Printf("  Enter '2' for recognizing Romaji-Katakana pairs: typing Hiragana (somewhat easy, useful)\n")
		fmt.Printf("  Enter '3' to practice recognizing Katakana, type either Hiragana or Romaji (very versatile)\n")
		fmt.Printf("  Enter '5' to practice typing drill lines\n")
		fmt.Printf("  Enter '6' to practice Hiragana\n")
		fmt.Printf("  Enter '7' to practice Most-Difficult ones\n")
		fmt.Printf("  Enter 'exit' to quit\n\n\n")
		_, _ = fmt.Scan(&mainMenuSelection)

		if mainMenuSelection == "1" {
			selectedExorcise := "Romaji_Prompt" // 1
			log_to_JapLog_file_inception_time(selectedExorcise)
			display_Romaji_instructions_BeginExorcise(selectedExorcise) // 1
		} else if mainMenuSelection == "2" {
			selectedExorcise := "Romaji+Kata_Prompt" // 2
			log_to_JapLog_file_inception_time(selectedExorcise)
			display_Romaji_plus_Kata_instructions_BeginExorcise(selectedExorcise) // 2
		} else if mainMenuSelection == "3" {
			selectedExorcise := "Kata_Prompt-Respond-w-Hira|Romaji" // 3 or 4
			log_to_JapLog_file_inception_time(selectedExorcise)
			display_KataExorciseInstructions_BeginExorcise(selectedExorcise) // 3 or 4
		} else if mainMenuSelection == "5" {
			selectedExorcise := "drillLines"
			log_to_JapLog_file_inception_time(selectedExorcise)
			drillLines() // 5
		} else if mainMenuSelection == "6" {
			selectedExorcise := "Hira_prompt" // 3 or 4
			log_to_JapLog_file_inception_time(selectedExorcise)
			display_Hira_instructions_BeginExorcise(selectedExorcise) // 3 or 4
		} else if mainMenuSelection == "7" {
			selectedExorcise := "Most_Difficult" // 3 or 4
			//log_to_JapLog_file_inception_time(selectedExorcise)
			display_Difficult_instructions_BeginExorcise(selectedExorcise) // 3 or 4
		} else if mainMenuSelection == "exit" || mainMenuSelection == "quit" {
			os.Exit(1)
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

// 6
//
//goland:noinspection ALL
func display_Hira_instructions_BeginExorcise(selectedExorcise string) {
	body_of_Hira_instructions()
	TouchTypingExorcise(selectedExorcise)
}

//goland:noinspection ALL
func body_of_Hira_instructions() {
	fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
	fmt.Println("Practicing touch-typing (TT) Romaji in response to Hiragana prompts:\n")
	fmt.Println("Type the Romaji corresponding to the Hiragana prompt\n")
	display_ListingOf_OptionsThese_AllHave_inCommon() // The func is located at the end of this file
	fmt.Println("Type the Romaji corresponding to the Hiragana prompt: \n")
}

//goland:noinspection ALL
func reDisplay_Hira_instructions() {
	body_of_Romaji_instructions()
}

// 7
//
//goland:noinspection ALL
func display_Difficult_instructions_BeginExorcise(selectedExorcise string) {
	body_of_Difficult_instructions()
	TouchTypingExorciseDifficult(selectedExorcise)
}

//goland:noinspection ALL
func body_of_Difficult_instructions() {
	fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
	fmt.Println("Practicing recognizing Katakana chars: Using Hiragana-input-mode on your system ...")
	fmt.Println("... enter the Hiragana chars that correspond to the Katakana.\n")
	fmt.Println("Or, alternatively, type the Romaji that corresponds to the Katakana\n")
	display_ListingOf_OptionsThese_AllHave_inCommon() // The func is located at the end of this file
	fmt.Println("Type either the Hiragana or Romaji that corresponds to the Katakana prompt: \n")
}

//goland:noinspection ALL
func reDisplay_Difficult_instructions() {
	body_of_Difficult_instructions()
}

// Things to do after an activity, and before beginning another activity
func do_betweenMainMenuSelectionsTTE(selectedExorcise string) {
	currentTime := time.Now()
	if selectedExorcise == "Romaji_Prompt" {
		fileHandleBig, err := os.OpenFile("JapLog.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
		check(err)                                                                                 // ... gets a file handle to JapLog.txt
		//defer fileHandleBig.Close() // It’s idiomatic to defer a Close immediately after opening a file.
		_, err2 := fmt.Fprintf(fileHandleBig, "\nTransition from exorcise 1 'Romaji_Prompt' occured at: %s \n",
			currentTime.Format("15:04:05 on Monday 01-02-2006"))
		check(err2)
		//createAndWrite_seedFile()
	} else if selectedExorcise == "Romaji+Kata_Prompt" {
		fileHandleBig, err := os.OpenFile("JapLog.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
		check(err)                                                                                 // ... gets a file handle to JapLog.txt
		//defer fileHandleBig.Close() // It’s idiomatic to defer a Close immediately after opening a file.
		_, err2 := fmt.Fprintf(fileHandleBig, "\nTransition from exorcise 2 'Romaji+Kata_Prompt' occured at: %s \n",
			currentTime.Format("15:04:05 on Monday 01-02-2006"))
		check(err2)
		//createAndWrite_seedFile()
	} else if selectedExorcise == "Kata_Prompt-Respond-w-Hira|Romaji" {
		fileHandleBig, err := os.OpenFile("JapLog.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
		check(err)                                                                                 // ... gets a file handle to JapLog.txt
		//defer fileHandleBig.Close() // It’s idiomatic to defer a Close immediately after opening a file.
		_, err2 := fmt.Fprintf(fileHandleBig, "\nTransition from exorcise 3 or 4 'Kata_Prompt-Respond-w-Hira|Romaji' occured at: %s \n",
			currentTime.Format("15:04:05 on Monday 01-02-2006"))
		check(err2)
		//createAndWrite_seedFile()
	} else if selectedExorcise == "drillLines" {
		fileHandleBig, err := os.OpenFile("JapLog.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
		check(err)                                                                                 // ... gets a file handle to JapLog.txt
		//defer fileHandleBig.Close() // It’s idiomatic to defer a Close immediately after opening a file.
		_, err2 := fmt.Fprintf(fileHandleBig, "\nTransition from exorcise 5 'Drill Lines' occured at: %s \n",
			currentTime.Format("01-02-2006 15:04:05 Monday"))
		check(err2)
		//createAndWrite_seedFile()
	}
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
