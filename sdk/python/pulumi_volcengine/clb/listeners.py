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
    'ListenersResult',
    'AwaitableListenersResult',
    'listeners',
    'listeners_output',
]

@pulumi.output_type
class ListenersResult:
    """
    A collection of values returned by Listeners.
    """
    def __init__(__self__, id=None, ids=None, listener_name=None, listeners=None, load_balancer_id=None, name_regex=None, output_file=None, total_count=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if listener_name and not isinstance(listener_name, str):
            raise TypeError("Expected argument 'listener_name' to be a str")
        pulumi.set(__self__, "listener_name", listener_name)
        if listeners and not isinstance(listeners, list):
            raise TypeError("Expected argument 'listeners' to be a list")
        pulumi.set(__self__, "listeners", listeners)
        if load_balancer_id and not isinstance(load_balancer_id, str):
            raise TypeError("Expected argument 'load_balancer_id' to be a str")
        pulumi.set(__self__, "load_balancer_id", load_balancer_id)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
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
    @pulumi.getter
    def ids(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "ids")

    @property
    @pulumi.getter(name="listenerName")
    def listener_name(self) -> Optional[str]:
        """
        The name of the Listener.
        """
        return pulumi.get(self, "listener_name")

    @property
    @pulumi.getter
    def listeners(self) -> Sequence['outputs.ListenersListenerResult']:
        """
        The collection of Listener query.
        """
        return pulumi.get(self, "listeners")

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
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        """
        The total count of Listener query.
        """
        return pulumi.get(self, "total_count")


class AwaitableListenersResult(ListenersResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListenersResult(
            id=self.id,
            ids=self.ids,
            listener_name=self.listener_name,
            listeners=self.listeners,
            load_balancer_id=self.load_balancer_id,
            name_regex=self.name_regex,
            output_file=self.output_file,
            total_count=self.total_count)


def listeners(ids: Optional[Sequence[str]] = None,
              listener_name: Optional[str] = None,
              load_balancer_id: Optional[str] = None,
              name_regex: Optional[str] = None,
              output_file: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListenersResult:
    """
    Use this data source to query detailed information of listeners
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    foo_zones = volcengine.ecs.zones()
    foo_vpc = volcengine.vpc.Vpc("fooVpc",
        vpc_name="acc-test-vpc",
        cidr_block="172.16.0.0/16")
    foo_subnet = volcengine.vpc.Subnet("fooSubnet",
        subnet_name="acc-test-subnet",
        cidr_block="172.16.0.0/24",
        zone_id=foo_zones.zones[0].id,
        vpc_id=foo_vpc.id)
    foo_clb = volcengine.clb.Clb("fooClb",
        type="public",
        subnet_id=foo_subnet.id,
        load_balancer_spec="small_1",
        description="acc0Demo",
        load_balancer_name="acc-test-create",
        eip_billing_config=volcengine.clb.ClbEipBillingConfigArgs(
            isp="BGP",
            eip_billing_type="PostPaidByBandwidth",
            bandwidth=1,
        ))
    foo_server_group = volcengine.clb.ServerGroup("fooServerGroup",
        load_balancer_id=foo_clb.id,
        server_group_name="acc-test-create",
        description="hello demo11")
    foo_listener = volcengine.clb.Listener("fooListener",
        load_balancer_id=foo_clb.id,
        listener_name="acc-test-listener",
        protocol="HTTP",
        port=90,
        server_group_id=foo_server_group.id,
        health_check=volcengine.clb.ListenerHealthCheckArgs(
            enabled="on",
            interval=10,
            timeout=3,
            healthy_threshold=5,
            un_healthy_threshold=2,
            domain="volcengine.com",
            http_code="http_2xx",
            method="GET",
            uri="/",
        ),
        enabled="on")
    foo_listeners = volcengine.clb.listeners_output(ids=[foo_listener.id])
    ```


    :param Sequence[str] ids: A list of Listener IDs.
    :param str listener_name: The name of the Listener.
    :param str load_balancer_id: The id of the Clb.
    :param str name_regex: A Name Regex of Listener.
    :param str output_file: File name where to save data source results.
    """
    __args__ = dict()
    __args__['ids'] = ids
    __args__['listenerName'] = listener_name
    __args__['loadBalancerId'] = load_balancer_id
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('volcengine:clb/listeners:Listeners', __args__, opts=opts, typ=ListenersResult).value

    return AwaitableListenersResult(
        id=__ret__.id,
        ids=__ret__.ids,
        listener_name=__ret__.listener_name,
        listeners=__ret__.listeners,
        load_balancer_id=__ret__.load_balancer_id,
        name_regex=__ret__.name_regex,
        output_file=__ret__.output_file,
        total_count=__ret__.total_count)


@_utilities.lift_output_func(listeners)
def listeners_output(ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                     listener_name: Optional[pulumi.Input[Optional[str]]] = None,
                     load_balancer_id: Optional[pulumi.Input[Optional[str]]] = None,
                     name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                     output_file: Optional[pulumi.Input[Optional[str]]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[ListenersResult]:
    """
    Use this data source to query detailed information of listeners
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    foo_zones = volcengine.ecs.zones()
    foo_vpc = volcengine.vpc.Vpc("fooVpc",
        vpc_name="acc-test-vpc",
        cidr_block="172.16.0.0/16")
    foo_subnet = volcengine.vpc.Subnet("fooSubnet",
        subnet_name="acc-test-subnet",
        cidr_block="172.16.0.0/24",
        zone_id=foo_zones.zones[0].id,
        vpc_id=foo_vpc.id)
    foo_clb = volcengine.clb.Clb("fooClb",
        type="public",
        subnet_id=foo_subnet.id,
        load_balancer_spec="small_1",
        description="acc0Demo",
        load_balancer_name="acc-test-create",
        eip_billing_config=volcengine.clb.ClbEipBillingConfigArgs(
            isp="BGP",
            eip_billing_type="PostPaidByBandwidth",
            bandwidth=1,
        ))
    foo_server_group = volcengine.clb.ServerGroup("fooServerGroup",
        load_balancer_id=foo_clb.id,
        server_group_name="acc-test-create",
        description="hello demo11")
    foo_listener = volcengine.clb.Listener("fooListener",
        load_balancer_id=foo_clb.id,
        listener_name="acc-test-listener",
        protocol="HTTP",
        port=90,
        server_group_id=foo_server_group.id,
        health_check=volcengine.clb.ListenerHealthCheckArgs(
            enabled="on",
            interval=10,
            timeout=3,
            healthy_threshold=5,
            un_healthy_threshold=2,
            domain="volcengine.com",
            http_code="http_2xx",
            method="GET",
            uri="/",
        ),
        enabled="on")
    foo_listeners = volcengine.clb.listeners_output(ids=[foo_listener.id])
    ```


    :param Sequence[str] ids: A list of Listener IDs.
    :param str listener_name: The name of the Listener.
    :param str load_balancer_id: The id of the Clb.
    :param str name_regex: A Name Regex of Listener.
    :param str output_file: File name where to save data source results.
    """
    ...
