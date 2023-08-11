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
    'BucketsResult',
    'AwaitableBucketsResult',
    'buckets',
    'buckets_output',
]

@pulumi.output_type
class BucketsResult:
    """
    A collection of values returned by Buckets.
    """
    def __init__(__self__, bucket_name=None, buckets=None, id=None, name_regex=None, output_file=None, total_count=None):
        if bucket_name and not isinstance(bucket_name, str):
            raise TypeError("Expected argument 'bucket_name' to be a str")
        pulumi.set(__self__, "bucket_name", bucket_name)
        if buckets and not isinstance(buckets, list):
            raise TypeError("Expected argument 'buckets' to be a list")
        pulumi.set(__self__, "buckets", buckets)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
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
    @pulumi.getter(name="bucketName")
    def bucket_name(self) -> Optional[str]:
        return pulumi.get(self, "bucket_name")

    @property
    @pulumi.getter
    def buckets(self) -> Sequence['outputs.BucketsBucketResult']:
        """
        The collection of TOS bucket query.
        """
        return pulumi.get(self, "buckets")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

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
        The total count of TOS bucket query.
        """
        return pulumi.get(self, "total_count")


class AwaitableBucketsResult(BucketsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return BucketsResult(
            bucket_name=self.bucket_name,
            buckets=self.buckets,
            id=self.id,
            name_regex=self.name_regex,
            output_file=self.output_file,
            total_count=self.total_count)


def buckets(bucket_name: Optional[str] = None,
            name_regex: Optional[str] = None,
            output_file: Optional[str] = None,
            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableBucketsResult:
    """
    Use this data source to query detailed information of tos buckets
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.tos.buckets(name_regex="test")
    ```


    :param str bucket_name: The name the TOS bucket.
    :param str name_regex: A Name Regex of TOS bucket.
    :param str output_file: File name where to save data source results.
    """
    __args__ = dict()
    __args__['bucketName'] = bucket_name
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('volcengine:tos/buckets:Buckets', __args__, opts=opts, typ=BucketsResult).value

    return AwaitableBucketsResult(
        bucket_name=pulumi.get(__ret__, 'bucket_name'),
        buckets=pulumi.get(__ret__, 'buckets'),
        id=pulumi.get(__ret__, 'id'),
        name_regex=pulumi.get(__ret__, 'name_regex'),
        output_file=pulumi.get(__ret__, 'output_file'),
        total_count=pulumi.get(__ret__, 'total_count'))


@_utilities.lift_output_func(buckets)
def buckets_output(bucket_name: Optional[pulumi.Input[Optional[str]]] = None,
                   name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                   output_file: Optional[pulumi.Input[Optional[str]]] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[BucketsResult]:
    """
    Use this data source to query detailed information of tos buckets
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.tos.buckets(name_regex="test")
    ```


    :param str bucket_name: The name the TOS bucket.
    :param str name_regex: A Name Regex of TOS bucket.
    :param str output_file: File name where to save data source results.
    """
    ...
