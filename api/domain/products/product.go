package products
 
import "time"
 
// Product - defines product info uploaded by user
type Product struct {
	ID        uint64 `json:"id"`
	Name      string `json:"name"`
	Detail    string `json:"detail"`
	Price     uint64 `json:"price"`
	Img       []byte `json:"img"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}