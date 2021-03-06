// Package doubleclickbidmanager provides access to the DoubleClick Bid Manager API.
//
// See https://developers.google.com/bid-manager/
//
// Usage example:
//
//   import "google.golang.org/api/doubleclickbidmanager/v1"
//   ...
//   doubleclickbidmanagerService, err := doubleclickbidmanager.New(oauthHttpClient)
package doubleclickbidmanager // import "google.golang.org/api/doubleclickbidmanager/v1"

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/net/context"
	"golang.org/x/net/context/ctxhttp"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/internal"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = internal.MarshalJSON

const apiId = "doubleclickbidmanager:v1"
const apiName = "doubleclickbidmanager"
const apiVersion = "v1"
const basePath = "https://www.googleapis.com/doubleclickbidmanager/v1/"

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Lineitems = NewLineitemsService(s)
	s.Queries = NewQueriesService(s)
	s.Reports = NewReportsService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Lineitems *LineitemsService

	Queries *QueriesService

	Reports *ReportsService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewLineitemsService(s *Service) *LineitemsService {
	rs := &LineitemsService{s: s}
	return rs
}

type LineitemsService struct {
	s *Service
}

func NewQueriesService(s *Service) *QueriesService {
	rs := &QueriesService{s: s}
	return rs
}

type QueriesService struct {
	s *Service
}

func NewReportsService(s *Service) *ReportsService {
	rs := &ReportsService{s: s}
	return rs
}

type ReportsService struct {
	s *Service
}

