package usecase

type KeywordUsecaseInterface interface {
	ListKeywords() ([]string, error)
	ListKeywordsByPageID(bookID int) ([]string, error)
}

type keywordUsecase struct{}

func NewKeywordUsecase() KeywordUsecaseInterface {
	return &keywordUsecase{}
}

func (*keywordUsecase) ListKeywords() ([]string, error) {
	return nil, nil
}

func (*keywordUsecase) ListKeywordsByPageID(bookID int) ([]string, error) {
	return nil, nil
}
