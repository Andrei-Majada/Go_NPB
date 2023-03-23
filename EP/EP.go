package EP

import (
	"fmt"
	"log"
	"sync"
	"os"
	"math"
	"time"
	"runtime"
	"strconv"
)

const(
	r23 = (0.5*0.5*0.5*0.5*0.5*0.5*0.5*0.5*0.5*0.5*0.5*0.5*0.5*0.5*0.5*0.5*0.5*0.5*0.5*0.5*0.5*0.5*0.5)//0,5 ^23
	r46 = r23 * r23
	t23 = (2.0*2.0*2.0*2.0*2.0*2.0*2.0*2.0*2.0*2.0*2.0*2.0*2.0*2.0*2.0*2.0*2.0*2.0*2.0*2.0*2.0*2.0*2.0)//2 ^23
	t46 = t23 * t23
)

const MK = 16
const NK = 1 << MK
const NQ = 10
const EPSILON = 1.0e-8
const A = 1220703125.0
const S = 271828183.0
const NK_PLUS = ((2*NK)+1)

type Results struct{
	//float64 is a version of float that stores decimal values using a total of 64-bits of data.
	qqR [NQ]float64
	sxx float64
	syy float64
}

func Ep(M int){
	var MM = M - MK 
	var NN = 1 << MM //1 * 2 ^(MM)
	
	var x = [NK_PLUS]float64{}
	var q = [NQ]float64{}
	var Mops, sx, sy ,an, gc, t1 float64
	var np int
	var checkSX float64
	var checkSY float64	
	dum := [3]float64{1.0,1.0,1.0}
	var t time.Duration
	var verify bool
	var sx_err float64
	var sy_err float64

	
	verify = false
	
	np = NN
	
	rand(0, &dum[0], dum[1], []float64{dum[2]})
	dum[0] = randlc(&dum[1], dum[2])
	for i := 0; i < NK_PLUS; i++{
		x[i] = -1.0e99
	}
	Mops = math.Log(math.Sqrt(math.Abs(math.Max(1.0,1.0))))

	var aux string
	if runtime.GOOS == "windows" {
		aux = "..\\bin\\"
	} else {
		aux = "../bin/"
	}

	t1 = A
	rand(0,&t1,A,x[:])
	
	for i := 0; i < MK+1; i++{
		randlc(&t1,t1)
	}
	
	an = t1
	t = S
	gc = 0.0
	sx = 0.0
	sy = 0.0
	
	for i := 0; i <= NQ-1; i++{
		q[i] = 0.0
	}
	
	var rr Results

	// Channels sends and receives block until the other side is ready. 
	// This allows goroutines to synchronize once all goroutines have completed their computation.
	// Buffered channels accept a limited number of values without a corresponding receiver for those values.
	result := make(chan Results,np) 

	//To wait for multiple goroutines to finish, we can use a wait group. 
	var wg sync.WaitGroup 

	start := time.Now()

	//Add np number to the WaitGroup counter. If the counter becomes zero, all goroutines blocked on Wait are released.
	wg.Add(np)

	for k := 1; k <= np; k++{
		//Goroutine is like a thread where you can spawn a new goroutine to run code simultaneously using the go keyword:
		//Our goroutine is independent and unknown to the main goroutine. Therefore, the main thread terminates even before executing our goroutine.
		//we need to ask the main thread to wait for the goroutine
		go func(k int){
			//
			defer wg.Done()
			var SX, SY float64
			var t1,t2,t3,t4,x1,x2 float64
			var kk, ik, l int
			var qq = [NQ]float64{}
			var x = [NK_PLUS]float64{}
			var rrTemp Results
			kk = (-1) + k
			t1 = S
			t2 = an
			
			for i:=0;i<NQ-1;i++{
				qq[i] = 0.0
			}
			SX = 0.0
			SY = 0.0
			
			for i:=0; i <=100; i++{
				ik = kk/2
				if ((2*ik) != kk){
					t3 = randlc(&t1,t2)
				}
				if (ik == 0){
					break
				}
				t3 = randlc(&t2,t2)
				kk = ik
			}
			rand((2*NK), &t1, A, x[:])
			
			for i := 0; i< NK; i++{
				x1 = 2.0 * x[2*i] - 1.0
				x2 = 2.0 * x[2*i+1] - 1.0
				t1 = math.Pow(x1,2) + math.Pow(x2,2)
				if (t1 <= 1.0){
					t2 = math.Sqrt(-2.0 * math.Log(t1) / t1)
					t3 = (x1 * t2)
					t4 = (x2 * t2)
					l = int(math.Max(math.Abs(t3), math.Abs(t4)))
					qq[l] += 1.0
					SX += t3
					SY += t4
				}
			}
			rrTemp.qqR = qq
			rrTemp.syy = SY
			rrTemp.sxx = SX
			result <- rrTemp
		}(k)
	} 
	for i := 1; i<= np; i++{
		rr = <-result
		sx += rr.sxx
		sy += rr.syy
		for j := range q{
			q[j] += rr.qqR[j]
		}
	}

	stop := time.Now()
	t = stop.Sub(start)
	//A sender can close a channel to indicate that no more values will be sent. 
	close(result)
	wg.Wait()
	
	for i := 0; i < NQ-1; i ++{
		gc = gc + q[i]
	}
	
	verify = true
	var n string
	if M == 24 {
		checkSX = -3.247834652034740e+3
		checkSY = -6.958407078382297e+3
		n = "S"
	}else if  M == 25 {
		checkSX = -2.863319731645753e+3
		checkSY = -6.320053679109499e+3
		n = "W"
	}else if M == 28 {
		checkSX = -4.295875165629892e+3
		checkSY = -1.580732573678431e+4
		n = "A"
	}else if M == 30 {
		checkSX =  4.033815542441498e+4
		checkSY = -2.660669192809235e+4
		n = "B"
	}else if M == 32 {
		checkSX =  4.764367927995374e+4
		checkSY = -8.084072988043731e+4
		n = "C"
	}else if M == 36 {
		checkSX =  1.982481200946593e+5
		checkSY = -1.020596636361769e+5
		n = "D"
	}else if M == 40 {
		checkSX = -5.319717441530e+05
		checkSY = -3.688834557731e+05
		n = "E"
	}else {
		verify = false
	}

	if verify {
		sx_err = math.Abs((sx - checkSX) / checkSX)
		sy_err = math.Abs((sy - checkSY) / checkSY)
		verify = ((sx_err <= EPSILON) && (sy_err <= EPSILON))
	}
	
	Mops = math.Pow(2.0, float64(M+1))/(t.Seconds())/1000000.0	
		
	fmt.Println("\n\nEP Benchmark Results available on bin/EP_" + n + "!\n")

	if verify {
		fmt.Println("Verification = SUCCESSFUL")
	} else {
		fmt.Println("Verification = UNSUCCESSFUL")
	}
	fmt.Printf("\n\n")
	fmt.Println("**************************************************************************")
	fmt.Println("*                           BENCHMARK COMPLETED                          *")
	fmt.Println("**************************************************************************")
	fmt.Printf("\n\n")

	f, err := os.Create(aux + "EP_" + n + ".txt")
	if err != nil {
		log.Fatal(err)
	}

	f.WriteString("NAS Parallel Benchmark Parallel GO version\n")
	f.WriteString("Number of random numbers generated: " +  fmt.Sprint(math.Pow(2.0,float64(M+1))) + "\n\n")
	f.WriteString("Benchmark EP Results:" + "\n")
	f.WriteString("CPU Time = " + fmt.Sprint(t) + "\n")
	f.WriteString("N = " + strconv.Itoa(M) + "\n")
	f.WriteString("No. Gaussian Pairs = " + fmt.Sprint(gc) + "\n")
	f.WriteString("Sums = " + fmt.Sprint(sx) + "  " + fmt.Sprint(sy) + "\n")
	f.WriteString("Counts: \n")
	for i := 0; i < NQ-1; i++ {
		f.WriteString(strconv.Itoa(i) + " - " + fmt.Sprint(q[i]) + "\n")
	}
	f.WriteString("Class NPB = " + n + "\n")
	f.WriteString("Total threads = " + strconv.Itoa(runtime.NumCPU()) + "\n")
	f.WriteString("Mop/s total = " + fmt.Sprint(Mops) + "\n")
	f.WriteString("Operation type = EP - Generation of Random Numbers\n")

	if verify {
		f.WriteString("Verification = SUCCESSFUL\n")
	} else {
		f.WriteString("Verification = UNSUCCESSFUL\n")
	}

	f.WriteString("Compiler Version = " + fmt.Sprint(runtime.Version()))

	defer f.Close()
}

