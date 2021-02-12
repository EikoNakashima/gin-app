package app
 
import "github.com/gouser/gin-app/api/controllers/products"
 
func mapUrls() {
	router.GET("/products/:product_id", products.GetProduct)
	router.POST("/products", products.CreateProduct)
}