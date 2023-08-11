# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'VpcEndpointConnectionZone',
    'VpcEndpointConnectionsConnectionResult',
    'VpcEndpointConnectionsConnectionZoneResult',
    'VpcEndpointServicePermissionsPermissionResult',
    'VpcEndpointServiceResource',
    'VpcEndpointServicesServiceResult',
    'VpcEndpointServicesServiceResourceResult',
    'VpcEndpointZonesVpcEndpointZoneResult',
    'VpcEndpointsVpcEndpointResult',
]

@pulumi.output_type
class VpcEndpointConnectionZone(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "networkInterfaceId":
            suggest = "network_interface_id"
        elif key == "networkInterfaceIp":
            suggest = "network_interface_ip"
        elif key == "resourceId":
            suggest = "resource_id"
        elif key == "subnetId":
            suggest = "subnet_id"
        elif key == "zoneDomain":
            suggest = "zone_domain"
        elif key == "zoneId":
            suggest = "zone_id"
        elif key == "zoneStatus":
            suggest = "zone_status"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in VpcEndpointConnectionZone. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        VpcEndpointConnectionZone.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        VpcEndpointConnectionZone.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 network_interface_id: Optional[str] = None,
                 network_interface_ip: Optional[str] = None,
                 resource_id: Optional[str] = None,
                 subnet_id: Optional[str] = None,
                 zone_domain: Optional[str] = None,
                 zone_id: Optional[str] = None,
                 zone_status: Optional[str] = None):
        """
        :param str network_interface_id: The id of the network interface.
        :param str network_interface_ip: The ip address of the network interface.
        :param str resource_id: The id of the resource.
        :param str subnet_id: The id of the subnet.
        :param str zone_domain: The domain of the zone.
        :param str zone_id: The id of the zone.
        :param str zone_status: The status of the zone.
        """
        if network_interface_id is not None:
            pulumi.set(__self__, "network_interface_id", network_interface_id)
        if network_interface_ip is not None:
            pulumi.set(__self__, "network_interface_ip", network_interface_ip)
        if resource_id is not None:
            pulumi.set(__self__, "resource_id", resource_id)
        if subnet_id is not None:
            pulumi.set(__self__, "subnet_id", subnet_id)
        if zone_domain is not None:
            pulumi.set(__self__, "zone_domain", zone_domain)
        if zone_id is not None:
            pulumi.set(__self__, "zone_id", zone_id)
        if zone_status is not None:
            pulumi.set(__self__, "zone_status", zone_status)

    @property
    @pulumi.getter(name="networkInterfaceId")
    def network_interface_id(self) -> Optional[str]:
        """
        The id of the network interface.
        """
        return pulumi.get(self, "network_interface_id")

    @property
    @pulumi.getter(name="networkInterfaceIp")
    def network_interface_ip(self) -> Optional[str]:
        """
        The ip address of the network interface.
        """
        return pulumi.get(self, "network_interface_ip")

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> Optional[str]:
        """
        The id of the resource.
        """
        return pulumi.get(self, "resource_id")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> Optional[str]:
        """
        The id of the subnet.
        """
        return pulumi.get(self, "subnet_id")

    @property
    @pulumi.getter(name="zoneDomain")
    def zone_domain(self) -> Optional[str]:
        """
        The domain of the zone.
        """
        return pulumi.get(self, "zone_domain")

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> Optional[str]:
        """
        The id of the zone.
        """
        return pulumi.get(self, "zone_id")

    @property
    @pulumi.getter(name="zoneStatus")
    def zone_status(self) -> Optional[str]:
        """
        The status of the zone.
        """
        return pulumi.get(self, "zone_status")


