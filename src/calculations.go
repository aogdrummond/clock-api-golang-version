package src

// Explanation of how to calculate
func CalculateAngleBetweenArrows(n map[string]int) int {
	hour := n["hours"]
	minute := n["minutes"]
	hourAngle := calculateHourDegreeFromZero(hour, minute)
	minuteAngle := calculateMinuteDegreeFromZero(minute)
	smallerAngle := calculateSmallerAngleBetweenArrows(hourAngle, minuteAngle)
	return smallerAngle
}

func calculateHourDegreeFromZero(hour int, minute int) int {
	degPerHour := 360.0 / 12.0
	anglePerMinute := 30.0 / 60.0
	minuteAngle := float64(minute) * anglePerMinute
	return int(hour*int(degPerHour)) + int(minuteAngle)
}

func calculateMinuteDegreeFromZero(minute int) int {

	return int(minute * int(360/60))
}

func calculateSmallerAngleBetweenArrows(arrowOnePosition int, arrowTwoPosition int) int {

	angleOne := abs(arrowOnePosition - arrowTwoPosition)
	if angleOne < 180 {
		return angleOne
	} else {
		return 360 - angleOne
	}

}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
