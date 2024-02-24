package redis
// https://github.com/go-redis/redismock

// import (
// 	"errors"
// 	"fmt"
// 	"testing"
// 	"time"

// 	red "github.com/go-redis/redis"
// 	"github.com/go-redis/redismock/v9"
// )

// func NewsInfoForCache(redisDB *Client, newsID int) (info string, err error) {
// 	cacheKey := fmt.Sprintf("news_redis_cache_%d", newsID)
// 	info, err = redisDB.Get(ctx, cacheKey).Result()
// 	if err == red.Nil {
// 		// info, err = call api()
// 		info = "test"
// 		err = redisDB.Set(ctx, cacheKey, info, 30*time.Minute).Err()
// 	}
// 	return
// }
// func TestNewsInfoForCache(t *testing.T) {
// 	db, mock := redismock.NewClientMock()
// 	newsID := 123456789
// 	key := fmt.Sprintf("news_redis_cache_%d", newsID)

// 	// mock ignoring `call api()`

// 	mock.ExpectGet(key).RedisNil()
// 	mock.Regexp().ExpectSet(key, `[a-z]+`, 30*time.Minute).SetErr(errors.New("FAIL"))

// 	_, err := NewsInfoForCache(db, newsID)
// 	if err == nil || err.Error() != "FAIL" {
// 		t.Error("wrong error")
// 	}

// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Error(err)
// 	}
// }
