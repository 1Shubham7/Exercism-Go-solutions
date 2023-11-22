package chessboard

import (
    // "strconv"
)

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools

type File []bool

type Chessboard map[string]File

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
    count := 0
    for _,i := range cb[file] {
        if i{
            count++
    }}
	return count
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
    if rank < 1 || rank > 8 {
        return 0
    }
        count := 0
    
		for _,value := range cb{
            // since indexing starts from 0, taking length - 1
            if value[rank-1]{
                count++
            }
	}
	return count
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
    count := 0
    for _,value := range cb{
        count = count + len(value)
    }
	return count
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
        count := 0
    for _,value := range cb{
        // value is []bool
        for _,boolValue := range value{
            if boolValue{
                count++
            }
        }
    }
	return count
}
