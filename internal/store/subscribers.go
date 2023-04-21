package store

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/redis/go-redis/v9"
)

const (
	setKey    = "subscribers"
	hostname  = "db"
	redisPort = "6379"
)

var (
	ctx                = context.Background()
	redisServerAddress = fmt.Sprint(hostname, ":", redisPort)
	emptyPassword      = ""
)

type Subscribers struct {
	persister *redis.Client
}

func NewSubscribers() Subscribers {
	persister := redis.NewClient(&redis.Options{
		Addr:     redisServerAddress,
		Password: emptyPassword,
		DB:       0,
	})
	return Subscribers{persister: persister}
}

func (m *Subscribers) All() []int64 {
	array := m.persister.SMembers(ctx, setKey)
	result := []int64{}
	for _, a := range array.Val() {
		userID, err := strconv.ParseInt(a, 10, 64)
		if err != nil {
			log.Printf("Failed to parse user id %s to int64", a)
		}
		result = append(result, userID)
	}
	log.Printf("All subscribers are %s\n", array.Val())
	return result
}

func (m *Subscribers) Add(user int64) {
	log.Printf("Add user %d to subscribers; set of subscribers is %s\n", user, m.persister.SMembers(ctx, setKey))
	m.persister.SAdd(ctx, setKey, user)
	log.Printf("Set after adding user %d = %s\n", user, m.persister.SMembers(ctx, setKey))
}

func (m *Subscribers) Rm(user int64) {
	log.Printf("Start rm user %d from subscribers; subscribers set is %v\n", user, m.persister.SMembers(ctx, setKey))
	ctx := context.Background()
	m.persister.SRem(ctx, setKey, user)
	log.Printf("Set after drop user %d is %s\n", user, m.persister.SMembers(ctx, setKey))
}
