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
    'VolumesResult',
    'AwaitableVolumesResult',
    'volumes',
    'volumes_output',
]

@pulumi.output_type
class VolumesResult:
    """
    A collection of values returned by Volumes.
    """
    def __init__(__self__, id=None, ids=None, instance_id=None, kind=None, name_regex=None, output_file=None, total_count=None, volume_name=None, volume_status=None, volume_type=None, volumes=None, zone_id=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if total_count and not isinstance(total_count, int):
            raise TypeError("Expected argument 'total_count' to be a int")
        pulumi.set(__self__, "total_count", total_count)
        if volume_name and not isinstance(volume_name, str):
            raise TypeError("Expected argument 'volume_name' to be a str")
        pulumi.set(__self__, "volume_name", volume_name)
        if volume_status and not isinstance(volume_status, str):
            raise TypeError("Expected argument 'volume_status' to be a str")
        pulumi.set(__self__, "volume_status", volume_status)
        if volume_type and not isinstance(volume_type, str):
            raise TypeError("Expected argument 'volume_type' to be a str")
        pulumi.set(__self__, "volume_type", volume_type)
        if volumes and not isinstance(volumes, list):
            raise TypeError("Expected argument 'volumes' to be a list")
        pulumi.set(__self__, "volumes", volumes)
        if zone_id and not isinstance(zone_id, str):
            raise TypeError("Expected argument 'zone_id' to be a str")
        pulumi.set(__self__, "zone_id", zone_id)

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
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[str]:
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def kind(self) -> Optional[str]:
        return pulumi.get(self, "kind")

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
        The total count of Volume query.
        """
        return pulumi.get(self, "total_count")

    @property
    @pulumi.getter(name="volumeName")
    def volume_name(self) -> Optional[str]:
        return pulumi.get(self, "volume_name")

    @property
    @pulumi.getter(name="volumeStatus")
    def volume_status(self) -> Optional[str]:
        return pulumi.get(self, "volume_status")

    @property
    @pulumi.getter(name="volumeType")
    def volume_type(self) -> Optional[str]:
        return pulumi.get(self, "volume_type")

    @property
    @pulumi.getter
    def volumes(self) -> Sequence['outputs.VolumesVolumeResult']:
        """
        The collection of Volume query.
        """
        return pulumi.get(self, "volumes")

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> Optional[str]:
        return pulumi.get(self, "zone_id")


class AwaitableVolumesResult(VolumesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return VolumesResult(
            id=self.id,
            ids=self.ids,
            instance_id=self.instance_id,
            kind=self.kind,
            name_regex=self.name_regex,
            output_file=self.output_file,
            total_count=self.total_count,
            volume_name=self.volume_name,
            volume_status=self.volume_status,
            volume_type=self.volume_type,
            volumes=self.volumes,
            zone_id=self.zone_id)


def volumes(ids: Optional[Sequence[str]] = None,
            instance_id: Optional[str] = None,
            kind: Optional[str] = None,
            name_regex: Optional[str] = None,
            output_file: Optional[str] = None,
            volume_name: Optional[str] = None,
            volume_status: Optional[str] = None,
            volume_type: Optional[str] = None,
            zone_id: Optional[str] = None,
            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableVolumesResult:
    """
    Use this data source to query detailed information of volumes
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.ebs.volumes(ids=["vol-3tzg6y5imn3b9fchkedb"])
    ```


    :param Sequence[str] ids: A list of Volume IDs.
    :param str instance_id: The Id of instance.
    :param str kind: The Kind of Volume.
    :param str name_regex: A Name Regex of Volume.
    :param str output_file: File name where to save data source results.
    :param str volume_name: The name of Volume.
    :param str volume_status: The Status of Volume, the value can be `available` or `attaching` or `attached` or `detaching` or `creating` or `deleting` or `error` or `extending`.
    :param str volume_type: The type of Volume.
    :param str zone_id: The Id of Zone.
    """
    __args__ = dict()
    __args__['ids'] = ids
    __args__['instanceId'] = instance_id
    __args__['kind'] = kind
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['volumeName'] = volume_name
    __args__['volumeStatus'] = volume_status
    __args__['volumeType'] = volume_type
    __args__['zoneId'] = zone_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('volcengine:ebs/volumes:Volumes', __args__, opts=opts, typ=VolumesResult).value

    return AwaitableVolumesResult(
        id=__ret__.id,
        ids=__ret__.ids,
        instance_id=__ret__.instance_id,
        kind=__ret__.kind,
        name_regex=__ret__.name_regex,
        output_file=__ret__.output_file,
        total_count=__ret__.total_count,
        volume_name=__ret__.volume_name,
        volume_status=__ret__.volume_status,
        volume_type=__ret__.volume_type,
        volumes=__ret__.volumes,
        zone_id=__ret__.zone_id)


@_utilities.lift_output_func(volumes)
def volumes_output(ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                   instance_id: Optional[pulumi.Input[Optional[str]]] = None,
                   kind: Optional[pulumi.Input[Optional[str]]] = None,
                   name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                   output_file: Optional[pulumi.Input[Optional[str]]] = None,
                   volume_name: Optional[pulumi.Input[Optional[str]]] = None,
                   volume_status: Optional[pulumi.Input[Optional[str]]] = None,
                   volume_type: Optional[pulumi.Input[Optional[str]]] = None,
                   zone_id: Optional[pulumi.Input[Optional[str]]] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[VolumesResult]:
    """
    Use this data source to query detailed information of volumes
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.ebs.volumes(ids=["vol-3tzg6y5imn3b9fchkedb"])
    ```


    :param Sequence[str] ids: A list of Volume IDs.
    :param str instance_id: The Id of instance.
    :param str kind: The Kind of Volume.
    :param str name_regex: A Name Regex of Volume.
    :param str output_file: File name where to save data source results.
    :param str volume_name: The name of Volume.
    :param str volume_status: The Status of Volume, the value can be `available` or `attaching` or `attached` or `detaching` or `creating` or `deleting` or `error` or `extending`.
    :param str volume_type: The type of Volume.
    :param str zone_id: The Id of Zone.
    """
    ...
