// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Fabric A Fibre Channel (FC) fabric REST object provides information about an FC network (fabric) connected to the cluster. Logically, the FC fabric also contains FC switches and the FC zones that comprise the active zoneset of the fabric. FC switch and zone infromation is not reported directly in the FC fabric REST object for reasons of scale and flexibility; they are found by querying the FC switches and FC zones REST endpoints.
//
// swagger:model fabric
type Fabric struct {

	// links
	Links *FabricInlineLinks `json:"_links,omitempty"`

	// cache
	Cache *FabricInlineCache `json:"cache,omitempty"`

	// An array of the connections between the cluster and the switches Fibre Channel fabric.
	//
	FabricInlineConnections []*FabricInlineConnectionsInlineArrayItem `json:"connections,omitempty"`

	// The world wide name (WWN) of the primary switch of the Fibre Channel (FC) fabric. This is used as a unique identifier for the FC fabric.
	//
	// Example: 10:00:c1:c2:c3:c4:c5:c6
	// Read Only: true
	Name *string `json:"name,omitempty"`

	// zoneset
	Zoneset *FabricInlineZoneset `json:"zoneset,omitempty"`
}

// Validate validates this fabric
func (m *Fabric) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCache(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFabricInlineConnections(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateZoneset(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Fabric) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *Fabric) validateCache(formats strfmt.Registry) error {
	if swag.IsZero(m.Cache) { // not required
		return nil
	}

	if m.Cache != nil {
		if err := m.Cache.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cache")
			}
			return err
		}
	}

	return nil
}

