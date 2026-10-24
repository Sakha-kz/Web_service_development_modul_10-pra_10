# Модуль 10. Практическая работа: Sales Analyzer на Golang

Практическая работа по дисциплине «Разработка веб-сервисов» (CSE5032).  
Кафедра Программной инженерии, Институт Автоматики и информационных технологий.

---

## 🎯 Цель работы
Закрепить навыки работы с массивами, срезами (`slice`), ассоциативными картами (`map`), функциями и указателями (`pointers`) в языке Go.

---

## 📂 Структура проекта

```text
Web_service_development_modul_10-pra_10/
├── .gitignore
├── go.mod
├── main.go
├── README.md
└── DESIGN_AND_ANSWERS.md
```

---

## 📋 Реализованные функции и алгоритмы
1. `CalculateRevenue(sales []Sale) float64` — вычисление суммарной выручки по всем продажам.
2. `GetTopProduct(sales []Sale) (string, int)` — группировка продаж по товарам через `map[string]int` и поиск лидера продаж.
3. `RevenueByCategory(sales []Sale) map[string]float64` — подсчет выручки по категориям товаров.
4. `ApplyDiscount(sale *Sale, discount float64)` — применение скидки непосредственно по указателю к объекту структуры.
5. Использование фиксированного массива `[3]float64{5, 10, 15}` для выбора процента скидки.

---

## 🚀 Запуск программы

```powershell
cd E:\Web_service_development_RK1\week_10\Web_service_development_modul_10-pra_10
go run main.go
```
