/*
 * Copyright (c) 2016 Stewart Buskirk <mitymauser@gmail.com>
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */
package goeasypost

import "time"



type Fee struct {
	Type     string `json:"type"`
	Amount   string `json:"amount"`
	Charged  bool   `json:"charges"`
	Refunded bool   `json:"refunded"`
}

type TrackerList struct {
	Trackers []Tracker `json:"trackers"`
	HasMore  bool      `json:"has_more"`
}

type TrackingLocation struct {
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
	Zip     string `json:"zip"`
}

type TrackingDetail struct {
	Status           string           `json:"status"`
	Message          string           `json:"message"`
	Datetime         time.Time        `json:"datetime"`
	TrackingLocation TrackingLocation `json:"tracking_location"`
}

type CarrierDetail struct {
	Service              string     `json:"service"`
	ContainerType        string     `json:"container_type"`
	EstDeliveryDateLocal *time.Time `json:"est_delivery_date_local"`
	EstDeliveryTimeLocal *time.Time `json:"est_delivery_time_local"`
}

type Tracker struct {
	Id              string           `json:"id"`
	Mode            string           `json:"mode"`
	TrackingCode    string           `json:"tracking_code"`
	Status          string           `json:"status"`
	ShipmentId      string           `json:"shipment_id"`
	Carrier         string           `json:"carrier"`
	TrackingDetails []TrackingDetail `json:"tracking_details"`
	Weight          float32          `json:"weight"`
	EstDeliveryDate *time.Time       `json:"est_delivery_date"`
	CreatedAt       time.Time        `json:"created_at"`
	UpdatedAt       time.Time        `json:"updated_at"`
	SignedBy        string           `json:"signed_by"`
	CarrierDetail   CarrierDetail    `json:"carrier_detail"`
	Fees            []Fee            `json:"fees"`
}

type Address struct {
	VerifyRequest   *VerifyRequest `json:"-"`
	Id              string         `json:"id"`
	Object          string         `json:"object"`
	Mode            string         `json:"mode"`
	Street1         string         `json:"street1"`
	Street2         string         `json:"street2"`
	City            string         `json:"city"`
	State           string         `json:"state"`
	Zip             string         `json:"zip"`
	Country         string         `json:"country"`
	Residential     bool           `json:"residential"`
	CarrierFacility string         `json:"carrier_facility"`
	Name            string         `json:"name"`
	Company         string         `json:"company"`
	Phone           string         `json:"phone"`
	Email           string         `json:"email"`
	FederalTaxId    string         `json:"federal_tax_id"`
	StateTaxId      string         `json:"state_tax_id"`
	Verifications   *Verifications `json:"verifications"`
}

type VerifyRequest struct {
	Strict bool
	Type   string
}

type Verifications struct {
	Zip4     *Verification `json:"zip4"`
	Delivery *Verification `json:"delivery"`
}

type Verification struct {
	Success bool         `json:"success"`
	Errors  [][][]string `json:"errors"`
}

type Parcel struct {
	Id                string    `json:"id"`
	Object            string    `json:"object"`
	Mode              string    `json:"mode"`
	Length            float32   `json:"length"`
	Width             float32   `json:"width"`
	Height            float32   `json:"height"`
	PredefinedPackage string    `json:"predefined_package"`
	Weight            float32   `json:"weight"`
	CreatedAt         *time.Time `json:"created_at"`
	UpdatedAt         *time.Time `json:"updated_at"`
}

type Message struct {
Carrier 	string `json:"carrier"`
Type 	string 	`json:"type"`
Message 	string 	`json:"message"`
}

 type CustomsItem struct {
	 Id                string    `json:"id"`
	 Object            string    `json:"object"`
Description 	string 	 `json:"description"`
Quantity 	float32 	 `json:"quantity"`
Value 	float32 	 `json:"value"`
Weight 	float32 	 `json:"weight"`
HsTariffNumber 	string 	 `json:"hs_tariff_number"`
OriginCountry string 	 `json:"origin_country"`
Currency 	string 	 `json:"currency"`
CreatedAt         time.Time `json:"created_at"`
UpdatedAt         time.Time `json:"updated_at"`
 }

type CustomsInfo struct {
	CustomsCertify 	bool      `json:"customs_certify"`
	CustomsSigner 	string  `json:"customs_signer"`
	ContentsType 	string `json:"contents_type"`
	RestrictionType 	string    `json:"restriction_type"`
	EelPfc 	string				`json:"eel_pfc"`
	CustomsItems 	[]CustomsItem   `json:"customs_items"`
}

type Rate struct {
	Id            string    `json:"id"`
	Object        string    `json:"object"`
	Mode          string    `json:"mode"`
Service 	string 	  `json:"service"`
Carrier 	string 	  `json:"carrier"`
CarrierAccountId 	string  `json:"carrier_account_id"`
ShipmentId 	string 	  `json:"shipment_id"`
Rate 	string 	  `json:"rate"`
Currency 	string 	  `json:"currency"`
RetailRate string	  `json:"retail_rate"`
RetailCurrency 	string  `json:"retail_currency"`
ListRate 	string   `json:"list_rate"`
ListCurrency 	string 	  `json:"list_currency"`
DeliveryDays 	int32 	  `json:"delivery_days"`
EstDeliveryDays 	int32 	  `json:"est_delivery_days"`
DeliveryDate 	string 	  `json:"delivery_date"`
DeliveryDateGuaranteed 	bool 	  `json:"delivery_date_guaranteed"`
	CreatedAt     *time.Time     `json:"created_at"`
	UpdatedAt     *time.Time    `json:"updated_at"`
}

