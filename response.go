package shiprocket

import "time"

type ClientResponse struct {
	Token string `json:"token"`
}

type ErrorResponse struct {
	StatusCode int                 `json:"status_code"`
	Message    string              `json:"message"`
	Errors     map[string][]string `json:"errors"`
}

type CourierAvailabityResponse struct {
	CompanyAutoShipmentInsuranceSetting bool `json:"company_auto_shipment_insurance_setting"`
	CovidZones                          struct {
		DeliveryZone interface{} `json:"delivery_zone"`
		PickupZone   interface{} `json:"pickup_zone"`
	} `json:"covid_zones"`
	Currency string `json:"currency"`
	Data     struct {
		AvailableCourierCompanies []struct {
			AirMaxWeight           string      `json:"air_max_weight"`
			AssuredAmount          int64       `json:"assured_amount"`
			BaseCourierID          interface{} `json:"base_courier_id"`
			BaseWeight             string      `json:"base_weight"`
			Blocked                int64       `json:"blocked"`
			CallBeforeDelivery     string      `json:"call_before_delivery"`
			ChargeWeight           float64     `json:"charge_weight"`
			City                   string      `json:"city"`
			Cod                    int64       `json:"cod"`
			CodCharges             float64     `json:"cod_charges"`
			CodMultiplier          float64     `json:"cod_multiplier"`
			Cost                   string      `json:"cost"`
			CourierCompanyID       int64       `json:"courier_company_id"`
			CourierName            string      `json:"courier_name"`
			CourierType            string      `json:"courier_type"`
			CoverageCharges        int64       `json:"coverage_charges"`
			CutoffTime             string      `json:"cutoff_time"`
			DeliveryBoyContact     string      `json:"delivery_boy_contact"`
			DeliveryPerformance    float64     `json:"delivery_performance"`
			Description            string      `json:"description"`
			Edd                    string      `json:"edd"`
			EntryTax               int64       `json:"entry_tax"`
			EstimatedDeliveryDays  string      `json:"estimated_delivery_days"`
			Etd                    string      `json:"etd"`
			EtdHours               int64       `json:"etd_hours"`
			FreightCharge          float64     `json:"freight_charge"`
			ID                     int64       `json:"id"`
			IsCustomRate           int64       `json:"is_custom_rate"`
			IsHyperlocal           bool        `json:"is_hyperlocal"`
			IsInternational        int64       `json:"is_international"`
			IsRtoAddressAvailable  bool        `json:"is_rto_address_available"`
			IsSurface              bool        `json:"is_surface"`
			LocalRegion            int64       `json:"local_region"`
			Metro                  int64       `json:"metro"`
			MinWeight              float64     `json:"min_weight"`
			Mode                   int64       `json:"mode"`
			Odablock               bool        `json:"odablock"`
			OtherCharges           int64       `json:"other_charges"`
			Others                 string      `json:"others"`
			PickupAvailability     string      `json:"pickup_availability"`
			PickupPerformance      float64     `json:"pickup_performance"`
			PickupPriority         string      `json:"pickup_priority"`
			PickupSupressHours     int64       `json:"pickup_supress_hours"`
			PodAvailable           string      `json:"pod_available"`
			Postcode               string      `json:"postcode"`
			QcCourier              int64       `json:"qc_courier"`
			Rank                   string      `json:"rank"`
			Rate                   float64     `json:"rate"`
			Rating                 float64     `json:"rating"`
			RealtimeTracking       string      `json:"realtime_tracking"`
			Region                 int64       `json:"region"`
			RtoCharges             float64     `json:"rto_charges"`
			RtoPerformance         float64     `json:"rto_performance"`
			SecondsLeftForPickup   int64       `json:"seconds_left_for_pickup"`
			SecureShipmentDisabled bool        `json:"secure_shipment_disabled"`
			ShipType               int64       `json:"ship_type"`
			State                  string      `json:"state"`
			SuppressDate           string      `json:"suppress_date"`
			SuppressText           string      `json:"suppress_text"`
			SuppressionDates       interface{} `json:"suppression_dates"`
			SurfaceMaxWeight       string      `json:"surface_max_weight"`
			TrackingPerformance    float64     `json:"tracking_performance"`
			VolumetricMaxWeight    *int64      `json:"volumetric_max_weight"`
			WeightCases            float64     `json:"weight_cases"`
			Zone                   string      `json:"zone"`
		} `json:"available_courier_companies"`
		ChildCourierID            interface{} `json:"child_courier_id"`
		IsRecommendationEnabled   int64       `json:"is_recommendation_enabled"`
		RecommendationAdvanceRule int64       `json:"recommendation_advance_rule"`
		RecommendedBy             struct {
			ID    int64  `json:"id"`
			Title string `json:"title"`
		} `json:"recommended_by"`
		RecommendedCourierCompanyID    int64 `json:"recommended_courier_company_id"`
		ShiprocketRecommendedCourierID int64 `json:"shiprocket_recommended_courier_id"`
	} `json:"data"`
	DgCourier                    int64         `json:"dg_courier"`
	EligibleForInsurance         int64         `json:"eligible_for_insurance"`
	InsuraceOptedAtOrderCreation bool          `json:"insurace_opted_at_order_creation"`
	IsAllowTemplatizedPricing    bool          `json:"is_allow_templatized_pricing"`
	IsLatlong                    int64         `json:"is_latlong"`
	IsOldZoneOpted               bool          `json:"is_old_zone_opted"`
	IsZoneFromMongo              bool          `json:"is_zone_from_mongo"`
	LabelGenerateType            int64         `json:"label_generate_type"`
	OnNewZone                    int64         `json:"on_new_zone"`
	SellerAddress                []interface{} `json:"seller_address"`
	Status                       int64         `json:"status"`
	UserInsuranceManadatory      bool          `json:"user_insurance_manadatory"`
}

