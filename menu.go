package main

// **do-this**
import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"os"
	//"math/rand"
)

/*
add a cyclic array to the logging of misses, and use it every few times to practice error-prone ones
and, another to store frequently gotten, and gotten right; then skip those more often
add a variety of fonts ??
develop a method for deploying hints for naked romaji using just one of the four card lines
... finish hints lines in main char struct (one line must be reserved for naked romaji hints)
? eventually, delete unused files hint4kata.go(noUsages); and singleQuestionMarksHints.go after salvaging hints etc.
*/

// then, **do-this** these:
// try refining drill lines, Kata, Hira, and Romaji (respond to each with Hiragana)
// start adding actual Japanese words (with Kanji) to the menus etc.
// build a struct and a file of cards with kanji, both words and phrases to implement an activity similar to 1-4
/*
// 46 15 25 21 = 107 chars from the blue card if we include all the suffixed combination "chars".
46*2=92 for only the complete simple hiragana plus katakana set of base chars.
// ... doubling it all (the 107 from above) for katakana gives a grand total of 214 Japanese chars excluding the various punctuation characters.
*/

func mainMenuPromptScanSelectAndBeginSelectedExorcise() {
	var mainMenuSelection int
	seedFileMaker()
	for {
		fmt.Printf("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
		fmt.Printf("  Main Menu: \n\n")
		fmt.Printf("  Enter 1 to practice recognizing Romaji: and typing Hiragana (simple: quite useful)\n")
		fmt.Printf("  Enter 2 for recognizing Romaji-Katakana pairs: typing Hiragana (somewhat easy: very useful)\n")
		fmt.Printf("  Enter 3 to practice recognizing Katakana, and typing Hiragana (fairly easy)\n")
		fmt.Printf("  Enter 4 to practice recognizing Katakana, and typing Romaji (more challenging)\n")
		fmt.Printf("  Enter 5 to practice typing drill lines\n\n\n")

		_, _ = fmt.Scan(&mainMenuSelection)

		if mainMenuSelection == 1 {
			selectedExorcise := "RomajiNakedPrompt"
			displayNakedRomajiOptionsAndBeginExorcise(selectedExorcise)
		} else if mainMenuSelection == 2 {
			selectedExorcise := "RomajiKataPrompt"
			displayRomajiKataOptionsAndBeginExorcise(selectedExorcise)
		} else if mainMenuSelection == 3 {
			selectedExorcise := "KataNakedPrompt"
			displayNakedKataOptionsAndBeginExorcise(selectedExorcise)
		} else if mainMenuSelection == 4 {
			selectedExorcise := "RespondWithRomajiToNakedKata"
			displayNakedKataOptionsAndBeginRomajiExorcise(selectedExorcise)
		} else if mainMenuSelection == 5 {
			drillLines()
		}
	}
}

const seedFile = "randomSeed.dat"

var seed int64

func seedFileMaker() {
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
	if selectedExorcise == "RomajiNakedPrompt" {
		//println("between 1 and another exorcise")
		// Save seed before exiting
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
	} else if selectedExorcise == "RomajiKataPrompt" {
		//println("between 2 and another exorcise")
		// Save seed before exiting
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
	} else if selectedExorcise == "KataNakedPrompt" {
		//println("between 3 and another exorcise")
		// Save seed before exiting
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
	} else if selectedExorcise == "RespondWithRomajiToNakedKata" {
		//println("between 4 and another exorcise")
		// Save seed before exiting
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
}

func displayCommonOptions() {
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

//goland:noinspection ALL
func displayNakedKataOptionsAndBeginExorcise(selectedExorcise string) {
	fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
	fmt.Println("Practicing recognizing naked Katakana chars, using touch-typed (TT) Hiragana chars:\n")
	fmt.Println("Using Hiragana-input-mode on your system, Type the Hiragana that corresponds to the Katakana prompt\n")
	displayCommonOptions()
	fmt.Println("Using Hiragana-input-mode, Type the Hiragana that corresponds to the Katakana prompt: \n")
	TouchTypingExorcise(selectedExorcise)
}

//goland:noinspection ALL
func displayNakedKataOptionsAndBeginRomajiExorcise(selectedExorcise string) {
	fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
	fmt.Println("Practicing recognizing naked Katakana chars, using touch-typed (TT) Romaji:\n")
	fmt.Println("Using Alpha-input-mode on your system, Type the Romaji that corresponds to the Katakana prompt\n")
	displayCommonOptions()
	fmt.Println("Using Alpha-input-mode, Type the Romaji that corresponds to the Katakana prompt: \n")
	TouchTypingExorcise(selectedExorcise)
}
func reDisplayNakedKataOptions() {
	fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
	fmt.Println("Practicing recognizing naked Katakana chars, using touch-typed (TT) Hiragana chars:\n")
	fmt.Println("Using Hiragana-input-mode on your system, Type the Hiragana that corresponds to the Katakana prompt\n")
	displayCommonOptions()
	fmt.Println("Using Hiragana-input-mode, Type the Hiragana that corresponds to the Katakana prompt: \n")
}

func displayRomajiKataOptionsAndBeginExorcise(selectedExorcise string) {
	fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
	fmt.Println("Practicing touch-typing (TT) Hiragana in response to Romaji-Katakana prompts:\n")
	fmt.Println("Using Hiragana-input-mode on your sys, Type the Hiragana corresponding to the Romaji-Katakana prompt\n")
	displayCommonOptions()
	fmt.Println("Using Hiragana-input-mode, Type the Hiragana corresponding to the Romaji-Katakana prompt: \n")
	TouchTypingExorcise(selectedExorcise)
}
func reDisplayRomajiKataOptions() {
	fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
	fmt.Println("Practicing touch-typing (TT) Hiragana in response to Romaji-Katakana prompts:\n")
	fmt.Println("Using Hiragana-input-mode on your sys, Type the Hiragana corresponding to the Romaji-Katakana prompt\n")
	displayCommonOptions()
	fmt.Println("Using Hiragana-input-mode, Type the Hiragana corresponding to the Romaji-Katakana prompt: \n")
}

func displayNakedRomajiOptionsAndBeginExorcise(selectedExorcise string) {
	fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
	fmt.Println("Practicing touch-typing (TT) Hiragana in response to Naked Romaji prompts:\n")
	fmt.Println("Using Hiragana-input-mode on your sys, Type the Hiragana corresponding to the Romaji prompt\n")
	displayCommonOptions()
	fmt.Println("Using Hiragana-input-mode, Type the Hiragana corresponding to the Romaji prompt: \n")
	TouchTypingExorcise(selectedExorcise)
}
func reDisplayNakedRomajiOptions() {
	fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
	fmt.Println("Practicing touch-typing (TT) Hiragana in response to Naked Romaji prompts:\n")
	fmt.Println("Using Hiragana-input-mode on your sys, Type the Hiragana corresponding to the Romaji prompt\n")
	displayCommonOptions()
	fmt.Println("Using Hiragana-input-mode, Type the Hiragana corresponding to the Romaji prompt: \n")
}
