package main

import (
	"errors"
	"fmt"
)

// DynamicArray представляет динамический массив строк
type DynamicArray struct {
	data []string
}

// NewDynamicArray создает новый динамический массив
func NewDynamicArray() *DynamicArray {
	return &DynamicArray{
		data: make([]string, 0, 10), // Начальная емкость 10
	}
}

// Add добавляет элемент в конец
func (da *DynamicArray) Add(value string) {
	da.data = append(da.data, value)
}

// Insert вставляет элемент по индексу
func (da *DynamicArray) Insert(index int, value string) error {
	if index < 0 || index > len(da.data) {
		return errors.New("индекс вне диапазона")
	}
	
	// Расширяем слайс
	da.data = append(da.data, "")
	// Сдвигаем элементы
	copy(da.data[index+1:], da.data[index:])
	// Вставляем значение
	da.data[index] = value
	
	return nil
}

// Get возвращает элемент по индексу
func (da *DynamicArray) Get(index int) (string, error) {
	if index < 0 || index >= len(da.data) {
		return "", errors.New("индекс вне диапазона")
	}
	return da.data[index], nil
}

// Remove удаляет элемент по индексу
func (da *DynamicArray) Remove(index int) error {
	if index < 0 || index >= len(da.data) {
		return errors.New("индекс вне диапазона")
	}
	
	// Сдвигаем элементы
	da.data = append(da.data[:index], da.data[index+1:]...)
	return nil
}

// Set заменяет элемент по индексу
func (da *DynamicArray) Set(index int, value string) error {
	if index < 0 || index >= len(da.data) {
		return errors.New("индекс вне диапазона")
	}
	
	da.data[index] = value
	return nil
}

// Size возвращает размер массива
func (da *DynamicArray) Size() int {
	return len(da.data)
}

// Capacity возвращает емкость массива
func (da *DynamicArray) Capacity() int {
	return cap(da.data)
}

// Print выводит все элементы
func (da *DynamicArray) Print() {
	fmt.Println("=== Элементы массива ===")
	for i, v := range da.data {
		fmt.Printf("[%d] %s\n", i, v)
	}
	fmt.Println()
}

func main() {
	da := NewDynamicArray()
	
	// Добавление элементов
	da.Add("Яблоко")
	da.Add("Банан")
	da.Add("Апельсин")
	
	fmt.Println("Начальный массив:")
	da.Print()
	
	// Вставка
	if err := da.Insert(1, "Груша"); err != nil {
		fmt.Println("Ошибка вставки:", err)
	}
	fmt.Println("После вставки 'Груша' на позицию 1:")
	da.Print()
	
	// Получение
	if value, err := da.Get(2); err != nil {
		fmt.Println("Ошибка получения:", err)
	} else {
		fmt.Println("Элемент с индексом 2:", value)
	}
	
	// Замена
	if err := da.Set(0, "Киви"); err != nil {
		fmt.Println("Ошибка замены:", err)
	}
	fmt.Println("После замены элемента 0 на 'Киви':")
	da.Print()
	
	// Удаление
	if err := da.Remove(1); err != nil {
		fmt.Println("Ошибка удаления:", err)
	}
	fmt.Println("После удаления элемента 1:")
	da.Print()
	
	// Информация
	fmt.Printf("Размер: %d\n", da.Size())
	fmt.Printf("Емкость: %d\n", da.Capacity())
}
