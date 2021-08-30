package main

const (
        DEUCE = iota
        THREE
        FOUR
        FIVE
        SIX
        SEVEN
        EIGHT
        NINE
        TEN
        JACK
        QUEEN
        KING
        ACE
)

const (
        SPADES = iota
        HEARTS
        DIAMONDS
        CLUBS
)

const (
        HIGH_CARD = iota
        PAIR
        TWO_PAIRS
        THREE_OF_KIND
        STRAIGHT
        FLUSH
        FULL_HOUSE
        QUADS
        STRAIGHT_FLUSH
        ROYAL_FLUSH
)

const ( //                              A KQJT 9876 5432
        STRAIGHT_ACE_HIGH       = 0B_0001_1111_0000_0000
        STRAIGHT_KING_HIGH      = 0B_0000_1111_1000_0000
        STRAIGHT_QUEEN_HIGH     = 0B_0000_0111_1100_0000
        STRAIGHT_JACK_HIGH      = 0B_0000_0011_1110_0000
        STRAIGHT_TEN_HIGH       = 0B_0000_0001_1111_0000
        STRAIGHT_NINE_HIGH      = 0B_0000_0000_1111_1000
        STRAIGHT_EIGHT_HIGH     = 0B_0000_0000_0111_1100
        STRAIGHT_SEVEN_HIGH     = 0B_0000_0000_0011_1110
        STRAIGHT_SIX_HIGH       = 0B_0000_0000_0001_1111
        STRAIGHT_FIVE_HIGH      = 0B_0001_0000_0001_1111
)

var CardRankShortName = []rune {
        '2',
        '3',
        '4',
        '5',
        '6',
        '7',
        '8',
        '9',
        'T',
        'J',
        'Q',
        'K',
        'A',
}

var CardRankLongName = []string {
        "Deuce",
        "Three",
        "Four",
        "Five",
        "Six",
        "Seven",
        "Eight",
        "Nine",
        "Ten",
        "Jack",
        "Queen",
        "King",
        "Ace",
}

var CardSuitShortName = []rune {
        /*
        '\u2660',
        '\u2665',
        '\u2666',
        '\u2663',
        */
        // For incomplete console fonts
        's',
        'h',
        'd',
        'c',
}

var CardSuitLongName = []string {
        "Spades",
        "Hearts",
        "Diamonds",
        "Clubs",
}

var HandValueName = []string {
        "High card",
        "Pair",
        "Two pairs",
        "Three of a kind",
        "Straight",
        "Flush",
        "Full house",
        "Four of a kind",
        "Straight flush",
        "Royal flush",
}

type Card struct {
        Rank int
        Suit int
}

type Hand struct {
        Cards  [2]Card
}

type Board struct {
        Cards  [5]Card
}

type Deck struct {
        Slots [52]bool
}

type HandValue struct {
        Straight        int
        Suits   [4]int
        Ranks   [13]int
}



