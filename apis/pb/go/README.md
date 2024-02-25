# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [coordinator.proto](#coordinator.proto)
    - [GetCurationRequest](#coordinator_service.GetCurationRequest)
  
    - [CoordinatorService](#coordinator_service.CoordinatorService)
  
- [entry_point.proto](#entry_point.proto)
    - [EntryPoint](#entry_point.EntryPoint)
  
- [finance.proto](#finance.proto)
    - [Bill](#finance.Bill)
    - [ExperienceTuple](#finance.ExperienceTuple)
    - [Finance](#finance.Finance)
    - [Invoice](#finance.Invoice)
    - [LineItem](#finance.LineItem)
    - [LineItem.AttributesEntry](#finance.LineItem.AttributesEntry)
    - [Payment](#finance.Payment)
    - [PaymentAuthorized](#finance.PaymentAuthorized)
    - [PaymentCaptured](#finance.PaymentCaptured)
    - [PaymentIntent](#finance.PaymentIntent)
    - [PaymentMethod](#finance.PaymentMethod)
    - [PaymentMethodCreated](#finance.PaymentMethodCreated)
    - [PaymentRefunded](#finance.PaymentRefunded)
    - [PaymentUpdated](#finance.PaymentUpdated)
    - [Payroll](#finance.Payroll)
    - [Statement](#finance.Statement)
    - [Transaction](#finance.Transaction)
    - [Transfer](#finance.Transfer)
  
    - [PaymentFailureCode](#finance.PaymentFailureCode)
    - [PaymentMethodType](#finance.PaymentMethodType)
    - [PaymentStatus](#finance.PaymentStatus)
  
- [finance_service.proto](#finance_service.proto)
    - [AddPaymentMethodRequest](#finance_service.AddPaymentMethodRequest)
    - [AddPaymentMethodResponse](#finance_service.AddPaymentMethodResponse)
    - [CaptureFundsRequest](#finance_service.CaptureFundsRequest)
    - [CaptureFundsResponse](#finance_service.CaptureFundsResponse)
    - [CreateRatesRequest](#finance_service.CreateRatesRequest)
    - [CreateRatesResponse](#finance_service.CreateRatesResponse)
    - [GetPaymentMethodsRequest](#finance_service.GetPaymentMethodsRequest)
    - [GetPaymentMethodsResponse](#finance_service.GetPaymentMethodsResponse)
    - [GetPaymentsRequest](#finance_service.GetPaymentsRequest)
    - [GetPaymentsResponse](#finance_service.GetPaymentsResponse)
    - [GetPayrollRequest](#finance_service.GetPayrollRequest)
    - [GetPayrollResponse](#finance_service.GetPayrollResponse)
    - [GetRatesRequest](#finance_service.GetRatesRequest)
    - [GetRatesResponse](#finance_service.GetRatesResponse)
    - [GetTransactionRequest](#finance_service.GetTransactionRequest)
    - [GetTransactionResponse](#finance_service.GetTransactionResponse)
    - [RemovePaymentMethodRequest](#finance_service.RemovePaymentMethodRequest)
    - [RemovePaymentMethodResponse](#finance_service.RemovePaymentMethodResponse)
    - [SaveAuthorizationRequest](#finance_service.SaveAuthorizationRequest)
    - [SaveAuthorizationResponse](#finance_service.SaveAuthorizationResponse)
    - [StartCheckoutRequest](#finance_service.StartCheckoutRequest)
    - [StartCheckoutResponse](#finance_service.StartCheckoutResponse)
  
    - [FinanceService](#finance_service.FinanceService)
  
- [global.proto](#global.proto)
    - [Address](#global.Address)
    - [Amenity](#global.Amenity)
    - [Calendar](#global.Calendar)
    - [Caveat](#global.Caveat)
    - [Comment](#global.Comment)
    - [Confirmation](#global.Confirmation)
    - [Content](#global.Content)
    - [Content.MetadataEntry](#global.Content.MetadataEntry)
    - [ContentPacket](#global.ContentPacket)
    - [Customer](#global.Customer)
    - [Entity](#global.Entity)
    - [Error](#global.Error)
    - [Event](#global.Event)
    - [Experience](#global.Experience)
    - [ExperienceConfirmationsRequest](#global.ExperienceConfirmationsRequest)
    - [ExperienceConfirmationsResponse](#global.ExperienceConfirmationsResponse)
    - [Feature](#global.Feature)
    - [Geo](#global.Geo)
    - [Impression](#global.Impression)
    - [Location](#global.Location)
    - [Note](#global.Note)
    - [Option](#global.Option)
    - [Permission](#global.Permission)
    - [Query](#global.Query)
    - [Rate](#global.Rate)
    - [Rate.MetadataEntry](#global.Rate.MetadataEntry)
    - [Rating](#global.Rating)
    - [Rule](#global.Rule)
    - [Size](#global.Size)
    - [Source](#global.Source)
    - [Space](#global.Space)
    - [Space.AttributesEntry](#global.Space.AttributesEntry)
    - [TaxRate](#global.TaxRate)
    - [TelematicUpdate](#global.TelematicUpdate)
    - [TelematicUpdate.MetadataEntry](#global.TelematicUpdate.MetadataEntry)
    - [Term](#global.Term)
  
    - [ConfirmationStatus](#global.ConfirmationStatus)
    - [ContentType](#global.ContentType)
    - [CreatorType](#global.CreatorType)
    - [Currency](#global.Currency)
    - [Day](#global.Day)
    - [EntityType](#global.EntityType)
    - [ExperienceType](#global.ExperienceType)
    - [LocationType](#global.LocationType)
    - [Order](#global.Order)
    - [RateFrequency](#global.RateFrequency)
    - [RateType](#global.RateType)
    - [ResourceType](#global.ResourceType)
    - [SpaceSubType](#global.SpaceSubType)
    - [SpaceType](#global.SpaceType)
  
- [health_check.proto](#health_check.proto)
    - [HealthCheckRequest](#health_check.HealthCheckRequest)
    - [HealthCheckResponse](#health_check.HealthCheckResponse)
  
    - [HealthCheckResponse.ServingStatus](#health_check.HealthCheckResponse.ServingStatus)
  
    - [Health](#health_check.Health)
  
- [Scalar Value Types](#scalar-value-types)



<a name="coordinator.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## coordinator.proto



<a name="coordinator_service.GetCurationRequest"></a>

### GetCurationRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| query | [global.Query](#global.Query) |  |  |
| page_token | [string](#string) |  | st joun home with 4 beds for a group of 8 ranging between 200 - 500 for the week. |
| page | [int32](#int32) |  |  |
| page_size | [int32](#int32) |  | Example query: &#34;rate_range&#34;: &#34;[500-1000]&#34;, &#34;rate_interval&#34;: &#34;hourly|daily|weekly|monthly&#34;, &#34;rate_currency&#34;: &#34;USD|EUR|GBP|BTC|ETH&#34;, &#34;location_names&#34;: &#34;st. john, st. thomas, jamaica, ...&#34;, &#34;location_geos&#34;: &#34;[18.12345,64.35465],[18.12345,64.35465],...&#34;, &#34;location_raduis&#34;: &#34;50mi,25mi...&#34;, |





 

 

 


<a name="coordinator_service.CoordinatorService"></a>

### CoordinatorService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|

 



<a name="entry_point.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## entry_point.proto


 

 

 


<a name="entry_point.EntryPoint"></a>

### EntryPoint


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|

 



<a name="finance.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## finance.proto



<a name="finance.Bill"></a>

### Bill



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| total | [int32](#int32) |  |  |
| line_items | [LineItem](#finance.LineItem) | repeated |  |






<a name="finance.ExperienceTuple"></a>

### ExperienceTuple



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |
| quantity | [int32](#int32) |  |  |
| unit_price | [int32](#int32) |  |  |
| amount | [int32](#int32) |  |  |






<a name="finance.Finance"></a>

### Finance



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="finance.Invoice"></a>

### Invoice



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| total | [int32](#int32) |  |  |
| line_items | [LineItem](#finance.LineItem) | repeated |  |






<a name="finance.LineItem"></a>

### LineItem



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| account | [string](#string) |  |  |
| code | [string](#string) |  |  |
| quantity | [int32](#int32) |  |  |
| unit_price | [int32](#int32) |  |  |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |
| amount | [int32](#int32) |  |  |
| currency | [global.Currency](#global.Currency) |  |  |
| item_rate | [global.Rate](#global.Rate) |  | TODO: maybe this should just be an item as it already has a type of tax or should/could it be a list of rates? |
| tax_rate | [global.Rate](#global.Rate) |  |  |
| reference | [string](#string) |  | FIXME: ?? reference to what? |
| attributes | [LineItem.AttributesEntry](#finance.LineItem.AttributesEntry) | repeated |  |






<a name="finance.LineItem.AttributesEntry"></a>

### LineItem.AttributesEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="finance.Payment"></a>

### Payment



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| guest_id | [string](#string) |  |  |
| journey_id | [string](#string) |  |  |
| customer_id | [string](#string) |  |  |
| payment_id | [string](#string) |  |  |
| amount | [int64](#int64) |  |  |
| amount_refunded | [int64](#int64) |  |  |
| amount_captured | [int64](#int64) |  |  |
| auth_date | [int64](#int64) |  |  |
| capture_date | [int64](#int64) |  |  |
| refund_date | [int64](#int64) |  |  |
| status | [PaymentStatus](#finance.PaymentStatus) |  |  |
| failure_code | [PaymentFailureCode](#finance.PaymentFailureCode) |  |  |
| failure_message | [string](#string) |  |  |
| payment_intent | [PaymentIntent](#finance.PaymentIntent) |  |  |






<a name="finance.PaymentAuthorized"></a>

### PaymentAuthorized
event source events


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| payment | [Payment](#finance.Payment) |  |  |






<a name="finance.PaymentCaptured"></a>

### PaymentCaptured



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| payment | [Payment](#finance.Payment) |  |  |






<a name="finance.PaymentIntent"></a>

### PaymentIntent



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| amount | [int64](#int64) |  |  |
| amount_capturable | [int64](#int64) |  |  |
| status | [string](#string) |  |  |






<a name="finance.PaymentMethod"></a>

### PaymentMethod



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| guest_id | [string](#string) |  |  |
| customer_id | [string](#string) |  |  |
| vendor_id | [string](#string) |  |  |
| last_four | [string](#string) |  |  |
| brand | [string](#string) |  |  |
| country | [string](#string) |  |  |
| exp_month | [int32](#int32) |  |  |
| exp_year | [int32](#int32) |  |  |
| type | [PaymentMethodType](#finance.PaymentMethodType) |  |  |






<a name="finance.PaymentMethodCreated"></a>

### PaymentMethodCreated



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| payment_method | [PaymentMethod](#finance.PaymentMethod) |  |  |






<a name="finance.PaymentRefunded"></a>

### PaymentRefunded



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| payment | [Payment](#finance.Payment) |  |  |






<a name="finance.PaymentUpdated"></a>

### PaymentUpdated



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| payment | [Payment](#finance.Payment) |  |  |






<a name="finance.Payroll"></a>

### Payroll



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| crew_id | [string](#string) |  |  |






<a name="finance.Statement"></a>

### Statement



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| total | [int32](#int32) |  |  |
| line_items | [LineItem](#finance.LineItem) | repeated |  |






<a name="finance.Transaction"></a>

### Transaction



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| total | [int32](#int32) |  |  |
| line_items | [LineItem](#finance.LineItem) | repeated |  |






<a name="finance.Transfer"></a>

### Transfer






 


<a name="finance.PaymentFailureCode"></a>

### PaymentFailureCode
https://stripe.com/docs/error-codes

| Name | Number | Description |
| ---- | ------ | ----------- |
| PAYMENT_FAILURE_CODE_UNSPECIFIED | 0 | default value, so no failure |
| PAYMENT_FAILURE_CODE_NONE | 1 | default value, so no failure |
| PAYMENT_FAILURE_CODE_OTHER | 2 | not one of the items listed below |
| PAYMENT_FAILURE_CODE_CLOSED | 3 |  |
| PAYMENT_FAILURE_CODE_ACCOUNT_INVALID | 4 |  |
| PAYMENT_FAILURE_CODE_ACCOUNT_NUMBER_INVALID | 5 |  |
| PAYMENT_FAILURE_CODE_ACCOUNT_COUNTRY_INVALID_ADDRESS | 6 |  |
| PAYMENT_FAILURE_CODE_ACCOUNT_INFORMATION_MISMATCH | 7 |  |
| PAYMENT_FAILURE_CODE_AMOUNT_TOO_LARGE | 8 |  |
| PAYMENT_FAILURE_CODE_CAPTURE_CHARGE_AUTHORIZATION_EXPIRE | 9 |  |
| PAYMENT_FAILURE_CODE_CAPTURE_UNAUTHORIZED_PAYMENT | 10 |  |
| PAYMENT_FAILURE_CODE_CARD_DECLINED | 11 |  |
| PAYMENT_FAILURE_CODE_CHARGE_ALREADY_CAPTURED | 12 |  |
| PAYMENT_FAILURE_CODE_CHARGE_ALREADY_REFUNDED | 13 |  |
| PAYMENT_FAILURE_CODE_CHARGE_DISPUTED | 14 |  |
| PAYMENT_FAILURE_CODE_CHARGE_EXPIRED_FOR_CAPTURE | 15 |  |
| PAYMENT_FAILURE_CODE_CHARGE_NOT_REFUNDABLE | 16 |  |
| PAYMENT_FAILURE_CODE_COUNTRY_UNSUPPORTED | 17 |  |
| PAYMENT_FAILURE_CODE_DEBIT_NOT_AUTHORIZED | 18 |  |
| PAYMENT_FAILURE_CODE_EXPIRED_CARD | 19 |  |
| PAYMENT_FAILURE_CODE_INCORRECT_ADDRESS | 20 |  |
| PAYMENT_FAILURE_CODE_INCORRECT_CVC | 21 |  |
| PAYMENT_FAILURE_CODE_INCORRECT_NUMBER | 22 |  |
| PAYMENT_FAILURE_CODE_INCORRECT_ZIP | 23 |  |
| PAYMENT_FAILURE_CODE_INSUFFICIENT_FUNDS | 24 |  |
| PAYMENT_FAILURE_CODE_INVALID_CVC | 25 |  |
| PAYMENT_FAILURE_CODE_INVALID_EXPIRY_MONTH | 26 |  |
| PAYMENT_FAILURE_CODE_INVALID_EXPIRY_YEAR | 27 |  |
| PAYMENT_FAILURE_CODE_INVALID_NUMBER | 28 |  |



<a name="finance.PaymentMethodType"></a>

### PaymentMethodType


| Name | Number | Description |
| ---- | ------ | ----------- |
| PAYMENT_METHOD_TYPE_UNSPECIFIED | 0 |  |
| PAYMENT_METHOD_TYPE_CARD | 1 |  |



<a name="finance.PaymentStatus"></a>

### PaymentStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| PAYMENT_STATUS_UNSPECIFIED | 0 |  |
| PAYMENT_STATUS_PENDING | 1 | has been authorized and awaiting capture |
| PAYMENT_STATUS_SUCCEEDED | 2 | captured |
| PAYMENT_STATUS_FAILED | 3 |  |
| PAYMENT_STATUS_REFUNDED | 4 |  |
| PAYMENT_STATUS_CANCELED | 5 |  |


 

 

 



<a name="finance_service.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## finance_service.proto



<a name="finance_service.AddPaymentMethodRequest"></a>

### AddPaymentMethodRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| payment_method | [finance.PaymentMethod](#finance.PaymentMethod) |  |  |






<a name="finance_service.AddPaymentMethodResponse"></a>

### AddPaymentMethodResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| payment_method | [finance.PaymentMethod](#finance.PaymentMethod) |  |  |






<a name="finance_service.CaptureFundsRequest"></a>

### CaptureFundsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| amount | [int64](#int64) |  |  |






<a name="finance_service.CaptureFundsResponse"></a>

### CaptureFundsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| payment | [finance.Payment](#finance.Payment) |  |  |






<a name="finance_service.CreateRatesRequest"></a>

### CreateRatesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| rates | [global.Rate](#global.Rate) | repeated |  |






<a name="finance_service.CreateRatesResponse"></a>

### CreateRatesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| rate_ids | [string](#string) | repeated |  |






<a name="finance_service.GetPaymentMethodsRequest"></a>

### GetPaymentMethodsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| query | [global.Query](#global.Query) |  |  |






<a name="finance_service.GetPaymentMethodsResponse"></a>

### GetPaymentMethodsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| payment_methods | [finance.PaymentMethod](#finance.PaymentMethod) | repeated |  |






<a name="finance_service.GetPaymentsRequest"></a>

### GetPaymentsRequest







<a name="finance_service.GetPaymentsResponse"></a>

### GetPaymentsResponse







<a name="finance_service.GetPayrollRequest"></a>

### GetPayrollRequest







<a name="finance_service.GetPayrollResponse"></a>

### GetPayrollResponse







<a name="finance_service.GetRatesRequest"></a>

### GetRatesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| query | [global.Query](#global.Query) |  |  |






<a name="finance_service.GetRatesResponse"></a>

### GetRatesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| rates | [global.Rate](#global.Rate) | repeated |  |






<a name="finance_service.GetTransactionRequest"></a>

### GetTransactionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| query | [global.Query](#global.Query) |  |  |






<a name="finance_service.GetTransactionResponse"></a>

### GetTransactionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| transactions | [finance.Transaction](#finance.Transaction) | repeated |  |






<a name="finance_service.RemovePaymentMethodRequest"></a>

### RemovePaymentMethodRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| guest_id | [string](#string) |  |  |
| payment_method_id | [string](#string) |  |  |






<a name="finance_service.RemovePaymentMethodResponse"></a>

### RemovePaymentMethodResponse







<a name="finance_service.SaveAuthorizationRequest"></a>

### SaveAuthorizationRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| guest_id | [string](#string) |  |  |
| payment_id | [string](#string) |  |  |
| journey_id | [string](#string) |  |  |
| amount | [int64](#int64) |  |  |






<a name="finance_service.SaveAuthorizationResponse"></a>

### SaveAuthorizationResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| payment | [finance.Payment](#finance.Payment) |  |  |






<a name="finance_service.StartCheckoutRequest"></a>

### StartCheckoutRequest







<a name="finance_service.StartCheckoutResponse"></a>

### StartCheckoutResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| session_link | [string](#string) |  |  |





 

 

 


<a name="finance_service.FinanceService"></a>

### FinanceService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| StartCheckout | [StartCheckoutRequest](#finance_service.StartCheckoutRequest) | [StartCheckoutResponse](#finance_service.StartCheckoutResponse) |  |
| GetPayments | [GetPaymentsRequest](#finance_service.GetPaymentsRequest) | [GetPaymentsResponse](#finance_service.GetPaymentsResponse) |  |
| AddPaymentMethod | [AddPaymentMethodRequest](#finance_service.AddPaymentMethodRequest) | [AddPaymentMethodResponse](#finance_service.AddPaymentMethodResponse) |  |
| GetPaymentMethods | [GetPaymentMethodsRequest](#finance_service.GetPaymentMethodsRequest) | [GetPaymentMethodsResponse](#finance_service.GetPaymentMethodsResponse) |  |
| RemovePaymentMethod | [RemovePaymentMethodRequest](#finance_service.RemovePaymentMethodRequest) | [RemovePaymentMethodResponse](#finance_service.RemovePaymentMethodResponse) |  |
| SaveAuthorization | [SaveAuthorizationRequest](#finance_service.SaveAuthorizationRequest) | [SaveAuthorizationResponse](#finance_service.SaveAuthorizationResponse) |  |
| CaptureFunds | [CaptureFundsRequest](#finance_service.CaptureFundsRequest) | [CaptureFundsResponse](#finance_service.CaptureFundsResponse) |  |
| GetPayroll | [GetPayrollRequest](#finance_service.GetPayrollRequest) | [GetPayrollResponse](#finance_service.GetPayrollResponse) |  |

 



<a name="global.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## global.proto



<a name="global.Address"></a>

### Address



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| street | [string](#string) |  |  |
| street_2 | [string](#string) |  |  |
| city | [string](#string) |  |  |
| state | [string](#string) |  |  |
| country | [string](#string) |  |  |
| zip | [string](#string) |  |  |






<a name="global.Amenity"></a>

### Amenity



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |
| type | [string](#string) |  |  |






<a name="global.Calendar"></a>

### Calendar
A calendar is a collection of related events, as well as additional 
metadata that defines it. Each  
calendar is identified by an ID which is an email address. Calendars 
can have multiple owners. When a reservation is created a calendar is 
created as well and tied to the reservation.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| owner_id | [string](#string) |  |  |
| resource_id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |
| created_at | [int64](#int64) |  |  |
| updated_at | [int64](#int64) |  |  |
| rules | [Rule](#global.Rule) | repeated |  |
| events | [Event](#global.Event) | repeated | owner_id repeated Guests owners = x; reservation_id rules repeated Event events Days/Slots repeated Guest viewers string time_zone string description |






<a name="global.Caveat"></a>

### Caveat



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |
| start_time | [int64](#int64) |  |  |
| end_time | [int64](#int64) |  |  |






<a name="global.Comment"></a>

### Comment



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| owner_id | [string](#string) |  |  |
| partner_id | [string](#string) |  |  |
| entity_id | [string](#string) |  |  |
| entity_type | [string](#string) |  |  |
| body | [string](#string) |  |  |
| created_at | [string](#string) |  |  |
| updated_at | [string](#string) |  |  |
| deleted_at | [string](#string) |  |  |






<a name="global.Confirmation"></a>

### Confirmation



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| experience_id | [string](#string) |  |  |
| experience_total | [int64](#int64) |  |  |
| share_of_total | [int64](#int64) |  |  |
| payment_method_id | [string](#string) |  |  |
| status | [ConfirmationStatus](#global.ConfirmationStatus) |  |  |
| client_secret | [string](#string) |  |  |
| currency | [Currency](#global.Currency) |  |  |






<a name="global.Content"></a>

### Content



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |
| url | [string](#string) |  |  |
| mime_type | [string](#string) |  |  |
| size | [int64](#int64) |  | in bytes |
| data | [bytes](#bytes) |  |  |
| creator_id | [string](#string) |  |  |
| creator_type | [CreatorType](#global.CreatorType) |  |  |
| created_at | [int64](#int64) |  |  |
| impressions | [Impression](#global.Impression) | repeated |  |
| content_type | [ContentType](#global.ContentType) |  |  |
| metadata | [Content.MetadataEntry](#global.Content.MetadataEntry) | repeated |  |
| associations | [string](#string) | repeated | TODO: work on |






<a name="global.Content.MetadataEntry"></a>

### Content.MetadataEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="global.ContentPacket"></a>

### ContentPacket
TODO: do we need this?


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| current_packets_total | [int64](#int64) |  |  |
| media | [Content](#global.Content) |  | int64 current_packets_sent = 1; |






<a name="global.Customer"></a>

### Customer



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| user_name | [string](#string) |  |  |
| first_name | [string](#string) |  |  |
| last_name | [string](#string) |  |  |
| email | [string](#string) |  |  |
| phone | [string](#string) |  |  |
| date_of_birth | [int64](#int64) |  |  |
| created_at | [int64](#int64) |  |  |
| updated_at | [int64](#int64) |  |  |
| deleted_at | [int64](#int64) |  |  |
| address | [Address](#global.Address) |  | guests can update their addresses |
| profile_image_url | [string](#string) |  |  |
| role_ids | [string](#string) | repeated | add alergies as top line item repeated Preference preferences = 13; Note above the persona property. Persona persona = 14; // Part fo the truth engine string gender = 15; |
| policy_ids | [string](#string) | repeated |  |
| crew_ids | [string](#string) | repeated | TODO: add permssions |






<a name="global.Entity"></a>

### Entity



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | Optional when applied to a rule as an option |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |
| required | [bool](#bool) |  |  |
| type | [EntityType](#global.EntityType) |  | Always required |
| role | [string](#string) |  |  |
| start_time | [int64](#int64) |  | the epoch time in seconds when the rule should start |
| end_time | [int64](#int64) |  | the epoch time in seconds when the rule should end |
| permissions | [Permission](#global.Permission) | repeated | relation |






<a name="global.Error"></a>

### Error



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| code | [string](#string) |  |  |
| message | [string](#string) |  |  |
| timestamp | [int64](#int64) |  |  |






<a name="global.Event"></a>

### Event
An event is an object associated with a specific date or time range.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| primary_calendar | [Calendar](#global.Calendar) |  | The calendar the governs the event. |
| associated_calendars | [Calendar](#global.Calendar) | repeated | Calendars where the event would show i.e. calendars of attending guests.

start_date end_date .... string description Status status repeated Reminder reminders repeated Attatchment attachments int64 start_time = x; int64 end_time = x; Location location = x; Guest organizer = x; repeated Guest guests = x; |






<a name="global.Experience"></a>

### Experience



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| partner_id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |
| type | [ExperienceType](#global.ExperienceType) |  |  |
| media | [Content](#global.Content) | repeated |  |
| start_date | [int64](#int64) |  |  |
| end_date | [int64](#int64) |  |  |
| tags | [string](#string) | repeated |  |
| rules | [Rule](#global.Rule) | repeated |  |
| rates | [Rate](#global.Rate) | repeated |  |
| features | [Feature](#global.Feature) | repeated |  |
| spaces | [Space](#global.Space) | repeated |  |
| impressions | [Impression](#global.Impression) | repeated |  |
| comments | [Comment](#global.Comment) | repeated |  |
| locations | [Location](#global.Location) | repeated |  |
| calendars | [Calendar](#global.Calendar) | repeated |  |
| customers | [Customer](#global.Customer) | repeated |  |






<a name="global.ExperienceConfirmationsRequest"></a>

### ExperienceConfirmationsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| journey_id | [string](#string) |  |  |
| confirmations | [Confirmation](#global.Confirmation) | repeated |  |






<a name="global.ExperienceConfirmationsResponse"></a>

### ExperienceConfirmationsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| journey_id | [string](#string) |  |  |
| confirmations | [Confirmation](#global.Confirmation) | repeated |  |






<a name="global.Feature"></a>

### Feature



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |
| type | [string](#string) |  | service, location, etc, |
| created_at | [string](#string) |  |  |
| updated_at | [string](#string) |  |  |
| deleted_at | [string](#string) |  |  |






<a name="global.Geo"></a>

### Geo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| latitude | [float](#float) |  |  |
| longitude | [float](#float) |  |  |
| altitude | [float](#float) |  |  |
| speed | [float](#float) |  |  |
| heading | [float](#float) |  |  |
| accuracy | [float](#float) |  |  |






<a name="global.Impression"></a>

### Impression
TODO: look in to analytics tracking


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| duration | [int64](#int64) |  |  |
| source | [Source](#global.Source) |  | TODO: or are guests the only sources for impressions? |






<a name="global.Location"></a>

### Location



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |
| address | [Address](#global.Address) |  |  |
| geo | [Geo](#global.Geo) |  |  |
| type | [LocationType](#global.LocationType) |  | TODO: add prev &amp;&amp; next bearfing?? |






<a name="global.Note"></a>

### Note



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| guest_id | [string](#string) |  |  |
| partner_id | [string](#string) |  |  |
| project_id | [string](#string) |  |  |
| title | [string](#string) |  |  |
| body | [string](#string) |  |  |
| created_at | [string](#string) |  |  |
| updated_at | [string](#string) |  |  |
| deleted_at | [string](#string) |  | TODO: implement private notes |






<a name="global.Option"></a>

### Option
boat entity and captain entity


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |
| required | [bool](#bool) |  |  |
| entity | [Entity](#global.Entity) |  |  |






<a name="global.Permission"></a>

### Permission



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |
| allow | [bool](#bool) |  |  |
| start_time | [int64](#int64) |  |  |
| end_time | [int64](#int64) |  |  |






<a name="global.Query"></a>

### Query
Meant to be super flexible to allow for any type of query
still with some level of static typing.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| terms | [Term](#global.Term) | repeated |  |
| page | [int32](#int32) |  |  |
| page_size | [int32](#int32) |  |  |
| order | [Order](#global.Order) |  |  |






<a name="global.Rate"></a>

### Rate



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| partner_id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |
| experiences | [string](#string) | repeated | type:id |
| rate_type | [RateType](#global.RateType) |  |  |
| start_date | [int64](#int64) |  |  |
| end_date | [int64](#int64) |  |  |
| amount | [int64](#int64) |  | allways in cents |
| currency | [Currency](#global.Currency) |  |  |
| frequency | [RateFrequency](#global.RateFrequency) |  | TODO: add number of guests, number of hours, etc. |
| metadata | [Rate.MetadataEntry](#global.Rate.MetadataEntry) | repeated |  |
| created_at | [int64](#int64) |  |  |
| updated_at | [int64](#int64) |  |  |






<a name="global.Rate.MetadataEntry"></a>

### Rate.MetadataEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="global.Rating"></a>

### Rating



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| guest_id | [string](#string) |  |  |
| partner_id | [string](#string) |  |  |
| project_id | [string](#string) |  |  |
| score | [int32](#int32) |  |  |
| comment | [string](#string) |  |  |
| created_at | [string](#string) |  |  |
| updated_at | [string](#string) |  |  |
| deleted_at | [string](#string) |  |  |
| sample | [string](#string) |  | HotelRating hotel_rating = 10; RoomRating room_rating = 11; |






<a name="global.Rule"></a>

### Rule



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| name | [string](#string) |  | boats required |
| description | [string](#string) |  | describes what the rule is for |
| start_time | [int64](#int64) |  | always in milliseconds |
| end_time | [int64](#int64) |  | -1 means never ends |
| principal | [Entity](#global.Entity) |  | applied to |
| options | [Option](#global.Option) | repeated | type?? |






<a name="global.Size"></a>

### Size



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| width | [int32](#int32) |  |  |
| length | [int32](#int32) |  |  |
| height | [int32](#int32) |  |  |
| depth | [int32](#int32) |  |  |






<a name="global.Source"></a>

### Source



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | oneof type { users.Guest guest = 3; |






<a name="global.Space"></a>

### Space



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |
| type | [SpaceType](#global.SpaceType) |  |  |
| sub_type | [SpaceSubType](#global.SpaceSubType) |  |  |
| guest_capacity | [int32](#int32) |  |  |
| size | [Size](#global.Size) |  |  |
| level | [int64](#int64) |  |  |
| media | [Content](#global.Content) | repeated |  |
| amenities | [Amenity](#global.Amenity) | repeated |  |
| spaces | [Space](#global.Space) | repeated | Some spaces have inner inner spaces such as a master bedroom having a bathroom.

* string building_id = 4; string floor_id = 5; int32 floor_level = 6; |
| attributes | [Space.AttributesEntry](#global.Space.AttributesEntry) | repeated | data is used to store the specific data for a space type. |






<a name="global.Space.AttributesEntry"></a>

### Space.AttributesEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="global.TaxRate"></a>

### TaxRate



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | tax props... [country, rate, created_at...] |






<a name="global.TelematicUpdate"></a>

### TelematicUpdate



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| type | [string](#string) |  |  |
| description | [string](#string) |  |  |
| timestamp | [int64](#int64) |  |  |
| location | [Location](#global.Location) |  |  |
| prev_update_id | [string](#string) |  |  |
| next_update_id | [string](#string) |  |  |
| metadata | [TelematicUpdate.MetadataEntry](#global.TelematicUpdate.MetadataEntry) | repeated |  |






<a name="global.TelematicUpdate.MetadataEntry"></a>

### TelematicUpdate.MetadataEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="global.Term"></a>

### Term



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  | TODO: make type an enum: [string, int, float, bool, date, time, datetime, geo, address, phone, email, url, uuid, id]\ |
| value | [string](#string) |  | int64 weight = 4; |





 


<a name="global.ConfirmationStatus"></a>

### ConfirmationStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| CONFIRMATION_STATUS_UNSPECIFIED | 0 |  |
| CONFIRMATION_STATUS_INITIATED | 2 |  |
| CONFIRMATION_STATUS_PENDING | 3 |  |
| CONFIRMATION_STATUS_CONFIRMED | 4 |  |
| CONFIRMATION_STATUS_CANCELLED | 5 |  |



<a name="global.ContentType"></a>

### ContentType


| Name | Number | Description |
| ---- | ------ | ----------- |
| CONTENT_TYPE_UNSPECIFIED | 0 |  |
| CONTENT_TYPE_IMAGE | 1 |  |
| CONTENT_TYPE_VIDEO | 2 |  |
| CONTENT_TYPE_AUDIO | 3 |  |
| CONTENT_TYPE_DOCUMENT | 4 |  |



<a name="global.CreatorType"></a>

### CreatorType
TODO: update to EntityType

| Name | Number | Description |
| ---- | ------ | ----------- |
| CREATOR_TYPE_UNSPECIFIED | 0 |  |
| CREATOR_TYPE_GUEST | 1 |  |
| CREATOR_TYPE_PARTNER | 2 |  |
| CREATOR_TYPE_CONCIERGE | 3 |  |



<a name="global.Currency"></a>

### Currency


| Name | Number | Description |
| ---- | ------ | ----------- |
| CURRENCY_UNSPECIFIED | 0 |  |
| CURRENCY_USD | 1 |  |
| CURRENCY_EUR | 2 |  |
| CURRENCY_XCD | 3 | CURRENCY_GBP = 3; CURRENCY_AUD = 4; CURRENCY_CAD = 5; CURRENCY_CHF = 6; CURRENCY_CNY = 7; CURRENCY_DKK = 8; CURRENCY_HKD = 9; CURRENCY_INR = 10; CURRENCY_JPY = 11; CURRENCY_MXN = 12; CURRENCY_NOK = 13; CURRENCY_NZD = 14; CURRENCY_PLN = 15; CURRENCY_SEK = 16; CURRENCY_SGD = 17; CURRENCY_ZAR = 18; |



<a name="global.Day"></a>

### Day


| Name | Number | Description |
| ---- | ------ | ----------- |
| DAY_UNSPECIFIED | 0 |  |
| DAY_SUNDAY | 1 |  |
| DAY_MONDAY | 2 |  |
| DAY_TUESDAY | 3 |  |
| DAY_WEDNESDAY | 4 |  |
| DAY_THURSDAY | 5 |  |
| DAY_FRIDAY | 6 |  |
| DAY_SATURDAY | 7 |  |



<a name="global.EntityType"></a>

### EntityType


| Name | Number | Description |
| ---- | ------ | ----------- |
| ENTITY_TYPE_UNSPECIFIED | 0 | option, rate, guest, resource, experience, journey, community, etc... |
| ENTITY_TYPE_GUEST | 1 |  |
| ENTITY_TYPE_PARTNER | 2 |  |
| ENTITY_TYPE_EVENT | 3 |  |
| ENTITY_TYPE_PROJECT | 4 |  |
| ENTITY_TYPE_RESOURCE | 5 |  |
| ENTITY_TYPE_NOTIFICATION | 6 |  |
| ENTITY_TYPE_JOURNEY | 7 |  |
| ENTITY_TYPE_EXPERIENCE | 8 |  |
| ENTITY_TYPE_QUALIFICATION | 9 | ENTITY_TYPE_APP = 3; ENTITY_TYPE_CREW = 4; ENTITY_TYPE_CIRCLE = 5; ENTITY_TYPE_POLICY = 6; ENTITY_TYPE_ROLE = 7; ENTITY_TYPE_RESOURCE = 8; ENTITY_TYPE_RULE = 9; ENTITY_TYPE_SESSION = 10; |



<a name="global.ExperienceType"></a>

### ExperienceType


| Name | Number | Description |
| ---- | ------ | ----------- |
| EXPERIENCE_TYPE_UNSPECIFIED | 0 |  |
| EXPERIENCE_TYPE_STAY | 1 |  |
| EXPERIENCE_TYPE_EATS | 2 |  |
| EXPERIENCE_TYPE_ACTIVITY | 3 |  |
| EXPERIENCE_TYPE_EVENT | 4 | EXPERIENCE_TYPE_EXCURSION = 3; |
| EXPERIENCE_TYPE_CAR_RENTAL | 5 |  |
| EXPERIENCE_TYPE_RIDE | 6 | TODO: are flights and rides the same thing? Don&#39;t think people feel that way |
| EXPERIENCE_TYPE_FLIGHT | 7 |  |
| EXPERIENCE_TYPE_SERVICE | 8 | EXPERIENCE_TYPE_YACHT_CHARTER |



<a name="global.LocationType"></a>

### LocationType


| Name | Number | Description |
| ---- | ------ | ----------- |
| LOCATION_TYPE_UNSPECIFIED | 0 |  |
| LOCATION_TYPE_BASE | 1 |  |
| LOCATION_TYPE_IN_TRANSIT | 2 |  |
| LOCATION_TYPE_IN_WAY_POINT | 3 |  |
| LOCATION_TYPE_DESTINATION | 4 |  |
| LOCATION_TYPE_PICKUP | 5 |  |
| LOCATION_TYPE_DROPOFF | 6 |  |
| LOCATION_TYPE_OTHER | 7 |  |



<a name="global.Order"></a>

### Order


| Name | Number | Description |
| ---- | ------ | ----------- |
| ORDER_UNSPECIFIED | 0 |  |
| ORDER_ASC | 1 |  |
| ORDER_DESC | 2 |  |



<a name="global.RateFrequency"></a>

### RateFrequency


| Name | Number | Description |
| ---- | ------ | ----------- |
| RATE_FREQUENCY_UNSPECIFIED | 0 |  |
| RATE_FREQUENCY_ONCE | 1 |  |
| RATE_FREQUENCY_DAILY | 2 |  |
| RATE_FREQUENCY_WEEKLY | 3 |  |
| RATE_FREQUENCY_BIWEEKLY | 4 |  |
| RATE_FREQUENCY_MONTHLY | 5 |  |
| RATE_FREQUENCY_QUARTERLY | 6 |  |
| RATE_FREQUENCY_SEMIANNUALLY | 7 |  |
| RATE_FREQUENCY_ANNUALLY | 8 |  |
| RATE_FREQUENCY_CUSTOM | 9 |  |



<a name="global.RateType"></a>

### RateType


| Name | Number | Description |
| ---- | ------ | ----------- |
| RATE_TYPE_UNSPECIFIED | 0 |  |
| RATE_TYPE_EXPERIENCE | 1 |  |
| RATE_TYPE_TAX | 2 |  |
| RATE_TYPE_DISCOUNT | 3 |  |
| RATE_TYPE_EXPERIENCE_FEE | 4 |  |
| RATE_TYPE_DEPOSIT | 5 |  |
| RATE_TYPE_CREDIT | 6 |  |
| RATE_TYPE_ADJUSTMENT | 7 |  |
| RATE_TYPE_OTHER | 8 |  |



<a name="global.ResourceType"></a>

### ResourceType


| Name | Number | Description |
| ---- | ------ | ----------- |
| RESOURCE_TYPE_UNSPECIFIED | 0 |  |
| RESOURCE_TYPE_PROPERTY | 1 | hotel, venue, home, etc. |
| RESOURCE_TYPE_VESSEL | 2 |  |
| RESOURCE_TYPE_VEHICLE | 3 |  |
| RESOURCE_TYPE_AIRCRAFT | 4 |  |
| RESOURCE_TYPE_EQUIPMENT | 5 |  |
| RESOURCE_TYPE_MEMBER | 6 |  |
| RESOURCE_TYPE_OTHER | 14 |  |



<a name="global.SpaceSubType"></a>

### SpaceSubType


| Name | Number | Description |
| ---- | ------ | ----------- |
| SPACE_SUB_TYPE_UNSPECIFIED | 0 |  |
| SPACE_SUB_TYPE_PARKING | 1 |  |
| SPACE_SUB_TYPE_ROOM | 2 |  |
| SPACE_SUB_TYPE_PENTHOUSE | 3 |  |
| SPACE_SUB_TYPE_ROOF | 4 |  |
| SPACE_SUB_TYPE_GROUND | 5 |  |
| SPACE_SUB_TYPE_BASEMENT | 6 |  |
| SPACE_SUB_TYPE_LOBBY | 7 |  |
| SPACE_SUB_TYPE_STORY | 8 |  |



<a name="global.SpaceType"></a>

### SpaceType


| Name | Number | Description |
| ---- | ------ | ----------- |
| SPACE_TYPE_UNSPECIFIED | 0 |  |
| SPACE_TYPE_BEDROOM | 1 |  |
| SPACE_TYPE_BATHROOM | 2 |  |
| SPACE_TYPE_KITCHEN | 3 |  |
| SPACE_TYPE_LIVING | 4 |  |
| SPACE_TYPE_POOL | 5 |  |
| SPACE_TYPE_AREA | 6 |  |
| SPACE_TYPE_SEAT | 7 |  |
| SPACE_TYPE_LOFT | 8 |  |
| SPACE_TYPE_DECK | 9 |  |


 

 

 



<a name="health_check.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## health_check.proto



<a name="health_check.HealthCheckRequest"></a>

### HealthCheckRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service | [string](#string) |  |  |






<a name="health_check.HealthCheckResponse"></a>

### HealthCheckResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [HealthCheckResponse.ServingStatus](#health_check.HealthCheckResponse.ServingStatus) |  |  |





 


<a name="health_check.HealthCheckResponse.ServingStatus"></a>

### HealthCheckResponse.ServingStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| SERVING_STATUS_UNSPECIFIED | 0 |  |
| SERVING_STATUS_SERVING | 1 |  |
| SERVING_STATUS_NOT_SERVING | 2 |  |
| SERVING_STATUS_SERVICE_UNKNOWN | 3 |  |


 

 


<a name="health_check.Health"></a>

### Health


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Check | [HealthCheckRequest](#health_check.HealthCheckRequest) | [HealthCheckResponse](#health_check.HealthCheckResponse) |  |
| Watch | [HealthCheckRequest](#health_check.HealthCheckRequest) | [HealthCheckResponse](#health_check.HealthCheckResponse) stream |  |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

