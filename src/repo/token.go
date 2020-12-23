package repo

import (
	"fmt"
	redisUtils "github.com/aaronchen2k/openstc/src/libs/redis"
	"github.com/aaronchen2k/openstc/src/models"
	"github.com/gomodule/redigo/redis"
	"gorm.io/gorm"
	"strings"
	"time"
)

type TokenRepo struct {
	CommonRepo
	DB *gorm.DB `inject:""`
}

func NewTokenRepo() *TokenRepo {
	return &TokenRepo{}
}

func (r *TokenRepo) GetRedisSessionV2(conn *redisUtils.RedisCluster, token string) (*models.RedisSessionV2, error) {
	sKey := models.ZXW_SESSION_TOKEN_PREFIX + token
	if !conn.Exists(sKey) {
		return nil, models.ERR_TOKEN_INVALID
	}
	pp := new(models.RedisSessionV2)
	if err := r.loadRedisHashToStruct(conn, sKey, pp); err != nil {
		return nil, err
	}
	return pp, nil
}

func (r *TokenRepo) loadRedisHashToStruct(conn *redisUtils.RedisCluster, sKey string, pst interface{}) error {
	vals, err := redis.Values(conn.HGetAll(sKey))
	if err != nil {
		return err
	}
	err = redis.ScanStruct(vals, pst)
	if err != nil {
		return err
	}
	return nil
}

func (r *TokenRepo) IsUserTokenOver(userId string) bool {
	conn := redisUtils.GetRedisClusterClient()
	defer conn.Close()
	if r.getUserTokenCount(conn, userId) >= r.getUserTokenMaxCount(conn) {
		return true
	}
	return false
}

func (r *TokenRepo) getUserTokenCount(conn *redisUtils.RedisCluster, userId string) int {
	count, err := redis.Int(conn.Scard(models.ZXW_SESSION_USER_PREFIX + userId))
	if err != nil {
		fmt.Println(fmt.Sprintf("getUserTokenCount error :%+v", err))
		return 0
	}
	return count
}

func (r *TokenRepo) getUserTokenMaxCount(conn *redisUtils.RedisCluster) int {
	count, err := redis.Int(conn.GetKey(models.ZXW_SESSION_USER_MAX_TOKEN_PREFIX))
	if err != nil {
		return models.ZXW_SESSION_USER_MAX_TOKEN_DEFAULT
	}
	return count
}

func (r *TokenRepo) UserTokenExpired(token string) {
	conn := redisUtils.GetRedisClusterClient()
	defer conn.Close()

	uKey := models.ZXW_SESSION_BIND_USER_PREFIX + token
	sKeys, err := redis.Strings(conn.Members(uKey))
	if err != nil {
		fmt.Println(fmt.Sprintf("conn.Members key %s error :%+v", uKey, err))
		return
	}
	for _, v := range sKeys {
		if !strings.Contains(v, models.ZXW_SESSION_USER_PREFIX) {
			continue
		}
		_, err := conn.Do("SREM", v, token)
		if err != nil {
			fmt.Println(fmt.Sprintf("conn.SREM key %s token %s  error :%+v", v, token, err))
			return
		}
	}
	if _, err := conn.Del(uKey); err != nil {
		fmt.Println(fmt.Sprintf("conn.Del key %s error :%+v", uKey, err))
	}
	return
}

func (r *TokenRepo) GetUserScope(userType string) uint64 {
	switch userType {
	case "admin":
		return models.AdminScope
	}
	return models.NoneScope
}

func (r *TokenRepo) ToCache(conn *redisUtils.RedisCluster, rs models.RedisSessionV2, token string) error {
	sKey := models.ZXW_SESSION_TOKEN_PREFIX + token

	if _, err := conn.HMSet(sKey,
		"user_id", rs.UserId,
		"login_type", rs.LoginType,
		"auth_type", rs.AuthType,
		"creation_data", rs.CreationDate,
		"expires_in", rs.ExpiresIn,
		"scope", rs.Scope,
	); err != nil {
		fmt.Println(fmt.Sprintf("conn.ToCache error :%+v", err))
		return err
	}
	return nil
}

