package fakers

import (
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/Romiz-Lab/BE-go-ecommerce/app/models"
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

func ProductFaker(db *gorm.DB) *models.Product {
	user := UserFaker(db)
	err := db.Create(&user).Error
	if err != nil {
    log.Fatal(err)
  }

	name := faker.Name()
	return &models.Product{
		ID: 				uuid.New().String(),
		UserID: 		user.ID,
		Sku: 				slug.Make(name),
		Name: 			name,
		Slug: 			slug.Make(name),
		Price: 			decimal.NewFromFloat(FakePrice()),
		Stock: 			rand.Intn(100),
		Weight: 		decimal.NewFromFloat(rand.Float64()),
		ShortDescription: faker.Paragraph(),
		Description: faker.Paragraph(),
		Status: 		1,
		CreatedAt: 	time.Time{},
		UpdatedAt: 	time.Time{},
		DeletedAt:   gorm.DeletedAt{},
	}
}

func precision(f float64, places int) float64 {
  shift := math.Pow10(places)
	return float64(int64(f * shift)) / shift
}

func FakePrice() float64 {
	return precision(rand.Float64() *math.Pow10(rand.Intn(8)), rand.Intn(2) + 1)
}

