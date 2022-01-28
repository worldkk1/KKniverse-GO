package coin_hunter

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BrazePayloadEvent struct {
	ExternalID string                 `json:"external_id"`
	Name       string                 `json:"name"`
	Time       string                 `json:"time"`
	Properties map[string]interface{} `json:"properties"`
}

type BrazePayload struct {
	APIKey     string                   `json:"api_key"`
	Events     []BrazePayloadEvent      `json:"events"`
	Attributes []map[string]interface{} `json:"attributes"`
}

func readBody(reader io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)

	s := buf.String()
	return s
}

func Braze(c *gin.Context) {
	var input BrazePayload
	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Println("[debug]err.Error()", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("[debug]input", input)
	fmt.Println("[debug]props", input.Events[0].Properties)
	if input.Events[0].Name == "Treasurehunt - Play game" {
		amount := input.Events[0].Properties["Amount"]
		level := input.Events[0].Properties["Level"]
		fmt.Println("[debug]amount", amount)
		fmt.Println("[debug]level", level)
	}
	// buf, _ := ioutil.ReadAll(c.Request.Body)
	// rdr1 := ioutil.NopCloser(bytes.NewBuffer(buf))
	// rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf)) //We have to create a new Buffer, because rdr1 will be read.

	// fmt.Println(readBody(rdr1)) // Print request body

	// c.Request.Body = rdr2

	c.JSON(http.StatusOK, "")
}
