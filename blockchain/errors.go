package blockchain

import "fmt"

type AssetNotFoundError struct {
	AssetID string
}

func (e *AssetNotFoundError) Error() string {
	return fmt.Sprintf("asset with id '%s' not found", e.AssetID)
}

type AssetAlreadyExistsError struct {
	AssetID string
}

func (e *AssetAlreadyExistsError) Error() string {
	return fmt.Sprintf("asset with id '%s' already exists", e.AssetID)
}
