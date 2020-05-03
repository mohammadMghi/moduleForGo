package contracts

import (
	"github.com/mohammadMghi/moduleForGo/entities"
	"github.com/mohammadMghi/moduleForGo/payloads"
)

type ISettingRepository interface {
	Set(setting entities.Setting) entities.Setting
	SetBatch(setting []entities.Setting) []entities.Setting
	Get(payload payloads.SettingGetPayload) entities.Setting
	List() []entities.Setting
}
