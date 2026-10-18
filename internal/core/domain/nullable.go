package domain

// Нужен троичной логики
// 1.Set==false значение не передали
// 2.Set==true, Value=nil передали nil значит удаляем номер
// 3.Set==true, Value=&v добавляем или меняем номер
type Nullable[T any] struct {
	Value *T
	Set   bool
}
