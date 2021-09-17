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
                fmt.Println(err)
                return h, err
        }
        h.Cards[1], err = StringToCard(s2)
        if err != nil {
                fmt.Println(err)
                return h, err
        }
        return h, nil
}

func InputRange() (Range, error) {
        var s string

        fmt.Print("Enter range: ")
        fmt.Scanln(&s)
        r, err := StringToRange(s)
        if err != nil {
                fmt.Println(err)
                return nil, err
        }
        return r, nil
}

func StringToRange(s string) (Range, error) {
        var r Range
        var rank1, rank2 int

        if len(s) < 2 {
                return nil, fmt.Errorf("Invalid input - string length is off")
        }

        s = strings.ToUpper(s)

        found := false
        for i, v := range CardRankShortName {
                if rune(s[0]) == v {
                        rank1 = i
                        found = true
                }
        }
        if !found {
                return nil, fmt.Errorf("Invalid rank %c. Valid ranks are 2-9,T,J,Q,K,A\n", s[0])
        }

        found = false
        for i, v := range CardRankShortName {
                if rune(s[1]) == v {
                        rank2 = i
                        found = true
                }
        }
        if !found {
                return nil, fmt.Errorf("Invalid rank %c. Valid ranks are 2-9,T,J,Q,K,A\n", s[1])
        }

        if rank1 > rank2 {
                rank1, rank2 = rank2, rank1
        }

        suited := BOTH
        plus   := false

        if len(s) >= 3 {
                switch s[2] {
                case 'S':
                        suited = SUITED
                case 'O':
                        suited = OFFSUIT
                case '+':
                        plus = true
                default:
                        return nil, fmt.Errorf("Invalid qualifier %c. Valid qualifiers are s,o,+\n", s[2])
                }
                
                if len(s) == 4 {
                        if s[3] == '+' {
                                plus = true
                        } else {
                                return nil, fmt.Errorf("Invalid range extension %c. Valid extension is +\n", s[2])
                        }
                }
        }

        if rank1 == rank2 { //return pair hands
                left, right := rank1, rank2
                if plus {
                        left = 0        //TT+ means AA,KK,QQ,JJ,TT
                }
                
                for rank := left; rank <= right; rank += 1 {
                        for suit1 := 0; suit1 < MAX_SUIT; suit1 += 1 {
                                for suit2 := suit1 + 1; suit2 < MAX_SUIT; suit2 += 1 {
                                        h := &Hand{}
                                        h.Cards[0] = Card{ rank, suit1}
                                        h.Cards[1] = Card{ rank, suit2}
                                        r = append(r, h)
                                }
                        }
                }
        } else if suited == SUITED {
                left, right := rank2, rank2
                if plus {
                        left = rank1 + 1        //AJs+ means AKs, AQs, AJs
                }
                for rank := left; rank <= right; rank += 1 {
                        for suit := 0; suit < MAX_SUIT; suit += 1 {
                                h := &Hand{}
                                h.Cards[0] = Card{ rank1, suit}
                                h.Cards[1] = Card{ rank, suit}
                                r = append(r, h)
                        }
                }
        } else if suited == OFFSUIT {
                left, right := rank2, rank2
                if plus {
                        left = rank1 + 1        //AJo+ means AKo, AQo, AJo
                }
                for rank := left; rank <= right; rank += 1 {
                        for suit1 := 0; suit1 < MAX_SUIT; suit1 += 1 {
                                for suit2 := 0; suit2 < MAX_SUIT; suit2 += 1 {
                                        if suit1 == suit2 {
                                                continue
                                        }
                                        h := &Hand{}
                                        h.Cards[0] = Card{ rank1, suit1}
                                        h.Cards[1] = Card{ rank, suit2}
                                        r = append(r, h)
                                }
                        }
                }
        } else {
                left, right := rank2, rank2
                if plus {
                        left = rank1 + 1        //AK+ means AKs, AKo, AQs, AQo, AJs, AJo
                }
                for rank := left; rank <= right; rank += 1 {
                        for suit1 := 0; suit1 < MAX_SUIT; suit1 += 1 {
                                for suit2 := 0; suit2 < MAX_SUIT; suit2 += 1 {
                                        h := &Hand{}
                                        h.Cards[0] = Card{ rank1, suit1}
                                        h.Cards[1] = Card{ rank, suit2}
                                        r = append(r, h)
                                }
                        }
                }
        }
        return r, nil
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
        return fmt.Sprintf("<%s%s>", h.Cards[0].GetShortName(), h.Cards[1].GetShortName())
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

func (r Range) GetShortName() string {
        var sb strings.Builder
        for i, v := range r {
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



