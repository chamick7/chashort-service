package url

import (
	"context"
	"net/http"
	"os"

	"github.com/chamick7/short-service/sqlc"
	"github.com/chamick7/short-service/utils/validate"
	"github.com/labstack/echo/v4"
	"github.com/teris-io/shortid"
)

type UrlService interface {
	ShortURL(ectx echo.Context) error
	RedirectTO(ectx echo.Context) error
}

type urlService struct {
	conn *sqlc.Queries
	ctx  context.Context
}

func New(ctx context.Context, conn *sqlc.Queries) *urlService {
	return &urlService{
		ctx:  ctx,
		conn: conn,
	}
}

func (s *urlService) ShortURL(ectx echo.Context) error {
	input := new(ShortRequestDTO)
	errBind := validate.BindAndValidate(ectx, input)
	if errBind != nil {
		return errBind
	}

	shortId, err := shortid.Generate()
	if err != nil {
		return ectx.JSON(http.StatusInternalServerError, map[string]string{"error": "generate shortId error"})
	}

	var shortUrl string = os.Getenv("DOMAIN") + shortId

	insertedUrl, err := s.conn.CreateShortUrl(s.ctx, sqlc.CreateShortUrlParams{
		ShortID:  shortId,
		ShortUrl: shortUrl,
		LongUrl:  input.Url,
	})
	if err != nil {
		return ectx.JSON(http.StatusInternalServerError, map[string]string{"error": "Insert error"})
	}

	return ectx.JSON(http.StatusCreated, map[string]string{"shortUrl": insertedUrl.ShortUrl})

}

func (s *urlService) RedirectTO(ectx echo.Context) error {
	shortId := ectx.Param("shortId")

	url, err := s.conn.GetUrlFromShortId(s.ctx, shortId)
	if err != nil {
		return ectx.JSON(http.StatusInternalServerError, map[string]string{"error": "Query error"})
	}

	s.conn.UpdateUrlCount(s.ctx, sqlc.UpdateUrlCountParams{ShortID: shortId, RedirectCount: url.RedirectCount + 1})

	return ectx.Redirect(http.StatusMovedPermanently, url.LongUrl)
}
