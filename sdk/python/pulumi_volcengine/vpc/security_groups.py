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
    'SecurityGroupsResult',
    'AwaitableSecurityGroupsResult',
    'security_groups',
    'security_groups_output',
]

@pulumi.output_type
class SecurityGroupsResult:
    """
    A collection of values returned by SecurityGroups.
    """
    def __init__(__self__, id=None, ids=None, name_regex=None, output_file=None, security_groups=None, total_count=None):
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
        if security_groups and not isinstance(security_groups, list):
            raise TypeError("Expected argument 'security_groups' to be a list")
        pulumi.set(__self__, "security_groups", security_groups)
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
    @pulumi.getter(name="securityGroups")
    def security_groups(self) -> Sequence['outputs.SecurityGroupsSecurityGroupResult']:
        """
        The collection of SecurityGroup query.
        """
        return pulumi.get(self, "security_groups")

    @property
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        """
        The total count of SecurityGroup query.
        """
        return pulumi.get(self, "total_count")


class AwaitableSecurityGroupsResult(SecurityGroupsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return SecurityGroupsResult(
            id=self.id,
            ids=self.ids,
            name_regex=self.name_regex,
            output_file=self.output_file,
            security_groups=self.security_groups,
            total_count=self.total_count)


def security_groups(ids: Optional[Sequence[str]] = None,
                    name_regex: Optional[str] = None,
                    output_file: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableSecurityGroupsResult:
    """
    Use this data source to query detailed information of security groups
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.Vpc.security_groups(ids=["sg-273ycgql3ig3k7fap8t3dyvqx"])
    ```


    :param Sequence[str] ids: A list of SecurityGroup IDs.
    :param str name_regex: A Name Regex of SecurityGroup.
    :param str output_file: File name where to save data source results.
    """
    __args__ = dict()
    __args__['ids'] = ids
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('volcengine:Vpc/securityGroups:SecurityGroups', __args__, opts=opts, typ=SecurityGroupsResult).value

    return AwaitableSecurityGroupsResult(
        id=__ret__.id,
        ids=__ret__.ids,
        name_regex=__ret__.name_regex,
        output_file=__ret__.output_file,
        security_groups=__ret__.security_groups,
        total_count=__ret__.total_count)


@_utilities.lift_output_func(security_groups)
def security_groups_output(ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                           name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                           output_file: Optional[pulumi.Input[Optional[str]]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[SecurityGroupsResult]:
    """
    Use this data source to query detailed information of security groups
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.Vpc.security_groups(ids=["sg-273ycgql3ig3k7fap8t3dyvqx"])
    ```


    :param Sequence[str] ids: A list of SecurityGroup IDs.
    :param str name_regex: A Name Regex of SecurityGroup.
    :param str output_file: File name where to save data source results.
    """
    ...
