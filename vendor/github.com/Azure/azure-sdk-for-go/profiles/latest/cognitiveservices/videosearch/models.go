// +build go1.9

// Copyright 2019 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package videosearch

import original "github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v1.0/videosearch"

const (
	DefaultEndpoint = original.DefaultEndpoint
)

type ErrorCode = original.ErrorCode

const (
	InsufficientAuthorization ErrorCode = original.InsufficientAuthorization
	InvalidAuthorization      ErrorCode = original.InvalidAuthorization
	InvalidRequest            ErrorCode = original.InvalidRequest
	None                      ErrorCode = original.None
	RateLimitExceeded         ErrorCode = original.RateLimitExceeded
	ServerError               ErrorCode = original.ServerError
)

type ErrorSubCode = original.ErrorSubCode

const (
	AuthorizationDisabled   ErrorSubCode = original.AuthorizationDisabled
	AuthorizationExpired    ErrorSubCode = original.AuthorizationExpired
	AuthorizationMissing    ErrorSubCode = original.AuthorizationMissing
	AuthorizationRedundancy ErrorSubCode = original.AuthorizationRedundancy
	Blocked                 ErrorSubCode = original.Blocked
	HTTPNotAllowed          ErrorSubCode = original.HTTPNotAllowed
	NotImplemented          ErrorSubCode = original.NotImplemented
	ParameterInvalidValue   ErrorSubCode = original.ParameterInvalidValue
	ParameterMissing        ErrorSubCode = original.ParameterMissing
	ResourceError           ErrorSubCode = original.ResourceError
	UnexpectedError         ErrorSubCode = original.UnexpectedError
)

type Freshness = original.Freshness

const (
	Day   Freshness = original.Day
	Month Freshness = original.Month
	Week  Freshness = original.Week
)

type SafeSearch = original.SafeSearch

const (
	Moderate SafeSearch = original.Moderate
	Off      SafeSearch = original.Off
	Strict   SafeSearch = original.Strict
)

type TextFormat = original.TextFormat

const (
	HTML TextFormat = original.HTML
	Raw  TextFormat = original.Raw
)

type Type = original.Type

const (
	TypeAnswer              Type = original.TypeAnswer
	TypeCreativeWork        Type = original.TypeCreativeWork
	TypeErrorResponse       Type = original.TypeErrorResponse
	TypeIdentifiable        Type = original.TypeIdentifiable
	TypeImageObject         Type = original.TypeImageObject
	TypeMediaObject         Type = original.TypeMediaObject
	TypeResponse            Type = original.TypeResponse
	TypeResponseBase        Type = original.TypeResponseBase
	TypeSearchResultsAnswer Type = original.TypeSearchResultsAnswer
	TypeThing               Type = original.TypeThing
	TypeTrendingVideos      Type = original.TypeTrendingVideos
	TypeVideoDetails        Type = original.TypeVideoDetails
	TypeVideoObject         Type = original.TypeVideoObject
	TypeVideos              Type = original.TypeVideos
)

type VideoInsightModule = original.VideoInsightModule

const (
	All           VideoInsightModule = original.All
	RelatedVideos VideoInsightModule = original.RelatedVideos
	VideoResult   VideoInsightModule = original.VideoResult
)

type VideoLength = original.VideoLength

const (
	VideoLengthAll    VideoLength = original.VideoLengthAll
	VideoLengthLong   VideoLength = original.VideoLengthLong
	VideoLengthMedium VideoLength = original.VideoLengthMedium
	VideoLengthShort  VideoLength = original.VideoLengthShort
)

type VideoPricing = original.VideoPricing

const (
	VideoPricingAll  VideoPricing = original.VideoPricingAll
	VideoPricingFree VideoPricing = original.VideoPricingFree
	VideoPricingPaid VideoPricing = original.VideoPricingPaid
)

type VideoQueryScenario = original.VideoQueryScenario

const (
	List                VideoQueryScenario = original.List
	SingleDominantVideo VideoQueryScenario = original.SingleDominantVideo
)

type VideoResolution = original.VideoResolution

const (
	VideoResolutionAll     VideoResolution = original.VideoResolutionAll
	VideoResolutionHD1080p VideoResolution = original.VideoResolutionHD1080p
	VideoResolutionHD720p  VideoResolution = original.VideoResolutionHD720p
	VideoResolutionSD480p  VideoResolution = original.VideoResolutionSD480p
)

type Answer = original.Answer
type BaseClient = original.BaseClient
type BasicAnswer = original.BasicAnswer
type BasicCreativeWork = original.BasicCreativeWork
type BasicIdentifiable = original.BasicIdentifiable
type BasicMediaObject = original.BasicMediaObject
type BasicResponse = original.BasicResponse
type BasicResponseBase = original.BasicResponseBase
type BasicSearchResultsAnswer = original.BasicSearchResultsAnswer
type BasicThing = original.BasicThing
type CreativeWork = original.CreativeWork
type Error = original.Error
type ErrorResponse = original.ErrorResponse
type Identifiable = original.Identifiable
type ImageObject = original.ImageObject
type MediaObject = original.MediaObject
type PivotSuggestions = original.PivotSuggestions
type Query = original.Query
type QueryContext = original.QueryContext
type Response = original.Response
type ResponseBase = original.ResponseBase
type SearchResultsAnswer = original.SearchResultsAnswer
type Thing = original.Thing
type TrendingVideos = original.TrendingVideos
type TrendingVideosCategory = original.TrendingVideosCategory
type TrendingVideosSubcategory = original.TrendingVideosSubcategory
type TrendingVideosTile = original.TrendingVideosTile
type VideoDetails = original.VideoDetails
type VideoObject = original.VideoObject
type Videos = original.Videos
type VideosClient = original.VideosClient
type VideosModule = original.VideosModule

func New() BaseClient {
	return original.New()
}
func NewVideosClient() VideosClient {
	return original.NewVideosClient()
}
func NewWithoutDefaults(endpoint string) BaseClient {
	return original.NewWithoutDefaults(endpoint)
}
func PossibleErrorCodeValues() []ErrorCode {
	return original.PossibleErrorCodeValues()
}
func PossibleErrorSubCodeValues() []ErrorSubCode {
	return original.PossibleErrorSubCodeValues()
}
func PossibleFreshnessValues() []Freshness {
	return original.PossibleFreshnessValues()
}
func PossibleSafeSearchValues() []SafeSearch {
	return original.PossibleSafeSearchValues()
}
func PossibleTextFormatValues() []TextFormat {
	return original.PossibleTextFormatValues()
}
func PossibleTypeValues() []Type {
	return original.PossibleTypeValues()
}
func PossibleVideoInsightModuleValues() []VideoInsightModule {
	return original.PossibleVideoInsightModuleValues()
}
func PossibleVideoLengthValues() []VideoLength {
	return original.PossibleVideoLengthValues()
}
func PossibleVideoPricingValues() []VideoPricing {
	return original.PossibleVideoPricingValues()
}
func PossibleVideoQueryScenarioValues() []VideoQueryScenario {
	return original.PossibleVideoQueryScenarioValues()
}
func PossibleVideoResolutionValues() []VideoResolution {
	return original.PossibleVideoResolutionValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
