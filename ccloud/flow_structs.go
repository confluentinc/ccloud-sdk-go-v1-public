package ccloud

import (
	_ "github.com/confluentinc/proto-go-setter"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
)

type AuthenticateRequest struct {
	Email                string            `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty" db:"email,omitempty" url:"email,omitempty"`
	Password             string            `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty" redact:"-" db:"password,omitempty" url:"password,omitempty"`
	RequestCarrier       map[string]string `protobuf:"bytes,3,rep,name=request_carrier,json=requestCarrier,proto3" json:"request_carrier,omitempty" db:"request_carrier,omitempty" url:"request_carrier,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	IdToken              string            `protobuf:"bytes,4,opt,name=id_token,json=idToken,proto3" json:"id_token,omitempty" db:"id_token,omitempty" url:"id_token,omitempty"`
	ExternalOauthToken   string            `protobuf:"bytes,5,opt,name=external_oauth_token,json=externalOauthToken,proto3" json:"external_oauth_token,omitempty" db:"external_oauth_token,omitempty" url:"external_oauth_token,omitempty"`
	OrgResourceId        string            `protobuf:"bytes,6,opt,name=org_resource_id,json=orgResourceId,proto3" json:"org_resource_id,omitempty" db:"org_resource_id,omitempty" url:"org_resource_id,omitempty"`
	IdentityPool         string            `protobuf:"bytes,7,opt,name=identity_pool,json=identityPool,proto3" json:"identity_pool,omitempty" db:"identity_pool,omitempty" url:"identity_pool,omitempty"`
	RefreshToken         string            `protobuf:"bytes,8,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty" db:"refresh_token,omitempty" url:"refresh_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *AuthenticateRequest) Reset()         { *m = AuthenticateRequest{} }
func (m *AuthenticateRequest) String() string { return proto.CompactTextString(m) }
func (*AuthenticateRequest) ProtoMessage()    {}

type AuthenticateReply struct {
	Token                string        `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty" db:"token,omitempty" url:"token,omitempty"`
	Error                *Error        `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty" db:"error,omitempty" url:"error,omitempty"`
	User                 *User         `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty" db:"user,omitempty" url:"user,omitempty"`
	Organization         *Organization `protobuf:"bytes,4,opt,name=organization,proto3" json:"organization,omitempty" db:"organization,omitempty" url:"organization,omitempty"`
	RefreshToken         string        `protobuf:"bytes,5,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty" db:"refresh_token,omitempty" url:"refresh_token,omitempty"`
	IdentityProvider     string        `protobuf:"bytes,6,opt,name=identity_provider,json=identityProvider,proto3" json:"identity_provider,omitempty" db:"identity_provider,omitempty" url:"identity_provider,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *AuthenticateReply) Reset()         { *m = AuthenticateReply{} }
func (m *AuthenticateReply) String() string { return proto.CompactTextString(m) }
func (*AuthenticateReply) ProtoMessage()    {}

func (m *AuthenticateReply) GetOrganization() *Organization {
	if m != nil {
		return m.Organization
	}
	return nil
}

func (m *AuthenticateReply) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *AuthenticateReply) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

type GetMeReply struct {
	Error                *Error        `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty" db:"error,omitempty" url:"error,omitempty"`
	User                 *User         `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty" db:"user,omitempty" url:"user,omitempty"`
	Organization         *Organization `protobuf:"bytes,3,opt,name=organization,proto3" json:"organization,omitempty" db:"organization,omitempty" url:"organization,omitempty"`
	Accounts             []*Account    `protobuf:"bytes,4,rep,name=accounts,proto3" json:"accounts,omitempty" db:"accounts,omitempty" url:"accounts,omitempty"`
	Account              *Account      `protobuf:"bytes,5,opt,name=account,proto3" json:"account,omitempty" db:"account,omitempty" url:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetMeReply) Reset()         { *m = GetMeReply{} }
func (m *GetMeReply) String() string { return proto.CompactTextString(m) }
func (*GetMeReply) ProtoMessage()    {}

func (m *GetMeReply) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *GetMeReply) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *GetMeReply) GetOrganization() *Organization {
	if m != nil {
		return m.Organization
	}
	return nil
}

func (m *GetMeReply) GetAccounts() []*Account {
	if m != nil {
		return m.Accounts
	}
	return nil
}

type CreateConnectAuthTokenReply struct {
	Error                *Error   `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty" db:"error,omitempty" url:"error,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty" db:"token,omitempty" url:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateConnectAuthTokenReply) Reset()         { *m = CreateConnectAuthTokenReply{} }
func (m *CreateConnectAuthTokenReply) String() string { return proto.CompactTextString(m) }
func (*CreateConnectAuthTokenReply) ProtoMessage()    {}

type CreateDecisionEngineAuthTokenReply struct {
	Error                *Error   `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty" db:"error,omitempty" url:"error,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty" db:"token,omitempty" url:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateDecisionEngineAuthTokenReply) Reset()         { *m = CreateDecisionEngineAuthTokenReply{} }
func (m *CreateDecisionEngineAuthTokenReply) String() string { return proto.CompactTextString(m) }
func (*CreateDecisionEngineAuthTokenReply) ProtoMessage()    {}
