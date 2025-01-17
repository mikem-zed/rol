package mappers

import (
	"rol/app/errors"
	"rol/domain"
	"rol/dtos"
)

//MapDtoToEntity map a DTO to its corresponding entity
//Params
//	dto - DTO struct
//	entity - pointer to the entity
//Return
//  error - if error occurs return error, otherwise nil
func MapDtoToEntity(dto interface{}, entity interface{}) error {
	switch dto.(type) {
	// EthernetSwitch
	case dtos.EthernetSwitchCreateDto:
		MapEthernetSwitchCreateDto(dto.(dtos.EthernetSwitchCreateDto), entity.(*domain.EthernetSwitch))
	case dtos.EthernetSwitchUpdateDto:
		MapEthernetSwitchUpdateDto(dto.(dtos.EthernetSwitchUpdateDto), entity.(*domain.EthernetSwitch))
	//EthernetSwitchPort
	case dtos.EthernetSwitchPortCreateDto:
		MapEthernetSwitchPortCreateDto(dto.(dtos.EthernetSwitchPortCreateDto), entity.(*domain.EthernetSwitchPort))
	case dtos.EthernetSwitchPortUpdateDto:
		MapEthernetSwitchPortUpdateDto(dto.(dtos.EthernetSwitchPortUpdateDto), entity.(*domain.EthernetSwitchPort))
	//HostNetworkVlan
	case dtos.HostNetworkVlanCreateDto:
		MapHostNetworkVlanCreateDtoToEntity(dto.(dtos.HostNetworkVlanCreateDto), entity.(*domain.HostNetworkVlan))
	case dtos.HostNetworkVlanUpdateDto:
		MapHostNetworkVlanUpdateDtoToEntity(dto.(dtos.HostNetworkVlanUpdateDto), entity.(*domain.HostNetworkVlan))
	//HostNetworkBridge
	case dtos.HostNetworkBridgeCreateDto:
		MapHostNetworkBridgeCreateDtoToEntity(dto.(dtos.HostNetworkBridgeCreateDto), entity.(*domain.HostNetworkBridge))
	case dtos.HostNetworkBridgeUpdateDto:
		MapHostNetworkBridgeUpdateDtoToEntity(dto.(dtos.HostNetworkBridgeUpdateDto), entity.(*domain.HostNetworkBridge))
	//EthernetSwitchVLAN
	case dtos.EthernetSwitchVLANCreateDto:
		MapEthernetSwitchVLANCreateDto(dto.(dtos.EthernetSwitchVLANCreateDto), entity.(*domain.EthernetSwitchVLAN))
	case dtos.EthernetSwitchVLANUpdateDto:
		MapEthernetSwitchVLANUpdateDto(dto.(dtos.EthernetSwitchVLANUpdateDto), entity.(*domain.EthernetSwitchVLAN))
	//DHCPServer
	case dtos.DHCP4ServerCreateDto:
		MapDHCP4ServerCreateDtoToEntity(dto.(dtos.DHCP4ServerCreateDto), entity.(*domain.DHCP4Config))
	case dtos.DHCP4ServerUpdateDto:
		MapDHCP4ServerUpdateDtoToEntity(dto.(dtos.DHCP4ServerUpdateDto), entity.(*domain.DHCP4Config))
	case dtos.DHCP4LeaseCreateDto:
		MapDHCP4LeaseCreateDtoToEntity(dto.(dtos.DHCP4LeaseCreateDto), entity.(*domain.DHCP4Lease))
	case dtos.DHCP4LeaseUpdateDto:
		MapDHCP4LeaseUpdateDtoToEntity(dto.(dtos.DHCP4LeaseUpdateDto), entity.(*domain.DHCP4Lease))
	default:
		return errors.Internal.Newf("can't find route for map dto %+v to entity %+v", dto, entity)
	}
	return nil
}

//MapEntityToDto map entity to DTO
//Params
//	entity - Entity struct
//	dto - dest DTO
//Return
//  error - if error occurs return error, otherwise nil
func MapEntityToDto(entity interface{}, dto interface{}) error {
	switch entity.(type) {
	// EthernetSwitch
	case domain.EthernetSwitch:
		MapEthernetSwitchToDto(entity.(domain.EthernetSwitch), dto.(*dtos.EthernetSwitchDto))
	//EthernetSwitchPort
	case domain.EthernetSwitchPort:
		MapEthernetSwitchPortToDto(entity.(domain.EthernetSwitchPort), dto.(*dtos.EthernetSwitchPortDto))
	//HTTPLog
	case domain.HTTPLog:
		MapHTTPLogEntityToDto(entity.(domain.HTTPLog), dto.(*dtos.HTTPLogDto))
	//AppLog
	case domain.AppLog:
		MapAppLogEntityToDto(entity.(domain.AppLog), dto.(*dtos.AppLogDto))
	//DeviceTemplate
	case domain.DeviceTemplate:
		MapDeviceTemplateToDto(entity.(domain.DeviceTemplate), dto.(*dtos.DeviceTemplateDto))
	//HostNetworkVlan
	case domain.HostNetworkVlan:
		MapHostNetworkVlanToDto(entity.(domain.HostNetworkVlan), dto.(*dtos.HostNetworkVlanDto))
	//HostNetworkBridge
	case domain.HostNetworkBridge:
		MapHostNetworkBridgeToDto(entity.(domain.HostNetworkBridge), dto.(*dtos.HostNetworkBridgeDto))
	//EthernetSwitchVLAN
	case domain.EthernetSwitchVLAN:
		MapEthernetSwitchVLANToDto(entity.(domain.EthernetSwitchVLAN), dto.(*dtos.EthernetSwitchVLANDto))
	//DHCP4Server
	case domain.DHCP4Config:
		MapDHCP4ServerToDto(entity.(domain.DHCP4Config), dto.(*dtos.DHCP4ServerDto))
	//DHCP4Lease
	case domain.DHCP4Lease:
		MapDHCP4LeaseToDto(entity.(domain.DHCP4Lease), dto.(*dtos.DHCP4LeaseDto))

	default:
		return errors.Internal.Newf("can't find route for map entity %+v to dto %+v", dto, entity)
	}
	return nil
}
