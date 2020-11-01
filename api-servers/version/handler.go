package version

import (
	"api-servers/es"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	m := r.Method
	if m != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	from := 0
	size := 1000
	name := strings.Split(r.URL.EscapedPath(), "/")[2]
	for {
		metas, e := es.SearchAllVersions(name, from, size)
		if e != nil {
			log.Println(e)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		for i := range metas {
			b, _ := json.Marshal(metas[i])
			w.Write(b)
			w.Write([]byte("\n"))
		}
		if len(metas) != size {
			return
		}
		from += size
	}
}
func Get(c *gin.Context) {

	from := 0
	size := 1000
	name := c.Param("name")
	for {
		metas, e := es.SearchAllVersions(name, from, size)
		if e != nil {
			log.Println(e)
			c.Writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		for i := range metas {
			b, _ := json.Marshal(metas[i])
			c.Writer.Write(b)
			c.Writer.Write([]byte("\n"))
		}
		if len(metas) != size {
			return
		}
		from += size
	}
}
