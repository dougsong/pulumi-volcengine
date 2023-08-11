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
    'Ipv6GatewaysResult',
    'AwaitableIpv6GatewaysResult',
    'ipv6_gateways',
    'ipv6_gateways_output',
]

@pulumi.output_type
class Ipv6GatewaysResult:
    """
    A collection of values returned by Ipv6Gateways.
    """
    def __init__(__self__, id=None, ids=None, ipv6_gateways=None, name=None, name_regex=None, output_file=None, total_count=None, vpc_ids=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if ipv6_gateways and not isinstance(ipv6_gateways, list):
            raise TypeError("Expected argument 'ipv6_gateways' to be a list")
        pulumi.set(__self__, "ipv6_gateways", ipv6_gateways)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if total_count and not isinstance(total_count, int):
            raise TypeError("Expected argument 'total_count' to be a int")
        pulumi.set(__self__, "total_count", total_count)
        if vpc_ids and not isinstance(vpc_ids, list):
            raise TypeError("Expected argument 'vpc_ids' to be a list")
        pulumi.set(__self__, "vpc_ids", vpc_ids)

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
    @pulumi.getter(name="ipv6Gateways")
    def ipv6_gateways(self) -> Sequence['outputs.Ipv6GatewaysIpv6GatewayResult']:
        """
        The collection of Ipv6Gateway query.
        """
        return pulumi.get(self, "ipv6_gateways")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The Name of the Ipv6Gateway.
        """
        return pulumi.get(self, "name")

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
        The total count of Ipv6Gateway query.
        """
        return pulumi.get(self, "total_count")

    @property
    @pulumi.getter(name="vpcIds")
    def vpc_ids(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "vpc_ids")


class AwaitableIpv6GatewaysResult(Ipv6GatewaysResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return Ipv6GatewaysResult(
            id=self.id,
            ids=self.ids,
            ipv6_gateways=self.ipv6_gateways,
            name=self.name,
            name_regex=self.name_regex,
            output_file=self.output_file,
            total_count=self.total_count,
            vpc_ids=self.vpc_ids)


def ipv6_gateways(ids: Optional[Sequence[str]] = None,
                  name: Optional[str] = None,
                  name_regex: Optional[str] = None,
                  output_file: Optional[str] = None,
                  vpc_ids: Optional[Sequence[str]] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableIpv6GatewaysResult:
    """
    Use this data source to query detailed information of vpc ipv6 gateways
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.vpc.ipv6_gateways(ids=["ipv6gw-12bcapllb5ukg17q7y2sd3thx"])
    ```


    :param Sequence[str] ids: The ID list of the Ipv6Gateways.
    :param str name: The name of the Ipv6Gateway.
    :param str name_regex: A Name Regex of the Ipv6Gateway.
    :param str output_file: File name where to save data source results.
    :param Sequence[str] vpc_ids: The ID list of the VPC which the Ipv6Gateway belongs to.
    """
    __args__ = dict()
    __args__['ids'] = ids
    __args__['name'] = name
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['vpcIds'] = vpc_ids
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('volcengine:vpc/ipv6Gateways:Ipv6Gateways', __args__, opts=opts, typ=Ipv6GatewaysResult).value

    return AwaitableIpv6GatewaysResult(
        id=pulumi.get(__ret__, 'id'),
        ids=pulumi.get(__ret__, 'ids'),
        ipv6_gateways=pulumi.get(__ret__, 'ipv6_gateways'),
        name=pulumi.get(__ret__, 'name'),
        name_regex=pulumi.get(__ret__, 'name_regex'),
        output_file=pulumi.get(__ret__, 'output_file'),
        total_count=pulumi.get(__ret__, 'total_count'),
        vpc_ids=pulumi.get(__ret__, 'vpc_ids'))


@_utilities.lift_output_func(ipv6_gateways)
def ipv6_gateways_output(ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                         name: Optional[pulumi.Input[Optional[str]]] = None,
                         name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                         output_file: Optional[pulumi.Input[Optional[str]]] = None,
                         vpc_ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[Ipv6GatewaysResult]:
    """
    Use this data source to query detailed information of vpc ipv6 gateways
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.vpc.ipv6_gateways(ids=["ipv6gw-12bcapllb5ukg17q7y2sd3thx"])
    ```


    :param Sequence[str] ids: The ID list of the Ipv6Gateways.
    :param str name: The name of the Ipv6Gateway.
    :param str name_regex: A Name Regex of the Ipv6Gateway.
    :param str output_file: File name where to save data source results.
    :param Sequence[str] vpc_ids: The ID list of the VPC which the Ipv6Gateway belongs to.
    """
    ...
