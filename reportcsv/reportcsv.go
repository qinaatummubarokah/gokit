package reportcsv

import "context"

// type Csvdata struct{
// 	id
// 	fare
// 	extra
// 	discount_amt
// 	paid_amount
// 	promo_code
// 	payment_token
// 	transaction_time
// 	identifier
// 	payment_type
// 	vehicle_id
// 	vehicle_name
// 	service_type
// 	driver_id
// 	pick_up_suburb
// 	pick_up_area
// 	destination_area
// 	pick_up_latitude
// 	pick_up_lng
// 	payment_profile_id
// 	state
// 	released_at
// 	completed_at
// 	created_at
// 	updated_at
// 	cc_identifier
// 	account_id
// 	sap_sent_at
// 	sap_state
// 	msaku_state
// 	cv_number
// 	validity_period
// 	itop_id
// 	order_id
// 	pickup_adress
// 	picked_up_at
// 	trip_purpose
// 	msaku_transaction_id
// 	trip_purposedriver_name
// 	external_order_id
// 	route_image
// 	department_name
// 	account_code
// 	user_name
// 	invoice_number
// 	posting_date
// 	distance
// 	other_information
// }

type Csvdata struct {
	ID       string `json:"id,omitempty"`
	UserName string `json:"user_name"`
}

type Repository interface {
	GetData(ctx context.Context, id string) ([]Csvdata, error)
}
