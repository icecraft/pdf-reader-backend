package api

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/icecraft/pdf-reader-backend/pkg/log"
	"github.com/icecraft/pdf-reader-backend/pkg/types"
)

func (s *Server) GetCiba(c *gin.Context) {
	word := strings.ToLower(c.Param("word"))
	ret := types.TranslateResp{Word: word, EN: []string{}, Examples: []string{}, Synomyms: []string{}}

	resp, docId, err := s.es.RetrieveWordTrans(word)
	if err != nil {
		log.Error(err, "failed to get translate word meaning", "word", c.Param("word"))
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	// means not found from es
	if resp == nil {
		cibaResp, err := s.ciba.Fetch(word)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		// Synomyms
		syns := make([]string, 0)
		for _, v := range cibaResp.Message.Synonym {
			for _, vv := range v.Means {
				for i := range vv.Cis {
					syns = append(syns, vv.Cis[i])
				}
			}
		}
		ret.Synomyms = syns

		//
		strCn := make([]string, 0)
		strEn := make([]string, 0)
		examples := make([]string, 0)
		for _, v := range cibaResp.Message.Collins {
			for _, vv := range v.Entry {
				strCn = append(strCn, vv.Tran)
				strEn = append(strEn, vv.Def)
				for _, example := range vv.Example {
					examples = append(examples, example.Ex)
				}
			}
		}
		if len(strCn) > 3 {
			strCn = strCn[:3]
		}

		ret.CN = strings.Join(strCn, ", ")
		ret.EN = strEn
		ret.Examples = examples

		// put to es
		resp = &types.RetrieveItem{
			Word:     word,
			CN:       ret.CN,
			EN:       ret.EN,
			Examples: ret.Examples,
			Synomyms: ret.Synomyms,
		}
		if err := s.es.Put(resp); err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

	} else {
		if err := s.es.IncrHit(docId); err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		ret.CN = resp.CN
		ret.EN = resp.EN
		ret.Examples = resp.Examples
		ret.Synomyms = resp.Synomyms
	}

	if len(ret.EN) > 3 {
		ret.EN = ret.EN[:3]
	}
	if len(ret.Examples) > 3 {
		ret.Examples = ret.Examples[:3]
	}
	if len(ret.Synomyms) > 3 {
		ret.Synomyms = ret.Synomyms[:3]
	}

	// put to es, or update es hit infomation
	if ret.EN == nil {
		ret.EN = []string{}
	}
	if ret.Examples == nil {
		ret.Examples = []string{}
	}
	if ret.Synomyms == nil {
		ret.Synomyms = []string{}
	}

	c.JSON(http.StatusOK, ret)
}
