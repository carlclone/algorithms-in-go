package array_and_hashtable

//Given four lists A, B, C, D of integer values, compute how many tuples (i, j, k, l) there are such that A[i] + B[j] + C[k] + D[l] is zero.
//
//To make problem a bit easier, all A, B, C, D have same length of N where 0 ≤ N ≤ 500. All integers are in the range of -228 to 228 - 1 and the result is guaranteed to be at most 231 - 1.
//
//Example:
//
//Input:
//A = [ 1, 2]
//B = [-2,-1]
//C = [-1, 2]
//D = [ 0, 2]
//
//Output:
//2
//
//Explanation:
//The two tuples are:
//1. (0, 0, 0, 1) -> A[0] + B[0] + C[0] + D[1] = 1 + (-2) + (-1) + 2 = 0
//2. (1, 1, 0, 0) -> A[1] + B[1] + C[0] + D[0] = 2 + (-1) + (-1) + 0 = 0

func fourSumCount(A []int, B []int, C []int, D []int) int {
	/*
		iterate all A[i] B[j] C[k] D[l] combo , if equal 0 , then res++
	*/
	/*
	 put A+B combo into map , iterate all combo ,
	*/

	res := 0
	map1 := make(map[int]int)
	for _, i := range A {
		for _, j := range B {
			map1[i+j]++
		}
	}

	var tmp int
	for _, k := range C {
		for _, l := range D {
			tmp = -(k + l)
			if _, ok := map1[tmp]; ok {
				res += map1[tmp]
			}
		}
	}
	return res
}
