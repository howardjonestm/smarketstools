package smarketstools

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type BidPrice struct {
	gorm.Model
	Time        int
	ContractID  string
	BidPrice    int
	BidQuantity int
}

type OfferPrice struct {
	gorm.Model
	Time          int
	ContractID    string
	OfferPrice    int
	OfferQuantity int
}

func InsertQuotesMysql(marketID, connection string) {

	token := ReadToken()
	sclient := Client{ApiToken: token}
	quoteCollection, _ := GetQuoteCollection(marketID, sclient)

	db, err := gorm.Open("mysql", connection)
	checkErr(err)
	defer db.Close()

	db.AutoMigrate(&BidPrice{})
	db.AutoMigrate(&OfferPrice{})

	for _, k := range quoteCollection {
		t := int(time.Now().Unix())
		for _, v := range k.Bids {
			db.Create(&BidPrice{Time: t, ContractID: k.ContractID, BidPrice: v.Price, BidQuantity: v.Quantity})
		}
		for _, v := range k.Offers {
			db.Create(&OfferPrice{Time: t, ContractID: k.ContractID, OfferPrice: v.Price, OfferQuantity: v.Quantity})
		}
	}

}
