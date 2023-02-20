package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gitlab.shlab.tech/xurui/pdf-reader-backend/pkg/log"
)

func (s *Server) GetCiba(c *gin.Context) {
	resp, err := s.ciba.Fetch(c.Param("word"))
	if err != nil {
		log.Error(err, "failed to get translate word meaning", "word", c.Param("word"))
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	//Get from es
	// if not then get from ciba

	// put to es, or update es hit infomation

	c.JSON(http.StatusOK, resp)
}
