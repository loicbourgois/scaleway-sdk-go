// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package test provides methods and message types of the test v1 API.
package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/scaleway/scaleway-sdk-go/internal/errors"
	"github.com/scaleway/scaleway-sdk-go/internal/marshaler"
	"github.com/scaleway/scaleway-sdk-go/internal/parameter"
	"github.com/scaleway/scaleway-sdk-go/namegenerator"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

// always import dependencies
var (
	_ fmt.Stringer
	_ json.Unmarshaler
	_ url.URL
	_ net.IP
	_ http.Header
	_ bytes.Reader
	_ time.Time

	_ scw.ScalewayRequest
	_ marshaler.Duration
	_ scw.File
	_ = parameter.AddToQuery
	_ = namegenerator.GetRandomName
)

// API no Auth Service for end-to-end testing.
type API struct {
	client *scw.Client
}

// NewAPI returns a API object from a Scaleway client.
func NewAPI(client *scw.Client) *API {
	return &API{
		client: client,
	}
}

type EyeColors string

const (
	// EyeColorsUnknown is [insert doc].
	EyeColorsUnknown = EyeColors("unknown")
	// EyeColorsAmber is [insert doc].
	EyeColorsAmber = EyeColors("amber")
	// EyeColorsBlue is [insert doc].
	EyeColorsBlue = EyeColors("blue")
	// EyeColorsBrown is [insert doc].
	EyeColorsBrown = EyeColors("brown")
	// EyeColorsGray is [insert doc].
	EyeColorsGray = EyeColors("gray")
	// EyeColorsGreen is [insert doc].
	EyeColorsGreen = EyeColors("green")
	// EyeColorsHazel is [insert doc].
	EyeColorsHazel = EyeColors("hazel")
	// EyeColorsRed is [insert doc].
	EyeColorsRed = EyeColors("red")
	// EyeColorsViolet is [insert doc].
	EyeColorsViolet = EyeColors("violet")
)

func (enum EyeColors) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum EyeColors) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *EyeColors) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = EyeColors(EyeColors(tmp).String())
	return nil
}

type HumanStatus string

const (
	// HumanStatusUnknown is [insert doc].
	HumanStatusUnknown = HumanStatus("unknown")
	// HumanStatusStopped is [insert doc].
	HumanStatusStopped = HumanStatus("stopped")
	// HumanStatusRunning is [insert doc].
	HumanStatusRunning = HumanStatus("running")
)

func (enum HumanStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum HumanStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *HumanStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = HumanStatus(HumanStatus(tmp).String())
	return nil
}

type ListHumansRequestOrderBy string

const (
	// ListHumansRequestOrderByCreatedAtAsc is [insert doc].
	ListHumansRequestOrderByCreatedAtAsc = ListHumansRequestOrderBy("created_at_asc")
	// ListHumansRequestOrderByCreatedAtDesc is [insert doc].
	ListHumansRequestOrderByCreatedAtDesc = ListHumansRequestOrderBy("created_at_desc")
	// ListHumansRequestOrderByUpdatedAtAsc is [insert doc].
	ListHumansRequestOrderByUpdatedAtAsc = ListHumansRequestOrderBy("updated_at_asc")
	// ListHumansRequestOrderByUpdatedAtDesc is [insert doc].
	ListHumansRequestOrderByUpdatedAtDesc = ListHumansRequestOrderBy("updated_at_desc")
	// ListHumansRequestOrderByHeightAsc is [insert doc].
	ListHumansRequestOrderByHeightAsc = ListHumansRequestOrderBy("height_asc")
	// ListHumansRequestOrderByHeightDesc is [insert doc].
	ListHumansRequestOrderByHeightDesc = ListHumansRequestOrderBy("height_desc")
)

func (enum ListHumansRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListHumansRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListHumansRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListHumansRequestOrderBy(ListHumansRequestOrderBy(tmp).String())
	return nil
}

type Human struct {
	ID string `json:"id"`

	OrganizationID string `json:"organization_id"`

	CreatedAt time.Time `json:"created_at"`

	UpdatedAt time.Time `json:"updated_at"`

	Height float64 `json:"height"`

	ShoeSize float32 `json:"shoe_size"`

	AltitudeInMeter int32 `json:"altitude_in_meter"`

	AltitudeInMillimeter int64 `json:"altitude_in_millimeter"`

	FingersCount uint32 `json:"fingers_count"`

	HairCount uint64 `json:"hair_count"`

	IsHappy bool `json:"is_happy"`
	// EyesColor
	//
	// Default value: unknown
	EyesColor EyeColors `json:"eyes_color"`
	// Status
	//
	// Default value: unknown
	Status HumanStatus `json:"status"`

	Region scw.Region `json:"region"`
}

