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

func (d *Deck) Remove(c Card) {
        idx := c.Suit * 13 + c.Rank
        d.Slots[idx] = false
}

func (d *Deck) IsInDeck(c Card) bool {
        idx := c.Suit * 13 + c.Rank
        return d.Slots[idx]
}

func (c Card) IsInDeck(d *Deck) bool {
        idx := c.Suit * 13 + c.Rank
        return d.Slots[idx]
}

func (c Card) Remove(d *Deck) {
        idx := c.Suit * 13 + c.Rank
        d.Slots[idx] = false
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
                card := Card{rand.Intn(13), rand.Intn(4)}
                if d.IsInDeck(card) {
                        d.Remove(card)
                        return card
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

