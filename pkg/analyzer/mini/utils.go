package mini

import (
	"encoding/json"
	"os"
)

func StrToMetric(str string) Metric {
	return mapStrMetric[str]
}

func StrToFinger(str string) uint16 {
	return mapStrFinger[str]
}

func MetricToStr(metric Metric) string {
	return mapMetricStr[metric]
}

func FingerToStr(finger uint16) string {
	return mapFingerStr[finger]
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
	// TODO: Add error handling back in
	_ = json.Unmarshal(jsonData, &rawTable)

	table := [4096]Metric{Unknown}
	for fingerStr, metricStr := range rawTable {
		finger0 := StrToFinger(fingerStr[0:2])
		finger1 := StrToFinger(fingerStr[2:4])
		finger2 := StrToFinger(fingerStr[4:6])
		index := finger0<<8 | finger1<<4 | finger2
		table[index] = StrToMetric(metricStr)
	}
	return table
}

func LoadCorpus(path string) []Ngram {
	// TODO: Add error handling back in
	jsonData, _ := os.ReadFile(path)

	var rawCorpus map[string]float64
	// TODO: Add error handling back in
	_ = json.Unmarshal(jsonData, &rawCorpus)

	var corpus []Ngram
	for chars, freq := range rawCorpus {
		corpus = append(corpus, Ngram{[]rune(chars), freq})
	}
	return corpus
}
