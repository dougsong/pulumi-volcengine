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
    'ZonesResult',
    'AwaitableZonesResult',
    'zones',
    'zones_output',
]

@pulumi.output_type
class ZonesResult:
    """
    A collection of values returned by Zones.
    """
    def __init__(__self__, id=None, output_file=None, region_id=None, total_count=None, zones=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if region_id and not isinstance(region_id, str):
            raise TypeError("Expected argument 'region_id' to be a str")
        pulumi.set(__self__, "region_id", region_id)
        if total_count and not isinstance(total_count, int):
            raise TypeError("Expected argument 'total_count' to be a int")
        pulumi.set(__self__, "total_count", total_count)
        if zones and not isinstance(zones, list):
            raise TypeError("Expected argument 'zones' to be a list")
        pulumi.set(__self__, "zones", zones)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="regionId")
    def region_id(self) -> str:
        return pulumi.get(self, "region_id")

    @property
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        """
        The total count of zone query.
        """
        return pulumi.get(self, "total_count")

    @property
    @pulumi.getter
    def zones(self) -> Sequence['outputs.ZonesZoneResult']:
        """
        The collection of zone query.
        """
        return pulumi.get(self, "zones")


class AwaitableZonesResult(ZonesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ZonesResult(
            id=self.id,
            output_file=self.output_file,
            region_id=self.region_id,
            total_count=self.total_count,
            zones=self.zones)


def zones(output_file: Optional[str] = None,
          region_id: Optional[str] = None,
          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableZonesResult:
    """
    Use this data source to query detailed information of escloud zones
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.escloud.zones(region_id="xxx")
    ```


    :param str output_file: File name where to save data source results.
    :param str region_id: The Id of Region.
    """
    __args__ = dict()
    __args__['outputFile'] = output_file
    __args__['regionId'] = region_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('volcengine:escloud/zones:Zones', __args__, opts=opts, typ=ZonesResult).value

    return AwaitableZonesResult(
        id=__ret__.id,
        output_file=__ret__.output_file,
        region_id=__ret__.region_id,
        total_count=__ret__.total_count,
        zones=__ret__.zones)


@_utilities.lift_output_func(zones)
def zones_output(output_file: Optional[pulumi.Input[Optional[str]]] = None,
                 region_id: Optional[pulumi.Input[str]] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[ZonesResult]:
    """
    Use this data source to query detailed information of escloud zones
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.escloud.zones(region_id="xxx")
    ```


    :param str output_file: File name where to save data source results.
    :param str region_id: The Id of Region.
    """
    ...
