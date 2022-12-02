package ccloud

import (
	"time"

	_ "github.com/confluentinc/proto-go-setter"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	_ "github.com/travisjeffery/proto-go-sql"
)

type AccountMetadata struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" db:"id,omitempty" url:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountMetadata) Reset()         { *m = AccountMetadata{} }
func (m *AccountMetadata) String() string { return proto.CompactTextString(m) }
func (*AccountMetadata) ProtoMessage()    {}

type Region struct {
	Id    string              `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" db:"id,omitempty" url:"id,omitempty"`
	Cloud string              `protobuf:"bytes,2,opt,name=cloud,proto3" json:"cloud,omitempty" db:"cloud,omitempty" url:"cloud,omitempty"`
	Zones []*AvailabilityZone `protobuf:"bytes,3,rep,name=zones,proto3" json:"zones,omitempty" db:"zones,omitempty" url:"zones,omitempty"`
	Name  string              `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty" db:"name,omitempty" url:"name,omitempty"`
	// This field will be calcuated at runtime. Default to false
	// e.g. when there is available physical cluster for logical cluster provision, it will be set to true
	IsSchedulable bool `protobuf:"varint,5,opt,name=is_schedulable,json=isSchedulable,proto3" json:"is_schedulable,omitempty" db:"is_schedulable,omitempty" url:"is_schedulable,omitempty"`
	// This field will be calcuated at runtime. Default to false
	// e.g. when there is available physical cluster for logical cluster provision, and its Durability==High, it will be set to true
	IsMultizoneEnabled bool           `protobuf:"varint,6,opt,name=is_multizone_enabled,json=isMultizoneEnabled,proto3" json:"is_multizone_enabled,omitempty" db:"is_multizone_enabled,omitempty" url:"is_multizone_enabled,omitempty"`
	Config             *Region_Config `protobuf:"bytes,7,opt,name=config,proto3" json:"config,omitempty" db:"config,omitempty" url:"config,omitempty"`
	// This is a vanity field. This means that it is not persisted to the DB
	// but computed at runtime.
	Schedulability       *Schedulability `protobuf:"bytes,8,opt,name=schedulability,proto3" json:"schedulability,omitempty" db:"schedulability,omitempty" url:"schedulability,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Region) Reset()         { *m = Region{} }
func (m *Region) String() string { return proto.CompactTextString(m) }
func (*Region) ProtoMessage()    {}

func (m *Region) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Region) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Region) GetIsSchedulable() bool {
	if m != nil {
		return m.IsSchedulable
	}
	return false
}

type AvailabilityZone struct {
	// us-west-2
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" db:"name,omitempty" url:"name,omitempty"`
	// AWS:   zone id
	// GCP:   zone name
	// AZURE: zone name
	ZoneId string `protobuf:"bytes,3,opt,name=zone_id,json=zoneId,proto3" json:"zone_id,omitempty" db:"zone_id,omitempty" url:"zone_id,omitempty"`
	// The internal mothership zone id
	Id                   string           `protobuf:"bytes,4,opt,name=id,proto3" json:"id,omitempty" db:"id,omitempty" url:"id,omitempty"`
	RegionId             string           `protobuf:"bytes,5,opt,name=region_id,json=regionId,proto3" json:"region_id,omitempty" db:"region_id,omitempty" url:"region_id,omitempty"`
	SniEnabled           *types.BoolValue `protobuf:"bytes,6,opt,name=sni_enabled,json=sniEnabled,proto3" json:"sni_enabled,omitempty" db:"sni_enabled,omitempty" url:"sni_enabled,omitempty"`
	Schedulable          *types.BoolValue `protobuf:"bytes,7,opt,name=schedulable,proto3" json:"schedulable,omitempty" db:"schedulable,omitempty" url:"schedulable,omitempty"`
	Created              *time.Time       `protobuf:"bytes,8,opt,name=created,proto3,stdtime" json:"created,omitempty" db:"created,omitempty" url:"created,omitempty"`
	Modified             *time.Time       `protobuf:"bytes,9,opt,name=modified,proto3,stdtime" json:"modified,omitempty" db:"modified,omitempty" url:"modified,omitempty"`
	Deactivated          *time.Time       `protobuf:"bytes,10,opt,name=deactivated,proto3,stdtime" json:"deactivated,omitempty" db:"deactivated,omitempty" url:"deactivated,omitempty"`
	Realm                string           `protobuf:"bytes,12,opt,name=realm,proto3" json:"realm,omitempty" db:"realm,omitempty" url:"realm,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *AvailabilityZone) Reset()         { *m = AvailabilityZone{} }
func (m *AvailabilityZone) String() string { return proto.CompactTextString(m) }
func (*AvailabilityZone) ProtoMessage()    {}

type Region_Config struct {
	Docker               *Region_Docker `protobuf:"bytes,1,opt,name=docker,proto3" json:"docker,omitempty" db:"docker,omitempty" url:"docker,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Region_Config) Reset()         { *m = Region_Config{} }
func (m *Region_Config) String() string { return proto.CompactTextString(m) }
func (*Region_Config) ProtoMessage()    {}

type Region_Docker struct {
	Repo                 string   `protobuf:"bytes,1,opt,name=repo,proto3" json:"repo,omitempty" db:"repo,omitempty" url:"repo,omitempty"`
	ImagePrefix          string   `protobuf:"bytes,2,opt,name=image_prefix,json=imagePrefix,proto3" json:"image_prefix,omitempty" db:"image_prefix,omitempty" url:"image_prefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Region_Docker) Reset()         { *m = Region_Docker{} }
func (m *Region_Docker) String() string { return proto.CompactTextString(m) }
func (*Region_Docker) ProtoMessage()    {}

type Schedulability struct {
	SharedNetwork        *Schedulability_Tenancy `protobuf:"bytes,1,opt,name=shared_network,json=sharedNetwork,proto3" json:"shared_network,omitempty" db:"shared_network,omitempty" url:"shared_network,omitempty"`
	DedicatedNetwork     *Schedulability_Tenancy `protobuf:"bytes,2,opt,name=dedicated_network,json=dedicatedNetwork,proto3" json:"dedicated_network,omitempty" db:"dedicated_network,omitempty" url:"dedicated_network,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *Schedulability) Reset()         { *m = Schedulability{} }
func (m *Schedulability) String() string { return proto.CompactTextString(m) }
func (*Schedulability) ProtoMessage()    {}

type Schedulability_Tenancy struct {
	SharedCluster        *Schedulability_Tenancy_Durability `protobuf:"bytes,1,opt,name=shared_cluster,json=sharedCluster,proto3" json:"shared_cluster,omitempty" db:"shared_cluster,omitempty" url:"shared_cluster,omitempty"`
	DedicatedCluster     *Schedulability_Tenancy_Durability `protobuf:"bytes,2,opt,name=dedicated_cluster,json=dedicatedCluster,proto3" json:"dedicated_cluster,omitempty" db:"dedicated_cluster,omitempty" url:"dedicated_cluster,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *Schedulability_Tenancy) Reset()         { *m = Schedulability_Tenancy{} }
func (m *Schedulability_Tenancy) String() string { return proto.CompactTextString(m) }
func (*Schedulability_Tenancy) ProtoMessage()    {}

type Schedulability_Tenancy_Durability struct {
	Low                  []NetworkType `protobuf:"varint,1,rep,packed,name=low,proto3,enum=kafka.scheduler.v1.NetworkType" json:"LOW" db:"low,omitempty" url:"low,omitempty"`
	High                 []NetworkType `protobuf:"varint,2,rep,packed,name=high,proto3,enum=kafka.scheduler.v1.NetworkType" json:"HIGH" db:"high,omitempty" url:"high,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Schedulability_Tenancy_Durability) Reset()         { *m = Schedulability_Tenancy_Durability{} }
func (m *Schedulability_Tenancy_Durability) String() string { return proto.CompactTextString(m) }
func (*Schedulability_Tenancy_Durability) ProtoMessage()    {}

type NetworkType int32

const (
	NetworkType_UNKNOWN_TYPE    NetworkType = 0
	NetworkType_PUBLIC          NetworkType = 1
	NetworkType_PRIVATE_LINK    NetworkType = 2
	NetworkType_TRANSIT_GATEWAY NetworkType = 3
	NetworkType_VPC_PEERING     NetworkType = 4
	// INTERNAL requires that source traffic origniates from the NetworkRegion's requested CIDR.
	NetworkType_INTERNAL NetworkType = 5
)

var NetworkType_name = map[int32]string{
	0: "UNKNOWN_TYPE",
	1: "PUBLIC",
	2: "PRIVATE_LINK",
	3: "TRANSIT_GATEWAY",
	4: "VPC_PEERING",
	5: "INTERNAL",
}

var NetworkType_value = map[string]int32{
	"UNKNOWN_TYPE":    0,
	"PUBLIC":          1,
	"PRIVATE_LINK":    2,
	"TRANSIT_GATEWAY": 3,
	"VPC_PEERING":     4,
	"INTERNAL":        5,
}

func (x NetworkType) String() string {
	return proto.EnumName(NetworkType_name, int32(x))
}

type CloudMetadata struct {
	Id                   string             `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" db:"id,omitempty" url:"id,omitempty"`
	Regions              []*Region          `protobuf:"bytes,2,rep,name=regions,proto3" json:"regions,omitempty" db:"regions,omitempty" url:"regions,omitempty"`
	Name                 string             `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty" db:"name,omitempty" url:"name,omitempty"`
	Accounts             []*AccountMetadata `protobuf:"bytes,4,rep,name=accounts,proto3" json:"accounts,omitempty" db:"accounts,omitempty" url:"accounts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *CloudMetadata) Reset()         { *m = CloudMetadata{} }
func (m *CloudMetadata) String() string { return proto.CompactTextString(m) }
func (*CloudMetadata) ProtoMessage()    {}

func (m *CloudMetadata) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CloudMetadata) GetRegions() []*Region {
	if m != nil {
		return m.Regions
	}
	return nil
}

func (m *CloudMetadata) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CloudMetadata) GetAccounts() []*AccountMetadata {
	if m != nil {
		return m.Accounts
	}
	return nil
}

type GetEnvironmentMetadataReply struct {
	Error                   *Error                    `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty" db:"error,omitempty" url:"error,omitempty"`
	Clouds                  []*CloudMetadata          `protobuf:"bytes,2,rep,name=clouds,proto3" json:"clouds,omitempty" db:"clouds,omitempty" url:"clouds,omitempty"`
	Status                  *EnvironmentStatus        `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty" db:"status,omitempty" url:"status,omitempty"`
	SchemaRegistryLocations []*SchemaRegistryLocation `protobuf:"bytes,4,rep,name=schema_registry_locations,json=schemaRegistryLocations,proto3" json:"schema_registry_locations,omitempty" db:"schema_registry_locations,omitempty" url:"schema_registry_locations,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}                  `json:"-"`
	XXX_unrecognized        []byte                    `json:"-"`
	XXX_sizecache           int32                     `json:"-"`
}

func (m *GetEnvironmentMetadataReply) Reset()         { *m = GetEnvironmentMetadataReply{} }
func (m *GetEnvironmentMetadataReply) String() string { return proto.CompactTextString(m) }
func (*GetEnvironmentMetadataReply) ProtoMessage()    {}

func (m *GetEnvironmentMetadataReply) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *GetEnvironmentMetadataReply) GetClouds() []*CloudMetadata {
	if m != nil {
		return m.Clouds
	}
	return nil
}

func (m *GetEnvironmentMetadataReply) GetStatus() *EnvironmentStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *GetEnvironmentMetadataReply) GetSchemaRegistryLocations() []*SchemaRegistryLocation {
	if m != nil {
		return m.SchemaRegistryLocations
	}
	return nil
}

type EnvironmentStatus struct {
	Info                 string   `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty" db:"info,omitempty" url:"info,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty" db:"error,omitempty" url:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EnvironmentStatus) Reset()         { *m = EnvironmentStatus{} }
func (m *EnvironmentStatus) String() string { return proto.CompactTextString(m) }
func (*EnvironmentStatus) ProtoMessage()    {}

type SchemaRegistryLocation struct {
	Id                    GlobalSchemaRegistryLocation `protobuf:"varint,1,opt,name=id,proto3,enum=kafka.scheduler.v1.GlobalSchemaRegistryLocation" json:"id,omitempty" db:"id,omitempty" url:"id,omitempty"` // Deprecated: Do not use.
	Name                  string                       `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" db:"name,omitempty" url:"name,omitempty"`                                               // Deprecated: Do not use.
	ClusterId             string                       `protobuf:"bytes,3,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty" db:"cluster_id,omitempty" url:"cluster_id,omitempty"`
	ServiceProvider       string                       `protobuf:"bytes,4,opt,name=service_provider,json=serviceProvider,proto3" json:"service_provider,omitempty" db:"service_provider,omitempty" url:"service_provider,omitempty"`
	Package               []string                     `protobuf:"bytes,5,rep,name=package,proto3" json:"package,omitempty" db:"package,omitempty" url:"package,omitempty"`
	ServiceProviderRegion string                       `protobuf:"bytes,6,opt,name=service_provider_region,json=serviceProviderRegion,proto3" json:"service_provider_region,omitempty" db:"service_provider_region,omitempty" url:"service_provider_region,omitempty"`
	RegionDisplayName     string                       `protobuf:"bytes,7,opt,name=region_display_name,json=regionDisplayName,proto3" json:"region_display_name,omitempty" db:"region_display_name,omitempty" url:"region_display_name,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}                     `json:"-"`
	XXX_unrecognized      []byte                       `json:"-"`
	XXX_sizecache         int32                        `json:"-"`
}

func (m *SchemaRegistryLocation) Reset()         { *m = SchemaRegistryLocation{} }
func (m *SchemaRegistryLocation) String() string { return proto.CompactTextString(m) }
func (*SchemaRegistryLocation) ProtoMessage()    {}

type GlobalSchemaRegistryLocation int32

const (
	GlobalSchemaRegistryLocation_NONE GlobalSchemaRegistryLocation = 0
	GlobalSchemaRegistryLocation_US   GlobalSchemaRegistryLocation = 1
	GlobalSchemaRegistryLocation_EU   GlobalSchemaRegistryLocation = 2
	GlobalSchemaRegistryLocation_APAC GlobalSchemaRegistryLocation = 3
)

var GlobalSchemaRegistryLocation_name = map[int32]string{
	0: "NONE",
	1: "US",
	2: "EU",
	3: "APAC",
}

var GlobalSchemaRegistryLocation_value = map[string]int32{
	"NONE": 0,
	"US":   1,
	"EU":   2,
	"APAC": 3,
}

func (x GlobalSchemaRegistryLocation) String() string {
	return proto.EnumName(GlobalSchemaRegistryLocation_name, int32(x))
}
