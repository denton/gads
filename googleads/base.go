package v201809

import (
	"bytes"
	"compress/gzip"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"
)

const (
	version               = "v201809"
	rootUrl               = "https://adwords.google.com/api/adwords/cm/"
	baseUrl               = "https://adwords.google.com/api/adwords/cm/" + version
	rootBillingUrl        = "https://adwords.google.com/api/adwords/billing/"
	baseBillingUrl        = "https://adwords.google.com/api/adwords/billing/" + version
	rootMcmUrl            = "https://adwords.google.com/api/adwords/mcm/"
	baseMcmUrl            = "https://adwords.google.com/api/adwords/mcm/" + version
	rootRemarketingUrl    = "https://adwords.google.com/api/adwords/rm/"
	baseRemarketingUrl    = "https://adwords.google.com/api/adwords/rm/" + version
	rootReportDownloadUrl = "https://adwords.google.com/api/adwords/reportdownload/"
	baseReportDownloadUrl = "https://adwords.google.com/api/adwords/reportdownload/" + version
	rootTrafficUrl        = "https://adwords.google.com/api/adwords/o/"
	baseTrafficUrl        = "https://adwords.google.com/api/adwords/o/" + version
	baseSyncUrl           = "https://adwords.google.com/api/adwords/ch/" + version
)

type ServiceUrl struct {
	Url  string
	Name string
}

// exceptions
var (
	ERROR_NOT_YET_IMPLEMENTED = fmt.Errorf("Not yet implemented")
)

var (

	// service urls
	adGroupAdServiceUrl          = ServiceUrl{baseUrl, "AdGroupAdService"}
	adGroupBidModifierServiceUrl = ServiceUrl{
		baseUrl,
		"AdGroupBidModifierService",
	}
	adGroupCriterionServiceUrl = ServiceUrl{
		baseUrl,
		"AdGroupCriterionService",
	}
	adGroupExtensionSettingServiceUrl = ServiceUrl{
		baseUrl,
		"AdGroupExtensionSettingService",
	}
	adGroupFeedServiceUrl = ServiceUrl{
		baseUrl,
		"AdGroupFeedService",
	}
	adGroupServiceUrl         = ServiceUrl{baseUrl, "AdGroupService"}
	adParamServiceUrl         = ServiceUrl{baseUrl, "AdParamService"}
	adwordsUserListServiceUrl = ServiceUrl{
		baseRemarketingUrl,
		"AdwordsUserListService",
	}
	batchJobServiceUrl        = ServiceUrl{baseUrl, "BatchJobService"}
	biddingStrategyServiceUrl = ServiceUrl{
		baseUrl,
		"BiddingStrategyService",
	}
	budgetOrderServiceUrl  = ServiceUrl{baseBillingUrl, "BudgetOrderService"}
	budgetServiceUrl       = ServiceUrl{baseUrl, "BudgetService"}
	campaignBidModifierUrl = ServiceUrl{
		baseUrl,
		"CampaignBidModifierService",
	}
	campaignExtensionSettingUrl = ServiceUrl{
		baseUrl,
		"CampaignExtensionSettingService",
	}
	campaignCriterionServiceUrl = ServiceUrl{
		baseUrl,
		"CampaignCriterionService",
	}
	campaignFeedServiceUrl = ServiceUrl{
		baseUrl,
		"CampaignFeedService",
	}
	campaignServiceUrl          = ServiceUrl{baseUrl, "CampaignService"}
	campaignSharedSetServiceUrl = ServiceUrl{
		baseUrl,
		"CampaignSharedSetService",
	}
	constantDataServiceUrl = ServiceUrl{
		baseUrl,
		"ConstantDataService",
	}
	conversionTrackerServiceUrl = ServiceUrl{
		baseUrl,
		"ConversionTrackerService",
	}
	customerFeedServiceUrl = ServiceUrl{
		baseUrl,
		"CustomerFeedService",
	}
	customerServiceUrl = ServiceUrl{
		baseMcmUrl,
		"CustomerService",
	}
	customerSyncServiceUrl = ServiceUrl{
		baseSyncUrl,
		"CustomerSyncService",
	}
	dataServiceUrl        = ServiceUrl{baseUrl, "DataService"}
	experimentServiceUrl  = ServiceUrl{baseUrl, "ExperimentService"}
	feedItemServiceUrl    = ServiceUrl{baseUrl, "FeedItemService"}
	feedMappingServiceUrl = ServiceUrl{
		baseUrl,
		"FeedMappingService",
	}
	feedServiceUrl        = ServiceUrl{baseUrl, "FeedService"}
	geoLocationServiceUrl = ServiceUrl{
		baseUrl,
		"GeoLocationService",
	}
	labelServiceUrl             = ServiceUrl{baseUrl, "LabelService"}
	locationCriterionServiceUrl = ServiceUrl{
		baseUrl,
		"LocationCriterionService",
	}
	managedCustomerServiceUrl = ServiceUrl{
		baseMcmUrl,
		"ManagedCustomerService",
	}
	mediaServiceUrl                 = ServiceUrl{baseUrl, "MediaService"}
	mutateJobServiceUrl             = ServiceUrl{baseUrl, "MutateJobService"}
	offlineConversionFeedServiceUrl = ServiceUrl{
		baseUrl,
		"OfflineConversionFeedService",
	}
	reportDefinitionServiceUrl = ServiceUrl{
		baseUrl,
		"ReportDefinitionService",
	}
	reportDownloadServiceUrl  = ServiceUrl{baseReportDownloadUrl, ""}
	sharedCriterionServiceUrl = ServiceUrl{
		baseUrl,
		"SharedCriterionService",
	}
	sharedSetServiceUrl     = ServiceUrl{baseUrl, "SharedSetService"}
	targetingIdeaServiceUrl = ServiceUrl{
		baseTrafficUrl,
		"TargetingIdeaService",
	}
	trafficEstimatorServiceUrl = ServiceUrl{
		baseTrafficUrl,
		"TrafficEstimatorService",
	}
)
var (
	knownErrors = []string{
		"InternalApiError.UNEXPECTED_INTERNAL_API_ERROR",
		"Service Unavailable",
	}
)

