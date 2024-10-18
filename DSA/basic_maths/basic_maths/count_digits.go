// Algorithm
// Step 1: Initialise a variable to store the count of digits of the number.
// Step 2: The count of digits can be calculated using log10 N + 1.


package basic_maths
import "math"
func CountDigits(num int) int {
    if num == 0 {
        return 1
    }
    return int(math.Floor(math.Log10(float64(num))) + 1)
}