// DownloadLineItemsRequest: Request to fetch stored line items.
type DownloadLineItemsRequest struct {
	// FilterIds: Ids of the specified filter type used to filter line items
	// to fetch. If omitted, all the line items will be returned.
	FilterIds googleapi.Int64s `json:"filterIds,omitempty"`

	// FilterType: Filter type used to filter line items to fetch.
	//
	// Possible values:
	//   "ADVERTISER_ID"
	//   "INSERTION_ORDER_ID"
	//   "LINE_ITEM_ID"
	FilterType string `json:"filterType,omitempty"`

	// Format: Format in which the line items will be returned. Default to
	// CSV.
	//
	// Possible values:
	//   "CSV"
	Format string `json:"format,omitempty"`

	// ForceSendFields is a list of field names (e.g. "FilterIds") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *DownloadLineItemsRequest) MarshalJSON() ([]byte, error) {
	type noMethod DownloadLineItemsRequest
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// DownloadLineItemsResponse: Download line items response.
type DownloadLineItemsResponse struct {
	// LineItems: Retrieved line items in CSV format. Refer to  Entity Write
	// File Format for more information on file format.
	LineItems string `json:"lineItems,omitempty"`

	// ForceSendFields is a list of field names (e.g. "LineItems") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *DownloadLineItemsResponse) MarshalJSON() ([]byte, error) {
	type noMethod DownloadLineItemsResponse
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// FilterPair: Filter used to match traffic data in your report.
type FilterPair struct {
	// Type: Filter type.
	//
	// Possible values:
	//   "FILTER_ACTIVE_VIEW_EXPECTED_VIEWABILITY"
	//   "FILTER_ACTIVITY_ID"
	//   "FILTER_ADVERTISER"
	//   "FILTER_ADVERTISER_CURRENCY"
	//   "FILTER_ADVERTISER_TIMEZONE"
	//   "FILTER_AD_POSITION"
	//   "FILTER_AGE"
	//   "FILTER_BRANDSAFE_CHANNEL_ID"
	//   "FILTER_BROWSER"
	//   "FILTER_CAMPAIGN_DAILY_FREQUENCY"
	//   "FILTER_CARRIER"
	//   "FILTER_CHANNEL_ID"
	//   "FILTER_CITY"
	//   "FILTER_CONVERSION_DELAY"
	//   "FILTER_COUNTRY"
	//   "FILTER_CREATIVE_ID"
	//   "FILTER_CREATIVE_SIZE"
	//   "FILTER_CREATIVE_TYPE"
	//   "FILTER_DATA_PROVIDER"
	//   "FILTER_DATE"
	//   "FILTER_DAY_OF_WEEK"
	//   "FILTER_DMA"
	//   "FILTER_EXCHANGE_ID"
	//   "FILTER_FLOODLIGHT_PIXEL_ID"
	//   "FILTER_GENDER"
	//   "FILTER_INSERTION_ORDER"
	//   "FILTER_INVENTORY_FORMAT"
	//   "FILTER_INVENTORY_SOURCE"
	//   "FILTER_INVENTORY_SOURCE_TYPE"
	//   "FILTER_KEYWORD"
	//   "FILTER_LINE_ITEM"
	//   "FILTER_LINE_ITEM_DAILY_FREQUENCY"
	//   "FILTER_LINE_ITEM_LIFETIME_FREQUENCY"
	//   "FILTER_LINE_ITEM_TYPE"
	//   "FILTER_MOBILE_DEVICE_MAKE"
	//   "FILTER_MOBILE_DEVICE_MAKE_MODEL"
	//   "FILTER_MOBILE_DEVICE_TYPE"
	//   "FILTER_MOBILE_GEO"
	//   "FILTER_MONTH"
	//   "FILTER_MRAID_SUPPORT"
	//   "FILTER_NIELSEN_AGE"
	//   "FILTER_NIELSEN_COUNTRY_CODE"
	//   "FILTER_NIELSEN_DEVICE_ID"
	//   "FILTER_NIELSEN_GENDER"
	//   "FILTER_ORDER_ID"
	//   "FILTER_OS"
	//   "FILTER_PAGE_CATEGORY"
	//   "FILTER_PAGE_LAYOUT"
	//   "FILTER_PARTNER"
	//   "FILTER_PARTNER_CURRENCY"
	//   "FILTER_PUBLIC_INVENTORY"
	//   "FILTER_QUARTER"
	//   "FILTER_REGION"
	//   "FILTER_REGULAR_CHANNEL_ID"
	//   "FILTER_SITE_ID"
	//   "FILTER_SITE_LANGUAGE"
	//   "FILTER_TARGETED_USER_LIST"
	//   "FILTER_TIME_OF_DAY"
	//   "FILTER_TRUEVIEW_CONVERSION_TYPE"
	//   "FILTER_UNKNOWN"
	//   "FILTER_USER_LIST"
	//   "FILTER_USER_LIST_FIRST_PARTY"
	//   "FILTER_USER_LIST_THIRD_PARTY"
	//   "FILTER_VIDEO_AD_POSITION_IN_STREAM"
	//   "FILTER_VIDEO_COMPANION_SIZE"
	//   "FILTER_VIDEO_COMPANION_TYPE"
	//   "FILTER_VIDEO_CREATIVE_DURATION"
	//   "FILTER_VIDEO_CREATIVE_DURATION_SKIPPABLE"
	//   "FILTER_VIDEO_DURATION_SECONDS"
	//   "FILTER_VIDEO_FORMAT_SUPPORT"
	//   "FILTER_VIDEO_INVENTORY_TYPE"
	//   "FILTER_VIDEO_PLAYER_SIZE"
	//   "FILTER_VIDEO_RATING_TIER"
	//   "FILTER_VIDEO_SKIPPABLE_SUPPORT"
	//   "FILTER_VIDEO_VPAID_SUPPORT"
	//   "FILTER_WEEK"
	//   "FILTER_YEAR"
	//   "FILTER_YOUTUBE_VERTICAL"
	//   "FILTER_ZIP_CODE"
	Type string `json:"type,omitempty"`

	// Value: Filter value.
	Value string `json:"value,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Type") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *FilterPair) MarshalJSON() ([]byte, error) {
	type noMethod FilterPair
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// ListQueriesResponse: List queries response.
type ListQueriesResponse struct {
	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "doubleclickbidmanager#listQueriesResponse".
	Kind string `json:"kind,omitempty"`

	// Queries: Retrieved queries.
	Queries []*Query `json:"queries,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Kind") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *ListQueriesResponse) MarshalJSON() ([]byte, error) {
	type noMethod ListQueriesResponse
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// ListReportsResponse: List reports response.
type ListReportsResponse struct {
	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "doubleclickbidmanager#listReportsResponse".
	Kind string `json:"kind,omitempty"`

	// Reports: Retrieved reports.
	Reports []*Report `json:"reports,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Kind") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *ListReportsResponse) MarshalJSON() ([]byte, error) {
	type noMethod ListReportsResponse
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// Parameters: Parameters of a query or report.
type Parameters struct {
	// Filters: Filters used to match traffic data in your report.
	Filters []*FilterPair `json:"filters,omitempty"`

	// GroupBys: Data is grouped by the filters listed in this field.
	//
	// Possible values:
	//   "FILTER_ACTIVE_VIEW_EXPECTED_VIEWABILITY"
	//   "FILTER_ACTIVITY_ID"
	//   "FILTER_ADVERTISER"
	//   "FILTER_ADVERTISER_CURRENCY"
	//   "FILTER_ADVERTISER_TIMEZONE"
	//   "FILTER_AD_POSITION"
	//   "FILTER_AGE"
	//   "FILTER_BRANDSAFE_CHANNEL_ID"
	//   "FILTER_BROWSER"
	//   "FILTER_CAMPAIGN_DAILY_FREQUENCY"
	//   "FILTER_CARRIER"
	//   "FILTER_CHANNEL_ID"
	//   "FILTER_CITY"
	//   "FILTER_CONVERSION_DELAY"
	//   "FILTER_COUNTRY"
	//   "FILTER_CREATIVE_ID"
	//   "FILTER_CREATIVE_SIZE"
	//   "FILTER_CREATIVE_TYPE"
	//   "FILTER_DATA_PROVIDER"
	//   "FILTER_DATE"
	//   "FILTER_DAY_OF_WEEK"
	//   "FILTER_DMA"
	//   "FILTER_EXCHANGE_ID"
	//   "FILTER_FLOODLIGHT_PIXEL_ID"
	//   "FILTER_GENDER"
	//   "FILTER_INSERTION_ORDER"
	//   "FILTER_INVENTORY_FORMAT"
	//   "FILTER_INVENTORY_SOURCE"
	//   "FILTER_INVENTORY_SOURCE_TYPE"
	//   "FILTER_KEYWORD"
	//   "FILTER_LINE_ITEM"
	//   "FILTER_LINE_ITEM_DAILY_FREQUENCY"
	//   "FILTER_LINE_ITEM_LIFETIME_FREQUENCY"
	//   "FILTER_LINE_ITEM_TYPE"
	//   "FILTER_MOBILE_DEVICE_MAKE"
	//   "FILTER_MOBILE_DEVICE_MAKE_MODEL"
	//   "FILTER_MOBILE_DEVICE_TYPE"
	//   "FILTER_MOBILE_GEO"
	//   "FILTER_MONTH"
	//   "FILTER_MRAID_SUPPORT"
	//   "FILTER_NIELSEN_AGE"
	//   "FILTER_NIELSEN_COUNTRY_CODE"
	//   "FILTER_NIELSEN_DEVICE_ID"
	//   "FILTER_NIELSEN_GENDER"
	//   "FILTER_ORDER_ID"
	//   "FILTER_OS"
	//   "FILTER_PAGE_CATEGORY"
	//   "FILTER_PAGE_LAYOUT"
	//   "FILTER_PARTNER"
	//   "FILTER_PARTNER_CURRENCY"
	//   "FILTER_PUBLIC_INVENTORY"
	//   "FILTER_QUARTER"
	//   "FILTER_REGION"
	//   "FILTER_REGULAR_CHANNEL_ID"
	//   "FILTER_SITE_ID"
	//   "FILTER_SITE_LANGUAGE"
	//   "FILTER_TARGETED_USER_LIST"
	//   "FILTER_TIME_OF_DAY"
	//   "FILTER_TRUEVIEW_CONVERSION_TYPE"
	//   "FILTER_UNKNOWN"
	//   "FILTER_USER_LIST"
	//   "FILTER_USER_LIST_FIRST_PARTY"
	//   "FILTER_USER_LIST_THIRD_PARTY"
	//   "FILTER_VIDEO_AD_POSITION_IN_STREAM"
	//   "FILTER_VIDEO_COMPANION_SIZE"
	//   "FILTER_VIDEO_COMPANION_TYPE"
	//   "FILTER_VIDEO_CREATIVE_DURATION"
	//   "FILTER_VIDEO_CREATIVE_DURATION_SKIPPABLE"
	//   "FILTER_VIDEO_DURATION_SECONDS"
	//   "FILTER_VIDEO_FORMAT_SUPPORT"
	//   "FILTER_VIDEO_INVENTORY_TYPE"
	//   "FILTER_VIDEO_PLAYER_SIZE"
	//   "FILTER_VIDEO_RATING_TIER"
	//   "FILTER_VIDEO_SKIPPABLE_SUPPORT"
	//   "FILTER_VIDEO_VPAID_SUPPORT"
	//   "FILTER_WEEK"
	//   "FILTER_YEAR"
	//   "FILTER_YOUTUBE_VERTICAL"
	//   "FILTER_ZIP_CODE"
	GroupBys []string `json:"groupBys,omitempty"`

	// IncludeInviteData: Whether to include data from Invite Media.
	IncludeInviteData bool `json:"includeInviteData,omitempty"`

	// Metrics: Metrics to include as columns in your report.
	//
	// Possible values:
	//   "METRIC_BID_REQUESTS"
	//   "METRIC_BILLABLE_COST_ADVERTISER"
	//   "METRIC_BILLABLE_COST_PARTNER"
	//   "METRIC_BILLABLE_COST_USD"
	//   "METRIC_CLICKS"
	//   "METRIC_CLICK_TO_POST_CLICK_CONVERSION_RATE"
	//   "METRIC_COMSCORE_VCE_AUDIENCE_AVG_FREQUENCY"
	//   "METRIC_COMSCORE_VCE_AUDIENCE_IMPRESSIONS"
	//   "METRIC_COMSCORE_VCE_AUDIENCE_IMPRESSIONS_SHARE"
	//   "METRIC_COMSCORE_VCE_AUDIENCE_REACH_PCT"
	//   "METRIC_COMSCORE_VCE_AUDIENCE_SHARE_PCT"
	//   "METRIC_COMSCORE_VCE_GROSS_RATING_POINTS"
	//   "METRIC_COMSCORE_VCE_POPULATION"
	//   "METRIC_COMSCORE_VCE_UNIQUE_AUDIENCE"
	//   "METRIC_CONVERSIONS_PER_MILLE"
	//   "METRIC_CPM_FEE1_ADVERTISER"
	//   "METRIC_CPM_FEE1_PARTNER"
	//   "METRIC_CPM_FEE1_USD"
	//   "METRIC_CPM_FEE2_ADVERTISER"
	//   "METRIC_CPM_FEE2_PARTNER"
	//   "METRIC_CPM_FEE2_USD"
	//   "METRIC_CPM_FEE3_ADVERTISER"
	//   "METRIC_CPM_FEE3_PARTNER"
	//   "METRIC_CPM_FEE3_USD"
	//   "METRIC_CPM_FEE4_ADVERTISER"
	//   "METRIC_CPM_FEE4_PARTNER"
	//   "METRIC_CPM_FEE4_USD"
	//   "METRIC_CPM_FEE5_ADVERTISER"
	//   "METRIC_CPM_FEE5_PARTNER"
	//   "METRIC_CPM_FEE5_USD"
	//   "METRIC_CTR"
	//   "METRIC_DATA_COST_ADVERTISER"
	//   "METRIC_DATA_COST_PARTNER"
	//   "METRIC_DATA_COST_USD"
	//   "METRIC_FEE10_ADVERTISER"
	//   "METRIC_FEE10_PARTNER"
	//   "METRIC_FEE10_USD"
	//   "METRIC_FEE11_ADVERTISER"
	//   "METRIC_FEE11_PARTNER"
	//   "METRIC_FEE11_USD"
	//   "METRIC_FEE12_ADVERTISER"
	//   "METRIC_FEE12_PARTNER"
	//   "METRIC_FEE12_USD"
	//   "METRIC_FEE13_ADVERTISER"
	//   "METRIC_FEE13_PARTNER"
	//   "METRIC_FEE13_USD"
	//   "METRIC_FEE14_ADVERTISER"
	//   "METRIC_FEE14_PARTNER"
	//   "METRIC_FEE14_USD"
	//   "METRIC_FEE15_ADVERTISER"
	//   "METRIC_FEE15_PARTNER"
	//   "METRIC_FEE15_USD"
	//   "METRIC_FEE16_ADVERTISER"
	//   "METRIC_FEE16_PARTNER"
	//   "METRIC_FEE16_USD"
	//   "METRIC_FEE17_ADVERTISER"
	//   "METRIC_FEE17_PARTNER"
	//   "METRIC_FEE17_USD"
	//   "METRIC_FEE18_ADVERTISER"
	//   "METRIC_FEE18_PARTNER"
	//   "METRIC_FEE18_USD"
	//   "METRIC_FEE19_ADVERTISER"
	//   "METRIC_FEE19_PARTNER"
	//   "METRIC_FEE19_USD"
	//   "METRIC_FEE20_ADVERTISER"
	//   "METRIC_FEE20_PARTNER"
	//   "METRIC_FEE20_USD"
	//   "METRIC_FEE2_ADVERTISER"
	//   "METRIC_FEE2_PARTNER"
	//   "METRIC_FEE2_USD"
	//   "METRIC_FEE3_ADVERTISER"
	//   "METRIC_FEE3_PARTNER"
	//   "METRIC_FEE3_USD"
	//   "METRIC_FEE4_ADVERTISER"
	//   "METRIC_FEE4_PARTNER"
	//   "METRIC_FEE4_USD"
	//   "METRIC_FEE5_ADVERTISER"
	//   "METRIC_FEE5_PARTNER"
	//   "METRIC_FEE5_USD"
	//   "METRIC_FEE6_ADVERTISER"
	//   "METRIC_FEE6_PARTNER"
	//   "METRIC_FEE6_USD"
	//   "METRIC_FEE7_ADVERTISER"
	//   "METRIC_FEE7_PARTNER"
	//   "METRIC_FEE7_USD"
	//   "METRIC_FEE8_ADVERTISER"
	//   "METRIC_FEE8_PARTNER"
	//   "METRIC_FEE8_USD"
	//   "METRIC_FEE9_ADVERTISER"
	//   "METRIC_FEE9_PARTNER"
	//   "METRIC_FEE9_USD"
	//   "METRIC_IMPRESSIONS"
	//   "METRIC_IMPRESSIONS_TO_CONVERSION_RATE"
	//   "METRIC_LAST_CLICKS"
	//   "METRIC_LAST_IMPRESSIONS"
	//   "METRIC_MEDIA_COST_ADVERTISER"
	//   "METRIC_MEDIA_COST_ECPAPC_ADVERTISER"
	//   "METRIC_MEDIA_COST_ECPAPC_PARTNER"
	//   "METRIC_MEDIA_COST_ECPAPC_USD"
	//   "METRIC_MEDIA_COST_ECPAPV_ADVERTISER"
	//   "METRIC_MEDIA_COST_ECPAPV_PARTNER"
	//   "METRIC_MEDIA_COST_ECPAPV_USD"
	//   "METRIC_MEDIA_COST_ECPA_ADVERTISER"
	//   "METRIC_MEDIA_COST_ECPA_PARTNER"
	//   "METRIC_MEDIA_COST_ECPA_USD"
	//   "METRIC_MEDIA_COST_ECPCV_ADVERTISER"
	//   "METRIC_MEDIA_COST_ECPCV_PARTNER"
	//   "METRIC_MEDIA_COST_ECPCV_USD"
	//   "METRIC_MEDIA_COST_ECPC_ADVERTISER"
	//   "METRIC_MEDIA_COST_ECPC_PARTNER"
	//   "METRIC_MEDIA_COST_ECPC_USD"
	//   "METRIC_MEDIA_COST_ECPM_ADVERTISER"
	//   "METRIC_MEDIA_COST_ECPM_PARTNER"
	//   "METRIC_MEDIA_COST_ECPM_USD"
	//   "METRIC_MEDIA_COST_PARTNER"
	//   "METRIC_MEDIA_COST_USD"
	//   "METRIC_MEDIA_FEE1_ADVERTISER"
	//   "METRIC_MEDIA_FEE1_PARTNER"
	//   "METRIC_MEDIA_FEE1_USD"
	//   "METRIC_MEDIA_FEE2_ADVERTISER"
	//   "METRIC_MEDIA_FEE2_PARTNER"
	//   "METRIC_MEDIA_FEE2_USD"
	//   "METRIC_MEDIA_FEE3_ADVERTISER"
	//   "METRIC_MEDIA_FEE3_PARTNER"
	//   "METRIC_MEDIA_FEE3_USD"
	//   "METRIC_MEDIA_FEE4_ADVERTISER"
	//   "METRIC_MEDIA_FEE4_PARTNER"
	//   "METRIC_MEDIA_FEE4_USD"
	//   "METRIC_MEDIA_FEE5_ADVERTISER"
	//   "METRIC_MEDIA_FEE5_PARTNER"
	//   "METRIC_MEDIA_FEE5_USD"
	//   "METRIC_PIXEL_LOADS"
	//   "METRIC_PLATFORM_FEE_ADVERTISER"
	//   "METRIC_PLATFORM_FEE_PARTNER"
	//   "METRIC_PLATFORM_FEE_USD"
	//   "METRIC_POST_CLICK_DFA_REVENUE"
	//   "METRIC_POST_VIEW_DFA_REVENUE"
	//   "METRIC_PROFIT_ADVERTISER"
	//   "METRIC_PROFIT_ECPAPC_ADVERTISER"
	//   "METRIC_PROFIT_ECPAPC_PARTNER"
	//   "METRIC_PROFIT_ECPAPC_USD"
	//   "METRIC_PROFIT_ECPAPV_ADVERTISER"
	//   "METRIC_PROFIT_ECPAPV_PARTNER"
	//   "METRIC_PROFIT_ECPAPV_USD"
	//   "METRIC_PROFIT_ECPA_ADVERTISER"
	//   "METRIC_PROFIT_ECPA_PARTNER"
	//   "METRIC_PROFIT_ECPA_USD"
	//   "METRIC_PROFIT_ECPC_ADVERTISER"
	//   "METRIC_PROFIT_ECPC_PARTNER"
	//   "METRIC_PROFIT_ECPC_USD"
	//   "METRIC_PROFIT_ECPM_ADVERTISER"
	//   "METRIC_PROFIT_ECPM_PARTNER"
	//   "METRIC_PROFIT_ECPM_USD"
	//   "METRIC_PROFIT_MARGIN"
	//   "METRIC_PROFIT_PARTNER"
	//   "METRIC_PROFIT_USD"
	//   "METRIC_REVENUE_ADVERTISER"
	//   "METRIC_REVENUE_ECPAPC_ADVERTISER"
	//   "METRIC_REVENUE_ECPAPC_PARTNER"
	//   "METRIC_REVENUE_ECPAPC_USD"
	//   "METRIC_REVENUE_ECPAPV_ADVERTISER"
	//   "METRIC_REVENUE_ECPAPV_PARTNER"
	//   "METRIC_REVENUE_ECPAPV_USD"
	//   "METRIC_REVENUE_ECPA_ADVERTISER"
	//   "METRIC_REVENUE_ECPA_PARTNER"
	//   "METRIC_REVENUE_ECPA_USD"
	//   "METRIC_REVENUE_ECPCV_ADVERTISER"
	//   "METRIC_REVENUE_ECPCV_PARTNER"
	//   "METRIC_REVENUE_ECPCV_USD"
	//   "METRIC_REVENUE_ECPC_ADVERTISER"
	//   "METRIC_REVENUE_ECPC_PARTNER"
	//   "METRIC_REVENUE_ECPC_USD"
	//   "METRIC_REVENUE_ECPM_ADVERTISER"
	//   "METRIC_REVENUE_ECPM_PARTNER"
	//   "METRIC_REVENUE_ECPM_USD"
	//   "METRIC_REVENUE_PARTNER"
	//   "METRIC_REVENUE_USD"
	//   "METRIC_RICH_MEDIA_VIDEO_COMPLETIONS"
	//   "METRIC_RICH_MEDIA_VIDEO_FIRST_QUARTILE_COMPLETES"
	//   "METRIC_RICH_MEDIA_VIDEO_FULL_SCREENS"
	//   "METRIC_RICH_MEDIA_VIDEO_MIDPOINTS"
	//   "METRIC_RICH_MEDIA_VIDEO_MUTES"
	//   "METRIC_RICH_MEDIA_VIDEO_PAUSES"
	//   "METRIC_RICH_MEDIA_VIDEO_PLAYS"
	//   "METRIC_RICH_MEDIA_VIDEO_SKIPS"
	//   "METRIC_RICH_MEDIA_VIDEO_THIRD_QUARTILE_COMPLETES"
	//   "METRIC_TEA_TRUEVIEW_IMPRESSIONS"
	//   "METRIC_TEA_TRUEVIEW_UNIQUE_COOKIES"
	//   "METRIC_TEA_TRUEVIEW_UNIQUE_PEOPLE"
	//   "METRIC_TOTAL_CONVERSIONS"
	//   "METRIC_TOTAL_MEDIA_COST_ADVERTISER"
	//   "METRIC_TOTAL_MEDIA_COST_ECPAPC_ADVERTISER"
	//   "METRIC_TOTAL_MEDIA_COST_ECPAPC_PARTNER"
	//   "METRIC_TOTAL_MEDIA_COST_ECPAPC_USD"
	//   "METRIC_TOTAL_MEDIA_COST_ECPAPV_ADVERTISER"
	//   "METRIC_TOTAL_MEDIA_COST_ECPAPV_PARTNER"
	//   "METRIC_TOTAL_MEDIA_COST_ECPAPV_USD"
	//   "METRIC_TOTAL_MEDIA_COST_ECPA_ADVERTISER"
	//   "METRIC_TOTAL_MEDIA_COST_ECPA_PARTNER"
	//   "METRIC_TOTAL_MEDIA_COST_ECPA_USD"
	//   "METRIC_TOTAL_MEDIA_COST_ECPCV_ADVERTISER"
	//   "METRIC_TOTAL_MEDIA_COST_ECPCV_PARTNER"
	//   "METRIC_TOTAL_MEDIA_COST_ECPCV_USD"
	//   "METRIC_TOTAL_MEDIA_COST_ECPC_ADVERTISER"
	//   "METRIC_TOTAL_MEDIA_COST_ECPC_PARTNER"
	//   "METRIC_TOTAL_MEDIA_COST_ECPC_USD"
	//   "METRIC_TOTAL_MEDIA_COST_ECPM_ADVERTISER"
	//   "METRIC_TOTAL_MEDIA_COST_ECPM_PARTNER"
	//   "METRIC_TOTAL_MEDIA_COST_ECPM_USD"
	//   "METRIC_TOTAL_MEDIA_COST_PARTNER"
	//   "METRIC_TOTAL_MEDIA_COST_USD"
	//   "METRIC_TRUEVIEW_CONVERSION_MANY_PER_VIEW"
	//   "METRIC_TRUEVIEW_CONVERSION_ONE_PER_VIEW"
	//   "METRIC_TRUEVIEW_CONVERSION_RATE_ONE_PER_VIEW"
	//   "METRIC_TRUEVIEW_COST_CONVERSION_MANY_PER_VIEW_RATIO"
	//   "METRIC_TRUEVIEW_COST_CONVERSION_ONE_PER_VIEW_RATIO"
	//   "METRIC_TRUEVIEW_CPV_ADVERTISER"
	//   "METRIC_TRUEVIEW_CPV_PARTNER"
	//   "METRIC_TRUEVIEW_CPV_USD"
	//   "METRIC_TRUEVIEW_EARNED_LIKES"
	//   "METRIC_TRUEVIEW_EARNED_PLAYLIST_ADDITIONS"
	//   "METRIC_TRUEVIEW_EARNED_SHARES"
	//   "METRIC_TRUEVIEW_EARNED_SUBSCRIBERS"
	//   "METRIC_TRUEVIEW_EARNED_VIEWS"
	//   "METRIC_TRUEVIEW_IMPRESSION_SHARE"
	//   "METRIC_TRUEVIEW_LOST_IS_BUDGET"
	//   "METRIC_TRUEVIEW_LOST_IS_RANK"
	//   "METRIC_TRUEVIEW_TOTAL_CONVERSION_VALUE"
	//   "METRIC_TRUEVIEW_UNIQUE_VIEWERS"
	//   "METRIC_TRUEVIEW_VALUE_CONVERSION_MANY_PER_VIEW_RATIO"
	//   "METRIC_TRUEVIEW_VALUE_CONVERSION_ONE_PER_VIEW_RATIO"
	//   "METRIC_TRUEVIEW_VIEWS"
	//   "METRIC_TRUEVIEW_VIEW_RATE"
	//   "METRIC_TRUEVIEW_VIEW_THROUGH_CONVERSION"
	//   "METRIC_UNIQUE_VISITORS_COOKIES"
	//   "METRIC_UNKNOWN"
	//   "METRIC_VIDEO_COMPANION_CLICKS"
	//   "METRIC_VIDEO_COMPANION_IMPRESSIONS"
	//   "METRIC_VIDEO_COMPLETION_RATE"
	Metrics []string `json:"metrics,omitempty"`

	// Type: Report type.
	//
	// Possible values:
	//   "TYPE_ACTIVE_GRP"
	//   "TYPE_AUDIENCE_COMPOSITION"
	//   "TYPE_AUDIENCE_PERFORMANCE"
	//   "TYPE_CLIENT_SAFE"
	//   "TYPE_COMSCORE_VCE"
	//   "TYPE_CROSS_FEE"
	//   "TYPE_CROSS_PARTNER"
	//   "TYPE_CROSS_PARTNER_THIRD_PARTY_DATA_PROVIDER"
	//   "TYPE_ESTIMATED_CONVERSION"
	//   "TYPE_FEE"
	//   "TYPE_GENERAL"
	//   "TYPE_INVENTORY_AVAILABILITY"
	//   "TYPE_KEYWORD"
	//   "TYPE_NIELSEN_AUDIENCE_PROFILE"
	//   "TYPE_NIELSEN_DAILY_REACH_BUILD"
	//   "TYPE_NIELSEN_SITE"
	//   "TYPE_ORDER_ID"
	//   "TYPE_PAGE_CATEGORY"
	//   "TYPE_PIXEL_LOAD"
	//   "TYPE_REACH_AND_FREQUENCY"
	//   "TYPE_THIRD_PARTY_DATA_PROVIDER"
	//   "TYPE_TRUEVIEW"
	//   "TYPE_TRUEVIEW_IAR"
	//   "TYPE_VERIFICATION"
	//   "TYPE_YOUTUBE_VERTICAL"
	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Filters") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Parameters) MarshalJSON() ([]byte, error) {
	type noMethod Parameters
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// Query: Represents a query.
type Query struct {
	// Kind: Identifies what kind of resource this is. Value: the fixed
	// string "doubleclickbidmanager#query".
	Kind string `json:"kind,omitempty"`

	// Metadata: Query metadata.
	Metadata *QueryMetadata `json:"metadata,omitempty"`

	// Params: Query parameters.
	Params *Parameters `json:"params,omitempty"`

	// QueryId: Query ID.
	QueryId int64 `json:"queryId,omitempty,string"`

	// ReportDataEndTimeMs: The ending time for the data that is shown in
	// the report. Note, reportDataEndTimeMs is required if
	// metadata.dataRange is CUSTOM_DATES and ignored otherwise.
	ReportDataEndTimeMs int64 `json:"reportDataEndTimeMs,omitempty,string"`

	// ReportDataStartTimeMs: The starting time for the data that is shown
	// in the report. Note, reportDataStartTimeMs is required if
	// metadata.dataRange is CUSTOM_DATES and ignored otherwise.
	ReportDataStartTimeMs int64 `json:"reportDataStartTimeMs,omitempty,string"`

	// Schedule: Information on how often and when to run a query.
	Schedule *QuerySchedule `json:"schedule,omitempty"`

	// TimezoneCode: Canonical timezone code for report data time. Defaults
	// to America/New_York.
	TimezoneCode string `json:"timezoneCode,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Kind") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Query) MarshalJSON() ([]byte, error) {
	type noMethod Query
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// QueryMetadata: Query metadata.
type QueryMetadata struct {
	// DataRange: Range of report data.
	//
	// Possible values:
	//   "ALL_TIME"
	//   "CURRENT_DAY"
	//   "CUSTOM_DATES"
	//   "LAST_14_DAYS"
	//   "LAST_30_DAYS"
	//   "LAST_365_DAYS"
	//   "LAST_7_DAYS"
	//   "LAST_90_DAYS"
	//   "MONTH_TO_DATE"
	//   "PREVIOUS_DAY"
	//   "PREVIOUS_HALF_MONTH"
	//   "PREVIOUS_MONTH"
	//   "PREVIOUS_QUARTER"
	//   "PREVIOUS_WEEK"
	//   "PREVIOUS_YEAR"
	//   "QUARTER_TO_DATE"
	//   "WEEK_TO_DATE"
	//   "YEAR_TO_DATE"
	DataRange string `json:"dataRange,omitempty"`

	// Format: Format of the generated report.
	//
	// Possible values:
	//   "CSV"
	//   "EXCEL_CSV"
	//   "XLSX"
	Format string `json:"format,omitempty"`

	// GoogleCloudStoragePathForLatestReport: The path to the location in
	// Google Cloud Storage where the latest report is stored.
	GoogleCloudStoragePathForLatestReport string `json:"googleCloudStoragePathForLatestReport,omitempty"`

	// GoogleDrivePathForLatestReport: The path in Google Drive for the
	// latest report.
	GoogleDrivePathForLatestReport string `json:"googleDrivePathForLatestReport,omitempty"`

	// LatestReportRunTimeMs: The time when the latest report started to
	// run.
	LatestReportRunTimeMs int64 `json:"latestReportRunTimeMs,omitempty,string"`

	// Locale: Locale of the generated reports. Valid values are cs CZECH de
	// GERMAN en ENGLISH es SPANISH fr FRENCH it ITALIAN ja JAPANESE ko
	// KOREAN pl POLISH pt-BR BRAZILIAN_PORTUGUESE ru RUSSIAN tr TURKISH uk
	// UKRAINIAN zh-CN CHINA_CHINESE zh-TW TAIWAN_CHINESE
	//
	// An locale string not in the list above will generate reports in
	// English.
	Locale string `json:"locale,omitempty"`

	// ReportCount: Number of reports that have been generated for the
	// query.
	ReportCount int64 `json:"reportCount,omitempty"`

	// Running: Whether the latest report is currently running.
	Running bool `json:"running,omitempty"`

	// SendNotification: Whether to send an email notification when a report
	// is ready. Default to false.
	SendNotification bool `json:"sendNotification,omitempty"`

	// ShareEmailAddress: List of email addresses which are sent email
	// notifications when the report is finished. Separate from
	// sendNotification.
	ShareEmailAddress []string `json:"shareEmailAddress,omitempty"`

	// Title: Query title. It is used to name the reports generated from
	// this query.
	Title string `json:"title,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DataRange") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *QueryMetadata) MarshalJSON() ([]byte, error) {
	type noMethod QueryMetadata
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// QuerySchedule: Information on how frequently and when to run a query.
type QuerySchedule struct {
	// EndTimeMs: Datetime to periodically run the query until.
	EndTimeMs int64 `json:"endTimeMs,omitempty,string"`

	// Frequency: How often the query is run.
	//
	// Possible values:
	//   "DAILY"
	//   "MONTHLY"
	//   "ONE_TIME"
	//   "QUARTERLY"
	//   "SEMI_MONTHLY"
	//   "WEEKLY"
	Frequency string `json:"frequency,omitempty"`

	// NextRunMinuteOfDay: Time of day at which a new report will be
	// generated, represented as minutes past midnight. Range is 0 to 1439.
	// Only applies to scheduled reports.
	NextRunMinuteOfDay int64 `json:"nextRunMinuteOfDay,omitempty"`

	// NextRunTimezoneCode: Canonical timezone code for report generation
	// time. Defaults to America/New_York.
	NextRunTimezoneCode string `json:"nextRunTimezoneCode,omitempty"`

	// ForceSendFields is a list of field names (e.g. "EndTimeMs") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *QuerySchedule) MarshalJSON() ([]byte, error) {
	type noMethod QuerySchedule
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// Report: Represents a report.
type Report struct {
	// Key: Key used to identify a report.
	Key *ReportKey `json:"key,omitempty"`

	// Metadata: Report metadata.
	Metadata *ReportMetadata `json:"metadata,omitempty"`

	// Params: Report parameters.
	Params *Parameters `json:"params,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Key") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Report) MarshalJSON() ([]byte, error) {
	type noMethod Report
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// ReportFailure: An explanation of a report failure.
type ReportFailure struct {
	// ErrorCode: Error code that shows why the report was not created.
	//
	// Possible values:
	//   "AUTHENTICATION_ERROR"
	//   "DEPRECATED_REPORTING_INVALID_QUERY"
	//   "REPORTING_BUCKET_NOT_FOUND"
	//   "REPORTING_CREATE_BUCKET_FAILED"
	//   "REPORTING_DELETE_BUCKET_FAILED"
	//   "REPORTING_FATAL_ERROR"
	//   "REPORTING_ILLEGAL_FILENAME"
	//   "REPORTING_IMCOMPATIBLE_METRICS"
	//   "REPORTING_INVALID_QUERY_MISSING_PARTNER_AND_ADVERTISER_FILTERS"
	//   "REPORTING_INVALID_QUERY_TITLE_MISSING"
	//   "REPORTING_INVALID_QUERY_TOO_MANY_UNFILTERED_LARGE_GROUP_BYS"
	//   "REPORTING_QUERY_NOT_FOUND"
	//   "REPORTING_TRANSIENT_ERROR"
	//   "REPORTING_UPDATE_BUCKET_PERMISSION_FAILED"
	//   "REPORTING_WRITE_BUCKET_OBJECT_FAILED"
	//   "SERVER_ERROR"
	//   "UNAUTHORIZED_API_ACCESS"
	//   "VALIDATION_ERROR"
	ErrorCode string `json:"errorCode,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ErrorCode") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *ReportFailure) MarshalJSON() ([]byte, error) {
	type noMethod ReportFailure
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// ReportKey: Key used to identify a report.
type ReportKey struct {
	// QueryId: Query ID.
	QueryId int64 `json:"queryId,omitempty,string"`

	// ReportId: Report ID.
	ReportId int64 `json:"reportId,omitempty,string"`

	// ForceSendFields is a list of field names (e.g. "QueryId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *ReportKey) MarshalJSON() ([]byte, error) {
	type noMethod ReportKey
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// ReportMetadata: Report metadata.
type ReportMetadata struct {
	// GoogleCloudStoragePath: The path to the location in Google Cloud
	// Storage where the report is stored.
	GoogleCloudStoragePath string `json:"googleCloudStoragePath,omitempty"`

	// ReportDataEndTimeMs: The ending time for the data that is shown in
	// the report.
	ReportDataEndTimeMs int64 `json:"reportDataEndTimeMs,omitempty,string"`

	// ReportDataStartTimeMs: The starting time for the data that is shown
	// in the report.
	ReportDataStartTimeMs int64 `json:"reportDataStartTimeMs,omitempty,string"`

	// Status: Report status.
	Status *ReportStatus `json:"status,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "GoogleCloudStoragePath") to unconditionally include in API requests.
	// By default, fields with empty values are omitted from API requests.
	// However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *ReportMetadata) MarshalJSON() ([]byte, error) {
	type noMethod ReportMetadata
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// ReportStatus: Report status.
type ReportStatus struct {
	// Failure: If the report failed, this records the cause.
	Failure *ReportFailure `json:"failure,omitempty"`

	// FinishTimeMs: The time when this report either completed successfully
	// or failed.
	FinishTimeMs int64 `json:"finishTimeMs,omitempty,string"`

	// Format: The file type of the report.
	//
	// Possible values:
	//   "CSV"
	//   "EXCEL_CSV"
	//   "XLSX"
	Format string `json:"format,omitempty"`

	// State: The state of the report.
	//
	// Possible values:
	//   "DONE"
	//   "FAILED"
	//   "RUNNING"
	State string `json:"state,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Failure") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *ReportStatus) MarshalJSON() ([]byte, error) {
	type noMethod ReportStatus
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// RowStatus: Represents the upload status of a row in the request.
type RowStatus struct {
	// Changed: Whether the stored entity is changed as a result of upload.
	Changed bool `json:"changed,omitempty"`

	// EntityId: Entity Id.
	EntityId int64 `json:"entityId,omitempty,string"`

	// EntityName: Entity name.
	EntityName string `json:"entityName,omitempty"`

	// Errors: Reasons why the entity can't be uploaded.
	Errors []string `json:"errors,omitempty"`

	// Persisted: Whether the entity is persisted.
	Persisted bool `json:"persisted,omitempty"`

	// RowNumber: Row number.
	RowNumber int64 `json:"rowNumber,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Changed") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *RowStatus) MarshalJSON() ([]byte, error) {
	type noMethod RowStatus
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// RunQueryRequest: Request to run a stored query to generate a report.
type RunQueryRequest struct {
	// DataRange: Report data range used to generate the report.
	//
	// Possible values:
	//   "ALL_TIME"
	//   "CURRENT_DAY"
	//   "CUSTOM_DATES"
	//   "LAST_14_DAYS"
	//   "LAST_30_DAYS"
	//   "LAST_365_DAYS"
	//   "LAST_7_DAYS"
	//   "LAST_90_DAYS"
	//   "MONTH_TO_DATE"
	//   "PREVIOUS_DAY"
	//   "PREVIOUS_HALF_MONTH"
	//   "PREVIOUS_MONTH"
	//   "PREVIOUS_QUARTER"
	//   "PREVIOUS_WEEK"
	//   "PREVIOUS_YEAR"
	//   "QUARTER_TO_DATE"
	//   "WEEK_TO_DATE"
	//   "YEAR_TO_DATE"
	DataRange string `json:"dataRange,omitempty"`

	// ReportDataEndTimeMs: The ending time for the data that is shown in
	// the report. Note, reportDataEndTimeMs is required if dataRange is
	// CUSTOM_DATES and ignored otherwise.
	ReportDataEndTimeMs int64 `json:"reportDataEndTimeMs,omitempty,string"`

	// ReportDataStartTimeMs: The starting time for the data that is shown
	// in the report. Note, reportDataStartTimeMs is required if dataRange
	// is CUSTOM_DATES and ignored otherwise.
	ReportDataStartTimeMs int64 `json:"reportDataStartTimeMs,omitempty,string"`

	// TimezoneCode: Canonical timezone code for report data time. Defaults
	// to America/New_York.
	TimezoneCode string `json:"timezoneCode,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DataRange") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *RunQueryRequest) MarshalJSON() ([]byte, error) {
	type noMethod RunQueryRequest
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// UploadLineItemsRequest: Request to upload line items.
type UploadLineItemsRequest struct {
	// DryRun: Set to true to get upload status without actually persisting
	// the line items.
	DryRun bool `json:"dryRun,omitempty"`

	// Format: Format the line items are in. Default to CSV.
	//
	// Possible values:
	//   "CSV"
	Format string `json:"format,omitempty"`

	// LineItems: Line items in CSV to upload. Refer to  Entity Write File
	// Format for more information on file format.
	LineItems string `json:"lineItems,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DryRun") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *UploadLineItemsRequest) MarshalJSON() ([]byte, error) {
	type noMethod UploadLineItemsRequest
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// UploadLineItemsResponse: Upload line items response.
type UploadLineItemsResponse struct {
	// UploadStatus: Status of upload.
	UploadStatus *UploadStatus `json:"uploadStatus,omitempty"`

	// ForceSendFields is a list of field names (e.g. "UploadStatus") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *UploadLineItemsResponse) MarshalJSON() ([]byte, error) {
	type noMethod UploadLineItemsResponse
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// UploadStatus: Represents the status of upload.
type UploadStatus struct {
	// Errors: Reasons why upload can't be completed.
	Errors []string `json:"errors,omitempty"`

	// RowStatus: Per-row upload status.
	RowStatus []*RowStatus `json:"rowStatus,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Errors") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *UploadStatus) MarshalJSON() ([]byte, error) {
	type noMethod UploadStatus
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// method id "doubleclickbidmanager.lineitems.downloadlineitems":

type LineitemsDownloadlineitemsCall struct {
	s                        *Service
	downloadlineitemsrequest *DownloadLineItemsRequest
	opt_                     map[string]interface{}
	ctx_                     context.Context
}

// Downloadlineitems: Retrieves line items in CSV format.
func (r *LineitemsService) Downloadlineitems(downloadlineitemsrequest *DownloadLineItemsRequest) *LineitemsDownloadlineitemsCall {
	c := &LineitemsDownloadlineitemsCall{s: r.s, opt_: make(map[string]interface{})}
	c.downloadlineitemsrequest = downloadlineitemsrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *LineitemsDownloadlineitemsCall) Fields(s ...googleapi.Field) *LineitemsDownloadlineitemsCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// Context sets the context to be used in this call's Do method.
// Any pending HTTP request will be aborted if the provided context
// is canceled.
func (c *LineitemsDownloadlineitemsCall) Context(ctx context.Context) *LineitemsDownloadlineitemsCall {
	c.ctx_ = ctx
	return c
}

func (c *LineitemsDownloadlineitemsCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.downloadlineitemsrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "lineitems/downloadlineitems")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *LineitemsDownloadlineitemsCall) Do() (*DownloadLineItemsResponse, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *DownloadLineItemsResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves line items in CSV format.",
	//   "httpMethod": "POST",
	//   "id": "doubleclickbidmanager.lineitems.downloadlineitems",
	//   "path": "lineitems/downloadlineitems",
	//   "request": {
	//     "$ref": "DownloadLineItemsRequest"
	//   },
	//   "response": {
	//     "$ref": "DownloadLineItemsResponse"
	//   }
	// }

}

// method id "doubleclickbidmanager.lineitems.uploadlineitems":

type LineitemsUploadlineitemsCall struct {
	s                      *Service
	uploadlineitemsrequest *UploadLineItemsRequest
	opt_                   map[string]interface{}
	ctx_                   context.Context
}

// Uploadlineitems: Uploads line items in CSV format.
func (r *LineitemsService) Uploadlineitems(uploadlineitemsrequest *UploadLineItemsRequest) *LineitemsUploadlineitemsCall {
	c := &LineitemsUploadlineitemsCall{s: r.s, opt_: make(map[string]interface{})}
	c.uploadlineitemsrequest = uploadlineitemsrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *LineitemsUploadlineitemsCall) Fields(s ...googleapi.Field) *LineitemsUploadlineitemsCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// Context sets the context to be used in this call's Do method.
// Any pending HTTP request will be aborted if the provided context
// is canceled.
func (c *LineitemsUploadlineitemsCall) Context(ctx context.Context) *LineitemsUploadlineitemsCall {
	c.ctx_ = ctx
	return c
}

func (c *LineitemsUploadlineitemsCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.uploadlineitemsrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "lineitems/uploadlineitems")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *LineitemsUploadlineitemsCall) Do() (*UploadLineItemsResponse, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *UploadLineItemsResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Uploads line items in CSV format.",
	//   "httpMethod": "POST",
	//   "id": "doubleclickbidmanager.lineitems.uploadlineitems",
	//   "path": "lineitems/uploadlineitems",
	//   "request": {
	//     "$ref": "UploadLineItemsRequest"
	//   },
	//   "response": {
	//     "$ref": "UploadLineItemsResponse"
	//   }
	// }

}

// method id "doubleclickbidmanager.queries.createquery":

type QueriesCreatequeryCall struct {
	s     *Service
	query *Query
	opt_  map[string]interface{}
	ctx_  context.Context
}

// Createquery: Creates a query.
func (r *QueriesService) Createquery(query *Query) *QueriesCreatequeryCall {
	c := &QueriesCreatequeryCall{s: r.s, opt_: make(map[string]interface{})}
	c.query = query
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *QueriesCreatequeryCall) Fields(s ...googleapi.Field) *QueriesCreatequeryCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// Context sets the context to be used in this call's Do method.
// Any pending HTTP request will be aborted if the provided context
// is canceled.
func (c *QueriesCreatequeryCall) Context(ctx context.Context) *QueriesCreatequeryCall {
	c.ctx_ = ctx
	return c
}

func (c *QueriesCreatequeryCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.query)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "query")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *QueriesCreatequeryCall) Do() (*Query, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Query
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a query.",
	//   "httpMethod": "POST",
	//   "id": "doubleclickbidmanager.queries.createquery",
	//   "path": "query",
	//   "request": {
	//     "$ref": "Query"
	//   },
	//   "response": {
	//     "$ref": "Query"
	//   }
	// }

}

// method id "doubleclickbidmanager.queries.deletequery":

type QueriesDeletequeryCall struct {
	s       *Service
	queryId int64
	opt_    map[string]interface{}
	ctx_    context.Context
}

// Deletequery: Deletes a stored query as well as the associated stored
// reports.
func (r *QueriesService) Deletequery(queryId int64) *QueriesDeletequeryCall {
	c := &QueriesDeletequeryCall{s: r.s, opt_: make(map[string]interface{})}
	c.queryId = queryId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *QueriesDeletequeryCall) Fields(s ...googleapi.Field) *QueriesDeletequeryCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// Context sets the context to be used in this call's Do method.
// Any pending HTTP request will be aborted if the provided context
// is canceled.
func (c *QueriesDeletequeryCall) Context(ctx context.Context) *QueriesDeletequeryCall {
	c.ctx_ = ctx
	return c
}

func (c *QueriesDeletequeryCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "query/{queryId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"queryId": strconv.FormatInt(c.queryId, 10),
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *QueriesDeletequeryCall) Do() error {
	res, err := c.doRequest("json")
	if err != nil {
		return err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return err
	}
	return nil
	// {
	//   "description": "Deletes a stored query as well as the associated stored reports.",
	//   "httpMethod": "DELETE",
	//   "id": "doubleclickbidmanager.queries.deletequery",
	//   "parameterOrder": [
	//     "queryId"
	//   ],
	//   "parameters": {
	//     "queryId": {
	//       "description": "Query ID to delete.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "query/{queryId}"
	// }

}

// method id "doubleclickbidmanager.queries.getquery":

type QueriesGetqueryCall struct {
	s       *Service
	queryId int64
	opt_    map[string]interface{}
	ctx_    context.Context
}

// Getquery: Retrieves a stored query.
func (r *QueriesService) Getquery(queryId int64) *QueriesGetqueryCall {
	c := &QueriesGetqueryCall{s: r.s, opt_: make(map[string]interface{})}
	c.queryId = queryId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *QueriesGetqueryCall) Fields(s ...googleapi.Field) *QueriesGetqueryCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// Context sets the context to be used in this call's Do method.
// Any pending HTTP request will be aborted if the provided context
// is canceled.
func (c *QueriesGetqueryCall) Context(ctx context.Context) *QueriesGetqueryCall {
	c.ctx_ = ctx
	return c
}

func (c *QueriesGetqueryCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "query/{queryId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"queryId": strconv.FormatInt(c.queryId, 10),
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *QueriesGetqueryCall) Do() (*Query, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Query
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves a stored query.",
	//   "httpMethod": "GET",
	//   "id": "doubleclickbidmanager.queries.getquery",
	//   "parameterOrder": [
	//     "queryId"
	//   ],
	//   "parameters": {
	//     "queryId": {
	//       "description": "Query ID to retrieve.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "query/{queryId}",
	//   "response": {
	//     "$ref": "Query"
	//   }
	// }

}

// method id "doubleclickbidmanager.queries.listqueries":

type QueriesListqueriesCall struct {
	s    *Service
	opt_ map[string]interface{}
	ctx_ context.Context
}

// Listqueries: Retrieves stored queries.
func (r *QueriesService) Listqueries() *QueriesListqueriesCall {
	c := &QueriesListqueriesCall{s: r.s, opt_: make(map[string]interface{})}
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *QueriesListqueriesCall) Fields(s ...googleapi.Field) *QueriesListqueriesCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// Context sets the context to be used in this call's Do method.
// Any pending HTTP request will be aborted if the provided context
// is canceled.
func (c *QueriesListqueriesCall) Context(ctx context.Context) *QueriesListqueriesCall {
	c.ctx_ = ctx
	return c
}

func (c *QueriesListqueriesCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "queries")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *QueriesListqueriesCall) Do() (*ListQueriesResponse, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ListQueriesResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves stored queries.",
	//   "httpMethod": "GET",
	//   "id": "doubleclickbidmanager.queries.listqueries",
	//   "path": "queries",
	//   "response": {
	//     "$ref": "ListQueriesResponse"
	//   }
	// }

}

// method id "doubleclickbidmanager.queries.runquery":

type QueriesRunqueryCall struct {
	s               *Service
	queryId         int64
	runqueryrequest *RunQueryRequest
	opt_            map[string]interface{}
	ctx_            context.Context
}

// Runquery: Runs a stored query to generate a report.
func (r *QueriesService) Runquery(queryId int64, runqueryrequest *RunQueryRequest) *QueriesRunqueryCall {
	c := &QueriesRunqueryCall{s: r.s, opt_: make(map[string]interface{})}
	c.queryId = queryId
	c.runqueryrequest = runqueryrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *QueriesRunqueryCall) Fields(s ...googleapi.Field) *QueriesRunqueryCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// Context sets the context to be used in this call's Do method.
// Any pending HTTP request will be aborted if the provided context
// is canceled.
func (c *QueriesRunqueryCall) Context(ctx context.Context) *QueriesRunqueryCall {
	c.ctx_ = ctx
	return c
}

func (c *QueriesRunqueryCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.runqueryrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "query/{queryId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"queryId": strconv.FormatInt(c.queryId, 10),
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *QueriesRunqueryCall) Do() error {
	res, err := c.doRequest("json")
	if err != nil {
		return err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return err
	}
	return nil
	// {
	//   "description": "Runs a stored query to generate a report.",
	//   "httpMethod": "POST",
	//   "id": "doubleclickbidmanager.queries.runquery",
	//   "parameterOrder": [
	//     "queryId"
	//   ],
	//   "parameters": {
	//     "queryId": {
	//       "description": "Query ID to run.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "query/{queryId}",
	//   "request": {
	//     "$ref": "RunQueryRequest"
	//   }
	// }

}

// method id "doubleclickbidmanager.reports.listreports":

type ReportsListreportsCall struct {
	s       *Service
	queryId int64
	opt_    map[string]interface{}
	ctx_    context.Context
}

// Listreports: Retrieves stored reports.
func (r *ReportsService) Listreports(queryId int64) *ReportsListreportsCall {
	c := &ReportsListreportsCall{s: r.s, opt_: make(map[string]interface{})}
	c.queryId = queryId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ReportsListreportsCall) Fields(s ...googleapi.Field) *ReportsListreportsCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// Context sets the context to be used in this call's Do method.
// Any pending HTTP request will be aborted if the provided context
// is canceled.
func (c *ReportsListreportsCall) Context(ctx context.Context) *ReportsListreportsCall {
	c.ctx_ = ctx
	return c
}

func (c *ReportsListreportsCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "queries/{queryId}/reports")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"queryId": strconv.FormatInt(c.queryId, 10),
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *ReportsListreportsCall) Do() (*ListReportsResponse, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ListReportsResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves stored reports.",
	//   "httpMethod": "GET",
	//   "id": "doubleclickbidmanager.reports.listreports",
	//   "parameterOrder": [
	//     "queryId"
	//   ],
	//   "parameters": {
	//     "queryId": {
	//       "description": "Query ID with which the reports are associated.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "queries/{queryId}/reports",
	//   "response": {
	//     "$ref": "ListReportsResponse"
	//   }
	// }

}
