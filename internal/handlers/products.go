package handlers

/*
import (
	"net/http"

	"github.com/PubliusVergilius/ecom-go-api/internal/services"
)



var Module = fx.Module(
	"server",
	fx.Provide(
		New,
	),
)

Route("/articles", func(r chi.Router) {
    r.With(paginate).Get("/", listArticles)                           // GET /articles
    r.With(paginate).Get("/{month}-{day}-{year}", listArticlesByDate) // GET /articles/01-16-2017

    r.Post("/", createArticle)                                        // POST /articles
    r.Get("/search", searchArticles)                                  // GET /articles/search

    // Regexp url parameters:
    r.Get("/{articleSlug:[a-z-]+}", getArticleBySlug)                // GET /articles/home-is-toronto

    // Subrouters:
    r.Route("/{articleID}", func(r chi.Router) {
      r.Use(ArticleCtx)
      r.Get("/", getArticle)                                          // GET /articles/123
      r.Put("/", updateArticle)                                       // PUT /articles/123
      r.Delete("/", deleteArticle)                                    // DELETE /articles/123
    })
})

func ArticleCtx(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    articleID := chi.URLParam(r, "articleID")
    article, err := dbGetArticle(articleID)
    if err != nil {
      http.Error(w, http.StatusText(404), 404)
      return
    }
    ctx := context.WithValue(r.Context(), "article", article)
    next.ServeHTTP(w, r.WithContext(ctx))
  })
}

func getArticle(w http.ResponseWriter, r *http.Request) {
  ctx := r.Context()
  article, ok := ctx.Value("article").(*Article)
  if !ok {
    http.Error(w, http.StatusText(422), 422)
    return
  }
  w.Write([]byte(fmt.Sprintf("title:%s", article.Title)))
}

// A completely separate router for administrator routes
func adminRouter() http.Handler {
  r := chi.NewRouter()
  r.Use(AdminOnly)
  r.Get("/", adminIndex)
  r.Get("/accounts", adminListAccounts)
  return r
}

func AdminOnly(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()
    perm, ok := ctx.Value("acl.permission").(YourPermissionType)
    if !ok || !perm.IsAdmin() {
      http.Error(w, http.StatusText(403), 403)
      return
    }
    next.ServeHTTP(w, r)
  })
}
*/
