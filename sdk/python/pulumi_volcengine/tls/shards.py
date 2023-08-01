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
    'ShardsResult',
    'AwaitableShardsResult',
    'shards',
    'shards_output',
]

@pulumi.output_type
class ShardsResult:
    """
    A collection of values returned by Shards.
    """
    def __init__(__self__, id=None, output_file=None, shards=None, topic_id=None, total_count=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if shards and not isinstance(shards, list):
            raise TypeError("Expected argument 'shards' to be a list")
        pulumi.set(__self__, "shards", shards)
        if topic_id and not isinstance(topic_id, str):
            raise TypeError("Expected argument 'topic_id' to be a str")
        pulumi.set(__self__, "topic_id", topic_id)
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
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter
    def shards(self) -> Sequence['outputs.ShardsShardResult']:
        """
        The collection of query.
        """
        return pulumi.get(self, "shards")

    @property
    @pulumi.getter(name="topicId")
    def topic_id(self) -> str:
        """
        The ID of topic.
        """
        return pulumi.get(self, "topic_id")

    @property
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        """
        The total count of query.
        """
        return pulumi.get(self, "total_count")


class AwaitableShardsResult(ShardsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ShardsResult(
            id=self.id,
            output_file=self.output_file,
            shards=self.shards,
            topic_id=self.topic_id,
            total_count=self.total_count)


def shards(output_file: Optional[str] = None,
           topic_id: Optional[str] = None,
           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableShardsResult:
    """
    Use this data source to query detailed information of tls shards
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.tls.shards(topic_id="edf051ed-3c46-49ba-9339-bea628fedc15")
    ```


    :param str output_file: File name where to save data source results.
    :param str topic_id: The id of topic.
    """
    __args__ = dict()
    __args__['outputFile'] = output_file
    __args__['topicId'] = topic_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('volcengine:tls/shards:Shards', __args__, opts=opts, typ=ShardsResult).value

    return AwaitableShardsResult(
        id=__ret__.id,
        output_file=__ret__.output_file,
        shards=__ret__.shards,
        topic_id=__ret__.topic_id,
        total_count=__ret__.total_count)


@_utilities.lift_output_func(shards)
def shards_output(output_file: Optional[pulumi.Input[Optional[str]]] = None,
                  topic_id: Optional[pulumi.Input[str]] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[ShardsResult]:
    """
    Use this data source to query detailed information of tls shards
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.tls.shards(topic_id="edf051ed-3c46-49ba-9339-bea628fedc15")
    ```


    :param str output_file: File name where to save data source results.
    :param str topic_id: The id of topic.
    """
    ...
