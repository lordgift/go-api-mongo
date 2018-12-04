package service

import (
	"merchants/persistence"
	"time"
	"net/http"
	"fmt"

	"gopkg.in/mgo.v2"
	"github.com/gin-gonic/gin"
)

func noSslConnect() (*mgo.Session, error) {
	// session, err := mgo.Dial("localhost:27017")

	info := &mgo.DialInfo{
		Addrs:    []string{"localhost"},
		Timeout:  60 * time.Second,
		Database: "merchants",
		Username: "lordgift",
		Password: "gotraining",
	}

	session, err := mgo.DialWithInfo(info)
	if err != nil {
		panic(err)
	}
	return session, err
}

type Service struct {
	// collection      *mgo.Collection
	merchantService persistence.MerchantService
	productService persistence.ProductService
	sellService persistence.SellService
	// bankAccountService persistence.BankAccountService
}

type Merchant struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	BankAccount string `json:"bank_account"`
}
func CreateService() *Service {
	session, _ := noSslConnect()
	db := session.DB("merchants")
	// err := db.C("test").Insert(Merchant{1,"MYSHOP", "xxxxx-x"})
	// if err != nil {
	// 	panic(err)
	// }

	s := &Service{
		// collection: c,
		merchantService: &persistence.MerchantServiceImp{
			Collection: db.C("merchants"),
		},
		productService: &persistence.ProductServiceImp{
			Collection: db.C("products"),
		},
		
		sellService: &persistence.SellServiceImp{
			Collection: db.C("sell"),
		},
	}

	return s
}

func SetupRoute(s *Service) *gin.Engine {
	r := gin.Default()

	root := r.Group("/")
	root.POST("/merchant/register", s.All)
	root.GET("/merchants/:id", s.All)
	root.POST("/merchants/:id", s.All)
	root.GET("/merchants/:id/products", s.All)
	root.POST("/merchants/:id/product", s.All)
	root.POST("/merchants/:id/product/:product_id", s.All)
	root.DELETE("/merchants/:id/product/:product_id", s.All)
	root.POST("/merchants/:id/report", s.All)
	root.POST("/buy/product", s.All)

	return r
}

func (s *Service) All(c *gin.Context) {
	merchants, err := s.merchantService.All()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"object":  "error",
			"message": fmt.Sprintf("db: query error: %s", err),
		})
		return
	}
	c.JSON(http.StatusOK, merchants)
}

// func (s *Server) All(c *gin.Context) {
// 	users, err := s.userService.All()
// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
// 			"object":  "error",
// 			"message": fmt.Sprintf("db: query error: %s", err),
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusOK, users)
// }

// func (s *Server) FindByID(c *gin.Context) {
// 	id, _ := strconv.Atoi(c.Param("id"))
// 	users, err := s.userService.FindByID(id)
// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
// 			"object":  "error",
// 			"message": fmt.Sprintf("db: query error: %s", err),
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusOK, users)
// }
