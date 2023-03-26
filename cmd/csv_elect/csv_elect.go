package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"github.com/google/uuid"
	"github.com/mbd98/go-vote/lib/graph"
	"github.com/mbd98/go-vote/lib/methods"
	"github.com/mbd98/go-vote/lib/primitives"
	"log"
	"os"
	"strconv"
)

var schulze = flag.Bool("schulze", false, "use the Schulze Condorcet method")
var rankedpairs = flag.Bool("rankedpairs", false, "use the Ranked Pairs Condorcet method")
var irv = flag.Bool("irv", false, "use the Instant-Runoff method")
var margin = flag.Bool("margin", false, "use margins in Condorcet calculations instead of winnings")
var help = flag.Bool("help", false, "get help")

func main() {
	flag.Parse()
	if *help {
		fmt.Fprintln(os.Stderr, "usage: $ csv_elect [flags] < [input.csv]")
		fmt.Fprintln(os.Stderr, "Determine the results of an election from ballots in a CSV file")
		fmt.Fprintln(os.Stderr, "The CSV file header should list the names of the alternatives")
		fmt.Fprintln(os.Stderr, "Each subsequent row corresponds to the rankings of the corresponding column of the ballot")
		fmt.Fprintln(os.Stderr, "Flags:")
		flag.PrintDefaults()
		return
	}
	csvIn := csv.NewReader(os.Stdin)
	header, err := csvIn.Read()
	if err != nil {
		log.Fatalln(err)
	}
	body, err := csvIn.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}
	alts := make([]primitives.Alternative, len(header))
	prefs := make([]primitives.PreferentialBallot, len(body))
	for j := 0; j < len(body); j++ {
		prefs[j] = make(primitives.PreferentialBallot, len(alts))
	}
	for i, h := range header {
		alts[i] = primitives.Alternative{
			Id:   uuid.New(),
			Name: h,
		}
		for j, b := range body {
			p, err := strconv.Atoi(b[i])
			if err != nil {
				log.Fatalf("error parsing row %d column %d: %v\n", j+2, i+1, err)
			}
			prefs[j][alts[i]] = p
		}
	}
	election := graph.NewElectionGraph(alts, prefs)
	if *schulze {
		strength, winner, dom := methods.Schulze(election, *margin)
		if *margin {
			fmt.Println("\n Schulze margin results:")
		} else {
			fmt.Println("\n Schulze winning results:")
		}
		fmt.Println("Path strength matrix:")
		fmt.Print(strength.PrettyString())
		fmt.Print("Winners: ")
		for alt, win := range winner {
			if win {
				fmt.Printf("{%v} ", alt)
			}
		}
		fmt.Printf("\nDominance: %v\n", dom)
	}
	if *rankedpairs {
		if *margin {
			fmt.Println("\n Ranked pairs margin results:")
		} else {
			fmt.Println("\n Ranked pairs winning results:")
		}
		dom, order, err := methods.RankedPairs(election, *margin)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("Dominance matrix:")
		fmt.Print(dom.PrettyString())
		fmt.Printf("Topological order: %v\n", order)
	}
	if *irv {
		fmt.Println("\nInstant-Runoff results:")
		fmt.Printf("Winner: %v\n", methods.InstantRunoff(alts, prefs))
	}
}
