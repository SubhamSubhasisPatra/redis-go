package db

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

type User struct {
	Username string `json:"username" binding:"required"`
	Points   int    `json:"points" binding:"required"`
	Rank     int    `json:"rank"`
}

/*
	Function : saveUser -> save the user data to the redis cacshe
*/

func (db *Database) SaveUser(user *User) error {
	member := &redis.Z{
		Score:  float64(user.Points),
		Member: user.Username,
	}
	pipe := db.Client.TxPipeline()
	pipe.ZAdd(Ctx, "leaderboard", member)
	rank := pipe.ZRank(Ctx, leaderboardKey, user.Username)
	_, err := pipe.Exec(Ctx)
	if err != nil {
		return err
	}
	fmt.Println(rank.Val(), err)
	user.Rank = int(rank.Val())
	return nil
}
