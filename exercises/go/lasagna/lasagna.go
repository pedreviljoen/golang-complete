package lasagna

const (
	OVENTIME  int = 40
	LAYERTIME int = 2
)

func OvenTime() int {
	return OVENTIME
}

func RemainingOvenTime(time int) int {
	return OVENTIME - time
}

func PreparationTime(layers int) int {
	return LAYERTIME * layers
}

func ElapsedTime(layers int, timeInOven int) int {
	return PreparationTime(layers) + timeInOven
}
