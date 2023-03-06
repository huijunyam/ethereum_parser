package apis

import (
	"encoding/json"
	"errors"
	"ethereum_parser/cron"
	"ethereum_parser/subscription"
	"net/http"
)

func subscribe(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		bindErrorResponse(w, errors.New("invalid request method"), http.StatusMethodNotAllowed)
		return
	}

	var s SubscriptionReq
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		bindErrorResponse(w, errors.New("address is required"), http.StatusBadRequest)
		return
	}

	if s.Address == "" {
		bindErrorResponse(w, errors.New("address cannot be empty"), http.StatusBadRequest)
		return
	}

	if s.User == "" {
		s.User = "default_test_user"
	}

	var monitor *subscription.AddressMonitor
	cron.AddressSubscriptionLock.Lock()
	defer cron.AddressSubscriptionLock.Unlock()
	if m, ok := cron.AddressSubscription[s.Address]; ok {
		monitor = m
	} else {
		monitor = subscription.NewAddressMonitor(s.Address)
		cron.AddressSubscription[s.Address] = monitor
	}

	observer := subscription.NewAddressObserver(s.User)
	isAttached, err := monitor.Attach(observer)
	if err != nil {
		bindErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	bindResponse(w, isAttached, nil)
}
