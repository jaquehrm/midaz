// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.0
// source: proto/account/account.proto

package account

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AccountID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AccountID) Reset() {
	*x = AccountID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_account_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountID) ProtoMessage() {}

func (x *AccountID) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_account_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountID.ProtoReflect.Descriptor instead.
func (*AccountID) Descriptor() ([]byte, []int) {
	return file_proto_account_account_proto_rawDescGZIP(), []int{0}
}

func (x *AccountID) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type AccountAlias struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Alias string `protobuf:"bytes,1,opt,name=alias,proto3" json:"alias,omitempty"`
}

func (x *AccountAlias) Reset() {
	*x = AccountAlias{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_account_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountAlias) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountAlias) ProtoMessage() {}

func (x *AccountAlias) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_account_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountAlias.ProtoReflect.Descriptor instead.
func (*AccountAlias) Descriptor() ([]byte, []int) {
	return file_proto_account_account_proto_rawDescGZIP(), []int{1}
}

func (x *AccountAlias) GetAlias() string {
	if x != nil {
		return x.Alias
	}
	return ""
}

type Balance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Available float64 `protobuf:"fixed64,1,opt,name=available,proto3" json:"available,omitempty"`
	OnHold    float64 `protobuf:"fixed64,2,opt,name=on_hold,json=onHold,proto3" json:"on_hold,omitempty"`
	Scale     float64 `protobuf:"fixed64,3,opt,name=scale,proto3" json:"scale,omitempty"`
}

func (x *Balance) Reset() {
	*x = Balance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_account_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Balance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Balance) ProtoMessage() {}

func (x *Balance) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_account_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Balance.ProtoReflect.Descriptor instead.
func (*Balance) Descriptor() ([]byte, []int) {
	return file_proto_account_account_proto_rawDescGZIP(), []int{2}
}

func (x *Balance) GetAvailable() float64 {
	if x != nil {
		return x.Available
	}
	return 0
}

func (x *Balance) GetOnHold() float64 {
	if x != nil {
		return x.OnHold
	}
	return 0
}

func (x *Balance) GetScale() float64 {
	if x != nil {
		return x.Scale
	}
	return 0
}

type Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value map[string]string `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Metadata) Reset() {
	*x = Metadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_account_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata) ProtoMessage() {}

func (x *Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_account_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadata.ProtoReflect.Descriptor instead.
func (*Metadata) Descriptor() ([]byte, []int) {
	return file_proto_account_account_proto_rawDescGZIP(), []int{3}
}

func (x *Metadata) GetValue() map[string]string {
	if x != nil {
		return x.Value
	}
	return nil
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code           string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Description    string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	AllowSending   bool   `protobuf:"varint,3,opt,name=allow_sending,json=allowSending,proto3" json:"allow_sending,omitempty"`
	AllowReceiving bool   `protobuf:"varint,4,opt,name=allow_receiving,json=allowReceiving,proto3" json:"allow_receiving,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_account_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_account_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_proto_account_account_proto_rawDescGZIP(), []int{4}
}

func (x *Status) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Status) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Status) GetAllowSending() bool {
	if x != nil {
		return x.AllowSending
	}
	return false
}

func (x *Status) GetAllowReceiving() bool {
	if x != nil {
		return x.AllowReceiving
	}
	return false
}

