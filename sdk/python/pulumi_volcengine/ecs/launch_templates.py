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
    'LaunchTemplatesResult',
    'AwaitableLaunchTemplatesResult',
    'launch_templates',
    'launch_templates_output',
]

@pulumi.output_type
class LaunchTemplatesResult:
    """
    A collection of values returned by LaunchTemplates.
    """
    def __init__(__self__, id=None, ids=None, launch_template_names=None, launch_templates=None, name_regex=None, output_file=None, total_count=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if launch_template_names and not isinstance(launch_template_names, list):
            raise TypeError("Expected argument 'launch_template_names' to be a list")
        pulumi.set(__self__, "launch_template_names", launch_template_names)
        if launch_templates and not isinstance(launch_templates, list):
            raise TypeError("Expected argument 'launch_templates' to be a list")
        pulumi.set(__self__, "launch_templates", launch_templates)
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
    @pulumi.getter(name="launchTemplateNames")
    def launch_template_names(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "launch_template_names")

    @property
    @pulumi.getter(name="launchTemplates")
    def launch_templates(self) -> Sequence['outputs.LaunchTemplatesLaunchTemplateResult']:
        """
        The collection of launch templates.
        """
        return pulumi.get(self, "launch_templates")

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
        The total count of scaling policy query.
        """
        return pulumi.get(self, "total_count")


class AwaitableLaunchTemplatesResult(LaunchTemplatesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return LaunchTemplatesResult(
            id=self.id,
            ids=self.ids,
            launch_template_names=self.launch_template_names,
            launch_templates=self.launch_templates,
            name_regex=self.name_regex,
            output_file=self.output_file,
            total_count=self.total_count)


def launch_templates(ids: Optional[Sequence[str]] = None,
                     launch_template_names: Optional[Sequence[str]] = None,
                     name_regex: Optional[str] = None,
                     output_file: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableLaunchTemplatesResult:
    """
    Use this data source to query detailed information of ecs launch templates
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    foo_launch_template = volcengine.ecs.LaunchTemplate("fooLaunchTemplate",
        description="acc-test-desc",
        eip_bandwidth=1,
        eip_billing_type="PostPaidByBandwidth",
        eip_isp="ChinaMobile",
        host_name="tf-host-name",
        hpc_cluster_id="hpcCluster-l8u24ovdmoab6opf",
        image_id="image-ycjwwciuzy5pkh54xx8f",
        instance_charge_type="PostPaid",
        instance_name="tf-acc-name",
        instance_type_id="ecs.g1.large",
        key_pair_name="tf-key-pair",
        launch_template_name="tf-acc-template")
    foo_launch_templates = volcengine.ecs.launch_templates_output(ids=[foo_launch_template.id])
    ```


    :param Sequence[str] ids: A list of launch template ids.
    :param Sequence[str] launch_template_names: A list of launch template names.
    :param str name_regex: A Name Regex of scaling policy.
    :param str output_file: File name where to save data source results.
    """
    __args__ = dict()
    __args__['ids'] = ids
    __args__['launchTemplateNames'] = launch_template_names
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('volcengine:ecs/launchTemplates:LaunchTemplates', __args__, opts=opts, typ=LaunchTemplatesResult).value

    return AwaitableLaunchTemplatesResult(
        id=__ret__.id,
        ids=__ret__.ids,
        launch_template_names=__ret__.launch_template_names,
        launch_templates=__ret__.launch_templates,
        name_regex=__ret__.name_regex,
        output_file=__ret__.output_file,
        total_count=__ret__.total_count)


@_utilities.lift_output_func(launch_templates)
def launch_templates_output(ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                            launch_template_names: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                            name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                            output_file: Optional[pulumi.Input[Optional[str]]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[LaunchTemplatesResult]:
    """
    Use this data source to query detailed information of ecs launch templates
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    foo_launch_template = volcengine.ecs.LaunchTemplate("fooLaunchTemplate",
        description="acc-test-desc",
        eip_bandwidth=1,
        eip_billing_type="PostPaidByBandwidth",
        eip_isp="ChinaMobile",
        host_name="tf-host-name",
        hpc_cluster_id="hpcCluster-l8u24ovdmoab6opf",
        image_id="image-ycjwwciuzy5pkh54xx8f",
        instance_charge_type="PostPaid",
        instance_name="tf-acc-name",
        instance_type_id="ecs.g1.large",
        key_pair_name="tf-key-pair",
        launch_template_name="tf-acc-template")
    foo_launch_templates = volcengine.ecs.launch_templates_output(ids=[foo_launch_template.id])
    ```


    :param Sequence[str] ids: A list of launch template ids.
    :param Sequence[str] launch_template_names: A list of launch template names.
    :param str name_regex: A Name Regex of scaling policy.
    :param str output_file: File name where to save data source results.
    """
    ...
