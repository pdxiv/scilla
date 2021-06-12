package main

func main() {

}

type gameData struct {
	general       general
	vocabulary    []vocabularyEntry
	message       []string
	systemMessage []string
	location      []locationEntry
	object        []objectEntry
	response      []wordConditionAction
	process       []wordConditionAction
	graphics      []graphicsEntry
	udgData       []byte
	charsetData   []byte
	charset1      []byte
	patchData     patch
}

type patch struct {
	defaultPrintingLine int
	cursorCharacter     string
	promptCharacter     string
	nonInversePrompt    bool
	nonFlashingCursor   bool
}

type general struct {
	locations           int
	objects             int
	conveyableObjects   int
	messages            int
	systemMessages      int
	characterSets       int
	graphicsCount       int
	defaultInkColor     int
	defaultPaperColor   int
	defaultFlashState   bool
	defaultBrightState  bool
	defaultInverseState bool
	defaultOverState    bool
	defaultBorderColor  int
	illustratorUsed     bool
	patchUsed           bool
	databaseVersion     string
}

type vocabularyEntry struct {
	text      string
	wordIndex int
}

type locationEntry struct {
	description string
	connection  []wordAndLocationPair
}

type wordAndLocationPair struct {
	wordIndex     int
	locationIndex int
}

type objectEntry struct {
	word        int
	description string
	initiallyAt int // 252: not created, 253: worn, 254: carried
}

type wordConditionAction struct {
	verb      int
	noun      int
	condition []conditionEntry
	action    []actionEntry
}

type conditionEntry struct {
	command    int
	parameter1 int
	parameter2 int
}

type actionEntry struct {
	command    int
	parameter1 int
	parameter2 int
}

// It would be good to be able to identify each
// graphics entry somehow when using a subroutine
type graphicsEntry struct {
	isSubroutine bool
	ink          int
	paper        int
	bit6         bool
	command      []graphicsCommand
}

type graphicsCommand struct {
	command    int
	parameter1 int
	parameter2 int
	parameter3 int
	parameter4 int
}

/*
Drawing commands:
ABS MOVE \d \d
BLOCK \d \d \d \d
BRIGHT \d
BSHADE \d \d \d
END
FILL \d \d
FLASH \d
GOSUB sc=\d \d
INK \d
LINE \d \d
PAPER \d
PLOT \d \d
REL MOVE \d \d
RPLOT \d \d
SHADE \d \d \d

BSHADE i \d \d \d
LINE i \d \d
LINE o \d \d
PLOT i \d \d
PLOT o \d \d
RPLOT i \d \d
RPLOT o \d \d
*/
