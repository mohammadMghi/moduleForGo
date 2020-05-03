package contracts

import (
	"github.com/mohammadMghi/moduleForGo/entities"
	"github.com/mohammadMghi/moduleForGo/payloads"
)

type ISettingService interface {
	Set(payload payloads.SettingSetPayload) entities.Setting
	Get(payload payloads.SettingGetPayload) entities.Setting
}
