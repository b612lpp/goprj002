package repository

import (
	"github.com/b612lpp/goprj002/domain"
)

// Интерфейс для хранилища бизнес данных
type Repo interface {
	// Возвращает последнюю запись по ключу юзер(u)+тип счетчика(t)
	GetLast(string, string) (domain.MeterReading, error)
	Save(reading domain.MeterReading) error
	GetTitle() string
}

// Интерфейс для хранилища событий из бизнес слоя
type EventStore interface {
	AddEvent(domain.EventuallyAppliedData) error
	GetTitle() string
}

// Интерфейс объединяющий хранилища для событий двух типов. Эвент и чтение\запись в бд
type ReadingStorage interface {
	GetLast(string, string) (domain.MeterReading, error)
	Save(reading domain.MeterReading) error
	AddEvent(domain.EventuallyAppliedData) error
	GetTitle() string
}

// Структура варехауса удовлетворяющая интерфейсу ReadingStorage. содержит в себе необходимые в бизнес сценарии интерфейсы
type WareHouse struct {
	appData   Repo
	eventData EventStore
}

// Функция создания экземпляра хранилища отвечающего интерфейсу ReadingStorage. ЮЗ кейс бработки входящих клиентских данных по счетчикам
func NewWareHouse(r Repo, e EventStore) ReadingStorage {
	return &WareHouse{appData: r, eventData: e}
}

func (wh *WareHouse) AddEvent(v domain.EventuallyAppliedData) error {
	return wh.eventData.AddEvent(v)

}

func (wh *WareHouse) GetLast(u string, t string) (domain.MeterReading, error) {
	return wh.appData.GetLast(u, t)
}

func (wh *WareHouse) Save(reading domain.MeterReading) error {
	return wh.appData.Save(reading)
}

func (wh *WareHouse) GetTitle() string {

	return "Подключены " + wh.appData.GetTitle() + " " + wh.eventData.GetTitle()
}
