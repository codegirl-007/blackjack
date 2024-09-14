package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"log"
	"math/rand"
	"time"
)

type Suit int

func (e Suit) String() string {
	switch e {
	case club:
		return "\u2663"
	case spade:
		return "\u2660"
	case heart:
		return "\u2665"
	case diamond:
		return "\u2666"
	default:
		return fmt.Sprintf("%d", int(e))
	}
}

const (
	club Suit = iota
	spade
	heart
	diamond
)

type model struct {
	deck  Deck
	hands []hand
}

func (m model) Init() tea.Cmd {
	return nil // I don't know either
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {
	// The header
	s := "Lets play some blackjack.\n\n"

	return s
}

type Deck struct {
	cards []Card
}

func (d *Deck) shuffle() Deck {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(d.cards), func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})

	return *d
}

func (d *Deck) deal() *Card {
	if len(d.cards) == 0 {
		return nil
	}

	card := d.cards[len(d.cards)-1]
	d.cards = d.cards[0 : len(d.cards)-1]

	fmt.Printf("something something")
	return &card
}

type hand struct {
	cards []Card
}

func (h *hand) take(c *Card) *hand {
	if h.cards == nil {
		h.cards = []Card{}
	}

	h.cards = append(h.cards, *c)
	return h
}

type Card struct {
	Suit Suit
	Val  int
}

func initialModel() model {
	cards := make([]Card, 0, 52)
	hands := make([]hand, 0, 5) // dummy players for now
	for i := 0; i < 4; i++ {
		for j := 1; j < 15; j++ {
			cards = append(cards, Card{Suit: Suit(i), Val: j})
		}
	}

	for i := 0; i < 5; i++ {
		hands = append(hands, hand{cards: cards})
	}

	fmt.Println("yo")
	deck := Deck{cards: cards}
	deck = deck.shuffle()
	fmt.Println("shuffled cards")

	fmt.Println("length of hands", len(hands))
	// deal cards to each hand in game
	for i := 0; i < len(hands); i++ {
		fmt.Printf("dealing cards")
		hands[i].take(deck.deal())
	}

	return model{
		deck:  deck,
		hands: hands,
	}
}

func main() {
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		log.Fatalf("error logging to `debug.log`: %v", err)
	}
	defer f.Close()
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		log.Fatalf("error running program: %v", err)
	}
}
