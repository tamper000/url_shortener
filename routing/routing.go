package routing

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
	"url_shortener/database"

	"github.com/labstack/echo/v4"
	"github.com/speps/go-hashids/v2"
)

func GetShorten(c echo.Context) error {
	shortID := c.Param("shortID")

	long_url := database.Get(shortID)
	if long_url == "" {
		c.Redirect(http.StatusPermanentRedirect, "/")
	}

	return c.Redirect(http.StatusPermanentRedirect, database.Get(shortID))
}

func Create(c echo.Context) error {
	long_url := c.FormValue("url")
	u, err := url.Parse(long_url)
	if u.Scheme == "" || u.Host == "" || !strings.Contains(u.Host, ".") || err != nil {
		return c.String(http.StatusBadRequest, "error")
	}

	hd := hashids.NewData()
	hd.Salt = long_url
	id_generator, _ := hashids.NewWithData(hd)
	short_id, _ := id_generator.Encode([]int{int(time.Now().Unix())})

	database.Add(short_id, long_url)

	data := map[string]string{
		"id":  short_id,
		"url": fmt.Sprintf("%s://%s/%s", c.Scheme(), c.Request().Host, short_id),
	}
	return c.JSON(http.StatusOK, data)
}
