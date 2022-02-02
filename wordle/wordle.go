package wordle

import (
	"crypto/rand"
	"errors"
	"math/big"
	"strings"
)

const (
	GuessCorrect = 'o'
	GuessPresent = '+'
	GuessAbsent  = 'x'
)

type GuessResult = []byte

type Guess struct {
	Word   string
	Result GuessResult
}

type Wordle interface {
	SubmitGuess(guess string) (GuessResult, error)
	GetGuesses() []Guess
}

type wordleImpl struct {
	guesses  []Guess
	answer   string
	maxTries int
}

func NewWordle() (Wordle, error) {
	answer, err := randomItem(validWordles)
	if err != nil {
		return nil, err
	}

	return &wordleImpl{
		guesses:  []Guess{},
		answer:   answer,
		maxTries: 6,
	}, nil
}

func randomItem(stuff map[string]struct{}) (string, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(len(stuff))))
	if err != nil {
		return "", err
	}
	n64 := n.Int64()

	for k := range stuff {
		if n64 == 0 {
			return k, nil
		}
		n64--
	}
	return "", errors.New("ran out of stuff")
}

func CheckGuess(guess string, answer string) (Guess, error) {
	// Lowercase the guess and answer
	guess = strings.TrimSpace(strings.ToLower(guess))
	answer = strings.TrimSpace(strings.ToLower(answer))

	// Make sure they're both the same length
	if len(guess) != len(answer) {
		return Guess{}, errors.New("guess must be the same length as the word")
	}

	// Make sure guess is in word list
	if _, ok := validWordles[guess]; !ok {
		if _, ok := validGuesses[answer]; !ok {
			return Guess{}, errors.New("guess must be a valid word")
		}
	}

	result := make(GuessResult, len(answer))

	// Evaluate correct and absent letters
	for i := range answer {
		if guess[i] == answer[i] {
			result[i] = GuessCorrect
		} else if !strings.Contains(answer, string(guess[i])) {
			result[i] = GuessAbsent
		}
	}

	// Evaluate present letters
	for i := range answer {
		// Skip if already evaluated
		if result[i] == GuessAbsent || result[i] == GuessCorrect {
			continue
		}

		// Go through the answer to find letters which are not
		// assigned a result yet
		for j := range answer {
			if result[j] == GuessCorrect {
				continue
			}

			// If we find the letter, mark it as present
			if guess[i] == answer[j] {
				result[i] = GuessPresent
				break
			}
		}

		// Otherwise, mark it as absent
		if result[i] != GuessPresent {
			result[i] = GuessAbsent
		}
	}

	return Guess{
		Word:   guess,
		Result: result,
	}, nil
}

func (w *wordleImpl) SubmitGuess(guess string) (GuessResult, error) {
	if len(w.guesses) >= w.maxTries {
		return nil, errors.New("you have used up all your guesses")
	}

	result, err := CheckGuess(guess, w.answer)
	if err != nil {
		return nil, err
	}

	w.guesses = append(w.guesses, result)

	return result.Result, nil
}

func (w *wordleImpl) GetGuesses() []Guess {
	return w.guesses
}
