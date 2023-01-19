package main

import (
	"context"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "pricetracker/pkg/build/pkg/proto"

	log "github.com/sirupsen/logrus"
)

type LoginForm struct {
	UserName string `form:"username"`
}

type AddProductForm struct {
	Shop string `form:"shop"`
	Name string `form:"name"`
	Url  string `form:"url"`
}

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	sdconn, err := grpc.Dial(users, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.WithError(err).Fatal("unable to connect to service dispatcher")
		// FIXME add retry policy
	}
	defer sdconn.Close()
	user := pb.NewUsersClient(sdconn)

	r.GET("/", func(c *gin.Context) {
		session := sessions.Default(c)
		if session.Get("username") != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/products")
		} else {
			c.HTML(http.StatusOK, "index.html.tmpl", gin.H{})
		}
	})

	r.GET("/products", func(c *gin.Context) {
		session := sessions.Default(c)
		var username string = session.Get("username").(string)

		list, err := user.GetProducts(context.Background(), &pb.User{
			Name: username,
		})

		if err == nil {
			c.HTML(http.StatusOK, "products.html.tmpl", gin.H{
				"sessionUsername": session.Get("username"),
				"products":        list.GetProductsList(),
			})
		} else {
			log.WithError(err).WithField("username", username).Info("could not fetched products list for user")
			c.AbortWithError(http.StatusInternalServerError, err)
		}
	})

	r.POST("/addProduct", func(c *gin.Context) {
		session := sessions.Default(c)
		var username string = session.Get("username").(string)

		addProductForm := AddProductForm{}
		if err := c.ShouldBind(&addProductForm); err == nil {
			list, err := user.AddProduct(context.Background(), &pb.UserProduct{
				User: &pb.User{
					Name: username,
				},
				Product: &pb.Product{
					Shop: addProductForm.Shop,
					Name: addProductForm.Name,
					Url:  addProductForm.Url,
				},
			})
			if err == nil {
				c.HTML(http.StatusOK, "products.html.tmpl", gin.H{
					"sessionUsername": session.Get("username"),
					"products":        list.GetProductsList(),
				})
			} else {
				log.WithError(err).WithField("addProductForm", addProductForm).Info("could not add new product")
				c.Redirect(http.StatusTemporaryRedirect, "/products")
			}
		} else {
			log.WithError(err).Error("unable to parse addProduct form")
			c.Redirect(http.StatusTemporaryRedirect, "/products")
		}
	})

	r.GET("/track", func(c *gin.Context) {
		session := sessions.Default(c)
		c.HTML(http.StatusOK, "track.html.tmpl", gin.H{
			"sessionUsername": session.Get("username"),
		})
	})

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html.tmpl", gin.H{})
	})

	r.POST("/login", func(c *gin.Context) {
		session := sessions.Default(c)
		loginForm := LoginForm{}
		if err := c.ShouldBind(&loginForm); err == nil {
			session.Set("username", loginForm.UserName)
			session.Save()
			c.Redirect(http.StatusMovedPermanently, "/")
		} else {
			log.WithError(err).Error("unable to parse username from form")
			c.Redirect(http.StatusTemporaryRedirect, "/login")
		}
	})

	r.GET("/logout", func(c *gin.Context) {
		session := sessions.Default(c)
		session.Clear()
		session.Save()
		c.Redirect(http.StatusTemporaryRedirect, "/")
	})

	r.Run("localhost:8080")
}

const users = "localhost:8081"
