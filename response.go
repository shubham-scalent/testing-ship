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
	CompanyAutoShipmentInsuranceSetting interface{} `json:"company_auto_shipment_insurance_setting"`
	CovidZones                          struct {
		DeliveryZone interface{} `json:"delivery_zone"`
		PickupZone   interface{} `json:"pickup_zone"`
	} `json:"covid_zones"`
	Currency interface{} `json:"currency"`
	Data     struct {
		AvailableCourierCompanies []struct {
			AirMaxWeight           interface{}  `json:"air_max_weight"`
			AssuredAmount          interface{}  `json:"assured_amount"`
			BaseCourierID          interface{}  `json:"base_courier_id"`
			BaseWeight             interface{}  `json:"base_weight"`
			Blocked                interface{}  `json:"blocked"`
			CallBeforeDelivery     interface{}  `json:"call_before_delivery"`
			ChargeWeight           interface{}  `json:"charge_weight"`
			City                   interface{}  `json:"city"`
			Cod                    interface{}  `json:"cod"`
			CodCharges             interface{}  `json:"cod_charges"`
			CodMultiplier          interface{}  `json:"cod_multiplier"`
			Cost                   interface{}  `json:"cost"`
			CourierCompanyID       interface{}  `json:"courier_company_id"`
			CourierName            interface{}  `json:"courier_name"`
			CourierType            interface{}  `json:"courier_type"`
			CoverageCharges        interface{}  `json:"coverage_charges"`
			CutoffTime             interface{}  `json:"cutoff_time"`
			DeliveryBoyContact     interface{}  `json:"delivery_boy_contact"`
			DeliveryPerformance    interface{}  `json:"delivery_performance"`
			Description            interface{}  `json:"description"`
			Edd                    interface{}  `json:"edd"`
			EntryTax               interface{}  `json:"entry_tax"`
			EstimatedDeliveryDays  interface{}  `json:"estimated_delivery_days"`
			Etd                    interface{}  `json:"etd"`
			EtdHours               interface{}  `json:"etd_hours"`
			FreightCharge          interface{}  `json:"freight_charge"`
			ID                     interface{}  `json:"id"`
			IsCustomRate           interface{}  `json:"is_custom_rate"`
			IsHyperlocal           interface{}  `json:"is_hyperlocal"`
			IsInternational        interface{}  `json:"is_international"`
			IsRtoAddressAvailable  interface{}  `json:"is_rto_address_available"`
			IsSurface              interface{}  `json:"is_surface"`
			LocalRegion            interface{}  `json:"local_region"`
			Metro                  interface{}  `json:"metro"`
			MinWeight              interface{}  `json:"min_weight"`
			Mode                   interface{}  `json:"mode"`
			Odablock               interface{}  `json:"odablock"`
			OtherCharges           interface{}  `json:"other_charges"`
			Others                 interface{}  `json:"others"`
			PickupAvailability     interface{}  `json:"pickup_availability"`
			PickupPerformance      interface{}  `json:"pickup_performance"`
			PickupPriority         interface{}  `json:"pickup_priority"`
			PickupSupressHours     interface{}  `json:"pickup_supress_hours"`
			PodAvailable           interface{}  `json:"pod_available"`
			Postcode               interface{}  `json:"postcode"`
			QcCourier              interface{}  `json:"qc_courier"`
			Rank                   interface{}  `json:"rank"`
			Rate                   interface{}  `json:"rate"`
			Rating                 interface{}  `json:"rating"`
			RealtimeTracking       interface{}  `json:"realtime_tracking"`
			Region                 interface{}  `json:"region"`
			RtoCharges             interface{}  `json:"rto_charges"`
			RtoPerformance         interface{}  `json:"rto_performance"`
			SecondsLeftForPickup   interface{}  `json:"seconds_left_for_pickup"`
			SecureShipmentDisabled interface{}  `json:"secure_shipment_disabled"`
			ShipType               interface{}  `json:"ship_type"`
			State                  interface{}  `json:"state"`
			SuppressDate           interface{}  `json:"suppress_date"`
			SuppressText           interface{}  `json:"suppress_text"`
			SuppressionDates       interface{}  `json:"suppression_dates"`
			SurfaceMaxWeight       interface{}  `json:"surface_max_weight"`
			TrackingPerformance    interface{}  `json:"tracking_performance"`
			VolumetricMaxWeight    *interface{} `json:"volumetric_max_weight"`
			WeightCases            interface{}  `json:"weight_cases"`
			Zone                   interface{}  `json:"zone"`
		} `json:"available_courier_companies"`
		ChildCourierID            interface{} `json:"child_courier_id"`
		IsRecommendationEnabled   interface{} `json:"is_recommendation_enabled"`
		RecommendationAdvanceRule interface{} `json:"recommendation_advance_rule"`
		RecommendedBy             struct {
			ID    interface{} `json:"id"`
			Title interface{} `json:"title"`
		} `json:"recommended_by"`
		RecommendedCourierCompanyID    interface{} `json:"recommended_courier_company_id"`
		ShiprocketRecommendedCourierID interface{} `json:"shiprocket_recommended_courier_id"`
	} `json:"data"`
	DgCourier                    interface{}   `json:"dg_courier"`
	EligibleForInsurance         interface{}   `json:"eligible_for_insurance"`
	InsuraceOptedAtOrderCreation interface{}   `json:"insurace_opted_at_order_creation"`
	IsAllowTemplatizedPricing    interface{}   `json:"is_allow_templatized_pricing"`
	IsLatlong                    interface{}   `json:"is_latlong"`
	IsOldZoneOpted               interface{}   `json:"is_old_zone_opted"`
	IsZoneFromMongo              interface{}   `json:"is_zone_from_mongo"`
	LabelGenerateType            interface{}   `json:"label_generate_type"`
	OnNewZone                    interface{}   `json:"on_new_zone"`
	SellerAddress                []interface{} `json:"seller_address"`
	Status                       interface{}   `json:"status"`
	UserInsuranceManadatory      interface{}   `json:"user_insurance_manadatory"`
}

