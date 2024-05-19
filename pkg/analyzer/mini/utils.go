package mini

import (
	"encoding/json"
	"os"
)

func strToMetric(str string) Metric {
	return mapStrMetric[str]
}

func strToFinger(str string) uint16 {
	return mapStrFinger[str]
}

func getFingerHash(keyMap map[rune]uint16, gram0, gram1, gram2 rune) (uint16, bool) {
	finger0, found := keyMap[gram0]
	if !found {
		return 0, false
	}
	finger1, found := keyMap[gram1]
	if !found {
		return 0, false
	}
	finger2, found := keyMap[gram2]
	if !found {
		return 0, false
	}
	return finger0<<8 | finger1<<4 | finger2, true
}

func LoadTable() [4096]Metric {
	jsonData, _ := os.ReadFile("table.json")

	var rawTable map[string]string
	_ = json.Unmarshal(jsonData, &rawTable)

	table := [4096]Metric{Unknown}
	for fingerStr, metricStr := range rawTable {
		finger0 := strToFinger(fingerStr[0:2])
		finger1 := strToFinger(fingerStr[2:4])
		finger2 := strToFinger(fingerStr[4:6])
		index := finger0<<8 | finger1<<4 | finger2
		table[index] = strToMetric(metricStr)
	}
	return table
}

func LoadCorpus(path string) []Ngram {
	jsonData, _ := os.ReadFile(path)

	var rawCorpus map[string]float64
	_ = json.Unmarshal(jsonData, &rawCorpus)

	var corpus []Ngram
	for chars, freq := range rawCorpus {
		corpus = append(corpus, Ngram{[]rune(chars), freq})
	}
	return corpus
}
