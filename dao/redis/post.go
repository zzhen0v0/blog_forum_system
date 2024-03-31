package redis

import (
	"strconv"
	"time"
	"web_app/models"

	"github.com/go-redis/redis"
)

func GetPostList(p *models.ParamPost) ([]string, error) {
	var key string = KeyPostTimeZSet
	if p.Order == models.OrderScore {
		key = KeyPostScoreZSet
	}
	start := (p.Page - 1) * p.Size
	end := start + p.Size - 1

	return client.ZRevRange(key, start, end).Result()
}

func GetPostListCommunityID(p *models.ParamPostCommunity) (ids []string, err error) {
	ckey := KeyCommunityPostSetPrefix + strconv.Itoa(int(p.CommunityID))
	var orderKey string = KeyPostTimeZSet
	if p.Order == models.OrderScore {
		orderKey = KeyPostScoreZSet
	}
	key := orderKey + strconv.Itoa(int(p.CommunityID))
	if client.Exists(key).Val() < 1 {
		pipeline := client.TxPipeline()
		pipeline.ZInterStore(key, redis.ZStore{
			Aggregate: "max",
		}, ckey, orderKey)
		pipeline.Expire(key, 60*time.Second)
		_, err := pipeline.Exec()
		if err != nil {
			return nil, err
		}
	}
	start := (p.Page - 1) * p.Size
	end := start + p.Size - 1
	return client.ZRevRange(key, start, end).Result()
}

func GetVoteData(ids []string) (data []int64, err error) {

	pipeline := client.TxPipeline()
	for _, id := range ids {
		pipeline.ZCount(KeyPostVotedZSetPrefix+id, "1", "1")
	}
	cmders, err := pipeline.Exec()
	if err != nil {
		return
	}
	data = make([]int64, 0, len(ids))
	for _, cmder := range cmders {
		v := cmder.(*redis.IntCmd).Val()
		data = append(data, v)
	}
	return
}