func (r *TokenRepo) SyncUserTokenCache(conn *redisUtils.RedisCluster, rs models.RedisSessionV2, token string) error {
	sKey := models.ZXW_SESSION_USER_PREFIX + token
	if _, err := conn.Sadd(sKey, token); err != nil {
		fmt.Println(fmt.Sprintf("conn.SyncUserTokenCache1 error :%+v", err))
		return err
	}
	sKey2 := models.ZXW_SESSION_BIND_USER_PREFIX + token
	_, err := conn.Sadd(sKey2, sKey)
	if err != nil {
		fmt.Println(fmt.Sprintf("conn.SyncUserTokenCache2 error :%+v", err))
		return err
	}
	return nil
}

func (r *TokenRepo) UpdateUserTokenCacheExpire(conn *redisUtils.RedisCluster, rs models.RedisSessionV2, token string) error {
	if _, err := conn.Expire(models.ZXW_SESSION_TOKEN_PREFIX+token, int(r.GetTokenExpire(rs).Seconds())); err != nil {
		fmt.Println(fmt.Sprintf("conn.UpdateUserTokenCacheExpire error :%+v", err))
		return err
	}
	return nil
}

func (r *TokenRepo) GetTokenExpire(rs models.RedisSessionV2) time.Duration {
	timeout := models.RedisSessionTimeoutApp
	if rs.LoginType == models.LoginTypeWeb {
		timeout = models.RedisSessionTimeoutWeb
	} else if rs.LoginType == models.LoginTypeWx {
		timeout = models.RedisSessionTimeoutWx
	} else if rs.LoginType == models.LoginTypeAlipay {
		timeout = models.RedisSessionTimeoutWx
	}
	return timeout
}

func (r *TokenRepo) DelUserTokenCache(conn *redisUtils.RedisCluster, rs models.RedisSessionV2, token string) error {
	sKey := models.ZXW_SESSION_USER_PREFIX + rs.UserId
	_, err := conn.Do("SREM", sKey, token)
	if err != nil {
		fmt.Println(fmt.Sprintf("conn.DelUserTokenCache1 error :%+v", err))
		return err
	}
	err = r.DelTokenCache(conn, rs, token)
	if err != nil {
		fmt.Println(fmt.Sprintf("conn.DelUserTokenCache2 error :%+v", err))
		return err
	}

	return nil
}
func (r *TokenRepo) DelTokenCache(conn *redisUtils.RedisCluster, rs models.RedisSessionV2, token string) error {
	sKey2 := models.ZXW_SESSION_BIND_USER_PREFIX + token
	_, err := conn.Del(sKey2)
	if err != nil {
		fmt.Println(fmt.Sprintf("conn.DelUserTokenCache2 error :%+v", err))
		return err
	}
	sKey3 := models.ZXW_SESSION_TOKEN_PREFIX + token
	_, err = conn.Del(sKey3)
	if err != nil {
		fmt.Println(fmt.Sprintf("conn.DelUserTokenCache3 error :%+v", err))
		return err
	}

	return nil
}

func (r *TokenRepo) CleanUserTokenCache(conn *redisUtils.RedisCluster, rs models.RedisSessionV2) error {
	sKey := models.ZXW_SESSION_USER_PREFIX + rs.UserId
	allTokens, err := redis.Strings(conn.Members(sKey))
	if err != nil {
		fmt.Println(fmt.Sprintf("conn.CleanUserTokenCache1 error :%+v", err))
		return err
	}
	_, err = conn.Del(sKey)
	if err != nil {
		fmt.Println(fmt.Sprintf("conn.CleanUserTokenCache2 error :%+v", err))
		return err
	}

	for _, token := range allTokens {
		err = r.DelTokenCache(conn, rs, token)
		if err != nil {
			fmt.Println(fmt.Sprintf("conn.DelUserTokenCache2 error :%+v", err))
			return err
		}
	}
	return nil
}