type Account struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name            string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ParentAccountId string    `protobuf:"bytes,3,opt,name=parent_account_id,json=parentAccountId,proto3" json:"parent_account_id,omitempty"`
	EntityId        string    `protobuf:"bytes,4,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	AssetCode       string    `protobuf:"bytes,5,opt,name=asset_code,json=assetCode,proto3" json:"asset_code,omitempty"`
	OrganizationId  string    `protobuf:"bytes,6,opt,name=organization_id,json=organizationId,proto3" json:"organization_id,omitempty"`
	LedgerId        string    `protobuf:"bytes,7,opt,name=ledger_id,json=ledgerId,proto3" json:"ledger_id,omitempty"`
	PortfolioId     string    `protobuf:"bytes,8,opt,name=portfolio_id,json=portfolioId,proto3" json:"portfolio_id,omitempty"`
	ProductId       string    `protobuf:"bytes,9,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Balance         *Balance  `protobuf:"bytes,10,opt,name=balance,proto3" json:"balance,omitempty"`
	Status          *Status   `protobuf:"bytes,11,opt,name=status,proto3" json:"status,omitempty"`
	Alias           string    `protobuf:"bytes,12,opt,name=alias,proto3" json:"alias,omitempty"`
	Type            string    `protobuf:"bytes,13,opt,name=type,proto3" json:"type,omitempty"`
	CreatedAt       string    `protobuf:"bytes,14,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt       string    `protobuf:"bytes,15,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt       string    `protobuf:"bytes,16,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	Metadata        *Metadata `protobuf:"bytes,17,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *Account) Reset() {
	*x = Account{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_account_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Account) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account) ProtoMessage() {}

func (x *Account) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_account_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Account.ProtoReflect.Descriptor instead.
func (*Account) Descriptor() ([]byte, []int) {
	return file_proto_account_account_proto_rawDescGZIP(), []int{5}
}

func (x *Account) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Account) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Account) GetParentAccountId() string {
	if x != nil {
		return x.ParentAccountId
	}
	return ""
}

func (x *Account) GetEntityId() string {
	if x != nil {
		return x.EntityId
	}
	return ""
}

func (x *Account) GetAssetCode() string {
	if x != nil {
		return x.AssetCode
	}
	return ""
}

func (x *Account) GetOrganizationId() string {
	if x != nil {
		return x.OrganizationId
	}
	return ""
}

func (x *Account) GetLedgerId() string {
	if x != nil {
		return x.LedgerId
	}
	return ""
}

func (x *Account) GetPortfolioId() string {
	if x != nil {
		return x.PortfolioId
	}
	return ""
}

func (x *Account) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *Account) GetBalance() *Balance {
	if x != nil {
		return x.Balance
	}
	return nil
}

func (x *Account) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *Account) GetAlias() string {
	if x != nil {
		return x.Alias
	}
	return ""
}

func (x *Account) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Account) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Account) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Account) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

func (x *Account) GetMetadata() *Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Alias   string   `protobuf:"bytes,2,opt,name=alias,proto3" json:"alias,omitempty"`
	Balance *Balance `protobuf:"bytes,3,opt,name=balance,proto3" json:"balance,omitempty"`
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_account_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_account_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_proto_account_account_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateRequest) GetAlias() string {
	if x != nil {
		return x.Alias
	}
	return ""
}

func (x *UpdateRequest) GetBalance() *Balance {
	if x != nil {
		return x.Balance
	}
	return nil
}

type ManyAccountsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Accounts []*Account `protobuf:"bytes,1,rep,name=accounts,proto3" json:"accounts,omitempty"`
}

func (x *ManyAccountsResponse) Reset() {
	*x = ManyAccountsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_account_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManyAccountsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManyAccountsResponse) ProtoMessage() {}

func (x *ManyAccountsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_account_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManyAccountsResponse.ProtoReflect.Descriptor instead.
func (*ManyAccountsResponse) Descriptor() ([]byte, []int) {
	return file_proto_account_account_proto_rawDescGZIP(), []int{7}
}

func (x *ManyAccountsResponse) GetAccounts() []*Account {
	if x != nil {
		return x.Accounts
	}
	return nil
}

type ManyAccountsID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []*AccountID `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (x *ManyAccountsID) Reset() {
	*x = ManyAccountsID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_account_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManyAccountsID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManyAccountsID) ProtoMessage() {}

func (x *ManyAccountsID) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_account_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManyAccountsID.ProtoReflect.Descriptor instead.
func (*ManyAccountsID) Descriptor() ([]byte, []int) {
	return file_proto_account_account_proto_rawDescGZIP(), []int{8}
}

func (x *ManyAccountsID) GetIds() []*AccountID {
	if x != nil {
		return x.Ids
	}
	return nil
}

type ManyAccountsAlias struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Aliases []*AccountAlias `protobuf:"bytes,1,rep,name=aliases,proto3" json:"aliases,omitempty"`
}

func (x *ManyAccountsAlias) Reset() {
	*x = ManyAccountsAlias{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_account_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManyAccountsAlias) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManyAccountsAlias) ProtoMessage() {}

func (x *ManyAccountsAlias) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_account_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManyAccountsAlias.ProtoReflect.Descriptor instead.
func (*ManyAccountsAlias) Descriptor() ([]byte, []int) {
	return file_proto_account_account_proto_rawDescGZIP(), []int{9}
}

func (x *ManyAccountsAlias) GetAliases() []*AccountAlias {
	if x != nil {
		return x.Aliases
	}
	return nil
}

var File_proto_account_account_proto protoreflect.FileDescriptor

