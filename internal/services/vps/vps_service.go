package vps

import (
	"context"
	"vball/internal/models"
	vps "vball/internal/repositories/vps"
)

func GetRegions_Service() ([]models.Region, error) {
	regions, err := vps.GetAllRegions_Repo(context.Background())
	if err != nil {
		return nil, err
	}
	return regions, nil
}
