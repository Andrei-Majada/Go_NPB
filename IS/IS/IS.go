package IS

import (
	"fmt"
	"sync"
	"time" //Package time provides functionality for measuring and displaying time.
	// "runtime"
	"os"
	"math"
	"sort"
	// "math/rand"
)

const (
	r23 = (0.5*0.5*0.5*0.5*0.5*0.5*0.5*0.5*0.5*0.5*0.5*0.5*0.5*0.5*0.5*0.5*0.5*0.5*0.5*0.5*0.5*0.5*0.5)//0,5 ^23
	r46 = r23 * r23
	t23 = (2.0*2.0*2.0*2.0*2.0*2.0*2.0*2.0*2.0*2.0*2.0*2.0*2.0*2.0*2.0*2.0*2.0*2.0*2.0*2.0*2.0*2.0*2.0)//2 ^23
	t46 = t23 * t23
)


func Randlc(x *float64, a float64) float64 {
	var t1, t2, t3, t4, a1, a2, x1, x2, z float64
	var aux float64
	t1 = r23 * a
	a1 = float64(int(t1))
	a2 = a - t23*a1

	t1 = r23 * (*x)
	x1 = float64(int(t1))
	x2 = (*x) - t23*x1
	t1 = a1*x2 + a2*x1
	t2 = float64(int(r23 * t1))
	z = t1 - t23*t2
	t3 = t23*z + a2*x2
	t4 = float64(int(r46 * t3))
	aux = t3 - t46*t4
	(*x) = aux
	return (r46 * (*x))
}

//definição de constantes de controle
const (
	T_BENCHMARKING = 0
	T_INITIALIZATION = 1
	T_SORTING = 2
	T_TOTAL_EXECUTION = 3
)

//classe escolhida para execução
const (
	CLASS_S = 'S'
	CLASS_W = 'W'
	CLASS_A = 'A'
	CLASS_B = 'B'
	CLASS_C = 'C'
	CLASS_D = 'D'
)

//definição de classe
const (
	CLASS_TYPE = CLASS_S
)

//definição de variaveis de controle
var (
	TOTAL_KEYS_LOG_2 int
	MAX_KEY_LOG_2 int
	NUM_BUCKETS_LOG_2 int
)

var (
	TOTAL_KEYS int
	MAX_KEY = 1000000
	NUM_BUCKETS = 256
	NUM_KEYS int
	USE_BUCKETS bool = false
	SIZE_OF_BUFFERS int
)

const (
	MAX_ITERATIONS = 10
	TEST_ARRAY_SIZE = 5
)

var (
	key_array []int
	key_buff1 []int
	key_buff2 []int
	partial_verify_vals []int
)

var (
	// Partial verif info
	test_index_array [TEST_ARRAY_SIZE]int
	test_rank_array [TEST_ARRAY_SIZE]int

	S_test_index_array = [TEST_ARRAY_SIZE]int{48427, 17148, 23627, 62548, 4431}
	S_test_rank_array = [TEST_ARRAY_SIZE]int{0, 18, 346, 64917, 65463}

	W_test_index_array = [TEST_ARRAY_SIZE]int{357773, 934767, 875723, 898999, 404505}
	W_test_rank_array = [TEST_ARRAY_SIZE]int{1249, 11698, 1039987, 1043896, 1048018}

	A_test_index_array = [TEST_ARRAY_SIZE]int{2112377, 662041, 5336171, 3642833, 4250760}
	A_test_rank_array = [TEST_ARRAY_SIZE]int{104, 17523, 123928, 8288932, 8388264}

	B_test_index_array = [TEST_ARRAY_SIZE]int{41869, 812306, 5102857, 18232239, 26860214}
	B_test_rank_array = [TEST_ARRAY_SIZE]int{33422937, 10244, 59149, 33135281, 99}

	C_test_index_array = [TEST_ARRAY_SIZE]int{44172927, 72999161, 74326391, 129606274, 21736814}
	C_test_rank_array = [TEST_ARRAY_SIZE]int{61147, 882988, 266290, 133997595, 133525895}

	D_test_index_array = [TEST_ARRAY_SIZE]int{1317351170, 995930646, 1157283250, 1503301535, 1453734525}
	D_test_rank_array = [TEST_ARRAY_SIZE]int{1, 36538729, 1978098519, 2145192618, 2147425337}
)

var (
	key_buff_ptr_global []int// usado para verificar os valores passados
	passed_verification int///flag de verificação
)

type Bucket struct{
	size [][]int
	ptrs [][]int
}

