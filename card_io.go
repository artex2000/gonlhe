package main

import (
        "fmt"
        "strings"
)

func InputCard() (Card, error) {
        var card Card
        try := 0
        for {
                if try > 3 {
                        return Card{0, 0}, fmt.Errorf("Three strikes, you're out")
                }
                try += 1

                var c string
                fmt.Print("Enter two letter card (rank+suit, ranks 2-9,T,J,Q,K,A, suits S,H,D,C): ")
                fmt.Scanln(&c)
                if len(c) != 2 {
                        fmt.Println("Invalid input - too many letters")
                        continue
                }

                c = strings.ToUpper(c)

                found := false
                for i, v := range CardRankShortName {
                        if rune(c[0]) == v {
                                card.Rank = i
                                found = true
                        }
                }
                if !found {
                        fmt.Printf("Invalid rank %c. Valid ranks are 2-9,T,J,Q,K,A\n", c[0])
                        continue
                }

                switch c[1] {
                case 'S':
                        card.Suit = SPADES
                case 'H':
                        card.Suit = HEARTS
                case 'D':
                        card.Suit = DIAMONDS
                case 'C':
                        card.Suit = CLUBS
                default:
                        fmt.Printf("Invalid suit %c. Valid suits are S,H,D,C\n", c[1])
                        continue
                }

                break
        }

        return card, nil
}

func (c Card) GetShortName() string {
        if c.Rank == -1 {
                return "Something went wrong"
        }
        return fmt.Sprintf("%c%c", CardRankShortName[c.Rank], CardSuitShortName[c.Suit])
}

func (c Card) GetLongName() string {
        if c.Rank == -1 {
                return "Something went wrong"
        }
        return fmt.Sprintf("%s of %s", CardRankLongName[c.Rank], CardSuitLongName[c.Suit])
}

func (h *Hand) GetShortName() string {
        return fmt.Sprintf("%s %s", h.Cards[0].GetShortName(), h.Cards[1].GetShortName())
}

func (b *Board) GetShortName() string {
        var sb strings.Builder
        for i, v := range b.Cards {
                if i > 0 {
                        sb.WriteRune(' ')
                }
                sb.WriteString(v.GetShortName())
        }
        return sb.String()
}
