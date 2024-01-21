package bidgame

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math/rand"
	"os"
	"sort"

	"github.com/aleph-design/startbridge/internal/config"
	"github.com/aleph-design/startbridge/internal/models"
)

// provide site wide access to app config
var app *config.AppConfig

// actually populate this package with site-wide
// access to AppConfig by transferring a pointer
// *config.AppConfig and call this function from
// main()
func NewBidGame(a *config.AppConfig) {
	app = a
}

const (
	cardWidth  = 209
	cardHeight = 303
	cardSpace  = 70
	extraFill  = 30
	cardFill   = 109
	outFile    = "static/images/cardFile.png"
	spriteFile = "static/images/cardsprite.png"
)

// @ g
//   - pointer to models.Game
//     the new Game struct has been created in handler Bidgame
func InitNewGame(g *models.Game) error {

	// create a slice of int and populate it with card numbers ranging from 0 to 51
	cDeck := getCardNumbers()
	// fmt.Println("\ncDeck: ", cDeck)

	// generate hands for each player in this game,
	err := generateHands(cDeck, g)
	if err != nil {
		fmt.Println("52 - bidgame-generateHands: ", err)
	}

	// create card images and /static/images/cardFile.png
	drawSouthHandImages(g)

	return nil
}

// get card ID's or numbers, it's just a slice of int 0 ... 51
// actually this is a one time operation; not necessay for every new bidround
func getCardNumbers() []int {
	return []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
}

// generate card hands for each player
// ===================================
func generateHands(deck []int, g *models.Game) error {
	// shuffle the slice of numbers; there are no cards in it yet
	rand.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})
	// fmt.Println("\nShuffled-Deck: ", deck)

	// create []cards with data for each card and store it
	// variable cards contains all shuffled card data
	cards := newCards(deck)
	// PrintCards(*cards)

	// distribute an unsorted ungrouped bunch of cards over each
	// players hand: N, E, S, W
	handsDistribution(cards, g)
	// PrintHands(g)

	distributeSuitsByHand(g)
	// PrintSouthSuits(g)

	setHCPtotalForHands(g)
	// ShowSouthHand(g)


	return nil
} // generateHandsImages() =====================================

// add card attributes to an allready shuffled deck of card numbers
// @ deck
// -	slice of numbered elements from 0 ... 51
// @ return
// -	slice of shuffled cards with card data, not distributed over hands
func newCards(deck []int) *[]models.Card {
	c := models.Card{}   // one individual card
	s := []models.Card{} // create a slice of cards

	for _, value := range deck {
		c.Index = value           // card index [0 <= value < 52]
		c.Pos.Col = value % 13    // card column in cardsprite image
		c.Pos.Row = value / 13    // card row in cardsprite image
		c.Suit = c.Pos.Row        // suit: 0=spades, 1=hearts, 2=clubs, 3=diamonds
		c.Hcp = addHCP(c.Pos.Col) // 0 <= rank <= 5 depending on column
		// iets := fmt.Sprintf("value: %d - row: %d - col: %d\n", value, c.pos.row, c.pos.col)
		// fmt.Print(iets)
		s = append(s, c)
	}
	return &s
} // newCards =============

// add High Card Point value to each individual card
// @ col
//   - indicates position of the card in sprite row
//     0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12
//
// @ return
//   - High Card Points of specified position
func addHCP(col int) int {
	switch col {
	case 9: // de Boer
		return 1
	case 10: // de Vrouw
		return 2
	case 11: // de Koning
		return 3
	case 12: // het Aas
		return 4
	default: // andere kaarten
		return 0
	}
} // addHCP ==========

// distribute the shuffled slice of cards over hands
// for each player: N, E, S, W
// @ cards
//   - slice of shuffled cards to be distributed over hands
//
// The result is that each hand contains a bunch of 13 unsorted
// and ungrouped cards.
func handsDistribution(cards *[]models.Card, g *models.Game) {
	// fmt.Println("\nlen(c): ", len(c))
	// fmt.Println("\nc[0]: ", c[0])
	// every fourth card goes to a different hand, simulates
	// the process of dealing cards
	for index, card := range *cards {
		suitPtr := index % 4
		// iets := fmt.Sprintf("index: %d - suit: %v\n", index, suit)
		// fmt.Print(iets)
		// the indexed card goes to the hand of a player
		switch suitPtr {
		case 0: // north
			g.North.Cards = append(g.North.Cards, card)
		case 1: // east
			g.East.Cards = append(g.East.Cards, card)
		case 2: // south
			g.South.Cards = append(g.South.Cards, card)
		case 3: // west
			g.West.Cards = append(g.West.Cards, card)
		}
	} // for loop
	// fmt.Println("\ng.South.Hand: ", g.South.Cards)
	// fmt.Println("\ng.South.Pack: ", g.South.Pack[2])
} // distributedHands() =============================

/*
====================================================
	BEGIN INSERTED NEW FUNCTIONS
====================================================
*/

