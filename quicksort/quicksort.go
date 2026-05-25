package quicksort

import (
	"errors"
	"math"
	"math/rand"
)

type (
	Slice         []float64
	PivotPosition int
)

const (
	PivotFirst PivotPosition = iota
	PivotLast
	PivotMiddle
	PivotRandom
)

var (
	ArrayOutOfBounds     error = errors.New("Array lengh out of bounds!")
	PivotOutOfBounds     error = errors.New("Pivot position out of bounds!")
	PartitionOutOfBounds error = errors.New("Partition imposible!")
)

func GenerateSlice() Slice {
	var s Slice
	for i := 0; i < int(math.Pow(2, 16)); i++ {
		s = append(s, float64(rand.Intn(int(math.Pow(2, 32)))))
	}

	return s
}

func (s Slice) QuickSort(pivotPos PivotPosition) error {

	if len(s) == 1 {
		return nil
	}

	if len(s) == 0 {
		return ArrayOutOfBounds
	}

	err := quickSort(s, 0, len(s)-1, pivotPos)
	if err != nil {
		return err
	}

	return nil
}

func quickSort(slice Slice, low, high int, pivotPos PivotPosition) error {
	if low < high {
		pivot, err := partition(slice, low, high, pivotPos)
		if err != nil {
			return err
		}

		quickSort(slice, low, pivot-1, pivotPos)
		quickSort(slice, pivot+1, high, pivotPos)

		return nil
	}

	return PartitionOutOfBounds
}

func partition(slice Slice, low, high int, pivotPos PivotPosition) (int, error) {
	pivotIndex, err := choosePivotIndex(low, high, pivotPos)

	if err != nil {
		return -1, err
	}

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

func choosePivotIndex(low, high int, pivotPos PivotPosition) (int, error) {
	switch pivotPos {

	case PivotFirst:
		return low, nil

	case PivotMiddle:
		return int((low + high) / 2), nil

	case PivotLast:
		return high, nil

	case PivotRandom:
		return rand.Intn(high-low) + low, nil

	default:
		return -1, PivotOutOfBounds
	}
}

func swap(var1, var2 *float64) {
	temp := *var1
	*var1 = *var2
	*var2 = temp
}
