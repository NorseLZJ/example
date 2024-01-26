package dbm

import (
	"log"
	"sync"
)

type userManager struct {
	userMap map[uint64]*DbUser
	sync.Mutex
}

var (
	UserManager *userManager
)

func LoadAllUser() {
	users := make([]*DbUser, 0)
	if err := JavaDb().Find(&users).Error; err != nil {
		log.Fatal(err)
	}
	userMap := map[uint64]*DbUser{}
	for _, user := range users {
		userMap[uint64(user.UserId)] = user
	}
	UserManager = &userManager{userMap: userMap}
}

func (u *userManager) CollectUsers(ids []uint64) map[int]*DbUser {
	result := make(map[int]*DbUser)
	UserManager.Lock()
	for _, id := range ids {
		if u.userMap[id] != nil {
			result[int(id)] = u.userMap[id]
		}
	}
	UserManager.Unlock()
	return result
}
