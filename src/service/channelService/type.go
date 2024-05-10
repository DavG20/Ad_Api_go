package channelService

import (
	"adtec/backend/src/service"
	"time"
)

type CreateInput struct {
	OwnerID          string     `json:"ownerId" validate:"required"`
	UserName         string     `json:"userName" validate:"required,min=4"`
	Name             *string    `json:"name,omitempty" validate:"omitempty,min=4"`
	Description      *string    `json:"description,omitempty" validate:"omitempty,min=10"`
	BotAddAsAdmin    *bool      `json:"botAddAsAdmin,omitempty"`
	ChannelCreatedAt *time.Time `json:"channelCreatedAt,omitempty" validate:"omitempty,ValidateTime"`
	Country          *string    `json:"country,omitempty" validate:"omitempty,iso3166_1_alpha3"` // Todo: validate country , I assume a three letter country code during validation
	Currency         *string    `json:"currency,omitempty" validate:"omitempty,iso4217"`
}
type UpdateInput struct {
	ID            string  `json:"id" validate:"required"`
	OwnerID       *string `json:"ownerId,omitempty" validate:"omitempty"`
	UserName      *string `json:"userName,omitempty" validate:"omitempty,min=4"`
	Name          *string `json:"name,omitempty" validate:"omitempty,min=4"`
	Description   *string `json:"description,omitempty" validate:"omitempty,min=10"`
	BotAddAsAdmin *bool   `json:"botAddAsAdmin,omitempty"`
	Country       *string `json:"country,omitempty" validate:"omitempty,iso3166_1_alpha3"` // TODO: validate country , I assume a three letter country code during validation
	Currency      *string `json:"currency,omitempty" validate:"omitempty,iso4217"`
}
type ChannelFilter struct {
	Name           *string                    `json:"name,omitempty"`
	Country        *string                    `json:"country,omitempty"`
	Categories     []string                   `json:"categories,omitempty"`
	Languages      []string                   `json:"languages,omitempty"`
	SubCount       *service.MinMaxFilterInput `json:"subCount,omitempty"`
	CpmFilter      *service.MinMaxFilterInput `json:"cpmFilter,omitempty"`
	PostViewFilter *service.MinMaxFilterInput `json:"postViewFilter,omitempty"`
	Filter         *service.FilterData        `json:"filter,omitempty"`
}
