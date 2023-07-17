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
    'InstanceTypesResult',
    'AwaitableInstanceTypesResult',
    'instance_types',
    'instance_types_output',
]

@pulumi.output_type
class InstanceTypesResult:
    """
    A collection of values returned by InstanceTypes.
    """
    def __init__(__self__, id=None, instance_type_configs=None, output_file=None, total_count=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_type_configs and not isinstance(instance_type_configs, list):
            raise TypeError("Expected argument 'instance_type_configs' to be a list")
        pulumi.set(__self__, "instance_type_configs", instance_type_configs)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
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
    @pulumi.getter(name="instanceTypeConfigs")
    def instance_type_configs(self) -> Sequence['outputs.InstanceTypesInstanceTypeConfigResult']:
        """
        The collection of instance types query.
        """
        return pulumi.get(self, "instance_type_configs")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        """
        The total count of instance types query.
        """
        return pulumi.get(self, "total_count")


class AwaitableInstanceTypesResult(InstanceTypesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return InstanceTypesResult(
            id=self.id,
            instance_type_configs=self.instance_type_configs,
            output_file=self.output_file,
            total_count=self.total_count)


def instance_types(output_file: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableInstanceTypesResult:
    """
    Use this data source to query detailed information of veenedge instance types
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.veenedge.instance_types()
    ```


    :param str output_file: File name where to save data source results.
    """
    __args__ = dict()
    __args__['outputFile'] = output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('volcengine:veenedge/instanceTypes:InstanceTypes', __args__, opts=opts, typ=InstanceTypesResult).value

    return AwaitableInstanceTypesResult(
        id=__ret__.id,
        instance_type_configs=__ret__.instance_type_configs,
        output_file=__ret__.output_file,
        total_count=__ret__.total_count)


@_utilities.lift_output_func(instance_types)
def instance_types_output(output_file: Optional[pulumi.Input[Optional[str]]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[InstanceTypesResult]:
    """
    Use this data source to query detailed information of veenedge instance types
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.veenedge.instance_types()
    ```


    :param str output_file: File name where to save data source results.
    """
    ...
