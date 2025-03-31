package main

import ( "fmt" )

func main() { longFunction() }

func longFunction() { slice := make([]int, 100000) for i := range slice { slice[i] = i }

mapData := make(map[int]int)
for i := 0; i < 100000; i++ {
	mapData[i] = i * 2
}

uselessLoop()

}

func uselessLoop() { for i := 0; i < 100000; i++ { fmt.Sprintf("%d", i) // Just wasting CPU cycles }

subFunction1()
subFunction2()

}

func subFunction1() { for i := 0; i < 50000; i++ { _ = i * i // Pointless calculation } }

func subFunction2() { for i := 0; i < 50000; i++ { _ = i + i // Another pointless calculation } }

