package main

import (
	"fmt"
	"math/rand"
	"time"
)

type List interface {
	print(origem bool)
	createV(n int, sorted bool)
	selectionsort()
	bubblesort()
	inserctionsort()
	mergesort()
	quicksort()
	countingsort()
}

type vetor struct {
	origem   []int
	ordenado []int
}

// Função print para aux vizualização
func (l *vetor) print(origem bool, i int) {
	if origem {
		fmt.Print("[")
		for x := 0; x < i; x++ {
			fmt.Print(l.origem[x], ", ")
		}
		fmt.Print("... ")
		for x := i; x > 1; x-- {
			fmt.Print(l.origem[len(l.origem)-x], ", ")
		}
		fmt.Print(l.origem[len(l.origem)-1])
		fmt.Print("]\n")

	} else {
		fmt.Print("[")
		for x := 0; x < i; x++ {
			fmt.Print(l.ordenado[x], ", ")
		}
		fmt.Print("... ")
		for x := i; x > 1; x-- {
			fmt.Print(l.ordenado[len(l.ordenado)-x], ", ")
		}
		fmt.Print(l.ordenado[len(l.ordenado)-1])
		fmt.Print("]\n")
	}
}

// Função de criação de vetor aleatorio ou ordenado
func (l *vetor) createV(n int, sorted bool) {
	l.origem = make([]int, n)
	l.ordenado = make([]int, n)
	if sorted {
		for i := 0; i < n; i++ {
			l.origem[i] = i
		}
	} else {
		for i := 0; i < n; i++ {
			l.origem[i] = rand.Intn(n)
		}
	}
}

//Algoritmos de ordenação

func (l *vetor) selectionsort() {
	copy(l.ordenado, l.origem)
	start := time.Now()

	for i := 0; i < len(l.ordenado)-1; i++ {
		menor := i
		for j := i; j < len(l.ordenado); j++ {
			if l.ordenado[j] < l.ordenado[menor] {
				menor = j
			}
		}
		l.ordenado[i], l.ordenado[menor] = l.ordenado[menor], l.ordenado[i]
	}
	tempo := time.Since(start)
	println("\n\nTempo de ordenação do selection: ", tempo.Seconds(), "segundos\n")
}

func (l *vetor) bubblesort() {
	copy(l.ordenado, l.origem)
	start := time.Now()

	for i := 0; i < len(l.ordenado)-1; i++ {
		trocou := false
		for j := 0; j < len(l.ordenado)-1-i; j++ {
			if l.ordenado[j] > l.ordenado[j+1] {
				l.ordenado[j], l.ordenado[j+1] = l.ordenado[j+1], l.ordenado[j]
				trocou = true
			}
		}
		if !trocou {
			tempo := time.Since(start)
			println("\n\nTempo de ordenação do bubblesort: ", tempo.Seconds(), "segundos\n")
			return
		}
	}
}

func (l *vetor) inserctionsort() {
	copy(l.ordenado, l.origem)
	start := time.Now()

	for i := 1; i < len(l.ordenado); i++ {
		temp := l.ordenado[i]
		k := i
		for ; k >= 1 && l.ordenado[k-1] > temp; k-- {
			l.ordenado[k] = l.ordenado[k-1]
		}
		l.ordenado[k] = temp
	}
	tempo := time.Since(start)
	println("\n\nTempo de ordenação do inserctionsort: ", tempo.Seconds(), "segundos\n")
}

func aux_sort(lista []int) []int {

	if len(lista) <= 1 {
		return lista
	}
	meio := len(lista) / 2

	esquerda := aux_sort(lista[:meio])
	direita := aux_sort(lista[meio:])

	return aux_merge(esquerda, direita)
}

func aux_merge(esquerda []int, direita []int) []int {
	resultado := make([]int, 0, len(esquerda)+len(direita))
	i, j := 0, 0
	for i < len(esquerda) && j < len(direita) {
		if esquerda[i] <= direita[j] {
			resultado = append(resultado, esquerda[i])
			i++
		} else {
			resultado = append(resultado, direita[j])
			j++
		}
	}
	resultado = append(resultado, esquerda[i:]...)
	resultado = append(resultado, direita[j:]...)
	return resultado
}

func (l *vetor) mergesort() {
	copy(l.ordenado, l.origem)
	start := time.Now()

	l.ordenado = aux_sort(l.ordenado)

	tempo := time.Since(start)
	println("\n\nTempo de ordenação do mergesort: ", tempo.Seconds(), "segundos\n")
}

func fix_auxquicksort(v []int, left int, right int) {
	if left < right {
		pivotindex := aux_partition(v, left, right)
		fix_auxquicksort(v, left, pivotindex-1)
		fix_auxquicksort(v, pivotindex+1, right)
	}
}

func rand_auxquicksort(v []int, left int, right int) {
	if left < right {
		random := rand.Intn(right)
		v[right], v[random] = v[random], v[right]
		pivotindex := aux_partition(v, left, right)
		rand_auxquicksort(v, left, pivotindex-1)
		rand_auxquicksort(v, pivotindex+1, right)
	}
}

func aux_partition(v []int, left int, right int) int {
	pivot := v[right]
	i := left
	for j := left; j < right; j++ {
		if v[j] < pivot {
			v[i], v[j] = v[j], v[i]
			i++
		}
	}
	v[i], v[right] = v[right], v[i]
	return i
}

func (l *vetor) quicksort() {
	copy(l.ordenado, l.origem)

	start := time.Now()
	fix_auxquicksort(l.ordenado, 0, len(l.ordenado)-1)
	tempo := time.Since(start)
	println("\n\nTempo de ordenação do quicksort fixo: ", tempo.Seconds(), "segundos\n")

	start = time.Now()
	rand_auxquicksort(l.ordenado, 0, (len(l.ordenado) - 1))
	tempo = time.Since(start)
	println("Tempo de ordenação do quicksort randomizado: ", tempo.Seconds(), "segundos\n")
}

func (l *vetor) countingsort() {
	copy(l.ordenado, l.origem)
	start := time.Now()

	if len(l.ordenado) == 0 {
		return
	}

	max := l.ordenado[0]
	for v := 0; v < len(l.ordenado); v++ {
		if l.ordenado[v] > max {
			max = l.ordenado[v]
		}
	}
	count := make([]int, max+1)

	for v := 0; v < len(l.ordenado); v++ {
		count[l.ordenado[v]]++
	}

	index := 0
	for i, c := range count {
		for c > 0 {
			l.ordenado[index] = i
			index++
			c--
		}
	}
	tempo := time.Since(start)
	println("\n\nTempo de ordenação do countingsort fixo: ", tempo.Seconds(), "segundos\n")
}

//Execução

func main() {
	i := 10
	l := &vetor{}
	l.createV(100000, true)
	l.print(true, i)

	l.selectionsort()
	l.print(false, i)

	l.bubblesort()
	l.print(false, i)

	l.inserctionsort()
	l.print(false, i)

	l.mergesort()
	l.print(false, i)

	l.quicksort()
	l.print(false, i)

	l.countingsort()
	l.print(false, i)
}
