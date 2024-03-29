// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package ndc_gen

import (
	"fmt"
	"io"
	"strconv"
)

type Node interface {
	IsNode()
}

type AdviseMessage struct {
	Code          string           `json:"code"`
	Description   string           `json:"description"`
	Level         int              `json:"level"`
	External      *ExternalMessage `json:"external"`
	CorrelationID string           `json:"correlationID"`
}

type AirShoppingRQBasic struct {
	OriginDestination []*OriginDestination `json:"originDestination"`
	Pax               []*IPax              `json:"pax"`
	Carrier           []*string            `json:"carrier"`
}

type AirShoppingRs struct {
	Errors   []*Error   `json:"errors"`
	Response *ResponseX `json:"response"`
}

type AmountType struct {
	Amount  *float64 `json:"amount"`
	CurCode *string  `json:"curCode"`
}

type CarrierOffers struct {
	Carrier *string  `json:"carrier"`
	Offers  []*Offer `json:"offers"`
}

type DataLists struct {
	PaxList        []*FlightPax  `json:"paxList"`
	PaxJourneyList []*PaxJourney `json:"paxJourneyList"`
	PaxSegmentList []*PaxSegment `json:"paxSegmentList"`
}

type DatedMarketingSegmentType struct {
	CarrierDesigCode                 *string `json:"carrierDesigCode"`
	MarketingCarrierFlightNumberText *string `json:"MarketingCarrierFlightNumberText"`
}

type Error struct {
	Code        string `json:"code"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

type ExternalMessage struct {
	Code    *string `json:"code"`
	Message string  `json:"message"`
}

type Fee struct {
	Amount   *AmountType `json:"amount"`
	DescText *TextType   `json:"descText"`
}

type FlightPax struct {
	Age   int              `json:"age"`
	Ptc   *IataPtcCodeType `json:"ptc"`
	PaxID string           `json:"paxID"`
}

type FlightXQuery struct {
	Search   *AirShoppingRs `json:"search"`
	Quote    *bool          `json:"quote"`
	Book     *bool          `json:"book"`
	Retrieve *bool          `json:"retrieve"`
	Cancel   *bool          `json:"cancel"`
	Issue    *bool          `json:"issue"`
}

type Individual struct {
	Code           string            `json:"code"`
	IndividualData []*IndividualData `json:"individualData"`
	CreatedAt      string            `json:"createdAt"`
	UpdatedAt      string            `json:"updatedAt"`
	AdviseMessage  []*AdviseMessage  `json:"adviseMessage"`
}

func (Individual) IsNode() {}

type IndividualData struct {
	Code      string  `json:"code"`
	Name      *string `json:"name"`
	Surname   string  `json:"surname"`
	BirthDate *string `json:"birthDate"`
	Title     *string `json:"title"`
}

type Offer struct {
	OfferID    *string      `json:"offerID"`
	TotalPrice *int         `json:"totalPrice"`
	OfferItems []*OfferItem `json:"offerItems"`
}

type OfferItem struct {
	OfferItemID *string    `json:"offerItemID"`
	PriceHotel  *Price     `json:"priceHotel"`
	PriceFlight *PriceX    `json:"priceFlight"`
	Services    []*Service `json:"services"`
}

type OriginDestination struct {
	OriginLoc *string `json:"originLoc"`
	DestLoc   *string `json:"destLoc"`
	Date      *string `json:"date"`
}

type PaxJourney struct {
	PaxJourneyID    string    `json:"paxJourneyID"`
	PaxSegmentRefID []*string `json:"paxSegmentRefID"`
}

type PaxSegment struct {
	PaxSegmentID         string                     `json:"paxSegmentID"`
	MarketingCarrierInfo *DatedMarketingSegmentType `json:"MarketingCarrierInfo"`
	Arrival              *TransportArrivalType      `json:"arrival"`
	Dep                  *TransportDepType          `json:"dep"`
}

type Price struct {
	Net   float64 `json:"net"`
	Gross float64 `json:"gross"`
}

type PriceX struct {
	BaseAmount  *AmountType       `json:"baseAmount"`
	Fee         []*Fee            `json:"fee"`
	Surcharge   []*SurchargeX     `json:"surcharge"`
	TaxSummary  []*TaxSummaryType `json:"taxSummary"`
	TotalAmount *AmountType       `json:"totalAmount"`
}

type ResponseX struct {
	OffersGroup []*CarrierOffers `json:"offersGroup"`
	DataLists   *DataLists       `json:"dataLists"`
}

type Service struct {
	ServiceID           *string              `json:"serviceID"`
	PaxRefID            *string              `json:"paxRefID"`
	ServiceAssociations *ServiceAssociations `json:"serviceAssociations"`
}

type ServiceAssociations struct {
	PaxJourneyRefID []*string `json:"paxJourneyRefID"`
}

type SurchargeX struct {
	TotalAmount *AmountType `json:"totalAmount"`
	Breakdown   []*Fee      `json:"breakdown"`
}

type TaxSummaryType struct {
	TotalTaxAmount *AmountType `json:"totalTaxAmount"`
	Tax            []*TaxType  `json:"tax"`
}

type TaxType struct {
	Amount   *AmountType `json:"amount"`
	DescText *TextType   `json:"descText"`
}

type TextType struct {
	Text *string `json:"text"`
}

type TransportArrivalType struct {
	IataLocationCode *string `json:"IATA_LocationCode"`
}

type TransportDepType struct {
	IataLocationCode          *string `json:"IATA_LocationCode"`
	AircraftScheduledDateTime *string `json:"AircraftScheduledDateTime"`
}

type IPax struct {
	Age int `json:"age"`
}

type IataPtcCodeType string

const (
	IataPtcCodeTypeAdt IataPtcCodeType = "ADT"
	IataPtcCodeTypeChd IataPtcCodeType = "CHD"
	IataPtcCodeTypeInf IataPtcCodeType = "INF"
)

var AllIataPtcCodeType = []IataPtcCodeType{
	IataPtcCodeTypeAdt,
	IataPtcCodeTypeChd,
	IataPtcCodeTypeInf,
}

func (e IataPtcCodeType) IsValid() bool {
	switch e {
	case IataPtcCodeTypeAdt, IataPtcCodeTypeChd, IataPtcCodeTypeInf:
		return true
	}
	return false
}

func (e IataPtcCodeType) String() string {
	return string(e)
}

func (e *IataPtcCodeType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = IataPtcCodeType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid IATA_PTC_CodeType", str)
	}
	return nil
}

func (e IataPtcCodeType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
