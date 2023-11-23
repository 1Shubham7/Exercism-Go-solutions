package logs

import (
    "strings"
    "unicode/utf8"
    )



// runes must be under single  quotes ''
// myRune := '¿'
// fmt.Printf("myRune value: %v\n", myRune)
// Output: myRune value: 191
// fmt.Printf("myRune Unicode character: %c\n", myRune)
// Output: myRune Unicode character: ¿
// a string in Go is often referred to as a sequence of runes. However, runes are stored as 1, 2, 3, or 4 bytes depending on the character. Due to this, strings are really just a sequence of bytes. Even though a string is just a slice of bytes, the range keyword iterates over a string's runes, not its bytes.

// Application identifies the application emitting the given log.
func Application(log string) string {
    // var uni rune
    for _,value := range log{

        switch value{
            case '❗':
        		return "recommendation"
            case '🔍':
        		return "search"
            case '☀':
        		return "weather"
            // the mistake I was making was that I was adding a default
        }
    }
return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	log = strings.ReplaceAll(log, string(oldRune), string(newRune))
	return log
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
    if utf8.RuneCountInString(log) <= limit{
        return true
    } else {
    return false
    } 
}
