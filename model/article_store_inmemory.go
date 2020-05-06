package model

type ArticleStoreInmemory struct {
	ArticleMap []Article
}

// func NewArticleStoreInmemory() ArticleStoreInmemory {
func NewArticleStoreInmemory() *ArticleStoreInmemory {
	return &ArticleStoreInmemory{
		ArticleMap: []Article{
			Article{ID: 1, Title: "Membuat website", Body: "Hallo ini bayu"},
			Article{ID: 2, Title: "Membuat website 2", Body: "Hallo ini bayu 2"},
		},
	}
}

func (store *ArticleStoreInmemory) Find(id int) *Article {
	article := Article{}
	for _, item := range store.ArticleMap {
		if item.ID == id {
			article = item
		}
	}

	return &article
}

func (store *ArticleStoreInmemory) Update(article *Article) error {
	for index, item := range store.ArticleMap {
		if item.ID == article.ID {
			store.ArticleMap[index] = *article
		}
	}

	return nil
}

func (store *ArticleStoreInmemory) Delete(article *Article) error {
	articles := []Article{}

	for _, item := range store.ArticleMap {
		if item.ID != article.ID {
			articles = append(articles, item)
		}
	}

	store.ArticleMap = articles
	return nil
}
