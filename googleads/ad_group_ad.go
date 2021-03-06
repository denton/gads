package v201809

import (
	"encoding/xml"
)

type AdGroupAdService struct {
	Auth
}

type AppUrl struct {
	Url    string `xml:"url"`
	OsType string `xml:"osType"` // "OS_TYPE_IOS", "OS_TYPE_ANDROID", "UNKNOWN"
}

type AdStrengthInfo struct {
	AdStrength  string   `xml:"adStrength"`
	ActionItems []string `xml:"actionItems"`
}

type CallOnlyAd struct {
	AdGroupId                  int64                   `xml:"-"`
	Id                         int64                   `xml:"id,omitempty"`
	DisplayUrl                 string                  `xml:"displayUrl"`
	FinalUrlSuffix             string                  `xml:"finalUrlSuffix"`
	DevicePreference           int64                   `xml:"devicePreference,omitempty"`
	Automated                  bool                    `xml:"automated,omitempty"`
	SystemManagedEntitySource  string                  `xml:"systemManagedEntitySource,omitempty"`
	Type                       string                  `xml:"type,omitempty"`
	CountryCode                string                  `xml:"countryCode"`
	PhoneNumber                string                  `xml:"phoneNumber"`
	BusinessName               string                  `xml:"businessName"`
	Description1               string                  `xml:"description1"`
	Description2               string                  `xml:"description2"`
	CallTracked                bool                    `xml:"callTracked"`
	DisableCallConversion      bool                    `xml:"disableCallConversion"`
	ConversionTypeId           int64                   `xml:"conversionTypeId"`
	PhoneNumberVerificationUrl string                  `xml:"phoneNumberVerificationUrl"`
	Status                     string                  `xml:"-"`
	Labels                     []Label                 `xml:"-"`
	PolicySummary              *AdGroupAdPolicySummary `xml:"-"`
	AdStrengthInfo             *AdStrengthInfo         `xml:"-"`
	BaseCampaignId             int64                   `xml:"-"`
	BaseAdGroupId              int64                   `xml:"-"`
}

type TextAd struct {
	AdGroupId                 int64                   `xml:"-"`
	Id                        int64                   `xml:"id,omitempty"`
	Url                       string                  `xml:"url"`
	DisplayUrl                string                  `xml:"displayUrl"`
	FinalUrls                 []string                `xml:"finalUrls,omitempty"`
	FinalUrlSuffix            string                  `xml:"finalUrlSuffix"`
	FinalMobileUrls           []string                `xml:"finalMobileUrls,omitempty"`
	FinalAppUrls              []AppUrl                `xml:"finalAppUrls,omitempty"`
	TrackingUrlTemplate       string                  `xml:"trackingUrlTemplate,omitempty"`
	UrlCustomParameters       *CustomParameters       `xml:"urlCustomParameters,omitempty"`
	DevicePreference          int64                   `xml:"devicePreference,omitempty"`
	Automated                 bool                    `xml:"automated,omitempty"`
	SystemManagedEntitySource string                  `xml:"systemManagedEntitySource,omitempty"`
	Type                      string                  `xml:"type,omitempty"`
	Headline                  string                  `xml:"headline"`
	Description1              string                  `xml:"description1"`
	Description2              string                  `xml:"description2"`
	Status                    string                  `xml:"-"`
	Labels                    []Label                 `xml:"-"`
	PolicySummary             *AdGroupAdPolicySummary `xml:"-"`
	AdStrengthInfo            *AdStrengthInfo         `xml:"-"`
	BaseCampaignId            int64                   `xml:"-"`
	BaseAdGroupId             int64                   `xml:"-"`
}

type PolicyTopicEvidence struct {
	PolicyTopicEvidenceType                        string   `xml:"description2"`
	EvidenceTextList                               []string `xml:"evidenceTextList"`
	PolicyTopicEvidenceDestinationMismatchUrlTypes []string `xml:"policyTopicEvidenceDestinationMismatchUrlTypes"`
}

type PolicyTopicEntry struct {
	PolicyTopicEntryType     string              `xml:"policyTopicEntryType"`
	PolicyTopicEvidences     PolicyTopicEvidence `xml:"policyTopicEvidences"`
	PolicyTopicConstraints   interface{}         `xml:"policyTopicConstraints"` //todo
	PolicyTopicId            string              `xml:"policyTopicId"`
	PolicyTopicName          string              `xml:"policyTopicName"`
	PolicyTopicHelpCenterUrl string              `xml:"policyTopicHelpCenterUrl"`
}

type AssetPolicySummaryInfo struct {
	PolicyTopicEntries     []PolicyTopicEntry `xml:"policyTopicEntries"`
	ReviewState            string             `xml:"reviewState"`
	DenormalizedStatus     string             `xml:"denormalizedStatus"`
	CombinedApprovalStatus string             `xml:"combinedApprovalStatus"`
	PolicySummaryInfo      string             `xml:"PolicySummaryInfo"`
}

type Asset struct {
	AssetId      int64  `xml:"assetId"`
	AssetName    string `xml:"assetName"`
	AssetSubtype string `xml:"assetSubtype"`
	AssetStatus  string `xml:"assetStatus"`
	Type         string `xml:"type"`
}

type TextAsset struct {
	Asset
	AssetText string `xml:"assetText"`
}

type AssetLink struct {
	Asset                  interface{}            `xml:"asset"` //todo
	PinnedField            string                 `xml:"pinnedField"`
	AssetPolicySummaryInfo AssetPolicySummaryInfo `xml:"assetPolicySummaryInfo"`
	AssetPerformanceLabel  string                 `xml:"assetPerformanceLabel"`
}

