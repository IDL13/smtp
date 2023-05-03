package unmarshal

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/labstack/echo/v4"
)

func New() *Context {
	return &Context{}
}

type Context struct {
	Email string `json:"email"`
	Msg   string `json:"msg"`
}

func (n *Context) Unmarshal(c echo.Context) *Context {

	defer c.Request().Body.Close()
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Printf("Failed reading the request body: %s", err)
	}

	err = json.Unmarshal(b, &n)
	if err != nil {
		log.Println(err)
	}

	return n
}