var file_proto_account_account_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x1b, 0x0a, 0x09, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x24, 0x0a, 0x0c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x6c,
	0x69, 0x61, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x22, 0x56, 0x0a, 0x07, 0x42, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62,
	0x6c, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x6f, 0x6e, 0x5f, 0x68, 0x6f, 0x6c, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x06, 0x6f, 0x6e, 0x48, 0x6f, 0x6c, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73,
	0x63, 0x61, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x73, 0x63, 0x61, 0x6c,
	0x65, 0x22, 0x78, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x32, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x1a, 0x38, 0x0a, 0x0a, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x8c, 0x01, 0x0a, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d,
	0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0c, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x53, 0x65, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x72, 0x65, 0x63, 0x65, 0x69,
	0x76, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x61, 0x6c, 0x6c, 0x6f,
	0x77, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x22, 0xa8, 0x04, 0x0a, 0x07, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x73, 0x73, 0x65, 0x74, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6f, 0x72, 0x67,
	0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6c,
	0x65, 0x64, 0x67, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x6f, 0x72, 0x74,
	0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x70, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x07, 0x62, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x07, 0x62,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x61, 0x6c, 0x69, 0x61, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x2d, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x61, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x12, 0x2a, 0x0a, 0x07,
	0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52,
	0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x44, 0x0a, 0x14, 0x4d, 0x61, 0x6e, 0x79,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2c, 0x0a, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x52, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x22, 0x36,
	0x0a, 0x0e, 0x4d, 0x61, 0x6e, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x49, 0x44,
	0x12, 0x24, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49,
	0x44, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x44, 0x0a, 0x11, 0x4d, 0x61, 0x6e, 0x79, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x41, 0x6c, 0x69, 0x61, 0x73, 0x12, 0x2f, 0x0a, 0x07, 0x61,
	0x6c, 0x69, 0x61, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x6c,
	0x69, 0x61, 0x73, 0x52, 0x07, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x65, 0x73, 0x32, 0xd5, 0x01, 0x0a,
	0x0c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x44, 0x0a,
	0x08, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x73, 0x12, 0x17, 0x2e, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x2e, 0x4d, 0x61, 0x6e, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73,
	0x49, 0x44, 0x1a, 0x1d, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x4d, 0x61, 0x6e,
	0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x42, 0x79, 0x41, 0x6c, 0x69, 0x61,
	0x73, 0x12, 0x1a, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x4d, 0x61, 0x6e, 0x79,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x41, 0x6c, 0x69, 0x61, 0x73, 0x1a, 0x1d, 0x2e,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x4d, 0x61, 0x6e, 0x79, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x34,
	0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x10, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_account_account_proto_rawDescOnce sync.Once
	file_proto_account_account_proto_rawDescData = file_proto_account_account_proto_rawDesc
)

func file_proto_account_account_proto_rawDescGZIP() []byte {
	file_proto_account_account_proto_rawDescOnce.Do(func() {
		file_proto_account_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_account_account_proto_rawDescData)
	})
	return file_proto_account_account_proto_rawDescData
}

var file_proto_account_account_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_account_account_proto_goTypes = []any{
	(*AccountID)(nil),            // 0: account.AccountID
	(*AccountAlias)(nil),         // 1: account.AccountAlias
	(*Balance)(nil),              // 2: account.Balance
	(*Metadata)(nil),             // 3: account.Metadata
	(*Status)(nil),               // 4: account.Status
	(*Account)(nil),              // 5: account.Account
	(*UpdateRequest)(nil),        // 6: account.UpdateRequest
	(*ManyAccountsResponse)(nil), // 7: account.ManyAccountsResponse
	(*ManyAccountsID)(nil),       // 8: account.ManyAccountsID
	(*ManyAccountsAlias)(nil),    // 9: account.ManyAccountsAlias
	nil,                          // 10: account.Metadata.ValueEntry
}
var file_proto_account_account_proto_depIdxs = []int32{
	10, // 0: account.Metadata.value:type_name -> account.Metadata.ValueEntry
	2,  // 1: account.Account.balance:type_name -> account.Balance
	4,  // 2: account.Account.status:type_name -> account.Status
	3,  // 3: account.Account.metadata:type_name -> account.Metadata
	2,  // 4: account.UpdateRequest.balance:type_name -> account.Balance
	5,  // 5: account.ManyAccountsResponse.accounts:type_name -> account.Account
	0,  // 6: account.ManyAccountsID.ids:type_name -> account.AccountID
	1,  // 7: account.ManyAccountsAlias.aliases:type_name -> account.AccountAlias
	8,  // 8: account.AccountProto.GetByIds:input_type -> account.ManyAccountsID
	9,  // 9: account.AccountProto.GetByAlias:input_type -> account.ManyAccountsAlias
	6,  // 10: account.AccountProto.Update:input_type -> account.UpdateRequest
	7,  // 11: account.AccountProto.GetByIds:output_type -> account.ManyAccountsResponse
	7,  // 12: account.AccountProto.GetByAlias:output_type -> account.ManyAccountsResponse
	5,  // 13: account.AccountProto.Update:output_type -> account.Account
	11, // [11:14] is the sub-list for method output_type
	8,  // [8:11] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_proto_account_account_proto_init() }
func file_proto_account_account_proto_init() {
	if File_proto_account_account_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_account_account_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*AccountID); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_account_account_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*AccountAlias); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_account_account_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Balance); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_account_account_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Metadata); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_account_account_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*Status); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_account_account_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*Account); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_account_account_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_account_account_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*ManyAccountsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_account_account_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*ManyAccountsID); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_account_account_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*ManyAccountsAlias); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_account_account_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_account_account_proto_goTypes,
		DependencyIndexes: file_proto_account_account_proto_depIdxs,
		MessageInfos:      file_proto_account_account_proto_msgTypes,
	}.Build()
	File_proto_account_account_proto = out.File
	file_proto_account_account_proto_rawDesc = nil
	file_proto_account_account_proto_goTypes = nil
	file_proto_account_account_proto_depIdxs = nil
}
