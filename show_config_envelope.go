package main

import (
	"log"

	"github.com/hyperledger/fabric/protoutil"

	"github.com/hyperledger/fabric/common/configtx"
	"github.com/hyperledger/fabric/protos/common"
)

func dealConfigEnvelope(payload *common.Payload) *ConfigEnvelope {

	configEnvelope, err := configtx.UnmarshalConfigEnvelope(payload.Data)
	if err != nil {
		log.Fatalf("error unmarshalling config, %v", err)
	}

	return convertConfigEnvelope(configEnvelope)
}

func convertConfigEnvelope(envelope *common.ConfigEnvelope) *ConfigEnvelope {
	ce := &ConfigEnvelope{
		Config:     convertConfig(envelope.Config),
		LastUpdate: convertLastUpdate(envelope.LastUpdate),
	}

	return ce
}

func convertLastUpdate(envelope *common.Envelope) *ConfigUpdateEnvelope {
	if envelope == nil {
		return nil
	}

	configUpdateEnv, err := protoutil.EnvelopeToConfigUpdate(envelope)
	if err != nil {
		log.Fatalf("err in EnvelopeToConfigUpdate: %v", err)
	}

	return &ConfigUpdateEnvelope{
		ConfigUpdate: convertConfigUpdate(configUpdateEnv.ConfigUpdate),
		Signatures:   convertConfigSignature(configUpdateEnv.Signatures),
	}
}

func convertConfigUpdate(bytes []byte) *ConfigUpdate {

	oldconfigUpdate, err := configtx.UnmarshalConfigUpdate(bytes)
	if err != nil {
		log.Fatalf("error in UnmarshalConfigUpdate, %v ", err)
	}

	cu := &ConfigUpdate{
		ChannelId:    oldconfigUpdate.ChannelId,
		ReadSet:      convertChannelGroup(oldconfigUpdate.ReadSet),
		WriteSet:     convertChannelGroup(oldconfigUpdate.WriteSet),
		IsolatedData: oldconfigUpdate.IsolatedData,
	}

	return cu
}

func convertConfigSignature(signatures []*common.ConfigSignature) []*ConfigSignature {
	var ss []*ConfigSignature
	for _, v := range signatures {
		s := &ConfigSignature{
			SignatureHeader: getSignatureHeader(v.SignatureHeader),
			Signature:       v.Signature,
		}
		ss = append(ss, s)
	}
	return ss
}

func convertConfig(config *common.Config) *Config {

	return &Config{
		Sequence:     config.Sequence,
		ChannelGroup: convertChannelGroup(config.ChannelGroup),
	}
}

func convertChannelGroup(group *common.ConfigGroup) *ConfigGroup {
	cg := &ConfigGroup{
		Version:   group.Version,
		Groups:    make(map[string]*ConfigGroup),
		Values:    make(map[string]*ConfigValue),
		Policies:  make(map[string]*ConfigPolicy),
		ModPolicy: group.ModPolicy,
	}

	for k, v := range group.Groups {
		cg.Groups[k] = convertChannelGroup(v)
	}

	for k, v := range group.Values {
		cg.Values[k] = convertConfigValue(k, v)
	}

	for k, v := range group.Policies {
		cg.Policies[k] = convertConfigConfigPolicy(v)
	}

	return cg
}

func convertConfigConfigPolicy(policy *common.ConfigPolicy) *ConfigPolicy {

	cp := &ConfigPolicy{
		Version:   policy.Version,
		Policy:    convertPolicy(policy.Policy),
		ModPolicy: policy.ModPolicy,
	}

	return cp

}

func convertPolicy(policy *common.Policy) *Policy {
	p := &Policy{}

	if policy != nil {
		p.Type = policy.Type
		p.Value = policy.Value
	}

	return p
}

func convertConfigValue(key string, value *common.ConfigValue) *ConfigValue {

	cv := &ConfigValue{
		Version:   value.Version,
		Value:     value.Value,
		ModPolicy: value.ModPolicy,
	}

	/*if key == "MSP" {  // TODO: how to decode the value
		cv.Value_role = showCreator(value.Value)
	}*/
	return cv
}
