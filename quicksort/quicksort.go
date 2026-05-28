package quicksort

import (
	"errors"
	"math"
	"math/rand"
	"sort"
)

type (
	Slice []float64

	PivotPosition int
	SliceOrder    int
)

const (
	size        = 1 << 16
	maxValue    = 1 << 32
	blockNumber = 16
)

var (
	ArrayOutOfBounds     error = errors.New("Array lengh out of bounds!")
	PivotOutOfBounds     error = errors.New("Pivot position out of bounds!")
	PartitionOutOfBounds error = errors.New("Partition imposible!")
	InvalidFlag          error = errors.New("No such flag or flag value exists!")

	Nan int = int(math.NaN())
)

func GenerateSlice(sliceFlag string) (Slice, error) {

	var (
		s    Slice
		step float64 = float64(maxValue) / float64(size)
	)

	switch sliceFlag {

	case "random":
		for i := 0; i < int(math.Pow(2, 16)); i++ {
			s = append(s, float64(rand.Intn(int(math.Pow(2, 32)))))
		}

	case "inverse":
		for i := range s {
			s[i] = float64(size-1-i) * step
		}

	case "nearly":

		for i := range s {
			s[i] = float64(i) * step
		}

		perturbations := size / 50
		for i := 0; i < perturbations; i++ {
			idx := rand.Intn(size)
			offset := rand.Intn(8) + 1
			if idx+offset < size {
				s[idx], s[idx+offset] = s[idx+offset], s[idx]
			}
		}

	case "block":

		blockSize := size / blockNumber
		for i := range s {
			s[i] = float64(i) * step
		}

		// Sort within each block (already sorted here, but works for any data)
		for b := 0; b < blockNumber; b++ {
			lo := b * blockSize
			hi := lo + blockSize
			sort.Slice(s[lo:hi], func(i, j int) bool { return s[lo+i] < s[lo+j] })
		}

		// Shuffle the block order so the overall slice is globally unsorted
		blockOrder := rand.Perm(blockNumber)
		tmp := make(Slice, size)
		for newPos, oldBlock := range blockOrder {
			lo := oldBlock * blockSize
			copy(tmp[newPos*blockSize:], s[lo:lo+blockSize])
		}

	default:
		return nil, InvalidFlag
	}

	return s, nil
}

func Pivot(low, high int, pivotFlag string) (int, error) {

	switch pivotFlag {

	case "first":
		return low, nil

	case "last":
		return high, nil

	case "middle":
		return (low + high) / 2, nil

	case "random":
		return rand.Intn(high-low) + low, nil

	default:
		return Nan, InvalidFlag
	}
}

func (s Slice) QuickSort(pivotFlag string) error {

	if len(s) == 1 {
		return nil
	}

	if len(s) == 0 {
		return ArrayOutOfBounds
	}

	pivotIndex, err := Pivot(0, len(s), pivotFlag)

	if err != nil {
		return err
	}

	err = quickSort(s, 0, len(s)-1, pivotIndex)
	if err != nil {
		return err
	}

	return nil
}

func quickSort(slice Slice, low, high int, pivotIndex int) error {
	if low < high {
		pivot, err := partition(slice, low, high, pivotIndex)
		if err != nil {
			return err
		}

		quickSort(slice, low, pivot-1, pivotIndex)
		quickSort(slice, pivot+1, high, pivotIndex)

		return nil
	}

	return PartitionOutOfBounds
}

func partition(slice Slice, low, high int, pivotIndex int) (int, error) {

	swap(&slice[pivotIndex], &slice[high])

	pivot := slice[high]
	i := low

	for j := low; j < high; j++ {
		if slice[j] < pivot {
			swap(&slice[i], &slice[j])
			i++
		}
	}

	swap(&slice[i], &slice[high])
	return i, nil
}

func swap(var1, var2 *float64) {
	temp := *var1
	*var1 = *var2
	*var2 = temp
}

// func Flags(pivotFlag, sliceFlag *string) (PivotPosition, SliceOrder) {

// 	var (
// 		pivot PivotPosition
// 		slice SliceOrder
// 	)

// 	switch *pivotFlag {

// 	case "first":
// 		pivot = PivotFirst
// 	case "last":
// 		pivot = PivotLast
// 	case "middle":
// 		pivot = PivotMiddle
// 	case "random":
// 		pivot = PivotRandom
// 	default:
// 		pivot = PivotError
// 	}

// 	switch *sliceFlag {

// 	case "random":
// 		slice = RandomOrder
// 	case "inverse":
// 		slice = InverseOrder
// 	case "nearly":
// 		slice = NearlySortedOrder
// 	case "block":
// 		slice = BlockOrder
// 	default:
// 		slice = SliceError
// 	}

// 	return pivot, slice
// }

// func choosePivotIndex(low, high int, pivotFlag *string) (int, error) {
// 	switch *pivotFlag {

// 	case PivotFirst:
// 		return low, nil

// 	case PivotMiddle:
// 		return int((low + high) / 2), nil

// 	case PivotLast:
// 		return high, nil

// 	case PivotRandom:
// 		return rand.Intn(high-low) + low, nil

// 	default:
// 		return -1, PivotOutOfBounds
// 	}
// }
