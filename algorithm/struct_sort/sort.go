package main

import (
	"fmt"
	"sort"
)

// 结构体切片按照多个字段排序，支持原生的sort实现和自定义排序算法实现

// Person 定义结构体
type Person struct {
	Name    string
	Age     int
	Sex     string
	City    string
	Balance float64
}

type People []Person

// Len 实现 sort.Interface 接口的 Len 方法
func (p People) Len() int { return len(p) }

// Less 实现 sort.Interface 接口的 Less 方法
func (p People) Less(i, j int) bool {
	// 先按照 Name 字段排序
	if p[i].Name != p[j].Name {
		return p[i].Name < p[j].Name
	}
	if p[i].Balance != p[j].Balance {
		return p[i].Balance < p[j].Balance
	}
	if p[i].Sex != p[j].Sex {
		return p[i].Sex < p[j].Sex
	}
	if p[i].Age != p[j].Age {
		return p[i].Age < p[j].Age
	}
	return p[i].City < p[j].City
}

// Swap 实现 sort.Interface 接口的 Swap 方法
func (p People) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func NativeSort(people People) {
	sort.Sort(people)
}

func QuickSort(people People, low, high int) {
	if low < high {
		pi := Partition(people, low, high)
		QuickSort(people, low, pi-1)
		QuickSort(people, pi+1, high)
	}
}

func Partition(people People, low, high int) int {
	pivot := people[high]
	i := low - 1
	for j := low; j < high; j++ {
		if people[j].Name < pivot.Name {
			i++
			people[i], people[j] = people[j], people[i]
		}
	}
	people[i+1], people[high] = people[high], people[i+1]
	return i + 1
}

func main() {
	// 初始化结构体切片
	people := People{Person{
		Name:    "zhangsan",
		Age:     23,
		Sex:     "male",
		City:    "Shanghai",
		Balance: 100.12,
	}, Person{
		Name:    "lisi",
		Age:     12,
		Sex:     "female",
		City:    "Shenzhen",
		Balance: 10099,
	}, Person{
		Name:    "wangwu",
		Age:     43,
		Sex:     "female",
		City:    "Beijing",
		Balance: 190.1,
	}, Person{
		Name:    "zhaoliu",
		Age:     12,
		Sex:     "female",
		City:    "Shanghai",
		Balance: 90,
	}, Person{
		Name:    "zhaoliu",
		Age:     23,
		Sex:     "female",
		City:    "Shanghai",
		Balance: 99,
	}, Person{
		Name:    "zhangsan",
		Age:     3,
		Sex:     "female",
		City:    "Chengdu",
		Balance: 100.12,
	}, Person{
		Name:    "zhangsan",
		Age:     4,
		Sex:     "female",
		City:    "Chengdu",
		Balance: 100.12,
	}, Person{
		Name:    "tianqi",
		Age:     19,
		Sex:     "female",
		City:    "Chongqing",
		Balance: 100.12,
	}}

	//NativeSort(people)
	QuickSort(people, 0, people.Len()-1)
	// 打印排序后的结果
	for _, person := range people {
		fmt.Printf("%s - %s - %d - %s - %f\n", person.Name, person.Sex, person.Age, person.City, person.Balance)
	}
}
