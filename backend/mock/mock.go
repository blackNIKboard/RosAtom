package mock

import (
	"errors"
	"math/rand"
	"sort"
)

/* WeightedProbability is the class to handle all of the weights etc
 * There will be a values and weights slice, think ArrayList in Java
 * Total will be the total value of weights
 * ReplayValues is a slice containing the replay sequence
 */
type WeightedProbability struct {
	ValueWeights []vw    //assume float 64 because with such a large input size you would need decimal percentages, ie .25%
	Total        float64 // total value of the weights
	ReplayValues []float64
	Initialized  bool
}

/* vw is a struct used to map the Key Value stores of the map to a slice
 */
type vw struct {
	Value  float64
	Weight float64
}

/* sortByWeight sorts a map by weight and return a  slice of struct vw and the total
 * Can then perform binary search on a slice of structs, can't perform binary search on map
 * @m float[64]float64
 */
func sortByWeight(m map[float64]float64) ([]vw, float64) {

	var mkv []vw
	total := 0.0
	for value, weight := range m {
		total += weight
		mkv = append(mkv, vw{value, total})
	}

	sort.Slice(mkv, func(i, j int) bool {
		return mkv[i].Weight < mkv[j].Weight
	})
	return mkv, total

}

/* Init
 * @numPool float[64]float64
 * numPool of tbe form [value] => [weight]
 * Ex: 7 => .25 so 7 will appear with 25% frequency
 * Total Probabilty can be over 100%
 * Algorithm takes O(N) to create the weights and values
 * Since using a Map there should be no duplicates except ones of form 7 vs 7.00
 */
func (wp *WeightedProbability) Init(numPool map[float64]float64) error {
	if numPool == nil {
		return errors.New("Number Pool is not initialized!")
	}
	valueWeights, total := sortByWeight(numPool)
	if total > 100.00 {
		return errors.New("Total is greater than 100!")
	}
	replayValues := []float64{}
	wp.ValueWeights = valueWeights
	wp.Total = total
	wp.ReplayValues = replayValues
	wp.Initialized = true
	return nil
}

/* GenerateRandomNumber
 * Returns an error if the struct is not initialized or if struct is unable to generate a random number
 * Sort.Search uses binary search to find the index of the first weight >= x
 */
func (wp *WeightedProbability) GenerateRandomNumber() (float64, error) {
	if !wp.Initialized {
		return 0, errors.New("Not initialized")
	}
	x := rand.Float64() * wp.Total
	// search the distribution, essentially the same as pythons bisect_left
	i := sort.Search(len(wp.ValueWeights), func(i int) bool {
		return wp.ValueWeights[i].Weight >= x
	})
	if i >= len(wp.ValueWeights) {
		return 0, errors.New("Index to big")
	}
	wp.ReplayValues = append(wp.ReplayValues, wp.ValueWeights[i].Value)
	return wp.ValueWeights[i].Value, nil
}

/* Replay
 * Since we were told we shouldn't write to a file, just store and return the slice
 */
func (wp *WeightedProbability) Replay() ([]float64, error) {
	if !wp.Initialized {
		return nil, errors.New("Not initialized")
	}
	return wp.ReplayValues, nil
}
