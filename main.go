package main

import (
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis"
)

const (
	sessionID = "54160123"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	response, err := client.Ping().Result()
	fmt.Println(response, err)

	//
	/*
		session := map[string]interface{}{
			"from":     "thai",
			"to":       "worldwide",
			"duration": "54160123:duration",
		}
		duration := map[string]interface{}{
			"start": "28/10/2018",
			"end":   "30/10/2018",
		}
		err = client.HMSet(sessionID, session).Err()
		fmt.Println(err)
		err = client.HMSet(sessionID+":duration", duration).Err()
		fmt.Println(err)

		sessionData, err := client.HGetAll(sessionID).Result()
		fmt.Println(sessionData, err)
		sessionData, err = client.HGetAll(sessionData["duration"]).Result()
		fmt.Println(sessionData, err)
	*/
	//
	session := map[string]interface{}{
		"from": "thai",
		"to":   "worldwide",
		"duration": map[string]interface{}{
			"start": "28/10/2018",
			"end":   "30/10/2018",
		},
	}
	data, _ := json.Marshal(session)
	err = client.Set(sessionID, string(data), 0).Err()
	fmt.Println(err)

	sessionData, err := client.Get(sessionID).Result()
	fmt.Println(sessionData, err)
}
