package main

import (
        "fmt"
        "strings"
)

func InputCard() (Card, error) {
        try := 0
        for {
                var s string
                fmt.Print("Enter two letter card (rank+suit, ranks 2-9,T,J,Q,K,A, suits S,H,D,C): ")
                fmt.Scanln(&s)
                card, err := StringToCard(s)
                if err != nil {
                        fmt.Println(err)
                        try += 1
                        if try > 3 {
                                break
                        }
                }
                return card, err
        }
        return Card{0, 0}, fmt.Errorf("Three strikes, you're out")
}

func InputHand() (*Hand, error) {
        h := &Hand{}
        var s1, s2 string
        var err error
        fmt.Print("Enter hand: ")
        fmt.Scanln(&s1, &s2)
        h.Cards[0], err = StringToCard(s1)
        if err != nil {
                return h, err
        }
        h.Cards[1], err = StringToCard(s2)
        if err != nil {
                return h, err
        }
        return h, nil
}

func StringToCard(s string) (Card, error) {
        card := Card{ -1, -1}
        if len(s) != 2 {
                return card, fmt.Errorf("Invalid input - string length is off")
        }

        s = strings.ToUpper(s)

        found := false
        for i, v := range CardRankShortName {
                if rune(s[0]) == v {
                        card.Rank = i
                        found = true
                }
        }
        if !found {
                return card, fmt.Errorf("Invalid rank %c. Valid ranks are 2-9,T,J,Q,K,A\n", s[0])
        }

        switch s[1] {
        case 'S':
                card.Suit = SPADES
        case 'H':
                card.Suit = HEARTS
        case 'D':
                card.Suit = DIAMONDS
        case 'C':
                card.Suit = CLUBS
        default:
                return card, fmt.Errorf("Invalid suit %c. Valid suits are S,H,D,C\n", s[1])
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

func (hv *HandValue) GetName() string {
        ValueName := HandValueName[hv.Value]
        switch hv.Value {
        case HIGH_CARD:
                return fmt.Sprintf("%s, %s", ValueName, CardRankLongName[hv.HighCard[0]])
        case PAIR:
                return fmt.Sprintf("%s of %s", ValueName, CardRankLongNamePlural[hv.HighPair])
        case TWO_PAIRS:
                return fmt.Sprintf("%s, %s over %s", ValueName, CardRankLongNamePlural[hv.HighPair], CardRankLongNamePlural[hv.LowPair])   
        case TRIPS:
                return fmt.Sprintf("%s, %s", ValueName, CardRankLongNamePlural[hv.Trips]) 
        case STRAIGHT:
                return fmt.Sprintf("%s, %s high", ValueName, CardRankLongName[hv.Straight]) 
        case FLUSH, STRAIGHT_FLUSH, ROYAL_FLUSH:
                return fmt.Sprintf("%s", ValueName)
        case FULL_HOUSE:
                return fmt.Sprintf("%s, %s full of %s", ValueName, CardRankLongNamePlural[hv.Trips], CardRankLongNamePlural[hv.HighPair])
        case QUADS:
                return fmt.Sprintf("%s, %s", ValueName, CardRankLongNamePlural[hv.Quads]) 
        }
        return "<Unknown>"
}



