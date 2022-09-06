package dummy

import (
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"github.com/martinustrinc/Shopping/Model"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

func ProductDummy(db *gorm.DB) *Model.Product {
	user := UserDummy(db)
	err := db.Create(&user).Error
	if err != nil {
		log.Fatal(err)
	}

	name := dummy.Name()
	return &Model.Product{
		ID:               uuid.New().String(),
		UserID:           user.ID,
		Sku:              slug.Make(name),
		Name:             name,
		Price:            decimal.NewFromFloat(dummyPrice()),
		Stock:            rand.Intn(100),
		Weight:           decimal.NewFromFloat(rand.Float64()),
		ShortDescription: dummy.Paragraph(),
		Description:      dummy.Paragraph(),
		Status:           1,
		CreatedAt:        time.Time{},
		UpdatedAt:        time.Time{},
		DeletedAt:        gorm.DeletedAt{},
	}
}

func dummyPrice() float64 {
	return precision(rand.Float64()*math.Pow10(rand.Intn(8)), rand.Intn(2)+1)
}

// precision | a helper function to set precision of price
func precision(val float64, pre int) float64 {
	div := math.Pow10(pre)
	return float64(int64(val*div)) / div
}