type CreateOrderResponse struct {
	AwbCode                interface{} `json:"awb_code"`
	CourierCompanyID       interface{} `json:"courier_company_id"`
	CourierName            interface{} `json:"courier_name"`
	OnboardingCompletedNow interface{} `json:"onboarding_completed_now"`
	OrderID                interface{} `json:"order_id"`
	ShipmentID             interface{} `json:"shipment_id"`
	Status                 interface{} `json:"status"`
	StatusCode             interface{} `json:"status_code"`
}

type GenerateInvoiceResponse struct {
	IsInvoiceCreated interface{}   `json:"is_invoice_created"`
	InvoiceURL       interface{}   `json:"invoice_url"`
	NotCreated       []interface{} `json:"not_created"`
}

type GenerateManifestResponse struct {
	ManifestURL interface{} `json:"manifest_url"`
	Status      interface{} `json:"status"`
}

type GenerateLabelRespone struct {
	LabelCreated interface{}   `json:"label_created"`
	LabelURL     interface{}   `json:"label_url"`
	NotCreated   []interface{} `json:"not_created"`
	Response     interface{}   `json:"response"`
}

type GenerateAWBResponse struct {
	AwbAssignStatus interface{} `json:"awb_assign_status"`
	Response        struct {
		Data struct {
			AppliedWeight    interface{} `json:"applied_weight"`
			AssignedDateTime struct {
				Date         interface{} `json:"date"`
				Timezone     interface{} `json:"timezone"`
				TimezoneType interface{} `json:"timezone_type"`
			} `json:"assigned_date_time"`
			AwbCode             interface{} `json:"awb_code"`
			AwbCodeStatus       interface{} `json:"awb_code_status"`
			ChildCourierName    interface{} `json:"child_courier_name"`
			Cod                 interface{} `json:"cod"`
			CompanyID           interface{} `json:"company_id"`
			CourierCompanyID    interface{} `json:"courier_company_id"`
			CourierName         interface{} `json:"courier_name"`
			InvoiceNo           interface{} `json:"invoice_no"`
			OrderID             interface{} `json:"order_id"`
			PickupScheduledDate interface{} `json:"pickup_scheduled_date"`
			RoutingCode         interface{} `json:"routing_code"`
			RtoRoutingCode      interface{} `json:"rto_routing_code"`
			ShipmentID          interface{} `json:"shipment_id"`
			ShippedBy           struct {
				Lat                       interface{} `json:"lat"`
				Long                      interface{} `json:"long"`
				RtoAddress1               interface{} `json:"rto_address_1"`
				RtoAddress2               interface{} `json:"rto_address_2"`
				RtoCity                   interface{} `json:"rto_city"`
				RtoCompanyName            interface{} `json:"rto_company_name"`
				RtoCountry                interface{} `json:"rto_country"`
				RtoEmail                  interface{} `json:"rto_email"`
				RtoPhone                  interface{} `json:"rto_phone"`
				RtoPostcode               interface{} `json:"rto_postcode"`
				RtoState                  interface{} `json:"rto_state"`
				ShipperAddress1           interface{} `json:"shipper_address_1"`
				ShipperAddress2           interface{} `json:"shipper_address_2"`
				ShipperCity               interface{} `json:"shipper_city"`
				ShipperCompanyName        interface{} `json:"shipper_company_name"`
				ShipperCountry            interface{} `json:"shipper_country"`
				ShipperEmail              interface{} `json:"shipper_email"`
				ShipperFirstMileActivated interface{} `json:"shipper_first_mile_activated"`
				ShipperPhone              interface{} `json:"shipper_phone"`
				ShipperPostcode           interface{} `json:"shipper_postcode"`
				ShipperState              interface{} `json:"shipper_state"`
			} `json:"shipped_by"`
			TransporterID   interface{} `json:"transporter_id"`
			TransporterName interface{} `json:"transporter_name"`
		} `json:"data"`
	} `json:"response"`
}

