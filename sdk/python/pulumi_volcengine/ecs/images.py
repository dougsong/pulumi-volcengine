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
    'ImagesResult',
    'AwaitableImagesResult',
    'images',
    'images_output',
]

@pulumi.output_type
class ImagesResult:
    """
    A collection of values returned by Images.
    """
    def __init__(__self__, id=None, ids=None, images=None, instance_type_id=None, is_support_cloud_init=None, name_regex=None, os_type=None, output_file=None, statuses=None, total_count=None, visibility=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if images and not isinstance(images, list):
            raise TypeError("Expected argument 'images' to be a list")
        pulumi.set(__self__, "images", images)
        if instance_type_id and not isinstance(instance_type_id, str):
            raise TypeError("Expected argument 'instance_type_id' to be a str")
        pulumi.set(__self__, "instance_type_id", instance_type_id)
        if is_support_cloud_init and not isinstance(is_support_cloud_init, bool):
            raise TypeError("Expected argument 'is_support_cloud_init' to be a bool")
        pulumi.set(__self__, "is_support_cloud_init", is_support_cloud_init)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if os_type and not isinstance(os_type, str):
            raise TypeError("Expected argument 'os_type' to be a str")
        pulumi.set(__self__, "os_type", os_type)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if statuses and not isinstance(statuses, list):
            raise TypeError("Expected argument 'statuses' to be a list")
        pulumi.set(__self__, "statuses", statuses)
        if total_count and not isinstance(total_count, int):
            raise TypeError("Expected argument 'total_count' to be a int")
        pulumi.set(__self__, "total_count", total_count)
        if visibility and not isinstance(visibility, str):
            raise TypeError("Expected argument 'visibility' to be a str")
        pulumi.set(__self__, "visibility", visibility)

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
    @pulumi.getter
    def images(self) -> Sequence['outputs.ImagesImageResult']:
        """
        The collection of Image query.
        """
        return pulumi.get(self, "images")

    @property
    @pulumi.getter(name="instanceTypeId")
    def instance_type_id(self) -> Optional[str]:
        return pulumi.get(self, "instance_type_id")

    @property
    @pulumi.getter(name="isSupportCloudInit")
    def is_support_cloud_init(self) -> Optional[bool]:
        """
        Whether the Image support cloud-init.
        """
        return pulumi.get(self, "is_support_cloud_init")

    @property
    @pulumi.getter(name="nameRegex")
    def name_regex(self) -> Optional[str]:
        return pulumi.get(self, "name_regex")

    @property
    @pulumi.getter(name="osType")
    def os_type(self) -> Optional[str]:
        """
        The operating system type of Image.
        """
        return pulumi.get(self, "os_type")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter
    def statuses(self) -> Optional[Sequence[str]]:
        """
        The status of Image.
        """
        return pulumi.get(self, "statuses")

    @property
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        """
        The total count of Image query.
        """
        return pulumi.get(self, "total_count")

    @property
    @pulumi.getter
    def visibility(self) -> Optional[str]:
        """
        The visibility of Image.
        """
        return pulumi.get(self, "visibility")


class AwaitableImagesResult(ImagesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ImagesResult(
            id=self.id,
            ids=self.ids,
            images=self.images,
            instance_type_id=self.instance_type_id,
            is_support_cloud_init=self.is_support_cloud_init,
            name_regex=self.name_regex,
            os_type=self.os_type,
            output_file=self.output_file,
            statuses=self.statuses,
            total_count=self.total_count,
            visibility=self.visibility)


def images(ids: Optional[Sequence[str]] = None,
           instance_type_id: Optional[str] = None,
           is_support_cloud_init: Optional[bool] = None,
           name_regex: Optional[str] = None,
           os_type: Optional[str] = None,
           output_file: Optional[str] = None,
           statuses: Optional[Sequence[str]] = None,
           visibility: Optional[str] = None,
           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableImagesResult:
    """
    Use this data source to query detailed information of images
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.Ecs.images(ids=[
        "image-cm9ssb4eqmhdas306zlp",
        "image-ybkzct2rtj4ay5rmlfc3",
    ])
    ```


    :param Sequence[str] ids: A list of Image IDs.
    :param str instance_type_id: The specification of  Instance.
    :param bool is_support_cloud_init: Whether the Image support cloud-init.
    :param str name_regex: A Name Regex of Image.
    :param str os_type: The operating system type of Image.
    :param str output_file: File name where to save data source results.
    :param Sequence[str] statuses: A list of Image status.
    :param str visibility: The visibility of Image.
    """
    __args__ = dict()
    __args__['ids'] = ids
    __args__['instanceTypeId'] = instance_type_id
    __args__['isSupportCloudInit'] = is_support_cloud_init
    __args__['nameRegex'] = name_regex
    __args__['osType'] = os_type
    __args__['outputFile'] = output_file
    __args__['statuses'] = statuses
    __args__['visibility'] = visibility
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('volcengine:Ecs/images:Images', __args__, opts=opts, typ=ImagesResult).value

    return AwaitableImagesResult(
        id=__ret__.id,
        ids=__ret__.ids,
        images=__ret__.images,
        instance_type_id=__ret__.instance_type_id,
        is_support_cloud_init=__ret__.is_support_cloud_init,
        name_regex=__ret__.name_regex,
        os_type=__ret__.os_type,
        output_file=__ret__.output_file,
        statuses=__ret__.statuses,
        total_count=__ret__.total_count,
        visibility=__ret__.visibility)


@_utilities.lift_output_func(images)
def images_output(ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                  instance_type_id: Optional[pulumi.Input[Optional[str]]] = None,
                  is_support_cloud_init: Optional[pulumi.Input[Optional[bool]]] = None,
                  name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                  os_type: Optional[pulumi.Input[Optional[str]]] = None,
                  output_file: Optional[pulumi.Input[Optional[str]]] = None,
                  statuses: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                  visibility: Optional[pulumi.Input[Optional[str]]] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[ImagesResult]:
    """
    Use this data source to query detailed information of images
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.Ecs.images(ids=[
        "image-cm9ssb4eqmhdas306zlp",
        "image-ybkzct2rtj4ay5rmlfc3",
    ])
    ```


    :param Sequence[str] ids: A list of Image IDs.
    :param str instance_type_id: The specification of  Instance.
    :param bool is_support_cloud_init: Whether the Image support cloud-init.
    :param str name_regex: A Name Regex of Image.
    :param str os_type: The operating system type of Image.
    :param str output_file: File name where to save data source results.
    :param Sequence[str] statuses: A list of Image status.
    :param str visibility: The visibility of Image.
    """
    ...
