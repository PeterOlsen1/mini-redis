package auth

import (
	"mini-redis/server/auth/authtypes"
	"sync"
)

var userMap map[string]*authtypes.User = make(map[string]*authtypes.User)
var mu sync.Mutex

func GetUser(username string) (*authtypes.User, bool) {
	mu.Lock()
	defer mu.Unlock()

	user, ok := userMap[username]
	return user, ok
}

func SetUser(user *authtypes.User) {
	mu.Lock()
	defer mu.Unlock()

	userMap[user.Username] = user
}

func GetAllUsers() []*authtypes.User {
	mu.Lock()
	defer mu.Unlock()

	out := make([]*authtypes.User, 0)
	for _, user := range userMap {
		out = append(out, user)
	}

	return out
}

func DeleteUser(user string) {
	mu.Lock()
	defer mu.Unlock()

	delete(userMap, user)
}
