package main

import (
	"fmt"
	"github.com/jdkato/prose/v3"
	"math"
)

func main() {
	text1 := "예시입니다. 한글 벡터"
	text2 := "한글 벡터 예시입니다."

	vector1 := textToVector(text1)
	vector2 := textToVector(text2)

	similarity := euclideanSimilarity(vector1, vector2)

	fmt.Printf("Similarity: %f\n", similarity)
}

func textToVector(text string) map[string]float64 {
	tokenizer := prose.NewIterTokenizer()

	tokens := tokenizer.Tokenize(text)

	tokenMap := make(map[string]float64)
	for _, token := range tokens {
		tokenMap[token.Text]++
	}

	return tokenMap
}

func euclideanSimilarity(vector1, vector2 map[string]float64) float64 {
	terms := make(map[string]bool)
	for term := range vector1 {
		terms[term] = true
	}
	for term := range vector2 {
		terms[term] = true
	}

	v1 := make([]float64, len(terms))
	v2 := make([]float64, len(terms))
	i := 0
	for term := range terms {
		v1[i] = vector1[term]
		v2[i] = vector2[term]
		i++
	}

	sumSquares := 0.0
	for i := 0; i < len(terms); i++ {
		diff := v1[i] - v2[i]
		sumSquares += diff * diff
	}

	distance := math.Sqrt(sumSquares)
	similarity := 1.0 / (1.0 + distance)

	return similarity
}