func (s ServiceUrl) String() string {
	if s.Name != "" {
		return s.Url + "/" + s.Name
	}
	return s.Url
}

type Auth struct {
	CustomerId     string
	DeveloperToken string
	UserAgent      string
	PartialFailure bool
	ValidateOnly   bool
	Testing        *testing.T `json:"-"`
	Client         HttpClient `json:"-"`
}

type HttpClient interface {
	Do(*http.Request) (*http.Response, error)
}

//
// Selector structs
//
type DateRange struct {
	Min string `xml:"min"`
	Max string `xml:"max"`
}

type Predicate struct {
	Field    string   `xml:"field"`
	Operator string   `xml:"operator"`
	Values   []string `xml:"values"`
}

type OrderBy struct {
	Field     string `xml:"field"`
	SortOrder string `xml:"sortOrder"`
}

type Paging struct {
	Offset int64 `xml:"https://adwords.google.com/api/adwords/cm/v201809 startIndex"`
	Limit  int64 `xml:"https://adwords.google.com/api/adwords/cm/v201809 numberResults"`
}

type Selector struct {
	XMLName    xml.Name
	Fields     []string    `xml:"fields,omitempty"`
	Predicates []Predicate `xml:"predicates"`
	DateRange  *DateRange  `xml:"dateRange,omitempty"`
	Ordering   []OrderBy   `xml:"ordering"`
	Paging     *Paging     `xml:"paging,omitempty"`
}

type AWQLQuery struct {
	XMLName xml.Name
	Query   string `xml:"query"`
}

// https://developers.google.com/adwords/api/docs/reference/v201809/AdGroupExtensionSettingService.DayOfWeek
// Days of the week.
// MONDAY, TUESDAY, WEDNESDAY, THURSDAY, FRIDAY, SATURDAY, SUNDAY
type DayOfWeek string

// https://developers.google.com/adwords/api/docs/reference/v201809/AdGroupExtensionSettingService.MinuteOfHour
// Minutes in an hour. Currently only 0, 15, 30, and 45 are supported
// ZERO, FIFTEEN, THIRTY, FORTY_FIVE
type MinuteOfHour string

// https://developers.google.com/adwords/api/docs/reference/v201809/AdGroupExtensionSettingService.GeoRestriction
// A restriction used to determine if the request context's geo should be matched.
// UNKNOWN, LOCATION_OF_PRESENCE
type GeoRestriction string

