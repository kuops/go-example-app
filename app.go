package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

var (
	httpRequestsTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Count of all HTTP requests",
	}, []string{"uri", "code", "method"})
)

type postData struct {
	Name string `json:"name" binding:"required"`
}

type secret struct {
	Secret string `yaml:"secret"`
}

func init() {
	_ = prometheus.Register(httpRequestsTotal)
}

func main() {
	r := gin.Default()
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	api := r.Group("/api")
	api.Handle(http.MethodGet, "/hello", helloHandler)
	api.Handle(http.MethodPost, "/hello", helloHandler)
	r.Run(":8080")
}

func getSecret(c string) string {
	var err error
	secret := secret{}
	secretconfig, err := ioutil.ReadFile(c)
	if err != nil {
		log.Fatalf("readfile %s error: %v",c,err)
		return ""
	}
	err = yaml.Unmarshal(secretconfig, &secret)
	if err != nil {
		log.Fatalf("unmarshal err: %v", err)
		return ""
	}
	return secret.Secret
}

func helloHandler(c *gin.Context) {
	secret := getSecret("./config/secret.yaml")
	var postdata = postData{}
	if c.Request.Method == "GET" {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "unsupport GET method",
		})
		httpRequestsTotal.WithLabelValues(
			c.Request.RequestURI, strconv.Itoa(http.StatusNotFound), c.Request.Method)
	} else {
		if err := c.BindJSON(&postdata); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"name":   postdata.Name,
				"secret": secret,
			})
			httpRequestsTotal.WithLabelValues(
				c.Request.RequestURI, strconv.Itoa(http.StatusOK), c.Request.Method)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "parameter not right",
			})
			httpRequestsTotal.WithLabelValues(
				c.Request.RequestURI, strconv.Itoa(http.StatusBadRequest), c.Request.Method)
		}
	}
}
