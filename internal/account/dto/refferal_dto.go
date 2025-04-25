package dto

import "github.com/evelinix/nusaloka/internal/account/model"

type ReferalCodeResponse struct {
	Code string `json:"code"`
}

type ReferalListResponse struct {
	Referals []model.Referal `json:"referals"`
}