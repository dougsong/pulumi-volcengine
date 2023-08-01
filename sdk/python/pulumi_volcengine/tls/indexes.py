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
    'IndexesResult',
    'AwaitableIndexesResult',
    'indexes',
    'indexes_output',
]

@pulumi.output_type
class IndexesResult:
    """
    A collection of values returned by Indexes.
    """
    def __init__(__self__, id=None, ids=None, output_file=None, tls_indexes=None, total_count=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if tls_indexes and not isinstance(tls_indexes, list):
            raise TypeError("Expected argument 'tls_indexes' to be a list")
        pulumi.set(__self__, "tls_indexes", tls_indexes)
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
    def ids(self) -> Sequence[str]:
        return pulumi.get(self, "ids")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="tlsIndexes")
    def tls_indexes(self) -> Sequence['outputs.IndexesTlsIndexResult']:
        """
        The collection of tls index query.
        """
        return pulumi.get(self, "tls_indexes")

    @property
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        """
        The total count of tls index query.
        """
        return pulumi.get(self, "total_count")


class AwaitableIndexesResult(IndexesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return IndexesResult(
            id=self.id,
            ids=self.ids,
            output_file=self.output_file,
            tls_indexes=self.tls_indexes,
            total_count=self.total_count)


def indexes(ids: Optional[Sequence[str]] = None,
            output_file: Optional[str] = None,
            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableIndexesResult:
    """
    Use this data source to query detailed information of tls indexes
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.tls.indexes(ids=["65d67d34-c5b4-4ec8-b3a9-175d3366****"])
    ```


    :param Sequence[str] ids: The list of topic id of tls index.
    :param str output_file: File name where to save data source results.
    """
    __args__ = dict()
    __args__['ids'] = ids
    __args__['outputFile'] = output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('volcengine:tls/indexes:Indexes', __args__, opts=opts, typ=IndexesResult).value

    return AwaitableIndexesResult(
        id=__ret__.id,
        ids=__ret__.ids,
        output_file=__ret__.output_file,
        tls_indexes=__ret__.tls_indexes,
        total_count=__ret__.total_count)


@_utilities.lift_output_func(indexes)
def indexes_output(ids: Optional[pulumi.Input[Sequence[str]]] = None,
                   output_file: Optional[pulumi.Input[Optional[str]]] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[IndexesResult]:
    """
    Use this data source to query detailed information of tls indexes
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.tls.indexes(ids=["65d67d34-c5b4-4ec8-b3a9-175d3366****"])
    ```


    :param Sequence[str] ids: The list of topic id of tls index.
    :param str output_file: File name where to save data source results.
    """
    ...
