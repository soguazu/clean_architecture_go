package main

import (
	"log"

	"github.com/soguazu/clean_arch/controllers"
	router "github.com/soguazu/clean_arch/http"
	"github.com/soguazu/clean_arch/repository"
	"github.com/soguazu/clean_arch/services"
)

var (
	muxDispatcher router.Router = router.NewMuxRouter()

	postRepository repository.Repository      = repository.NewSQLiteRepository()
	postService    services.PostService       = services.NewPostService(postRepository)
	controller     controllers.PostController = controllers.NewPostController(postService)

	carDetailService     services.CarDetailsService       = services.NewCarDetailsService()
	carDetailsController controllers.CarDetailsController = controllers.NewCarDetailsControllers(carDetailService)
	portNumber                                            = ":8000"
)

func main() {

	muxDispatcher.GET("/", controller.Home)
	muxDispatcher.GET("/posts", controller.GetPosts)
	muxDispatcher.POST("/posts", controller.AddPost)

	muxDispatcher.GET("/carDetails", carDetailsController.GetCarDetails)

	log.Fatalln(muxDispatcher.SERVE(portNumber))
}
