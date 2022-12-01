package enum

// ProjectStatus 广告项目状态
type ProjectStatus string

const (
	// PROJECT_STATUS_ENABLE 启用
	PROJECT_STATUS_ENABLE ProjectStatus = "PROJECT_STATUS_ENABLE"
	// PROJECT_STATUS_DISABLE 暂停
	PROJECT_STATUS_DISABLE ProjectStatus = "PROJECT_STATUS_DISABLE"
	// PROJECT_STATUS_DELETE 删除
	PROJECT_STATUS_DELETE ProjectStatus = "PROJECT_STATUS_DELETE"
	// PROJECT_STATUS_ALL 所有包含已删除
	PROJECT_STATUS_ALL ProjectStatus = "PROJECT_STATUS_ALL"
	// PROJECT_STATUS_NOT_DELETE 所有不包含已删除
	PROJECT_STATUS_NOT_DELETE ProjectStatus = "PROJECT_STATUS_NOT_DELETE"
	// PROJECT_STATUS_BUDGET_EXCEED 项目超出预算
	PROJECT_STATUS_BUDGET_EXCEED ProjectStatus = "PROJECT_STATUS_BUDGET_EXCEED"
	// PROJECT_STATUS_BUDGET_PRE_OFFLINE_BUDGET 项目接近预算
	PROJECT_STATUS_BUDGET_PRE_OFFLINE_BUDGET ProjectStatus = "PROJECT_STATUS_BUDGET_PRE_OFFLINE_BUDGET"
	// PROJECT_STATUS_NOT_START 未达投放时间
	PROJECT_STATUS_NOT_START ProjectStatus = "PROJECT_STATUS_NOT_START"
	// PROJECT_STATUS_DONE 已完成
	PROJECT_STATUS_DONE ProjectStatus = "PROJECT_STATUS_DONE"
	// PROJECT_STATUS_NO_SCHEDULE 不在投放时段
	PROJECT_STATUS_NO_SCHEDULE ProjectStatus = "PROJECT_STATUS_NO_SCHEDULE"
)
