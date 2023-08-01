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
    'TagsResult',
    'AwaitableTagsResult',
    'tags',
    'tags_output',
]

@pulumi.output_type
class TagsResult:
    """
    A collection of values returned by Tags.
    """
    def __init__(__self__, id=None, names=None, namespace=None, output_file=None, registry=None, repository=None, tags=None, total_count=None, types=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        pulumi.set(__self__, "names", names)
        if namespace and not isinstance(namespace, str):
            raise TypeError("Expected argument 'namespace' to be a str")
        pulumi.set(__self__, "namespace", namespace)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if registry and not isinstance(registry, str):
            raise TypeError("Expected argument 'registry' to be a str")
        pulumi.set(__self__, "registry", registry)
        if repository and not isinstance(repository, str):
            raise TypeError("Expected argument 'repository' to be a str")
        pulumi.set(__self__, "repository", repository)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if total_count and not isinstance(total_count, int):
            raise TypeError("Expected argument 'total_count' to be a int")
        pulumi.set(__self__, "total_count", total_count)
        if types and not isinstance(types, list):
            raise TypeError("Expected argument 'types' to be a list")
        pulumi.set(__self__, "types", types)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def names(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "names")

    @property
    @pulumi.getter
    def namespace(self) -> str:
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter
    def registry(self) -> str:
        return pulumi.get(self, "registry")

    @property
    @pulumi.getter
    def repository(self) -> str:
        return pulumi.get(self, "repository")

    @property
    @pulumi.getter
    def tags(self) -> Sequence['outputs.TagsTagResult']:
        """
        The collection of repository query.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        """
        The total count of tag query.
        """
        return pulumi.get(self, "total_count")

    @property
    @pulumi.getter
    def types(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "types")


class AwaitableTagsResult(TagsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return TagsResult(
            id=self.id,
            names=self.names,
            namespace=self.namespace,
            output_file=self.output_file,
            registry=self.registry,
            repository=self.repository,
            tags=self.tags,
            total_count=self.total_count,
            types=self.types)


def tags(names: Optional[Sequence[str]] = None,
         namespace: Optional[str] = None,
         output_file: Optional[str] = None,
         registry: Optional[str] = None,
         repository: Optional[str] = None,
         types: Optional[Sequence[str]] = None,
         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableTagsResult:
    """
    Use this data source to query detailed information of cr tags
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    foo = volcengine.cr.tags(namespace="test",
        registry="enterprise-1",
        repository="repo",
        types=["Image"])
    ```


    :param Sequence[str] names: The list of instance names.
    :param str namespace: The CR namespace.
    :param str output_file: File name where to save data source results.
    :param str registry: The CR instance name.
    :param str repository: The repository name.
    :param Sequence[str] types: The list of OCI product tag type.
    """
    __args__ = dict()
    __args__['names'] = names
    __args__['namespace'] = namespace
    __args__['outputFile'] = output_file
    __args__['registry'] = registry
    __args__['repository'] = repository
    __args__['types'] = types
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('volcengine:cr/tags:Tags', __args__, opts=opts, typ=TagsResult).value

    return AwaitableTagsResult(
        id=__ret__.id,
        names=__ret__.names,
        namespace=__ret__.namespace,
        output_file=__ret__.output_file,
        registry=__ret__.registry,
        repository=__ret__.repository,
        tags=__ret__.tags,
        total_count=__ret__.total_count,
        types=__ret__.types)


@_utilities.lift_output_func(tags)
def tags_output(names: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                namespace: Optional[pulumi.Input[str]] = None,
                output_file: Optional[pulumi.Input[Optional[str]]] = None,
                registry: Optional[pulumi.Input[str]] = None,
                repository: Optional[pulumi.Input[str]] = None,
                types: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[TagsResult]:
    """
    Use this data source to query detailed information of cr tags
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    foo = volcengine.cr.tags(namespace="test",
        registry="enterprise-1",
        repository="repo",
        types=["Image"])
    ```


    :param Sequence[str] names: The list of instance names.
    :param str namespace: The CR namespace.
    :param str output_file: File name where to save data source results.
    :param str registry: The CR instance name.
    :param str repository: The repository name.
    :param Sequence[str] types: The list of OCI product tag type.
    """
    ...
