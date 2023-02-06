package main

import (
	"context"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/url"

	pb "pricetracker/pkg/build/pkg/proto"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type LoginForm struct {
	UserName string `form:"username"`
}

type AddProductForm struct {
	Shop string `form:"shop"`
	Name string `form:"name"`
	Url  string `form:"url"`
}

type ChooseProductForm struct {
	Idx int `form:"idx"`
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
		var username string = session.Get("username").(string)

		list, err := user.GetProducts(context.Background(), &pb.User{
			Name: username,
		})

		if err == nil {
			c.HTML(http.StatusOK, "track.html.tmpl", gin.H{
				"sessionUsername": session.Get("username"),
				"products":        list.GetProductsList(),
			})
		} else {
			log.WithError(err).WithField("username", username).Info("could not fetched products list for user")
			c.AbortWithError(http.StatusInternalServerError, err)
		}
	})

	r.POST("/chooseProduct", func(c *gin.Context) {
		session := sessions.Default(c)
		var username string = session.Get("username").(string)

		chooseProductForm := ChooseProductForm{}
		if err := c.ShouldBind(&chooseProductForm); err == nil {
			list, err := user.GetProducts(context.Background(), &pb.User{
				Name: username,
			})

			if err == nil {
				product := list.GetProductsList()[chooseProductForm.Idx]
				log.WithField("product", product).WithField("idx", chooseProductForm.Idx).Info("choosed product")

				resp, err := http.Get(fmt.Sprintf("%s?shop=%s&name=%s&url=%s", diagramGenerator, url.QueryEscape(product.Shop), url.QueryEscape(product.Name), url.QueryEscape(product.Url)))

				if err == nil {
					defer resp.Body.Close()
					body, err := io.ReadAll(resp.Body)
					if err == nil {
						bodyStr := string(body[:])
						c.HTML(http.StatusOK, "track.html.tmpl", gin.H{
							"sessionUsername": session.Get("username").(string),
							"products":        list.GetProductsList(),
							"chart":           template.HTML(bodyStr),
						})
					} else {
						log.WithError(err).WithField("product", product).Info("could not read diagram response for product")
						c.Redirect(http.StatusTemporaryRedirect, "/track")
					}
				} else {
					log.WithError(err).WithField("product", product).Info("could not fetch diagram for product")
					c.Redirect(http.StatusTemporaryRedirect, "/track")
				}
			} else {
				log.WithError(err).WithField("username", username).Info("could not fetched products list for user")
				c.Redirect(http.StatusTemporaryRedirect, "/products")
			}
		} else {
			log.WithError(err).Error("unable to parse chooseProduct form")
			c.Redirect(http.StatusTemporaryRedirect, "/track")
		}
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

	r.Run(":8080")
}

const users = "10.104.130.163:8081"

const diagramGenerator = "http://10.104.130.165:8085/product"
