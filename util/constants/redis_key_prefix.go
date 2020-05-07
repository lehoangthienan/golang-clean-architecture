package constants

type RedisKeyPrefix string

func (rkp RedisKeyPrefix) String() string {
	return string(rkp)
}

const (
	CampaignTypeRedisKey RedisKeyPrefix = "campaign_type_"
)
