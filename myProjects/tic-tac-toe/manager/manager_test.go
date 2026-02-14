package manager

import "testing"

func TestCreateGame(t *testing.T) {
	m := New()
	err := m.CreateGame("game?!")

	if err != nil {
		t.Errorf("Expected no errror upon creation, got %s\n", err)
	}

	err = m.CreateGame("game?!")
	if err == nil {
		t.Errorf("Expected error for the same game, found nil")
	}

}

func TestGetGame(t *testing.T) {
	m := New()
	m.CreateGame("abc")

	g, err := m.GetGame("NoGameNoLife")
	if err == nil {
		t.Error("Expected error for nonexistent room")
	}
	if g != nil {
		t.Errorf("Expected nil game, found %s\n", g)
	}

	g, err = m.GetGame("abc")
	if err != nil {
		t.Errorf("Expected no error, found %s\n", err)
	}
	if g == nil {
		t.Error("Expected game, got nil")
	}
}

func TestDeleteGame(t *testing.T) {
	m := New()
	err := m.CreateGame("abc")

	if err != nil {
		t.Errorf("Error upon creation %s\n", err)
	}

	err = m.DeleteGame("abc")

	if err != nil {
		t.Errorf("Expected nil, got %s\n", err)
	}

	g, err := m.GetGame("abc")
	if g != nil {
		t.Error("Expected nil, instead found deleted game")
	}
	if err == nil {
		t.Error("Expected error, instead found nil")
	}

}
