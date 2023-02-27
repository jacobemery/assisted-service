// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FeatureSupportLevel feature support level
//
// swagger:model feature-support-level
type FeatureSupportLevel struct {

	// features
	Features []*FeatureSupportLevelFeaturesItems0 `json:"features"`

	// Version of the OpenShift cluster.
	OpenshiftVersion string `json:"openshift_version,omitempty"`
}

// Validate validates this feature support level
func (m *FeatureSupportLevel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFeatures(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FeatureSupportLevel) validateFeatures(formats strfmt.Registry) error {
	if swag.IsZero(m.Features) { // not required
		return nil
	}

	for i := 0; i < len(m.Features); i++ {
		if swag.IsZero(m.Features[i]) { // not required
			continue
		}

		if m.Features[i] != nil {
			if err := m.Features[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("features" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("features" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this feature support level based on the context it is used
func (m *FeatureSupportLevel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFeatures(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FeatureSupportLevel) contextValidateFeatures(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Features); i++ {

		if m.Features[i] != nil {
			if err := m.Features[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("features" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("features" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *FeatureSupportLevel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FeatureSupportLevel) UnmarshalBinary(b []byte) error {
	var res FeatureSupportLevel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FeatureSupportLevelFeaturesItems0 feature support level features items0
//
// swagger:model FeatureSupportLevelFeaturesItems0
type FeatureSupportLevelFeaturesItems0 struct {

	// The ID of the feature
	// Required: true
	// Enum: [ADDITIONAL_NTP_SOURCE REQUESTED_HOSTNAME PROXY SNO DAY2_HOSTS VIP_AUTO_ALLOC DISK_SELECTION OVN_NETWORK_TYPE SDN_NETWORK_TYPE PLATFORM_SELECTION SCHEDULABLE_MASTERS AUTO_ASSIGN_ROLE CUSTOM_MANIFEST DISK_ENCRYPTION CLUSTER_MANAGED_NETWORKING_WITH_VMS ARM64_ARCHITECTURE ARM64_ARCHITECTURE_WITH_CLUSTER_MANAGED_NETWORKING PPC64LE_ARCHITECTURE S390X_ARCHITECTURE SINGLE_NODE_EXPANSION LVM DUAL_STACK_NETWORKING MULTIARCH_RELEASE_IMAGE NUTANIX_INTEGRATION DUAL_STACK_VIPS USER_MANAGED_NETWORKING_WITH_MULTI_NODE]
	FeatureID *string `json:"feature_id"`

	// support level
	// Required: true
	// Enum: [supported unsupported tech-preview dev-preview]
	SupportLevel *string `json:"support_level"`
}

// Validate validates this feature support level features items0
func (m *FeatureSupportLevelFeaturesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFeatureID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSupportLevel(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var featureSupportLevelFeaturesItems0TypeFeatureIDPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ADDITIONAL_NTP_SOURCE","REQUESTED_HOSTNAME","PROXY","SNO","DAY2_HOSTS","VIP_AUTO_ALLOC","DISK_SELECTION","OVN_NETWORK_TYPE","SDN_NETWORK_TYPE","PLATFORM_SELECTION","SCHEDULABLE_MASTERS","AUTO_ASSIGN_ROLE","CUSTOM_MANIFEST","DISK_ENCRYPTION","CLUSTER_MANAGED_NETWORKING_WITH_VMS","ARM64_ARCHITECTURE","ARM64_ARCHITECTURE_WITH_CLUSTER_MANAGED_NETWORKING","PPC64LE_ARCHITECTURE","S390X_ARCHITECTURE","SINGLE_NODE_EXPANSION","LVM","DUAL_STACK_NETWORKING","MULTIARCH_RELEASE_IMAGE","NUTANIX_INTEGRATION","DUAL_STACK_VIPS","USER_MANAGED_NETWORKING_WITH_MULTI_NODE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		featureSupportLevelFeaturesItems0TypeFeatureIDPropEnum = append(featureSupportLevelFeaturesItems0TypeFeatureIDPropEnum, v)
	}
}

const (

	// FeatureSupportLevelFeaturesItems0FeatureIDADDITIONALNTPSOURCE captures enum value "ADDITIONAL_NTP_SOURCE"
	FeatureSupportLevelFeaturesItems0FeatureIDADDITIONALNTPSOURCE string = "ADDITIONAL_NTP_SOURCE"

	// FeatureSupportLevelFeaturesItems0FeatureIDREQUESTEDHOSTNAME captures enum value "REQUESTED_HOSTNAME"
	FeatureSupportLevelFeaturesItems0FeatureIDREQUESTEDHOSTNAME string = "REQUESTED_HOSTNAME"

	// FeatureSupportLevelFeaturesItems0FeatureIDPROXY captures enum value "PROXY"
	FeatureSupportLevelFeaturesItems0FeatureIDPROXY string = "PROXY"

	// FeatureSupportLevelFeaturesItems0FeatureIDSNO captures enum value "SNO"
	FeatureSupportLevelFeaturesItems0FeatureIDSNO string = "SNO"

	// FeatureSupportLevelFeaturesItems0FeatureIDDAY2HOSTS captures enum value "DAY2_HOSTS"
	FeatureSupportLevelFeaturesItems0FeatureIDDAY2HOSTS string = "DAY2_HOSTS"

	// FeatureSupportLevelFeaturesItems0FeatureIDVIPAUTOALLOC captures enum value "VIP_AUTO_ALLOC"
	FeatureSupportLevelFeaturesItems0FeatureIDVIPAUTOALLOC string = "VIP_AUTO_ALLOC"

	// FeatureSupportLevelFeaturesItems0FeatureIDDISKSELECTION captures enum value "DISK_SELECTION"
	FeatureSupportLevelFeaturesItems0FeatureIDDISKSELECTION string = "DISK_SELECTION"

	// FeatureSupportLevelFeaturesItems0FeatureIDOVNNETWORKTYPE captures enum value "OVN_NETWORK_TYPE"
	FeatureSupportLevelFeaturesItems0FeatureIDOVNNETWORKTYPE string = "OVN_NETWORK_TYPE"

	// FeatureSupportLevelFeaturesItems0FeatureIDSDNNETWORKTYPE captures enum value "SDN_NETWORK_TYPE"
	FeatureSupportLevelFeaturesItems0FeatureIDSDNNETWORKTYPE string = "SDN_NETWORK_TYPE"

	// FeatureSupportLevelFeaturesItems0FeatureIDPLATFORMSELECTION captures enum value "PLATFORM_SELECTION"
	FeatureSupportLevelFeaturesItems0FeatureIDPLATFORMSELECTION string = "PLATFORM_SELECTION"

	// FeatureSupportLevelFeaturesItems0FeatureIDSCHEDULABLEMASTERS captures enum value "SCHEDULABLE_MASTERS"
	FeatureSupportLevelFeaturesItems0FeatureIDSCHEDULABLEMASTERS string = "SCHEDULABLE_MASTERS"

	// FeatureSupportLevelFeaturesItems0FeatureIDAUTOASSIGNROLE captures enum value "AUTO_ASSIGN_ROLE"
	FeatureSupportLevelFeaturesItems0FeatureIDAUTOASSIGNROLE string = "AUTO_ASSIGN_ROLE"

	// FeatureSupportLevelFeaturesItems0FeatureIDCUSTOMMANIFEST captures enum value "CUSTOM_MANIFEST"
	FeatureSupportLevelFeaturesItems0FeatureIDCUSTOMMANIFEST string = "CUSTOM_MANIFEST"

	// FeatureSupportLevelFeaturesItems0FeatureIDDISKENCRYPTION captures enum value "DISK_ENCRYPTION"
	FeatureSupportLevelFeaturesItems0FeatureIDDISKENCRYPTION string = "DISK_ENCRYPTION"

	// FeatureSupportLevelFeaturesItems0FeatureIDCLUSTERMANAGEDNETWORKINGWITHVMS captures enum value "CLUSTER_MANAGED_NETWORKING_WITH_VMS"
	FeatureSupportLevelFeaturesItems0FeatureIDCLUSTERMANAGEDNETWORKINGWITHVMS string = "CLUSTER_MANAGED_NETWORKING_WITH_VMS"

	// FeatureSupportLevelFeaturesItems0FeatureIDARM64ARCHITECTURE captures enum value "ARM64_ARCHITECTURE"
	FeatureSupportLevelFeaturesItems0FeatureIDARM64ARCHITECTURE string = "ARM64_ARCHITECTURE"

	// FeatureSupportLevelFeaturesItems0FeatureIDARM64ARCHITECTUREWITHCLUSTERMANAGEDNETWORKING captures enum value "ARM64_ARCHITECTURE_WITH_CLUSTER_MANAGED_NETWORKING"
	FeatureSupportLevelFeaturesItems0FeatureIDARM64ARCHITECTUREWITHCLUSTERMANAGEDNETWORKING string = "ARM64_ARCHITECTURE_WITH_CLUSTER_MANAGED_NETWORKING"

	// FeatureSupportLevelFeaturesItems0FeatureIDPPC64LEARCHITECTURE captures enum value "PPC64LE_ARCHITECTURE"
	FeatureSupportLevelFeaturesItems0FeatureIDPPC64LEARCHITECTURE string = "PPC64LE_ARCHITECTURE"

	// FeatureSupportLevelFeaturesItems0FeatureIDS390XARCHITECTURE captures enum value "S390X_ARCHITECTURE"
	FeatureSupportLevelFeaturesItems0FeatureIDS390XARCHITECTURE string = "S390X_ARCHITECTURE"

	// FeatureSupportLevelFeaturesItems0FeatureIDSINGLENODEEXPANSION captures enum value "SINGLE_NODE_EXPANSION"
	FeatureSupportLevelFeaturesItems0FeatureIDSINGLENODEEXPANSION string = "SINGLE_NODE_EXPANSION"

	// FeatureSupportLevelFeaturesItems0FeatureIDLVM captures enum value "LVM"
	FeatureSupportLevelFeaturesItems0FeatureIDLVM string = "LVM"

	// FeatureSupportLevelFeaturesItems0FeatureIDDUALSTACKNETWORKING captures enum value "DUAL_STACK_NETWORKING"
	FeatureSupportLevelFeaturesItems0FeatureIDDUALSTACKNETWORKING string = "DUAL_STACK_NETWORKING"

	// FeatureSupportLevelFeaturesItems0FeatureIDMULTIARCHRELEASEIMAGE captures enum value "MULTIARCH_RELEASE_IMAGE"
	FeatureSupportLevelFeaturesItems0FeatureIDMULTIARCHRELEASEIMAGE string = "MULTIARCH_RELEASE_IMAGE"

	// FeatureSupportLevelFeaturesItems0FeatureIDNUTANIXINTEGRATION captures enum value "NUTANIX_INTEGRATION"
	FeatureSupportLevelFeaturesItems0FeatureIDNUTANIXINTEGRATION string = "NUTANIX_INTEGRATION"

	// FeatureSupportLevelFeaturesItems0FeatureIDDUALSTACKVIPS captures enum value "DUAL_STACK_VIPS"
	FeatureSupportLevelFeaturesItems0FeatureIDDUALSTACKVIPS string = "DUAL_STACK_VIPS"

	// FeatureSupportLevelFeaturesItems0FeatureIDUSERMANAGEDNETWORKINGWITHMULTINODE captures enum value "USER_MANAGED_NETWORKING_WITH_MULTI_NODE"
	FeatureSupportLevelFeaturesItems0FeatureIDUSERMANAGEDNETWORKINGWITHMULTINODE string = "USER_MANAGED_NETWORKING_WITH_MULTI_NODE"
)

// prop value enum
func (m *FeatureSupportLevelFeaturesItems0) validateFeatureIDEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, featureSupportLevelFeaturesItems0TypeFeatureIDPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FeatureSupportLevelFeaturesItems0) validateFeatureID(formats strfmt.Registry) error {

	if err := validate.Required("feature_id", "body", m.FeatureID); err != nil {
		return err
	}

	// value enum
	if err := m.validateFeatureIDEnum("feature_id", "body", *m.FeatureID); err != nil {
		return err
	}

	return nil
}

var featureSupportLevelFeaturesItems0TypeSupportLevelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["supported","unsupported","tech-preview","dev-preview"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		featureSupportLevelFeaturesItems0TypeSupportLevelPropEnum = append(featureSupportLevelFeaturesItems0TypeSupportLevelPropEnum, v)
	}
}

const (

	// FeatureSupportLevelFeaturesItems0SupportLevelSupported captures enum value "supported"
	FeatureSupportLevelFeaturesItems0SupportLevelSupported string = "supported"

	// FeatureSupportLevelFeaturesItems0SupportLevelUnsupported captures enum value "unsupported"
	FeatureSupportLevelFeaturesItems0SupportLevelUnsupported string = "unsupported"

	// FeatureSupportLevelFeaturesItems0SupportLevelTechPreview captures enum value "tech-preview"
	FeatureSupportLevelFeaturesItems0SupportLevelTechPreview string = "tech-preview"

	// FeatureSupportLevelFeaturesItems0SupportLevelDevPreview captures enum value "dev-preview"
	FeatureSupportLevelFeaturesItems0SupportLevelDevPreview string = "dev-preview"
)

// prop value enum
func (m *FeatureSupportLevelFeaturesItems0) validateSupportLevelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, featureSupportLevelFeaturesItems0TypeSupportLevelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FeatureSupportLevelFeaturesItems0) validateSupportLevel(formats strfmt.Registry) error {

	if err := validate.Required("support_level", "body", m.SupportLevel); err != nil {
		return err
	}

	// value enum
	if err := m.validateSupportLevelEnum("support_level", "body", *m.SupportLevel); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this feature support level features items0 based on context it is used
func (m *FeatureSupportLevelFeaturesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FeatureSupportLevelFeaturesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FeatureSupportLevelFeaturesItems0) UnmarshalBinary(b []byte) error {
	var res FeatureSupportLevelFeaturesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