type CreateOrderResponse struct {
	AwbCode                interface{} `json:"awb_code"`
	CourierCompanyID       interface{} `json:"courier_company_id"`
	CourierName            interface{} `json:"courier_name"`
	OnboardingCompletedNow int64       `json:"onboarding_completed_now"`
	OrderID                int64       `json:"order_id"`
	ShipmentID             int64       `json:"shipment_id"`
	Status                 string      `json:"status"`
	StatusCode             int64       `json:"status_code"`
}

type GenerateInvoiceResponse struct {
	IsInvoiceCreated bool          `json:"is_invoice_created"`
	InvoiceURL       string        `json:"invoice_url"`
	NotCreated       []interface{} `json:"not_created"`
}

type GenerateManifestResponse struct {
	ManifestURL string `json:"manifest_url"`
	Status      int64  `json:"status"`
}

type GenerateLabelRespone struct {
	LabelCreated int64         `json:"label_created"`
	LabelURL     string        `json:"label_url"`
	NotCreated   []interface{} `json:"not_created"`
	Response     string        `json:"response"`
}

type GenerateAWBResponse struct {
	AwbAssignStatus int64 `json:"awb_assign_status"`
	Response        struct {
		Data struct {
			AppliedWeight    float64 `json:"applied_weight"`
			AssignedDateTime struct {
				Date         string `json:"date"`
				Timezone     string `json:"timezone"`
				TimezoneType int64  `json:"timezone_type"`
			} `json:"assigned_date_time"`
			AwbCode             string      `json:"awb_code"`
			AwbCodeStatus       int64       `json:"awb_code_status"`
			ChildCourierName    interface{} `json:"child_courier_name"`
			Cod                 int64       `json:"cod"`
			CompanyID           int64       `json:"company_id"`
			CourierCompanyID    int64       `json:"courier_company_id"`
			CourierName         string      `json:"courier_name"`
			InvoiceNo           string      `json:"invoice_no"`
			OrderID             int64       `json:"order_id"`
			PickupScheduledDate string      `json:"pickup_scheduled_date"`
			RoutingCode         string      `json:"routing_code"`
			RtoRoutingCode      string      `json:"rto_routing_code"`
			ShipmentID          int64       `json:"shipment_id"`
			ShippedBy           struct {
				Lat                       string `json:"lat"`
				Long                      string `json:"long"`
				RtoAddress1               string `json:"rto_address_1"`
				RtoAddress2               string `json:"rto_address_2"`
				RtoCity                   string `json:"rto_city"`
				RtoCompanyName            string `json:"rto_company_name"`
				RtoCountry                string `json:"rto_country"`
				RtoEmail                  string `json:"rto_email"`
				RtoPhone                  string `json:"rto_phone"`
				RtoPostcode               string `json:"rto_postcode"`
				RtoState                  string `json:"rto_state"`
				ShipperAddress1           string `json:"shipper_address_1"`
				ShipperAddress2           string `json:"shipper_address_2"`
				ShipperCity               string `json:"shipper_city"`
				ShipperCompanyName        string `json:"shipper_company_name"`
				ShipperCountry            string `json:"shipper_country"`
				ShipperEmail              string `json:"shipper_email"`
				ShipperFirstMileActivated int64  `json:"shipper_first_mile_activated"`
				ShipperPhone              string `json:"shipper_phone"`
				ShipperPostcode           string `json:"shipper_postcode"`
				ShipperState              string `json:"shipper_state"`
			} `json:"shipped_by"`
			TransporterID   string `json:"transporter_id"`
			TransporterName string `json:"transporter_name"`
		} `json:"data"`
	} `json:"response"`
}

