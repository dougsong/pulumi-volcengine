# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = [
    'BandwidthPackagesResult',
    'AwaitableBandwidthPackagesResult',
    'bandwidth_packages',
    'bandwidth_packages_output',
]

@pulumi.output_type
class BandwidthPackagesResult:
    """
    A collection of values returned by BandwidthPackages.
    """
    def __init__(__self__, bandwidth_packages=None, cen_bandwidth_package_names=None, cen_id=None, id=None, ids=None, local_geographic_region_set_id=None, name_regex=None, output_file=None, peer_geographic_region_set_id=None, tags=None, total_count=None):
        if bandwidth_packages and not isinstance(bandwidth_packages, list):
            raise TypeError("Expected argument 'bandwidth_packages' to be a list")
        pulumi.set(__self__, "bandwidth_packages", bandwidth_packages)
        if cen_bandwidth_package_names and not isinstance(cen_bandwidth_package_names, list):
            raise TypeError("Expected argument 'cen_bandwidth_package_names' to be a list")
        pulumi.set(__self__, "cen_bandwidth_package_names", cen_bandwidth_package_names)
        if cen_id and not isinstance(cen_id, str):
            raise TypeError("Expected argument 'cen_id' to be a str")
        pulumi.set(__self__, "cen_id", cen_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if local_geographic_region_set_id and not isinstance(local_geographic_region_set_id, str):
            raise TypeError("Expected argument 'local_geographic_region_set_id' to be a str")
        pulumi.set(__self__, "local_geographic_region_set_id", local_geographic_region_set_id)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if peer_geographic_region_set_id and not isinstance(peer_geographic_region_set_id, str):
            raise TypeError("Expected argument 'peer_geographic_region_set_id' to be a str")
        pulumi.set(__self__, "peer_geographic_region_set_id", peer_geographic_region_set_id)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if total_count and not isinstance(total_count, int):
            raise TypeError("Expected argument 'total_count' to be a int")
        pulumi.set(__self__, "total_count", total_count)

    @property
    @pulumi.getter(name="bandwidthPackages")
    def bandwidth_packages(self) -> Sequence['outputs.BandwidthPackagesBandwidthPackageResult']:
        """
        The collection of cen bandwidth package query.
        """
        return pulumi.get(self, "bandwidth_packages")

    @property
    @pulumi.getter(name="cenBandwidthPackageNames")
    def cen_bandwidth_package_names(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "cen_bandwidth_package_names")

    @property
    @pulumi.getter(name="cenId")
    def cen_id(self) -> Optional[str]:
        return pulumi.get(self, "cen_id")

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
    @pulumi.getter(name="localGeographicRegionSetId")
    def local_geographic_region_set_id(self) -> Optional[str]:
        """
        The local geographic region set id of the cen bandwidth package.
        """
        return pulumi.get(self, "local_geographic_region_set_id")

    @property
    @pulumi.getter(name="nameRegex")
    def name_regex(self) -> Optional[str]:
        return pulumi.get(self, "name_regex")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="peerGeographicRegionSetId")
    def peer_geographic_region_set_id(self) -> Optional[str]:
        """
        The peer geographic region set id of the cen bandwidth package.
        """
        return pulumi.get(self, "peer_geographic_region_set_id")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.BandwidthPackagesTagResult']]:
        """
        Tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        """
        The total count of cen bandwidth package query.
        """
        return pulumi.get(self, "total_count")


class AwaitableBandwidthPackagesResult(BandwidthPackagesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return BandwidthPackagesResult(
            bandwidth_packages=self.bandwidth_packages,
            cen_bandwidth_package_names=self.cen_bandwidth_package_names,
            cen_id=self.cen_id,
            id=self.id,
            ids=self.ids,
            local_geographic_region_set_id=self.local_geographic_region_set_id,
            name_regex=self.name_regex,
            output_file=self.output_file,
            peer_geographic_region_set_id=self.peer_geographic_region_set_id,
            tags=self.tags,
            total_count=self.total_count)


def bandwidth_packages(cen_bandwidth_package_names: Optional[Sequence[str]] = None,
                       cen_id: Optional[str] = None,
                       ids: Optional[Sequence[str]] = None,
                       local_geographic_region_set_id: Optional[str] = None,
                       name_regex: Optional[str] = None,
                       output_file: Optional[str] = None,
                       peer_geographic_region_set_id: Optional[str] = None,
                       tags: Optional[Sequence[pulumi.InputType['BandwidthPackagesTagArgs']]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableBandwidthPackagesResult:
    """
    Use this data source to query detailed information of cen bandwidth packages
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    foo = volcengine.cen.bandwidth_packages(cen_id="cen-2bzrl3srxsv0g2dx0efyoojn3",
        ids=["cbp-2bzeew3s8p79c2dx0eeohej4x"])
    ```


    :param Sequence[str] cen_bandwidth_package_names: A list of cen bandwidth package names.
    :param str cen_id: A cen id.
    :param Sequence[str] ids: A list of cen bandwidth package IDs.
    :param str local_geographic_region_set_id: A local geographic region set id.
    :param str name_regex: A Name Regex of cen bandwidth package.
    :param str output_file: File name where to save data source results.
    :param str peer_geographic_region_set_id: A peer geographic region set id.
    :param Sequence[pulumi.InputType['BandwidthPackagesTagArgs']] tags: Tags.
    """
    __args__ = dict()
    __args__['cenBandwidthPackageNames'] = cen_bandwidth_package_names
    __args__['cenId'] = cen_id
    __args__['ids'] = ids
    __args__['localGeographicRegionSetId'] = local_geographic_region_set_id
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['peerGeographicRegionSetId'] = peer_geographic_region_set_id
    __args__['tags'] = tags
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('volcengine:cen/bandwidthPackages:BandwidthPackages', __args__, opts=opts, typ=BandwidthPackagesResult).value

    return AwaitableBandwidthPackagesResult(
        bandwidth_packages=__ret__.bandwidth_packages,
        cen_bandwidth_package_names=__ret__.cen_bandwidth_package_names,
        cen_id=__ret__.cen_id,
        id=__ret__.id,
        ids=__ret__.ids,
        local_geographic_region_set_id=__ret__.local_geographic_region_set_id,
        name_regex=__ret__.name_regex,
        output_file=__ret__.output_file,
        peer_geographic_region_set_id=__ret__.peer_geographic_region_set_id,
        tags=__ret__.tags,
        total_count=__ret__.total_count)


@_utilities.lift_output_func(bandwidth_packages)
def bandwidth_packages_output(cen_bandwidth_package_names: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                              cen_id: Optional[pulumi.Input[Optional[str]]] = None,
                              ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                              local_geographic_region_set_id: Optional[pulumi.Input[Optional[str]]] = None,
                              name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                              output_file: Optional[pulumi.Input[Optional[str]]] = None,
                              peer_geographic_region_set_id: Optional[pulumi.Input[Optional[str]]] = None,
                              tags: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['BandwidthPackagesTagArgs']]]]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[BandwidthPackagesResult]:
    """
    Use this data source to query detailed information of cen bandwidth packages
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    foo = volcengine.cen.bandwidth_packages(cen_id="cen-2bzrl3srxsv0g2dx0efyoojn3",
        ids=["cbp-2bzeew3s8p79c2dx0eeohej4x"])
    ```


    :param Sequence[str] cen_bandwidth_package_names: A list of cen bandwidth package names.
    :param str cen_id: A cen id.
    :param Sequence[str] ids: A list of cen bandwidth package IDs.
    :param str local_geographic_region_set_id: A local geographic region set id.
    :param str name_regex: A Name Regex of cen bandwidth package.
    :param str output_file: File name where to save data source results.
    :param str peer_geographic_region_set_id: A peer geographic region set id.
    :param Sequence[pulumi.InputType['BandwidthPackagesTagArgs']] tags: Tags.
    """
    ...
