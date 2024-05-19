package mini

import (
	"strings"

	"github.com/akl-infra/slf/v2"
	"github.com/charmbracelet/log"
)

type Corpus []Ngram

var Corpora map[string]Corpus

func Init() {
	// TODO: Read all corpora
	// Load corpora into local map/cache
	Corpora = make(map[string]Corpus)
	Corpora["monkeyracer"] = LoadCorpus("corpora/monkeyracer/trigrams.json")
	// Load table into local map
	Table = LoadTable()
}

// AnalyzeTrigrams
// - using cmini implementation
// - ignore spaces
// - case-insensitive
// - sfb is double counted

// TODO: Change corpus to `name string` to load from map
// func Analyze(layout *slf.Layout, corpus *[]Ngram) []float64 {
func Analyze(layout *slf.Layout, corpus string) []float64 {
	keyMap := make(map[rune]uint16)
	counter := make([]float64, MetricNum)

	for _, key := range layout.Keys {
		if key.Char == "" {
			continue
		}
		s := strings.ToLower(key.Char)
		char := []rune(s)[0]
		finger := uint16(key.Finger)
		keyMap[char] = finger
	}
	for _, trigram := range Corpora[corpus] {
		gram0 := trigram.chars[0]
		gram1 := trigram.chars[1]
		gram2 := trigram.chars[2]

		if gram0 == ' ' || gram1 == ' ' || gram2 == ' ' {
			continue
		}
		if gram0 == gram1 || gram1 == gram2 {
			counter[Sfr] += trigram.freq
			continue
		}

		fingerHash, ok := getFingerHash(keyMap, gram0, gram1, gram2)
		if !ok {
			counter[Unknown] += trigram.freq
			continue
		}
		gramType := Table[fingerHash]
		if gramType == Sfr {
			log.Info(trigram)
		}
		counter[gramType] += trigram.freq
	}

	var total float64
	for _, freq := range counter {
		total += freq
	}
	for index, freq := range counter {
		counter[index] = freq / total
	}
	counter[Sfb] /= 2
	return counter
}