type ResponsiveSearchAd struct {
	AdGroupId                 int64                   `xml:"-"`
	Id                        int64                   `xml:"id,omitempty"`
	FinalUrls                 []string                `xml:"finalUrls,omitempty"`
	FinalUrlSuffix            string                  `xml:"finalUrlSuffix"`
	FinalMobileUrls           []string                `xml:"finalMobileUrls,omitempty"`
	TrackingUrlTemplate       string                  `xml:"trackingUrlTemplate,omitempty"`
	UrlCustomParameters       *CustomParameters       `xml:"urlCustomParameters,omitempty"`
	Automated                 bool                    `xml:"automated,omitempty"`
	SystemManagedEntitySource string                  `xml:"systemManagedEntitySource,omitempty"`
	Type                      string                  `xml:"type,omitempty"`
	Headlines                 []AssetLink             `xml:"headlines"`
	Descriptions              []AssetLink             `xml:"descriptions"`
	Path1                     string                  `xml:"path1"`
	Path2                     string                  `xml:"path2"`
	Status                    string                  `xml:"-"`
	Labels                    []Label                 `xml:"-"`
	PolicySummary             *AdGroupAdPolicySummary `xml:"-"`
	AdStrengthInfo            *AdStrengthInfo         `xml:"-"`
	BaseCampaignId            int64                   `xml:"-"`
	BaseAdGroupId             int64                   `xml:"-"`
}

type ExpandedTextAd struct {
	AdGroupId                 int64                   `xml:"-"`
	Id                        int64                   `xml:"id,omitempty"`
	FinalUrls                 []string                `xml:"finalUrls,omitempty"`
	FinalUrlSuffix            string                  `xml:"finalUrlSuffix"`
	FinalMobileUrls           []string                `xml:"finalMobileUrls,omitempty"`
	TrackingUrlTemplate       string                  `xml:"trackingUrlTemplate,omitempty"`
	UrlCustomParameters       *CustomParameters       `xml:"urlCustomParameters,omitempty"`
	Automated                 bool                    `xml:"automated,omitempty"`
	SystemManagedEntitySource string                  `xml:"systemManagedEntitySource,omitempty"`
	Type                      string                  `xml:"type,omitempty"`
	HeadlinePart1             string                  `xml:"headlinePart1,omitempty"`
	HeadlinePart2             string                  `xml:"headlinePart2,omitempty"`
	HeadlinePart3             string                  `xml:"headlinePart3,omitempty"`
	Description               string                  `xml:"description,omitempty"`
	Description2              string                  `xml:"description2,omitempty"`
	Path1                     string                  `xml:"path1,omitempty"`
	Path2                     string                  `xml:"path2,omitempty"`
	Status                    string                  `xml:"-"`
	Labels                    []Label                 `xml:"-"`
	BaseCampaignId            int64                   `xml:"-"`
	BaseAdGroupId             int64                   `xml:"-"`
	PolicySummary             *AdGroupAdPolicySummary `xml:"-"`
	AdStrengthInfo            *AdStrengthInfo         `xml:"-"`
}

type BatchExpandedTextAd struct {
	AdGroupId           int64                   `xml:"-"`
	Id                  int64                   `xml:"id,omitempty"`
	FinalUrls           []string                `xml:"finalUrls,omitempty"`
	FinalUrlSuffix      string                  `xml:"finalUrlSuffix"`
	FinalMobileUrls     []string                `xml:"finalMobileUrls,omitempty"`
	TrackingUrlTemplate string                  `xml:"trackingUrlTemplate,omitempty"`
	UrlCustomParameters *CustomParameters       `xml:"urlCustomParameters,omitempty"`
	Type                string                  `xml:"type,omitempty"`
	HeadlinePart1       string                  `xml:"headlinePart1,omitempty"`
	HeadlinePart2       string                  `xml:"headlinePart2,omitempty"`
	HeadlinePart3       string                  `xml:"headlinePart3,omitempty"`
	Description         string                  `xml:"description,omitempty"`
	Description2        string                  `xml:"description2,omitempty"`
	Path1               string                  `xml:"path1,omitempty"`
	Path2               string                  `xml:"path2,omitempty"`
	Status              string                  `xml:"-"`
	Labels              []Label                 `xml:"-"`
	PolicySummary       *AdGroupAdPolicySummary `xml:"-"`
	AdStrengthInfo      *AdStrengthInfo         `xml:"-"`
	BaseCampaignId      int64                   `xml:"-"`
	BaseAdGroupId       int64                   `xml:"-"`
}

type BatchExpandedTextAdInner struct {
	HeadlinePart1       string   `xml:"headlinePart1,omitempty"`
	HeadlinePart2       string   `xml:"headlinePart2,omitempty"`
	Description         string   `xml:"description,omitempty"`
	Path1               string   `xml:"path1,omitempty"`
	Path2               string   `xml:"path2,omitempty"`
	FinalUrls           []string `xml:"finalUrls,omitempty"`
	FinalMobileUrls     []string `xml:"FinalMobileUrls,omitempty"`
	TrackingUrlTemplate string   `xml:"trackingUrlTemplate,omitempty"`
}

func (ad BatchExpandedTextAd) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = []xml.Attr{{Value: "operand"}}
	e.EncodeToken(start)

	e.EncodeElement(ad.Id, xml.StartElement{
		Name: xml.Name{Space: "", Local: "id"},
		Attr: []xml.Attr{},
	})
	e.EncodeElement(
		ad.AdGroupId,
		xml.StartElement{Name: xml.Name{Space: "", Local: "adGroupId"}},
	)
	e.EncodeElement(BatchExpandedTextAdInner{
		HeadlinePart1:       ad.HeadlinePart1,
		HeadlinePart2:       ad.HeadlinePart2,
		Description:         ad.Description,
		Path1:               ad.Path1,
		Path2:               ad.Path2,
		FinalUrls:           ad.FinalUrls,
		FinalMobileUrls:     []string{},
		TrackingUrlTemplate: ad.TrackingUrlTemplate,
	}, xml.StartElement{
		Name: xml.Name{Space: "", Local: "ad"},
		Attr: []xml.Attr{
			{
				Name:  xml.Name{Space: "http://www.w3.org/2001/XMLSchema-instance", Local: "type"},
				Value: "ExpandedTextAd",
			},
		},
	})
	e.EncodeElement(ad.Status, xml.StartElement{Name: xml.Name{Space: "", Local: "status"}})
	e.EncodeElement(ad.Labels, xml.StartElement{Name: xml.Name{Space: "", Local: "labels"}})

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

