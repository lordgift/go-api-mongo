package service

import (
	"merchants/util"
	"merchants/persistence"
	"time"
	"net/http"
	"log"
	"fmt"

	"gopkg.in/mgo.v2"
	"github.com/gin-gonic/gin"
)

func noSslConnect() (*mgo.Session, error) {
	
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
	merchantService persistence.MerchantService
	productService persistence.ProductService
	sellService persistence.SellService
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

	usr := util.RandStringRunes(5);
	pwd := util.RandStringRunes(10);
	log.Printf("** Generate Authorization **\nUsername : %s\nPassword : %s ", 
usr,pwd);

	root := r.Group("/")
	root.POST("/register", s.register)
	root.POST("/buy/product", s.all)

	merchant := root.Group("/merchant")
	merchant.Use(gin.BasicAuth(gin.Accounts{
		"admin": "1234",
		//FIXME use generated password
		//usr:pwd,
	}))
	merchant.GET("/:id", s.all)
	merchant.POST("/:id", s.all)
	merchant.GET("/:id/products", s.all)
	merchant.POST("/:id/product", s.all)
	merchant.POST("/:id/product/:product_id", s.all)
	merchant.DELETE("/:id/product/:product_id", s.all)
	merchant.POST("/:id/report", s.all)


	return r
}

func (s *Service) register(c *gin.Context) {

	var register persistence.Merchant
	err := c.ShouldBindJSON(&register)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	isDuplicated,err := s.merchantService.IsDuplicatedBankAccount(register.BankAccount)
	if err != nil || isDuplicated {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	log.Printf("bankAccount:%s \n isDup:%v", register.BankAccount, isDuplicated)
	if (!isDuplicated) {
		
		register.Username = util.RandStringRunes(5);
		register.Password = util.RandStringRunes(10);

		s.merchantService.Register(register)
	} 


}


func (s *Service) all(c *gin.Context) {
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