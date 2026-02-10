package manager

import (
	"errors"
	"sync"
	"tic-tac-toe/game"
)

type Manager struct {
	mu    sync.Mutex
	games map[string]*game.Game
}

func (m *Manager) CreateGame(name string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, exists := m.games[name]; exists {
		return errors.New("room already exists")
	}
	m.games[name] = game.New()
	return nil
}

func (m *Manager) GetGame(name string) (*game.Game, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, exists := m.games[name]; exists {
		return m.games[name], nil
	}
	return nil, errors.New("couldn't create game")
}

func New() *Manager {
	return &Manager{
		games: New().games,
	}
}
