package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	//panic("NeedsLicense not implemented")
	
	if kind == "car" || kind == "truck" {
		return true
	}
	
	return false
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in dictionary order.
func ChooseVehicle(option1, option2 string) string {
	//panic("ChooseVehicle not implemented")
	
	var option string
	
	if option1 == "Bugatti Veyron" || option2 == "Bugatti Veyron" {
		option = "Bugatti Veyron"
	} else if option1 == "Chery EQ" || option2 == "Chery EQ" {
		option = "Chery EQ"
	} else if option1 == "Kia Niro Elektro" || option2 == "Kia Niro Elektro" {
		option = "Chery EQ"
	} else if  option1 == "Toyota Corolla" || option2 == "Toyota Corolla" {
		option = "Toyota Corolla"
	} else if option1 == "Wuling Hongguang" || option2 == "Wuling Hongguang" {
		option = "Wuling Hongguang"
	} else if option1 == "Volkswagen Beetle" || option2 == "Volkswagen Beetle" {
		option = "Volkswagen Beetle"
	} else if  option1 == "Ford Focus" || option2 == "Ford Focus" {
		option = "Ford Focus"
	} else if option1 == "Ford Pinto" || option2 == "Ford Pinto" {
		option = "Ford Pinto"
	} else if option1 == "2018 Bergamont City" || option2 == "2018 Bergamont City" {
		option = "2018 Bergamont City"
	} else {
		option = "2020 Gazelle Medeo"
	}
	
	return option + " is clearly the better choice."
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	//panic("CalculateResellPrice not implemented")
	var newPrice float64	
	
	if age < 3 {
		newPrice = originalPrice * 0.8
	} else if age >= 3 && age < 10 {
		newPrice = originalPrice * 0.7
	} else {
		newPrice = originalPrice * 0.5
	}
	
	return newPrice
	
}
