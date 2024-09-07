package essearch

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/trim21/errgo"
	"go.uber.org/zap"

	"github.com/bangumi/server/web/accessor"
	"github.com/bangumi/server/web/req"
	"github.com/bangumi/server/web/res"
)

func (h Handler) Celebrity(c echo.Context) error {
	user := accessor.GetFromCtx(c)
	if !user.Login {
		return res.Unauthorized("Requires valid user token")
	}

	var reqData req.CelebrityQuery
	if err := c.Echo().JSONSerializer.Deserialize(c, &reqData); err != nil {
		return res.JSONError(c, err)
	}

	r, err := h.search.Celebrity(c.Request().Context(), reqData.Query, reqData.Type)
	if err != nil {
		h.log.Error("Performing celebrity search failed", zap.Error(err))
		return errgo.Wrap(err, "Performing celebrity search failed")
	}

	data, err := h.extractCelebrity(c.Request().Context(), &r)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res.SearchResult{
		Total:    r.Metadata.Total,
		Took:     r.Metadata.Took,
		Timedout: r.Metadata.Timedout,
		ScrollID: r.Metadata.ScrollID,
		Data:     data,
	})
}