// https://developers.google.com/adwords/api/docs/reference/v201809/AdGroupExtensionSettingService.PolicyData
// Approval and policy information attached to an entity.
type PolicyData struct {
	DisapprovalReasons []DisapprovalReason `xml:"https://adwords.google.com/api/adwords/cm/v201809 disapprovalReasons,omitempty"`
	PolicyDataType     string              `xml:"https://adwords.google.com/api/adwords/cm/v201809 PolicyData.Type,omitempty"`
}

// https://developers.google.com/adwords/api/docs/reference/v201809/AdGroupExtensionSettingService.DisapprovalReason
// Container for information about why an AdWords entity was disapproved.
type DisapprovalReason struct {
	ShortName string `xml:"https://adwords.google.com/api/adwords/cm/v201809 shortName,omitempty"`
}

// error parsers
func selectorError() (err error) {
	return err
}

func (a *Auth) do(
	serviceUrl ServiceUrl,
	action string,
	body, ret interface{},
) error {
	raw, err := a.doRequest(serviceUrl, action, body)
	if err != nil {
		return err
	}

	if err = xml.Unmarshal([]byte(raw), &ret); err != nil {
		return err
	}

	if level := os.Getenv("DEBUG"); level != "" {
		if resBody, err := xml.MarshalIndent(ret, "  ", "  "); err != nil {
			fmt.Println("warn: ", err)
		} else {
			fmt.Printf("response->\n%s\n", string(resBody))
		}
	}

	return nil
}

func (a *Auth) request(
	serviceUrl ServiceUrl,
	action string,
	body interface{},
) (respBody []byte, err error) {
	return a.doRequest(serviceUrl, action, body)
}

var (
	tokenForCache string
)

func SetCacheToken(t string) {
	tokenForCache = t
}

func (a *Auth) doRequest(
	serviceUrl ServiceUrl,
	action string,
	body interface{},
) (respBody []byte, err error) {
	var result []byte

	for i := 3; i >= 0; i-- {
		result, err := a.doRequestFunc(serviceUrl, action, body)
		if err != nil {
			if !isShouldRetry(err) {
				return result, err
			}

			if i > 0 {
				time.Sleep(time.Second * 5)
			}
		} else {
			return result, err
		}
	}
	return result, err
}

func isShouldRetry(err error) bool {
	errorText := err.Error()
	for _, v := range knownErrors {
		if strings.Contains(errorText, v) {
			return true
		}
	}
	return false
}

