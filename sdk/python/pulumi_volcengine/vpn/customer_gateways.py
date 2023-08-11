# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'CustomerGatewaysResult',
    'AwaitableCustomerGatewaysResult',
    'customer_gateways',
    'customer_gateways_output',
]

@pulumi.output_type
class CustomerGatewaysResult:
    """
    A collection of values returned by CustomerGateways.
    """
    def __init__(__self__, customer_gateway_names=None, customer_gateways=None, id=None, ids=None, ip_address=None, name_regex=None, output_file=None, total_count=None):
        if customer_gateway_names and not isinstance(customer_gateway_names, list):
            raise TypeError("Expected argument 'customer_gateway_names' to be a list")
        pulumi.set(__self__, "customer_gateway_names", customer_gateway_names)
        if customer_gateways and not isinstance(customer_gateways, list):
            raise TypeError("Expected argument 'customer_gateways' to be a list")
        pulumi.set(__self__, "customer_gateways", customer_gateways)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if ip_address and not isinstance(ip_address, str):
            raise TypeError("Expected argument 'ip_address' to be a str")
        pulumi.set(__self__, "ip_address", ip_address)
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
    @pulumi.getter(name="customerGatewayNames")
    def customer_gateway_names(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "customer_gateway_names")

    @property
    @pulumi.getter(name="customerGateways")
    def customer_gateways(self) -> Sequence['outputs.CustomerGatewaysCustomerGatewayResult']:
        """
        The collection of customer gateway query.
        """
        return pulumi.get(self, "customer_gateways")

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
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> Optional[str]:
        """
        The IP address of the customer gateway.
        """
        return pulumi.get(self, "ip_address")

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
        The total count of customer gateway query.
        """
        return pulumi.get(self, "total_count")


class AwaitableCustomerGatewaysResult(CustomerGatewaysResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return CustomerGatewaysResult(
            customer_gateway_names=self.customer_gateway_names,
            customer_gateways=self.customer_gateways,
            id=self.id,
            ids=self.ids,
            ip_address=self.ip_address,
            name_regex=self.name_regex,
            output_file=self.output_file,
            total_count=self.total_count)


def customer_gateways(customer_gateway_names: Optional[Sequence[str]] = None,
                      ids: Optional[Sequence[str]] = None,
                      ip_address: Optional[str] = None,
                      name_regex: Optional[str] = None,
                      output_file: Optional[str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableCustomerGatewaysResult:
    """
    Use this data source to query detailed information of customer gateways
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    foo = volcengine.vpn.customer_gateways(ids=["cgw-2d68c4zglycjk58ozfe96norh"])
    ```


    :param Sequence[str] customer_gateway_names: A list of customer gateway names.
    :param Sequence[str] ids: A list of customer gateway ids.
    :param str ip_address: A IP address of the customer gateway.
    :param str name_regex: A Name Regex of customer gateway.
    :param str output_file: File name where to save data source results.
    """
    __args__ = dict()
    __args__['customerGatewayNames'] = customer_gateway_names
    __args__['ids'] = ids
    __args__['ipAddress'] = ip_address
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('volcengine:vpn/customerGateways:CustomerGateways', __args__, opts=opts, typ=CustomerGatewaysResult).value

    return AwaitableCustomerGatewaysResult(
        customer_gateway_names=pulumi.get(__ret__, 'customer_gateway_names'),
        customer_gateways=pulumi.get(__ret__, 'customer_gateways'),
        id=pulumi.get(__ret__, 'id'),
        ids=pulumi.get(__ret__, 'ids'),
        ip_address=pulumi.get(__ret__, 'ip_address'),
        name_regex=pulumi.get(__ret__, 'name_regex'),
        output_file=pulumi.get(__ret__, 'output_file'),
        total_count=pulumi.get(__ret__, 'total_count'))


@_utilities.lift_output_func(customer_gateways)
def customer_gateways_output(customer_gateway_names: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                             ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                             ip_address: Optional[pulumi.Input[Optional[str]]] = None,
                             name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                             output_file: Optional[pulumi.Input[Optional[str]]] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[CustomerGatewaysResult]:
    """
    Use this data source to query detailed information of customer gateways
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    foo = volcengine.vpn.customer_gateways(ids=["cgw-2d68c4zglycjk58ozfe96norh"])
    ```


    :param Sequence[str] customer_gateway_names: A list of customer gateway names.
    :param Sequence[str] ids: A list of customer gateway ids.
    :param str ip_address: A IP address of the customer gateway.
    :param str name_regex: A Name Regex of customer gateway.
    :param str output_file: File name where to save data source results.
    """
    ...
