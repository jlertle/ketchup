// Code generated by protoc-gen-go.
// source: api.proto
// DO NOT EDIT!

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	api.proto
	content.proto
	data.proto
	export/export.proto
	metadata.proto
	page.proto
	route.proto
	theme.proto
	user.proto
	packages.proto
	struct.proto

It has these top-level messages:
	Error
	FieldError
	ListOptions
	UpdatePageRoutesRequest
	GetRenderedPageRequest
	GetRenderedPageResponse
	ListPageRequest
	ListPageResponse
	ListRouteResponse
	ListThemeResponse
	UpdateDataRequest
	ListDataResponse
	TLSSettingsReponse
	EnableTLSRequest
	InstallThemeRequest
	ContentMultiple
	ContentText
	ContentString
	Data
	Export
	Metadata
	Timestamp
	Author
	Page
	Content
	Route
	Theme
	ThemeTemplate
	ThemePlaceholder
	ThemeAsset
	User
	PackageRelease
	Package
	Registry
	Struct
	Value
	ListValue
*/
package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import ketchup_models "github.com/ketchuphq/ketchup/proto/ketchup/models"
import structpb "github.com/ketchuphq/ketchup/proto/structpb"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ErrorCode int32

const (
	ErrorCode_INTERNAL_SERVER_ERROR ErrorCode = 1
	ErrorCode_NOT_FOUND             ErrorCode = 2
	ErrorCode_FORBIDDEN             ErrorCode = 3
)

var ErrorCode_name = map[int32]string{
	1: "INTERNAL_SERVER_ERROR",
	2: "NOT_FOUND",
	3: "FORBIDDEN",
}
var ErrorCode_value = map[string]int32{
	"INTERNAL_SERVER_ERROR": 1,
	"NOT_FOUND":             2,
	"FORBIDDEN":             3,
}

func (x ErrorCode) Enum() *ErrorCode {
	p := new(ErrorCode)
	*p = x
	return p
}
func (x ErrorCode) String() string {
	return proto.EnumName(ErrorCode_name, int32(x))
}
func (x *ErrorCode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ErrorCode_value, data, "ErrorCode")
	if err != nil {
		return err
	}
	*x = ErrorCode(value)
	return nil
}
func (ErrorCode) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type GetRenderedPageRequest_RenderedPageFilter int32

const (
	GetRenderedPageRequest_all       GetRenderedPageRequest_RenderedPageFilter = 1
	GetRenderedPageRequest_published GetRenderedPageRequest_RenderedPageFilter = 2
	GetRenderedPageRequest_draft     GetRenderedPageRequest_RenderedPageFilter = 3
)

var GetRenderedPageRequest_RenderedPageFilter_name = map[int32]string{
	1: "all",
	2: "published",
	3: "draft",
}
var GetRenderedPageRequest_RenderedPageFilter_value = map[string]int32{
	"all":       1,
	"published": 2,
	"draft":     3,
}

func (x GetRenderedPageRequest_RenderedPageFilter) Enum() *GetRenderedPageRequest_RenderedPageFilter {
	p := new(GetRenderedPageRequest_RenderedPageFilter)
	*p = x
	return p
}
func (x GetRenderedPageRequest_RenderedPageFilter) String() string {
	return proto.EnumName(GetRenderedPageRequest_RenderedPageFilter_name, int32(x))
}
func (x *GetRenderedPageRequest_RenderedPageFilter) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(GetRenderedPageRequest_RenderedPageFilter_value, data, "GetRenderedPageRequest_RenderedPageFilter")
	if err != nil {
		return err
	}
	*x = GetRenderedPageRequest_RenderedPageFilter(value)
	return nil
}
func (GetRenderedPageRequest_RenderedPageFilter) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{4, 0}
}

type ListPageRequest_ListPageFilter int32

const (
	ListPageRequest_all       ListPageRequest_ListPageFilter = 1
	ListPageRequest_published ListPageRequest_ListPageFilter = 2
	ListPageRequest_draft     ListPageRequest_ListPageFilter = 3
)

