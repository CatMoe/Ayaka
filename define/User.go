package define

import "github.com/google/uuid"

// 等待考虑是否只使用OAuth接入
type User struct {
	UserName string    `json:"user_name"` // 昵称
	ID       uuid.UUID `json:"uuid"`      // UUID
	WeChatID string    `json:"wechat_id,omitempty"`
	Password string    `json:"password"` //应当为hash后结果

}
