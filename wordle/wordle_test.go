package wordle_test

import (
	"testing"

	"github.com/hizkifw/wordle-go/wordle"
	"github.com/stretchr/testify/assert"
)

func TestCheckGuesses(t *testing.T) {
	var (
		guess wordle.Guess
		err   error
	)

	// Test different length words
	_, err = wordle.CheckGuess("test", "wow")
	assert.Error(t, err)

	// Test word not in list
	_, err = wordle.CheckGuess("aaaaa", "bbbbb")
	assert.Error(t, err)

	// Test completely wrong word
	guess, err = wordle.CheckGuess("green", "moist")
	assert.NoError(t, err)
	assert.Equal(t, "green", guess.Word)
	assert.Equal(t, wordle.GuessResult{
		wordle.GuessAbsent,
		wordle.GuessAbsent,
		wordle.GuessAbsent,
		wordle.GuessAbsent,
		wordle.GuessAbsent,
	}, guess.Result)

	// Test correct word
	guess, err = wordle.CheckGuess("world", "world")
	assert.NoError(t, err)
	assert.Equal(t, "world", guess.Word)
	assert.Equal(t, wordle.GuessResult{
		wordle.GuessCorrect,
		wordle.GuessCorrect,
		wordle.GuessCorrect,
		wordle.GuessCorrect,
		wordle.GuessCorrect,
	}, guess.Result)

	// Half-correct guesses
	guess, err = wordle.CheckGuess("elder", "wider")
	assert.Equal(t, wordle.GuessResult{
		wordle.GuessAbsent,
		wordle.GuessAbsent,
		wordle.GuessCorrect,
		wordle.GuessCorrect,
		wordle.GuessCorrect,
	}, guess.Result)

	guess, err = wordle.CheckGuess("model", "lemon")
	assert.Equal(t, wordle.GuessResult{
		wordle.GuessPresent,
		wordle.GuessPresent,
		wordle.GuessAbsent,
		wordle.GuessPresent,
		wordle.GuessPresent,
	}, guess.Result)

	guess, err = wordle.CheckGuess("mango", "lemon")
	assert.Equal(t, wordle.GuessResult{
		wordle.GuessPresent,
		wordle.GuessAbsent,
		wordle.GuessPresent,
		wordle.GuessAbsent,
		wordle.GuessPresent,
	}, guess.Result)
}
