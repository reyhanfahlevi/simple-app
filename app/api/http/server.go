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
	"github.com/reyhanfahlevi/simple-app/app"
	"github.com/reyhanfahlevi/simple-app/app/api/http/blog"
)

var (
	blogHTTP *blog.Handler
)

// Init will initialize this http package
func Init(bloguc app.BlogUseCase) {
	blogHTTP = blog.New(bloguc)
}

// Run run http server
func Run(port string) {
	r := gin.Default()

	/* Register router here */

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "You know, for learning...")
	})

	SpawnBlogHTTPHandler(r)

	/* End of registering router */

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

// SpawnBlogHTTPHandler will spawn group of blog http handler
// will have pattern /blog/....
func SpawnBlogHTTPHandler(r *gin.Engine) {
	group := r.Group("/blog")

	blogHTTP.HandlerCreateBlogPost(group)
	blogHTTP.HandlerUpdateBlogPost(group)
	blogHTTP.HandlerGetBlogPost(group)
}