type TrackingThroughOrderIDResponse []struct {
	TrackingData struct {
		TrackStatus    interface{} `json:"track_status"`
		ShipmentStatus interface{} `json:"shipment_status"`
		ShipmentTrack  []struct {
			ID                  interface{} `json:"id"`
			AwbCode             interface{} `json:"awb_code"`
			CourierCompanyID    interface{} `json:"courier_company_id"`
			ShipmentID          interface{} `json:"shipment_id"`
			OrderID             interface{} `json:"order_id"`
			PickupDate          any         `json:"pickup_date"`
			DeliveredDate       any         `json:"delivered_date"`
			Weight              interface{} `json:"weight"`
			Packages            interface{} `json:"packages"`
			CurrentStatus       interface{} `json:"current_status"`
			DeliveredTo         interface{} `json:"delivered_to"`
			Destination         interface{} `json:"destination"`
			ConsigneeName       interface{} `json:"consignee_name"`
			Origin              interface{} `json:"origin"`
			CourierAgentDetails any         `json:"courier_agent_details"`
			Edd                 interface{} `json:"edd"`
		} `json:"shipment_track"`
		ShipmentTrackActivities []struct {
			Date     interface{} `json:"date"`
			Status   interface{} `json:"status"`
			Activity interface{} `json:"activity"`
			Location interface{} `json:"location"`
			SrStatus interface{} `json:"sr-status"`
		} `json:"shipment_track_activities"`
		TrackURL interface{} `json:"track_url"`
		Etd      interface{} `json:"etd"`
	} `json:"tracking_data"`
}

