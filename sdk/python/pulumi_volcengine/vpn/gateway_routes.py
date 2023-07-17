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
    'GatewayRoutesResult',
    'AwaitableGatewayRoutesResult',
    'gateway_routes',
    'gateway_routes_output',
]

@pulumi.output_type
class GatewayRoutesResult:
    """
    A collection of values returned by GatewayRoutes.
    """
    def __init__(__self__, destination_cidr_block=None, id=None, ids=None, next_hop_id=None, output_file=None, total_count=None, vpn_gateway_id=None, vpn_gateway_routes=None):
        if destination_cidr_block and not isinstance(destination_cidr_block, str):
            raise TypeError("Expected argument 'destination_cidr_block' to be a str")
        pulumi.set(__self__, "destination_cidr_block", destination_cidr_block)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if next_hop_id and not isinstance(next_hop_id, str):
            raise TypeError("Expected argument 'next_hop_id' to be a str")
        pulumi.set(__self__, "next_hop_id", next_hop_id)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if total_count and not isinstance(total_count, int):
            raise TypeError("Expected argument 'total_count' to be a int")
        pulumi.set(__self__, "total_count", total_count)
        if vpn_gateway_id and not isinstance(vpn_gateway_id, str):
            raise TypeError("Expected argument 'vpn_gateway_id' to be a str")
        pulumi.set(__self__, "vpn_gateway_id", vpn_gateway_id)
        if vpn_gateway_routes and not isinstance(vpn_gateway_routes, list):
            raise TypeError("Expected argument 'vpn_gateway_routes' to be a list")
        pulumi.set(__self__, "vpn_gateway_routes", vpn_gateway_routes)

    @property
    @pulumi.getter(name="destinationCidrBlock")
    def destination_cidr_block(self) -> Optional[str]:
        """
        The destination cidr block of the VPN gateway route.
        """
        return pulumi.get(self, "destination_cidr_block")

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
    @pulumi.getter(name="nextHopId")
    def next_hop_id(self) -> Optional[str]:
        """
        The next hop id of the VPN gateway route.
        """
        return pulumi.get(self, "next_hop_id")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        """
        The total count of VPN gateway route query.
        """
        return pulumi.get(self, "total_count")

    @property
    @pulumi.getter(name="vpnGatewayId")
    def vpn_gateway_id(self) -> Optional[str]:
        """
        The ID of the VPN gateway of the VPN gateway route.
        """
        return pulumi.get(self, "vpn_gateway_id")

    @property
    @pulumi.getter(name="vpnGatewayRoutes")
    def vpn_gateway_routes(self) -> Sequence['outputs.GatewayRoutesVpnGatewayRouteResult']:
        """
        The collection of VPN gateway route query.
        """
        return pulumi.get(self, "vpn_gateway_routes")


class AwaitableGatewayRoutesResult(GatewayRoutesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GatewayRoutesResult(
            destination_cidr_block=self.destination_cidr_block,
            id=self.id,
            ids=self.ids,
            next_hop_id=self.next_hop_id,
            output_file=self.output_file,
            total_count=self.total_count,
            vpn_gateway_id=self.vpn_gateway_id,
            vpn_gateway_routes=self.vpn_gateway_routes)


def gateway_routes(destination_cidr_block: Optional[str] = None,
                   ids: Optional[Sequence[str]] = None,
                   next_hop_id: Optional[str] = None,
                   output_file: Optional[str] = None,
                   vpn_gateway_id: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGatewayRoutesResult:
    """
    Use this data source to query detailed information of vpn gateway routes
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.vpn.gateway_routes(ids=["vgr-2byssu52dktts2dx0ee90r5hp]"])
    ```


    :param str destination_cidr_block: A destination cidr block.
    :param Sequence[str] ids: A list of VPN gateway route ids.
    :param str next_hop_id: An ID of next hop.
    :param str output_file: File name where to save data source results.
    :param str vpn_gateway_id: An ID of VPN gateway.
    """
    __args__ = dict()
    __args__['destinationCidrBlock'] = destination_cidr_block
    __args__['ids'] = ids
    __args__['nextHopId'] = next_hop_id
    __args__['outputFile'] = output_file
    __args__['vpnGatewayId'] = vpn_gateway_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('volcengine:vpn/gatewayRoutes:GatewayRoutes', __args__, opts=opts, typ=GatewayRoutesResult).value

    return AwaitableGatewayRoutesResult(
        destination_cidr_block=__ret__.destination_cidr_block,
        id=__ret__.id,
        ids=__ret__.ids,
        next_hop_id=__ret__.next_hop_id,
        output_file=__ret__.output_file,
        total_count=__ret__.total_count,
        vpn_gateway_id=__ret__.vpn_gateway_id,
        vpn_gateway_routes=__ret__.vpn_gateway_routes)


@_utilities.lift_output_func(gateway_routes)
def gateway_routes_output(destination_cidr_block: Optional[pulumi.Input[Optional[str]]] = None,
                          ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                          next_hop_id: Optional[pulumi.Input[Optional[str]]] = None,
                          output_file: Optional[pulumi.Input[Optional[str]]] = None,
                          vpn_gateway_id: Optional[pulumi.Input[Optional[str]]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GatewayRoutesResult]:
    """
    Use this data source to query detailed information of vpn gateway routes
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.vpn.gateway_routes(ids=["vgr-2byssu52dktts2dx0ee90r5hp]"])
    ```


    :param str destination_cidr_block: A destination cidr block.
    :param Sequence[str] ids: A list of VPN gateway route ids.
    :param str next_hop_id: An ID of next hop.
    :param str output_file: File name where to save data source results.
    :param str vpn_gateway_id: An ID of VPN gateway.
    """
    ...
