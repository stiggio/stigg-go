// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stigg

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/stiggio/stigg-go/internal/apijson"
	"github.com/stiggio/stigg-go/internal/requestconfig"
	"github.com/stiggio/stigg-go/option"
	"github.com/stiggio/stigg-go/packages/param"
)

// V1CustomerPaymentMethodService contains methods and other services that help
// with interacting with the stigg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1CustomerPaymentMethodService] method instead.
type V1CustomerPaymentMethodService struct {
	Options []option.RequestOption
}

// NewV1CustomerPaymentMethodService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV1CustomerPaymentMethodService(opts ...option.RequestOption) (r V1CustomerPaymentMethodService) {
	r = V1CustomerPaymentMethodService{}
	r.Options = opts
	return
}

// Attaches a payment method to a customer for billing. Required for paid
// subscriptions when integrated with a billing provider.
func (r *V1CustomerPaymentMethodService) Attach(ctx context.Context, id string, body V1CustomerPaymentMethodAttachParams, opts ...option.RequestOption) (res *CustomerResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/customers/%s/payment-method", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Removes the payment method from a customer. Ensure active paid subscriptions
// have an alternative payment method.
func (r *V1CustomerPaymentMethodService) Detach(ctx context.Context, id string, opts ...option.RequestOption) (res *CustomerResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/customers/%s/payment-method", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type V1CustomerPaymentMethodAttachParams struct {
	// Integration details
	IntegrationID string `json:"integrationId" api:"required"`
	// Billing provider payment method id
	PaymentMethodID string `json:"paymentMethodId" api:"required"`
	// The vendor identifier of integration
	//
	// Any of "AUTH0", "ZUORA", "STRIPE", "HUBSPOT", "AWS_MARKETPLACE", "SNOWFLAKE",
	// "SALESFORCE", "BIG_QUERY", "OPEN_FGA", "APP_STORE".
	VendorIdentifier V1CustomerPaymentMethodAttachParamsVendorIdentifier `json:"vendorIdentifier,omitzero" api:"required"`
	// Customers selected currency
	//
	// Any of "usd", "aed", "all", "amd", "ang", "aud", "awg", "azn", "bam", "bbd",
	// "bdt", "bgn", "bif", "bmd", "bnd", "bsd", "bwp", "byn", "bzd", "brl", "cad",
	// "cdf", "chf", "cny", "czk", "dkk", "dop", "dzd", "egp", "etb", "eur", "fjd",
	// "gbp", "gel", "gip", "gmd", "gyd", "hkd", "hrk", "htg", "idr", "ils", "inr",
	// "isk", "jmd", "jpy", "kes", "kgs", "khr", "kmf", "krw", "kyd", "kzt", "lbp",
	// "lkr", "lrd", "lsl", "mad", "mdl", "mga", "mkd", "mmk", "mnt", "mop", "mro",
	// "mvr", "mwk", "mxn", "myr", "mzn", "nad", "ngn", "nok", "npr", "nzd", "pgk",
	// "php", "pkr", "pln", "qar", "ron", "rsd", "rub", "rwf", "sar", "sbd", "scr",
	// "sek", "sgd", "sle", "sll", "sos", "szl", "thb", "tjs", "top", "try", "ttd",
	// "tzs", "uah", "uzs", "vnd", "vuv", "wst", "xaf", "xcd", "yer", "zar", "zmw",
	// "clp", "djf", "gnf", "ugx", "pyg", "xof", "xpf".
	BillingCurrency V1CustomerPaymentMethodAttachParamsBillingCurrency `json:"billingCurrency,omitzero"`
	paramObj
}

func (r V1CustomerPaymentMethodAttachParams) MarshalJSON() (data []byte, err error) {
	type shadow V1CustomerPaymentMethodAttachParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CustomerPaymentMethodAttachParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The vendor identifier of integration
type V1CustomerPaymentMethodAttachParamsVendorIdentifier string

const (
	V1CustomerPaymentMethodAttachParamsVendorIdentifierAuth0          V1CustomerPaymentMethodAttachParamsVendorIdentifier = "AUTH0"
	V1CustomerPaymentMethodAttachParamsVendorIdentifierZuora          V1CustomerPaymentMethodAttachParamsVendorIdentifier = "ZUORA"
	V1CustomerPaymentMethodAttachParamsVendorIdentifierStripe         V1CustomerPaymentMethodAttachParamsVendorIdentifier = "STRIPE"
	V1CustomerPaymentMethodAttachParamsVendorIdentifierHubspot        V1CustomerPaymentMethodAttachParamsVendorIdentifier = "HUBSPOT"
	V1CustomerPaymentMethodAttachParamsVendorIdentifierAwsMarketplace V1CustomerPaymentMethodAttachParamsVendorIdentifier = "AWS_MARKETPLACE"
	V1CustomerPaymentMethodAttachParamsVendorIdentifierSnowflake      V1CustomerPaymentMethodAttachParamsVendorIdentifier = "SNOWFLAKE"
	V1CustomerPaymentMethodAttachParamsVendorIdentifierSalesforce     V1CustomerPaymentMethodAttachParamsVendorIdentifier = "SALESFORCE"
	V1CustomerPaymentMethodAttachParamsVendorIdentifierBigQuery       V1CustomerPaymentMethodAttachParamsVendorIdentifier = "BIG_QUERY"
	V1CustomerPaymentMethodAttachParamsVendorIdentifierOpenFga        V1CustomerPaymentMethodAttachParamsVendorIdentifier = "OPEN_FGA"
	V1CustomerPaymentMethodAttachParamsVendorIdentifierAppStore       V1CustomerPaymentMethodAttachParamsVendorIdentifier = "APP_STORE"
)

// Customers selected currency
type V1CustomerPaymentMethodAttachParamsBillingCurrency string

const (
	V1CustomerPaymentMethodAttachParamsBillingCurrencyUsd V1CustomerPaymentMethodAttachParamsBillingCurrency = "usd"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyAed V1CustomerPaymentMethodAttachParamsBillingCurrency = "aed"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyAll V1CustomerPaymentMethodAttachParamsBillingCurrency = "all"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyAmd V1CustomerPaymentMethodAttachParamsBillingCurrency = "amd"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyAng V1CustomerPaymentMethodAttachParamsBillingCurrency = "ang"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyAud V1CustomerPaymentMethodAttachParamsBillingCurrency = "aud"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyAwg V1CustomerPaymentMethodAttachParamsBillingCurrency = "awg"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyAzn V1CustomerPaymentMethodAttachParamsBillingCurrency = "azn"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyBam V1CustomerPaymentMethodAttachParamsBillingCurrency = "bam"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyBbd V1CustomerPaymentMethodAttachParamsBillingCurrency = "bbd"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyBdt V1CustomerPaymentMethodAttachParamsBillingCurrency = "bdt"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyBgn V1CustomerPaymentMethodAttachParamsBillingCurrency = "bgn"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyBif V1CustomerPaymentMethodAttachParamsBillingCurrency = "bif"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyBmd V1CustomerPaymentMethodAttachParamsBillingCurrency = "bmd"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyBnd V1CustomerPaymentMethodAttachParamsBillingCurrency = "bnd"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyBsd V1CustomerPaymentMethodAttachParamsBillingCurrency = "bsd"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyBwp V1CustomerPaymentMethodAttachParamsBillingCurrency = "bwp"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyByn V1CustomerPaymentMethodAttachParamsBillingCurrency = "byn"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyBzd V1CustomerPaymentMethodAttachParamsBillingCurrency = "bzd"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyBrl V1CustomerPaymentMethodAttachParamsBillingCurrency = "brl"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyCad V1CustomerPaymentMethodAttachParamsBillingCurrency = "cad"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyCdf V1CustomerPaymentMethodAttachParamsBillingCurrency = "cdf"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyChf V1CustomerPaymentMethodAttachParamsBillingCurrency = "chf"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyCny V1CustomerPaymentMethodAttachParamsBillingCurrency = "cny"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyCzk V1CustomerPaymentMethodAttachParamsBillingCurrency = "czk"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyDkk V1CustomerPaymentMethodAttachParamsBillingCurrency = "dkk"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyDop V1CustomerPaymentMethodAttachParamsBillingCurrency = "dop"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyDzd V1CustomerPaymentMethodAttachParamsBillingCurrency = "dzd"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyEgp V1CustomerPaymentMethodAttachParamsBillingCurrency = "egp"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyEtb V1CustomerPaymentMethodAttachParamsBillingCurrency = "etb"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyEur V1CustomerPaymentMethodAttachParamsBillingCurrency = "eur"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyFjd V1CustomerPaymentMethodAttachParamsBillingCurrency = "fjd"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyGbp V1CustomerPaymentMethodAttachParamsBillingCurrency = "gbp"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyGel V1CustomerPaymentMethodAttachParamsBillingCurrency = "gel"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyGip V1CustomerPaymentMethodAttachParamsBillingCurrency = "gip"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyGmd V1CustomerPaymentMethodAttachParamsBillingCurrency = "gmd"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyGyd V1CustomerPaymentMethodAttachParamsBillingCurrency = "gyd"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyHkd V1CustomerPaymentMethodAttachParamsBillingCurrency = "hkd"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyHrk V1CustomerPaymentMethodAttachParamsBillingCurrency = "hrk"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyHtg V1CustomerPaymentMethodAttachParamsBillingCurrency = "htg"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyIdr V1CustomerPaymentMethodAttachParamsBillingCurrency = "idr"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyIls V1CustomerPaymentMethodAttachParamsBillingCurrency = "ils"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyInr V1CustomerPaymentMethodAttachParamsBillingCurrency = "inr"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyIsk V1CustomerPaymentMethodAttachParamsBillingCurrency = "isk"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyJmd V1CustomerPaymentMethodAttachParamsBillingCurrency = "jmd"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyJpy V1CustomerPaymentMethodAttachParamsBillingCurrency = "jpy"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyKes V1CustomerPaymentMethodAttachParamsBillingCurrency = "kes"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyKgs V1CustomerPaymentMethodAttachParamsBillingCurrency = "kgs"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyKhr V1CustomerPaymentMethodAttachParamsBillingCurrency = "khr"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyKmf V1CustomerPaymentMethodAttachParamsBillingCurrency = "kmf"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyKrw V1CustomerPaymentMethodAttachParamsBillingCurrency = "krw"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyKyd V1CustomerPaymentMethodAttachParamsBillingCurrency = "kyd"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyKzt V1CustomerPaymentMethodAttachParamsBillingCurrency = "kzt"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyLbp V1CustomerPaymentMethodAttachParamsBillingCurrency = "lbp"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyLkr V1CustomerPaymentMethodAttachParamsBillingCurrency = "lkr"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyLrd V1CustomerPaymentMethodAttachParamsBillingCurrency = "lrd"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyLsl V1CustomerPaymentMethodAttachParamsBillingCurrency = "lsl"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyMad V1CustomerPaymentMethodAttachParamsBillingCurrency = "mad"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyMdl V1CustomerPaymentMethodAttachParamsBillingCurrency = "mdl"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyMga V1CustomerPaymentMethodAttachParamsBillingCurrency = "mga"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyMkd V1CustomerPaymentMethodAttachParamsBillingCurrency = "mkd"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyMmk V1CustomerPaymentMethodAttachParamsBillingCurrency = "mmk"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyMnt V1CustomerPaymentMethodAttachParamsBillingCurrency = "mnt"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyMop V1CustomerPaymentMethodAttachParamsBillingCurrency = "mop"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyMro V1CustomerPaymentMethodAttachParamsBillingCurrency = "mro"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyMvr V1CustomerPaymentMethodAttachParamsBillingCurrency = "mvr"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyMwk V1CustomerPaymentMethodAttachParamsBillingCurrency = "mwk"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyMxn V1CustomerPaymentMethodAttachParamsBillingCurrency = "mxn"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyMyr V1CustomerPaymentMethodAttachParamsBillingCurrency = "myr"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyMzn V1CustomerPaymentMethodAttachParamsBillingCurrency = "mzn"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyNad V1CustomerPaymentMethodAttachParamsBillingCurrency = "nad"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyNgn V1CustomerPaymentMethodAttachParamsBillingCurrency = "ngn"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyNok V1CustomerPaymentMethodAttachParamsBillingCurrency = "nok"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyNpr V1CustomerPaymentMethodAttachParamsBillingCurrency = "npr"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyNzd V1CustomerPaymentMethodAttachParamsBillingCurrency = "nzd"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyPgk V1CustomerPaymentMethodAttachParamsBillingCurrency = "pgk"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyPhp V1CustomerPaymentMethodAttachParamsBillingCurrency = "php"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyPkr V1CustomerPaymentMethodAttachParamsBillingCurrency = "pkr"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyPln V1CustomerPaymentMethodAttachParamsBillingCurrency = "pln"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyQar V1CustomerPaymentMethodAttachParamsBillingCurrency = "qar"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyRon V1CustomerPaymentMethodAttachParamsBillingCurrency = "ron"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyRsd V1CustomerPaymentMethodAttachParamsBillingCurrency = "rsd"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyRub V1CustomerPaymentMethodAttachParamsBillingCurrency = "rub"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyRwf V1CustomerPaymentMethodAttachParamsBillingCurrency = "rwf"
	V1CustomerPaymentMethodAttachParamsBillingCurrencySar V1CustomerPaymentMethodAttachParamsBillingCurrency = "sar"
	V1CustomerPaymentMethodAttachParamsBillingCurrencySbd V1CustomerPaymentMethodAttachParamsBillingCurrency = "sbd"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyScr V1CustomerPaymentMethodAttachParamsBillingCurrency = "scr"
	V1CustomerPaymentMethodAttachParamsBillingCurrencySek V1CustomerPaymentMethodAttachParamsBillingCurrency = "sek"
	V1CustomerPaymentMethodAttachParamsBillingCurrencySgd V1CustomerPaymentMethodAttachParamsBillingCurrency = "sgd"
	V1CustomerPaymentMethodAttachParamsBillingCurrencySle V1CustomerPaymentMethodAttachParamsBillingCurrency = "sle"
	V1CustomerPaymentMethodAttachParamsBillingCurrencySll V1CustomerPaymentMethodAttachParamsBillingCurrency = "sll"
	V1CustomerPaymentMethodAttachParamsBillingCurrencySos V1CustomerPaymentMethodAttachParamsBillingCurrency = "sos"
	V1CustomerPaymentMethodAttachParamsBillingCurrencySzl V1CustomerPaymentMethodAttachParamsBillingCurrency = "szl"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyThb V1CustomerPaymentMethodAttachParamsBillingCurrency = "thb"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyTjs V1CustomerPaymentMethodAttachParamsBillingCurrency = "tjs"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyTop V1CustomerPaymentMethodAttachParamsBillingCurrency = "top"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyTry V1CustomerPaymentMethodAttachParamsBillingCurrency = "try"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyTtd V1CustomerPaymentMethodAttachParamsBillingCurrency = "ttd"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyTzs V1CustomerPaymentMethodAttachParamsBillingCurrency = "tzs"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyUah V1CustomerPaymentMethodAttachParamsBillingCurrency = "uah"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyUzs V1CustomerPaymentMethodAttachParamsBillingCurrency = "uzs"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyVnd V1CustomerPaymentMethodAttachParamsBillingCurrency = "vnd"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyVuv V1CustomerPaymentMethodAttachParamsBillingCurrency = "vuv"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyWst V1CustomerPaymentMethodAttachParamsBillingCurrency = "wst"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyXaf V1CustomerPaymentMethodAttachParamsBillingCurrency = "xaf"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyXcd V1CustomerPaymentMethodAttachParamsBillingCurrency = "xcd"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyYer V1CustomerPaymentMethodAttachParamsBillingCurrency = "yer"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyZar V1CustomerPaymentMethodAttachParamsBillingCurrency = "zar"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyZmw V1CustomerPaymentMethodAttachParamsBillingCurrency = "zmw"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyClp V1CustomerPaymentMethodAttachParamsBillingCurrency = "clp"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyDjf V1CustomerPaymentMethodAttachParamsBillingCurrency = "djf"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyGnf V1CustomerPaymentMethodAttachParamsBillingCurrency = "gnf"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyUgx V1CustomerPaymentMethodAttachParamsBillingCurrency = "ugx"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyPyg V1CustomerPaymentMethodAttachParamsBillingCurrency = "pyg"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyXof V1CustomerPaymentMethodAttachParamsBillingCurrency = "xof"
	V1CustomerPaymentMethodAttachParamsBillingCurrencyXpf V1CustomerPaymentMethodAttachParamsBillingCurrency = "xpf"
)
