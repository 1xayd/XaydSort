package main

import (
	"math/rand"
	"sort"
	"testing"
)

func QuickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	quickSortHelper(arr, 0, len(arr)-1)
}

func quickSortHelper(arr []int, low, high int) {
	if low < high {
		pi := partition(arr, low, high)
		quickSortHelper(arr, low, pi-1)
		quickSortHelper(arr, pi+1, high)
	}
}

func partition(arr []int, low, high int) int {
	mid := low + (high-low)>>1
	if arr[mid] < arr[low] {
		arr[low], arr[mid] = arr[mid], arr[low]
	}
	if arr[high] < arr[low] {
		arr[low], arr[high] = arr[high], arr[low]
	}
	if arr[mid] < arr[high] {
		arr[mid], arr[high] = arr[high], arr[mid]
	}

	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func isSorted(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			return false
		}
	}
	return true
}

func generateRandomArray(size int, seed int64) []int {
	r := rand.New(rand.NewSource(seed))
	arr := make([]int, size)
	for i := range arr {
		arr[i] = r.Intn(size * 10)
	}
	return arr
}

func generateNearlySorted(size int, seed int64) []int {
	r := rand.New(rand.NewSource(seed))
	arr := make([]int, size)
	for i := range arr {
		arr[i] = i
	}

	swaps := size / 20
	for i := 0; i < swaps; i++ {
		a, b := r.Intn(size), r.Intn(size)
		arr[a], arr[b] = arr[b], arr[a]
	}
	return arr
}

func TestXaydSort_Empty(t *testing.T) {
	arr := []int{}
	XaydSort(arr)
	if len(arr) != 0 {
		t.Error("Empty array failed")
	}
}

func TestXaydSort_Single(t *testing.T) {
	arr := []int{42}
	XaydSort(arr)
	if len(arr) != 1 || arr[0] != 42 {
		t.Error("Single element failed")
	}
}

func TestXaydSort_Random(t *testing.T) {
	sizes := []int{10, 50, 100, 500, 1000, 5000}
	for _, size := range sizes {
		arr := generateRandomArray(size, 12345)
		XaydSort(arr)
		if !isSorted(arr) {
			t.Errorf("Random array size %d failed", size)
		}
	}
}

func TestXaydSort_NearlySorted(t *testing.T) {
	arr := generateNearlySorted(1000, 12345)
	XaydSort(arr)
	if !isSorted(arr) {
		t.Error("Nearly sorted array failed")
	}
}

func TestXaydSort_Duplicates(t *testing.T) {
	arr := []int{5, 2, 8, 2, 9, 1, 5, 5, 2, 8}
	XaydSort(arr)
	if !isSorted(arr) {
		t.Error("Array with duplicates failed")
	}
}

func TestXaydSort_Reversed(t *testing.T) {
	arr := make([]int, 100)
	for i := range arr {
		arr[i] = 100 - i
	}
	XaydSort(arr)
	if !isSorted(arr) {
		t.Error("Reversed array failed")
	}
}

func BenchmarkXaydSort_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := generateRandomArray(100, int64(i))
		b.StartTimer()
		XaydSort(arr)
	}
}

func BenchmarkXaydSort_1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := generateRandomArray(1000, int64(i))
		b.StartTimer()
		XaydSort(arr)
	}
}

func BenchmarkXaydSort_10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := generateRandomArray(10000, int64(i))
		b.StartTimer()
		XaydSort(arr)
	}
}

func BenchmarkXaydSort_50000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := generateRandomArray(50000, int64(i))
		b.StartTimer()
		XaydSort(arr)
	}
}

func BenchmarkXaydSort_100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := generateRandomArray(100000, int64(i))
		b.StartTimer()
		XaydSort(arr)
	}
}

func BenchmarkQuickSort_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := generateRandomArray(100, int64(i))
		b.StartTimer()
		QuickSort(arr)
	}
}

func BenchmarkQuickSort_1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := generateRandomArray(1000, int64(i))
		b.StartTimer()
		QuickSort(arr)
	}
}

func BenchmarkQuickSort_10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := generateRandomArray(10000, int64(i))
		b.StartTimer()
		QuickSort(arr)
	}
}

func BenchmarkQuickSort_50000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := generateRandomArray(50000, int64(i))
		b.StartTimer()
		QuickSort(arr)
	}
}

func BenchmarkQuickSort_100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := generateRandomArray(100000, int64(i))
		b.StartTimer()
		QuickSort(arr)
	}
}

func BenchmarkGoSort_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := generateRandomArray(100, int64(i))
		b.StartTimer()
		sort.Ints(arr)
	}
}

func BenchmarkGoSort_1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := generateRandomArray(1000, int64(i))
		b.StartTimer()
		sort.Ints(arr)
	}
}

func BenchmarkGoSort_10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := generateRandomArray(10000, int64(i))
		b.StartTimer()
		sort.Ints(arr)
	}
}

func BenchmarkGoSort_50000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := generateRandomArray(50000, int64(i))
		b.StartTimer()
		sort.Ints(arr)
	}
}

func BenchmarkGoSort_100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := generateRandomArray(100000, int64(i))
		b.StartTimer()
		sort.Ints(arr)
	}
}

func BenchmarkXaydSort_NearlySorted_10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := generateNearlySorted(10000, int64(i))
		b.StartTimer()
		XaydSort(arr)
	}
}

func BenchmarkQuickSort_NearlySorted_10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := generateNearlySorted(10000, int64(i))
		b.StartTimer()
		QuickSort(arr)
	}
}
