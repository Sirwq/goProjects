package main

import (
	"testing"
)

func TestFreqCounter(t *testing.T) {
	result := freqCounter(`The quick brown Fox jumps over 
	the lazy dog. Fox? Yes, the fox is quick-very quick!
	123 123 123. The dog, however, is not... not at all.`)

	if result["the"] != 4 {
		t.Errorf("expected 'The' = 2, got %d", result["The"])
	}

	if result["quick"] != 3 {
		t.Errorf("expected 'quick' = 3, got %d", result["quick"])
	}

	if result["brown"] != 1 {
		t.Errorf("expected 'brown' = 1, got %d", result["brown"])
	}

	if result["dog"] != 2 {
		t.Errorf("expected 'dog' = 2, got %d", result["dog"])
	}

	//fmt.Println(result)

}
