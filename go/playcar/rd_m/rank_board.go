package route

import (
	"errors"
	"fmt"
	redis "github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"log"
	"playcar/dbm"
	"strconv"
	"time"
)

var (
	ctx = context.Background()
)

type PlayerScore struct {
	Rank     int
	PlayerID string
	Score    int
	NickName string
	Phone    string
}

type Leaderboard struct {
	client *redis.Client
	db     *gorm.DB
}

var (
	RankBoard *Leaderboard
)

func InitLeaderboard(addr, password string, db int) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		DB:       db,
		Password: password,
	})
	if err := client.Ping(ctx).Err(); err != nil {
		log.Fatal(err)
	}
	RankBoard = &Leaderboard{client: client}
}

func (lb *Leaderboard) AddScoreBig(key string, playerID uint64, score float64) {
	currentScore, err := lb.client.ZScore(ctx, key, fmt.Sprintf("%d", playerID)).Result()
	if err != nil && !errors.Is(redis.Nil, err) {
		return
	}
	if score > currentScore {
		err := lb.client.ZAdd(ctx, key, &redis.Z{Member: playerID, Score: score}).Err()
		if err != nil {
			log.Println(err)
		}
	}
}

func (lb *Leaderboard) AddScoreSmall(key string, playerID uint64, score float64) {
	currentScore, err := lb.client.ZScore(ctx, key, fmt.Sprintf("%d", playerID)).Result()
	if err != nil && !errors.Is(redis.Nil, err) {
		return
	}
	if currentScore == 0 || score < currentScore {
		err := lb.client.ZAdd(ctx, key, &redis.Z{Member: playerID, Score: score}).Err()
		if err != nil {
			log.Println(err)
		}
	}
}

func (lb *Leaderboard) GetTopN(key string, n int) []*PlayerScore {
	result, err := lb.client.ZRangeWithScores(ctx, key, 0, int64(n-1)).Result()
	//result, err := lb.client.ZRevRangeWithScores(ctx, "leaderboard", 0, int64(n-1)).Result()
	if err != nil {
		log.Fatal(err)
	}

	var topScores []*PlayerScore
	var ids []uint64
	for idx, z := range result {
		topScores = append(topScores, &PlayerScore{
			Rank:     idx + 1,
			PlayerID: z.Member.(string),
			Score:    int(z.Score),
		})
		uid, _ := strconv.Atoi(z.Member.(string))
		ids = append(ids, uint64(uid))
	}

	users := dbm.UserManager.CollectUsers(ids)
	for _, vv := range topScores {
		playerId, _ := strconv.Atoi(vv.PlayerID)
		if users[playerId] != nil {
			vv.NickName = users[playerId].NickName
			vv.Phone = users[playerId].PhoneNumber
		}
	}

	return topScores
}

func (lb *Leaderboard) AddOnline(playerID uint64, token string) {
	key := fmt.Sprintf("online:%d", playerID)
	err := lb.client.Set(ctx, key, token, time.Hour*24).Err()
	if err != nil {
		log.Println(err)
	}
}

func (lb *Leaderboard) IsOnline(playerID uint64, token string) bool {
	key := fmt.Sprintf("online:%d", playerID)
	value, err := lb.client.Get(ctx, key).Result()
	if err != nil {
		log.Println(err)
	}
	return value == token
}
