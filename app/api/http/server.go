package http

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/reyhanfahlevi/simple-app/app/api/http/blog"
	"github.com/reyhanfahlevi/simple-app/internal/service/post"
	"github.com/reyhanfahlevi/simple-app/internal/service/user"
	blogSvc "github.com/reyhanfahlevi/simple-app/internal/usecase/blog"
)

// Run run http server
func Run(port string) {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "You know, for learning...")
	})

	SpawnBlogHTTPHandler(r)

	srv := &http.Server{
		Addr:    port,
		Handler: r,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Graceful Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}

	log.Println("Server exiting")
}

// SpawnBlogHTTPHandler will spawn blog http handler
func SpawnBlogHTTPHandler(r *gin.Engine) {
	postSvc := post.New()
	userSvc := user.New()

	blogService := blogSvc.New(postSvc, userSvc)
	blogHttp := blog.New(blogService)

	group := r.Group("/blog")

	blogHttp.HandlerCreateBlogPost(group)
	blogHttp.HandlerUpdateBlogPost(group)
	blogHttp.HandlerGetBlogPost(group)
}
