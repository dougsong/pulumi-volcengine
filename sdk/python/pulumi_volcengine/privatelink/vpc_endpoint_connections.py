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
    'VpcEndpointConnectionsResult',
    'AwaitableVpcEndpointConnectionsResult',
    'vpc_endpoint_connections',
    'vpc_endpoint_connections_output',
]

@pulumi.output_type
class VpcEndpointConnectionsResult:
    """
    A collection of values returned by VpcEndpointConnections.
    """
    def __init__(__self__, connections=None, endpoint_id=None, endpoint_owner_account_id=None, id=None, output_file=None, service_id=None, total_count=None):
        if connections and not isinstance(connections, list):
            raise TypeError("Expected argument 'connections' to be a list")
        pulumi.set(__self__, "connections", connections)
        if endpoint_id and not isinstance(endpoint_id, str):
            raise TypeError("Expected argument 'endpoint_id' to be a str")
        pulumi.set(__self__, "endpoint_id", endpoint_id)
        if endpoint_owner_account_id and not isinstance(endpoint_owner_account_id, str):
            raise TypeError("Expected argument 'endpoint_owner_account_id' to be a str")
        pulumi.set(__self__, "endpoint_owner_account_id", endpoint_owner_account_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if service_id and not isinstance(service_id, str):
            raise TypeError("Expected argument 'service_id' to be a str")
        pulumi.set(__self__, "service_id", service_id)
        if total_count and not isinstance(total_count, int):
            raise TypeError("Expected argument 'total_count' to be a int")
        pulumi.set(__self__, "total_count", total_count)

    @property
    @pulumi.getter
    def connections(self) -> Sequence['outputs.VpcEndpointConnectionsConnectionResult']:
        """
        The list of query.
        """
        return pulumi.get(self, "connections")

    @property
    @pulumi.getter(name="endpointId")
    def endpoint_id(self) -> Optional[str]:
        """
        The id of the vpc endpoint.
        """
        return pulumi.get(self, "endpoint_id")

    @property
    @pulumi.getter(name="endpointOwnerAccountId")
    def endpoint_owner_account_id(self) -> Optional[str]:
        """
        The account id of the vpc endpoint.
        """
        return pulumi.get(self, "endpoint_owner_account_id")

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
    @pulumi.getter(name="serviceId")
    def service_id(self) -> str:
        """
        The id of the vpc endpoint service.
        """
        return pulumi.get(self, "service_id")

    @property
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        """
        Returns the total amount of the data list.
        """
        return pulumi.get(self, "total_count")


class AwaitableVpcEndpointConnectionsResult(VpcEndpointConnectionsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return VpcEndpointConnectionsResult(
            connections=self.connections,
            endpoint_id=self.endpoint_id,
            endpoint_owner_account_id=self.endpoint_owner_account_id,
            id=self.id,
            output_file=self.output_file,
            service_id=self.service_id,
            total_count=self.total_count)


def vpc_endpoint_connections(endpoint_id: Optional[str] = None,
                             endpoint_owner_account_id: Optional[str] = None,
                             output_file: Optional[str] = None,
                             service_id: Optional[str] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableVpcEndpointConnectionsResult:
    """
    Use this data source to query detailed information of privatelink vpc endpoint connections
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.privatelink.vpc_endpoint_connections(endpoint_id="ep-3rel74u229dz45zsk2i6l69qa",
        service_id="epsvc-2byz5mykk9y4g2dx0efs4aqz3")
    ```


    :param str endpoint_id: The id of the vpc endpoint.
    :param str endpoint_owner_account_id: The account id of the vpc endpoint.
    :param str output_file: File name where to save data source results.
    :param str service_id: The id of the vpc endpoint service.
    """
    __args__ = dict()
    __args__['endpointId'] = endpoint_id
    __args__['endpointOwnerAccountId'] = endpoint_owner_account_id
    __args__['outputFile'] = output_file
    __args__['serviceId'] = service_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('volcengine:privatelink/vpcEndpointConnections:VpcEndpointConnections', __args__, opts=opts, typ=VpcEndpointConnectionsResult).value

    return AwaitableVpcEndpointConnectionsResult(
        connections=__ret__.connections,
        endpoint_id=__ret__.endpoint_id,
        endpoint_owner_account_id=__ret__.endpoint_owner_account_id,
        id=__ret__.id,
        output_file=__ret__.output_file,
        service_id=__ret__.service_id,
        total_count=__ret__.total_count)


@_utilities.lift_output_func(vpc_endpoint_connections)
def vpc_endpoint_connections_output(endpoint_id: Optional[pulumi.Input[Optional[str]]] = None,
                                    endpoint_owner_account_id: Optional[pulumi.Input[Optional[str]]] = None,
                                    output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                    service_id: Optional[pulumi.Input[str]] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[VpcEndpointConnectionsResult]:
    """
    Use this data source to query detailed information of privatelink vpc endpoint connections
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.privatelink.vpc_endpoint_connections(endpoint_id="ep-3rel74u229dz45zsk2i6l69qa",
        service_id="epsvc-2byz5mykk9y4g2dx0efs4aqz3")
    ```


    :param str endpoint_id: The id of the vpc endpoint.
    :param str endpoint_owner_account_id: The account id of the vpc endpoint.
    :param str output_file: File name where to save data source results.
    :param str service_id: The id of the vpc endpoint service.
    """
    ...
