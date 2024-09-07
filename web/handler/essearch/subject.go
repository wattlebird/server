package essearch

import (
	"net/http"

	"github.com/bangumi/server/web/accessor"
	"github.com/bangumi/server/web/req"
	"github.com/bangumi/server/web/res"
	"github.com/labstack/echo/v4"
	"github.com/trim21/errgo"

	"go.uber.org/zap"
)

func (h Handler) Subject(c echo.Context) error {
	user := accessor.GetFromCtx(c)
	if !user.Login {
		return res.Unauthorized("Requires valid user token")
	}

	var reqData req.SubjectQuery
	if err := c.Echo().JSONSerializer.Deserialize(c, &reqData); err != nil {
		return res.JSONError(c, err)
	}
	dateRange := make(map[string]string)
	if reqData.DateRange.Gte != "" {
		dateRange["gte"] = reqData.DateRange.Gte
	}
	if reqData.DateRange.Lte != "" {
		dateRange["lte"] = reqData.DateRange.Lte
	}

	r, err := h.search.Subject(c.Request().Context(), reqData.Query, reqData.Type, reqData.Tags, dateRange)
	if err != nil {
		h.log.Error("Performing subject search failed", zap.Error(err))
		return errgo.Wrap(err, "Performing subject search failed")
	}

	data, err := h.extractSubject(c.Request().Context(), &r, user)
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
