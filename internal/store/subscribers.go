package store

import "log"

type Subscribers struct {
	subscribers []int64
}

func NewSubscribers() Subscribers {
	return Subscribers{subscribers: []int64{}}
}

func (m *Subscribers) All() []int64 {
	return m.subscribers
}

func (m *Subscribers) Add(user int64) Subscribers {
	log.Printf("Add user %v to subscribers", user)
	if contains(m.subscribers, user) {
		return *m
	}
	return Subscribers{subscribers: append(m.subscribers, user)}
}

func (m *Subscribers) Rm(user int64) Subscribers {
	log.Printf("Start rm user %v to subscribers", user)
	contains := false
	var index int
	for i, elem := range m.subscribers {
		if elem == user {
			log.Printf("Find user %v in subscribers, deleting...", user)
			index = i
			contains = true
			break
		}
	}
	if contains {
		log.Printf("Deleting user %v from subscribers...", user)
		return Subscribers{subscribers: remove(m.subscribers, index)}
	}
	return *m
}

func remove(slice []int64, s int) []int64 {
	return append(slice[:s], slice[s+1:]...)
}

func contains(xs []int64, x int64) bool {
	for _, elem := range xs {
		if elem == x {
			return true
		}
	}
	return false
}
