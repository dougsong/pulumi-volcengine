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
    'AllowListsResult',
    'AwaitableAllowListsResult',
    'allow_lists',
    'allow_lists_output',
]

@pulumi.output_type
class AllowListsResult:
    """
    A collection of values returned by AllowLists.
    """
    def __init__(__self__, allow_lists=None, id=None, instance_id=None, name_regex=None, output_file=None, region_id=None, total_count=None):
        if allow_lists and not isinstance(allow_lists, list):
            raise TypeError("Expected argument 'allow_lists' to be a list")
        pulumi.set(__self__, "allow_lists", allow_lists)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if region_id and not isinstance(region_id, str):
            raise TypeError("Expected argument 'region_id' to be a str")
        pulumi.set(__self__, "region_id", region_id)
        if total_count and not isinstance(total_count, int):
            raise TypeError("Expected argument 'total_count' to be a int")
        pulumi.set(__self__, "total_count", total_count)

    @property
    @pulumi.getter(name="allowLists")
    def allow_lists(self) -> Sequence['outputs.AllowListsAllowListResult']:
        """
        Information of list of allow list.
        """
        return pulumi.get(self, "allow_lists")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[str]:
        """
        Id of instance.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="nameRegex")
    def name_regex(self) -> Optional[str]:
        return pulumi.get(self, "name_regex")

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
        The total count of allow list query.
        """
        return pulumi.get(self, "total_count")


class AwaitableAllowListsResult(AllowListsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return AllowListsResult(
            allow_lists=self.allow_lists,
            id=self.id,
            instance_id=self.instance_id,
            name_regex=self.name_regex,
            output_file=self.output_file,
            region_id=self.region_id,
            total_count=self.total_count)


def allow_lists(instance_id: Optional[str] = None,
                name_regex: Optional[str] = None,
                output_file: Optional[str] = None,
                region_id: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableAllowListsResult:
    """
    Use this data source to query detailed information of redis allow lists
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.redis.allow_lists(region_id="cn-beijing")
    ```


    :param str instance_id: The Id of instance.
    :param str name_regex: A Name Regex of Allow List.
    :param str output_file: File name where to save data source results.
    :param str region_id: The Id of region.
    """
    __args__ = dict()
    __args__['instanceId'] = instance_id
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['regionId'] = region_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('volcengine:redis/allowLists:AllowLists', __args__, opts=opts, typ=AllowListsResult).value

    return AwaitableAllowListsResult(
        allow_lists=__ret__.allow_lists,
        id=__ret__.id,
        instance_id=__ret__.instance_id,
        name_regex=__ret__.name_regex,
        output_file=__ret__.output_file,
        region_id=__ret__.region_id,
        total_count=__ret__.total_count)


@_utilities.lift_output_func(allow_lists)
def allow_lists_output(instance_id: Optional[pulumi.Input[Optional[str]]] = None,
                       name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                       output_file: Optional[pulumi.Input[Optional[str]]] = None,
                       region_id: Optional[pulumi.Input[str]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[AllowListsResult]:
    """
    Use this data source to query detailed information of redis allow lists
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.redis.allow_lists(region_id="cn-beijing")
    ```


    :param str instance_id: The Id of instance.
    :param str name_regex: A Name Regex of Allow List.
    :param str output_file: File name where to save data source results.
    :param str region_id: The Id of region.
    """
    ...
