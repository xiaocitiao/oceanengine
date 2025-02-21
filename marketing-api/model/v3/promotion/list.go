package promotion

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// ListRequest 获取广告列表 API Request
type ListRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Filtering 过滤条件
	Filtering *ListFilter `json:"filtering,omitempty"`
	// Fields 查询字段集合，如果指定则应答结果仅返回指定字段
	// 可参考应答参数返回的指标字段
	Fields []string `json:"fields,omitempty"`
	// Page 页数默认值: 1，page必须大于0
	Page int `json:"page,omitempty"`
	// PageSize 页面大小默认值:10，page_size范围为1-10
	PageSize int `json:"page_size,omitempty"`
}

// ListFilter 过滤条件
type ListFilter struct {
	// IDs 按广告ID过滤，范围为1-100
	IDs []uint64 `json:"ids,omitempty"`
	// Name 广告名称，长度是1-50个字（两个英文字符占1个字，该字段采取模糊查询的方式）
	Name string `json:"name,omitempty"`
	// Status 广告状态过滤，允许值：
	// NOT_DELETED 不限 、ALL 不限（包含已删除）、OK 投放中、DELETED 已删除、PROJECT_OFFLINE_BUDGET 项目超出预算、PROJECT_PREOFFLINE_BUDGET 项目接近预算、TIME_NO_REACH 未到达投放时间、TIME_DONE 已完成、NO_SCHEDULE 不在投放时段、AUDIT 新建审核中、REAUDIT 修改审核中、FROZEN 已终止、AUDIT_DENY 审核不通过、OFFLINE_BUDGET 广告超出预算、OFFLINE_BALANCE 账户余额不足、PREOFFLINE_BUDGET 广告接近预算、DISABLED 已暂停、PROJECT_DISABLED 已被项目暂停、LIVE_ROOM_OFF 关联直播间不可投、PRODUCT_OFFLINE 关联商品不可投，、AWEME_ACCOUNT_DISABLED 关联抖音账号不可投、AWEME_ANCHOR_DISABLED 锚点不可投、DISABLE_BY_QUOTA 已暂停（配额达限）
	Status enum.PromotionStatus `json:"status,omitempty"`
	// PromotionCreateTime 广告创建时间，格式yyyy-mm-dd，表示过滤出当天创建的广告项目
	PromotionCreateTime string `json:"promotion_create_time,omitempty"`
	// PromotionModifyTime 广告更新时间，格式yyyy-mm-dd，表示过滤出当天更新的广告项目
	PromotionModifyTime string `json:"promotion_modify_time,omitempty"`
	// RejectReasonType 审核建议类型，允许值：NONE 无建议、REVIEW_REJECT 审核不通过、LOW_MATERAIL 低质素材、DISCOMFORT 引人不适、QUALITY_POOR 素材质量低、EXAGGERATION 夸大宣传、ELSE 其他
	RejectReasonType enum.PromotionRejectReasonType `json:"reject_reason_type,omitempty"`
	// LearningPhase 学习期状态 允许值：LEARNING（学习期中）、LEARNED（学习期结束）、LEARN_FAILED（学习期失败)
	LearningPhase []enum.LearningPhase `json:"learning_phase,omitempty"`
}

func (f ListFilter) GetIDs() []uint64 {
	return f.IDs
}
func (f ListFilter) GetName() string {
	return f.Name
}
func (f ListFilter) GetCreateTime() string {
	return f.PromotionCreateTime
}
func (f ListFilter) GetModifyTime() string {
	return f.PromotionModifyTime
}

// Encode implement GetRequest interface
func (r ListRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.Filtering != nil {
		filtering, _ := json.Marshal(r.Filtering)
		values.Set("filtering", string(filtering))
	}
	if len(r.Fields) > 0 {
		fields, _ := json.Marshal(r.Fields)
		values.Set("fields", string(fields))
	}
	if r.Page > 0 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// ListResponse 获取广告列表 API Response
type ListResponse struct {
	model.BaseResponse
	Data *ListResponseData `json:"data,omitempty"`
}

type ListResponseData struct {
	// List 项目列表
	List []Promotion `json:"list,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}