@pulumi.output_type
class VpcEndpointConnectionsConnectionResult(dict):
    def __init__(__self__, *,
                 connection_status: str,
                 creation_time: str,
                 endpoint_id: str,
                 endpoint_owner_account_id: str,
                 endpoint_vpc_id: str,
                 service_id: str,
                 update_time: str,
                 zones: Sequence['outputs.VpcEndpointConnectionsConnectionZoneResult']):
        """
        :param str connection_status: The status of the connection.
        :param str creation_time: The create time of the connection.
        :param str endpoint_id: The id of the vpc endpoint.
        :param str endpoint_owner_account_id: The account id of the vpc endpoint.
        :param str endpoint_vpc_id: The vpc id of the vpc endpoint.
        :param str service_id: The id of the vpc endpoint service.
        :param str update_time: The update time of the connection.
        :param Sequence['VpcEndpointConnectionsConnectionZoneArgs'] zones: The available zones.
        """
        pulumi.set(__self__, "connection_status", connection_status)
        pulumi.set(__self__, "creation_time", creation_time)
        pulumi.set(__self__, "endpoint_id", endpoint_id)
        pulumi.set(__self__, "endpoint_owner_account_id", endpoint_owner_account_id)
        pulumi.set(__self__, "endpoint_vpc_id", endpoint_vpc_id)
        pulumi.set(__self__, "service_id", service_id)
        pulumi.set(__self__, "update_time", update_time)
        pulumi.set(__self__, "zones", zones)

    @property
    @pulumi.getter(name="connectionStatus")
    def connection_status(self) -> str:
        """
        The status of the connection.
        """
        return pulumi.get(self, "connection_status")

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> str:
        """
        The create time of the connection.
        """
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter(name="endpointId")
    def endpoint_id(self) -> str:
        """
        The id of the vpc endpoint.
        """
        return pulumi.get(self, "endpoint_id")

    @property
    @pulumi.getter(name="endpointOwnerAccountId")
    def endpoint_owner_account_id(self) -> str:
        """
        The account id of the vpc endpoint.
        """
        return pulumi.get(self, "endpoint_owner_account_id")

    @property
    @pulumi.getter(name="endpointVpcId")
    def endpoint_vpc_id(self) -> str:
        """
        The vpc id of the vpc endpoint.
        """
        return pulumi.get(self, "endpoint_vpc_id")

    @property
    @pulumi.getter(name="serviceId")
    def service_id(self) -> str:
        """
        The id of the vpc endpoint service.
        """
        return pulumi.get(self, "service_id")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        The update time of the connection.
        """
        return pulumi.get(self, "update_time")

    @property
    @pulumi.getter
    def zones(self) -> Sequence['outputs.VpcEndpointConnectionsConnectionZoneResult']:
        """
        The available zones.
        """
        return pulumi.get(self, "zones")


@pulumi.output_type
class VpcEndpointConnectionsConnectionZoneResult(dict):
    def __init__(__self__, *,
                 network_interface_id: str,
                 network_interface_ip: str,
                 resource_id: str,
                 subnet_id: str,
                 zone_domain: str,
                 zone_id: str,
                 zone_status: str):
        """
        :param str network_interface_id: The id of the network interface.
        :param str network_interface_ip: The ip address of the network interface.
        :param str resource_id: The id of the resource.
        :param str subnet_id: The id of the subnet.
        :param str zone_domain: The domain of the zone.
        :param str zone_id: The id of the zone.
        :param str zone_status: The status of the zone.
        """
        pulumi.set(__self__, "network_interface_id", network_interface_id)
        pulumi.set(__self__, "network_interface_ip", network_interface_ip)
        pulumi.set(__self__, "resource_id", resource_id)
        pulumi.set(__self__, "subnet_id", subnet_id)
        pulumi.set(__self__, "zone_domain", zone_domain)
        pulumi.set(__self__, "zone_id", zone_id)
        pulumi.set(__self__, "zone_status", zone_status)

    @property
    @pulumi.getter(name="networkInterfaceId")
    def network_interface_id(self) -> str:
        """
        The id of the network interface.
        """
        return pulumi.get(self, "network_interface_id")

    @property
    @pulumi.getter(name="networkInterfaceIp")
    def network_interface_ip(self) -> str:
        """
        The ip address of the network interface.
        """
        return pulumi.get(self, "network_interface_ip")

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> str:
        """
        The id of the resource.
        """
        return pulumi.get(self, "resource_id")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> str:
        """
        The id of the subnet.
        """
        return pulumi.get(self, "subnet_id")

    @property
    @pulumi.getter(name="zoneDomain")
    def zone_domain(self) -> str:
        """
        The domain of the zone.
        """
        return pulumi.get(self, "zone_domain")

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> str:
        """
        The id of the zone.
        """
        return pulumi.get(self, "zone_id")

    @property
    @pulumi.getter(name="zoneStatus")
    def zone_status(self) -> str:
        """
        The status of the zone.
        """
        return pulumi.get(self, "zone_status")


@pulumi.output_type
class VpcEndpointServicePermissionsPermissionResult(dict):
    def __init__(__self__, *,
                 permit_account_id: str):
        """
        :param str permit_account_id: The Id of permit account.
        """
        pulumi.set(__self__, "permit_account_id", permit_account_id)

    @property
    @pulumi.getter(name="permitAccountId")
    def permit_account_id(self) -> str:
        """
        The Id of permit account.
        """
        return pulumi.get(self, "permit_account_id")


@pulumi.output_type
class VpcEndpointServiceResource(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "resourceId":
            suggest = "resource_id"
        elif key == "resourceType":
            suggest = "resource_type"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in VpcEndpointServiceResource. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        VpcEndpointServiceResource.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        VpcEndpointServiceResource.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 resource_id: str,
                 resource_type: str):
        """
        :param str resource_id: The id of resource.
        :param str resource_type: The type of resource.
        """
        pulumi.set(__self__, "resource_id", resource_id)
        pulumi.set(__self__, "resource_type", resource_type)

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> str:
        """
        The id of resource.
        """
        return pulumi.get(self, "resource_id")

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> str:
        """
        The type of resource.
        """
        return pulumi.get(self, "resource_type")


@pulumi.output_type
class VpcEndpointServicesServiceResult(dict):
    def __init__(__self__, *,
                 auto_accept_enabled: bool,
                 creation_time: str,
                 description: str,
                 id: str,
                 resources: Sequence['outputs.VpcEndpointServicesServiceResourceResult'],
                 service_domain: str,
                 service_id: str,
                 service_name: str,
                 service_resource_type: str,
                 service_type: str,
                 status: str,
                 update_time: str,
                 zone_ids: Sequence[str]):
        """
        :param bool auto_accept_enabled: Whether auto accept node connect.
        :param str creation_time: The create time of service.
        :param str description: The description of service.
        :param str id: The Id of service.
        :param Sequence['VpcEndpointServicesServiceResourceArgs'] resources: The resources info.
        :param str service_domain: The domain of service.
        :param str service_id: The Id of service.
        :param str service_name: The name of vpc endpoint service.
        :param str service_resource_type: The resource type of service.
        :param str service_type: The type of service.
        :param str status: The status of service.
        :param str update_time: The update time of service.
        :param Sequence[str] zone_ids: The IDs of zones.
        """
        pulumi.set(__self__, "auto_accept_enabled", auto_accept_enabled)
        pulumi.set(__self__, "creation_time", creation_time)
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "resources", resources)
        pulumi.set(__self__, "service_domain", service_domain)
        pulumi.set(__self__, "service_id", service_id)
        pulumi.set(__self__, "service_name", service_name)
        pulumi.set(__self__, "service_resource_type", service_resource_type)
        pulumi.set(__self__, "service_type", service_type)
        pulumi.set(__self__, "status", status)
        pulumi.set(__self__, "update_time", update_time)
        pulumi.set(__self__, "zone_ids", zone_ids)

    @property
    @pulumi.getter(name="autoAcceptEnabled")
    def auto_accept_enabled(self) -> bool:
        """
        Whether auto accept node connect.
        """
        return pulumi.get(self, "auto_accept_enabled")

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> str:
        """
        The create time of service.
        """
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The description of service.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The Id of service.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def resources(self) -> Sequence['outputs.VpcEndpointServicesServiceResourceResult']:
        """
        The resources info.
        """
        return pulumi.get(self, "resources")

    @property
    @pulumi.getter(name="serviceDomain")
    def service_domain(self) -> str:
        """
        The domain of service.
        """
        return pulumi.get(self, "service_domain")

    @property
    @pulumi.getter(name="serviceId")
    def service_id(self) -> str:
        """
        The Id of service.
        """
        return pulumi.get(self, "service_id")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> str:
        """
        The name of vpc endpoint service.
        """
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter(name="serviceResourceType")
    def service_resource_type(self) -> str:
        """
        The resource type of service.
        """
        return pulumi.get(self, "service_resource_type")

    @property
    @pulumi.getter(name="serviceType")
    def service_type(self) -> str:
        """
        The type of service.
        """
        return pulumi.get(self, "service_type")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        The status of service.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        The update time of service.
        """
        return pulumi.get(self, "update_time")

    @property
    @pulumi.getter(name="zoneIds")
    def zone_ids(self) -> Sequence[str]:
        """
        The IDs of zones.
        """
        return pulumi.get(self, "zone_ids")


