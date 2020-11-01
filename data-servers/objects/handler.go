package objects

import (
	"data-servers/config"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
)

func Get(c *gin.Context)  {
	name := c.Param("name")

	file, err := os.Open(config.StorageRoot+name)
	if err != nil {
		log.Println(err)
		c.Writer.WriteHeader(http.StatusNotFound)
		return
	}
	defer file.Close()
	io.Copy(c.Writer, file)
}
func Post(c *gin.Context)  {
	name := c.Param("name")
	file, err := os.Create(config.StorageRoot+name)
	if err != nil {
		log.Println(err)
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer file.Close()
	io.Copy(file, c.Request.Body)
}
