package main

type Symbol uint

const (
	ROCK Symbol = iota
	PAPER
	SCISSORS
)

var (
	symbols = map[string]Symbol{
		"A": ROCK,
		"B": PAPER,
		"C": SCISSORS,
		"X": ROCK,
		"Y": PAPER,
		"Z": SCISSORS,
	}
)

type Game struct {
	player   Symbol
	opponent Symbol
}

func (game *Game) Evaluate() int {
	var winningSlice = []Symbol{ROCK, PAPER, SCISSORS}
	if game.opponent == winningSlice[(SliceIndex(winningSlice, game.player)+2)%3] {
		return 6
	}
	if game.opponent == game.player {
		return 3
	}
	return 0
}

func (game *Game) Score() (score int) {
	var winningSlice = []Symbol{ROCK, PAPER, SCISSORS}
	score += SliceIndex(winningSlice, game.player) + 1
	score += game.Evaluate()
	return
}

func CalculateScore(games []Game) (score int) {
	for _, game := range games {
		score += game.Score()
	}
	return
}

func NewGame(opponent Symbol, player string) Game {
	var winningSlice = []Symbol{ROCK, PAPER, SCISSORS}
	var indices = []string{"Y", "Z", "X"}
	return Game{opponent: opponent, player: winningSlice[(SliceIndex(winningSlice, opponent)+SliceIndex(indices, player))%3]}
}
