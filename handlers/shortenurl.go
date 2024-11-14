package handlers

import (
	"net/http"
	"strings"
	"time"

	databases "github.com/mu-wahba/url-shortener/databases"
	utils "github.com/mu-wahba/url-shortener/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var body struct {
	URL              string        `json:"url" binding:"required"`
	CustomShortenURL string        `json:"custom_shorten_url"` //this is not the all url it's only the random value
	Expiry           time.Duration `json:"expiry"`             //in minutes
}

func ShortenUrl(c *gin.Context) {

	request := body
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Validation Error, url is required",
		})
		return
	}
	//RATE LIMIT
	client_ip := c.ClientIP()
	shouldpass := utils.ShouldPass(c, client_ip)
	if shouldpass != "ok" {
		return
	}

	//create keypair in redis , they key is the url and the vlue is shortenurl and TTL
	var randomId string
	if request.CustomShortenURL == "" {
		//create one
		randomId = generateRandom(6)
	} else {
		if utils.Isvalid(request.CustomShortenURL) {
			randomId = request.CustomShortenURL
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": "can't accept random id choose 6 alphanumeric chars",
			})
			return
		}

	}

	//handle capital letters
	randomId = strings.ToLower(randomId)

	var default_expiry time.Duration

	if request.Expiry.Seconds() == 0 {
		//set default
		default_expiry = 12 * 30 * 24 * time.Hour //defauls is 1 year
	} else {
		default_expiry = request.Expiry
	}

	//key:random value:url TTl:expiry
	//if exist true
	if utils.CheckInRedis(1, randomId) {
		c.JSON(http.StatusBadGateway, gin.H{
			"msg": randomId + " is taken already",
		})
	} else {
		//ADD it id db1
		r1, nil := databases.ConnectDB(1)
		err := r1.Set(databases.Ctx, randomId, request.URL, default_expiry*time.Second*60).Err()
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"err": "couldn't set in db ",
			})
			return
		} else {

			c.JSON(http.StatusOK, gin.H{
				"msg":    randomId + " is created successfully",
				"expiry": default_expiry,
			})
		}
	}

}

func generateRandom(l int) string {
	return uuid.New().String()[:l]
}
