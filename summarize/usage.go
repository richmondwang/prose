package summarize

// WordDensity returns a map of each word and its density.
func (d *Document) WordDensity() map[string]float64 {
	density := make(map[string]float64)
	for word, stats := range d.Words {
		density[word] = float64(stats[0]) / d.NumWords
	}
	return density
}

// MeanWordLength returns the mean number of characters per word.
func (d *Document) MeanWordLength() float64 {
	return d.NumCharacters / d.NumWords
}