func randlc(x *float64, a float64) float64 {
	
	var t1,t2,t3,t4,a1,a2,x1,x2,z float64

	t1 = r23*a
	a1 = float64(int(t1))
	a2 = a - t23 * a1

	t1 = r23 * (*x)
	x1 = float64(int(t1))
	x2 = (*x) - t23 * x1
	t1 = a1 * x2 + a2 * x1
	t2 = float64(int(r23 * t1))
	z = t1 - t23 * t2
	t3 = t23 * z + a2 * x2
	t4 = float64(int(r46 * t3))
	(*x) = t3 - t46 * t4

	return (r46 * (*x))
}

func rand(n int, x_speed *float64, a float64, y []float64){

	var x,t1,t2,t3,t4,a1,a2,x1,x2,z float64

	t1 = r23 * a
	a1 = float64(int(t1))
	a2 = a - t23 * a1
	x  = *x_speed

	for i:=0; i < n; i++{
		t1 = r23 * x
		x1 = float64(int(t1))
		x2 = x - t23 * x1
		t1 = a1 * x2 + a2 * x1
		t2 = float64(int(r23 * t1))
		z = t1 - t23 * t2
		t3 = t23 * z + a2 * x2
		t4 = float64(int(r46 * t3))
		x = t3 - t46 * t4
		y[i] = r46 * x
	}

	*x_speed = x
}