var ListPageRequest_ListPageFilter_name = map[int32]string{
	1: "all",
	2: "published",
	3: "draft",
}
var ListPageRequest_ListPageFilter_value = map[string]int32{
	"all":       1,
	"published": 2,
	"draft":     3,
}

func (x ListPageRequest_ListPageFilter) Enum() *ListPageRequest_ListPageFilter {
	p := new(ListPageRequest_ListPageFilter)
	*p = x
	return p
}
func (x ListPageRequest_ListPageFilter) String() string {
	return proto.EnumName(ListPageRequest_ListPageFilter_name, int32(x))
}
func (x *ListPageRequest_ListPageFilter) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ListPageRequest_ListPageFilter_value, data, "ListPageRequest_ListPageFilter")
	if err != nil {
		return err
	}
	*x = ListPageRequest_ListPageFilter(value)
	return nil
}
func (ListPageRequest_ListPageFilter) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{6, 0}
}

type Error struct {
	Code             *ErrorCode    `protobuf:"varint,1,opt,name=code,enum=ketchup.api.ErrorCode" json:"code,omitempty"`
	Title            *string       `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	Detail           *string       `protobuf:"bytes,3,opt,name=detail" json:"detail,omitempty"`
	Fields           []*FieldError `protobuf:"bytes,4,rep,name=fields" json:"fields,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Error) GetCode() ErrorCode {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return ErrorCode_INTERNAL_SERVER_ERROR
}

func (m *Error) GetTitle() string {
	if m != nil && m.Title != nil {
		return *m.Title
	}
	return ""
}

func (m *Error) GetDetail() string {
	if m != nil && m.Detail != nil {
		return *m.Detail
	}
	return ""
}

func (m *Error) GetFields() []*FieldError {
	if m != nil {
		return m.Fields
	}
	return nil
}

