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
    'RouteEntriesResult',
    'AwaitableRouteEntriesResult',
    'route_entries',
    'route_entries_output',
]

@pulumi.output_type
class RouteEntriesResult:
    """
    A collection of values returned by RouteEntries.
    """
    def __init__(__self__, destination_cidr_block=None, id=None, ids=None, next_hop_id=None, next_hop_type=None, output_file=None, route_entries=None, route_entry_name=None, route_entry_type=None, route_table_id=None, total_count=None):
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
        if next_hop_type and not isinstance(next_hop_type, str):
            raise TypeError("Expected argument 'next_hop_type' to be a str")
        pulumi.set(__self__, "next_hop_type", next_hop_type)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if route_entries and not isinstance(route_entries, list):
            raise TypeError("Expected argument 'route_entries' to be a list")
        pulumi.set(__self__, "route_entries", route_entries)
        if route_entry_name and not isinstance(route_entry_name, str):
            raise TypeError("Expected argument 'route_entry_name' to be a str")
        pulumi.set(__self__, "route_entry_name", route_entry_name)
        if route_entry_type and not isinstance(route_entry_type, str):
            raise TypeError("Expected argument 'route_entry_type' to be a str")
        pulumi.set(__self__, "route_entry_type", route_entry_type)
        if route_table_id and not isinstance(route_table_id, str):
            raise TypeError("Expected argument 'route_table_id' to be a str")
        pulumi.set(__self__, "route_table_id", route_table_id)
        if total_count and not isinstance(total_count, int):
            raise TypeError("Expected argument 'total_count' to be a int")
        pulumi.set(__self__, "total_count", total_count)

    @property
    @pulumi.getter(name="destinationCidrBlock")
    def destination_cidr_block(self) -> Optional[str]:
        """
        The destination CIDR block of the route entry.
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
        The id of the next hop.
        """
        return pulumi.get(self, "next_hop_id")

    @property
    @pulumi.getter(name="nextHopType")
    def next_hop_type(self) -> Optional[str]:
        """
        The type of the next hop.
        """
        return pulumi.get(self, "next_hop_type")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="routeEntries")
    def route_entries(self) -> Sequence['outputs.RouteEntriesRouteEntryResult']:
        """
        The collection of route tables.
        """
        return pulumi.get(self, "route_entries")

    @property
    @pulumi.getter(name="routeEntryName")
    def route_entry_name(self) -> Optional[str]:
        """
        The name of the route entry.
        """
        return pulumi.get(self, "route_entry_name")

    @property
    @pulumi.getter(name="routeEntryType")
    def route_entry_type(self) -> Optional[str]:
        return pulumi.get(self, "route_entry_type")

    @property
    @pulumi.getter(name="routeTableId")
    def route_table_id(self) -> str:
        """
        The id of the route table to which the route entry belongs.
        """
        return pulumi.get(self, "route_table_id")

    @property
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        """
        The total count of route entry query.
        """
        return pulumi.get(self, "total_count")


class AwaitableRouteEntriesResult(RouteEntriesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return RouteEntriesResult(
            destination_cidr_block=self.destination_cidr_block,
            id=self.id,
            ids=self.ids,
            next_hop_id=self.next_hop_id,
            next_hop_type=self.next_hop_type,
            output_file=self.output_file,
            route_entries=self.route_entries,
            route_entry_name=self.route_entry_name,
            route_entry_type=self.route_entry_type,
            route_table_id=self.route_table_id,
            total_count=self.total_count)


def route_entries(destination_cidr_block: Optional[str] = None,
                  ids: Optional[Sequence[str]] = None,
                  next_hop_id: Optional[str] = None,
                  next_hop_type: Optional[str] = None,
                  output_file: Optional[str] = None,
                  route_entry_name: Optional[str] = None,
                  route_entry_type: Optional[str] = None,
                  route_table_id: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableRouteEntriesResult:
    """
    Use this data source to query detailed information of route entries
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.vpc.route_entries(ids=[],
        route_table_id="vtb-274e19skkuhog7fap8u4i8ird")
    ```


    :param str destination_cidr_block: A destination CIDR block of route entry.
    :param Sequence[str] ids: A list of route entry ids.
    :param str next_hop_id: An id of next hop.
    :param str next_hop_type: A type of next hop, Optional choice contains `Instance`, `NetworkInterface`, `NatGW`, `VpnGW`.
    :param str output_file: File name where to save data source results.
    :param str route_entry_name: A name of route entry.
    :param str route_entry_type: A type of route entry.
    :param str route_table_id: An id of route table.
    """
    __args__ = dict()
    __args__['destinationCidrBlock'] = destination_cidr_block
    __args__['ids'] = ids
    __args__['nextHopId'] = next_hop_id
    __args__['nextHopType'] = next_hop_type
    __args__['outputFile'] = output_file
    __args__['routeEntryName'] = route_entry_name
    __args__['routeEntryType'] = route_entry_type
    __args__['routeTableId'] = route_table_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('volcengine:vpc/routeEntries:RouteEntries', __args__, opts=opts, typ=RouteEntriesResult).value

    return AwaitableRouteEntriesResult(
        destination_cidr_block=__ret__.destination_cidr_block,
        id=__ret__.id,
        ids=__ret__.ids,
        next_hop_id=__ret__.next_hop_id,
        next_hop_type=__ret__.next_hop_type,
        output_file=__ret__.output_file,
        route_entries=__ret__.route_entries,
        route_entry_name=__ret__.route_entry_name,
        route_entry_type=__ret__.route_entry_type,
        route_table_id=__ret__.route_table_id,
        total_count=__ret__.total_count)


@_utilities.lift_output_func(route_entries)
def route_entries_output(destination_cidr_block: Optional[pulumi.Input[Optional[str]]] = None,
                         ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                         next_hop_id: Optional[pulumi.Input[Optional[str]]] = None,
                         next_hop_type: Optional[pulumi.Input[Optional[str]]] = None,
                         output_file: Optional[pulumi.Input[Optional[str]]] = None,
                         route_entry_name: Optional[pulumi.Input[Optional[str]]] = None,
                         route_entry_type: Optional[pulumi.Input[Optional[str]]] = None,
                         route_table_id: Optional[pulumi.Input[str]] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[RouteEntriesResult]:
    """
    Use this data source to query detailed information of route entries
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.vpc.route_entries(ids=[],
        route_table_id="vtb-274e19skkuhog7fap8u4i8ird")
    ```


    :param str destination_cidr_block: A destination CIDR block of route entry.
    :param Sequence[str] ids: A list of route entry ids.
    :param str next_hop_id: An id of next hop.
    :param str next_hop_type: A type of next hop, Optional choice contains `Instance`, `NetworkInterface`, `NatGW`, `VpnGW`.
    :param str output_file: File name where to save data source results.
    :param str route_entry_name: A name of route entry.
    :param str route_entry_type: A type of route entry.
    :param str route_table_id: An id of route table.
    """
    ...
