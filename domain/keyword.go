package domain

import "github.com/marutaku/epub-index-creator/ent"

type Keyword struct {
	ID      int
	Keyword string
}

func NewKeywordFromEnt(entKeyword *ent.Keyword) *Keyword {
	return &Keyword{
		ID:      entKeyword.ID,
		Keyword: entKeyword.Keyword,
	}
}
