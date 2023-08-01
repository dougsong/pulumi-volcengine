# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = [
    'ProjectsResult',
    'AwaitableProjectsResult',
    'projects',
    'projects_output',
]

@pulumi.output_type
class ProjectsResult:
    """
    A collection of values returned by Projects.
    """
    def __init__(__self__, iam_project_name=None, id=None, is_full_name=None, name_regex=None, output_file=None, project_id=None, project_name=None, tags=None, tls_projects=None, total_count=None):
        if iam_project_name and not isinstance(iam_project_name, str):
            raise TypeError("Expected argument 'iam_project_name' to be a str")
        pulumi.set(__self__, "iam_project_name", iam_project_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if is_full_name and not isinstance(is_full_name, bool):
            raise TypeError("Expected argument 'is_full_name' to be a bool")
        pulumi.set(__self__, "is_full_name", is_full_name)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if project_name and not isinstance(project_name, str):
            raise TypeError("Expected argument 'project_name' to be a str")
        pulumi.set(__self__, "project_name", project_name)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if tls_projects and not isinstance(tls_projects, list):
            raise TypeError("Expected argument 'tls_projects' to be a list")
        pulumi.set(__self__, "tls_projects", tls_projects)
        if total_count and not isinstance(total_count, int):
            raise TypeError("Expected argument 'total_count' to be a int")
        pulumi.set(__self__, "total_count", total_count)

    @property
    @pulumi.getter(name="iamProjectName")
    def iam_project_name(self) -> Optional[str]:
        """
        The IAM project name of the tls project.
        """
        return pulumi.get(self, "iam_project_name")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="isFullName")
    def is_full_name(self) -> Optional[bool]:
        return pulumi.get(self, "is_full_name")

    @property
    @pulumi.getter(name="nameRegex")
    def name_regex(self) -> Optional[str]:
        return pulumi.get(self, "name_regex")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[str]:
        """
        The ID of the tls project.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="projectName")
    def project_name(self) -> Optional[str]:
        """
        The name of the tls project.
        """
        return pulumi.get(self, "project_name")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.ProjectsTagResult']]:
        """
        Tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tlsProjects")
    def tls_projects(self) -> Sequence['outputs.ProjectsTlsProjectResult']:
        """
        The collection of tls project query.
        """
        return pulumi.get(self, "tls_projects")

    @property
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        """
        The total count of tls project query.
        """
        return pulumi.get(self, "total_count")


class AwaitableProjectsResult(ProjectsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ProjectsResult(
            iam_project_name=self.iam_project_name,
            id=self.id,
            is_full_name=self.is_full_name,
            name_regex=self.name_regex,
            output_file=self.output_file,
            project_id=self.project_id,
            project_name=self.project_name,
            tags=self.tags,
            tls_projects=self.tls_projects,
            total_count=self.total_count)


def projects(iam_project_name: Optional[str] = None,
             is_full_name: Optional[bool] = None,
             name_regex: Optional[str] = None,
             output_file: Optional[str] = None,
             project_id: Optional[str] = None,
             project_name: Optional[str] = None,
             tags: Optional[Sequence[pulumi.InputType['ProjectsTagArgs']]] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableProjectsResult:
    """
    Use this data source to query detailed information of tls projects
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.tls.projects(project_id="e020c978-4f05-40e1-9167-0113d3ef****")
    ```


    :param str iam_project_name: The IAM project name of the tls project.
    :param bool is_full_name: Whether to match accurately when filtering based on ProjectName.
    :param str name_regex: A Name Regex of tls project.
    :param str output_file: File name where to save data source results.
    :param str project_id: The id of tls project. This field supports fuzzy queries. It is not supported to specify both ProjectName and ProjectId at the same time.
    :param str project_name: The name of tls project. This field supports fuzzy queries. It is not supported to specify both ProjectName and ProjectId at the same time.
    :param Sequence[pulumi.InputType['ProjectsTagArgs']] tags: Tags.
    """
    __args__ = dict()
    __args__['iamProjectName'] = iam_project_name
    __args__['isFullName'] = is_full_name
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['projectId'] = project_id
    __args__['projectName'] = project_name
    __args__['tags'] = tags
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('volcengine:tls/projects:Projects', __args__, opts=opts, typ=ProjectsResult).value

    return AwaitableProjectsResult(
        iam_project_name=__ret__.iam_project_name,
        id=__ret__.id,
        is_full_name=__ret__.is_full_name,
        name_regex=__ret__.name_regex,
        output_file=__ret__.output_file,
        project_id=__ret__.project_id,
        project_name=__ret__.project_name,
        tags=__ret__.tags,
        tls_projects=__ret__.tls_projects,
        total_count=__ret__.total_count)


@_utilities.lift_output_func(projects)
def projects_output(iam_project_name: Optional[pulumi.Input[Optional[str]]] = None,
                    is_full_name: Optional[pulumi.Input[Optional[bool]]] = None,
                    name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                    output_file: Optional[pulumi.Input[Optional[str]]] = None,
                    project_id: Optional[pulumi.Input[Optional[str]]] = None,
                    project_name: Optional[pulumi.Input[Optional[str]]] = None,
                    tags: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['ProjectsTagArgs']]]]] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[ProjectsResult]:
    """
    Use this data source to query detailed information of tls projects
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.tls.projects(project_id="e020c978-4f05-40e1-9167-0113d3ef****")
    ```


    :param str iam_project_name: The IAM project name of the tls project.
    :param bool is_full_name: Whether to match accurately when filtering based on ProjectName.
    :param str name_regex: A Name Regex of tls project.
    :param str output_file: File name where to save data source results.
    :param str project_id: The id of tls project. This field supports fuzzy queries. It is not supported to specify both ProjectName and ProjectId at the same time.
    :param str project_name: The name of tls project. This field supports fuzzy queries. It is not supported to specify both ProjectName and ProjectId at the same time.
    :param Sequence[pulumi.InputType['ProjectsTagArgs']] tags: Tags.
    """
    ...
