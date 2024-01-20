package main

/*
Гоше дали задание написать красивую сортировку слиянием. Поэтому Гоше обязательно надо реализовать отдельно функцию merge и функцию merge_sort.
Функция merge принимает один массив и три целочисленных индекса: left, mid, и right. Функция сливает две отсортированные части одного и того же массива в один отсортированный массив. Первая часть массива определяется полуинтервалом
[left,mid)
 массива array, а вторая часть – полуинтервалом
[mid,right)
 того же массива array. Функция возвращает сливаемый массив.
Функция merge_sort принимает некоторый подмассив, который нужно отсортировать. Подмассив задаётся полуинтервалом — его началом и концом. Функция должна отсортировать передаваемый в неё подмассив, она ничего не возвращает.
Функция merge_sort разбивает полуинтервал на две половинки и рекурсивно вызывает сортировку отдельно для каждой. Затем два отсортированных массива сливаются в один с помощью merge.
Заметьте, что в функции передаются именно полуинтервалы
[begin,end)
, то есть правый конец не включается. Например, если вызвать merge_sort(arr, 0, 4), где arr=[4,5,3,0,1,2]
, то будут отсортированы только первые четыре элемента, изменённый массив будет выглядеть как arr=[0,3,4,5,1,2].
Реализуйте эти две функции.
Используйте заготовки кода для данной задачи, расположенные по ссылкам: https://github.com/Yandex-Practicum/algorithms-templates/tree/main/go/sprint3/K
Передаваемый в функции массив состоит из целых чисел, не превосходящих по модулю 10^9
. Длина сортируемого диапазона не превосходит 10^5.

*/

import "reflect"

func merge_sort(arr []int, lf int, rg int) {
	if len(arr) == 1 || rg-lf <= 1 {
		return
	}
	mid := (rg + lf) / 2
	merge_sort(arr, lf, mid)
	merge_sort(arr, mid, rg)
	res := merge(arr, lf, mid, rg)
	for i, t := range res {
		arr[lf+i] = t
	}
}

func merge(arr []int, left int, mid int, right int) (result []int) {
	i := left
	j := mid
	for {
		if i >= mid || j >= right {
			break
		}
		if arr[i] <= arr[j] {
			result = append(result, arr[i])
			i++
		} else {
			result = append(result, arr[j])
			j++
		}
	}
	for i < mid {
		result = append(result, arr[i])
		i++
	}
	for j < right {
		result = append(result, arr[j])
		j++
	}
	return
}

func main() {
	a := []int{1, 4, 9, 2, 10, 11}
	b := merge(a, 0, 3, 6)
	expected := []int{1, 2, 4, 9, 10, 11}
	if !reflect.DeepEqual(b, expected) {
		panic("WA. Merge")
	}

	c := []int{1, 4, 2, 10, 1, 2}
	merge_sort(c, 0, 6)
	expected = []int{1, 1, 2, 2, 4, 10}
	if !reflect.DeepEqual(c, expected) {
		panic("WA. MergeSort")
	}
}
