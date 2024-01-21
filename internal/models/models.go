package models

// store "things" in the database

type Point struct {
	Col int // also the sorting order of cards
	Row int
}

type Card struct {
	Index int   // card index: 0 ... 51
	Pos   Point // card position in cardprite.png
	Suit  int   // 0=clubs, 1=hearts, 2=spades, 3=diamonds
	Hcp   int   // 0, 1, 2, 3, 4
}

type Suit struct {
	Cards  []Card // cards for this suit
	Name   string // name of this suit: schoppen, harten, etc
	Length int    // number of cards in this suit
	Hcp    int    // number of hcp points for this suit
	Reval  int    // re-valuated number of hcp when trump
}

type Hand struct {
	Pack  []Suit   // cards grouped by suit: 0=spades, 1=hearts, 2=clubs
	Cards []Card   // all cards for this hand => WE DON'T NEED THIS
	Hcp   int      // High Card Points for this hand	s = append(s, clubs)
	Dist  string   // "5332", "6510" dist as string
}

type Game struct {
	North  Hand // is a player and has a hand
	East   Hand
	South  Hand
	West   Hand
	KeyStr string // biddinglist for this game: "1H3H" etc.
}

type BidData struct {
	TopMsg string // text for msg-box-top
	Msg_1  string // text for msg-box-body
	Msg_2  string // text for msg-box-body
	Msg_3  string
	North  string // bid display list for each player
	East   string
	South  string
	West   string
}

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
}
