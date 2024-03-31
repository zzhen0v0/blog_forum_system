package redis

import (
	"math"
	"time"

	"github.com/go-redis/redis"
)

const (
	OneWeekInSeconds         = 7 * 24 * 3600
	VoteScore        float64 = 432
	PostPerAge               = 20
)

/* PostVote 为帖子投票
投票分为四种情况：1.投赞成票 2.投反对票 3.取消投票 4.反转投票

记录文章参与投票的人
更新文章分数：赞成票要加分；反对票减分

v=1时，有两种情况
	1.之前没投过票，现在要投赞成票
	2.之前投过反对票，现在要改为赞成票
v=0时，有两种情况
	1.之前投过赞成票，现在要取消
	2.之前投过反对票，现在要取消
v=-1时，有两种情况
	1.之前没投过票，现在要投反对票
	2.之前投过赞成票，现在要改为反对票
*/

func PostVote(postID, userID string, v float64) (err error) {
	//check vote is ok
	postTime := client.ZScore(KeyPostTimeZSet, postID).Val()
	if float64(time.Now().Unix())-postTime > OneWeekInSeconds {
		err = ErrorVoteTimeExpire
		return
	}

	//update score
	ov := client.ZScore(KeyPostVotedZSetPrefix+postID, userID).Val()
	if ov == v {
		return ErrorVoted
	}

	var op float64
	if v > ov {
		op = 1
	} else {
		op = -1
	}
	diff := math.Abs(ov - v)

	pipeline := client.TxPipeline()
	pipeline.ZIncrBy(KeyPostScoreZSet, op*diff*VoteScore, postID)
	//update user record
	if v == 0 {
		pipeline.ZRem(KeyPostVotedZSetPrefix+postID, userID)
	} else {
		pipeline.ZAdd(KeyPostVotedZSetPrefix+postID, redis.Z{
			Score:  v,
			Member: userID,
		})
	}
	_, err = pipeline.Exec()
	return
}

func CreatePost(postID, communityID string) (err error) {

	pipeline := client.TxPipeline()
	pipeline.ZAdd(KeyPostTimeZSet, redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})

	pipeline.ZAdd(KeyPostScoreZSet, redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})
	ckey := KeyCommunityPostSetPrefix + communityID
	pipeline.SAdd(ckey, postID)
	_, err = pipeline.Exec()

	return
}
