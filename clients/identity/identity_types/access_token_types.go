package identity_types

type AccessTokenType string

const (
	AccessTokenTypeUser    AccessTokenType = "userToken"
	AccessTokenTypeTenant  AccessTokenType = "tenantApiToken"
	AccessTokenTypeUserApi AccessTokenType = "userApiToken"
)
