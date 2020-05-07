package constants

import (
	"time"

	"github.com/lehoangthienan/marvel-heroes-backend/util/errors"
)

// Options type
type Options map[string]interface{}

// default map values
const (
	DefaultPage           int    = 1
	DefaultAmountPerPage  int    = 10
	DefaultOrder          string = "desc"
	DefaultOrderedBy      string = "updated_at"
	DefaultFromHours      int    = -1
	DefaultToHours        int    = -1
	DefaultCampaignTypeID string = ""
)

// Set set value at key
func (o Options) Set(key string, value interface{}) {
	o[key] = value
}

// Page returns key="page" value
func (o Options) Page() int {
	if p, has := o["page"]; has {
		if page, ok := p.(int); ok {
			return page
		}
		return DefaultPage
	}
	return DefaultPage
}

// AmountPerPage returns key="amountPerPage" value
func (o Options) AmountPerPage() int {
	if app, has := o["amountPerPage"]; has {
		if amountPerPage, ok := app.(int); ok {
			return amountPerPage
		}
		return DefaultAmountPerPage
	}
	return DefaultAmountPerPage
}

// Order returns key="order" value
func (o Options) Order() string {
	if od, has := o["order"]; has {
		if order, ok := od.(string); ok {
			if order == "" {
				return DefaultOrder
			}
			return order
		}
		return DefaultOrder
	}
	return DefaultOrder
}

// OrderedBy returns key="orderedBy" value
func (o Options) OrderedBy() string {
	if ob, has := o["orderedBy"]; has {
		if orderedBy, ok := ob.(string); ok {
			return orderedBy
		}
		return DefaultOrderedBy
	}
	return DefaultOrderedBy
}

// FromHours returns key="fromHours" value
func (o Options) FromHours() int {
	if fh, has := o["fromHours"]; has {
		if fromHours, ok := fh.(int); ok {
			return fromHours
		}
		return DefaultFromHours
	}
	return DefaultFromHours
}

// ToHours returns key="toHours" value
func (o Options) ToHours() int {
	if th, has := o["toHours"]; has {
		if toHours, ok := th.(int); ok {
			return toHours
		}
		return DefaultToHours
	}
	return DefaultToHours
}

// CampaignTypeID returns key="campaignTypeID" value
func (o Options) CampaignTypeID() string {
	if ctID, has := o["campaignTypeID"]; has {
		if campaignTypeID, ok := ctID.(string); ok {
			return campaignTypeID
		}
		return DefaultCampaignTypeID
	}
	return DefaultCampaignTypeID
}

// FromTime returns key="fromTime" value
func (o Options) FromTime() (*time.Time, error) {
	if ft, has := o["fromTime"]; has {
		if fromTime, ok := ft.(*time.Time); ok {
			return fromTime, nil
		}
		return nil, errors.TypeAssertionError
	}
	return nil, nil
}

// ToTime returns key="toTime" value
func (o Options) ToTime() (*time.Time, error) {
	if tt, has := o["toTime"]; has {
		if toTime, ok := tt.(*time.Time); ok {
			return toTime, nil
		}
		return nil, errors.TypeAssertionError
	}
	return nil, nil
}

// NowAvailible func
func (o Options) NowAvailible() (string, error) {
	if na, has := o["nowAvailible"].(string); has {
		return na, nil
	}
	return "", nil
}
