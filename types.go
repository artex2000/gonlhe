package main

const (
        ACE     = iota
        KING
        QUEEN
        JACK
        TEN
        NINE
        EIGHT
        SEVEN
        SIX
        FIVE
        FOUR
        THREE
        DEUCE

        MAX_RANK
)

const (
        SPADES = iota
        HEARTS
        DIAMONDS
        CLUBS

        MAX_SUIT
)

const (
        HIGH_CARD = iota
        PAIR
        TWO_PAIRS
        TRIPS
        STRAIGHT
        FLUSH
        FULL_HOUSE
        QUADS
        STRAIGHT_FLUSH
        ROYAL_FLUSH
)

const (
        FLOP  = 3
        TURN  = 4
        RIVER = 5
)

const (
        WIN  = iota
        TIE
        LOSE
)

const (
        SUITED  = iota
        OFFSUIT
        BOTH
)

var CardRankShortName = []rune {
        'A',
        'K',
        'Q',
        'J',
        'T',
        '9',
        '8',
        '7',
        '6',
        '5',
        '4',
        '3',
        '2',
}

var CardRankLongName = []string {
        "Ace",
        "King",
        "Queen",
        "Jack",
        "Ten",
        "Nine",
        "Eight",
        "Seven",
        "Six",
        "Five",
        "Four",
        "Three",
        "Deuce",
}

var CardRankLongNamePlural = []string {
        "Aces",
        "Kings",
        "Queens",
        "Jacks",
        "Tens",
        "Nines",
        "Eights",
        "Sevens",
        "Sixes",
        "Fives",
        "Fours",
        "Threes",
        "Deuces",
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

var StraightTypes = []int {
        //    2 3456 789T JQKA
        0B_0000_0000_0001_1111,
        0B_0000_0000_0011_1110,
        0B_0000_0000_0111_1100,
        0B_0000_0000_1111_1000,
        0B_0000_0001_1111_0000,
        0B_0000_0011_1110_0000,
        0B_0000_0111_1100_0000,
        0B_0000_1111_1000_0000,
        0B_0001_1111_0000_0000,
        0B_0001_1110_0000_0001,
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

type CardCombination struct {
        RankOfSuit     [4]int
        NumOfSuit      [4]int
        NumOfRank      [13]int
        AllRanks       int
}

type HandValue struct {
        Value           int
        Flush           int
        Straight        int
        Quads           int
        Trips           int
        HighPair        int
        LowPair         int
        HighCard        [5]int
}

type Range []*Hand

type Outcome struct {
        Wins int
        Ties int
        Loss int
}


