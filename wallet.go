package wallet
var ConRate=82.47

func CheckInrToUsdConRate( UserGuessConRate float64) bool{
    if(UserGuessConRate == ConRate){
		return true
	} else {
		return false
	}
}