@pulumi.output_type
class VpcEndpointServicesServiceResourceResult(dict):
    def __init__(__self__, *,
                 resource_id: str,
                 resource_type: str,
                 zone_id: str):
        """
        :param str resource_id: The id of resource.
        :param str resource_type: The type of resource.
        :param str zone_id: The zone id of resource.
        """
        pulumi.set(__self__, "resource_id", resource_id)
        pulumi.set(__self__, "resource_type", resource_type)
        pulumi.set(__self__, "zone_id", zone_id)

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> str:
        """
        The id of resource.
        """
        return pulumi.get(self, "resource_id")

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> str:
        """
        The type of resource.
        """
        return pulumi.get(self, "resource_type")

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> str:
        """
        The zone id of resource.
        """
        return pulumi.get(self, "zone_id")


@pulumi.output_type
class VpcEndpointZonesVpcEndpointZoneResult(dict):
    def __init__(__self__, *,
                 id: str,
                 network_interface_id: str,
                 network_interface_ip: str,
                 service_status: str,
                 subnet_id: str,
                 zone_domain: str,
                 zone_id: str,
                 zone_status: str):
        """
        :param str id: The Id of vpc endpoint zone.
        :param str network_interface_id: The network interface id of vpc endpoint.
        :param str network_interface_ip: The network interface ip of vpc endpoint.
        :param str service_status: The status of vpc endpoint service.
        :param str subnet_id: The subnet id of vpc endpoint.
        :param str zone_domain: The domain of vpc endpoint zone.
        :param str zone_id: The Id of vpc endpoint zone.
        :param str zone_status: The status of vpc endpoint zone.
        """
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "network_interface_id", network_interface_id)
        pulumi.set(__self__, "network_interface_ip", network_interface_ip)
        pulumi.set(__self__, "service_status", service_status)
        pulumi.set(__self__, "subnet_id", subnet_id)
        pulumi.set(__self__, "zone_domain", zone_domain)
        pulumi.set(__self__, "zone_id", zone_id)
        pulumi.set(__self__, "zone_status", zone_status)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The Id of vpc endpoint zone.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="networkInterfaceId")
    def network_interface_id(self) -> str:
        """
        The network interface id of vpc endpoint.
        """
        return pulumi.get(self, "network_interface_id")

    @property
    @pulumi.getter(name="networkInterfaceIp")
    def network_interface_ip(self) -> str:
        """
        The network interface ip of vpc endpoint.
        """
        return pulumi.get(self, "network_interface_ip")

    @property
    @pulumi.getter(name="serviceStatus")
    def service_status(self) -> str:
        """
        The status of vpc endpoint service.
        """
        return pulumi.get(self, "service_status")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> str:
        """
        The subnet id of vpc endpoint.
        """
        return pulumi.get(self, "subnet_id")

    @property
    @pulumi.getter(name="zoneDomain")
    def zone_domain(self) -> str:
        """
        The domain of vpc endpoint zone.
        """
        return pulumi.get(self, "zone_domain")

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> str:
        """
        The Id of vpc endpoint zone.
        """
        return pulumi.get(self, "zone_id")

    @property
    @pulumi.getter(name="zoneStatus")
    def zone_status(self) -> str:
        """
        The status of vpc endpoint zone.
        """
        return pulumi.get(self, "zone_status")


