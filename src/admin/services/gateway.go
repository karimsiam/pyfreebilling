package services

import (
	"pks.pyfreebilling.com/models"
	"pks.pyfreebilling.com/navigation"
)

// GatewayService
var (
	GatewaysService gatewaysServiceInterface = &gatewaysService{}
)

type gatewaysService struct{}

type gatewaysServiceInterface interface {
	CreateGateway(models.Gateway) (*models.Gateway, error)
	ListGateways(pageStr string, gatewaysPerPage int) (*models.Gateways, *navigation.Pagination, error)
}

// SaveGateway saves the gateway object and returns the saved object
func (s *gatewaysService) CreateGateway(gateway models.Gateway) (*models.Gateway, error) {
	if err := models.CreateGateway(&gateway); err != nil {
		return nil, err
	}

	return &gateway, nil
}

// ListGateways returns a paginated list of gateways
func (s *gatewaysService) ListGateways(pageStr string, gatewaysPerPage int) (*models.Gateways, *navigation.Pagination, error) {
	// Count total og gatesways in DB
	gatewayCount, err := models.CountGateways()
	if err != nil {
		return nil, nil, err
	}

	// Get pagination informations
	p, paginateErr := navigation.Paginate(pageStr, int(gatewayCount), gatewaysPerPage)
	if paginateErr != nil {
		// c.AbortWithStatus(http.StatusBadRequest)
		return nil, nil, paginateErr
	}

	// Get gateways list
	var gateways models.Gateways
	if err := models.GetGateways(&gateways, gatewaysPerPage, p.Offset); err != nil {
		return nil, nil, err
	}

	return &gateways, p, nil
}
