package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
    if kind == "car" || kind == "truck" {
        return true
    } else {
        return false
    }
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
    result := ""
	if option1 < option2 {
        result += option1 + " is clearly the better choice."       
    } else {
        result += option2 + " is clearly the better choice."
    }
    return result
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
    var value float64;;
	if age < 3 {
        value = originalPrice * 0.80
    } else if age >= 3 && age < 10 {
        value = originalPrice * 0.70
    } else {
        value = originalPrice * 0.50
    }
    return value
}
