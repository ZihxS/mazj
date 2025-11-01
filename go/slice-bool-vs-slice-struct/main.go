package main

import (
	"fmt"
	"runtime"
)

const N = 10_000_000

func getAlloc() uint64 {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return m.Alloc
}

func main() {
	runtime.GC()

	allocBefore := getAlloc()
	s1 := make([]bool, N)
	boolUsage := getAlloc() - allocBefore

	runtime.GC()

	allocBefore = getAlloc()
	s2 := make([]struct{}, N)
	structUsage := getAlloc() - allocBefore

	fmt.Print("\n\n\n\n\n")

	fmt.Printf("🌟 JUMLAH ELEMEN -> %d\n\n", N)

	fmt.Printf("👉 []bool -> %d bytes (kurang lebih %.2f MB)\n", boolUsage, float64(boolUsage)/1e6)
	fmt.Printf("👉 []struct{} -> %d bytes (kurang lebih %.2f MB)\n\n", structUsage, float64(structUsage)/1e6)

	if boolUsage > 0 {
		saved := float64(boolUsage-structUsage) / 1e6
		percent := (1 - float64(structUsage)/float64(boolUsage)) * 100
		fmt.Printf("✅ HEMAT -> kurang lebih %.2f MB (kurang lebih %.2f%%)\n", saved, percent)
	}

	fmt.Print("\n\n\n\n\n")

	_ = s1
	_ = s2
}
