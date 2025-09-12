package hamming

import (
    "errors"
)

func Distance(a, b string) (int, error) {
	ar := []rune(a)
    br := []rune(b)
    hammingDistance := 0
	length := len(ar)

	if len(ar) != len(br) {
		return 0, errors.New("strings must be of equal length")
	}
    
    for i := 0; i < length; i++{
    	if ar[i] != br[i] {
            hammingDistance++
        }
    }
    return hammingDistance, nil
}