package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/plutotree/toolskit/core/rand/randnum"
)

func strToNum(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}

	return num
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/rand_num/rand/:min/:max/:num", func(c *gin.Context) {
		min := strToNum(c.Params.ByName("min"))
		max := strToNum(c.Params.ByName("max"))
		num := strToNum(c.Params.ByName("num"))

		if min >= max || num <= 0 || num > 50 || max-min > 100 {
			c.JSON(http.StatusBadRequest, gin.H{"code": 101, "err_msg": "request param error"})
			return
		}

		r, err := randnum.Rand(min, max, num)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{"code": 0, "err_msg": "succ", "result": r})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"code": -1, "err_msg": "server error"})
		}
	})

	r.GET("/rand_num/rand_perm/:min/:max/:num", func(c *gin.Context) {
		min := strToNum(c.Params.ByName("min"))
		max := strToNum(c.Params.ByName("max"))
		num := strToNum(c.Params.ByName("num"))

		if min >= max || num <= 0 || num > 50 || max-min > 100 {
			c.JSON(http.StatusBadRequest, gin.H{"code": 101, "err_msg": "request param error"})
			return
		}
		r, err := randnum.RandPerm(min, max, num)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{"code": 0, "err_msg": "succ", "result": r})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"code": -1, "err_msg": "server error"})
		}
	})

	r.GET("/rand_num/rand_pick/:min/:max/:pick/:num", func(c *gin.Context) {
		min := strToNum(c.Params.ByName("min"))
		max := strToNum(c.Params.ByName("max"))
		pick := strToNum(c.Params.ByName("pick"))
		num := strToNum(c.Params.ByName("num"))

		if min >= max || num <= 0 || num > 50 || max-min > 100 || pick > max-min || pick <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"code": 101, "err_msg": "request param error"})
			return
		}

		r, err := randnum.RandPick(min, max, pick, num)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{"code": 0, "err_msg": "succ", "result": r})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"code": -1, "err_msg": "server error"})
		}
	})

	return r
}

func main() {
	fmt.Errorf("XXXX")
	r := setupRouter()
	r.Run("127.0.0.1:12345")
}
