// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package handler

import (
	"gitlab.com/shaninalex/flowreon/apps/org/domain"
)

type OrganizationHandler struct {
	api *domain.OrganizationApi
}

func NewOrganizationHandler() *OrganizationHandler {
	return &OrganizationHandler{
		api: domain.NewOrganizationApi(),
	}
}
