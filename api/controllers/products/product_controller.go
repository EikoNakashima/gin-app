

package products
 
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
 
	"github.com/gin-gonic/gin"
	"github.com/gouser/gin-app/api/domain/products"
)
 
// CreateProduct - Create product
func CreateProduct(c *gin.Context) {
	var product products.Product
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Println(err.Error())
		return
	}
	if err := json.Unmarshal(bytes, &product); err != nil {
		log.Println(err.Error())
		return
	}
	fmt.Println(string(bytes))
	fmt.Println(err)
	c.JSON(http.StatusOK, product)
}