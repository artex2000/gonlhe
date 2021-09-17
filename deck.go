package main

import (
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

func (b *Board) Remove(d *Deck) {
        for _, v := range b.Cards {
                v.Remove(d)
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

