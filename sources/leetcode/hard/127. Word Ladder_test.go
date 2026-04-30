package hard

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ladderLength(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		beginWord string
		endWord   string
		wordList  []string
		want      int
	}{
		{
			// hit -> hot -> dot -> dog -> cog
			name:      "classic example from problem statement",
			beginWord: "hit",
			endWord:   "cog",
			wordList:  []string{"hot", "dot", "dog", "lot", "log", "cog"},
			want:      5,
		},
		{
			name:      "endWord not in wordList",
			beginWord: "hit",
			endWord:   "cog",
			wordList:  []string{"hot", "dot", "dog", "lot", "log"},
			want:      0,
		},
		{
			name:      "no path exists endWord in list but unreachable",
			beginWord: "hit",
			endWord:   "xyz",
			wordList:  []string{"abc", "xyz"},
			want:      0,
		},
		{
			name:      "beginWord directly transforms to endWord chain length two",
			beginWord: "hot",
			endWord:   "dot",
			wordList:  []string{"dot", "lot"},
			want:      2,
		},
		{
			name:      "single path through multiple intermediate words",
			beginWord: "red",
			endWord:   "tax",
			wordList:  []string{"ted", "tex", "red", "tax", "tad", "den", "rex", "pee"},
			want:      4,
		},
		{
			name:      "wordList contains beginWord does not affect result",
			beginWord: "hit",
			endWord:   "cog",
			wordList:  []string{"hit", "hot", "dot", "dog", "lot", "log", "cog"},
			want:      5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := ladderLength(tt.beginWord, tt.endWord, tt.wordList)
			assert.Equal(t, tt.want, got)
		})
	}
}
