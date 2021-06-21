package models

// Это очень приближенная структура типа User. Она будет расширяться и изменяться по мере прохождения уроков
// и выяснения требований к интерфейсам типа. Также, я бы приземлял структуру типа на особенности выбранной
// СУБД для проекта и требований к API.
// Если подытожить, кажется преждевременно прорабатывать вид типа User, с точки зрения как надо.

type User struct {
	Id         uint64
	CalendarId uint64
	ResumeId   uint64
	Name       string
	Surname    string
	Patronymic string
	Email      string
}

type UserSearchParams struct {
	Count  uint64
	Offset uint64
}

type UserSearchResult struct {
	Items      []User
	NextOffset uint64
}
