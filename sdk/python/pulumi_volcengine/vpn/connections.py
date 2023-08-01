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
    'ConnectionsResult',
    'AwaitableConnectionsResult',
    'connections',
    'connections_output',
]

@pulumi.output_type
class ConnectionsResult:
    """
    A collection of values returned by Connections.
    """
    def __init__(__self__, customer_gateway_id=None, id=None, ids=None, name_regex=None, output_file=None, total_count=None, vpn_connection_names=None, vpn_connections=None, vpn_gateway_id=None):
        if customer_gateway_id and not isinstance(customer_gateway_id, str):
            raise TypeError("Expected argument 'customer_gateway_id' to be a str")
        pulumi.set(__self__, "customer_gateway_id", customer_gateway_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if total_count and not isinstance(total_count, int):
            raise TypeError("Expected argument 'total_count' to be a int")
        pulumi.set(__self__, "total_count", total_count)
        if vpn_connection_names and not isinstance(vpn_connection_names, list):
            raise TypeError("Expected argument 'vpn_connection_names' to be a list")
        pulumi.set(__self__, "vpn_connection_names", vpn_connection_names)
        if vpn_connections and not isinstance(vpn_connections, list):
            raise TypeError("Expected argument 'vpn_connections' to be a list")
        pulumi.set(__self__, "vpn_connections", vpn_connections)
        if vpn_gateway_id and not isinstance(vpn_gateway_id, str):
            raise TypeError("Expected argument 'vpn_gateway_id' to be a str")
        pulumi.set(__self__, "vpn_gateway_id", vpn_gateway_id)

    @property
    @pulumi.getter(name="customerGatewayId")
    def customer_gateway_id(self) -> Optional[str]:
        """
        The ID of the customer gateway.
        """
        return pulumi.get(self, "customer_gateway_id")

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
    @pulumi.getter(name="nameRegex")
    def name_regex(self) -> Optional[str]:
        return pulumi.get(self, "name_regex")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        """
        The total count of VPN connection query.
        """
        return pulumi.get(self, "total_count")

    @property
    @pulumi.getter(name="vpnConnectionNames")
    def vpn_connection_names(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "vpn_connection_names")

    @property
    @pulumi.getter(name="vpnConnections")
    def vpn_connections(self) -> Sequence['outputs.ConnectionsVpnConnectionResult']:
        """
        The collection of VPN connection query.
        """
        return pulumi.get(self, "vpn_connections")

    @property
    @pulumi.getter(name="vpnGatewayId")
    def vpn_gateway_id(self) -> Optional[str]:
        """
        The ID of the vpn gateway.
        """
        return pulumi.get(self, "vpn_gateway_id")


class AwaitableConnectionsResult(ConnectionsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ConnectionsResult(
            customer_gateway_id=self.customer_gateway_id,
            id=self.id,
            ids=self.ids,
            name_regex=self.name_regex,
            output_file=self.output_file,
            total_count=self.total_count,
            vpn_connection_names=self.vpn_connection_names,
            vpn_connections=self.vpn_connections,
            vpn_gateway_id=self.vpn_gateway_id)


def connections(customer_gateway_id: Optional[str] = None,
                ids: Optional[Sequence[str]] = None,
                name_regex: Optional[str] = None,
                output_file: Optional[str] = None,
                vpn_connection_names: Optional[Sequence[str]] = None,
                vpn_gateway_id: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableConnectionsResult:
    """
    Use this data source to query detailed information of vpn connections
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.vpn.connections(ids=["vgc-2d5wwids8cdts58ozfe63k2uq"])
    ```


    :param str customer_gateway_id: An ID of customer gateway.
    :param Sequence[str] ids: A list of VPN connection ids.
    :param str name_regex: A Name Regex of VPN connection.
    :param str output_file: File name where to save data source results.
    :param Sequence[str] vpn_connection_names: A list of VPN connection names.
    :param str vpn_gateway_id: An ID of VPN gateway.
    """
    __args__ = dict()
    __args__['customerGatewayId'] = customer_gateway_id
    __args__['ids'] = ids
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['vpnConnectionNames'] = vpn_connection_names
    __args__['vpnGatewayId'] = vpn_gateway_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('volcengine:vpn/connections:Connections', __args__, opts=opts, typ=ConnectionsResult).value

    return AwaitableConnectionsResult(
        customer_gateway_id=__ret__.customer_gateway_id,
        id=__ret__.id,
        ids=__ret__.ids,
        name_regex=__ret__.name_regex,
        output_file=__ret__.output_file,
        total_count=__ret__.total_count,
        vpn_connection_names=__ret__.vpn_connection_names,
        vpn_connections=__ret__.vpn_connections,
        vpn_gateway_id=__ret__.vpn_gateway_id)


@_utilities.lift_output_func(connections)
def connections_output(customer_gateway_id: Optional[pulumi.Input[Optional[str]]] = None,
                       ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                       name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                       output_file: Optional[pulumi.Input[Optional[str]]] = None,
                       vpn_connection_names: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                       vpn_gateway_id: Optional[pulumi.Input[Optional[str]]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[ConnectionsResult]:
    """
    Use this data source to query detailed information of vpn connections
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.vpn.connections(ids=["vgc-2d5wwids8cdts58ozfe63k2uq"])
    ```


    :param str customer_gateway_id: An ID of customer gateway.
    :param Sequence[str] ids: A list of VPN connection ids.
    :param str name_regex: A Name Regex of VPN connection.
    :param str output_file: File name where to save data source results.
    :param Sequence[str] vpn_connection_names: A list of VPN connection names.
    :param str vpn_gateway_id: An ID of VPN gateway.
    """
    ...
