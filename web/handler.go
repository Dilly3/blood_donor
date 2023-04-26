package web

import (
	"encoding/json"
	"io/ioutil"

	"github.com/dilly3/blood-donor/internal/models"
	"github.com/gin-gonic/gin"
)

func (cfig *Config) Test() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(200, "Welcome back")
	}
}

func (cfig *Config) HandleSaveCandidate() func(c *gin.Context) {
	return func(c *gin.Context) {
		cand := models.Candidate{}
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			cfig.Logger.Error(err.Error())
		}
		json.Unmarshal(body, &cand)
		b, err := cfig.Serv.SaveCandidate(cand)
		if err != nil {
			cfig.Logger.Error(err.Error())
		}
		c.JSON(201, b)
	}
}

func (cfig *Config) HandleGetAllCandidates() func(c *gin.Context) {
	return func(c *gin.Context) {
		cands, err := cfig.Serv.GetAllCandidates()
		if err != nil {
			cfig.Logger.Error(err.Error())
		}
		c.JSON(200, gin.H{"candidates": cands})
	}
}
func (cfig *Config) HandleGEtByNAme() func(c *gin.Context) {
	return func(c *gin.Context) {
		fullname := c.Param("fullname")
		cand, err := cfig.Serv.GetByFullname(fullname)
		if err != nil {
			cfig.Logger.Error(err.Error())
		}
		c.JSON(200, gin.H{cand.Id: cand})
	}
}

func (cfig *Config) HandleGEtById() func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		cand, err := cfig.Serv.GetById(id)
		if err != nil {
			cfig.Logger.Error(err.Error())
		}
		c.JSON(200, gin.H{cand.Id: cand})
	}
}
