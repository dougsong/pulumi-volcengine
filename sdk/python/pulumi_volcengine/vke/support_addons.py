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
    'SupportAddonsResult',
    'AwaitableSupportAddonsResult',
    'support_addons',
    'support_addons_output',
]

@pulumi.output_type
class SupportAddonsResult:
    """
    A collection of values returned by SupportAddons.
    """
    def __init__(__self__, addons=None, categories=None, deploy_modes=None, deploy_node_types=None, id=None, kubernetes_versions=None, name=None, necessaries=None, output_file=None, pod_network_modes=None, total_count=None):
        if addons and not isinstance(addons, list):
            raise TypeError("Expected argument 'addons' to be a list")
        pulumi.set(__self__, "addons", addons)
        if categories and not isinstance(categories, list):
            raise TypeError("Expected argument 'categories' to be a list")
        pulumi.set(__self__, "categories", categories)
        if deploy_modes and not isinstance(deploy_modes, list):
            raise TypeError("Expected argument 'deploy_modes' to be a list")
        pulumi.set(__self__, "deploy_modes", deploy_modes)
        if deploy_node_types and not isinstance(deploy_node_types, list):
            raise TypeError("Expected argument 'deploy_node_types' to be a list")
        pulumi.set(__self__, "deploy_node_types", deploy_node_types)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if kubernetes_versions and not isinstance(kubernetes_versions, list):
            raise TypeError("Expected argument 'kubernetes_versions' to be a list")
        pulumi.set(__self__, "kubernetes_versions", kubernetes_versions)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if necessaries and not isinstance(necessaries, list):
            raise TypeError("Expected argument 'necessaries' to be a list")
        pulumi.set(__self__, "necessaries", necessaries)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if pod_network_modes and not isinstance(pod_network_modes, list):
            raise TypeError("Expected argument 'pod_network_modes' to be a list")
        pulumi.set(__self__, "pod_network_modes", pod_network_modes)
        if total_count and not isinstance(total_count, int):
            raise TypeError("Expected argument 'total_count' to be a int")
        pulumi.set(__self__, "total_count", total_count)

    @property
    @pulumi.getter
    def addons(self) -> Sequence['outputs.SupportAddonsAddonResult']:
        """
        The collection of addons query.
        """
        return pulumi.get(self, "addons")

    @property
    @pulumi.getter
    def categories(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "categories")

    @property
    @pulumi.getter(name="deployModes")
    def deploy_modes(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "deploy_modes")

    @property
    @pulumi.getter(name="deployNodeTypes")
    def deploy_node_types(self) -> Optional[Sequence[str]]:
        """
        The deploy node types.
        """
        return pulumi.get(self, "deploy_node_types")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="kubernetesVersions")
    def kubernetes_versions(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "kubernetes_versions")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The name of addon.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def necessaries(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "necessaries")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="podNetworkModes")
    def pod_network_modes(self) -> Optional[Sequence[str]]:
        """
        The network modes of pod.
        """
        return pulumi.get(self, "pod_network_modes")

    @property
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        """
        The total count of addons query.
        """
        return pulumi.get(self, "total_count")


