package Mode

var REDIS_DOMAINLISK_KEY="domainList"

var REDIS_EXECMQ="domainpro"

var M10 int64=10*60
var M30 int64=30*60
var H1 int64=1*60*60
var H24 int64=24*60*60

type TaskInfo struct {
	DomainName string `json:"domain_name" binding:"required"`
	CheckRate  string `json:"check_rate" binding:"required"`
	Platform string `json:"platform" binding:"required"`
	DomainType string `json:"domain_type" binding:"required"`
	FailRate  float64 `json:"fail_rate" binding:"required"`
	AlertsType string `json:"alerts_type" binding:"required"`
	LastExecutedEvent int64 `json:"last_executed_event"`
}
