package smarketstools

//Only tested for back bets
//Converts conventional odds/stakes into exchange quanity/prices
func ConvertToExchange(pounds, odds float64) (int, int) {
	priceUnrounded := 10000 / odds
	price := int(priceUnrounded)

	quantity := int((100000000 * pounds) / priceUnrounded)

	return price, quantity
}

//Only tested for back bets
//Converts  exchange quanity/prices into conventional odds/stakes
func ConvertFromExchange(price, quantity int) (float64, float64) {
	priceF64 := float64(price)
	odds := 10000 / priceF64

	quantityF64 := float64(quantity)
	pounds := (quantityF64 * priceF64) / 100000000

	return pounds, odds
}
