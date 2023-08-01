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
    'SpecsResult',
    'AwaitableSpecsResult',
    'specs',
    'specs_output',
]

@pulumi.output_type
class SpecsResult:
    """
    A collection of values returned by Specs.
    """
    def __init__(__self__, id=None, output_file=None, region_id=None, specs=None, total_count=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if region_id and not isinstance(region_id, str):
            raise TypeError("Expected argument 'region_id' to be a str")
        pulumi.set(__self__, "region_id", region_id)
        if specs and not isinstance(specs, dict):
            raise TypeError("Expected argument 'specs' to be a dict")
        pulumi.set(__self__, "specs", specs)
        if total_count and not isinstance(total_count, int):
            raise TypeError("Expected argument 'total_count' to be a int")
        pulumi.set(__self__, "total_count", total_count)

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
    def region_id(self) -> Optional[str]:
        return pulumi.get(self, "region_id")

    @property
    @pulumi.getter
    def specs(self) -> 'outputs.SpecsSpecsResult':
        """
        A list of supported node specification information for MongoDB instances.
        """
        return pulumi.get(self, "specs")

    @property
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        """
        The total count of region query.
        """
        return pulumi.get(self, "total_count")


class AwaitableSpecsResult(SpecsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return SpecsResult(
            id=self.id,
            output_file=self.output_file,
            region_id=self.region_id,
            specs=self.specs,
            total_count=self.total_count)


def specs(output_file: Optional[str] = None,
          region_id: Optional[str] = None,
          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableSpecsResult:
    """
    Use this data source to query detailed information of mongodb specs
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    foo = volcengine.mongodb.specs()
    ```


    :param str output_file: File name where to save data source results.
    :param str region_id: The region ID to query.
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
    __ret__ = pulumi.runtime.invoke('volcengine:mongodb/specs:Specs', __args__, opts=opts, typ=SpecsResult).value

    return AwaitableSpecsResult(
        id=__ret__.id,
        output_file=__ret__.output_file,
        region_id=__ret__.region_id,
        specs=__ret__.specs,
        total_count=__ret__.total_count)


@_utilities.lift_output_func(specs)
def specs_output(output_file: Optional[pulumi.Input[Optional[str]]] = None,
                 region_id: Optional[pulumi.Input[Optional[str]]] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[SpecsResult]:
    """
    Use this data source to query detailed information of mongodb specs
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    foo = volcengine.mongodb.specs()
    ```


    :param str output_file: File name where to save data source results.
    :param str region_id: The region ID to query.
    """
    ...
