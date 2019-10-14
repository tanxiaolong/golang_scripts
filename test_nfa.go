package main

import (
	"fmt"
)

//
type state struct {
	symbol rune
	edge1  *state
	edge2  *state
} // state struct

// helper struct
type nfa struct {
	initial *state
	accept  *state
} // nfa struct

func poregtonfa(postfix string) *nfa {
	// an array of pointers to nfa's, the curley braces denote we want an empty one
	nfastack := []*nfa{}

	// loop throught the regular expression one rune at a time
	for _, r := range postfix {
		switch r {
		case '.':
			// get the last thing off the stack and store in frag2
			frag2 := nfastack[len(nfastack)-1]
			// get rid of the last thing on the stack, because it's already on frag2
			nfastack = nfastack[:len(nfastack)-1]
			// get the last thing off the stack and store in frag1
			frag1 := nfastack[len(nfastack)-1]
			// get rid of the last thing on the stack, because it's already on frag1
			nfastack = nfastack[:len(nfastack)-1]

			// join the fragments together by setting the accept state of frag1
			// to the initial state of frag2
			frag1.accept.edge1 = frag2.initial

			// then we append the new nfa we created above to the nfastack
			nfastack = append(nfastack, &nfa{initial: frag1.initial, accept: frag2.accept})
		case '|':
			// get the last thing off the stack and store in frag2
			frag2 := nfastack[len(nfastack)-1]
			// get rid of the last thing on the stack, because it's already on frag2
			nfastack = nfastack[:len(nfastack)-1]
			// get the last thing off the stack and store in frag1
			frag1 := nfastack[len(nfastack)-1]
			// get rid of the last thing on the stack, because it's already on frag1
			nfastack = nfastack[:len(nfastack)-1]

			// new accept state
			accept := state{}
			// the new initial state where it's two edges point to the two fragments initial states
			initial := state{edge1: frag1.initial, edge2: frag2.initial}
			// we then set the accept states of the fragments to the new accept state
			frag1.accept.edge1 = &accept
			frag2.accept.edge1 = &accept

			// then we append the new nfa accept state and initial state we created above to the nfastack
			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
		case '*':
			// get the last thing off the stack and store in frag1
			frag := nfastack[len(nfastack)-1]
			// get rid of the last thing on the stack, because it's already on frag1
			nfastack = nfastack[:len(nfastack)-1]

			accept := state{}
			// the new initial state that points to the initial of the fragment at edge1
			// and points to the new accept state at edge2
			initial := state{edge1: frag.initial, edge2: &accept}
			// join the fragment edge1 to it's initial state
			frag.accept.edge1 = frag.initial
			// join the fragment edge2 to the new accept state
			frag.accept.edge2 = &accept

			// then we append the new nfa accept state and initial state we created above to the nfastack
			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
		default:
			// new empty accept state
			accept := state{}
			// new initial state with the symbol value of r
			// and it's only edge points at the new accept state
			initial := state{symbol: r, edge1: &accept}

			// appending the new nfa to the stack
			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
		} // switch
	} // for

	// the only item will be the actual nfa that you want to return
	return nfastack[0]
} // poregtonfa

func main() {
	nfa := poregtonfa("ab.c*|")
	fmt.Printf("%+v\n", nfa)
	fmt.Printf("%+v\n", nfa.accept)
	fmt.Printf("%+v\n", nfa.initial)
} // main
