package model

var (
	DefaultName   = make(map[string]string)
	CC            = "custom_config"                                          // 自行配置的key
	CN            = "用户自定义默认key"                                             // 自定义key备注note
	TXSECRETID    = "tx_secretid"                                            // 腾讯SECRETID，账号需要开启[QCloudFinanceFullAccess、QcloudRedisFullAccess、QcloudMonitorFullAccess]权限
	TXSECRETKEY   = "tx_secretkey"                                           // 腾讯SECRETKEY
	TXAPIURL      = "tx_apiurl"                                              // 腾讯APIURL
	TXREGION      = "tx_region"                                              // 腾讯REGION
	BGSAVECOMMAND = "redis_bgsave"                                           //bgsave命令的别名
	CfgDefault    = [...]string{TXSECRETID, TXSECRETKEY, TXAPIURL, TXREGION} // 默认key列表
)

func init() {
	// DefaultName := make(map[string]string)
	DefaultName[TXSECRETID] = "腾讯SECRETID"
	DefaultName[TXSECRETKEY] = "腾讯TXSECRETKEY"
	DefaultName[TXAPIURL] = "腾讯APIURL"
	DefaultName[BGSAVECOMMAND] = "bgsave别名"
}
