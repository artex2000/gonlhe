package main

func EvalHand(h *Hand, b *Board, s int) *HandValue {
        c := &CardCombination{}
        r := h.Cards[0]
        c.NumOfSuit[r.Suit] += 1
        c.NumOfRank[r.Rank] += 1
        c.RankOfSuit[r.Suit] |= (1 << r.Rank)

        r = h.Cards[1]
        c.NumOfSuit[r.Suit] += 1
        c.NumOfRank[r.Rank] += 1
        c.RankOfSuit[r.Suit] |= (1 << r.Rank)

        for i := 0; i < s; i += 1 {
                r = b.Cards[i]
                c.NumOfSuit[r.Suit] += 1
                c.NumOfRank[r.Rank] += 1
                c.RankOfSuit[r.Suit] |= (1 << r.Rank)
        }

        for _, v := range c.RankOfSuit {
                c.AllRanks |= v
        }

        return EvalCombination(c)
}

func EvalCombination(c *CardCombination) *HandValue {
        hv := &HandValue{} 
        s, t := CheckStraightFlush(c)
        if s != -1 {    //we have straight flush
                if t == 0 {//Ace-High straight type
                        hv.Value = ROYAL_FLUSH
                } else {
                        hv.Value = STRAIGHT_FLUSH
                }
                hv.Flush = s    //suit
                hv.Straight = t

                //nothing to check further
                return hv
        }

        hv.Quads = CheckQuads(c)
        if hv.Quads != -1 {     //we have quads
                hv.Value = QUADS
                //let's figure kicker
                for i, v := range c.NumOfRank {
                        if v > 0 {
                                hv.HighCard[0] = i
                                break
                        }
                }

                //nothing to check further
                return hv
        }

        hv.Trips = CheckTrips(c)
        if hv.Trips != -1 {     //we have trips, check for full house
                hv.HighPair = CheckPair(c)
                if hv.HighPair != -1 {
                        hv.Value = FULL_HOUSE
                        return hv
                }
        }

        hv.Flush = CheckFlush(c)
        if hv.Flush != -1 {
                hv.Value = FLUSH
                //let's get exact cards for flush
                r := c.RankOfSuit[hv.Flush]
                idx := 0
                for i := 0; i < 13; i += 1 {
                        if (r & (1 << i)) != 0 {
                                hv.HighCard[idx] = i
                                idx += 1
                                if idx > 4 {
                                        return hv
                                }
                        }
                }
                panic("Not enough cards for flush in RankOfSuit")
        }

        hv.Straight = CheckStraight(c)
        if hv.Straight != -1 {
                hv.Value = STRAIGHT
                return hv
        }

        if hv.Trips != -1 {     //we have trips but not a full house
                //Get two high single cards
                idx := 0
                for i, v := range c.NumOfRank {
                        if v > 0 {
                                hv.HighCard[idx] = i
                                idx += 1
                                if idx > 1 {
                                        return hv
                                }
                        }
                }
                panic("Not enough high cards for trips in NumOfRank")
        }

        hv.HighPair = CheckPair(c)
        if hv.HighPair != -1 { //we have a pair
                hv.LowPair = CheckPair(c)
                if hv.LowPair != -1 { //we have two pairs
                        hv.Value = TWO_PAIRS
                        for i, v := range c.NumOfRank {
                                if v > 0 {
                                        hv.HighCard[0] = i
                                        return hv
                                }
                        }
                        panic("Not enough high cards for two pairs in NumOfRank")
                } else { //we just have one pair
                        hv.Value = PAIR
                        idx := 0
                        for i, v := range c.NumOfRank {
                                if v > 0 {
                                        hv.HighCard[idx] = i
                                        idx += 1
                                        if idx > 2 {
                                                return hv
                                        }
                                }
                        }
                        panic("Not enough high cards for pair in NumOfRank")
                }
        }

        hv.Value = HIGH_CARD
        idx := 0
        for i, v := range c.NumOfRank {
                if v > 0 {
                        hv.HighCard[idx] = i
                        idx += 1
                        if idx > 4 {
                                break
                        }
                }
        }
        if idx < 5 {
                panic("Not enough high cards for pair in NumOfRank")
        }
        return hv
}

func CheckStraightFlush(c *CardCombination) (int, int) {
        for i, v := range c.RankOfSuit {
                for j, t := range StraightTypes {
                        if (v & t) == t {
                                return i, j
                        }
                }
        }
        return -1, -1
}

func CheckFlush(c *CardCombination) int {
        for i, v := range c.NumOfSuit {
                if v >= 5 {
                        return i
                }
        }
        return -1
}

func CheckStraight(c *CardCombination) int {
        for i, v := range StraightTypes {
                if (v & c.AllRanks) == v {
                        return i
                }
        }
        return -1
}

func CheckQuads(c *CardCombination) int {
        for i, v := range c.NumOfRank {
                if v == 4 {
                        c.NumOfRank[i] = 0          //reset rank as accounted for
                        return i
                }
        }
        return -1
}

func CheckTrips(c *CardCombination) int {
        for i, v := range c.NumOfRank {
                if v == 3 {
                        c.NumOfRank[i] = 0          //reset rank as accounted for
                        return i
                }
        }
        return -1
}

func CheckPair(c *CardCombination) int {
        for i, v := range c.NumOfRank {
                if v >= 2 {     //we should account for low trips
                        c.NumOfRank[i] = 0          //reset rank as accounted for
                        return i
                }
        }
        return -1
}
