package v201809

import (
	"encoding/xml"
	"log"
)

//	"encoding/xml"
//	"fmt"

type BudgetOrderService struct {
	Auth
}

func NewBudgetOrderService(auth *Auth) *BudgetOrderService {
	return &BudgetOrderService{Auth: *auth}
}

type Field struct {
	XMLName xml.Name
	Value   string `xml:",innerxml"`
}

type BillingSelector struct {
	XMLName    xml.Name
	Fields     []Field     `xml:"fields,omitempty"`
	Predicates []Predicate `xml:"predicates"`
	DateRange  *DateRange  `xml:"dateRange,omitempty"`
	Ordering   []OrderBy   `xml:"ordering"`
	Paging     *Paging     `xml:"paging,omitempty"`
}

// A Budget represents an allotment of money to be spent over a fixed
// period of time.
type BudgetOrder struct {
	Id         int64  `xml:"budgetId,omitempty"`           // A unique identifier
	Name       string `xml:"name"`                         // A descriptive name
	Period     string `xml:"period,omitempty"`             // The period to spend the budget
	Amount     int64  `xml:"totalAdjustments>microAmount"` // The amount in cents
	Delivery   string `xml:"deliveryMethod,omitempty"`     // The rate at which the budget spent. valid options are STANDARD or ACCELERATED.
	References int64  `xml:"referenceCount,omitempty"`     // The number of campaigns using the budget
	Shared     bool   `xml:"isExplicitlyShared,omitempty"` // If this budget was created to be shared across campaigns
	Status     string `xml:"status,omitempty"`             // The status of the budget. can be ENABLED, REMOVED, UNKNOWN
}

type BillingAccount struct {
}

func (s *BudgetOrderService) Get(
	selector BillingSelector,
) (budgetOrders []BudgetOrder, totalCount int64, err error) {
	selector.XMLName = xml.Name{baseBillingUrl, "serviceSelector"}
	if len(selector.Fields) > 0 {
		for i := range selector.Fields {
			selector.Fields[i].XMLName = xml.Name{baseUrl, "fields"}
		}
	}
	respBody, err := s.Auth.request(
		budgetOrderServiceUrl,
		"get",
		struct {
			XMLName xml.Name
			Sel     BillingSelector
		}{
			XMLName: xml.Name{
				Space: baseBillingUrl,
				Local: "get",
			},
			Sel: selector,
		},
	)
	if err != nil {
		return budgetOrders, totalCount, err
	}
	getResp := struct {
		Size         int64         `xml:"rval>totalNumEntries"`
		BudgetOrders []BudgetOrder `xml:"rval>entries"`
	}{}
	log.Println(string(respBody))
	err = xml.Unmarshal([]byte(respBody), &getResp)
	if err != nil {
		return budgetOrders, totalCount, err
	}
	return getResp.BudgetOrders, getResp.Size, err
}

func (s *BudgetOrderService) GetBillingAccounts() (billingAccounts []BillingAccount, err error) {
	respBody, err := s.Auth.request(
		budgetOrderServiceUrl,
		"getBillingAccounts",
		struct {
			XMLName xml.Name
		}{
			XMLName: xml.Name{
				Space: baseBillingUrl,
				Local: "getBillingAccounts",
			},
		},
	)
	if err != nil {
		return billingAccounts, err
	}
	getResp := struct {
		BillingAccounts []BillingAccount `xml:"rval"`
	}{}
	err = xml.Unmarshal([]byte(respBody), &getResp)
	if err != nil {
		return billingAccounts, err
	}
	return getResp.BillingAccounts, nil
}
