package electionday

import (
    // "fmt"
)

// NewVoteCounter returns a new vote counter with
// a given number of initial votes.
func NewVoteCounter(initialVotes int) *int {
    x := &initialVotes
    return x
}

// VoteCount extracts the number of votes from a counter.
func VoteCount(counter *int) int {
    if (counter==nil){
        return 0
    } else {
    		return *counter
    }
}

// IncrementVoteCount increments the value in a vote counter.
func IncrementVoteCount(counter *int, increment int) {

    *counter = *counter + increment
    
}

// NewElectionResult creates a new election result.
func NewElectionResult(candidateName string, votes int) *ElectionResult {
	panic("Please implement the NewElectionResult() function")
}

// DisplayResult creates a message with the result to be displayed.
func DisplayResult(result *ElectionResult) string {
	panic("Please implement the DisplayResult() function")
}

// DecrementVotesOfCandidate decrements by one the vote count of a candidate in a map.
func DecrementVotesOfCandidate(results map[string]int, candidate string) {
	panic("Please implement the DecrementVotesOfCandidate() function")
}