func IS(M int){

	if M == 24 {//S
		TOTAL_KEYS_LOG_2 = 16
		MAX_KEY_LOG_2 = 11
		NUM_BUCKETS_LOG_2 = 9
	} else if M == 25 {//W
		TOTAL_KEYS_LOG_2 = 20
		MAX_KEY_LOG_2 = 16
		NUM_BUCKETS_LOG_2 = 10
	} else if M == 28 {//A
		TOTAL_KEYS_LOG_2 = 23
		MAX_KEY_LOG_2 = 19
		NUM_BUCKETS_LOG_2 = 10
	} else if M == 30 {//B
		TOTAL_KEYS_LOG_2 = 25
		MAX_KEY_LOG_2 = 21
		NUM_BUCKETS_LOG_2 = 10
	} else if M == 32 {//C
		TOTAL_KEYS_LOG_2 = 27
		MAX_KEY_LOG_2 = 23
		NUM_BUCKETS_LOG_2 = 10
	} else if M == 36 {//D
		TOTAL_KEYS_LOG_2 = 31
		MAX_KEY_LOG_2 = 27
		NUM_BUCKETS_LOG_2 = 10
	} else {
		fmt.Println("input error: type 'make help' for more info")
		os.Exit(1)
	}

	TOTAL_KEYS = 1 << TOTAL_KEYS_LOG_2//1 * 2 ^(TOTAL_KEYS_LOG_2)
	MAX_KEY = 1 << MAX_KEY_LOG_2//1 * 2 ^(MAX_KEY_LOG_2)
	NUM_KEYS = TOTAL_KEYS
	SIZE_OF_BUFFERS = NUM_KEYS

	key_array = make([]int, SIZE_OF_BUFFERS)
	key_buff1 = make([]int, MAX_KEY)
	key_buff2 = make([]int, SIZE_OF_BUFFERS)
	partial_verify_vals = make([]int, TEST_ARRAY_SIZE)

	var num_workers = 8
	var rr Bucket

	rr.ptrs = make([][]int, 0, num_workers)
	rr.size = make([][]int, 0, num_workers)

	var wg sync.WaitGroup

	var start = time.Now()

	wg.Add(num_workers)
	for i := 0; i < num_workers; i++ {
		go create_seq(314159265.00, 1220703125.00, i, num_workers, &wg)
	}
	wg.Wait()

	var stop = time.Now()
	var t = stop.Sub(start)

	fmt.Println("time:", t)

	sort.Ints(key_array)

	// bucketSort(key_array, 2)
	passed_verification = 0

	full_verify(0)

	if passed_verification > 0 {
		fmt.Println("Verification = SUCCESSFUL")
	} else {
		fmt.Println("Verification = UNSUCCESSFUL")
	}
	fmt.Printf("\n\n")
	fmt.Println("**************************************************************************")
	fmt.Println("*                           BENCHMARK COMPLETED                          *")
	fmt.Println("**************************************************************************")
	fmt.Printf("\n\n")
}

func create_seq(seed float64, a float64, id int, numWorkers int, wg *sync.WaitGroup) {
	(*wg).Done()
	// var wg sync.WaitGroup 
	var (
		x, s float64
		i, j, k int
		an = a // a = 5.00
		mq int
		k1, k2 int
	)
	//id vai de 0 a 8
	
	mq = (NUM_KEYS + numWorkers - 1) / numWorkers

	k1 = mq * id
	k2 = k1 + mq
	if k2 > NUM_KEYS {
			k2 = NUM_KEYS
	}

	s = find_my_seed(id,
		numWorkers,
		int64(4*NUM_KEYS),
		seed,
		an)

	k = MAX_KEY / 4
	// wg.Add(k2)
	
	for i = k1; i < k2; i++ {
		// go func(k1 int){

		x = Randlc(&s, an)

		x += Randlc(&s, an)

		x += Randlc(&s, an)
		fmt.Println("Resultado rand", x)

		x += Randlc(&s, an)
		key_array[j] = int(float64(k) * x)

		key_array[i] = int(x) * k
	}
}

func find_my_seed(kn, np int, nn int64, s, a float64) float64 {
	/*
	 * Create a random number sequence of total length nn residing
	 * on np number of processors.  Each processor will therefore have a
	 * subsequence of length nn/np.  This routine returns that random
	 * number which is the first random number for the subsequence belonging
	 * to processor rank kn, and which is used as seed for proc kn ran # gen.
	 */
	var t1, t2 float64
	var mq, nq, kk, ik int64

	if kn == 0 {
		return s
	}

	mq = int64(math.Ceil(float64(nn/4+int64(np)-1) / float64(np)))
	nq = mq * 4 * int64(kn)

	t1 = s
	t2 = a
	kk = nq
	for kk > 1 {
		ik = kk / 2
		if 2*ik == kk {
			t2 = Randlc(&t2, t2)
			kk = ik
		} else {
			t1 = Randlc(&t1, t2)
			kk = kk - 1
		}
	}

	t1 = Randlc(&t1, t2)
	return t1
}

func bucketSort(array []int, bucketSize int) []int {
	var max, min int
	for _, n := range array {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	nBuckets := (max-min)/bucketSize + 1
	buckets := make([][]int, nBuckets)
	for i := 0; i < nBuckets; i++ {
		buckets[i] = make([]int, 0)
	}

	for _, n := range array {
		idx := (n-min) / bucketSize
		buckets[idx] = append(buckets[idx], n)
	}

	sorted := make([]int, 0)
	for _, bucket := range buckets {
		if len(bucket) > 0 {
			sorted = append(sorted, bucket...)
		}
	}

	return sorted
}

func full_verify(numWorkers int) {
	var (
		j int
		// k, k1 int
	)

	// var rr Bucket

	// var wg sync.WaitGroup

	// wg.Add(numWorkers)

	// for i = 0; i < numWorkers; i++ {
	// 	go func(i int) {
	// 		for j=i; j<NUM_BUCKETS; j+=numWorkers {
	// 			if j > 0 {
	// 				k1 = rr.ptrs[i][j-1]
	// 			} else {
	// 				k1 = 0
	// 			}
	// 			for m = k1; m < rr.ptrs[i][j]; m++ {
	// 				k = key_buff_ptr_global[key_buff2[m]] -1
	// 				key_array[k] = key_buff2[m]
	// 			}
	// 		}
	// 	}(i)
	// }
	// wg.Wait()


	j = 0
	for i := 1; i < len(key_array); i++ {
		if key_array[i-1] > key_array[i] {
			j++
		}
	}
	if j != 0 {
		fmt.Println("Full_verify: number of keys out of sort: ", j)
	} else {
		passed_verification += 1
	}
}
