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
    'ServerGroupsResult',
    'AwaitableServerGroupsResult',
    'server_groups',
    'server_groups_output',
]

@pulumi.output_type
class ServerGroupsResult:
    """
    A collection of values returned by ServerGroups.
    """
    def __init__(__self__, groups=None, id=None, ids=None, load_balancer_id=None, name_regex=None, output_file=None, server_group_name=None, total_count=None):
        if groups and not isinstance(groups, list):
            raise TypeError("Expected argument 'groups' to be a list")
        pulumi.set(__self__, "groups", groups)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if load_balancer_id and not isinstance(load_balancer_id, str):
            raise TypeError("Expected argument 'load_balancer_id' to be a str")
        pulumi.set(__self__, "load_balancer_id", load_balancer_id)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if server_group_name and not isinstance(server_group_name, str):
            raise TypeError("Expected argument 'server_group_name' to be a str")
        pulumi.set(__self__, "server_group_name", server_group_name)
        if total_count and not isinstance(total_count, int):
            raise TypeError("Expected argument 'total_count' to be a int")
        pulumi.set(__self__, "total_count", total_count)

    @property
    @pulumi.getter
    def groups(self) -> Sequence['outputs.ServerGroupsGroupResult']:
        """
        The collection of ServerGroup query.
        """
        return pulumi.get(self, "groups")

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
    @pulumi.getter(name="loadBalancerId")
    def load_balancer_id(self) -> Optional[str]:
        return pulumi.get(self, "load_balancer_id")

    @property
    @pulumi.getter(name="nameRegex")
    def name_regex(self) -> Optional[str]:
        return pulumi.get(self, "name_regex")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="serverGroupName")
    def server_group_name(self) -> Optional[str]:
        """
        The name of the ServerGroup.
        """
        return pulumi.get(self, "server_group_name")

    @property
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        """
        The total count of ServerGroup query.
        """
        return pulumi.get(self, "total_count")


class AwaitableServerGroupsResult(ServerGroupsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ServerGroupsResult(
            groups=self.groups,
            id=self.id,
            ids=self.ids,
            load_balancer_id=self.load_balancer_id,
            name_regex=self.name_regex,
            output_file=self.output_file,
            server_group_name=self.server_group_name,
            total_count=self.total_count)


def server_groups(ids: Optional[Sequence[str]] = None,
                  load_balancer_id: Optional[str] = None,
                  name_regex: Optional[str] = None,
                  output_file: Optional[str] = None,
                  server_group_name: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableServerGroupsResult:
    """
    Use this data source to query detailed information of server groups
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.Clb.server_groups(ids=[
        "rsp-273yv0kir1vk07fap8tt9jtwg",
        "rsp-273yxuqfova4g7fap8tyemn6t",
        "rsp-273z9pt9lpdds7fap8sqdvfrf",
    ])
    ```


    :param Sequence[str] ids: A list of ServerGroup IDs.
    :param str load_balancer_id: The id of the Clb.
    :param str name_regex: A Name Regex of ServerGroup.
    :param str output_file: File name where to save data source results.
    :param str server_group_name: The name of the ServerGroup.
    """
    __args__ = dict()
    __args__['ids'] = ids
    __args__['loadBalancerId'] = load_balancer_id
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['serverGroupName'] = server_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('volcengine:Clb/serverGroups:ServerGroups', __args__, opts=opts, typ=ServerGroupsResult).value

    return AwaitableServerGroupsResult(
        groups=__ret__.groups,
        id=__ret__.id,
        ids=__ret__.ids,
        load_balancer_id=__ret__.load_balancer_id,
        name_regex=__ret__.name_regex,
        output_file=__ret__.output_file,
        server_group_name=__ret__.server_group_name,
        total_count=__ret__.total_count)


@_utilities.lift_output_func(server_groups)
def server_groups_output(ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                         load_balancer_id: Optional[pulumi.Input[Optional[str]]] = None,
                         name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                         output_file: Optional[pulumi.Input[Optional[str]]] = None,
                         server_group_name: Optional[pulumi.Input[Optional[str]]] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[ServerGroupsResult]:
    """
    Use this data source to query detailed information of server groups
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.Clb.server_groups(ids=[
        "rsp-273yv0kir1vk07fap8tt9jtwg",
        "rsp-273yxuqfova4g7fap8tyemn6t",
        "rsp-273z9pt9lpdds7fap8sqdvfrf",
    ])
    ```


    :param Sequence[str] ids: A list of ServerGroup IDs.
    :param str load_balancer_id: The id of the Clb.
    :param str name_regex: A Name Regex of ServerGroup.
    :param str output_file: File name where to save data source results.
    :param str server_group_name: The name of the ServerGroup.
    """
    ...