type TrackingThroughShipmentIDResponse struct {
	TrackingData struct {
		TrackStatus    interface{} `json:"track_status"`
		ShipmentStatus interface{} `json:"shipment_status"`
		ShipmentTrack  []struct {
			ID                  interface{} `json:"id"`
			AwbCode             interface{} `json:"awb_code"`
			CourierCompanyID    interface{} `json:"courier_company_id"`
			ShipmentID          interface{} `json:"shipment_id"`
			OrderID             interface{} `json:"order_id"`
			PickupDate          any         `json:"pickup_date"`
			DeliveredDate       any         `json:"delivered_date"`
			Weight              interface{} `json:"weight"`
			Packages            interface{} `json:"packages"`
			CurrentStatus       interface{} `json:"current_status"`
			DeliveredTo         interface{} `json:"delivered_to"`
			Destination         interface{} `json:"destination"`
			ConsigneeName       interface{} `json:"consignee_name"`
			Origin              interface{} `json:"origin"`
			CourierAgentDetails any         `json:"courier_agent_details"`
			Edd                 interface{} `json:"edd"`
		} `json:"shipment_track"`
		ShipmentTrackActivities []struct {
			Date          interface{} `json:"date"`
			Status        interface{} `json:"status"`
			Activity      interface{} `json:"activity"`
			Location      interface{} `json:"location"`
			SrStatus      interface{} `json:"sr-status"`
			SrStatusLabel interface{} `json:"sr-status-label"`
		} `json:"shipment_track_activities"`
		TrackURL interface{} `json:"track_url"`
		Etd      interface{} `json:"etd"`
	} `json:"tracking_data"`
}

type AddPickUpLocationResponse struct {
	Success interface{} `json:"success"`
	Address struct {
		CompanyID      interface{} `json:"company_id"`
		PickupCode     interface{} `json:"pickup_code"`
		UpdatedAt      time.Time   `json:"updated_at"`
		CreatedAt      time.Time   `json:"created_at"`
		ID             interface{} `json:"id"`
		Address        interface{} `json:"address"`
		Address2       interface{} `json:"address_2"`
		AddressType    any         `json:"address_type"`
		City           interface{} `json:"city"`
		State          interface{} `json:"state"`
		Country        interface{} `json:"country"`
		Gstin          any         `json:"gstin"`
		PinCode        interface{} `json:"pin_code"`
		Phone          interface{} `json:"phone"`
		Email          interface{} `json:"email"`
		Name           interface{} `json:"name"`
		AlternatePhone any         `json:"alternate_phone"`
		Lat            any         `json:"lat"`
		Long           any         `json:"long"`
		Status         interface{} `json:"status"`
		RtoAddressID   interface{} `json:"rto_address_id"`
		ExtraInfo      interface{} `json:"extra_info"`
		PhoneVerified  interface{} `json:"phone_verified"`
	} `json:"address"`
	PickupID    interface{} `json:"pickup_id"`
	CompanyName interface{} `json:"company_name"`
	FullName    interface{} `json:"full_name"`
}

type GetAllPickUpLocationResponse struct {
	Data struct {
		ShippingAddress []struct {
			ID                   interface{} `json:"id"`
			PickupLocation       interface{} `json:"pickup_location"`
			AddressType          interface{} `json:"address_type"`
			Address              interface{} `json:"address"`
			Address2             interface{} `json:"address_2"`
			UpdatedAddress       interface{} `json:"updated_address"`
			OldAddress           interface{} `json:"old_address"`
			OldAddress2          interface{} `json:"old_address2"`
			City                 interface{} `json:"city"`
			State                interface{} `json:"state"`
			Country              interface{} `json:"country"`
			PinCode              interface{} `json:"pin_code"`
			Email                interface{} `json:"email"`
			IsFirstMilePickup    interface{} `json:"is_first_mile_pickup"`
			Phone                interface{} `json:"phone"`
			Name                 interface{} `json:"name"`
			CompanyID            interface{} `json:"company_id"`
			Gstin                interface{} `json:"gstin"`
			VendorName           any         `json:"vendor_name"`
			Status               interface{} `json:"status"`
			PhoneVerified        interface{} `json:"phone_verified"`
			Lat                  interface{} `json:"lat"`
			Long                 interface{} `json:"long"`
			OpenTime             interface{} `json:"open_time"`
			CloseTime            interface{} `json:"close_time"`
			WarehouseCode        any         `json:"warehouse_code"`
			AlternatePhone       interface{} `json:"alternate_phone"`
			RtoAddressID         interface{} `json:"rto_address_id"`
			LatLongStatus        interface{} `json:"lat_long_status"`
			New                  interface{} `json:"new"`
			AssociatedRtoAddress any         `json:"associated_rto_address"`
		} `json:"shipping_address"`
		AllowMore        interface{} `json:"allow_more"`
		IsBlackboxSeller interface{} `json:"is_blackbox_seller"`
		CompanyName      interface{} `json:"company_name"`
		RecentAddresses  []any       `json:"recent_addresses"`
	} `json:"data"`
}