type DynamicSettings struct {
	LandscapeLogoImage Media  `xml:"landscapeLogoImage"`
	PricePrefix        string `xml:"pricePrefix"`
	PromoText          string `xml:"promoText"`
}

type ResponsiveDisplayAd struct {
	AdGroupId                 int64                   `xml:"-"`
	Id                        int64                   `xml:"id,omitempty"`
	FinalUrls                 []string                `xml:"finalUrls,omitempty"`
	FinalUrlSuffix            string                  `xml:"finalUrlSuffix"`
	FinalMobileUrls           []string                `xml:"finalMobileUrls,omitempty"`
	TrackingUrlTemplate       string                  `xml:"trackingUrlTemplate,omitempty"`
	UrlCustomParameters       *CustomParameters       `xml:"urlCustomParameters,omitempty"`
	Automated                 bool                    `xml:"automated,omitempty"`
	SystemManagedEntitySource string                  `xml:"systemManagedEntitySource,omitempty"`
	Type                      string                  `xml:"type,omitempty"`
	MarketingImage            Media                   `xml:"marketingImage"`
	LogoImage                 Media                   `xml:"logoImage"`
	SquareMarketingImage      Media                   `xml:"squareMarketingImage"`
	ShortHeadline             string                  `xml:"shortHeadline"`
	LongHeadline              string                  `xml:"longHeadline"`
	Description               string                  `xml:"description"`
	BusinessName              string                  `xml:"businessName"`
	MainColor                 string                  `xml:"mainColor"`
	AccentColor               string                  `xml:"accentColor"`
	AllowFlexibleColor        bool                    `xml:"allowFlexibleColor"`
	CallToActionText          string                  `xml:"callToActionText"`
	DynamicDisplayAdSettings  DynamicSettings         `xml:"dynamicDisplayAdSettings"`
	FormatSetting             string                  `xml:"formatSetting"`
	Status                    string                  `xml:"-"`
	Labels                    []Label                 `xml:"-"`
	BaseCampaignId            int64                   `xml:"-"`
	BaseAdGroupId             int64                   `xml:"-"`
	PolicySummary             *AdGroupAdPolicySummary `xml:"-"`
	AdStrengthInfo            *AdStrengthInfo         `xml:"-"`
}

type AdGroupExperimentData struct {
	ExperimentId          int64  `xml:"experimentId"`
	ExperimentDeltaStatus string `xml:"experimentDeltaStatus"`
	ExperimentDataStatus  string `xml:"experimentDataStatus"`
}

type ImageAd struct {
	AdGroupId                 int64                   `xml:"-"`
	Id                        int64                   `xml:"id,omitempty"`
	Url                       string                  `xml:"url"`
	DisplayUrl                string                  `xml:"displayUrl"`
	FinalUrls                 []string                `xml:"finalUrls,omitempty"`
	FinalUrlSuffix            string                  `xml:"finalUrlSuffix"`
	FinalMobileUrls           []string                `xml:"finalMobileUrls,omitempty"`
	FinalAppUrls              []AppUrl                `xml:"finalAppUrls,omitempty"`
	TrackingUrlTemplate       string                  `xml:"trackingUrlTemplate,omitempty"`
	UrlCustomParameters       *CustomParameters       `xml:"urlCustomParameters,omitempty"`
	DevicePreference          int64                   `xml:"devicePreference,omitempty"`
	Automated                 bool                    `xml:"automated,omitempty"`
	SystemManagedEntitySource string                  `xml:"systemManagedEntitySource,omitempty"`
	Type                      string                  `xml:"type,omitempty"`
	Image                     int64                   `xml:"imageId"`
	Name                      string                  `xml:"name"`
	AdToCopyImageFrom         int64                   `xml:"adToCopyImageFrom"`
	Status                    string                  `xml:"-"`
	Labels                    []Label                 `xml:"-"`
	PolicySummary             *AdGroupAdPolicySummary `xml:"-"`
	AdStrengthInfo            *AdStrengthInfo         `xml:"-"`
	BaseCampaignId            int64                   `xml:"-"`
	BaseAdGroupId             int64                   `xml:"-"`
}

/*
type MobileAd struct {
	AdGroupId           int64                   `xml:"-"`
	Id                  int64                   `xml:"id,omitempty"`
	Url                 string                  `xml:"url"`
	DisplayUrl          string                  `xml:"displayUrl"`
	FinalUrls           []string                `xml:"finalUrls,omitempty"`
	FinalMobileUrls     []string                `xml:"finalMobileUrls,omitempty"`
	FinalAppUrls        []AppUrl                `xml:"finalAppUrls,omitempty"`
	TrackingUrlTemplate string                  `xml:"trackingUrlTemplate,omitempty"`
	UrlCustomParameters *CustomParameters       `xml:"urlCustomParameters,omitempty"`
	Type                string                  `xml:"type,omitempty"`
	DevicePreference    int64                   `xml:"devicePreference,omitempty"`
	Headline            string                  `xml:"headline"`
	Description         string                  `xml:"description"`
	MarkupLanguages     []string                `xml:"markupLanguages"`
	MobileCarriers      []string                `xml:"mobileCarriers"`
	BusinessName        string                  `xml:"businessName"`
	CountryCode         string                  `xml:"countryCode"`
	PhoneNumber         string                  `xml:"phoneNumber"`
	Status              string                  `xml:"-"`
	Labels              []Label                 `xml:"-"`
	PolicySummary       *AdGroupAdPolicySummary `xml:"-"`
}
*/

type TemplateElementField struct {
	Name       string `xml:"name"`
	Type       string `xml:"type"`
	FieldText  string `xml:"fieldText"`
	FieldMedia string `xml:"fieldMedia"`
}

type TemplateElement struct {
	UniqueName string                 `xml:"uniqueName"`
	Fields     []TemplateElementField `xml:"fields"`
}