// For example, North.Cards is a buch of 13 unsorted
// ungrouped cards.
// So we range over the hands and within the hand, we
// range over those 13 cards to group them into suits.
// Mext step is to sort the suit.

// For each hand distribute cards into suits
func distributeSuitsByHand(g *models.Game) {
	for hand := 0; hand < 4; hand++ {
		switch hand {
		case 0:
			g.North.Pack = groupCardsBySuit(&g.North.Cards)
		case 1:
			g.East.Pack = groupCardsBySuit(&g.East.Cards)
		case 2:
			g.South.Pack = groupCardsBySuit(&g.South.Cards)
		case 3:
			g.West.Pack = groupCardsBySuit(&g.West.Cards)
		}
	}
} // distributeHandsBySuit() ========================

// distributed the bunch of 13 cards over suits
// returns ALL suits for one hand only
func groupCardsBySuit(cards *[]models.Card) []models.Suit {
	// suits are: 0=spades, 1=hearts, 2=clubs, 3=diamonds
	// yet unsorted but grouped by their suit

	s := []models.Card{} // create new slice for suit of spades
	h := []models.Card{} // new suit for hearts
	c := []models.Card{} // idem clubs
	d := []models.Card{} // idem diamonds

	for _, card := range *cards {
		// switch on field suit in the card
		switch card.Suit {
		case 0:
			s = append(s, card) // suit of spades
		case 1:
			h = append(h, card) // suit of hearts
		case 2:
			c = append(c, card) // suit of clubs
		case 3:
			d = append(d, card) // suit of diamonds
		}
	} // for loop

	// sort cards high to low within each suit
	s = sortSlice(s) // spades
	h = sortSlice(h) // hearts
	c = sortSlice(c) // clubs
	d = sortSlice(d) // diamonds

	pack := []models.Suit{}
	for suit := 0; suit < 4; suit++ {
		switch suit {
		case 0:
			pack = append(pack, getSuitData(suit, s))
		case 1:
			pack = append(pack, getSuitData(suit, h))
		case 2:
			pack = append(pack, getSuitData(suit, c))
		case 3:
			pack = append(pack, getSuitData(suit, d))
		}
	}

	// // North pack must contain four suits
	// pack = []models.Suit{}
	// g.North.Pack = append(g.North.Pack, suit)

	return pack
}

var name = map[int]string{
	0: "schoppen",
	1: "harten",
	2: "klaveren",
	3: "ruiten",
}

// THIS IS JUST A SINGLE SUIT
func getSuitData(id int, s []models.Card) models.Suit {
	suit := models.Suit{}

	suit.Cards = []models.Card{}
	suit.Cards = s
	suit.Hcp = getSuitHcp(&s)
	suit.Length = len(s)
	suit.Name = name[id] // when suit is spades
	suit.Reval = 0       // a re-evaluation for this suit WHEN IT'S TRUMP

	return suit
}

func getSuitHcp(cards *[]models.Card) int {
	p := 0
	for _, card := range *cards {
		p += card.Hcp
	}
	return p
}

/*
====================================================
	END INSERTED NEW FUNCTIONS
====================================================
*/

// sort the slice of cards high => low
func sortSlice(s []models.Card) []models.Card {
	sort.Slice(s, func(i, j int) bool {
		return s[i].Pos.Col > s[j].Pos.Col
	})
	// p, _ := fmt.Printf("%v", s)
	// fmt.Println(p)
	return s
} // sortSlice() ==========

// set High Card Point total for all hands.
// still using the []cards
func setHCPtotalForHands(g *models.Game) {
	for hand := 0; hand < 4; hand++ {
		switch hand {
		case 0:
			g.North.Hcp = setHCPhandTotal(&g.North.Cards)
		case 1:
			g.East.Hcp = setHCPhandTotal(&g.East.Cards)
		case 2:
			g.South.Hcp = setHCPhandTotal(&g.South.Cards)
		case 3:
			g.West.Hcp = setHCPhandTotal(&g.West.Cards)
		}
	}
} // setHCPtotalForHands() ============

// set High Card Point total for this hand
// @ cards
// -	a hand of cards destined to one player
// @ return
//   - an int value for the total high card points
//     for this player
func setHCPhandTotal(cards *[]models.Card) int {
	// c := *cards
	p := 0
	for _, card := range *cards {
		p += card.Hcp
	}
	return p
} // setHCPtotal() =========


/***********************
*	END GENERATING HANDS *
************************/

