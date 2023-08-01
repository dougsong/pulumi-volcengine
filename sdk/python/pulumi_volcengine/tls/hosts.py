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
    'HostsResult',
    'AwaitableHostsResult',
    'hosts',
    'hosts_output',
]

@pulumi.output_type
class HostsResult:
    """
    A collection of values returned by Hosts.
    """
    def __init__(__self__, heartbeat_status=None, host_group_id=None, host_infos=None, id=None, ip=None, output_file=None, total_count=None):
        if heartbeat_status and not isinstance(heartbeat_status, int):
            raise TypeError("Expected argument 'heartbeat_status' to be a int")
        pulumi.set(__self__, "heartbeat_status", heartbeat_status)
        if host_group_id and not isinstance(host_group_id, str):
            raise TypeError("Expected argument 'host_group_id' to be a str")
        pulumi.set(__self__, "host_group_id", host_group_id)
        if host_infos and not isinstance(host_infos, list):
            raise TypeError("Expected argument 'host_infos' to be a list")
        pulumi.set(__self__, "host_infos", host_infos)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ip and not isinstance(ip, str):
            raise TypeError("Expected argument 'ip' to be a str")
        pulumi.set(__self__, "ip", ip)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if total_count and not isinstance(total_count, int):
            raise TypeError("Expected argument 'total_count' to be a int")
        pulumi.set(__self__, "total_count", total_count)

    @property
    @pulumi.getter(name="heartbeatStatus")
    def heartbeat_status(self) -> Optional[int]:
        """
        The the heartbeat status.
        """
        return pulumi.get(self, "heartbeat_status")

    @property
    @pulumi.getter(name="hostGroupId")
    def host_group_id(self) -> str:
        """
        The id of host group.
        """
        return pulumi.get(self, "host_group_id")

    @property
    @pulumi.getter(name="hostInfos")
    def host_infos(self) -> Sequence['outputs.HostsHostInfoResult']:
        """
        The collection of query.
        """
        return pulumi.get(self, "host_infos")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ip(self) -> Optional[str]:
        """
        The ip address.
        """
        return pulumi.get(self, "ip")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        """
        The total count of query.
        """
        return pulumi.get(self, "total_count")


class AwaitableHostsResult(HostsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return HostsResult(
            heartbeat_status=self.heartbeat_status,
            host_group_id=self.host_group_id,
            host_infos=self.host_infos,
            id=self.id,
            ip=self.ip,
            output_file=self.output_file,
            total_count=self.total_count)


def hosts(heartbeat_status: Optional[int] = None,
          host_group_id: Optional[str] = None,
          ip: Optional[str] = None,
          output_file: Optional[str] = None,
          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableHostsResult:
    """
    Use this data source to query detailed information of tls hosts
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.tls.hosts(host_group_id="527102e2-1e4f-45f4-a990-751152125da7")
    ```


    :param int heartbeat_status: The the heartbeat status.
    :param str host_group_id: The id of host group.
    :param str ip: The ip address.
    :param str output_file: File name where to save data source results.
    """
    __args__ = dict()
    __args__['heartbeatStatus'] = heartbeat_status
    __args__['hostGroupId'] = host_group_id
    __args__['ip'] = ip
    __args__['outputFile'] = output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('volcengine:tls/hosts:Hosts', __args__, opts=opts, typ=HostsResult).value

    return AwaitableHostsResult(
        heartbeat_status=__ret__.heartbeat_status,
        host_group_id=__ret__.host_group_id,
        host_infos=__ret__.host_infos,
        id=__ret__.id,
        ip=__ret__.ip,
        output_file=__ret__.output_file,
        total_count=__ret__.total_count)


@_utilities.lift_output_func(hosts)
def hosts_output(heartbeat_status: Optional[pulumi.Input[Optional[int]]] = None,
                 host_group_id: Optional[pulumi.Input[str]] = None,
                 ip: Optional[pulumi.Input[Optional[str]]] = None,
                 output_file: Optional[pulumi.Input[Optional[str]]] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[HostsResult]:
    """
    Use this data source to query detailed information of tls hosts
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.tls.hosts(host_group_id="527102e2-1e4f-45f4-a990-751152125da7")
    ```


    :param int heartbeat_status: The the heartbeat status.
    :param str host_group_id: The id of host group.
    :param str ip: The ip address.
    :param str output_file: File name where to save data source results.
    """
    ...