type TrackingThroughOrderIDResponse []struct {
	TrackingData struct {
		TrackStatus    int `json:"track_status"`
		ShipmentStatus int `json:"shipment_status"`
		ShipmentTrack  []struct {
			ID                  int    `json:"id"`
			AwbCode             string `json:"awb_code"`
			CourierCompanyID    int    `json:"courier_company_id"`
			ShipmentID          int    `json:"shipment_id"`
			OrderID             int    `json:"order_id"`
			PickupDate          any    `json:"pickup_date"`
			DeliveredDate       any    `json:"delivered_date"`
			Weight              string `json:"weight"`
			Packages            int    `json:"packages"`
			CurrentStatus       string `json:"current_status"`
			DeliveredTo         string `json:"delivered_to"`
			Destination         string `json:"destination"`
			ConsigneeName       string `json:"consignee_name"`
			Origin              string `json:"origin"`
			CourierAgentDetails any    `json:"courier_agent_details"`
			Edd                 string `json:"edd"`
		} `json:"shipment_track"`
		ShipmentTrackActivities []struct {
			Date     string `json:"date"`
			Status   string `json:"status"`
			Activity string `json:"activity"`
			Location string `json:"location"`
			SrStatus int    `json:"sr-status"`
		} `json:"shipment_track_activities"`
		TrackURL string `json:"track_url"`
		Etd      string `json:"etd"`
	} `json:"tracking_data"`
}

type TrackingThroughShipmentIDResponse struct {
	TrackingData struct {
		TrackStatus    int `json:"track_status"`
		ShipmentStatus int `json:"shipment_status"`
		ShipmentTrack  []struct {
			ID                  int    `json:"id"`
			AwbCode             string `json:"awb_code"`
			CourierCompanyID    int    `json:"courier_company_id"`
			ShipmentID          int    `json:"shipment_id"`
			OrderID             int    `json:"order_id"`
			PickupDate          any    `json:"pickup_date"`
			DeliveredDate       any    `json:"delivered_date"`
			Weight              string `json:"weight"`
			Packages            int    `json:"packages"`
			CurrentStatus       string `json:"current_status"`
			DeliveredTo         string `json:"delivered_to"`
			Destination         string `json:"destination"`
			ConsigneeName       string `json:"consignee_name"`
			Origin              string `json:"origin"`
			CourierAgentDetails any    `json:"courier_agent_details"`
			Edd                 string `json:"edd"`
		} `json:"shipment_track"`
		ShipmentTrackActivities []struct {
			Date          string `json:"date"`
			Status        string `json:"status"`
			Activity      string `json:"activity"`
			Location      string `json:"location"`
			SrStatus      string `json:"sr-status"`
			SrStatusLabel string `json:"sr-status-label"`
		} `json:"shipment_track_activities"`
		TrackURL string `json:"track_url"`
		Etd      string `json:"etd"`
	} `json:"tracking_data"`
}

