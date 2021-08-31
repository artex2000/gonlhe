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
        //test_random()
        test_showdown()
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
        for i := 0; i < 20; i += 1 {
                if i > 0 {
                        d.Reset()
                }
                h := d.GetRandomHand()
                b := d.GetRandomBoard()
                fmt.Printf("I was dealt %s on %s board\n", h.GetShortName(), b.GetShortName())
                hv := EvalHand(h, b, RIVER)
                fmt.Printf("\t%s\n", hv.GetName())
        }
}

func test_showdown() {
        d := NewDeck()
        for i := 0; i < 5; i += 1 {
                if i > 0 {
                        d.Reset()
                }
                hh, err := InputHand()
                if err != nil {
                        fmt.Println("You're not serious, bye")
                        return
                }
                hh.Remove(d)

                vh := d.GetRandomHand()
                b  := d.GetRandomBoard()

                hero := EvalHand(hh, b, RIVER)
                vill := EvalHand(vh, b, RIVER)

                r := Showdown(hero, vill)

                fmt.Printf("Hero hand: %s\n", hh.GetShortName())
                fmt.Printf("Villain hand: %s\n", vh.GetShortName())
                fmt.Printf("Board runs: %s\n", b.GetShortName())
                fmt.Printf("Hero has %s\n", hero.GetName())
                fmt.Printf("Villain has %s\n", vill.GetName())

                switch r {
                case WIN:
                        fmt.Println("Hero wins")
                case LOSE:
                        fmt.Println("Hero loses")
                case TIE:
                        fmt.Println("Hero chops")
                }
                fmt.Println("---------")
        }
}
