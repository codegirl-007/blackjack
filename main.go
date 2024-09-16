package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"log"
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
	s := ""

	for _, h := range m.hands {
		for _, c := range h.cards {
			s = s + c.Suit.String()
		}
	}

	return s
}

func initialModel() model {
	hands := make([]hand, 0, 2)

	for i := 0; i < 2; i++ {
		hands = append(hands, hand{})
	}

	deck := MakeDeck()
	deck = deck.shuffle()

	// deal cards to each hand in game
	for i := 0; i < len(hands); i++ {
		d := deck.deal()
		hands[i].take(d)
		hands[i].take(deck.deal()) // give each play two cards to start
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
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatalf("error running program: %v", err)
	}
}
