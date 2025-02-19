package main

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/muesli/termenv"
)

var (
	apiEndpoint    = "https://api.coingecko.com/api/v3/coins/markets"
	defaultHeight  = 15
	minHeight      = 5
	fiatCurrencies = []string{"usd", "eur"}
)

type tickMsg int
type model struct {
	height    int
	cursor    int
	width     int
	fiatIndex int
	coins     Coins
}

var term = termenv.ColorProfile()

func main() {
	m := model{
		height:    defaultHeight,
		cursor:    0,
		width:     0,
		fiatIndex: 0,
		coins:     nil,
	}

	if err := tea.NewProgram(&m).Start(); err != nil {
		fmt.Println("Oh no, it didn't work:", err)
		os.Exit(1)
	}
}

func (m *model) Init() tea.Cmd {
	return tea.Tick(time.Duration(0), func(t time.Time) tea.Msg {
		return tickMsg(1)
	})
}

func tick() tea.Cmd {
	return tea.Tick(time.Duration(time.Second*30), func(t time.Time) tea.Msg {
		return tickMsg(1)
	})
}
