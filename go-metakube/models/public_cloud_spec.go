// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PublicCloudSpec PublicCloudSpec is a public counterpart of apiv1.CloudSpec.
//
// swagger:model PublicCloudSpec
type PublicCloudSpec struct {

	// datacenter name
	DatacenterName string `json:"dc,omitempty"`

	// alibaba
	Alibaba PublicAlibabaCloudSpec `json:"alibaba,omitempty"`

	// aws
	Aws PublicAWSCloudSpec `json:"aws,omitempty"`

	// azure
	Azure PublicAzureCloudSpec `json:"azure,omitempty"`

	// bringyourown
	Bringyourown PublicBringYourOwnCloudSpec `json:"bringyourown,omitempty"`

	// digitalocean
	Digitalocean PublicDigitaloceanCloudSpec `json:"digitalocean,omitempty"`

	// fake
	Fake PublicFakeCloudSpec `json:"fake,omitempty"`

	// gcp
	Gcp PublicGCPCloudSpec `json:"gcp,omitempty"`

	// hetzner
	Hetzner PublicHetznerCloudSpec `json:"hetzner,omitempty"`

	// kubevirt
	Kubevirt PublicKubevirtCloudSpec `json:"kubevirt,omitempty"`

	// openstack
	Openstack *PublicOpenstackCloudSpec `json:"openstack,omitempty"`

	// packet
	Packet PublicPacketCloudSpec `json:"packet,omitempty"`

	// vsphere
	Vsphere PublicVSphereCloudSpec `json:"vsphere,omitempty"`
}

// Validate validates this public cloud spec
func (m *PublicCloudSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOpenstack(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PublicCloudSpec) validateOpenstack(formats strfmt.Registry) error {

	if swag.IsZero(m.Openstack) { // not required
		return nil
	}

	if m.Openstack != nil {
		if err := m.Openstack.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("openstack")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PublicCloudSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PublicCloudSpec) UnmarshalBinary(b []byte) error {
	var res PublicCloudSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