type FieldError struct {
	Field            *string `protobuf:"bytes,1,opt,name=field" json:"field,omitempty"`
	Code             *string `protobuf:"bytes,2,opt,name=code" json:"code,omitempty"`
	Title            *string `protobuf:"bytes,3,opt,name=title" json:"title,omitempty"`
	Detail           *string `protobuf:"bytes,4,opt,name=detail" json:"detail,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *FieldError) Reset()                    { *m = FieldError{} }
func (m *FieldError) String() string            { return proto.CompactTextString(m) }
func (*FieldError) ProtoMessage()               {}
func (*FieldError) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *FieldError) GetField() string {
	if m != nil && m.Field != nil {
		return *m.Field
	}
	return ""
}

func (m *FieldError) GetCode() string {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return ""
}

func (m *FieldError) GetTitle() string {
	if m != nil && m.Title != nil {
		return *m.Title
	}
	return ""
}

func (m *FieldError) GetDetail() string {
	if m != nil && m.Detail != nil {
		return *m.Detail
	}
	return ""
}

type ListOptions struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *ListOptions) Reset()                    { *m = ListOptions{} }
func (m *ListOptions) String() string            { return proto.CompactTextString(m) }
func (*ListOptions) ProtoMessage()               {}
func (*ListOptions) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type UpdatePageRoutesRequest struct {
	Routes           []*ketchup_models.Route `protobuf:"bytes,1,rep,name=routes" json:"routes,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *UpdatePageRoutesRequest) Reset()                    { *m = UpdatePageRoutesRequest{} }
func (m *UpdatePageRoutesRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdatePageRoutesRequest) ProtoMessage()               {}
func (*UpdatePageRoutesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *UpdatePageRoutesRequest) GetRoutes() []*ketchup_models.Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

type GetRenderedPageRequest struct {
	Options          *GetRenderedPageRequest_RenderedPageOptions `protobuf:"bytes,1,opt,name=options" json:"options,omitempty"`
	XXX_unrecognized []byte                                      `json:"-"`
}

func (m *GetRenderedPageRequest) Reset()                    { *m = GetRenderedPageRequest{} }
func (m *GetRenderedPageRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRenderedPageRequest) ProtoMessage()               {}
func (*GetRenderedPageRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GetRenderedPageRequest) GetOptions() *GetRenderedPageRequest_RenderedPageOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

type GetRenderedPageRequest_RenderedPageOptions struct {
	Filter           *GetRenderedPageRequest_RenderedPageFilter `protobuf:"varint,1,opt,name=filter,enum=ketchup.api.GetRenderedPageRequest_RenderedPageFilter" json:"filter,omitempty"`
	XXX_unrecognized []byte                                     `json:"-"`
}

func (m *GetRenderedPageRequest_RenderedPageOptions) Reset() {
	*m = GetRenderedPageRequest_RenderedPageOptions{}
}
func (m *GetRenderedPageRequest_RenderedPageOptions) String() string {
	return proto.CompactTextString(m)
}
func (*GetRenderedPageRequest_RenderedPageOptions) ProtoMessage() {}
func (*GetRenderedPageRequest_RenderedPageOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{4, 0}
}

func (m *GetRenderedPageRequest_RenderedPageOptions) GetFilter() GetRenderedPageRequest_RenderedPageFilter {
	if m != nil && m.Filter != nil {
		return *m.Filter
	}
	return GetRenderedPageRequest_all
}

type GetRenderedPageResponse struct {
	Data             *structpb.Struct `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *GetRenderedPageResponse) Reset()                    { *m = GetRenderedPageResponse{} }
func (m *GetRenderedPageResponse) String() string            { return proto.CompactTextString(m) }
func (*GetRenderedPageResponse) ProtoMessage()               {}
func (*GetRenderedPageResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *GetRenderedPageResponse) GetData() *structpb.Struct {
	if m != nil {
		return m.Data
	}
	return nil
}

type ListPageRequest struct {
	List             *ListOptions                     `protobuf:"bytes,1,opt,name=list" json:"list,omitempty"`
	Options          *ListPageRequest_ListPageOptions `protobuf:"bytes,2,opt,name=options" json:"options,omitempty"`
	XXX_unrecognized []byte                           `json:"-"`
}

func (m *ListPageRequest) Reset()                    { *m = ListPageRequest{} }
func (m *ListPageRequest) String() string            { return proto.CompactTextString(m) }
func (*ListPageRequest) ProtoMessage()               {}
func (*ListPageRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ListPageRequest) GetList() *ListOptions {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *ListPageRequest) GetOptions() *ListPageRequest_ListPageOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

type ListPageRequest_ListPageOptions struct {
	Filter           *ListPageRequest_ListPageFilter `protobuf:"varint,1,opt,name=filter,enum=ketchup.api.ListPageRequest_ListPageFilter" json:"filter,omitempty"`
	XXX_unrecognized []byte                          `json:"-"`
}

func (m *ListPageRequest_ListPageOptions) Reset()         { *m = ListPageRequest_ListPageOptions{} }
func (m *ListPageRequest_ListPageOptions) String() string { return proto.CompactTextString(m) }
func (*ListPageRequest_ListPageOptions) ProtoMessage()    {}
func (*ListPageRequest_ListPageOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{6, 0}
}

func (m *ListPageRequest_ListPageOptions) GetFilter() ListPageRequest_ListPageFilter {
	if m != nil && m.Filter != nil {
		return *m.Filter
	}
	return ListPageRequest_all
}

type ListPageResponse struct {
	Pages            []*ketchup_models.Page `protobuf:"bytes,1,rep,name=pages" json:"pages,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *ListPageResponse) Reset()                    { *m = ListPageResponse{} }
func (m *ListPageResponse) String() string            { return proto.CompactTextString(m) }
func (*ListPageResponse) ProtoMessage()               {}
func (*ListPageResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ListPageResponse) GetPages() []*ketchup_models.Page {
	if m != nil {
		return m.Pages
	}
	return nil
}

type ListRouteResponse struct {
	Routes           []*ketchup_models.Route `protobuf:"bytes,1,rep,name=routes" json:"routes,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *ListRouteResponse) Reset()                    { *m = ListRouteResponse{} }
func (m *ListRouteResponse) String() string            { return proto.CompactTextString(m) }
func (*ListRouteResponse) ProtoMessage()               {}
func (*ListRouteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ListRouteResponse) GetRoutes() []*ketchup_models.Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

type ListThemeResponse struct {
	Themes           []*ketchup_models.Theme `protobuf:"bytes,1,rep,name=themes" json:"themes,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *ListThemeResponse) Reset()                    { *m = ListThemeResponse{} }
func (m *ListThemeResponse) String() string            { return proto.CompactTextString(m) }
func (*ListThemeResponse) ProtoMessage()               {}
func (*ListThemeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *ListThemeResponse) GetThemes() []*ketchup_models.Theme {
	if m != nil {
		return m.Themes
	}
	return nil
}

type UpdateDataRequest struct {
	Data             []*ketchup_models.Data `protobuf:"bytes,1,rep,name=data" json:"data,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *UpdateDataRequest) Reset()                    { *m = UpdateDataRequest{} }
func (m *UpdateDataRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateDataRequest) ProtoMessage()               {}
func (*UpdateDataRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *UpdateDataRequest) GetData() []*ketchup_models.Data {
	if m != nil {
		return m.Data
	}
	return nil
}

type ListDataResponse struct {
	Data             []*ketchup_models.Data `protobuf:"bytes,1,rep,name=data" json:"data,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *ListDataResponse) Reset()                    { *m = ListDataResponse{} }
func (m *ListDataResponse) String() string            { return proto.CompactTextString(m) }
func (*ListDataResponse) ProtoMessage()               {}
func (*ListDataResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *ListDataResponse) GetData() []*ketchup_models.Data {
	if m != nil {
		return m.Data
	}
	return nil
}

type TLSSettingsReponse struct {
	TlsEmail         *string `protobuf:"bytes,1,opt,name=tls_email,json=tlsEmail" json:"tls_email,omitempty"`
	TlsDomain        *string `protobuf:"bytes,2,opt,name=tls_domain,json=tlsDomain" json:"tls_domain,omitempty"`
	AgreedOn         *string `protobuf:"bytes,3,opt,name=agreed_on,json=agreedOn" json:"agreed_on,omitempty"`
	TermsOfService   *string `protobuf:"bytes,4,opt,name=terms_of_service,json=termsOfService" json:"terms_of_service,omitempty"`
	HasCertificate   *bool   `protobuf:"varint,5,opt,name=has_certificate,json=hasCertificate" json:"has_certificate,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TLSSettingsReponse) Reset()                    { *m = TLSSettingsReponse{} }
func (m *TLSSettingsReponse) String() string            { return proto.CompactTextString(m) }
func (*TLSSettingsReponse) ProtoMessage()               {}
func (*TLSSettingsReponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *TLSSettingsReponse) GetTlsEmail() string {
	if m != nil && m.TlsEmail != nil {
		return *m.TlsEmail
	}
	return ""
}

func (m *TLSSettingsReponse) GetTlsDomain() string {
	if m != nil && m.TlsDomain != nil {
		return *m.TlsDomain
	}
	return ""
}

func (m *TLSSettingsReponse) GetAgreedOn() string {
	if m != nil && m.AgreedOn != nil {
		return *m.AgreedOn
	}
	return ""
}

func (m *TLSSettingsReponse) GetTermsOfService() string {
	if m != nil && m.TermsOfService != nil {
		return *m.TermsOfService
	}
	return ""
}

func (m *TLSSettingsReponse) GetHasCertificate() bool {
	if m != nil && m.HasCertificate != nil {
		return *m.HasCertificate
	}
	return false
}

type EnableTLSRequest struct {
	TlsEmail         *string `protobuf:"bytes,1,opt,name=tls_email,json=tlsEmail" json:"tls_email,omitempty"`
	TlsDomain        *string `protobuf:"bytes,2,opt,name=tls_domain,json=tlsDomain" json:"tls_domain,omitempty"`
	Agreed           *bool   `protobuf:"varint,3,opt,name=agreed" json:"agreed,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *EnableTLSRequest) Reset()                    { *m = EnableTLSRequest{} }
func (m *EnableTLSRequest) String() string            { return proto.CompactTextString(m) }
func (*EnableTLSRequest) ProtoMessage()               {}
func (*EnableTLSRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *EnableTLSRequest) GetTlsEmail() string {
	if m != nil && m.TlsEmail != nil {
		return *m.TlsEmail
	}
	return ""
}

func (m *EnableTLSRequest) GetTlsDomain() string {
	if m != nil && m.TlsDomain != nil {
		return *m.TlsDomain
	}
	return ""
}

func (m *EnableTLSRequest) GetAgreed() bool {
	if m != nil && m.Agreed != nil {
		return *m.Agreed
	}
	return false
}

type InstallThemeRequest struct {
	Name             *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *InstallThemeRequest) Reset()                    { *m = InstallThemeRequest{} }
func (m *InstallThemeRequest) String() string            { return proto.CompactTextString(m) }
func (*InstallThemeRequest) ProtoMessage()               {}
func (*InstallThemeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *InstallThemeRequest) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Error)(nil), "ketchup.api.Error")
	proto.RegisterType((*FieldError)(nil), "ketchup.api.FieldError")
	proto.RegisterType((*ListOptions)(nil), "ketchup.api.ListOptions")
	proto.RegisterType((*UpdatePageRoutesRequest)(nil), "ketchup.api.UpdatePageRoutesRequest")
	proto.RegisterType((*GetRenderedPageRequest)(nil), "ketchup.api.GetRenderedPageRequest")
	proto.RegisterType((*GetRenderedPageRequest_RenderedPageOptions)(nil), "ketchup.api.GetRenderedPageRequest.RenderedPageOptions")
	proto.RegisterType((*GetRenderedPageResponse)(nil), "ketchup.api.GetRenderedPageResponse")
	proto.RegisterType((*ListPageRequest)(nil), "ketchup.api.ListPageRequest")
	proto.RegisterType((*ListPageRequest_ListPageOptions)(nil), "ketchup.api.ListPageRequest.ListPageOptions")
	proto.RegisterType((*ListPageResponse)(nil), "ketchup.api.ListPageResponse")
	proto.RegisterType((*ListRouteResponse)(nil), "ketchup.api.ListRouteResponse")
	proto.RegisterType((*ListThemeResponse)(nil), "ketchup.api.ListThemeResponse")
	proto.RegisterType((*UpdateDataRequest)(nil), "ketchup.api.UpdateDataRequest")
	proto.RegisterType((*ListDataResponse)(nil), "ketchup.api.ListDataResponse")
	proto.RegisterType((*TLSSettingsReponse)(nil), "ketchup.api.TLSSettingsReponse")
	proto.RegisterType((*EnableTLSRequest)(nil), "ketchup.api.EnableTLSRequest")
	proto.RegisterType((*InstallThemeRequest)(nil), "ketchup.api.InstallThemeRequest")
	proto.RegisterEnum("ketchup.api.ErrorCode", ErrorCode_name, ErrorCode_value)
	proto.RegisterEnum("ketchup.api.GetRenderedPageRequest_RenderedPageFilter", GetRenderedPageRequest_RenderedPageFilter_name, GetRenderedPageRequest_RenderedPageFilter_value)
	proto.RegisterEnum("ketchup.api.ListPageRequest_ListPageFilter", ListPageRequest_ListPageFilter_name, ListPageRequest_ListPageFilter_value)
}
func (m *Error) SetCode(v *ErrorCode) {
	m.Code = v
}

func (m *Error) SetTitle(v *string) {
	m.Title = v
}

func (m *Error) SetDetail(v *string) {
	m.Detail = v
}

func (m *Error) SetFields(v []*FieldError) {
	m.Fields = v
}

func (m *FieldError) SetField(v *string) {
	m.Field = v
}

func (m *FieldError) SetCode(v *string) {
	m.Code = v
}

func (m *FieldError) SetTitle(v *string) {
	m.Title = v
}

func (m *FieldError) SetDetail(v *string) {
	m.Detail = v
}

func (m *UpdatePageRoutesRequest) SetRoutes(v []*ketchup_models.Route) {
	m.Routes = v
}

func (m *GetRenderedPageRequest) SetOptions(v *GetRenderedPageRequest_RenderedPageOptions) {
	m.Options = v
}

func (m *GetRenderedPageResponse) SetData(v *structpb.Struct) {
	m.Data = v
}

func (m *ListPageRequest) SetList(v *ListOptions) {
	m.List = v
}

func (m *ListPageRequest) SetOptions(v *ListPageRequest_ListPageOptions) {
	m.Options = v
}

func (m *ListPageResponse) SetPages(v []*ketchup_models.Page) {
	m.Pages = v
}

func (m *ListRouteResponse) SetRoutes(v []*ketchup_models.Route) {
	m.Routes = v
}

func (m *ListThemeResponse) SetThemes(v []*ketchup_models.Theme) {
	m.Themes = v
}

func (m *UpdateDataRequest) SetData(v []*ketchup_models.Data) {
	m.Data = v
}

func (m *ListDataResponse) SetData(v []*ketchup_models.Data) {
	m.Data = v
}

func (m *TLSSettingsReponse) SetTlsEmail(v *string) {
	m.TlsEmail = v
}

func (m *TLSSettingsReponse) SetTlsDomain(v *string) {
	m.TlsDomain = v
}

func (m *TLSSettingsReponse) SetAgreedOn(v *string) {
	m.AgreedOn = v
}

func (m *TLSSettingsReponse) SetTermsOfService(v *string) {
	m.TermsOfService = v
}

func (m *TLSSettingsReponse) SetHasCertificate(v *bool) {
	m.HasCertificate = v
}

func (m *EnableTLSRequest) SetTlsEmail(v *string) {
	m.TlsEmail = v
}

func (m *EnableTLSRequest) SetTlsDomain(v *string) {
	m.TlsDomain = v
}

func (m *EnableTLSRequest) SetAgreed(v *bool) {
	m.Agreed = v
}

func (m *InstallThemeRequest) SetName(v *string) {
	m.Name = v
}

func init() { proto.RegisterFile("api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 755 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0x5d, 0x6f, 0xda, 0x48,
	0x14, 0x95, 0xf9, 0x0a, 0x5c, 0x36, 0xc4, 0x99, 0x24, 0xc0, 0x66, 0xb5, 0x12, 0xf2, 0xcb, 0xb2,
	0x49, 0xd6, 0x2b, 0x65, 0xa5, 0xcd, 0xcb, 0x6e, 0xa5, 0x26, 0x98, 0x36, 0x12, 0x82, 0x76, 0x20,
	0x79, 0xb5, 0x26, 0xf8, 0x02, 0x56, 0x8d, 0xed, 0x7a, 0x86, 0xfe, 0x89, 0xbe, 0x55, 0xfd, 0x41,
	0xfd, 0x69, 0x95, 0x67, 0xc6, 0x7c, 0x57, 0x4a, 0xd4, 0x37, 0xdf, 0x3b, 0xe7, 0x9e, 0x7b, 0xe6,
	0xcc, 0xe5, 0x02, 0x15, 0x16, 0xfb, 0x76, 0x9c, 0x44, 0x22, 0x22, 0xd5, 0x0f, 0x28, 0xc6, 0xb3,
	0x45, 0x6c, 0xb3, 0xd8, 0x3f, 0x87, 0x98, 0x4d, 0x51, 0x1d, 0x9c, 0x57, 0x93, 0x68, 0x21, 0x96,
	0x81, 0x98, 0xe1, 0x3c, 0x0b, 0x7e, 0xe1, 0x22, 0x59, 0x8c, 0x85, 0x8e, 0xc0, 0x63, 0x82, 0xa9,
	0x6f, 0xeb, 0x8b, 0x01, 0x45, 0x27, 0x49, 0xa2, 0x84, 0x5c, 0x40, 0x61, 0x1c, 0x79, 0xd8, 0x34,
	0x5a, 0x46, 0xbb, 0x76, 0x5d, 0xb7, 0xd7, 0xba, 0xd8, 0x12, 0x71, 0x17, 0x79, 0x48, 0x25, 0x86,
	0x9c, 0x42, 0x51, 0xf8, 0x22, 0xc0, 0x66, 0xae, 0x65, 0xb4, 0x2b, 0x54, 0x05, 0xa4, 0x0e, 0x25,
	0x0f, 0x05, 0xf3, 0x83, 0x66, 0x5e, 0xa6, 0x75, 0x44, 0xfe, 0x86, 0xd2, 0xc4, 0xc7, 0xc0, 0xe3,
	0xcd, 0x42, 0x2b, 0xdf, 0xae, 0x5e, 0x37, 0x36, 0xb8, 0xbb, 0xe9, 0x91, 0x6c, 0x40, 0x35, 0xcc,
	0xf2, 0x00, 0x56, 0xd9, 0xb4, 0x99, 0xcc, 0x4b, 0x65, 0x15, 0xaa, 0x02, 0x42, 0xb4, 0x5c, 0xa5,
	0x60, 0x4b, 0x56, 0x7e, 0xbf, 0xac, 0xc2, 0xba, 0x2c, 0xeb, 0x10, 0xaa, 0x3d, 0x9f, 0x8b, 0x41,
	0x2c, 0xfc, 0x28, 0xe4, 0xd6, 0x5b, 0x68, 0x3c, 0xc4, 0x1e, 0x13, 0xf8, 0x8e, 0x4d, 0x91, 0xa6,
	0x4e, 0x72, 0x8a, 0x1f, 0x17, 0xc8, 0x05, 0xf9, 0x0b, 0x4a, 0xd2, 0x5a, 0xde, 0x34, 0xe4, 0x05,
	0xce, 0x96, 0x17, 0x98, 0x47, 0x1e, 0x06, 0xdc, 0x96, 0x70, 0xaa, 0x41, 0xd6, 0xe7, 0x1c, 0xd4,
	0xdf, 0xa0, 0xa0, 0x18, 0x7a, 0x98, 0xa0, 0x27, 0xf9, 0x34, 0xd3, 0x7b, 0x38, 0x88, 0x54, 0x3f,
	0x79, 0x9b, 0xea, 0xf5, 0xcd, 0x86, 0x17, 0xfb, 0xab, 0xec, 0xf5, 0x9c, 0x96, 0x4b, 0x33, 0x9e,
	0x73, 0x84, 0x93, 0x3d, 0xe7, 0xa4, 0x9f, 0x9a, 0x1e, 0x08, 0x4c, 0xf4, 0x83, 0xfe, 0xfb, 0xd2,
	0x46, 0x5d, 0x59, 0x4d, 0x35, 0x8b, 0x75, 0x03, 0x64, 0xf7, 0x94, 0x1c, 0x40, 0x9e, 0x05, 0x81,
	0x69, 0x90, 0x43, 0xa8, 0xc4, 0x8b, 0xa7, 0xc0, 0xe7, 0x33, 0xf4, 0xcc, 0x1c, 0xa9, 0x40, 0xd1,
	0x4b, 0xd8, 0x44, 0x98, 0x79, 0xab, 0x0b, 0x8d, 0x9d, 0x6e, 0x3c, 0x8e, 0x42, 0x8e, 0xe4, 0x12,
	0x0a, 0xe9, 0x28, 0x6a, 0x2b, 0x1a, 0xf6, 0x34, 0x8a, 0xa6, 0x81, 0x9e, 0xd9, 0xa7, 0xc5, 0xc4,
	0x1e, 0xca, 0xa9, 0xa5, 0x12, 0x64, 0x7d, 0xcd, 0xc1, 0x51, 0xfa, 0x5e, 0xeb, 0x76, 0x5e, 0x41,
	0x21, 0xf0, 0xb9, 0xd0, 0x04, 0xcd, 0x8d, 0x2b, 0xae, 0xbd, 0x2d, 0x95, 0x28, 0xd2, 0x5d, 0x99,
	0x9f, 0x93, 0x05, 0x57, 0x3b, 0x05, 0xeb, 0x66, 0x64, 0xf1, 0x8e, 0xe3, 0x8f, 0x2b, 0x21, 0x99,
	0xdb, 0x77, 0x5b, 0x6e, 0x5f, 0x3e, 0x8b, 0x79, 0xcb, 0xe2, 0x7f, 0xa0, 0xb6, 0x79, 0xf2, 0x1c,
	0x7b, 0x5f, 0x81, 0xb9, 0xa2, 0xd7, 0xbe, 0x5e, 0x40, 0x31, 0x5d, 0x0b, 0xd9, 0xb8, 0x9e, 0x6e,
	0x8f, 0xab, 0x04, 0x2b, 0x88, 0x75, 0x0b, 0xc7, 0x69, 0xbd, 0x9a, 0xe0, 0x8c, 0xe0, 0x85, 0x03,
	0xaf, 0x39, 0x46, 0xe9, 0xc6, 0x59, 0xe7, 0x90, 0x2b, 0xe8, 0x87, 0x1c, 0x0a, 0xae, 0x41, 0xd6,
	0xff, 0x70, 0xac, 0x7e, 0x7e, 0x1d, 0x26, 0x58, 0xf6, 0xbe, 0xed, 0xe5, 0x80, 0xec, 0xbd, 0x87,
	0x84, 0xaa, 0xe9, 0xf8, 0x4f, 0xd9, 0xa0, 0x8a, 0xb5, 0x82, 0xe7, 0x57, 0x7f, 0x33, 0x80, 0x8c,
	0x7a, 0xc3, 0x21, 0x0a, 0xe1, 0x87, 0x53, 0x4e, 0x51, 0x11, 0xfc, 0x06, 0x15, 0x11, 0x70, 0x17,
	0xe7, 0xe9, 0xf2, 0x50, 0xdb, 0xa7, 0x2c, 0x02, 0xee, 0xa4, 0x31, 0xf9, 0x1d, 0x20, 0x3d, 0xf4,
	0xa2, 0x39, 0xf3, 0x43, 0xbd, 0x86, 0x52, 0x78, 0x47, 0x26, 0xd2, 0x5a, 0x36, 0x4d, 0x10, 0x3d,
	0x37, 0x0a, 0xf5, 0x3e, 0x2a, 0xab, 0xc4, 0x20, 0x24, 0x6d, 0x30, 0x05, 0x26, 0x73, 0xee, 0x46,
	0x13, 0x97, 0x63, 0xf2, 0xc9, 0x1f, 0xa3, 0x5e, 0x4e, 0x35, 0x99, 0x1f, 0x4c, 0x86, 0x2a, 0x4b,
	0xfe, 0x80, 0xa3, 0x19, 0xe3, 0xee, 0x18, 0x13, 0xe1, 0x4f, 0xfc, 0x31, 0x13, 0xd8, 0x2c, 0xb6,
	0x8c, 0x76, 0x99, 0xd6, 0x66, 0x8c, 0xdf, 0xad, 0xb2, 0xd6, 0x04, 0x4c, 0x27, 0x64, 0x4f, 0x01,
	0x8e, 0x7a, 0xc3, 0xcc, 0xbe, 0x9f, 0xd1, 0x5f, 0x87, 0x92, 0x92, 0x2b, 0xc5, 0x97, 0xa9, 0x8e,
	0xac, 0x3f, 0xe1, 0xe4, 0x3e, 0xe4, 0x82, 0x05, 0x81, 0x7e, 0x6e, 0xd5, 0x8a, 0x40, 0x21, 0x64,
	0x73, 0xd4, 0x5d, 0xe4, 0xf7, 0x45, 0x07, 0x2a, 0xcb, 0x3f, 0x0e, 0xf2, 0x2b, 0x9c, 0xdd, 0xf7,
	0x47, 0x0e, 0xed, 0xbf, 0xee, 0xb9, 0x43, 0x87, 0x3e, 0x3a, 0xd4, 0x75, 0x28, 0x1d, 0x50, 0x35,
	0xdc, 0xfd, 0xc1, 0xc8, 0xed, 0x0e, 0x1e, 0xfa, 0x1d, 0x33, 0x97, 0x86, 0xdd, 0x01, 0xbd, 0xbd,
	0xef, 0x74, 0x9c, 0xbe, 0x99, 0xff, 0x1e, 0x00, 0x00, 0xff, 0xff, 0x40, 0x70, 0x52, 0x50, 0xfa,
	0x06, 0x00, 0x00,
}
