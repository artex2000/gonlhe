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
        //test_showdown()
        //test_range()
        //test_board_sim()
        //test_eval()
        test_eval_verbose()
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
                hv := EvaluateHand(h, b, RIVER)
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

                hero := EvaluateHand(hh, b, RIVER)
                vill := EvaluateHand(vh, b, RIVER)

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

func test_eval() {
        d := NewDeck()
        hh := &Hand{}
        vh := &Hand{}

        for i := 0; i < 200; i += 1 {
                if i > 0 {
                        d.Reset()
                }
                hh.Cards[0].Rank = ACE
                hh.Cards[0].Suit = SPADES
                hh.Cards[1].Rank = ACE
                hh.Cards[1].Suit = CLUBS
                hh.Remove(d)

                vh.Cards[0].Rank = SEVEN
                vh.Cards[0].Suit = DIAMONDS
                vh.Cards[1].Rank = DEUCE
                vh.Cards[1].Suit = HEARTS
                vh.Remove(d)

                b  := d.GetRandomBoard()

                hero := EvaluateHand(hh, b, RIVER)
                vill := EvaluateHand(vh, b, RIVER)

                r := Showdown(hero, vill)

                if r == LOSE {
                        fmt.Printf("Board runs: %s\n", b.GetShortName())
                        fmt.Printf("Hero(%s): %s\n", hh.GetShortName(), hero.GetName())
                        fmt.Printf("Villain(%s): %s\n", vh.GetShortName(), vill.GetName())
                }

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

func test_eval_verbose() {
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
                fmt.Printf("-----> %s\n", hh.GetShortName())

                b, err := InputBoard()
                if err != nil {
                        fmt.Println("You're not serious, bye")
                        return
                }
                fmt.Printf("-----> %s\n", b.GetShortName())

                c := MergeHand(hh, b, RIVER)
                fmt.Printf("Combination \n%s\n", c.GetName())
                hv := EvaluateCombination(c)
                fmt.Printf("Hand value %s\n", hv.GetName())
        }
}


func test_range() {
        for i := 0; i < 5; i += 1 {
                r, err := InputRange();
                if err != nil {
                        fmt.Println("You're not serious, bye")
                        return
                }
                fmt.Printf("Range: %s\n", r.GetShortName())
        }
}

func test_board_sim() {
        d := NewDeck()
        for i := 0; i < 1; i += 1 {
                fmt.Println("Hero:")
                hh, err := InputHand()
                if err != nil {
                        fmt.Println("You're not serious, bye")
                        return
                }

                fmt.Println("Villain:")
                vh, err := InputHand()
                if err != nil {
                        fmt.Println("You're not serious, bye")
                        return
                }

                o, err := d.RunBoardSimulation(hh, vh)
                if err != nil {
                        fmt.Println(err)
                        continue
                }

                t := o.Wins + o.Ties + o.Loss
                fmt.Printf("Wins(%d) %.4f\n", o.Wins, float32(o.Wins)/float32(t))
                fmt.Printf("Loss(%d) %.4f\n", o.Loss, float32(o.Loss)/float32(t))
                fmt.Printf("Ties(%d) %.4f\n", o.Ties, float32(o.Ties)/float32(t))
        }
}


