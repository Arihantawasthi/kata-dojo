package algos

import "math"

func CrystalProblem(storeys []bool) int {
    height := len(storeys)
	stepSize := int(math.Sqrt(float64(height)))
    var broken_storey int
    var prev_step int
    var min_storey int

    for i := 0; i < height; i += stepSize {
        if storeys[i] {
            broken_storey = i
            prev_step = i - stepSize
            if prev_step < 0 {
                prev_step = 0
            }
            break;
        }
    }

    for i := prev_step; i <= broken_storey; i++ {
        if storeys[i] {
            min_storey = i
            break
        }
    }

    return min_storey;
}
