package contracts

import "github.com/mohammadMghi/moduleForGo/payloads"

type IApplicationRepository interface {
	Info() (payloads.ApplicationInfoPayload, error)
	Ping() (payloads.ApplicationPingPayload, error)
}
