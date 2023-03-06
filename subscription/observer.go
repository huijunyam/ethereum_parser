package subscription

import (
	"ethereum_parser/models"
	"log"
)

type Observer interface {
	Update(transaction models.Transaction)
	GetUser() string
}

type AddressObserver struct {
	user string
}

func NewAddressObserver(user string) *AddressObserver {
	return &AddressObserver{user: user}
}

func (t *AddressObserver) Update(transaction models.Transaction) {
	// todo do something: inform user when transaction occur
	log.Printf("notify user: %s, hash:%s", t.user, transaction.Hash)
}

func (t *AddressObserver) GetUser() string {
	return t.user
}
