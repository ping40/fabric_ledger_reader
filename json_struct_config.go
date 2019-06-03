package main

type ConfigEnvelope struct {
	Config     *Config               `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	LastUpdate *ConfigUpdateEnvelope `protobuf:"bytes,2,opt,name=last_update,json=lastUpdate,proto3" json:"last_update,omitempty"`
}

// Config represents the config for a particular channel
type Config struct {
	Sequence     uint64       `protobuf:"varint,1,opt,name=sequence,proto3" json:"sequence,omitempty"`
	ChannelGroup *ConfigGroup `protobuf:"bytes,2,opt,name=channel_group,json=channelGroup,proto3" json:"channel_group,omitempty"`
}

// ConfigGroup is the hierarchical data structure for holding config
type ConfigGroup struct {
	Version   uint64                   `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Groups    map[string]*ConfigGroup  `protobuf:"bytes,2,rep,name=groups,proto3" json:"groups,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Values    map[string]*ConfigValue  `protobuf:"bytes,3,rep,name=values,proto3" json:"values,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Policies  map[string]*ConfigPolicy `protobuf:"bytes,4,rep,name=policies,proto3" json:"policies,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ModPolicy string                   `protobuf:"bytes,5,opt,name=mod_policy,json=modPolicy,proto3" json:"mod_policy,omitempty"`
}

// ConfigValue represents an individual piece of config data
type ConfigValue struct {
	Version    uint64              `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Value      []byte              `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	ModPolicy  string              `protobuf:"bytes,3,opt,name=mod_policy,json=modPolicy,proto3" json:"mod_policy,omitempty"`
	Value_role *SerializedIdentity `json:"value_role,omitempty"`
}

type ConfigPolicy struct {
	Version   uint64  `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Policy    *Policy `protobuf:"bytes,2,opt,name=policy,proto3" json:"policy,omitempty"`
	ModPolicy string  `protobuf:"bytes,3,opt,name=mod_policy,json=modPolicy,proto3" json:"mod_policy,omitempty"`
}

// Policy expresses a policy which the orderer can evaluate, because there has been some desire expressed to support
// multiple policy engines, this is typed as a oneof for now
type Policy struct {
	Type  int32  `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

type ConfigUpdateEnvelope struct {
	ConfigUpdate *ConfigUpdate      `protobuf:"bytes,1,opt,name=config_update,json=configUpdate,proto3" json:"config_update,omitempty"`
	Signatures   []*ConfigSignature `protobuf:"bytes,2,rep,name=signatures,proto3" json:"signatures,omitempty"`
}

type ConfigSignature struct {
	SignatureHeader *SignatureHeader `protobuf:"bytes,1,opt,name=signature_header,json=signatureHeader,proto3" json:"signature_header,omitempty"`
	Signature       []byte           `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

type ConfigUpdate struct {
	ChannelId    string            `protobuf:"bytes,1,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	ReadSet      *ConfigGroup      `protobuf:"bytes,2,opt,name=read_set,json=readSet,proto3" json:"read_set,omitempty"`
	WriteSet     *ConfigGroup      `protobuf:"bytes,3,opt,name=write_set,json=writeSet,proto3" json:"write_set,omitempty"`
	IsolatedData map[string][]byte `protobuf:"bytes,5,rep,name=isolated_data,json=isolatedData,proto3" json:"isolated_data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}