type TemplateAd struct {
	Id                        int64                   `xml:"id,omitempty"`
	Url                       string                  `xml:"url"`
	DisplayUrl                string                  `xml:"displayUrl"`
	FinalUrls                 []string                `xml:"finalUrls,omitempty"`
	FinalUrlSuffix            string                  `xml:"finalUrlSuffix"`
	FinalMobileUrls           []string                `xml:"finalMobileUrls,omitempty"`
	FinalAppUrls              []AppUrl                `xml:"finalAppUrls,omitempty"`
	TrackingUrlTemplate       string                  `xml:"trackingUrlTemplate,omitempty"`
	UrlCustomParameters       *CustomParameters       `xml:"urlCustomParameters,omitempty"`
	Type                      string                  `xml:"type,omitempty"`
	DevicePreference          int64                   `xml:"devicePreference,omitempty"`
	Automated                 bool                    `xml:"automated,omitempty"`
	SystemManagedEntitySource string                  `xml:"systemManagedEntitySource,omitempty"`
	TemplateId                int64                   `xml:"templateId"`
	AdUnionId                 int64                   `xml:"adUnionId>id"`
	TemplateElements          []TemplateElement       `xml:"templateElements"`
	Dimensions                []Dimensions            `xml:"dimensions"`
	AdAsImage                 *Media                  `xml:"adAsImage"`
	Name                      string                  `xml:"name"`
	Duration                  int64                   `xml:"duration"`
	OriginAdId                *int64                  `xml:"originAdId"`
	AdGroupId                 int64                   `xml:"-"`
	Status                    string                  `xml:"-"`
	Labels                    []Label                 `xml:"-"`
	PolicySummary             *AdGroupAdPolicySummary `xml:"-"`
	AdStrengthInfo            *AdStrengthInfo         `xml:"-"`
	BaseCampaignId            int64                   `xml:"-"`
	BaseAdGroupId             int64                   `xml:"-"`
}

type DynamicSearchAd struct {
	Id                        int64                   `xml:"id,omitempty"`
	Url                       string                  `xml:"url"`
	DisplayUrl                string                  `xml:"displayUrl"`
	FinalUrls                 []string                `xml:"finalUrls,omitempty"`
	FinalUrlSuffix            string                  `xml:"finalUrlSuffix"`
	FinalMobileUrls           []string                `xml:"finalMobileUrls,omitempty"`
	FinalAppUrls              []AppUrl                `xml:"finalAppUrls,omitempty"`
	TrackingUrlTemplate       string                  `xml:"trackingUrlTemplate,omitempty"`
	UrlCustomParameters       *CustomParameters       `xml:"urlCustomParameters,omitempty"`
	Automated                 bool                    `xml:"automated,omitempty"`
	Type                      string                  `xml:"type,omitempty"`
	DevicePreference          int64                   `xml:"devicePreference,omitempty"`
	SystemManagedEntitySource string                  `xml:"systemManagedEntitySource,omitempty"`
	Description1              string                  `xml:"description1"`
	Description2              string                  `xml:"description2"`
	AdGroupId                 int64                   `xml:"-"`
	Status                    string                  `xml:"-"`
	Labels                    []Label                 `xml:"-"`
	PolicySummary             *AdGroupAdPolicySummary `xml:"-"`
	AdStrengthInfo            *AdStrengthInfo         `xml:"-"`
	BaseCampaignId            int64                   `xml:"-"`
	BaseAdGroupId             int64                   `xml:"-"`
}

type ProductAd struct {
	AdGroupId                 int64                   `xml:"-"`
	Id                        int64                   `xml:"id,omitempty"`
	Type                      string                  `xml:"type,omitempty"`
	FinalUrlSuffix            string                  `xml:"finalUrlSuffix,omitempty"`
	UrlData                   []*UrlData              `xml:"urlData,omitempty"`
	Automated                 bool                    `xml:"automated,omitempty"`
	DevicePreference          int64                   `xml:"devicePreference,omitempty"`
	SystemManagedEntitySource string                  `xml:"systemManagedEntitySource,omitempty"`
	Status                    string                  `xml:"-"`
	Labels                    []Label                 `xml:"-"`
	PolicySummary             *AdGroupAdPolicySummary `xml:"-"`
	AdStrengthInfo            *AdStrengthInfo         `xml:"-"`
	BaseCampaignId            int64                   `xml:"-"`
	BaseAdGroupId             int64                   `xml:"-"`
}

type MultiAssetResponsiveDisplayAd struct {
	AdGroupId                 int64                   `xml:"-"`
	Id                        int64                   `xml:"id,omitempty"`
	Type                      string                  `xml:"type,omitempty"`
	FinalUrls                 []string                `xml:"finalUrls,omitempty"`
	FinalMobileUrls           []string                `xml:"finalMobileUrls,omitempty"`
	TrackingUrlTemplate       string                  `xml:"trackingUrlTemplate,omitempty"`
	FinalUrlSuffix            string                  `xml:"finalUrlSuffix,omitempty"`
	UrlCustomParameters       *CustomParameters       `xml:"urlCustomParameters,omitempty"`
	UrlData                   []*UrlData              `xml:"urlData,omitempty"`
	Automated                 bool                    `xml:"automated,omitempty"`
	SystemManagedEntitySource string                  `xml:"systemManagedEntitySource,omitempty"`
	Status                    string                  `xml:"-"`
	Labels                    []Label                 `xml:"-"`
	PolicySummary             *AdGroupAdPolicySummary `xml:"-"`
	AdStrengthInfo            *AdStrengthInfo         `xml:"-"`
	BaseCampaignId            int64                   `xml:"-"`
	BaseAdGroupId             int64                   `xml:"-"`
}

