package contracts

import "github.com/mohammadMghi/moduleForGo/payloads"

type IApplicationService interface {
	Info() (payloads.ApplicationInfoPayload, error)
	Ping() (payloads.ApplicationPingPayload, error)
}
