package main

/*
Uma empresa de sistemas precisa analisar que algoritmos de ordenamento utilizar para seus
serviços.
Para eles é necessário instanciar 3 arrays com valores aleatórios não ordenados
- uma matriz de inteiros com 100 valores
- uma matriz de inteiros com 1000 valores
- uma matriz de inteiros com 10.000 valores

Para instanciar as variáveis, utilize o rand:

Cada um deve ser ordenado por:
- Inserção
- grupo
- seleção

Uma rotina para cada execução de classificação
Tenho que esperar terminar a ordenação de 100 números para continuar com a ordenação de
1000 e depois a ordenação de 10000.
Por fim, devo medir o tempo de cada um e mostrar o resultado na tela, para saber qual
ordenação ficou melhor para cada arranjo.
*/

import (
	"fmt"
	"math/rand"
	"time"
)

func insertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key
	}
	return arr
}

func bubbleSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func selectionSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
	return arr
}

func main() {
	arr100 := rand.Perm(100)
	arr1000 := rand.Perm(1000)
	arr10000 := rand.Perm(10000)

	start := time.Now()
	insertionSort(arr100)
	elapsed := time.Since(start)
	fmt.Printf("Tempo de execução para ordenação de 100 com inserção: %s\n", elapsed)

	start = time.Now()
	bubbleSort(arr100)
	elapsed = time.Since(start)
	fmt.Printf("Tempo de execução para ordenação de 100 com bubble sort: %s\n", elapsed)

	start = time.Now()
	selectionSort(arr100)
	elapsed = time.Since(start)
	fmt.Printf("Tempo de execução para ordenação de 100 com seleção: %s\n", elapsed)

	start = time.Now()
	insertionSort(arr1000)
	elapsed = time.Since(start)
	fmt.Printf("Tempo de execução para ordenação de 1000 com inserção: %s\n", elapsed)

	start = time.Now()
	bubbleSort(arr1000)
	elapsed = time.Since(start)
	fmt.Printf("Tempo de execução para ordenação de 1000 com bubble sort: %s\n", elapsed)

	start = time.Now()
	selectionSort(arr1000)
	elapsed = time.Since(start)
	fmt.Printf("Tempo de execução para ordenação de 1000 com seleção: %s\n", elapsed)

	start = time.Now()
	insertionSort(arr10000)
	elapsed = time.Since(start)
	fmt.Printf("Tempo de execução para ordenação de 10000 com inserção: %s\n", elapsed)
}
