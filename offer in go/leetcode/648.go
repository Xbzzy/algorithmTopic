package leetcode

import "strings"

func replaceWords(dictionary []string, sentence string) string {
	if sentence == "" {
		return ""
	}
	max, min := 0, 101
	wordMap := make(map[string]bool, len(dictionary))
	for _, val := range dictionary {
		wordMap[val] = true
		if len(val) > max {
			max = len(val)
		}
		if len(val) < min {
			min = len(val)
		}
	}
	sentenceWord := strings.Split(sentence, " ")
	for i, val := range sentenceWord {
		for j := min; j <= max; j++ {
			if len(val) >= j {
				if wordMap[val[:j]] == true {
					sentenceWord[i] = val[:j]
				}
			}
		}
	}
	return strings.Join(sentenceWord, " ")
}