type CancelShipmentResponse struct {
	Message interface{} `json:"message"`
}
type ShipmentPickUpResponse struct {
	PickupStatus interface{} `json:"pickup_status"`
	Response     struct {
		PickupScheduledDate interface{} `json:"pickup_scheduled_date"`
		PickupTokenNumber   interface{} `json:"pickup_token_number"`
		Status              interface{} `json:"status"`
		Others              interface{} `json:"others"`
		PickupGeneratedDate struct {
			Date         interface{} `json:"date"`
			TimezoneType interface{} `json:"timezone_type"`
			Timezone     interface{} `json:"timezone"`
		} `json:"pickup_generated_date"`
		Data interface{} `json:"data"`
	} `json:"response"`
}

type GetLocalityDetailsResponse struct {
	Success         interface{} `json:"success"`
	PostcodeDetails struct {
		Postcode  interface{}   `json:"postcode"`
		City      interface{}   `json:"city"`
		Locality  []interface{} `json:"locality"`
		State     interface{}   `json:"state"`
		StateCode interface{}   `json:"state_code"`
		Longitude interface{}   `json:"longitude"`
		Latitude  interface{}   `json:"latitude"`
		Country   interface{}   `json:"country"`
	} `json:"postcode_details"`
}

type GetWalletBalanceResponse struct {
	Data struct {
		BalanceAmount interface{} `json:"balance_amount"`
	} `json:"data"`
}

type GetSpecificNDRShipmentDetailsResponse struct {
	Data []struct {
		ID                interface{} `json:"id"`
		ShipmentID        interface{} `json:"shipment_id"`
		ChannelOrderID    interface{} `json:"channel_order_id"`
		CustomerName      interface{} `json:"customer_name"`
		CustomerEmail     interface{} `json:"customer_email"`
		CustomerPhone     interface{} `json:"customer_phone"`
		CustomerAddress   interface{} `json:"customer_address"`
		CustomerAddress2  interface{} `json:"customer_address_2"`
		CustomerCity      interface{} `json:"customer_city"`
		CustomerState     interface{} `json:"customer_state"`
		CustomerPincode   interface{} `json:"customer_pincode"`
		PaymentStatus     interface{} `json:"payment_status"`
		Status            interface{} `json:"status"`
		StatusCode        interface{} `json:"status_code"`
		PaymentMethod     interface{} `json:"payment_method"`
		CreatedAt         interface{} `json:"created_at"`
		Reason            interface{} `json:"reason"`
		Attempts          interface{} `json:"attempts"`
		NDRRaisedAt       interface{} `json:"ndr_raised_at"`
		Courier           interface{} `json:"courier"`
		AWBCode           interface{} `json:"awb_code"`
		EscalationStatus  interface{} `json:"escalation_status"`
		ProductName       interface{} `json:"product_name"`
		ProductPrice      interface{} `json:"product_price"`
		ShipmentChannelID interface{} `json:"shipment_channel_id"`
		History           []struct {
			ID                          interface{} `json:"id"`
			NdrID                       interface{} `json:"ndr_id"`
			NDRReason                   interface{} `json:"ndr_reason"`
			ActionBy                    interface{} `json:"action_by"`
			NDRAttempt                  interface{} `json:"ndr_attempt"`
			Medium                      interface{} `json:"medium"`
			NDRPushStatus               interface{} `json:"ndr_push_status"`
			Comment                     interface{} `json:"comment"`
			CallCenterCallRecording     interface{} `json:"call_center_call_recording"`
			CallCenterCallRecordingDate interface{} `json:"call_center_recording_date"`
			ProofRecording              interface{} `json:"proof_recording"`
			ProofImage                  interface{} `json:"proof_image"`
			SMSResponse                 interface{} `json:"sms_response"`
			NDRRaisedAt                 interface{} `json:"ndr_raised_at"`
		} `json:"history"`
		DeliveredDate      interface{} `json:"delivered_date"`
		CancellationReason interface{} `json:"cancellation_reason"`
	} `json:"data"`
	Meta struct {
		Pagination struct {
			Total       interface{} `json:"total"`
			Count       interface{} `json:"count"`
			PerPage     interface{} `json:"per_page"`
			CurrentPage interface{} `json:"current_page"`
			TotalPages  interface{} `json:"total_pages"`
			Links       interface{} `json:"links"`
		} `json:"pagination"`
	} `json:"meta"`
}

