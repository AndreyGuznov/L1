package main

import "fmt"

// Сторонний интерфейс который необходимо адаптировать путём реализации
type Target interface {
	Request() string
}

// Создаем структуры для которую необходимо адаптировать
type ForAdapting struct {
}

// Конструктор
func NewAdapt(adapt *ForAdapting) Target {
	return &Adapter{
		adapt,
	}
}

// Метод реализующий запрос адаптируемой структуры
func (a *ForAdapting) SpecificRequest() string {
	return "SomeRequest"
}

// Структура которая реализует интерфейс "стороннего" интерфейса Target
type Adapter struct {
	*ForAdapting
}

// Метод для выполнения запроса
func (a *Adapter) Request() string {
	return a.SpecificRequest()
}

func main() {
	adapt1 := NewAdapt(&ForAdapting{})
	req := adapt1.Request()
	fmt.Println(req)
}
