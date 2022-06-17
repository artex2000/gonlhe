package main

import (
        "fmt"
        "math/rand"
)

func NewDeck() *Deck {
        deck := Deck{}
        deck.Reset()
        return &deck
}

func (d *Deck) Reset() {
        for i, _ := range d.Slots {
                d.Slots[i] = true
        }
}

func (c Card) IsInDeck(d *Deck) bool {
        idx := c.Suit * 13 + c.Rank
        return d.Slots[idx]
}

func (c Card) Remove(d *Deck) {
        idx := c.Suit * 13 + c.Rank
        d.Slots[idx] = false
}

func (c Card) Restore(d *Deck) {
        idx := c.Suit * 13 + c.Rank
        d.Slots[idx] = true
}

func (h *Hand) IsInDeck(d *Deck) bool {
        if h.Cards[0].IsInDeck(d) && h.Cards[1].IsInDeck(d) {
                return true
        } else {
                return false
        }
}

func (h *Hand) Remove(d *Deck) {
        for _, v := range h.Cards {
                v.Remove(d)
        }
}

func (h *Hand) Restore(d *Deck) {
        for _, v := range h.Cards {
                v.Restore(d)
        }
}

func (b *Board) Remove(d *Deck) {
        for _, v := range b.Cards {
                v.Remove(d)
        }
}

func (b *Board) Restore(d *Deck) {
        for _, v := range b.Cards {
                v.Restore(d)
        }
}

func (d *Deck) GetRandomCard() Card {
        for {
                c := Card{rand.Intn(13), rand.Intn(4)}
                if c.IsInDeck(d) {
                        c.Remove(d)
                        return c
                }
        }
        return Card{-1, -1}
}

func (d *Deck) GetRandomHand() *Hand {
        h := &Hand{}
        for i, _ := range h.Cards {
                c := d.GetRandomCard()
                if c.Rank == -1 {
                        return nil
                }
                h.Cards[i] = c
        }
        return h
}

func (d *Deck) GetRandomBoard() *Board {
        b := &Board{}
        for i, _ := range b.Cards {
                c := d.GetRandomCard()
                if c.Rank == -1 {
                        return nil
                }
                b.Cards[i] = c
        }
        return b
}

func (d *Deck) RunBoardSimulation(hero, vill *Hand) (Outcome, error) {
        o := Outcome{0, 0, 0}
        d.Reset()
        hero.Remove(d)
        if vill.IsInDeck(d) {
                vill.Remove(d)
        } else {
                return o, fmt.Errorf("Can't run sim - hero and villain hands overlap (%s %s)\n", hero.GetShortName(), vill.GetShortName())
        }

        var flop1, flop2, flop3, turn, river Card
        for f1 := 0; f1 < 48; f1 += 1 {
                flop1.Suit = f1 / 13 
                flop1.Rank = f1 % 13 
                if !flop1.IsInDeck(d) {
                        continue
                }
                flop1.Remove(d)
                for f2 := f1 + 1; f2 < 49; f2 += 1 {
                        flop2.Suit = f2 / 13 
                        flop2.Rank = f2 % 13 
                        if !flop2.IsInDeck(d) {
                                continue
                        }
                        flop2.Remove(d)
                        for f3 := f2 + 1; f3 < 50; f3 += 1 {
                                flop3.Suit = f3 / 13 
                                flop3.Rank = f3 % 13 
                                if !flop3.IsInDeck(d) {
                                        continue
                                }
                                flop3.Remove(d)
                                for t := f3 + 1; t < 51; t += 1 {
                                        turn.Suit = t / 13 
                                        turn.Rank = t % 13 
                                        if !turn.IsInDeck(d) {
                                                continue
                                        }
                                        turn.Remove(d)
                                        for r := t + 1; r < 52; r += 1 {
                                                river.Suit = r / 13 
                                                river.Rank = r % 13 
                                                if !river.IsInDeck(d) {
                                                        continue
                                                }

                                                b := &Board{}
                                                b.Cards[0] = flop1
                                                b.Cards[1] = flop2
                                                b.Cards[2] = flop3
                                                b.Cards[3] = turn
                                                b.Cards[4] = river

                                                hv := EvaluateHand(hero, b, RIVER)
                                                vv := EvaluateHand(vill, b, RIVER)

                                                /*
                                                fmt.Printf("Board: %s\n", b.GetShortName())
                                                fmt.Printf("Hero(%s): %s\n", hero.GetShortName(), hv.GetName())
                                                fmt.Printf("Villain(%s): %s\n", vill.GetShortName(), vv.GetName())
                                                */

                                                rr := Showdown(hv, vv)
                                                switch rr {
                                                case WIN:
                                                        o.Wins += 1
                                                case LOSE:
                                                        o.Loss += 1
                                                case TIE:
                                                        o.Ties += 1
                                                }
                                        }
                                        turn.Restore(d)
                                }
                                flop3.Restore(d)
                        }
                        flop2.Restore(d)
                }
                flop1.Restore(d)
        }
        return o, nil
}