type ExpandedDynamicSearchAd struct {
	AdGroupId                 int64                   `xml:"-"`
	Id                        int64                   `xml:"id,omitempty"`
	Type                      string                  `xml:"type,omitempty"`
	Url                       string                  `xml:"url"`
	DisplayUrl                string                  `xml:"displayUrl"`
	FinalUrls                 []string                `xml:"finalUrls,omitempty"`
	FinalMobileUrls           []string                `xml:"finalMobileUrls,omitempty"`
	FinalAppUrls              []AppUrl                `xml:"finalAppUrls,omitempty"`
	TrackingUrlTemplate       string                  `xml:"trackingUrlTemplate,omitempty"`
	FinalUrlSuffix            string                  `xml:"finalUrlSuffix,omitempty"`
	UrlCustomParameters       *CustomParameters       `xml:"urlCustomParameters,omitempty"`
	UrlData                   []*UrlData              `xml:"urlData,omitempty"`
	Automated                 bool                    `xml:"automated,omitempty"`
	DevicePreference          int64                   `xml:"devicePreference,omitempty"`
	SystemManagedEntitySource string                  `xml:"systemManagedEntitySource,omitempty"`
	Description1              string                  `xml:"description1"`
	Description2              string                  `xml:"description2"`
	Status                    string                  `xml:"-"`
	Labels                    []Label                 `xml:"-"`
	PolicySummary             *AdGroupAdPolicySummary `xml:"-"`
	AdStrengthInfo            *AdStrengthInfo         `xml:"-"`
	BaseCampaignId            int64                   `xml:"-"`
	BaseAdGroupId             int64                   `xml:"-"`
}

type UrlList struct {
	Urls []string `xml:"urls"`
}

type UrlData struct {
	UrlId               string  `xml:"urlId"`
	FinalUrls           UrlList `xml:"finalUrls,omitempty"`
	FinalMobileUrls     UrlList `xml:"finalMobileUrls,omitempty"`
	TrackingUrlTemplate string  `xml:"trackingUrlTemplate,omitempty"`
}

type GoalOptimizedShoppingAd struct {
	Id                        int64                   `xml:"id,omitempty"`
	Type                      string                  `xml:"type,omitempty"`
	TrackingUrlTemplate       string                  `xml:"trackingUrlTemplate,omitempty"`
	FinalUrlSuffix            string                  `xml:"finalUrlSuffix,omitempty"`
	UrlCustomParameters       *CustomParameters       `xml:"urlCustomParameters,omitempty"`
	UrlData                   []*UrlData              `xml:"urlData,omitempty"`
	Automated                 bool                    `xml:"automated,omitempty"`
	DevicePreference          int64                   `xml:"devicePreference,omitempty"`
	SystemManagedEntitySource string                  `xml:"systemManagedEntitySource,omitempty"`
	AdGroupId                 int64                   `xml:"-"`
	Status                    string                  `xml:"-"`
	Labels                    []Label                 `xml:"-"`
	PolicySummary             *AdGroupAdPolicySummary `xml:"-"`
	AdStrengthInfo            *AdStrengthInfo         `xml:"-"`
	BaseCampaignId            int64                   `xml:"-"`
	BaseAdGroupId             int64                   `xml:"-"`
}

type UniversalAppAd struct {
	AdGroupId                 int64                   `xml:"-"`
	Id                        int64                   `xml:"id,omitempty"`
	Url                       string                  `xml:"url"`
	FinalUrls                 []string                `xml:"finalUrls,omitempty"`
	FinalUrlSuffix            string                  `xml:"finalUrlSuffix"`
	FinalMobileUrls           []string                `xml:"finalMobileUrls,omitempty"`
	TrackingUrlTemplate       string                  `xml:"trackingUrlTemplate,omitempty"`
	UrlData                   []*UrlData              `xml:"urlData,omitempty"`
	UrlCustomParameters       *CustomParameters       `xml:"urlCustomParameters,omitempty"`
	Automated                 bool                    `xml:"automated,omitempty"`
	SystemManagedEntitySource string                  `xml:"systemManagedEntitySource,omitempty"`
	Type                      string                  `xml:"type,omitempty"`
	Headlines                 []AssetLink             `xml:"headlines"`
	Descriptions              []AssetLink             `xml:"descriptions"`
	Images                    *[]AssetLink            `xml:"images"`
	Videos                    *[]AssetLink            `xml:"videos"`
	Html5MediaBundles         *[]AssetLink            `xml:"html5MediaBundles"`
	Status                    string                  `xml:"-"`
	Labels                    []Label                 `xml:"-"`
	PolicySummary             *AdGroupAdPolicySummary `xml:"-"`
	AdStrengthInfo            *AdStrengthInfo         `xml:"-"`
	BaseCampaignId            int64                   `xml:"-"`
	BaseAdGroupId             int64                   `xml:"-"`
}

type ShowcaseAd struct {
	AdGroupId                 int64                   `xml:"-"`
	Id                        int64                   `xml:"id,omitempty"`
	Url                       string                  `xml:"url"`
	DisplayUrl                string                  `xml:"displayUrl"`
	FinalUrls                 []string                `xml:"finalUrls,omitempty"`
	FinalUrlSuffix            string                  `xml:"finalUrlSuffix"`
	FinalMobileUrls           []string                `xml:"finalMobileUrls,omitempty"`
	FinalAppUrls              []AppUrl                `xml:"finalAppUrls,omitempty"`
	TrackingUrlTemplate       string                  `xml:"trackingUrlTemplate,omitempty"`
	UrlData                   []*UrlData              `xml:"urlData,omitempty"`
	UrlCustomParameters       *CustomParameters       `xml:"urlCustomParameters,omitempty"`
	DevicePreference          int64                   `xml:"devicePreference,omitempty"`
	Automated                 bool                    `xml:"automated,omitempty"`
	SystemManagedEntitySource string                  `xml:"systemManagedEntitySource,omitempty"`
	Type                      string                  `xml:"type,omitempty"`
	Name                      string                  `xml:"name,omitempty"`
	Headline                  string                  `xml:"headline"`
	Description               string                  `xml:"description"`
	CollapsedImage            *Media                  `xml:"collapsedImage"`
	ExpandedImage             *Media                  `xml:"expandedImage"`
	Status                    string                  `xml:"-"`
	Labels                    []Label                 `xml:"-"`
	PolicySummary             *AdGroupAdPolicySummary `xml:"-"`
	AdStrengthInfo            *AdStrengthInfo         `xml:"-"`
	BaseCampaignId            int64                   `xml:"-"`
	BaseAdGroupId             int64                   `xml:"-"`
}

