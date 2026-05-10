package main

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
	"time"
)

type JournalMode string

const (
	ModeFloat JournalMode = "float"
	ModeSwim  JournalMode = "swim"
)

var validMoods = []string{
	"grounded", "foggy", "tender", "light",
	"restless", "heavy", "still", "raw", "calm",
}

type Win struct {
	Title     string
	CreatedAt time.Time
}

func (w Win) String() string {
	return fmt.Sprintf("+ %s", w.Title)
}

type JournalEntry struct {
	Mood      string
	Body      string
	Mode      JournalMode
	Wins      []Win
	CreatedAt time.Time
}

func NewEntry(mood, body string, mode JournalMode) (*JournalEntry, error) {
	if !isMoodValid(mood) {
		return nil, fmt.Errorf("invalid mood %q: must be one of %s", mood, strings.Join(validMoods, ", "))
	}
	return &JournalEntry{
		Mood:      mood,
		Body:      body,
		Mode:      mode,
		CreatedAt: time.Now(),
	}, nil
}

func (e *JournalEntry) AddWin(title string) {
	e.Wins = append(e.Wins, Win{Title: title, CreatedAt: time.Now()})
}

func (e *JournalEntry) Summary() string {
	preview := e.Body
	if len(preview) > 60 {
		preview = preview[:60] + "..."
	}
	return fmt.Sprintf("[%s] %s (%d wins)", e.Mood, preview, len(e.Wins))
}

func isMoodValid(mood string) bool {
	for _, m := range validMoods {
		if m == mood {
			return true
		}
	}
	return false
}

// GenerateBlobSeed returns a deterministic uint32 seed for Mirror blob visuals.
func GenerateBlobSeed(mood, userID string) (uint32, error) {
	if mood == "" || userID == "" {
		return 0, errors.New("mood and userID must not be empty")
	}
	raw := fmt.Sprintf("%s:%s", mood, userID)
	sum := md5.Sum([]byte(raw))
	hex := hex.EncodeToString(sum[:4])
	var seed uint32
	fmt.Sscanf(hex, "%x", &seed)
	return seed, nil
}

func GetInsight(entry *JournalEntry) string {
	if entry.Body == "" {
		return "Nothing to reflect on yet."
	}
	words := strings.Fields(entry.Body)
	return fmt.Sprintf(
		"You wrote %d words in %s mode while feeling %s.",
		len(words), entry.Mode, entry.Mood,
	)
}

func main() {
	entry, err := NewEntry(
		"tender",
		"the morning was quiet enough to hear myself think",
		ModeSwim,
	)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	entry.AddWin("finished testing The Record")
	entry.AddWin("published Substack piece")

	fmt.Println(entry.Summary())
	fmt.Println(GetInsight(entry))

	seed, err := GenerateBlobSeed(entry.Mood, "user_karelle")
	if err != nil {
		fmt.Println("seed error:", err)
		return
	}
	fmt.Printf("blob seed: %d\n", seed)

	for _, win := range entry.Wins {
		fmt.Println(win)
	}
}