func (m *Fabric) validateFabricInlineConnections(formats strfmt.Registry) error {
	if swag.IsZero(m.FabricInlineConnections) { // not required
		return nil
	}

	for i := 0; i < len(m.FabricInlineConnections); i++ {
		if swag.IsZero(m.FabricInlineConnections[i]) { // not required
			continue
		}

		if m.FabricInlineConnections[i] != nil {
			if err := m.FabricInlineConnections[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("connections" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Fabric) validateZoneset(formats strfmt.Registry) error {
	if swag.IsZero(m.Zoneset) { // not required
		return nil
	}

	if m.Zoneset != nil {
		if err := m.Zoneset.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("zoneset")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this fabric based on the context it is used
func (m *Fabric) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCache(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFabricInlineConnections(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateZoneset(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Fabric) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *Fabric) contextValidateCache(ctx context.Context, formats strfmt.Registry) error {

	if m.Cache != nil {
		if err := m.Cache.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cache")
			}
			return err
		}
	}

	return nil
}

func (m *Fabric) contextValidateFabricInlineConnections(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.FabricInlineConnections); i++ {

		if m.FabricInlineConnections[i] != nil {
			if err := m.FabricInlineConnections[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("connections" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Fabric) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Fabric) contextValidateZoneset(ctx context.Context, formats strfmt.Registry) error {

	if m.Zoneset != nil {
		if err := m.Zoneset.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("zoneset")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Fabric) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Fabric) UnmarshalBinary(b []byte) error {
	var res Fabric
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FabricInlineCache Properties of Fibre Chanel fabric cache.
//
// swagger:model fabric_inline_cache
type FabricInlineCache struct {

	// The age of the Fibre Channel fabric data cache retrieved. If the FC fabric data cache has not been fully updated for a newly discovered fabric, or a fabric that has been re-discovered after being purged, a value for this property will not be retrieved. The value is in ISO 8601 duration format.
	//
	// Example: PT3M30S
	// Read Only: true
	Age *string `json:"age,omitempty"`

	// A boolean that indicates if the retrieved data is current relative to the `cache.maximum_age` value of the request. A value of `true` indicates that the data is no older than the requested maximum age. A value of `false` indicates that the data is older than the requested maximum age; if more current data is required, the caller should wait for some time for the cache update to complete and query the data again.
	//
	// Read Only: true
	IsCurrent *bool `json:"is_current,omitempty"`

	// The date and time at which the Fibre Channel fabric data cache retrieved was last updated. If the FC fabric data cache has not been fully updated for a newly discovered fabric, or a fabric that has been re-discovered after being purged, a value for this property will not be retrieved.
	//
	// Read Only: true
	// Format: date-time
	UpdateTime *strfmt.DateTime `json:"update_time,omitempty"`
}

// Validate validates this fabric inline cache
func (m *FabricInlineCache) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FabricInlineCache) validateUpdateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("cache"+"."+"update_time", "body", "date-time", m.UpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this fabric inline cache based on the context it is used
func (m *FabricInlineCache) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAge(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIsCurrent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdateTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FabricInlineCache) contextValidateAge(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "cache"+"."+"age", "body", m.Age); err != nil {
		return err
	}

	return nil
}

func (m *FabricInlineCache) contextValidateIsCurrent(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "cache"+"."+"is_current", "body", m.IsCurrent); err != nil {
		return err
	}

	return nil
}

func (m *FabricInlineCache) contextValidateUpdateTime(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "cache"+"."+"update_time", "body", m.UpdateTime); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FabricInlineCache) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FabricInlineCache) UnmarshalBinary(b []byte) error {
	var res FabricInlineCache
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FabricInlineConnectionsInlineArrayItem A connection between a cluster node Fibre Channel (FC) port and an FC switch port.
//
// swagger:model fabric_inline_connections_inline_array_item
type FabricInlineConnectionsInlineArrayItem struct {

	// cluster port
	ClusterPort *FabricInlineConnectionsInlineArrayItemInlineClusterPort `json:"cluster_port,omitempty"`

	// switch
	Switch *FabricInlineConnectionsInlineArrayItemInlineSwitch `json:"switch,omitempty"`
}

// Validate validates this fabric inline connections inline array item
func (m *FabricInlineConnectionsInlineArrayItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterPort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSwitch(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FabricInlineConnectionsInlineArrayItem) validateClusterPort(formats strfmt.Registry) error {
	if swag.IsZero(m.ClusterPort) { // not required
		return nil
	}

	if m.ClusterPort != nil {
		if err := m.ClusterPort.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster_port")
			}
			return err
		}
	}

	return nil
}

func (m *FabricInlineConnectionsInlineArrayItem) validateSwitch(formats strfmt.Registry) error {
	if swag.IsZero(m.Switch) { // not required
		return nil
	}

	if m.Switch != nil {
		if err := m.Switch.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("switch")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this fabric inline connections inline array item based on the context it is used
func (m *FabricInlineConnectionsInlineArrayItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClusterPort(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSwitch(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FabricInlineConnectionsInlineArrayItem) contextValidateClusterPort(ctx context.Context, formats strfmt.Registry) error {

	if m.ClusterPort != nil {
		if err := m.ClusterPort.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster_port")
			}
			return err
		}
	}

	return nil
}

func (m *FabricInlineConnectionsInlineArrayItem) contextValidateSwitch(ctx context.Context, formats strfmt.Registry) error {

	if m.Switch != nil {
		if err := m.Switch.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("switch")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FabricInlineConnectionsInlineArrayItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FabricInlineConnectionsInlineArrayItem) UnmarshalBinary(b []byte) error {
	var res FabricInlineConnectionsInlineArrayItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FabricInlineConnectionsInlineArrayItemInlineClusterPort The cluster Fibre Channel (FC) port that connects the FC fabric.
//
// swagger:model fabric_inline_connections_inline_array_item_inline_cluster_port
type FabricInlineConnectionsInlineArrayItemInlineClusterPort struct {

	// links
	Links *FabricInlineConnectionsInlineArrayItemInlineClusterPortInlineLinks `json:"_links,omitempty"`

	// The name of the cluster Fibre Channel port.
	//
	// Example: 0a
	Name *string `json:"name,omitempty"`

	// node
	Node *FabricInlineConnectionsInlineArrayItemInlineClusterPortInlineNode `json:"node,omitempty"`

	// The unique identifier of the cluster Fibre Channel port.
	//
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	UUID *string `json:"uuid,omitempty"`

	// The world wide port name (WWPN) of the cluster Fibre Channel port.
	//
	// Example: 50:0a:11:22:33:44:55:66
	// Read Only: true
	Wwpn *string `json:"wwpn,omitempty"`
}

// Validate validates this fabric inline connections inline array item inline cluster port
func (m *FabricInlineConnectionsInlineArrayItemInlineClusterPort) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FabricInlineConnectionsInlineArrayItemInlineClusterPort) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster_port" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

func (m *FabricInlineConnectionsInlineArrayItemInlineClusterPort) validateNode(formats strfmt.Registry) error {
	if swag.IsZero(m.Node) { // not required
		return nil
	}

	if m.Node != nil {
		if err := m.Node.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster_port" + "." + "node")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this fabric inline connections inline array item inline cluster port based on the context it is used
func (m *FabricInlineConnectionsInlineArrayItemInlineClusterPort) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWwpn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FabricInlineConnectionsInlineArrayItemInlineClusterPort) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster_port" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

func (m *FabricInlineConnectionsInlineArrayItemInlineClusterPort) contextValidateNode(ctx context.Context, formats strfmt.Registry) error {

	if m.Node != nil {
		if err := m.Node.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster_port" + "." + "node")
			}
			return err
		}
	}

	return nil
}

func (m *FabricInlineConnectionsInlineArrayItemInlineClusterPort) contextValidateWwpn(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "cluster_port"+"."+"wwpn", "body", m.Wwpn); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FabricInlineConnectionsInlineArrayItemInlineClusterPort) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FabricInlineConnectionsInlineArrayItemInlineClusterPort) UnmarshalBinary(b []byte) error {
	var res FabricInlineConnectionsInlineArrayItemInlineClusterPort
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FabricInlineConnectionsInlineArrayItemInlineClusterPortInlineLinks fabric inline connections inline array item inline cluster port inline links
//
// swagger:model fabric_inline_connections_inline_array_item_inline_cluster_port_inline__links
type FabricInlineConnectionsInlineArrayItemInlineClusterPortInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this fabric inline connections inline array item inline cluster port inline links
func (m *FabricInlineConnectionsInlineArrayItemInlineClusterPortInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FabricInlineConnectionsInlineArrayItemInlineClusterPortInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster_port" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this fabric inline connections inline array item inline cluster port inline links based on the context it is used
func (m *FabricInlineConnectionsInlineArrayItemInlineClusterPortInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FabricInlineConnectionsInlineArrayItemInlineClusterPortInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster_port" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FabricInlineConnectionsInlineArrayItemInlineClusterPortInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FabricInlineConnectionsInlineArrayItemInlineClusterPortInlineLinks) UnmarshalBinary(b []byte) error {
	var res FabricInlineConnectionsInlineArrayItemInlineClusterPortInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FabricInlineConnectionsInlineArrayItemInlineClusterPortInlineNode The node on which the cluster Fibre Channel port is located.
//
// swagger:model fabric_inline_connections_inline_array_item_inline_cluster_port_inline_node
type FabricInlineConnectionsInlineArrayItemInlineClusterPortInlineNode struct {

	// The name of the node on which the cluster Fibre Channel port is located.
	//
	// Example: node1
	// Read Only: true
	Name *string `json:"name,omitempty"`
}

// Validate validates this fabric inline connections inline array item inline cluster port inline node
func (m *FabricInlineConnectionsInlineArrayItemInlineClusterPortInlineNode) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this fabric inline connections inline array item inline cluster port inline node based on the context it is used
func (m *FabricInlineConnectionsInlineArrayItemInlineClusterPortInlineNode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FabricInlineConnectionsInlineArrayItemInlineClusterPortInlineNode) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "cluster_port"+"."+"node"+"."+"name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FabricInlineConnectionsInlineArrayItemInlineClusterPortInlineNode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FabricInlineConnectionsInlineArrayItemInlineClusterPortInlineNode) UnmarshalBinary(b []byte) error {
	var res FabricInlineConnectionsInlineArrayItemInlineClusterPortInlineNode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FabricInlineConnectionsInlineArrayItemInlineSwitch The Fibre Channel switch to which the cluster node port is connected.
//
// swagger:model fabric_inline_connections_inline_array_item_inline_switch
type FabricInlineConnectionsInlineArrayItemInlineSwitch struct {

	// links
	Links *FabricInlineConnectionsInlineArrayItemInlineSwitchInlineLinks `json:"_links,omitempty"`

	// port
	Port *FabricInlineConnectionsInlineArrayItemInlineSwitchInlinePort `json:"port,omitempty"`

	// The world-wide name (WWN) of the Fibre Channel switch to which the cluster node port is attached.
	//
	// Example: 10:00:b1:b2:b3:b4:b4:b6
	// Read Only: true
	Wwn *string `json:"wwn,omitempty"`
}

// Validate validates this fabric inline connections inline array item inline switch
func (m *FabricInlineConnectionsInlineArrayItemInlineSwitch) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FabricInlineConnectionsInlineArrayItemInlineSwitch) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("switch" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

func (m *FabricInlineConnectionsInlineArrayItemInlineSwitch) validatePort(formats strfmt.Registry) error {
	if swag.IsZero(m.Port) { // not required
		return nil
	}

	if m.Port != nil {
		if err := m.Port.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("switch" + "." + "port")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this fabric inline connections inline array item inline switch based on the context it is used
func (m *FabricInlineConnectionsInlineArrayItemInlineSwitch) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePort(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWwn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FabricInlineConnectionsInlineArrayItemInlineSwitch) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("switch" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

func (m *FabricInlineConnectionsInlineArrayItemInlineSwitch) contextValidatePort(ctx context.Context, formats strfmt.Registry) error {

	if m.Port != nil {
		if err := m.Port.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("switch" + "." + "port")
			}
			return err
		}
	}

	return nil
}

func (m *FabricInlineConnectionsInlineArrayItemInlineSwitch) contextValidateWwn(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "switch"+"."+"wwn", "body", m.Wwn); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FabricInlineConnectionsInlineArrayItemInlineSwitch) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FabricInlineConnectionsInlineArrayItemInlineSwitch) UnmarshalBinary(b []byte) error {
	var res FabricInlineConnectionsInlineArrayItemInlineSwitch
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FabricInlineConnectionsInlineArrayItemInlineSwitchInlineLinks fabric inline connections inline array item inline switch inline links
//
// swagger:model fabric_inline_connections_inline_array_item_inline_switch_inline__links
type FabricInlineConnectionsInlineArrayItemInlineSwitchInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this fabric inline connections inline array item inline switch inline links
func (m *FabricInlineConnectionsInlineArrayItemInlineSwitchInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FabricInlineConnectionsInlineArrayItemInlineSwitchInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("switch" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this fabric inline connections inline array item inline switch inline links based on the context it is used
func (m *FabricInlineConnectionsInlineArrayItemInlineSwitchInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FabricInlineConnectionsInlineArrayItemInlineSwitchInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("switch" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FabricInlineConnectionsInlineArrayItemInlineSwitchInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FabricInlineConnectionsInlineArrayItemInlineSwitchInlineLinks) UnmarshalBinary(b []byte) error {
	var res FabricInlineConnectionsInlineArrayItemInlineSwitchInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FabricInlineConnectionsInlineArrayItemInlineSwitchInlinePort The port of the Fibre Channel switch to which the cluster node port is connected.
//
// swagger:model fabric_inline_connections_inline_array_item_inline_switch_inline_port
type FabricInlineConnectionsInlineArrayItemInlineSwitchInlinePort struct {

	// The world wide port name (WWPN) of the Fibre Channel switch port.
	//
	// Example: 50:0a:a1:a2:a3:a4:a5:a6
	// Read Only: true
	Wwpn *string `json:"wwpn,omitempty"`
}

// Validate validates this fabric inline connections inline array item inline switch inline port
func (m *FabricInlineConnectionsInlineArrayItemInlineSwitchInlinePort) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this fabric inline connections inline array item inline switch inline port based on the context it is used
func (m *FabricInlineConnectionsInlineArrayItemInlineSwitchInlinePort) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateWwpn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FabricInlineConnectionsInlineArrayItemInlineSwitchInlinePort) contextValidateWwpn(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "switch"+"."+"port"+"."+"wwpn", "body", m.Wwpn); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FabricInlineConnectionsInlineArrayItemInlineSwitchInlinePort) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FabricInlineConnectionsInlineArrayItemInlineSwitchInlinePort) UnmarshalBinary(b []byte) error {
	var res FabricInlineConnectionsInlineArrayItemInlineSwitchInlinePort
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FabricInlineLinks fabric inline links
//
// swagger:model fabric_inline__links
type FabricInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this fabric inline links
func (m *FabricInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FabricInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this fabric inline links based on the context it is used
func (m *FabricInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FabricInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FabricInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FabricInlineLinks) UnmarshalBinary(b []byte) error {
	var res FabricInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FabricInlineZoneset The active Fibre Channel zoneset in the fabric.
//
// swagger:model fabric_inline_zoneset
type FabricInlineZoneset struct {

	// The name of the Fibre Channel zoneset.
	//
	// Example: zoneset1
	// Read Only: true
	Name *string `json:"name,omitempty"`
}

// Validate validates this fabric inline zoneset
func (m *FabricInlineZoneset) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this fabric inline zoneset based on the context it is used
func (m *FabricInlineZoneset) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FabricInlineZoneset) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "zoneset"+"."+"name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FabricInlineZoneset) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FabricInlineZoneset) UnmarshalBinary(b []byte) error {
	var res FabricInlineZoneset
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