type RichMediaAd struct {
	AdGroupId                 int64                   `xml:"-"`
	Id                        int64                   `xml:"id,omitempty"`
	Url                       string                  `xml:"url"`
	DisplayUrl                string                  `xml:"displayUrl"`
	FinalUrls                 []string                `xml:"finalUrls,omitempty"`
	FinalUrlSuffix            string                  `xml:"finalUrlSuffix"`
	FinalMobileUrls           []string                `xml:"finalMobileUrls,omitempty"`
	FinalAppUrls              []AppUrl                `xml:"finalAppUrls,omitempty"`
	TrackingUrlTemplate       string                  `xml:"trackingUrlTemplate,omitempty"`
	UrlData                   []*UrlData              `xml:"urlData,omitempty"`
	UrlCustomParameters       *CustomParameters       `xml:"urlCustomParameters,omitempty"`
	DevicePreference          int64                   `xml:"devicePreference,omitempty"`
	Automated                 bool                    `xml:"automated,omitempty"`
	SystemManagedEntitySource string                  `xml:"systemManagedEntitySource,omitempty"`
	Type                      string                  `xml:"type,omitempty"`
	Name                      string                  `xml:"name"`
	Dimensions                *Dimensions             `xml:"dimensions,omitempty"`
	Snippet                   string                  `xml:"snippet"`
	ImpressionBeaconUrl       string                  `xml:"impressionBeaconUrl"`
	AdDuration                int                     `xml:"adDuration"`
	CertifiedVendorFormatId   int64                   `xml:"certifiedVendorFormatId"`
	SourceUrl                 string                  `xml:"sourceUrl"`
	RichMediaAdType           string                  `xml:"richMediaAdType"`
	AdAttributes              []string                `xml:"adAttributes"`
	Status                    string                  `xml:"-"`
	Labels                    []Label                 `xml:"-"`
	PolicySummary             *AdGroupAdPolicySummary `xml:"-"`
	AdStrengthInfo            *AdStrengthInfo         `xml:"-"`
	BaseCampaignId            int64                   `xml:"-"`
	BaseAdGroupId             int64                   `xml:"-"`
}

type ThirdPartyRedirectAd struct {
	AdGroupId                 int64                   `xml:"-"`
	Id                        int64                   `xml:"id,omitempty"`
	Url                       string                  `xml:"url"`
	DisplayUrl                string                  `xml:"displayUrl"`
	FinalUrls                 []string                `xml:"finalUrls,omitempty"`
	FinalUrlSuffix            string                  `xml:"finalUrlSuffix"`
	FinalMobileUrls           []string                `xml:"finalMobileUrls,omitempty"`
	FinalAppUrls              []AppUrl                `xml:"finalAppUrls,omitempty"`
	TrackingUrlTemplate       string                  `xml:"trackingUrlTemplate,omitempty"`
	UrlData                   []*UrlData              `xml:"urlData,omitempty"`
	UrlCustomParameters       *CustomParameters       `xml:"urlCustomParameters,omitempty"`
	DevicePreference          int64                   `xml:"devicePreference,omitempty"`
	Automated                 bool                    `xml:"automated,omitempty"`
	SystemManagedEntitySource string                  `xml:"systemManagedEntitySource,omitempty"`
	Type                      string                  `xml:"type,omitempty"`
	Name                      string                  `xml:"name"`
	Dimensions                *Dimensions             `xml:"dimensions,omitempty"`
	Snippet                   string                  `xml:"snippet"`
	ImpressionBeaconUrl       string                  `xml:"impressionBeaconUrl"`
	AdDuration                int                     `xml:"adDuration"`
	CertifiedVendorFormatId   int64                   `xml:"certifiedVendorFormatId"`
	SourceUrl                 string                  `xml:"sourceUrl"`
	RichMediaAdType           string                  `xml:"richMediaAdType"`
	AdAttributes              []string                `xml:"adAttributes"`
	IsCookieTargeted          bool                    `xml:"isCookieTargeted"`
	IsUserInterestTargeted    bool                    `xml:"isUserInterestTargeted"`
	IsTagged                  bool                    `xml:"isTagged"`
	VideoTypes                []string                `xml:"videoTypes"`
	ExpandingDirections       []string                `xml:"expandingDirections"`
	Status                    string                  `xml:"-"`
	Labels                    []Label                 `xml:"-"`
	PolicySummary             *AdGroupAdPolicySummary `xml:"-"`
	AdStrengthInfo            *AdStrengthInfo         `xml:"-"`
	BaseCampaignId            int64                   `xml:"-"`
	BaseAdGroupId             int64                   `xml:"-"`
}

type Teaser struct {
	Headline     string `xml:"headline"`
	Description  string `xml:"description"`
	BusinessName string `xml:"businessName"`
	LogoImage    *Media `xml:"logoImage"`
}

type DisplayCallToAction struct {
	Text      string `xml:"text"`
	TextColor string `xml:"textColor"`
	UrlId     string `xml:"urlId"`
}

type ProductImage struct {
	ProductImage        *Media               `xml:"productImage"`
	Description         string               `xml:"description"`
	DisplayCallToAction *DisplayCallToAction `xml:"displayCallToAction"`
}

type GmailAd struct {
	AdGroupId                         int64                   `xml:"-"`
	Id                                int64                   `xml:"id,omitempty"`
	FinalUrls                         []string                `xml:"finalUrls,omitempty"`
	FinalUrlSuffix                    string                  `xml:"finalUrlSuffix"`
	FinalMobileUrls                   []string                `xml:"finalMobileUrls,omitempty"`
	TrackingUrlTemplate               string                  `xml:"trackingUrlTemplate,omitempty"`
	UrlData                           []*UrlData              `xml:"urlData,omitempty"`
	UrlCustomParameters               *CustomParameters       `xml:"urlCustomParameters,omitempty"`
	Automated                         bool                    `xml:"automated,omitempty"`
	SystemManagedEntitySource         string                  `xml:"systemManagedEntitySource,omitempty"`
	Type                              string                  `xml:"type,omitempty"`
	Teaser                            Teaser                  `xml:"teaser"`
	HeaderImage                       *Media                  `xml:"headerImage"`
	MarketingImage                    *Media                  `xml:"marketingImage"`
	MrketingImageHeadline             string                  `xml:"marketingImageHeadline"`
	MarketingImageDescription         string                  `xml:"marketingImageDescription"`
	MarketingImageDisplayCallToAction *DisplayCallToAction    `xml:"marketingImageDisplayCallToAction"`
	ProductImages                     *[]ProductImage         `xml:"productImages"`
	ProductVideoList                  *[]Media                `xml:"productVideoList"`
	Status                            string                  `xml:"-"`
	Labels                            []Label                 `xml:"-"`
	PolicySummary                     *AdGroupAdPolicySummary `xml:"-"`
	AdStrengthInfo                    *AdStrengthInfo         `xml:"-"`
	BaseCampaignId                    int64                   `xml:"-"`
	BaseAdGroupId                     int64                   `xml:"-"`
}