type AddPickUpLocationResponse struct {
	Success bool `json:"success"`
	Address struct {
		CompanyID      int       `json:"company_id"`
		PickupCode     string    `json:"pickup_code"`
		UpdatedAt      time.Time `json:"updated_at"`
		CreatedAt      time.Time `json:"created_at"`
		ID             int       `json:"id"`
		Address        string    `json:"address"`
		Address2       string    `json:"address_2"`
		AddressType    any       `json:"address_type"`
		City           string    `json:"city"`
		State          string    `json:"state"`
		Country        string    `json:"country"`
		Gstin          any       `json:"gstin"`
		PinCode        string    `json:"pin_code"`
		Phone          string    `json:"phone"`
		Email          string    `json:"email"`
		Name           string    `json:"name"`
		AlternatePhone any       `json:"alternate_phone"`
		Lat            any       `json:"lat"`
		Long           any       `json:"long"`
		Status         int       `json:"status"`
		RtoAddressID   int       `json:"rto_address_id"`
		ExtraInfo      string    `json:"extra_info"`
		PhoneVerified  int       `json:"phone_verified"`
	} `json:"address"`
	PickupID    int    `json:"pickup_id"`
	CompanyName string `json:"company_name"`
	FullName    string `json:"full_name"`
}

type GetAllPickUpLocationResponse struct {
	Data struct {
		ShippingAddress []struct {
			ID                   int    `json:"id"`
			PickupLocation       string `json:"pickup_location"`
			AddressType          string `json:"address_type"`
			Address              string `json:"address"`
			Address2             string `json:"address_2"`
			UpdatedAddress       bool   `json:"updated_address"`
			OldAddress           string `json:"old_address"`
			OldAddress2          string `json:"old_address2"`
			City                 string `json:"city"`
			State                string `json:"state"`
			Country              string `json:"country"`
			PinCode              string `json:"pin_code"`
			Email                string `json:"email"`
			IsFirstMilePickup    int    `json:"is_first_mile_pickup"`
			Phone                string `json:"phone"`
			Name                 string `json:"name"`
			CompanyID            int    `json:"company_id"`
			Gstin                string `json:"gstin"`
			VendorName           any    `json:"vendor_name"`
			Status               int    `json:"status"`
			PhoneVerified        int    `json:"phone_verified"`
			Lat                  string `json:"lat"`
			Long                 string `json:"long"`
			OpenTime             string `json:"open_time"`
			CloseTime            string `json:"close_time"`
			WarehouseCode        any    `json:"warehouse_code"`
			AlternatePhone       string `json:"alternate_phone"`
			RtoAddressID         int    `json:"rto_address_id"`
			LatLongStatus        int    `json:"lat_long_status"`
			New                  int    `json:"new"`
			AssociatedRtoAddress any    `json:"associated_rto_address"`
		} `json:"shipping_address"`
		AllowMore        string `json:"allow_more"`
		IsBlackboxSeller bool   `json:"is_blackbox_seller"`
		CompanyName      string `json:"company_name"`
		RecentAddresses  []any  `json:"recent_addresses"`
	} `json:"data"`
}

type CancelShipmentResponse struct {
	Message string `json:"message"`
}
type ShipmentPickUpResponse struct {
	PickupStatus int64 `json:"pickup_status"`
	Response     struct {
		PickupScheduledDate string `json:"pickup_scheduled_date"`
		PickupTokenNumber   string `json:"pickup_token_number"`
		Status              int64  `json:"status"`
		Others              string `json:"others"`
		PickupGeneratedDate struct {
			Date         string `json:"date"`
			TimezoneType int64  `json:"timezone_type"`
			Timezone     string `json:"timezone"`
		} `json:"pickup_generated_date"`
		Data string `json:"data"`
	} `json:"response"`
}

type GetLocalityDetailsResponse struct {
	Success         bool `json:"success"`
	PostcodeDetails struct {
		Postcode  string   `json:"postcode"`
		City      string   `json:"city"`
		Locality  []string `json:"locality"`
		State     string   `json:"state"`
		StateCode string   `json:"state_code"`
		Longitude string   `json:"longitude"`
		Latitude  string   `json:"latitude"`
		Country   string   `json:"country"`
	} `json:"postcode_details"`
}