func drawSouthHandImages(g *models.Game) {
	// create a "canvas" to receive card images
	canvas := createCardsImagesCanvas(g)

	// set background - kleur van het canvas
	bg := image.NewUniform(color.Transparent)

	// create the background - het canvas
	// draw.Draw(canvas, canvas.Bounds(), &image.Uniform{}, image.ZP, draw.Src)
	draw.Draw(canvas, canvas.Bounds(), bg, image.Point{0, 0}, draw.Src)

	// Read the whole sprite image from it's file
	spriteData := readSpriteData(spriteFile)

	// card rectangle
	cardRect := createCardRectangle()

	//here start the loop
	//===================
	xOff := 0 // X offset for each suit
	// range over the suit pack for the south hand
	for _, suit := range g.South.Pack {

		fmt.Println("\nsuitLength: ", suit.Length, "suitName: ", suit.Name)

		if 0 < suit.Length {
			xPos := 0 // X position for each card
			// range over each card in the suit
			for cdx, card := range suit.Cards {

				fmt.Println("cdx: ", cdx, "card: ", card)

				// cardPos is used to position cardRect on the canvas
				// X and Y depends on the cards position in the hand
				xPos = cdx*cardSpace + xOff
				cardPos := image.Point{xPos, 0}
				// calculate upperLeft sprite source: a Point{X, Y}
				cardLeft := cardWidth * card.Pos.Col
				cardTop := cardHeight * card.Pos.Row
				// select a card image from the sprite.
				// sprite source coordinates for the selected sprite part
				// you'll need the cards column and row to calculate these
				// values: col=3 and row=2 evaluate to 418, 303
				spriteSrc := image.Point{cardLeft, cardTop}
				// Draw(dst draw.Image,
				//      r image.Rectangle,	the rectangle in which we draw
				//      src image.Image, 		the sprite from which we take a part
				//      sp image.Point, 		upperleft point of the selected part
				//      op draw.Op)					the way we draw, scr is fast and
				//													overwrite all pixels in destiny image
				draw.Draw(canvas, cardRect.Bounds().Add(cardPos), spriteData, spriteSrc, draw.Src)
				// xPos =
				// iets := fmt.Sprintf("bdx %d - cdx %d - xOff: %d - xPos: %d", bdx, cdx, xOff, xPos)
				// fmt.Println(iets)
			}
			// this is done only when the suit is NOT empty
			xOff = xPos + cardSpace + extraFill
		} // end if
	} // end for-loop

	// write image to file
	writeImageToFile(canvas)
} // createCardsImageFile() ========================

func createCardsImagesCanvas(g *models.Game) *image.RGBA {
	canvasWidth := calculateCanvasWidth(g)
	// fmt.Println("width: ", canvasWidth)
	return image.NewRGBA(image.Rect(0, 0, canvasWidth, cardHeight))
}

func calculateCanvasWidth(g *models.Game) int {
	width := 0
	// we create images for the south hand.
	// for _, suit := range g.South.Suits {
	// 	// number of cards in the suit
	// 	width += len(suit)*cardSpace + extraFill
	// }
	return width + cardFill
}

func readSpriteData(f string) image.Image {
	spriteFile, err := os.Open(f)
	if err != nil {
		fmt.Println("Err open file: ", err)
		panic(err)
	}
	defer spriteFile.Close()

	spriteData, err := png.Decode(spriteFile)
	if err != nil {
		fmt.Println("Err decoding image file: ", err)
		panic(err)
	}
	return spriteData
}

func createCardRectangle() image.Rectangle {
	// constant for every card image
	min := image.Point{0, 0}
	max := image.Point{209, 303}
	// type Rectangle struct {
	//   Min, Max Point
	// }
	return image.Rectangle{min, max}
}

func writeImageToFile(image *image.RGBA) {
	cardFile, err := os.Create(outFile)
	if err != nil {
		fmt.Println("Err creating outFile: ", err)
		return
	}
	// fmt.Println("write image to file")
	png.Encode(cardFile, image)
	cardFile.Close()
}

// ==============================
// Available outside this package
// ==============================

func PrintCards(c *[]models.Card) {
	fmt.Println("Cards for each hand =================")
	b := *c
	for _, card := range b {
		iets := fmt.Sprintf("index: %d- color: %d- hcp: %d", card.Index, card.Suit, card.Hcp)
		fmt.Println(iets)
	}
} // PrintCards() ======

func PrintHands(g *models.Game) {
	fmt.Println("The hands ======================")
	for i := 0; i < 4; i++ {
		switch i {
		case 0:
			fmt.Println("\nNorth: ", g.North.Cards)
		case 1:
			fmt.Println("\nEast:  ", g.East.Cards)
		case 2:
			fmt.Println("\nSouth: ", g.South.Cards)
		case 3:
			fmt.Println("\nWest:  ", g.West.Cards)
		}
	}
} // PrintHands() ======

func PrintSouthSuits(g *models.Game) {
	fmt.Println("\nSouth suits ======================")
	for suit := 0; suit < 4; suit++ {
		fmt.Println("Name  : ", g.South.Pack[suit].Name)
		fmt.Println("Cards : ", g.South.Pack[suit].Cards)
		fmt.Println("Length: ", g.South.Pack[suit].Length)
		fmt.Println("Hcp   : ", g.South.Pack[suit].Hcp)
	}
} // PrintHands() ======

func ShowSouthHand(g *models.Game) {
	// s := g.South
	// iets := fmt.Sprintf("south-hand:\n- cards %v\n- suits: %v\n- hcp: %d\n- dist: %s\n", s.Cards, s.Suits, s.Hcp, s.Dist)
	// fmt.Println(iets)
} // PrintSouthSuits() ====