type ListHumansResponse struct {
	TotalCount uint32 `json:"total_count"`

	Humans []*Human `json:"humans"`
}

type RegisterResponse struct {
	SecretKey string `json:"secret_key"`

	AccessKey string `json:"access_key"`
}

// Service API

type RegisterRequest struct {
	Region scw.Region `json:"-"`

	Username string `json:"username"`
}

func (s *API) Register(req *RegisterRequest, opts ...scw.RequestOption) (*RegisterResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/test/v1/regions/" + fmt.Sprint(req.Region) + "/register",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp RegisterResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListHumansRequest struct {
	Region scw.Region `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
	// OrderBy
	//
	// Default value: created_at_asc
	OrderBy ListHumansRequestOrderBy `json:"-"`

	OrganizationID *string `json:"-"`
}

// ListHumans list all your humans
func (s *API) ListHumans(req *ListHumansRequest, opts ...scw.RequestOption) (*ListHumansResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/test/v1/regions/" + fmt.Sprint(req.Region) + "/humans",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListHumansResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListHumansResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListHumansResponse) UnsafeAppend(res interface{}) (uint32, scw.SdkError) {
	results, ok := res.(*ListHumansResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Humans = append(r.Humans, results.Humans...)
	r.TotalCount += uint32(len(results.Humans))
	return uint32(len(results.Humans)), nil
}

type GetHumanRequest struct {
	Region scw.Region `json:"-"`

	HumanID string `json:"-"`
}

// GetHuman get human details
//
// Get the human details associated with the given id.
func (s *API) GetHuman(req *GetHumanRequest, opts ...scw.RequestOption) (*Human, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.HumanID) == "" {
		return nil, errors.New("field HumanID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/test/v1/regions/" + fmt.Sprint(req.Region) + "/humans/" + fmt.Sprint(req.HumanID) + "",
		Headers: http.Header{},
	}

	var resp Human

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type CreateHumanRequest struct {
	Region scw.Region `json:"-"`

	Height float64 `json:"height"`

	ShoeSize float32 `json:"shoe_size"`

	AltitudeInMeter int32 `json:"altitude_in_meter"`

	AltitudeInMillimeter int64 `json:"altitude_in_millimeter"`

	FingersCount uint32 `json:"fingers_count"`

	HairCount uint64 `json:"hair_count"`

	IsHappy bool `json:"is_happy"`
	// EyesColor
	//
	// Default value: unknown
	EyesColor EyeColors `json:"eyes_color"`

	OrganizationID string `json:"organization_id"`
}

// CreateHuman create a new human
func (s *API) CreateHuman(req *CreateHumanRequest, opts ...scw.RequestOption) (*Human, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/test/v1/regions/" + fmt.Sprint(req.Region) + "/humans",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Human

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpdateHumanRequest struct {
	Region scw.Region `json:"-"`

	HumanID string `json:"-"`

	Height *float64 `json:"height"`

	ShoeSize *float32 `json:"shoe_size"`

	AltitudeInMeter *int32 `json:"altitude_in_meter"`

	AltitudeInMillimeter *int64 `json:"altitude_in_millimeter"`

	FingersCount *uint32 `json:"fingers_count"`

	HairCount *uint64 `json:"hair_count"`

	IsHappy *bool `json:"is_happy"`
	// EyesColor
	//
	// Default value: unknown
	EyesColor EyeColors `json:"eyes_color"`
}

// UpdateHuman update an existing human
//
// Update the human associated with the given id.
func (s *API) UpdateHuman(req *UpdateHumanRequest, opts ...scw.RequestOption) (*Human, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.HumanID) == "" {
		return nil, errors.New("field HumanID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PATCH",
		Path:    "/test/v1/regions/" + fmt.Sprint(req.Region) + "/humans/" + fmt.Sprint(req.HumanID) + "",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Human

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteHumanRequest struct {
	Region scw.Region `json:"-"`

	HumanID string `json:"-"`
}

// DeleteHuman delete an existing human
//
// Delete the human associated with the given id.
func (s *API) DeleteHuman(req *DeleteHumanRequest, opts ...scw.RequestOption) (*Human, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.HumanID) == "" {
		return nil, errors.New("field HumanID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/test/v1/regions/" + fmt.Sprint(req.Region) + "/humans/" + fmt.Sprint(req.HumanID) + "",
		Headers: http.Header{},
	}

	var resp Human

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type RunHumanRequest struct {
	Region scw.Region `json:"-"`

	HumanID string `json:"-"`
}

func (s *API) RunHuman(req *RunHumanRequest, opts ...scw.RequestOption) (*Human, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.HumanID) == "" {
		return nil, errors.New("field HumanID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/test/v1/regions/" + fmt.Sprint(req.Region) + "/humans/" + fmt.Sprint(req.HumanID) + "/run",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Human

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