@pulumi.output_type
class VpcEndpointsVpcEndpointResult(dict):
    def __init__(__self__, *,
                 business_status: str,
                 connection_status: str,
                 creation_time: str,
                 deleted_time: str,
                 description: str,
                 endpoint_domain: str,
                 endpoint_id: str,
                 endpoint_name: str,
                 endpoint_type: str,
                 id: str,
                 service_id: str,
                 service_name: str,
                 status: str,
                 update_time: str,
                 vpc_id: str):
        """
        :param str business_status: Whether the vpc endpoint is locked.
        :param str connection_status: The connection  status of vpc endpoint.
        :param str creation_time: The create time of vpc endpoint.
        :param str deleted_time: The delete time of vpc endpoint.
        :param str description: The description of vpc endpoint.
        :param str endpoint_domain: The domain of vpc endpoint.
        :param str endpoint_id: The Id of vpc endpoint.
        :param str endpoint_name: The name of vpc endpoint.
        :param str endpoint_type: The type of vpc endpoint.
        :param str id: The Id of vpc endpoint.
        :param str service_id: The Id of vpc endpoint service.
        :param str service_name: The name of vpc endpoint service.
        :param str status: The status of vpc endpoint. Valid values: `Creating`, `Pending`, `Available`, `Deleting`, `Inactive`.
        :param str update_time: The update time of vpc endpoint.
        :param str vpc_id: The vpc id of vpc endpoint.
        """
        pulumi.set(__self__, "business_status", business_status)
        pulumi.set(__self__, "connection_status", connection_status)
        pulumi.set(__self__, "creation_time", creation_time)
        pulumi.set(__self__, "deleted_time", deleted_time)
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "endpoint_domain", endpoint_domain)
        pulumi.set(__self__, "endpoint_id", endpoint_id)
        pulumi.set(__self__, "endpoint_name", endpoint_name)
        pulumi.set(__self__, "endpoint_type", endpoint_type)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "service_id", service_id)
        pulumi.set(__self__, "service_name", service_name)
        pulumi.set(__self__, "status", status)
        pulumi.set(__self__, "update_time", update_time)
        pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="businessStatus")
    def business_status(self) -> str:
        """
        Whether the vpc endpoint is locked.
        """
        return pulumi.get(self, "business_status")

    @property
    @pulumi.getter(name="connectionStatus")
    def connection_status(self) -> str:
        """
        The connection  status of vpc endpoint.
        """
        return pulumi.get(self, "connection_status")

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> str:
        """
        The create time of vpc endpoint.
        """
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter(name="deletedTime")
    def deleted_time(self) -> str:
        """
        The delete time of vpc endpoint.
        """
        return pulumi.get(self, "deleted_time")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The description of vpc endpoint.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="endpointDomain")
    def endpoint_domain(self) -> str:
        """
        The domain of vpc endpoint.
        """
        return pulumi.get(self, "endpoint_domain")

    @property
    @pulumi.getter(name="endpointId")
    def endpoint_id(self) -> str:
        """
        The Id of vpc endpoint.
        """
        return pulumi.get(self, "endpoint_id")

    @property
    @pulumi.getter(name="endpointName")
    def endpoint_name(self) -> str:
        """
        The name of vpc endpoint.
        """
        return pulumi.get(self, "endpoint_name")

    @property
    @pulumi.getter(name="endpointType")
    def endpoint_type(self) -> str:
        """
        The type of vpc endpoint.
        """
        return pulumi.get(self, "endpoint_type")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The Id of vpc endpoint.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="serviceId")
    def service_id(self) -> str:
        """
        The Id of vpc endpoint service.
        """
        return pulumi.get(self, "service_id")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> str:
        """
        The name of vpc endpoint service.
        """
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        The status of vpc endpoint. Valid values: `Creating`, `Pending`, `Available`, `Deleting`, `Inactive`.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        The update time of vpc endpoint.
        """
        return pulumi.get(self, "update_time")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> str:
        """
        The vpc id of vpc endpoint.
        """
        return pulumi.get(self, "vpc_id")


