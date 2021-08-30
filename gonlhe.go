package main

import (
        "fmt"
        "math/rand"
        "time"
)

func main () {
        rand.Seed(time.Now().UnixNano())
        //test_input()
        //test_random_card()
        test_random()
}

func test_input() {
        for {
                c, err := InputCard()
                if err != nil {
                        fmt.Println("You're not serious, bye")
                        return
                }
                fmt.Printf("You've entered %s\n", c.GetShortName())
                fmt.Printf("Nice choice: %s\n", c.GetLongName())
        }
}

func test_random_card() {
        d := NewDeck()
        for i := 0; i < 5; i += 1 {
                if i > 0 {
                        d.Reset()
                }
                c := d.GetRandomCard()
                fmt.Printf("I was dealt %s (%s)\n", c.GetShortName(), c.GetLongName())
        }
}

func test_random() {
        d := NewDeck()
        for i := 0; i < 5; i += 1 {
                if i > 0 {
                        d.Reset()
                }
                h := d.GetRandomHand()
                b := d.GetRandomBoard()
                fmt.Printf("I was dealt %s on %s board\n", h.GetShortName(), b.GetShortName())
        }
}
