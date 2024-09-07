package essearch

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/trim21/errgo"

	"github.com/bangumi/server/web/accessor"
	"github.com/bangumi/server/web/res"
)

func (h Handler) AutoComplete(c echo.Context) error {
	user := accessor.GetFromCtx(c)
	if !user.Login {
		return res.Unauthorized("Requires valid user token")
	}

	prefix := c.QueryParam("prefix")
	typ := c.QueryParam("type")

	r, err := h.search.AutoComplete(c.Request().Context(), prefix, typ)
	if err != nil {
		return errgo.Wrap(err, "AutoComplete failed")
	}

	return c.JSON(http.StatusOK, r)
}
