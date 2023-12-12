package scrabble

import (
    "strings"
)

func Score(word string) int {
    var i, score int

    word = strings.ToUpper(word)
    
    for i=0;i<len(word);i++{
        switch word[i]{
            case 'A','E','I','O','U','L','N','R','S','T':
				score++
            case 'D','G':
        		score = score+2
            case 'B','C','M','P':
        		score = score + 3
            case 'F','H','V','W','Y':
        		score = score + 4
            case 'K':
        		score = score + 5
            case 'J','X':
        		score = score + 8
            case 'Q','Z':
        		score = score + 10
            default:
        		score = score
        }
    }
	return score
}
