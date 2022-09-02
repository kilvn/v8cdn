// Package enum 
package enum

type EnumUserAuthType string

const (
    EnumUserAuthTypeApiKey EnumUserAuthType = "user_api_key"

    EnumUserAuthTypePartner EnumUserAuthType = "partner"
)
