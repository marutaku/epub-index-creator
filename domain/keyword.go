package domain

type Keyword string

func (k Keyword) String() string {
	return string(k)
}
