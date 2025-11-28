package comparator_test

import (
	"fmt"
	"github.com/QAQandOwO/godget/comparator"
)

type product struct {
	name  string
	count uint
	desc  string
}

func ExampleComparator_Compare() {
	comparator := comparator.New[product](comparator.By(func(p product) string { return p.name })).
		ThenComparing(func(p1, p2 product) int { return int(p1.count) - int(p2.count) })

	p1 := product{name: "product1", count: 10}
	p2 := product{name: "product2", count: 5}
	p3 := product{name: "product1", count: 20}
	p4 := product{name: "product1", count: 10, desc: "product1 description"}

	fmt.Println(comparator.Compare(p1, p2))
	fmt.Println(comparator.Compare(p1, p3))
	fmt.Println(comparator.Compare(p2, p3))
	fmt.Println(comparator.Compare(p1, p4))

	// Output:
	// -1
	// -1
	// 1
	// 0
}

func ExampleComparator_Less() {
	comparator := comparator.New[product](comparator.By(func(p product) string { return p.name })).
		ThenComparing(func(p1, p2 product) int { return int(p1.count) - int(p2.count) })

	p1 := product{name: "product1", count: 10}
	p2 := product{name: "product2", count: 5}
	p3 := product{name: "product1", count: 20}
	p4 := product{name: "product1", count: 10, desc: "product1 description"}

	fmt.Println(comparator.Less(p1, p2))
	fmt.Println(comparator.Less(p1, p3))
	fmt.Println(comparator.Less(p2, p3))
	fmt.Println(comparator.Less(p1, p4))

	// Output:
	// true
	// true
	// false
	// false
}

func ExampleComparator_Greater() {
	comparator := comparator.New[product](comparator.By(func(p product) string { return p.name })).
		ThenComparing(func(p1, p2 product) int { return int(p1.count) - int(p2.count) })

	p1 := product{name: "product1", count: 10}
	p2 := product{name: "product2", count: 5}
	p3 := product{name: "product1", count: 20}
	p4 := product{name: "product1", count: 10, desc: "product1 description"}

	fmt.Println(comparator.Greater(p1, p2))
	fmt.Println(comparator.Greater(p1, p3))
	fmt.Println(comparator.Greater(p2, p3))
	fmt.Println(comparator.Greater(p1, p4))

	// Output:
	// false
	// false
	// true
	// false
}

func ExampleComparator_Equal() {
	comparator := comparator.New[product](comparator.By(func(p product) string { return p.name })).
		ThenComparing(func(p1, p2 product) int { return int(p1.count) - int(p2.count) })

	p1 := product{name: "product1", count: 10}
	p2 := product{name: "product2", count: 5}
	p3 := product{name: "product1", count: 20}
	p4 := product{name: "product1", count: 10, desc: "product1 description"}

	fmt.Println(comparator.Equal(p1, p2))
	fmt.Println(comparator.Equal(p1, p3))
	fmt.Println(comparator.Equal(p2, p3))
	fmt.Println(comparator.Equal(p1, p4))

	// Output:
	// false
	// false
	// false
	// true
}

func ExampleComparator_Min() {
	comparator := comparator.New[product](comparator.By(func(p product) string { return p.name })).
		ThenComparing(func(p1, p2 product) int { return int(p1.count) - int(p2.count) })

	products := []product{
		{name: "product1", count: 10},
		{name: "product2", count: 5},
		{name: "product1", count: 20},
		{name: "product1", count: 10, desc: "product1 description"},
	}

	fmt.Println(comparator.Min(products...))

	// Output:
	// {product1 10 }
}

func ExampleComparator_Max() {
	comparator := comparator.New[product](comparator.By(func(p product) string { return p.name })).
		ThenComparing(func(p1, p2 product) int { return int(p1.count) - int(p2.count) })

	products := []product{
		{name: "product1", count: 10},
		{name: "product2", count: 5},
		{name: "product1", count: 20},
		{name: "product1", count: 10, desc: "product1 description"},
	}

	fmt.Println(comparator.Max(products...))

	// Output:
	// {product2 5 }
}

func ExampleComparator_ReverseLast() {
	comparator := comparator.New[product](comparator.By(func(p product) string { return p.name })).
		ThenComparing(func(p1, p2 product) int { return int(p1.count) - int(p2.count) })

	p1 := product{name: "product1", count: 10}
	p2 := product{name: "product2", count: 5}
	p3 := product{name: "product1", count: 20}
	fmt.Println(comparator.Compare(p1, p2))
	fmt.Println(comparator.Compare(p1, p3))

	_ = comparator.ReverseLast()
	fmt.Println(comparator.Compare(p1, p2))
	fmt.Println(comparator.Compare(p1, p3))

	// Output:
	// -1
	// -1
	// -1
	// 1
}

func ExampleComparator_ReverseAll() {
	comparator := comparator.New[product](comparator.By(func(p product) string { return p.name })).
		ThenComparing(func(p1, p2 product) int { return int(p1.count) - int(p2.count) })

	p1 := product{name: "product1", count: 10}
	p2 := product{name: "product2", count: 5}
	p3 := product{name: "product1", count: 20}
	fmt.Println(comparator.Compare(p1, p2))
	fmt.Println(comparator.Compare(p1, p3))

	_ = comparator.ReverseAll()
	fmt.Println(comparator.Compare(p1, p2))
	fmt.Println(comparator.Compare(p1, p3))

	// Output:
	// -1
	// -1
	// 1
	// 1
}

func ExampleComparator_SortSlice() {
	comparator := comparator.New[product](comparator.By(func(p product) string { return p.name })).
		ThenComparing(func(p1, p2 product) int { return int(p1.count) - int(p2.count) })

	products := []product{
		{name: "product1", count: 10},
		{name: "product2", count: 5},
		{name: "product1", count: 20},
		{name: "product1", count: 10, desc: "product1 description"},
	}

	comparator.SortSlice(products)
	for _, p := range products {
		fmt.Println(p)
	}

	// Output:
	// {product1 10 }
	// {product1 10 product1 description}
	// {product1 20 }
	// {product2 5 }
}

func ExampleComparator_SliceIsSorted() {
	comparator := comparator.New[product](comparator.By(func(p product) string { return p.name })).
		ThenComparing(func(p1, p2 product) int { return int(p1.count) - int(p2.count) })

	products1 := []product{
		{name: "product1", count: 10},
		{name: "product2", count: 5},
		{name: "product1", count: 20},
		{name: "product1", count: 10, desc: "product1 description"},
	}
	products2 := []product{
		{name: "product1", count: 10, desc: "product1 description"},
		{name: "product1", count: 10},
		{name: "product1", count: 20},
		{name: "product2", count: 5},
	}

	fmt.Println(comparator.SliceIsSorted(products1))
	fmt.Println(comparator.SliceIsSorted(products2))

	// Output:
	// false
	// true
}
