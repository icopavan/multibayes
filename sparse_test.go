package multibayes

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestSparseBag(t *testing.T) {
	testdata := getTestData()
	tokenizer, err := NewTokenizer(&TokenizerConf{
		NGramSize: 1,
	})
	assert.Equalf(t, err, nil, "Error creating new tokenizer")

	sparse := NewSparseMatrix()

	for _, document := range testdata {
		ngrams := tokenizer.Parse(document.Text)
		sparse.Add(ngrams, document.Classes)
	}
}

func TestToFeatureMatrix(t *testing.T) {
	testdata := getTestData()
	tokenizer, err := NewTokenizer(&TokenizerConf{
		NGramSize: 1,
	})
	assert.Equalf(t, err, nil, "Error creating new tokenizer")

	sparse := NewSparseMatrix()

	for _, document := range testdata {
		ngrams := tokenizer.Parse(document.Text)
		sparse.Add(ngrams, document.Classes)
	}

	matrices := sparse.ToFeatureMatrix()

	assert.Equal(t, len(sparse.Classes), len(matrices["classes"].Data), "Wrong length")
	assert.Equal(t, len(sparse.Tokens), len(matrices["tokens"].Data), "Wrong length")
}