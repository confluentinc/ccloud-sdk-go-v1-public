package ccloud

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
	io "io"
	math "math"

	_ "github.com/confluentinc/proto-go-setter"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
)

type TaxAddress struct {
	Line1 string `protobuf:"bytes,1,opt,name=line1,proto3" json:"line1,omitempty" db:"line1,omitempty" url:"line1,omitempty"`
	Line2 string `protobuf:"bytes,2,opt,name=line2,proto3" json:"line2,omitempty" db:"line2,omitempty" url:"line2,omitempty"`
	City  string `protobuf:"bytes,3,opt,name=city,proto3" json:"city,omitempty" db:"city,omitempty" url:"city,omitempty"`
	State string `protobuf:"bytes,4,opt,name=state,proto3" json:"state,omitempty" db:"state,omitempty" url:"state,omitempty"`
	// country must follow the ISO 3166-1 alpha-2 standard
	// https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2
	Country              string   `protobuf:"bytes,5,opt,name=country,proto3" json:"country,omitempty" db:"country,omitempty" url:"country,omitempty"`
	PostalCode           string   `protobuf:"bytes,6,opt,name=postal_code,json=postalCode,proto3" json:"postal_code,omitempty" db:"postal_code,omitempty" url:"postal_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

type TaxId struct {
	// details of taxId definition is here:
	// https://stripe.com/docs/api/customer_tax_ids/object
	TypeEnum             string   `protobuf:"bytes,1,opt,name=type_enum,json=typeEnum,proto3" json:"type_enum,omitempty" db:"type_enum,omitempty" url:"type_enum,omitempty"`
	TaxId                string   `protobuf:"bytes,2,opt,name=tax_id,json=taxId,proto3" json:"tax_id,omitempty" db:"tax_id,omitempty" url:"tax_id,omitempty"`
	StripeObjectId       string   `protobuf:"bytes,3,opt,name=stripe_object_id,json=stripeObjectId,proto3" json:"stripe_object_id,omitempty" redact:"-" db:"stripe_object_id,omitempty" url:"stripe_object_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

type PromoCodeClaim struct {
	OrgId                int32            `protobuf:"varint,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty" db:"org_id,omitempty" url:"org_id,omitempty"`
	Code                 string           `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty" db:"code,omitempty" url:"code,omitempty"`
	Amount               int64            `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty" db:"amount,omitempty" url:"amount,omitempty"`
	Balance              int64            `protobuf:"varint,4,opt,name=balance,proto3" json:"balance,omitempty" db:"balance,omitempty" url:"balance,omitempty"`
	ClaimDate            *types.Timestamp `protobuf:"bytes,5,opt,name=claim_date,json=claimDate,proto3" json:"claim_date,omitempty" db:"claim_date,omitempty" url:"claim_date,omitempty"`
	CreditExpirationDate *types.Timestamp `protobuf:"bytes,6,opt,name=credit_expiration_date,json=creditExpirationDate,proto3" json:"credit_expiration_date,omitempty" db:"credit_expiration_date,omitempty" url:"credit_expiration_date,omitempty"`
	ClaimedBy            int32            `protobuf:"varint,7,opt,name=claimed_by,json=claimedBy,proto3" json:"claimed_by,omitempty" db:"claimed_by,omitempty" url:"claimed_by,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *PromoCodeClaim) Reset()         { *m = PromoCodeClaim{} }
func (m *PromoCodeClaim) String() string { return proto.CompactTextString(m) }
func (*PromoCodeClaim) ProtoMessage()    {}

type Card struct {
	Cardholder           string   `protobuf:"bytes,1,opt,name=cardholder,proto3" json:"cardholder,omitempty" db:"cardholder,omitempty" url:"cardholder,omitempty"`
	Brand                string   `protobuf:"bytes,2,opt,name=brand,proto3" json:"brand,omitempty" db:"brand,omitempty" url:"brand,omitempty"`
	Last4                string   `protobuf:"bytes,3,opt,name=last4,proto3" json:"last4,omitempty" db:"last4,omitempty" url:"last4,omitempty"`
	ExpMonth             string   `protobuf:"bytes,4,opt,name=exp_month,json=expMonth,proto3" json:"exp_month,omitempty" db:"exp_month,omitempty" url:"exp_month,omitempty"`
	ExpYear              string   `protobuf:"bytes,5,opt,name=exp_year,json=expYear,proto3" json:"exp_year,omitempty" db:"exp_year,omitempty" url:"exp_year,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Card) Reset()         { *m = Card{} }
func (m *Card) String() string { return proto.CompactTextString(m) }
func (*Card) ProtoMessage()    {}

func (m *Card) GetBrand() string {
	if m != nil {
		return m.Brand
	}
	return ""
}

func (m *Card) GetLast4() string {
	if m != nil {
		return m.Last4
	}
	return ""
}

type PriceTable struct {
	PriceTable           map[string]*UnitPrices `protobuf:"bytes,1,rep,name=price_table,json=priceTable,proto3" json:"price_table,omitempty" db:"price_table,omitempty" url:"price_table,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *PriceTable) Reset()         { *m = PriceTable{} }
func (m *PriceTable) String() string { return proto.CompactTextString(m) }
func (*PriceTable) ProtoMessage()    {}

type UnitPrices struct {
	Prices               map[string]float64 `protobuf:"bytes,1,rep,name=prices,proto3" json:"prices,omitempty" db:"prices,omitempty" url:"prices,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
	Unit                 string             `protobuf:"bytes,2,opt,name=unit,proto3" json:"unit,omitempty" db:"unit,omitempty" url:"unit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *UnitPrices) Reset()         { *m = UnitPrices{} }
func (m *UnitPrices) String() string { return proto.CompactTextString(m) }
func (*UnitPrices) ProtoMessage()    {}

type GetPriceTableReply struct {
	PriceTable           *PriceTable `protobuf:"bytes,1,opt,name=price_table,json=priceTable,proto3" json:"price_table,omitempty" db:"price_table,omitempty" url:"price_table,omitempty"`
	Error                *Error      `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty" db:"error,omitempty" url:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GetPriceTableReply) Reset()         { *m = GetPriceTableReply{} }
func (m *GetPriceTableReply) String() string { return proto.CompactTextString(m) }
func (*GetPriceTableReply) ProtoMessage()    {}

func (m *GetPriceTableReply) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type GetPaymentInfoReply struct {
	Card                 *Card         `protobuf:"bytes,1,opt,name=card,proto3" json:"card,omitempty" db:"card,omitempty" url:"card,omitempty"`
	Organization         *Organization `protobuf:"bytes,3,opt,name=organization,proto3" json:"organization,omitempty" db:"organization,omitempty" url:"organization,omitempty"`
	Error                *Error        `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty" db:"error,omitempty" url:"error,omitempty"`
	TaxAddress           *TaxAddress   `protobuf:"bytes,4,opt,name=tax_address,json=taxAddress,proto3" json:"tax_address,omitempty" db:"tax_address,omitempty" url:"tax_address,omitempty"`
	TaxIds               []*TaxId      `protobuf:"bytes,5,rep,name=tax_ids,json=taxIds,proto3" json:"tax_ids,omitempty" db:"tax_ids,omitempty" url:"tax_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetPaymentInfoReply) Reset()         { *m = GetPaymentInfoReply{} }
func (m *GetPaymentInfoReply) String() string { return proto.CompactTextString(m) }
func (*GetPaymentInfoReply) ProtoMessage()    {}

func (m *GetPaymentInfoReply) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type UpdatePaymentInfoRequest struct {
	StripeToken          string            `protobuf:"bytes,1,opt,name=stripe_token,json=stripeToken,proto3" json:"stripe_token,omitempty" redact:"-" db:"stripe_token,omitempty" url:"stripe_token,omitempty"`
	Organization         *Organization     `protobuf:"bytes,2,opt,name=organization,proto3" json:"organization,omitempty" db:"organization,omitempty" url:"organization,omitempty"`
	RequestCarrier       map[string]string `protobuf:"bytes,3,rep,name=request_carrier,json=requestCarrier,proto3" json:"request_carrier,omitempty" db:"request_carrier,omitempty" url:"request_carrier,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *UpdatePaymentInfoRequest) Reset()         { *m = UpdatePaymentInfoRequest{} }
func (m *UpdatePaymentInfoRequest) String() string { return proto.CompactTextString(m) }
func (*UpdatePaymentInfoRequest) ProtoMessage()    {}

type UpdatePaymentInfoReply struct {
	Card                 *Card         `protobuf:"bytes,1,opt,name=card,proto3" json:"card,omitempty" db:"card,omitempty" url:"card,omitempty"`
	Organization         *Organization `protobuf:"bytes,3,opt,name=organization,proto3" json:"organization,omitempty" db:"organization,omitempty" url:"organization,omitempty"`
	Error                *Error        `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty" db:"error,omitempty" url:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *UpdatePaymentInfoReply) Reset()         { *m = UpdatePaymentInfoReply{} }
func (m *UpdatePaymentInfoReply) String() string { return proto.CompactTextString(m) }
func (*UpdatePaymentInfoReply) ProtoMessage()    {}

func (m *UpdatePaymentInfoReply) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

// promo code api object for users
type ClaimPromoCodeRequest struct {
	OrgId                int32             `protobuf:"varint,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty" db:"org_id,omitempty" url:"org_id,omitempty"`
	Code                 string            `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty" db:"code,omitempty" url:"code,omitempty"`
	ClaimedBy            int32             `protobuf:"varint,3,opt,name=claimed_by,json=claimedBy,proto3" json:"claimed_by,omitempty" db:"claimed_by,omitempty" url:"claimed_by,omitempty"`
	RequestCarrier       map[string]string `protobuf:"bytes,4,rep,name=request_carrier,json=requestCarrier,proto3" json:"request_carrier,omitempty" db:"request_carrier,omitempty" url:"request_carrier,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ClaimPromoCodeRequest) Reset()         { *m = ClaimPromoCodeRequest{} }
func (m *ClaimPromoCodeRequest) String() string { return proto.CompactTextString(m) }
func (*ClaimPromoCodeRequest) ProtoMessage()    {}

type ClaimPromoCodeReply struct {
	Claim                *PromoCodeClaim `protobuf:"bytes,1,opt,name=claim,proto3" json:"claim,omitempty" db:"claim,omitempty" url:"claim,omitempty"`
	Error                *Error          `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty" db:"error,omitempty" url:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ClaimPromoCodeReply) Reset()         { *m = ClaimPromoCodeReply{} }
func (m *ClaimPromoCodeReply) String() string { return proto.CompactTextString(m) }
func (*ClaimPromoCodeReply) ProtoMessage()    {}

func (m *ClaimPromoCodeReply) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type GetPromoCodeClaimsRequest struct {
	OrgId                int32             `protobuf:"varint,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty" db:"org_id,omitempty" url:"org_id,omitempty"`
	RequestCarrier       map[string]string `protobuf:"bytes,2,rep,name=request_carrier,json=requestCarrier,proto3" json:"request_carrier,omitempty" db:"request_carrier,omitempty" url:"request_carrier,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ExcludeExpired       bool              `protobuf:"varint,3,opt,name=exclude_expired,json=excludeExpired,proto3" json:"exclude_expired,omitempty" db:"exclude_expired,omitempty" url:"exclude_expired,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GetPromoCodeClaimsRequest) Reset()         { *m = GetPromoCodeClaimsRequest{} }
func (m *GetPromoCodeClaimsRequest) String() string { return proto.CompactTextString(m) }
func (*GetPromoCodeClaimsRequest) ProtoMessage()    {}

type GetPromoCodeClaimsReply struct {
	Claims               []*PromoCodeClaim `protobuf:"bytes,1,rep,name=claims,proto3" json:"claims,omitempty" db:"claims,omitempty" url:"claims,omitempty"`
	Error                *Error            `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty" db:"error,omitempty" url:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GetPromoCodeClaimsReply) Reset()         { *m = GetPromoCodeClaimsReply{} }
func (m *GetPromoCodeClaimsReply) String() string { return proto.CompactTextString(m) }
func (*GetPromoCodeClaimsReply) ProtoMessage()    {}

func (m *GetPromoCodeClaimsReply) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *TaxId) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TypeEnum)
	if l > 0 {
		n += 1 + l + sovBilling(uint64(l))
	}
	l = len(m.TaxId)
	if l > 0 {
		n += 1 + l + sovBilling(uint64(l))
	}
	l = len(m.StripeObjectId)
	if l > 0 {
		n += 1 + l + sovBilling(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}
func (m *PriceTable) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.PriceTable) > 0 {
		for k, v := range m.PriceTable {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovBilling(uint64(l))
			}
			mapEntrySize := 1 + len(k) + sovBilling(uint64(len(k))) + l
			n += mapEntrySize + 1 + sovBilling(uint64(mapEntrySize))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}
func (m *UnitPrices) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Prices) > 0 {
		for k, v := range m.Prices {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovBilling(uint64(len(k))) + 1 + 8
			n += mapEntrySize + 1 + sovBilling(uint64(mapEntrySize))
		}
	}
	l = len(m.Unit)
	if l > 0 {
		n += 1 + l + sovBilling(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}
func (m *Card) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Cardholder)
	if l > 0 {
		n += 1 + l + sovBilling(uint64(l))
	}
	l = len(m.Brand)
	if l > 0 {
		n += 1 + l + sovBilling(uint64(l))
	}
	l = len(m.Last4)
	if l > 0 {
		n += 1 + l + sovBilling(uint64(l))
	}
	l = len(m.ExpMonth)
	if l > 0 {
		n += 1 + l + sovBilling(uint64(l))
	}
	l = len(m.ExpYear)
	if l > 0 {
		n += 1 + l + sovBilling(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}
func (m *TaxAddress) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Line1)
	if l > 0 {
		n += 1 + l + sovBilling(uint64(l))
	}
	l = len(m.Line2)
	if l > 0 {
		n += 1 + l + sovBilling(uint64(l))
	}
	l = len(m.City)
	if l > 0 {
		n += 1 + l + sovBilling(uint64(l))
	}
	l = len(m.State)
	if l > 0 {
		n += 1 + l + sovBilling(uint64(l))
	}
	l = len(m.Country)
	if l > 0 {
		n += 1 + l + sovBilling(uint64(l))
	}
	l = len(m.PostalCode)
	if l > 0 {
		n += 1 + l + sovBilling(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}
func (m *TaxAddress) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBilling
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TaxAddress: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TaxAddress: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Line1", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBilling
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBilling
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBilling
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Line1 = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Line2", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBilling
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBilling
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBilling
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Line2 = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field City", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBilling
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBilling
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBilling
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.City = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBilling
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBilling
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBilling
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.State = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Country", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBilling
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBilling
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBilling
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Country = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PostalCode", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBilling
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBilling
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBilling
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PostalCode = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBilling(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBilling
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthBilling
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TaxId) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBilling
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TaxId: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TaxId: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TypeEnum", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBilling
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBilling
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBilling
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TypeEnum = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TaxId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBilling
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBilling
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBilling
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TaxId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StripeObjectId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBilling
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBilling
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBilling
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StripeObjectId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBilling(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBilling
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthBilling
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PromoCodeClaim) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBilling
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PromoCodeClaim: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PromoCodeClaim: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrgId", wireType)
			}
			m.OrgId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBilling
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OrgId |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBilling
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBilling
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBilling
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Code = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			m.Amount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBilling
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Amount |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Balance", wireType)
			}
			m.Balance = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBilling
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Balance |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimDate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBilling
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBilling
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBilling
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ClaimDate == nil {
				m.ClaimDate = &types.Timestamp{}
			}
			if err := m.ClaimDate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreditExpirationDate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBilling
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBilling
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBilling
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CreditExpirationDate == nil {
				m.CreditExpirationDate = &types.Timestamp{}
			}
			if err := m.CreditExpirationDate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimedBy", wireType)
			}
			m.ClaimedBy = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBilling
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ClaimedBy |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBilling(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBilling
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthBilling
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Card) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBilling
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Card: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Card: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cardholder", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBilling
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBilling
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBilling
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Cardholder = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Brand", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBilling
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBilling
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBilling
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Brand = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Last4", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBilling
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBilling
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBilling
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Last4 = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpMonth", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBilling
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBilling
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBilling
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ExpMonth = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpYear", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBilling
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBilling
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBilling
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ExpYear = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBilling(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBilling
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthBilling
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PriceTable) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBilling
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PriceTable: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PriceTable: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PriceTable", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBilling
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBilling
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBilling
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PriceTable == nil {
				m.PriceTable = make(map[string]*UnitPrices)
			}
			var mapkey string
			var mapvalue *UnitPrices
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowBilling
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowBilling
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthBilling
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthBilling
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowBilling
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthBilling
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthBilling
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &UnitPrices{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipBilling(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthBilling
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.PriceTable[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBilling(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBilling
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthBilling
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetPriceTableReply) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBilling
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetPriceTableReply: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetPriceTableReply: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PriceTable", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBilling
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBilling
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBilling
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PriceTable == nil {
				m.PriceTable = &PriceTable{}
			}
			if err := m.PriceTable.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBilling
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBilling
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBilling
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Error == nil {
				m.Error = &Error{}
			}
			if err := m.Error.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBilling(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBilling
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthBilling
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetPaymentInfoReply) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBilling
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetPaymentInfoReply: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetPaymentInfoReply: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Card", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBilling
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBilling
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBilling
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Card == nil {
				m.Card = &Card{}
			}
			if err := m.Card.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBilling
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBilling
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBilling
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Error == nil {
				m.Error = &Error{}
			}
			if err := m.Error.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Organization", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBilling
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBilling
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBilling
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Organization == nil {
				m.Organization = &Organization{}
			}
			if err := m.Organization.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TaxAddress", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBilling
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBilling
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBilling
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TaxAddress == nil {
				m.TaxAddress = &TaxAddress{}
			}
			if err := m.TaxAddress.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TaxIds", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBilling
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBilling
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBilling
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TaxIds = append(m.TaxIds, &TaxId{})
			if err := m.TaxIds[len(m.TaxIds)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBilling(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBilling
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthBilling
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PromoCodeClaim) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.OrgId != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintBilling(dAtA, i, uint64(m.OrgId))
	}
	if len(m.Code) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintBilling(dAtA, i, uint64(len(m.Code)))
		i += copy(dAtA[i:], m.Code)
	}
	if m.Amount != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintBilling(dAtA, i, uint64(m.Amount))
	}
	if m.Balance != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintBilling(dAtA, i, uint64(m.Balance))
	}
	if m.ClaimDate != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintBilling(dAtA, i, uint64(m.ClaimDate.Size()))
		n68, err := m.ClaimDate.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n68
	}
	if m.CreditExpirationDate != nil {
		dAtA[i] = 0x32
		i++
		i = encodeVarintBilling(dAtA, i, uint64(m.CreditExpirationDate.Size()))
		n69, err := m.CreditExpirationDate.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n69
	}
	if m.ClaimedBy != 0 {
		dAtA[i] = 0x38
		i++
		i = encodeVarintBilling(dAtA, i, uint64(m.ClaimedBy))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}
func (m *Card) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Cardholder) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintBilling(dAtA, i, uint64(len(m.Cardholder)))
		i += copy(dAtA[i:], m.Cardholder)
	}
	if len(m.Brand) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintBilling(dAtA, i, uint64(len(m.Brand)))
		i += copy(dAtA[i:], m.Brand)
	}
	if len(m.Last4) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintBilling(dAtA, i, uint64(len(m.Last4)))
		i += copy(dAtA[i:], m.Last4)
	}
	if len(m.ExpMonth) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintBilling(dAtA, i, uint64(len(m.ExpMonth)))
		i += copy(dAtA[i:], m.ExpMonth)
	}
	if len(m.ExpYear) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintBilling(dAtA, i, uint64(len(m.ExpYear)))
		i += copy(dAtA[i:], m.ExpYear)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}
func (m *PriceTable) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.PriceTable) > 0 {
		for k, _ := range m.PriceTable {
			dAtA[i] = 0xa
			i++
			v := m.PriceTable[k]
			msgSize := 0
			if v != nil {
				msgSize = v.Size()
				msgSize += 1 + sovBilling(uint64(msgSize))
			}
			mapSize := 1 + len(k) + sovBilling(uint64(len(k))) + msgSize
			i = encodeVarintBilling(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintBilling(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			if v != nil {
				dAtA[i] = 0x12
				i++
				i = encodeVarintBilling(dAtA, i, uint64(v.Size()))
				n89, err := v.MarshalTo(dAtA[i:])
				if err != nil {
					return 0, err
				}
				i += n89
			}
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}
func (m *UnitPrices) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBilling
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: UnitPrices: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UnitPrices: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Prices", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBilling
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBilling
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBilling
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Prices == nil {
				m.Prices = make(map[string]float64)
			}
			var mapkey string
			var mapvalue float64
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowBilling
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowBilling
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthBilling
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthBilling
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapvaluetemp uint64
					if (iNdEx + 8) > l {
						return io.ErrUnexpectedEOF
					}
					mapvaluetemp = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
					iNdEx += 8
					mapvalue = math.Float64frombits(mapvaluetemp)
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipBilling(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthBilling
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Prices[mapkey] = mapvalue
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Unit", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBilling
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBilling
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBilling
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Unit = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBilling(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBilling
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthBilling
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TaxAddress) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Line1) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintBilling(dAtA, i, uint64(len(m.Line1)))
		i += copy(dAtA[i:], m.Line1)
	}
	if len(m.Line2) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintBilling(dAtA, i, uint64(len(m.Line2)))
		i += copy(dAtA[i:], m.Line2)
	}
	if len(m.City) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintBilling(dAtA, i, uint64(len(m.City)))
		i += copy(dAtA[i:], m.City)
	}
	if len(m.State) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintBilling(dAtA, i, uint64(len(m.State)))
		i += copy(dAtA[i:], m.State)
	}
	if len(m.Country) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintBilling(dAtA, i, uint64(len(m.Country)))
		i += copy(dAtA[i:], m.Country)
	}
	if len(m.PostalCode) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintBilling(dAtA, i, uint64(len(m.PostalCode)))
		i += copy(dAtA[i:], m.PostalCode)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}
func (m *UnitPrices) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Prices) > 0 {
		for k, _ := range m.Prices {
			dAtA[i] = 0xa
			i++
			v := m.Prices[k]
			mapSize := 1 + len(k) + sovBilling(uint64(len(k))) + 1 + 8
			i = encodeVarintBilling(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintBilling(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x11
			i++
			encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(v))))
			i += 8
		}
	}
	if len(m.Unit) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintBilling(dAtA, i, uint64(len(m.Unit)))
		i += copy(dAtA[i:], m.Unit)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}
func (m *GetPriceTableReply) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.PriceTable != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintBilling(dAtA, i, uint64(m.PriceTable.Size()))
		n49, err := m.PriceTable.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n49
	}
	if m.Error != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintBilling(dAtA, i, uint64(m.Error.Size()))
		n50, err := m.Error.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n50
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}
func (m *TaxId) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.TypeEnum) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintBilling(dAtA, i, uint64(len(m.TypeEnum)))
		i += copy(dAtA[i:], m.TypeEnum)
	}
	if len(m.TaxId) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintBilling(dAtA, i, uint64(len(m.TaxId)))
		i += copy(dAtA[i:], m.TaxId)
	}
	if len(m.StripeObjectId) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintBilling(dAtA, i, uint64(len(m.StripeObjectId)))
		i += copy(dAtA[i:], m.StripeObjectId)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}
func (m *GetPaymentInfoReply) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Card != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintBilling(dAtA, i, uint64(m.Card.Size()))
		n108, err := m.Card.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n108
	}
	if m.Error != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintBilling(dAtA, i, uint64(m.Error.Size()))
		n109, err := m.Error.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n109
	}
	if m.Organization != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintBilling(dAtA, i, uint64(m.Organization.Size()))
		n110, err := m.Organization.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n110
	}
	if m.TaxAddress != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintBilling(dAtA, i, uint64(m.TaxAddress.Size()))
		n111, err := m.TaxAddress.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n111
	}
	if len(m.TaxIds) > 0 {
		for _, msg := range m.TaxIds {
			dAtA[i] = 0x2a
			i++
			i = encodeVarintBilling(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintBilling(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}

func sovBilling(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}

func skipBilling(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBilling
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowBilling
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowBilling
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthBilling
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthBilling
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowBilling
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipBilling(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthBilling
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthBilling = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBilling   = fmt.Errorf("proto: integer overflow")
)
