package bidmap

// NT doesn't need a suit, neither does "RD", "", "DB" and "PA"
type Bid struct {
	ID   int    // 0 ... 39 [1 clubs ... PASS]
	Suit int    // 0=Spades, 1=Hearts, 2=Clubs, 3=Diamonds are valid references to hand suit
	Str  string // multiple use field contains: "1H", "1S", "3H" o.i.d.
	Img  string // display for this bid
}

type BidIndex struct {
	Count int    // counter for each individual bid
	List  []*Bid // each individual bid
}

// The KEY is an integer!
var B = map[int]*Bid{

	0: {0, 2, "1C", `1♣️`},
	1: {1, 3, "1D", `1♦️`},
	2: {2, 1, "1H", `1♥️`},
	3: {3, 0, "1S", `1♠️`},
	4: {4, 0, "1N", "1N"}, // dummy value for suit
	5: {5, 2, "2C", `1♣️`},
	6: {6, 3, "2D", `1♦️`},
	7: {7, 1, "2H", `2♥️`},
	8: {8, 0, "2S", `2♠️`},
	9: {9, 0, "2N", "2N;"}, // Dummy suit value

	10: {10, 0, "3C", `3♣️`},
	11: {11, 1, "3D", `3♦️`},
	12: {12, 2, "3H", `3♥️`},
	13: {13, 3, "3S", `3♠️`},
	14: {14, 0, "3N", "3N;"}, // Dummy suit value
	15: {15, 0, "4C", `4♣️`},
	16: {16, 1, "4D", `4♦️`},
	17: {17, 2, "4H", `4♥️`},
	18: {18, 3, "4S", `4♠️`},
	19: {19, 0, "4N", "4N;"}, // Dummy suit value

	20: {20, 0, "5C", `5♣️`},
	21: {21, 1, "5D", `5♦️`},
	22: {22, 2, "5H", `5♥️`},
	23: {23, 3, "5S", `5♠️`},
	24: {24, 0, "5N", "5N;"}, // Dummy suit value
	25: {25, 0, "6C", `6♣️`},
	26: {26, 1, "6D", `6♦️`},
	27: {27, 2, "6H", `6♥️`},
	28: {28, 3, "6S", `6♠️`},
	29: {29, 0, "6N", "6N;"}, // Dummy suit value

	30: {30, 0, "7C", `7♣️`},
	31: {31, 1, "7D", `7♦️`},
	32: {32, 2, "7H", `7♥️`},
	33: {33, 3, "7S", `7♠️`},
	34: {34, 0, "7N", "7N;"}, // Dummy suit value
	35: {35, 0, "RD", "RD"},  // RDBL// Dummy suit value
	36: {36, 0, "", ""},      // Dummy suit value
	37: {37, 0, "DB", "DB"},  // DBL// Dummy suit value
	38: {38, 0, "", ""},      // Dummy suit value
	39: {39, 0, "PA", "PA"},  // PASS // Dummy suit value
}

func NewBid(id int) *Bid {
	bid := &Bid{}
	bid.ID = id
	bid.Suit = B[id].Suit
	bid.Str = B[id].Str
	bid.Img = B[id].Img
	return bid
}