func (a *Auth) doRequestFunc(
	serviceUrl ServiceUrl,
	action string,
	body interface{},
) (respBody []byte, err error) {

	startTime := time.Now()

	type devToken struct {
		XMLName xml.Name
	}
	type soapReqHeader struct {
		XMLName          xml.Name
		UserAgent        string `xml:"userAgent"`
		DeveloperToken   string `xml:"developerToken"`
		ClientCustomerId string `xml:"clientCustomerId,omitempty"`
		PartialFailure   bool   `xml:"partialFailure,omitempty"`
		ValidateOnly     bool   `xml:"validateOnly,omitempty"`
	}

	type soapReqBody struct {
		Body interface{}
	}

	type soapReqEnvelope struct {
		XMLName xml.Name
		Header  soapReqHeader `xml:"Header>RequestHeader"`
		Body    soapReqBody   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
	}

	reqHead := soapReqHeader{
		XMLName:          xml.Name{serviceUrl.Url, "RequestHeader"},
		UserAgent:        a.UserAgent,
		DeveloperToken:   a.DeveloperToken,
		ClientCustomerId: a.CustomerId,
	}

	// https://developers.google.com/adwords/api/docs/guides/partial-failure
	if a.PartialFailure {
		reqHead.PartialFailure = true
	}
	if a.ValidateOnly {
		reqHead.ValidateOnly = true
	}

	reqBody, err := xml.MarshalIndent(
		soapReqEnvelope{
			XMLName: xml.Name{
				"http://schemas.xmlsoap.org/soap/envelope/",
				"Envelope",
			},
			Header: reqHead,
			Body:   soapReqBody{body},
		},
		"  ", "  ")
	if err != nil {
		return []byte{}, err
	}

	// load cache
	cacheResp, ok := []byte{}, false
	if cache_ENABLED {
		cacheResp, ok = cache.Get([]string{
			serviceUrl.String(),
			tokenForCache,
			action,
			string(reqBody),
		})
	}

	respStatusCode := 0

	if ok && cache_ENABLED {
		respBody = cacheResp
		respStatusCode = 200
	} else {
		req, err := http.NewRequest("POST", serviceUrl.String(), bytes.NewReader(reqBody))
		req.Header.Add("Accept", "text/xml")
		req.Header.Add("User-Agent", "gads (gzip)")
		req.Header.Add("Accept-Encoding", "gzip")
		req.Header.Add("Accept", "multipart/*")
		req.Header.Add("Content-Type", "text/xml;charset=UTF-8")
		contentLength := fmt.Sprintf("%d", len(reqBody))
		req.Header.Add("Content-length", contentLength)
		req.Header.Add("SOAPAction", action)
		//if a.Testing != nil {
		//	a.Testing.Logf("request ->\n%s\n%#v\n%s\n", req.URL.String(), req.Header, string(reqBody))
		//}

		// Added some logging/"poor man's" debugging to inspect outbound SOAP requests
		if level := os.Getenv("DEBUG"); level != "" {
			fmt.Printf("request ->\n%s\n%#v\n%s\n", req.URL.String(), req.Header, string(reqBody))
		}

		resp, err := a.Client.Do(req)
		if err != nil {
			return []byte{}, err
		}
		defer resp.Body.Close()

		var reader io.ReadCloser
		switch resp.Header.Get("Content-Encoding") {
		case "gzip":
			reader, err = gzip.NewReader(resp.Body)
			defer reader.Close()
		default:
			reader = resp.Body
		}

		respBody, err = ioutil.ReadAll(reader)
		if err != nil {
			return []byte{}, err
		}
		respStatusCode = resp.StatusCode
	}

	// save cache
	if !ok && cache_ENABLED {
		cache.Set(
			[]string{
				serviceUrl.String(),
				action,
				string(reqBody),
			}, respBody,
		)
	}

	defer stat.count(serviceUrl.Name, ok, cache_MEM, time.Since(startTime))

	// Added some logging/"poor man's" debugging to inspect outbound SOAP requests
	if level := os.Getenv("DEBUG"); level != "" {
		fmt.Printf("response ->\n%s\n", string(respBody))
	}

	if a.Testing != nil {
		a.Testing.Logf(
			"respBody ->\n%s\n%s\n",
			string(respBody),
			fmt.Sprintf("%d", respStatusCode),
		)
	}

	type soapRespHeader struct {
		RequestId    string `xml:"requestId"`
		ServiceName  string `xml:"serviceName"`
		MethodName   string `xml:"methodName"`
		Operations   int64  `xml:"operations"`
		ResponseTime int64  `xml:"responseTime"`
	}

	type soapRespBody struct {
		Response []byte `xml:",innerxml"`
	}

	soapResp := struct {
		XMLName xml.Name       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
		Header  soapRespHeader `xml:"Header>RequestHeader"`
		Body    soapRespBody   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
	}{}

	err = xml.Unmarshal([]byte(respBody), &soapResp)
	if err != nil {
		return respBody, err
	}
	if respStatusCode == 400 || respStatusCode == 401 || respStatusCode == 403 || respStatusCode == 405 ||
		respStatusCode == 500 {
		fault := Fault{}
		err = xml.Unmarshal(soapResp.Body.Response, &fault)
		if err != nil {
			return respBody, err
		}

		for i := range fault.Errors.ApiExceptionFaults {
			switch fault.Errors.ApiExceptionFaults[i].ErrorsType {
			case "AuthenticationError",
				"RateExceededError",
				"DatabaseError",
				"InternalApiError":
				return soapResp.Body.Response, &baseError{
					code:    fault.Errors.ApiExceptionFaults[i].Reason,
					origErr: &fault.Errors,
				}
			}
		}

		if fault.Errors.ApiExceptionFaults == nil {
			return soapResp.Body.Response, errors.New(fault.FaultString)
		}

		return soapResp.Body.Response, &fault.Errors
	}
	return soapResp.Body.Response, err
}
