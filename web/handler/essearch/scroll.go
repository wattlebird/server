package essearch

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/trim21/errgo"

	"github.com/bangumi/server/web/accessor"
	"github.com/bangumi/server/web/res"
)

func (h Handler) Scroll(c echo.Context) error {
	user := accessor.GetFromCtx(c)
	if !user.Login {
		return res.Unauthorized("Requires valid user token")
	}

	scrollId := c.QueryParam("scroll_id")

	r, err := h.search.Scroll(c.Request().Context(), scrollId)
	if err != nil {
		return errgo.Wrap(err, fmt.Sprintf("Scroll on search result %s failed", scrollId))
	}

	rtn := res.SearchResult{
		Total:    r.Metadata.Total,
		Took:     r.Metadata.Took,
		Timedout: r.Metadata.Timedout,
		ScrollID: r.Metadata.ScrollID,
	}
	if len(r.Celebrities) > 0 {
		rtn.Data, err = h.extractCelebrity(c.Request().Context(), &r)
		if err != nil {
			return err
		}
	} else if len(r.Subjects) > 0 {
		rtn.Data, err = h.extractSubject(c.Request().Context(), &r, user)
		if err != nil {
			return err
		}
	}

	return c.JSON(http.StatusOK, r)
}
