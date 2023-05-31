package services

const MAX_LAYERS = 10 // Hard limit to prevent hypothetical endless recursive searching

// Ensures that degrees of separation don't exceed max layers supported by this API
func getMaxLayers(degreesOfSeparation int) int {
	maxLayers := degreesOfSeparation
	if maxLayers > MAX_LAYERS {
		maxLayers = MAX_LAYERS
	}
	return maxLayers
}
