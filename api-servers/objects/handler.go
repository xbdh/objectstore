package objects

import (
	"api-servers/es"
	"api-servers/heartbeat"
	"api-servers/locate"
	"api-servers/objectstream"
	"api-servers/utils"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
)

func putStream(name string) (*objectstream.PutStream, error) {
	server := heartbeat.ChooseRandomDataServer()
	if server == "" {
		return nil, fmt.Errorf("cannot find any dataServer")
	}

	return objectstream.NewPutStream(server, name), nil
}

func storeObject(r io.Reader, name string) (int, error) {
	stream, err := putStream(name)
	if err != nil {
		return http.StatusServiceUnavailable, err
	}

	io.Copy(stream, r)
	err = stream.Close()
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

func getStream(name string) (io.Reader, error) {
	server := locate.Locate(name)
	if server == "" {
		return nil, fmt.Errorf("object %s locate fail", name)
	}
	return objectstream.NewGetStream(server, name)
}

func Get(c *gin.Context) {
	name := c.Param("name")
	//versionId := r.URL.Query()["version"]
	versionId := c.Query("version")
	version := 0
	var e error
	if len(versionId) != 0 {
		version, e = strconv.Atoi(versionId)
		if e != nil {
			log.Println(e)
			c.Writer.WriteHeader(http.StatusBadRequest)
			return
		}
	}
	meta, e := es.GetMetadata(name, version)
	if e != nil {
		log.Println(e)
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	if meta.Hash == "" {
		c.Writer.WriteHeader(http.StatusNotFound)
		return
	}
	object := url.PathEscape(meta.Hash)
	stream, e := getStream(object)
	if e != nil {
		log.Println(e)
		c.Writer.WriteHeader(http.StatusNotFound)
		return
	}
	io.Copy(c.Writer, stream)
}

func Post(c *gin.Context) {
	hash := utils.GetHashFromHeader(c.Request.Header)
	if hash == "" {
		log.Println("missing object hash in digest header")
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	code, e := storeObject(c.Request.Body, url.PathEscape(hash))
	if e != nil {
		log.Println(e)
		c.Writer.WriteHeader(code)
		return
	}
	if code != http.StatusOK {
		c.Writer.WriteHeader(code)
		return
	}

	name := c.Param("name")
	size := utils.GetSizeFromHeader(c.Request.Header)
	e = es.AddVersion(name, hash, size)
	if e != nil {
		log.Println(e)
		c.Writer.WriteHeader(http.StatusInternalServerError)
	}
}
func Delete(c *gin.Context) {
	name := c.Param("name")
	version, e := es.SearchLatestVersion(name)
	if e != nil {
		log.Println(e)
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	e = es.PutMetadata(name, version.Version+1, 0, "")
	if e != nil {
		log.Println(e)
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

}
