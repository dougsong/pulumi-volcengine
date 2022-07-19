# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'NetworkInterfacesResult',
    'AwaitableNetworkInterfacesResult',
    'network_interfaces',
    'network_interfaces_output',
]

@pulumi.output_type
class NetworkInterfacesResult:
    """
    A collection of values returned by NetworkInterfaces.
    """
    def __init__(__self__, id=None, ids=None, instance_id=None, network_interface_name=None, network_interfaces=None, output_file=None, primary_ip_addresses=None, security_group_id=None, status=None, subnet_id=None, total_count=None, type=None, vpc_id=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if network_interface_name and not isinstance(network_interface_name, str):
            raise TypeError("Expected argument 'network_interface_name' to be a str")
        pulumi.set(__self__, "network_interface_name", network_interface_name)
        if network_interfaces and not isinstance(network_interfaces, list):
            raise TypeError("Expected argument 'network_interfaces' to be a list")
        pulumi.set(__self__, "network_interfaces", network_interfaces)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if primary_ip_addresses and not isinstance(primary_ip_addresses, list):
            raise TypeError("Expected argument 'primary_ip_addresses' to be a list")
        pulumi.set(__self__, "primary_ip_addresses", primary_ip_addresses)
        if security_group_id and not isinstance(security_group_id, str):
            raise TypeError("Expected argument 'security_group_id' to be a str")
        pulumi.set(__self__, "security_group_id", security_group_id)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if subnet_id and not isinstance(subnet_id, str):
            raise TypeError("Expected argument 'subnet_id' to be a str")
        pulumi.set(__self__, "subnet_id", subnet_id)
        if total_count and not isinstance(total_count, int):
            raise TypeError("Expected argument 'total_count' to be a int")
        pulumi.set(__self__, "total_count", total_count)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if vpc_id and not isinstance(vpc_id, str):
            raise TypeError("Expected argument 'vpc_id' to be a str")
        pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ids(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "ids")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[str]:
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="networkInterfaceName")
    def network_interface_name(self) -> Optional[str]:
        """
        The name of the ENI.
        """
        return pulumi.get(self, "network_interface_name")

    @property
    @pulumi.getter(name="networkInterfaces")
    def network_interfaces(self) -> Sequence['outputs.NetworkInterfacesNetworkInterfaceResult']:
        """
        The collection of ENI.
        """
        return pulumi.get(self, "network_interfaces")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="primaryIpAddresses")
    def primary_ip_addresses(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "primary_ip_addresses")

    @property
    @pulumi.getter(name="securityGroupId")
    def security_group_id(self) -> Optional[str]:
        return pulumi.get(self, "security_group_id")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        The status of the ENI.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> Optional[str]:
        """
        The id of the subnet to which the ENI is connected.
        """
        return pulumi.get(self, "subnet_id")

    @property
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        """
        The total count of ENI query.
        """
        return pulumi.get(self, "total_count")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        """
        The type of the ENI.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[str]:
        """
        The id of the virtual private cloud (VPC) to which the ENI belongs.
        """
        return pulumi.get(self, "vpc_id")


class AwaitableNetworkInterfacesResult(NetworkInterfacesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return NetworkInterfacesResult(
            id=self.id,
            ids=self.ids,
            instance_id=self.instance_id,
            network_interface_name=self.network_interface_name,
            network_interfaces=self.network_interfaces,
            output_file=self.output_file,
            primary_ip_addresses=self.primary_ip_addresses,
            security_group_id=self.security_group_id,
            status=self.status,
            subnet_id=self.subnet_id,
            total_count=self.total_count,
            type=self.type,
            vpc_id=self.vpc_id)


def network_interfaces(ids: Optional[Sequence[str]] = None,
                       instance_id: Optional[str] = None,
                       network_interface_name: Optional[str] = None,
                       output_file: Optional[str] = None,
                       primary_ip_addresses: Optional[Sequence[str]] = None,
                       security_group_id: Optional[str] = None,
                       status: Optional[str] = None,
                       subnet_id: Optional[str] = None,
                       type: Optional[str] = None,
                       vpc_id: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableNetworkInterfacesResult:
    """
    Use this data source to query detailed information of network interfaces
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.Vpc.network_interfaces(ids=["eni-2744htx2w0j5s7fap8t3ivwze"])
    ```


    :param Sequence[str] ids: A list of ENI ids.
    :param str instance_id: An id of the instance to which the ENI is bound.
    :param str network_interface_name: A name of ENI.
    :param str output_file: File name where to save data source results.
    :param Sequence[str] primary_ip_addresses: A list of primary IP address of ENI.
    :param str security_group_id: An id of the security group to which the secondary ENI belongs.
    :param str status: A status of ENI.
    :param str subnet_id: An id of the subnet to which the ENI is connected.
    :param str type: A type of ENI.
    :param str vpc_id: An id of the virtual private cloud (VPC) to which the ENI belongs.
    """
    __args__ = dict()
    __args__['ids'] = ids
    __args__['instanceId'] = instance_id
    __args__['networkInterfaceName'] = network_interface_name
    __args__['outputFile'] = output_file
    __args__['primaryIpAddresses'] = primary_ip_addresses
    __args__['securityGroupId'] = security_group_id
    __args__['status'] = status
    __args__['subnetId'] = subnet_id
    __args__['type'] = type
    __args__['vpcId'] = vpc_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('volcengine:Vpc/networkInterfaces:NetworkInterfaces', __args__, opts=opts, typ=NetworkInterfacesResult).value

    return AwaitableNetworkInterfacesResult(
        id=__ret__.id,
        ids=__ret__.ids,
        instance_id=__ret__.instance_id,
        network_interface_name=__ret__.network_interface_name,
        network_interfaces=__ret__.network_interfaces,
        output_file=__ret__.output_file,
        primary_ip_addresses=__ret__.primary_ip_addresses,
        security_group_id=__ret__.security_group_id,
        status=__ret__.status,
        subnet_id=__ret__.subnet_id,
        total_count=__ret__.total_count,
        type=__ret__.type,
        vpc_id=__ret__.vpc_id)


@_utilities.lift_output_func(network_interfaces)
def network_interfaces_output(ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                              instance_id: Optional[pulumi.Input[Optional[str]]] = None,
                              network_interface_name: Optional[pulumi.Input[Optional[str]]] = None,
                              output_file: Optional[pulumi.Input[Optional[str]]] = None,
                              primary_ip_addresses: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                              security_group_id: Optional[pulumi.Input[Optional[str]]] = None,
                              status: Optional[pulumi.Input[Optional[str]]] = None,
                              subnet_id: Optional[pulumi.Input[Optional[str]]] = None,
                              type: Optional[pulumi.Input[Optional[str]]] = None,
                              vpc_id: Optional[pulumi.Input[Optional[str]]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[NetworkInterfacesResult]:
    """
    Use this data source to query detailed information of network interfaces
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.Vpc.network_interfaces(ids=["eni-2744htx2w0j5s7fap8t3ivwze"])
    ```


    :param Sequence[str] ids: A list of ENI ids.
    :param str instance_id: An id of the instance to which the ENI is bound.
    :param str network_interface_name: A name of ENI.
    :param str output_file: File name where to save data source results.
    :param Sequence[str] primary_ip_addresses: A list of primary IP address of ENI.
    :param str security_group_id: An id of the security group to which the secondary ENI belongs.
    :param str status: A status of ENI.
    :param str subnet_id: An id of the subnet to which the ENI is connected.
    :param str type: A type of ENI.
    :param str vpc_id: An id of the virtual private cloud (VPC) to which the ENI belongs.
    """
    ...
