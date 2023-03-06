package subscription

import (
	"errors"
	"ethereum_parser/models"
)

type Subject interface {
	Attach(Observer) (bool, error)
	Detach(Observer) (bool, error)
	Notify(models.Transaction) (bool, error)
}

type AddressMonitor struct {
	address     string
	observers   []Observer
	observerIdx map[string]int
}

func NewAddressMonitor(address string) *AddressMonitor {
	return &AddressMonitor{
		address:     address,
		observers:   make([]Observer, 0),
		observerIdx: make(map[string]int),
	}
}

func (a *AddressMonitor) Attach(o Observer) (bool, error) {
	user := o.GetUser()
	if _, ok := a.observerIdx[user]; ok {
		return false, errors.New("observer already exists")
	}
	a.observers = append(a.observers, o)
	a.observerIdx[user] = len(a.observers) - 1
	return true, nil
}

func (a *AddressMonitor) Detach(o Observer) (bool, error) {
	user := o.GetUser()
	if _, ok := a.observerIdx[user]; !ok {
		return false, errors.New("observer not found")
	}
	idx := a.observerIdx[user]
	a.observers = append(a.observers[:idx], a.observers[idx+1:]...)
	delete(a.observerIdx, user)
	return true, nil
}

func (a *AddressMonitor) Notify(t models.Transaction) (bool, error) {
	for _, observer := range a.observers {
		observer.Update(t)
	}
	return true, nil
}

func (a *AddressMonitor) SetTransaction(t models.Transaction) {
	a.Notify(t)
}
