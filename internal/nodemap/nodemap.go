package nodemap

// type Txt struct {
// 	T1 string
// 	T2 string
// 	T3 string
// 	T4 string
// }

// type Node struct {
// 	PtnMin  int    // minimum points needed for this bid
// 	PtnMax  int    // maximum points needed for this bid
// 	Cards   int    // Minimum number or cards in this suit needed
// 	Suit    int    // Suit: 0=Spades, 1=Hearts, 2=Clubs, 3=Diamonds
// 	Balance bool   // Evenwichtige hand
// 	Msg     *Txt   // pointer to text struct
// 	BidTwo  string // norths second bijbod
// }

// // Message text
// var t50 = &Txt{T1: "Heel verstandig,", T2: "Er zijn voldoende punten: BIEDEN!"}
// var t51 = &Txt{T1: "Er zijn 12-19 punten nodig om te openen", T2: "Teveel punten.", T3: "Er ontbreekt een 5+ kaart!"}
// var t52 = &Txt{T1: "Er zijn 12-19 punten nodig om te openen", T2: "Teveel punten.", T3: "Er ontbreekt een 2+ kaart!"}
// var t53 = &Txt{T1: "Te weinig punten", T2: "Teveel punten.", T3: "Te weinig harten."}
// var t54 = &Txt{T1: "Te weinig punten", T2: "Teveel punten.", T3: "Te weinig ruiten."}
// var t55 = &Txt{T1: "Te weinig punten", T2: "Teveel punten.", T3: "Te weinig schoppen."}
// var t56 = &Txt{T1: "VERPLICHT BIEDEN!"}
// var t57 = &Txt{T1: "Te weinig punten", T2: "Teveel punten.", T3: "Geen evenwichtige verdeling!"}
// var t58 = &Txt{T1: "Er zijn 12-19 punten nodig om te openen", T2: "Teveel punten.", T3: "Er ontbreekt een 4+ kaart!"}
// var t59 = &Txt{T1: "Er zijn 15-17 punten nodig om SA te openen", T2: "Teveel punten.", T3: "Géén evenwichtige verdeling!"}

// // OPENING op 1-niveau: 12-19 punten; voorlopig alleen: PA, 1H en 1S
// // =================================================================
// var n100 = &Node{PtnMin: 0, PtnMax: 12, Cards: 0, Msg: t50}       // PA
// var n101 = &Node{PtnMin: 12, PtnMax: 19, Cards: 5, Msg: t51}      // 1H, 1S
// var n102 = &Node{PtnMin: 12, PtnMax: 19, Cards: 2, Msg: t52}      // 1C
// var n103 = &Node{PtnMin: 12, PtnMax: 19, Cards: 4, Msg: t58}      // 1D
// var n104 = &Node{PtnMin: 15, PtnMax: 17, Balance: true, Msg: t59} // 1N

// // 1E-REBID: Er zijn altijd al minstens 12 punten bij de openaar
// // =============================================================
// // BidTwo
// var n150 = &Node{PtnMin: 12, PtnMax: 15, Cards: 0, Msg: t51, BidTwo: "PA"} // 1H2HPA
// var n151 = &Node{PtnMin: 16, PtnMax: 17, Cards: 0, Msg: t51} // 1H2H3H
// var n152 = &Node{PtnMin: 18, PtnMax: 19, Cards: 0, Msg: t51, BidTwo: "PA"} // 1H2H4HPA needs another PA
// var n153 = &Node{PtnMin: 12, PtnMax: 14, Cards: 0, Msg: t51, BidTwo: "PA"}
// var n154 = &Node{PtnMin: 15, PtnMax: 19, Cards: 0, Msg: t51, BidTwo: "PA"} // 1H3H4H
// var n155 = &Node{PtnMin: 12, PtnMax: 15, Cards: 6, Suit: 1, Msg: t53}
// var n156 = &Node{PtnMin: 16, PtnMax: 17, Cards: 6, Suit: 1, Msg: t53}
// var n157 = &Node{PtnMin: 18, PtnMax: 19, Cards: 6, Suit: 1, Msg: t53, BidTwo: "PA"}
// var n158 = &Node{PtnMin: 12, PtnMax: 17, Cards: 4, Suit: 3, Msg: t54}
// var n159 = &Node{PtnMin: 18, PtnMax: 19, Cards: 4, Suit: 3, Msg: t54}
// var n160 = &Node{PtnMin: 12, PtnMax: 17, Cards: 4, Suit: 0, Msg: t55}
// var n161 = &Node{PtnMin: 18, PtnMax: 19, Cards: 4, Suit: 0, Msg: t55}
// var n162 = &Node{PtnMin: 12, PtnMax: 13, Cards: 0, Msg: t55, BidTwo: "PA"}
// var n163 = &Node{PtnMin: 14, PtnMax: 40, Cards: 0, Msg: t55, BidTwo: "PA"}
// var n164 = &Node{PtnMin: 12, PtnMax: 13, Cards: 6, Suit: 1, Msg: t53, BidTwo: "PA"}
// var n165 = &Node{PtnMin: 14, PtnMax: 40, Cards: 6, Suit: 1, Msg: t53, BidTwo: "PA"}
// var n166 = &Node{PtnMin: 12, PtnMax: 13, Cards: 4, Suit: 3, Msg: t53}
// var n167 = &Node{PtnMin: 12, PtnMax: 40, Cards: 2, Suit: 1, Msg: t53, BidTwo: "PA"}
// var n168 = &Node{PtnMin: 14, PtnMax: 40, Cards: 6, Suit: 1, Msg: t53, BidTwo: "PA"}

// var n180 = &Node{PtnMin: 40, PtnMax: 40, Cards: 0, Msg: t56}
// var n181 = &Node{PtnMin: 12, PtnMax: 14, Cards: 6, Suit: 1, Balance: true, Msg: t57}

// // Biedverloop 1H of 1S opening
// var N = map[string]*Node{
// 	//													COMPUTER					COMPUTER
// 	// 								 Opening	1eBijbod	Herbod	2eBijbod
// 	"PA":     n100, // PASS
// 	"1H":     n101, // 1 Hart
// 	"1S":     n101, // *****************************
// 	"1C":     n102,
// 	"1D":     n103,
// 	"1N":     n104,
// 	"1H2HPA": n150,
// 	"1S2SPA": n150,
// 	"1H2H3H": n151,
// 	"1H2H4H": n152,
// 	"1H3HPA": n153,
// 	"1H3H4H": n154,
// 	"1H4HPA": n150,
// 	"1H1NPA": n153,
// 	"1H1N2N": n151,
// 	"1H1N3N": n152,
// 	"1H1N2H": n150,
// 	"1H1N3H": n156,
// 	"1H1N4H": n157,
// 	"1H1N2D": n158,
// 	"1H1N3D": n159,
// 	"1H1N2S": n160,
// 	"1H1N3S": n161,
// 	"1H2NPA": n162,
// 	"1H2N3N": n163,
// 	"1H2N3H": n164,
// 	"1H2N4H": n165,
// 	"1H2N3D": n166,
// 	"1H3NPA": n167,
// 	"1H3N4H": n168,

// 	"1H1SPA": n180,
// 	"1H1S1N": n181,
// }
