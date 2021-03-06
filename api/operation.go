package api

import (
	"fmt"
)

// Operation custom type for PA API operations
type Operation string

// ErrInvalidResource is thrown when a resource is not compatible with an operation
type ErrInvalidResource struct {
	resource  Resource
	operation Operation
}

func (e ErrInvalidResource) Error() string {
	return fmt.Sprintf(`Invalid resource "%s" for operation "%s"`, e.resource, e.operation)
}

const (
	// GetBrowseNodes constant for GetBrowseNodes operation
	GetBrowseNodes Operation = "GetBrowseNodes"
	// GetItems constant for GetItems operation
	GetItems Operation = "GetItems"
	// GetVariations constant for GetVariations operation
	GetVariations Operation = "GetVariations"
	// SearchItems constant for SearchItems operation
	SearchItems Operation = "SearchItems"
)

// resourceOperationsMap maps resources to their valid operations
var resourceOperationsMap = map[Resource][]Operation{
	BrowseNodesAncestor:                                    []Operation{GetBrowseNodes},
	BrowseNodesChildren:                                    []Operation{GetBrowseNodes},
	BrowseNodeInfoBrowseNodes:                              []Operation{GetItems, GetVariations, SearchItems},
	BrowseNodeInfoBrowseNodesAncestor:                      []Operation{GetItems, GetVariations, SearchItems},
	BrowseNodeInfoBrowseNodesSalesRank:                     []Operation{GetItems, GetVariations, SearchItems},
	BrowseNodeInfoWebsiteSalesRank:                         []Operation{GetItems, GetVariations, SearchItems},
	CustomerReviewsCount:                                   []Operation{GetItems, GetVariations, SearchItems},
	CustomerReviewsStarRating:                              []Operation{GetItems, GetVariations, SearchItems},
	ImagesPrimarySmall:                                     []Operation{GetItems, GetVariations, SearchItems},
	ImagesPrimaryMedium:                                    []Operation{GetItems, GetVariations, SearchItems},
	ImagesPrimaryLarge:                                     []Operation{GetItems, GetVariations, SearchItems},
	ImagesVariantsSmall:                                    []Operation{GetItems, GetVariations, SearchItems},
	ImagesVariantsMedium:                                   []Operation{GetItems, GetVariations, SearchItems},
	ImagesVariantsLarge:                                    []Operation{GetItems, GetVariations, SearchItems},
	ItemInfoByLineInfo:                                     []Operation{GetItems, GetVariations, SearchItems},
	ItemInfoContentInfo:                                    []Operation{GetItems, GetVariations, SearchItems},
	ItemInfoContentRating:                                  []Operation{GetItems, GetVariations, SearchItems},
	ItemInfoClassifications:                                []Operation{GetItems, GetVariations, SearchItems},
	ItemInfoExternalIds:                                    []Operation{GetItems, GetVariations, SearchItems},
	ItemInfoFeatures:                                       []Operation{GetItems, GetVariations, SearchItems},
	ItemInfoManufactureInfo:                                []Operation{GetItems, GetVariations, SearchItems},
	ItemInfoProductInfo:                                    []Operation{GetItems, GetVariations, SearchItems},
	ItemInfoTechnicalInfo:                                  []Operation{GetItems, GetVariations, SearchItems},
	ItemInfoTitle:                                          []Operation{GetItems, GetVariations, SearchItems},
	ItemInfoTradeInInfo:                                    []Operation{GetItems, GetVariations, SearchItems},
	OffersListingsAvailabilityMaxOrderQuantity:             []Operation{GetItems, GetVariations, SearchItems},
	OffersListingsAvailabilityMessage:                      []Operation{GetItems, GetVariations, SearchItems},
	OffersListingsAvailabilityMinOrderQuantity:             []Operation{GetItems, GetVariations, SearchItems},
	OffersListingsAvailabilityType:                         []Operation{GetItems, GetVariations, SearchItems},
	OffersListingsCondition:                                []Operation{GetItems, GetVariations, SearchItems},
	OffersListingsConditionSubCondition:                    []Operation{GetItems, GetVariations, SearchItems},
	OffersListingsDeliveryInfoIsAmazonFulfilled:            []Operation{GetItems, GetVariations, SearchItems},
	OffersListingsDeliveryInfoIsFreeShippingEligible:       []Operation{GetItems, GetVariations, SearchItems},
	OffersListingsDeliveryInfoIsPrimeEligible:              []Operation{GetItems, GetVariations, SearchItems},
	OffersListingsDeliveryInfoShippingCharges:              []Operation{GetItems, GetVariations, SearchItems},
	OffersListingsIsBuyBoxWinner:                           []Operation{GetItems, GetVariations, SearchItems},
	OffersListingsLoyaltyPointsPoints:                      []Operation{GetItems, GetVariations, SearchItems},
	OffersListingsMerchantInfo:                             []Operation{GetItems, GetVariations, SearchItems},
	OffersListingsPrice:                                    []Operation{GetItems, GetVariations, SearchItems},
	OffersListingsProgramEligibilityIsPrimeExclusive:       []Operation{GetItems, GetVariations, SearchItems},
	OffersListingsProgramEligibilityIsPrimePantry:          []Operation{GetItems, GetVariations, SearchItems},
	OffersListingsPromotions:                               []Operation{GetItems, GetVariations, SearchItems},
	OffersListingsSavingBasis:                              []Operation{GetItems, GetVariations, SearchItems},
	OffersSummariesHighestPrice:                            []Operation{GetItems, GetVariations, SearchItems},
	OffersSummariesLowestPrice:                             []Operation{GetItems, GetVariations, SearchItems},
	OffersSummariesOfferCount:                              []Operation{GetItems, GetVariations, SearchItems},
	ParentASIN:                                             []Operation{GetItems, GetVariations, SearchItems},
	RentalOffersListingsAvailabilityMaxOrderQuantity:       []Operation{GetItems, GetVariations, SearchItems},
	RentalOffersListingsAvailabilityMessage:                []Operation{GetItems, GetVariations, SearchItems},
	RentalOffersListingsAvailabilityMinOrderQuantity:       []Operation{GetItems, GetVariations, SearchItems},
	RentalOffersListingsAvailabilityType:                   []Operation{GetItems, GetVariations, SearchItems},
	RentalOffersListingsBasePrice:                          []Operation{GetItems, GetVariations, SearchItems},
	RentalOffersListingsCondition:                          []Operation{GetItems, GetVariations, SearchItems},
	RentalOffersListingsConditionSubCondition:              []Operation{GetItems, GetVariations, SearchItems},
	RentalOffersListingsDeliveryInfoIsAmazonFulfilled:      []Operation{GetItems, GetVariations, SearchItems},
	RentalOffersListingsDeliveryInfoIsFreeShippingEligible: []Operation{GetItems, GetVariations, SearchItems},
	RentalOffersListingsDeliveryInfoIsPrimeEligible:        []Operation{GetItems, GetVariations, SearchItems},
	RentalOffersListingsDeliveryInfoShippingCharges:        []Operation{GetItems, GetVariations, SearchItems},
	RentalOffersListingsMerchantInfo:                       []Operation{GetItems, GetVariations, SearchItems},
	VariationSummaryPriceHighestPrice:                      []Operation{GetVariations},
	VariationSummaryPriceLowestPrice:                       []Operation{GetVariations},
	VariationSummaryVariationDimension:                     []Operation{GetVariations},
	SearchRefinements:                                      []Operation{SearchItems},
}

// Validate validates resources for an operation
func (o Operation) Validate(resources []Resource) error {
	for _, r := range resources {
		validOperations := resourceOperationsMap[r]
		if !existsInOperations(o, validOperations) {
			return ErrInvalidResource{r, o}
		}
	}

	return nil
}

// existsInOperations checks if an Operation exists in a slice of Operations
func existsInOperations(operation Operation, operations []Operation) bool {
	for _, o := range operations {
		if o == operation {
			return true
		}
	}

	return false
}