type PostageLabel struct {

	Id            string    `json:"id"`
	Object        string    `json:"object"`
IntegratedForm  string    `json:"integrated_form"`
LabelDate 		*time.Time    `json:"label_date"`
	labelEpl2Url  string    `json:"label_epl2_url"`
LabelFileType  string    `json:"label_file_type"`
LabelPdfUrl string    `json:"label_pdf_url"`
LabelResolution int32     `json:"label_resolution"`
LabelSize string    `json:"label_size"`
LabelType string    `json:"label_type"`
LabelUrl string    `json:"label_url"`
LabelZplUrl string    `json:"label_zpl_url"`
	CreatedAt     time.Time     `json:"created_at"`
	UpdatedAt     time.Time    `json:"updated_at"`
}


type Options struct {
	AdditionalHandling	bool	`json:"additional_handling"`
	AddressValidationLevel	string	`json:"address_validation_level"`
	Alcohol	bool	`json:"alcohol"`
	BillReceiverAccount	string	`json:"bill_receiver_account"`
	BillReceiverPostalCode	string	`json:"bill_receiver_postal_code"`
	BillThirdPartyAccount	string	`json:"bill_third_party_account"`
	BillThirdPartyCountry	string	`json:"bill_third_party_country"`
	BillThirdPartyPostalCode	string	`json:"bill_third_party_postal_code"`
	ByDrone	bool	`json:"by_drone"`
	CarbonNeutral	bool	`json:"carbon_neutral"`
	CodAmount	string	`json:"cod_amount"`
	CodMethod	string	`json:"cod_method"`
	Currency	string	`json:"currency"`
	DeliveredDutyPaid	bool	`json:"delivered_duty_paid"`
	DeliveryConfirmation	string	`json:"delivery_confirmation"`
	DryIce	bool	`json:"dry_ice"`
	DryIceMedical	string	`json:"dry_ice_medical"`
	DryIceWeight	string	`json:"dry_ice_weight"`
	FreightCharge	float32	`json:"freight_charge"`
	HandlingInstructions	string	`json:"handling_instructions"`
	Hazmat	string	`json:"hazmat"`
	HoldForPickup	bool	`json:"hold_for_pickup"`
	InvoiceNumber	string	`json:"invoice_number"`
	LabelDate	string	`json:"label_date"`
	LabelFormat	string	`json:"label_format"`
	Machinable	bool	`json:"machinable"`
	PrintCustom1	string	`json:"print_custom_1"`
	PrintCustom1Barcode	bool	`json:"print_custom_1_barcode"`
	PrintCustom1Code	string	`json:"print_custom_1_code"`
	PrintCustom2	string	`json:"print_custom_2"`
	PrintCustom2Barcode	bool	`json:"print_custom_2_barcode"`
	PrintCustom2Code	string	`json:"print_custom_2_code"`
	PrintCustom3	string	`json:"print_custom_3"`
	PrintCustom3Barcode	bool	`json:"print_custom_3_barcode"`
	PrintCustom3Code	string	`json:"print_custom_3_code"`
	SaturdayDelivery	bool	`json:"saturday_delivery"`
	SmartpostHub	string	`json:"smartpost_hub"`
	SmartpostManifest	string	`json:"smartpost_manifest"`
	SpecialRatesEligibility	string	`json:"special_rates_eligibility"`
}
type Shipment struct {
	Id            string    `json:"id"`
	Object        string    `json:"object"`
	Mode          string    `json:"mode"`
	Reference     string    `json:"reference"`
	ToAddress     *Address    `json:"to_address"`
	FromAddress   *Address    `json:"from_address"`
	ReturnAddress *Address    `json:"return_address"`
	BuyerAddress  *Address    `json:"buyer_address"`
	Parcel        Parcel    `json:"parcel"`
	CustomsInfo   *CustomsInfo    `json:"customs_info"`
	ScanForm      string    `json:"scan_form"`
	Forms         []string    `json:"forms"`
	Insurance     string    `json:"insurance"`
	Rates         []Rate    `json:"rates"`
	SelectedRate  *Rate    `json:"selected_rate"`
	PostageLabel  *PostageLabel    `json:"postage_label"`
	Messages      []Message    `json:"messages"`
	Options       *Options    `json:"options"`
	IsReturn      bool    `json:"is_return"`
	TrackingCode  string    `json:"tracking_code"`
	UspsZone      int32 `json:"usps_zone"`
	Status        string    `json:"status"`
	Tracker       *Tracker    `json:"tracker"`
	Fees          []Fee    `json:"fees"`
	RefundStatus  string   `json:"refund_status"`
	BatchId       string    `json:"batch_id"`
	BatchStatus   string    `json:"batch_status"`
	BatchMessage  string    `json:"batch_message"`
	CreatedAt     *time.Time     `json:"created_at"`
	UpdatedAt     *time.Time    `json:"updated_at"`
}