type DeprecatedAd struct {
	AdGroupId                 int64                   `xml:"-"`
	Id                        int64                   `xml:"id,omitempty"`
	Url                       string                  `xml:"url"`
	DisplayUrl                string                  `xml:"displayUrl"`
	FinalUrls                 []string                `xml:"finalUrls,omitempty"`
	FinalUrlSuffix            string                  `xml:"finalUrlSuffix"`
	FinalMobileUrls           []string                `xml:"finalMobileUrls,omitempty"`
	FinalAppUrls              []AppUrl                `xml:"finalAppUrls,omitempty"`
	TrackingUrlTemplate       string                  `xml:"trackingUrlTemplate,omitempty"`
	UrlData                   []*UrlData              `xml:"urlData,omitempty"`
	UrlCustomParameters       *CustomParameters       `xml:"urlCustomParameters,omitempty"`
	DevicePreference          int64                   `xml:"devicePreference,omitempty"`
	Automated                 bool                    `xml:"automated,omitempty"`
	SystemManagedEntitySource string                  `xml:"systemManagedEntitySource,omitempty"`
	Type                      string                  `xml:"type,omitempty"`
	Name                      string                  `xml:"name,omitempty"`
	DeprecatedAdType          string                  `xml:"deprecatedAdType"`
	Status                    string                  `xml:"-"`
	Labels                    []Label                 `xml:"-"`
	PolicySummary             *AdGroupAdPolicySummary `xml:"-"`
	AdStrengthInfo            *AdStrengthInfo         `xml:"-"`
	BaseCampaignId            int64                   `xml:"-"`
	BaseAdGroupId             int64                   `xml:"-"`
}

//

type AdGroupAdOperations map[string]AdGroupAds

func NewAdGroupAdService(auth *Auth) *AdGroupAdService {
	return &AdGroupAdService{Auth: *auth}
}

func NewTextAd(adGroupId int64, url, displayUrl, headline, description1, description2, status string) TextAd {
	return TextAd{
		AdGroupId:    adGroupId,
		Url:          url,
		DisplayUrl:   displayUrl,
		Headline:     headline,
		Description1: description1,
		Description2: description2,
		Status:       status,
	}
}

type AdGroupAdLabel struct {
	AdGroupAdId int64 `xml:"adGroupAdId"`
	LabelId     int64 `xml:"labelId"`
}

type AdGroupAdLabelOperations map[string][]AdGroupAdLabel

type AdUrlUpgrade struct {
	AdId                int64  `xml:"adId"`
	FinalUrl            string `xml:"finalUrl"`
	FinalMobileUrl      string `xml:"finalMobileUrl"`
	TrackingUrlTemplate string `xml:"trackingUrlTemplate"`
}

type AdGroupAdPolicySummary struct {
	ReviewState            string `xml:"reviewState"`
	DenormalizedStatus     string `xml:"denormalizedStatus"`
	CombinedApprovalStatus string `xml:"combinedApprovalStatus"`
}

// Get returns an array of ad's and the total number of ad's matching
// the selector.
//
// Example
//
//   ads, totalCount, err := adGroupAdService.Get(
//     gads.Selector{
//       Fields: []string{
//         "AdGroupId",
//         "Status",
//         "AdGroupCreativeApprovalStatus",
//         "AdGroupAdDisapprovalReasons",
//         "AdGroupAdTrademarkDisapproved",
//       },
//       Predicates: []gads.Predicate{
//         {"AdGroupId", "EQUALS", []string{adGroupId}},
//       },
//     },
//   )
//
// Selectable fields are
//   "AdGroupId", "Id", "Url", "DisplayUrl", "CreativeFinalUrls", "CreativeFinalMobileUrls",
//   "CreativeFinalAppUrls", "CreativeTrackingUrlTemplate", "CreativeUrlCustomParameters",
//   "DevicePreference", "Status", "AdGroupCreativeApprovalStatus", "AdGroupAdDisapprovalReasons"
//   "AdGroupAdTrademarkDisapproved", "Labels"
//
//   TextAd
//     "Headline", "Description1", "Description2"
//
//   ImageAd
//     "ImageCreativeName"
//
// filterable fields are
//   "AdGroupId", "Id", "Url", "DisplayUrl", "CreativeFinalUrls", "CreativeFinalMobileUrls",
//   "CreativeFinalAppUrls", "CreativeTrackingUrlTemplate", "CreativeUrlCustomParameters",
//   "DevicePreference", "Status", "AdGroupCreativeApprovalStatus", "AdGroupAdDisapprovalReasons"
//   "Labels"
//
//   TextAd specific fields
//     "Headline", "Description1", "Description2"
//
//   ImageAd specific fields
//     "ImageCreativeName"
//
// Relevant documentation
//
//     https://developers.google.com/adwords/api/docs/reference/v201409/AdGroupAdService#get
//
func (s AdGroupAdService) Get(selector Selector) (adGroupAds AdGroupAds, totalCount int64, err error) {
	selector.XMLName = xml.Name{baseUrl, "serviceSelector"}
	respBody, err := s.Auth.request(
		adGroupAdServiceUrl,
		"get",
		struct {
			XMLName xml.Name
			Sel     Selector
		}{
			XMLName: xml.Name{
				Space: baseUrl,
				Local: "get",
			},
			Sel: selector,
		},
	)
	if err != nil {
		return adGroupAds, totalCount, err
	}
	getResp := struct {
		Size       int64      `xml:"rval>totalNumEntries"`
		AdGroupAds AdGroupAds `xml:"rval>entries"`
	}{}
	err = xml.Unmarshal([]byte(respBody), &getResp)
	if err != nil {
		return adGroupAds, totalCount, err
	}
	return getResp.AdGroupAds, getResp.Size, err
}

