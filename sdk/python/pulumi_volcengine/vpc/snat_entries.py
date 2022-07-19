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
    'SnatEntriesResult',
    'AwaitableSnatEntriesResult',
    'snat_entries',
    'snat_entries_output',
]

@pulumi.output_type
class SnatEntriesResult:
    """
    A collection of values returned by SnatEntries.
    """
    def __init__(__self__, eip_id=None, id=None, ids=None, nat_gateway_id=None, output_file=None, snat_entries=None, snat_entry_name=None, subnet_id=None, total_count=None):
        if eip_id and not isinstance(eip_id, str):
            raise TypeError("Expected argument 'eip_id' to be a str")
        pulumi.set(__self__, "eip_id", eip_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if nat_gateway_id and not isinstance(nat_gateway_id, str):
            raise TypeError("Expected argument 'nat_gateway_id' to be a str")
        pulumi.set(__self__, "nat_gateway_id", nat_gateway_id)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if snat_entries and not isinstance(snat_entries, list):
            raise TypeError("Expected argument 'snat_entries' to be a list")
        pulumi.set(__self__, "snat_entries", snat_entries)
        if snat_entry_name and not isinstance(snat_entry_name, str):
            raise TypeError("Expected argument 'snat_entry_name' to be a str")
        pulumi.set(__self__, "snat_entry_name", snat_entry_name)
        if subnet_id and not isinstance(subnet_id, str):
            raise TypeError("Expected argument 'subnet_id' to be a str")
        pulumi.set(__self__, "subnet_id", subnet_id)
        if total_count and not isinstance(total_count, int):
            raise TypeError("Expected argument 'total_count' to be a int")
        pulumi.set(__self__, "total_count", total_count)

    @property
    @pulumi.getter(name="eipId")
    def eip_id(self) -> Optional[str]:
        """
        The id of the public ip address used by the SNAT entry.
        """
        return pulumi.get(self, "eip_id")

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
    @pulumi.getter(name="natGatewayId")
    def nat_gateway_id(self) -> Optional[str]:
        """
        The id of the nat gateway to which the entry belongs.
        """
        return pulumi.get(self, "nat_gateway_id")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="snatEntries")
    def snat_entries(self) -> Sequence['outputs.SnatEntriesSnatEntryResult']:
        """
        The collection of snat entries.
        """
        return pulumi.get(self, "snat_entries")

    @property
    @pulumi.getter(name="snatEntryName")
    def snat_entry_name(self) -> Optional[str]:
        """
        The name of the SNAT entry.
        """
        return pulumi.get(self, "snat_entry_name")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> Optional[str]:
        """
        The id of the subnet that is required to access the internet.
        """
        return pulumi.get(self, "subnet_id")

    @property
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        """
        The total count of snat entries query.
        """
        return pulumi.get(self, "total_count")


class AwaitableSnatEntriesResult(SnatEntriesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return SnatEntriesResult(
            eip_id=self.eip_id,
            id=self.id,
            ids=self.ids,
            nat_gateway_id=self.nat_gateway_id,
            output_file=self.output_file,
            snat_entries=self.snat_entries,
            snat_entry_name=self.snat_entry_name,
            subnet_id=self.subnet_id,
            total_count=self.total_count)


def snat_entries(eip_id: Optional[str] = None,
                 ids: Optional[Sequence[str]] = None,
                 nat_gateway_id: Optional[str] = None,
                 output_file: Optional[str] = None,
                 snat_entry_name: Optional[str] = None,
                 subnet_id: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableSnatEntriesResult:
    """
    Use this data source to query detailed information of snat entries
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.Vpc.snat_entries(ids=["snat-274zl8b1kxzb47fap8u35uune"])
    ```


    :param str eip_id: An id of the public ip address used by the SNAT entry.
    :param Sequence[str] ids: A list of SNAT entry ids.
    :param str nat_gateway_id: An id of the nat gateway to which the entry belongs.
    :param str output_file: File name where to save data source results.
    :param str snat_entry_name: A name of SNAT entry.
    :param str subnet_id: An id of the subnet that is required to access the Internet.
    """
    __args__ = dict()
    __args__['eipId'] = eip_id
    __args__['ids'] = ids
    __args__['natGatewayId'] = nat_gateway_id
    __args__['outputFile'] = output_file
    __args__['snatEntryName'] = snat_entry_name
    __args__['subnetId'] = subnet_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('volcengine:Vpc/snatEntries:SnatEntries', __args__, opts=opts, typ=SnatEntriesResult).value

    return AwaitableSnatEntriesResult(
        eip_id=__ret__.eip_id,
        id=__ret__.id,
        ids=__ret__.ids,
        nat_gateway_id=__ret__.nat_gateway_id,
        output_file=__ret__.output_file,
        snat_entries=__ret__.snat_entries,
        snat_entry_name=__ret__.snat_entry_name,
        subnet_id=__ret__.subnet_id,
        total_count=__ret__.total_count)


@_utilities.lift_output_func(snat_entries)
def snat_entries_output(eip_id: Optional[pulumi.Input[Optional[str]]] = None,
                        ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                        nat_gateway_id: Optional[pulumi.Input[Optional[str]]] = None,
                        output_file: Optional[pulumi.Input[Optional[str]]] = None,
                        snat_entry_name: Optional[pulumi.Input[Optional[str]]] = None,
                        subnet_id: Optional[pulumi.Input[Optional[str]]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[SnatEntriesResult]:
    """
    Use this data source to query detailed information of snat entries
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.Vpc.snat_entries(ids=["snat-274zl8b1kxzb47fap8u35uune"])
    ```


    :param str eip_id: An id of the public ip address used by the SNAT entry.
    :param Sequence[str] ids: A list of SNAT entry ids.
    :param str nat_gateway_id: An id of the nat gateway to which the entry belongs.
    :param str output_file: File name where to save data source results.
    :param str snat_entry_name: A name of SNAT entry.
    :param str subnet_id: An id of the subnet that is required to access the Internet.
    """
    ...
