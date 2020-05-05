package model

type Article struct {
	ID          int
	Title, Body string
}

type ArticleStoreInMemory struct {
	ArticleMap []Article
}

func NewArticleStoreInMemory() *ArticleStoreInMemory {
	return &ArticleStoreInMemory{
		ArticleMap: []Article{
			Article{ID: 1, Title: "Membuat Website", Body: "Hallo"},
		},
	}
}

func CreateArticle(title, body string) (*Article, error) {
	return &Article{
		Title: title,
		Body:  body,
	}, nil
}

func (store *ArticleStoreInMemory) Save(article *Article) error {
	// 1. calculate current length
	lastID := len(store.ArticleMap)

	// set article id
	article.ID = lastID + 1

	// push to article map slices
	store.ArticleMap = append(store.ArticleMap, *article)

	return nil
}

func (store *ArticleStoreInMemory) Remove(id int) error {
	store.ArticleMap = append(store.ArticleMap[:id-1], store.ArticleMap[id:]...)
	return nil
}
func (store *ArticleStoreInMemory) Update(title, body string, id int) error {
	store.ArticleMap[id-1] = Article{ID: id, Title: title, Body: body}
	return nil
}
