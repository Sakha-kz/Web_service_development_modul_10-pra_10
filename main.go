package main

import (
	"fmt"
	"strings"
)

// Sale представляет информацию о продаже
type Sale struct {
	ID       int
	Product  string
	Category string
	Price    float64
	Quantity int
}

// 1. CalculateRevenue рассчитывает общую сумму продаж с использованием среза и цикла for range
func CalculateRevenue(sales []Sale) float64 {
	var total float64
	for _, s := range sales {
		total += s.Price * float64(s.Quantity)
	}
	return total
}

// 2. GetTopProduct определяет товар с наибольшим суммарным количеством проданных единиц
func GetTopProduct(sales []Sale) (string, int) {
	counts := make(map[string]int)
	for _, s := range sales {
		counts[s.Product] += s.Quantity
	}

	var topProduct string
	var maxQty int
	for prod, qty := range counts {
		if qty > maxQty {
			maxQty = qty
			topProduct = prod
		}
	}
	return topProduct, maxQty
}

// 3. RevenueByCategory рассчитывает общую выручку по каждой категории
func RevenueByCategory(sales []Sale) map[string]float64 {
	byCat := make(map[string]float64)
	for _, s := range sales {
		byCat[s.Category] += s.Price * float64(s.Quantity)
	}
	return byCat
}

// 4. ApplyDiscount применяет процент скидки к цене через указатель на структуру Sale
func ApplyDiscount(sale *Sale, discountPercent float64) {
	sale.Price = sale.Price * (1.0 - discountPercent/100.0)
}

func formatCurrency(val float64) string {
	// Форматирование числа с разделением тысяч пробелами
	intVal := int(val)
	str := fmt.Sprintf("%d", intVal)
	var parts []string
	for len(str) > 3 {
		parts = append([]string{str[len(str)-3:]}, parts...)
		str = str[:len(str)-3]
	}
	parts = append([]string{str}, parts...)
	return strings.Join(parts, " ") + " тг"
}

func main() {
	// 7. Использование массива скидок согласно заданию
	discountOptions := [3]float64{5, 10, 15}

	// Инициализация среза продаж согласно заданию
	sales := []Sale{
		{ID: 1, Product: "Ноутбук", Category: "Электроника", Price: 350000, Quantity: 2},
		{ID: 2, Product: "Мышь", Category: "Аксессуары", Price: 5000, Quantity: 10},
		{ID: 3, Product: "Монитор", Category: "Электроника", Price: 85000, Quantity: 3},
		{ID: 4, Product: "Клавиатура", Category: "Аксессуары", Price: 12000, Quantity: 5},
		{ID: 5, Product: "Наушники", Category: "Аксессуары", Price: 25000, Quantity: 4},
		{ID: 6, Product: "Ноутбук", Category: "Электроника", Price: 350000, Quantity: 1},
	}

	// Подсчет общего количества проданных единиц
	totalUnits := 0
	for _, s := range sales {
		totalUnits += s.Quantity
	}

	revenueBefore := CalculateRevenue(sales)
	avgSale := revenueBefore / float64(len(sales))
	topProd, topQty := GetTopProduct(sales)
	catRevenue := RevenueByCategory(sales)

	fmt.Println("============ SALES ANALYZER ============")
	fmt.Printf("Количество записей: %d\n", len(sales))
	fmt.Printf("Продано единиц товара: %d\n", totalUnits)
	fmt.Printf("Общая выручка: %s\n", formatCurrency(revenueBefore))
	fmt.Printf("Средняя сумма продажи: %s\n", formatCurrency(avgSale))
	fmt.Println("Самый продаваемый товар:")
	fmt.Printf("%s — %d шт.\n", topProd, topQty)
	fmt.Println("Выручка по категориям:")
	fmt.Printf("Электроника: %s\n", formatCurrency(catRevenue["Электроника"]))
	fmt.Printf("Аксессуары: %s\n", formatCurrency(catRevenue["Аксессуары"]))

	// Применение скидки 10% (второй элемент массива discountOptions[1]) к первой записи о ноутбуке (&sales[0])
	oldPrice := sales[0].Price
	chosenDiscount := discountOptions[1] // 10%
	ApplyDiscount(&sales[0], chosenDiscount)
	newPrice := sales[0].Price

	revenueAfter := CalculateRevenue(sales)

	fmt.Printf("\nПрименение скидки %.0f%%:\n", chosenDiscount)
	fmt.Printf("Товар: %s\n", sales[0].Product)
	fmt.Printf("Старая цена: %s\n", formatCurrency(oldPrice))
	fmt.Printf("Новая цена: %s\n", formatCurrency(newPrice))
	fmt.Printf("Выручка после скидки: %s\n", formatCurrency(revenueAfter))
}