// Mutate allows you to add, modify and remove ads, returning the
// modified ads.
//
// Example
//
//  ads, err := adGroupAdService.Mutate(
//    gads.AdGroupAdOperations{
//      "ADD": {
//        gads.NewTextAd(
//          adGroup.Id,
//          "https://classdo.com/en",
//          "classdo.com",
//          "test headline",
//          "test line one",
//          "test line two",
//          "PAUSED",
//        ),
//      },
//      "SET": {
//        modifiedAd,
//      },
//      "REMOVE": {
//        adNeedingRemoval,
//      },
//    },
//  )
//
// Relevant documentation
//
//     https://developers.google.com/adwords/api/docs/reference/v201409/AdGroupAdService#mutate
//
func (s *AdGroupAdService) Mutate(adGroupAdOperations AdGroupAdOperations) (adGroupAds AdGroupAds, err error) {
	type adGroupAdOperation struct {
		Action    string     `xml:"operator"`
		AdGroupAd AdGroupAds `xml:"operand"`
	}
	operations := []adGroupAdOperation{}
	for action, adGroupAds := range adGroupAdOperations {
		for _, adGroupAd := range adGroupAds {
			ad := []interface{}{adGroupAd}
			operations = append(operations,
				adGroupAdOperation{
					Action:    action,
					AdGroupAd: ad,
				},
			)
		}
	}
	mutation := struct {
		XMLName xml.Name
		Ops     []adGroupAdOperation `xml:"operations"`
	}{
		XMLName: xml.Name{
			Space: baseUrl,
			Local: "mutate",
		},
		Ops: operations,
	}

	respBody, err := s.Auth.request(adGroupAdServiceUrl, "mutate", mutation)
	if err != nil {
		return adGroupAds, err
	}
	mutateResp := struct {
		AdGroupAds AdGroupAds `xml:"rval>value"`
	}{}
	err = xml.Unmarshal([]byte(respBody), &mutateResp)
	if err != nil {
		return adGroupAds, err
	}
	return mutateResp.AdGroupAds, err
}

// MutateLabel allows you to add and removes labels from ads.
//
// Example
//
//  ads, err := adGroupAdService.MutateLabel(
//    gads.AdGroupAdLabelOperations{
//      "ADD": {
//        gads.AdGroupAdLabel{AdGroupAdId: 3200, LabelId: 5353},
//        gads.AdGroupAdLabel{AdGroupAdId: 4320, LabelId: 5643},
//      },
//      "REMOVE": {
//        gads.AdGroupAdLabel{AdGroupAdId: 3653, LabelId: 5653},
//      },
//    }
//
// Relevant documentation
//
//     https://developers.google.com/adwords/api/docs/reference/v201409/AdGroupAdService#mutateLabel
//
func (s *AdGroupAdService) MutateLabel(adGroupAdLabelOperations AdGroupAdLabelOperations) (adGroupAdLabels []AdGroupAdLabel, err error) {
	type adGroupAdLabelOperation struct {
		Action         string         `xml:"operator"`
		AdGroupAdLabel AdGroupAdLabel `xml:"operand"`
	}
	operations := []adGroupAdLabelOperation{}
	for action, adGroupAdLabels := range adGroupAdLabelOperations {
		for _, adGroupAdLabel := range adGroupAdLabels {
			operations = append(operations,
				adGroupAdLabelOperation{
					Action:         action,
					AdGroupAdLabel: adGroupAdLabel,
				},
			)
		}
	}
	mutation := struct {
		XMLName xml.Name
		Ops     []adGroupAdLabelOperation `xml:"operations"`
	}{
		XMLName: xml.Name{
			Space: baseUrl,
			Local: "mutateLabel",
		},
		Ops: operations}
	respBody, err := s.Auth.request(adGroupAdServiceUrl, "mutateLabel", mutation)
	if err != nil {
		return adGroupAdLabels, err
	}
	mutateResp := struct {
		AdGroupAdLabels []AdGroupAdLabel `xml:"rval>value"`
	}{}
	err = xml.Unmarshal([]byte(respBody), &mutateResp)
	if err != nil {
		return adGroupAdLabels, err
	}

	return mutateResp.AdGroupAdLabels, err
}

// Query is not yet implemented
//
// Relevant documentation
//
//     https://developers.google.com/adwords/api/docs/reference/v201409/AdGroupAdService#query
//
func (s *AdGroupAdService) Query(query string) (adGroupAds AdGroupAds, totalCount int64, err error) {

	respBody, err := s.Auth.request(
		adGroupAdServiceUrl,
		"query",
		AWQLQuery{
			XMLName: xml.Name{
				Space: baseUrl,
				Local: "query",
			},
			Query: query,
		},
	)

	if err != nil {
		return adGroupAds, totalCount, err
	}

	getResp := struct {
		Size       int64      `xml:"rval>totalNumEntries"`
		AdGroupAds AdGroupAds `xml:"rval>entries"`
	}{}

	err = xml.Unmarshal([]byte(respBody), &getResp)
	if err != nil {
		return adGroupAds, totalCount, err
	}
	return getResp.AdGroupAds, getResp.Size, err

	//return adGroupAds, totalCount, ERROR_NOT_YET_IMPLEMENTED
}

// Query is not yet implemented
//
// Relevant documentation
//
//     https://developers.google.com/adwords/api/docs/reference/v201409/AdGroupAdService#upgradeUrl
//
func (s *AdGroupAdService) UpgradeUrl(adUrlUpgrades []AdUrlUpgrade) (adGroupAds AdGroupAds, err error) {
	return adGroupAds, ERROR_NOT_YET_IMPLEMENTED
}