type GetAllNDRShipmentsResponse struct {
	Data []struct {
		ID                interface{} `json:"id"`
		ShipmentID        interface{} `json:"shipment_id"`
		ChannelOrderID    interface{} `json:"channel_order_id"`
		CustomerName      interface{} `json:"customer_name"`
		CustomerEmail     interface{} `json:"customer_email"`
		CustomerPhone     interface{} `json:"customer_phone"`
		CustomerAddress   interface{} `json:"customer_address"`
		CustomerAddress2  interface{} `json:"customer_address_2"`
		CustomerCity      interface{} `json:"customer_city"`
		CustomerState     interface{} `json:"customer_state"`
		CustomerPincode   interface{} `json:"customer_pincode"`
		PaymentStatus     interface{} `json:"payment_status"`
		Status            interface{} `json:"status"`
		StatusCode        interface{} `json:"status_code"`
		PaymentMethod     interface{} `json:"payment_method"`
		CreatedAt         interface{} `json:"created_at"`
		Reason            interface{} `json:"reason"`
		Attempts          interface{} `json:"attempts"`
		NDRRaisedAt       interface{} `json:"ndr_raised_at"`
		Courier           interface{} `json:"courier"`
		AWBCode           interface{} `json:"awb_code"`
		EscalationStatus  interface{} `json:"escalation_status"`
		ProductName       interface{} `json:"product_name"`
		ProductPrice      interface{} `json:"product_price"`
		ShipmentChannelID interface{} `json:"shipment_channel_id"`
		History           struct {
			ID                          interface{} `json:"id"`
			NdrID                       interface{} `json:"ndr_id"`
			NDRReason                   interface{} `json:"ndr_reason"`
			ActionBy                    interface{} `json:"action_by"`
			NDRAttempt                  interface{} `json:"ndr_attempt"`
			Medium                      interface{} `json:"medium"`
			NDRPushStatus               interface{} `json:"ndr_push_status"`
			Comment                     interface{} `json:"comment"`
			CallCenterCallRecording     interface{} `json:"call_center_call_recording"`
			CallCenterCallRecordingDate interface{} `json:"call_center_recording_date"`
			ProofRecording              interface{} `json:"proof_recording"`
			ProofImage                  interface{} `json:"proof_image"`
			SMSResponse                 interface{} `json:"sms_response"`
			NDRRaisedAt                 interface{} `json:"ndr_raised_at"`
		} `json:"history"`
		DeliveredDate interface{} `json:"delivered_date"`
	} `json:"data"`
	Meta struct {
		Pagination struct {
			Total       interface{} `json:"total"`
			Count       interface{} `json:"count"`
			PerPage     interface{} `json:"per_page"`
			CurrentPage interface{} `json:"current_page"`
			TotalPages  interface{} `json:"total_pages"`
			Links       struct {
				Next interface{} `json:"next"`
			} `json:"links"`
		} `json:"pagination"`
	} `json:"meta"`
}

type NDRActionResponse struct {
	Status interface{} `json:"status"`
}

type GetDiscrepancyDataResponse struct {
	Status interface{} `json:"status"`
	Data   []struct {
	} `json:"data"`
	UpperFoldText interface{} `json:"upper_fold_text"`
	LowerFildText interface{} `json:"lower_fild_text"`
}