class AwaitableSupportAddonsResult(SupportAddonsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return SupportAddonsResult(
            addons=self.addons,
            categories=self.categories,
            deploy_modes=self.deploy_modes,
            deploy_node_types=self.deploy_node_types,
            id=self.id,
            kubernetes_versions=self.kubernetes_versions,
            name=self.name,
            necessaries=self.necessaries,
            output_file=self.output_file,
            pod_network_modes=self.pod_network_modes,
            total_count=self.total_count)


def support_addons(categories: Optional[Sequence[str]] = None,
                   deploy_modes: Optional[Sequence[str]] = None,
                   deploy_node_types: Optional[Sequence[str]] = None,
                   kubernetes_versions: Optional[Sequence[str]] = None,
                   name: Optional[str] = None,
                   necessaries: Optional[Sequence[str]] = None,
                   output_file: Optional[str] = None,
                   pod_network_modes: Optional[Sequence[str]] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableSupportAddonsResult:
    """
    Use this data source to query detailed information of vke support addons
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.vke.support_addons(categories=["Monitor"],
        name="metrics-server")
    ```


    :param Sequence[str] categories: The categories of addons, the value is `Storage` or `Network` or `Monitor` or `Scheduler` or `Dns` or `Security` or `Gpu` or `Image`.
    :param Sequence[str] deploy_modes: The deploy model, the value is `Managed` or `Unmanaged`.
    :param Sequence[str] deploy_node_types: The deploy node types, the value is `Node` or `VirtualNode`. Only effected when deploy_mode is `Unmanaged`.
    :param Sequence[str] kubernetes_versions: A list of Kubernetes Versions.
    :param str name: The name of the addon.
    :param Sequence[str] necessaries: The necessaries of addons, the value is `Required` or `Recommended` or `OnDemand`.
    :param str output_file: File name where to save data source results.
    :param Sequence[str] pod_network_modes: The container network model, the value is `Flannel` or `VpcCniShared`. Flannel: Flannel network model, an independent Underlay container network solution, combined with the global routing capability of VPC, to achieve a high-performance network experience for the cluster. VpcCniShared: VPC-CNI network model, an Underlay container network solution based on the ENI of the private network elastic network card, with high network communication performance.
    """
    __args__ = dict()
    __args__['categories'] = categories
    __args__['deployModes'] = deploy_modes
    __args__['deployNodeTypes'] = deploy_node_types
    __args__['kubernetesVersions'] = kubernetes_versions
    __args__['name'] = name
    __args__['necessaries'] = necessaries
    __args__['outputFile'] = output_file
    __args__['podNetworkModes'] = pod_network_modes
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('volcengine:vke/supportAddons:SupportAddons', __args__, opts=opts, typ=SupportAddonsResult).value

    return AwaitableSupportAddonsResult(
        addons=pulumi.get(__ret__, 'addons'),
        categories=pulumi.get(__ret__, 'categories'),
        deploy_modes=pulumi.get(__ret__, 'deploy_modes'),
        deploy_node_types=pulumi.get(__ret__, 'deploy_node_types'),
        id=pulumi.get(__ret__, 'id'),
        kubernetes_versions=pulumi.get(__ret__, 'kubernetes_versions'),
        name=pulumi.get(__ret__, 'name'),
        necessaries=pulumi.get(__ret__, 'necessaries'),
        output_file=pulumi.get(__ret__, 'output_file'),
        pod_network_modes=pulumi.get(__ret__, 'pod_network_modes'),
        total_count=pulumi.get(__ret__, 'total_count'))


@_utilities.lift_output_func(support_addons)
def support_addons_output(categories: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                          deploy_modes: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                          deploy_node_types: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                          kubernetes_versions: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                          name: Optional[pulumi.Input[Optional[str]]] = None,
                          necessaries: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                          output_file: Optional[pulumi.Input[Optional[str]]] = None,
                          pod_network_modes: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[SupportAddonsResult]:
    """
    Use this data source to query detailed information of vke support addons
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.vke.support_addons(categories=["Monitor"],
        name="metrics-server")
    ```


    :param Sequence[str] categories: The categories of addons, the value is `Storage` or `Network` or `Monitor` or `Scheduler` or `Dns` or `Security` or `Gpu` or `Image`.
    :param Sequence[str] deploy_modes: The deploy model, the value is `Managed` or `Unmanaged`.
    :param Sequence[str] deploy_node_types: The deploy node types, the value is `Node` or `VirtualNode`. Only effected when deploy_mode is `Unmanaged`.
    :param Sequence[str] kubernetes_versions: A list of Kubernetes Versions.
    :param str name: The name of the addon.
    :param Sequence[str] necessaries: The necessaries of addons, the value is `Required` or `Recommended` or `OnDemand`.
    :param str output_file: File name where to save data source results.
    :param Sequence[str] pod_network_modes: The container network model, the value is `Flannel` or `VpcCniShared`. Flannel: Flannel network model, an independent Underlay container network solution, combined with the global routing capability of VPC, to achieve a high-performance network experience for the cluster. VpcCniShared: VPC-CNI network model, an Underlay container network solution based on the ENI of the private network elastic network card, with high network communication performance.
    """
    ...
