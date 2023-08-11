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
from ._inputs import *

__all__ = [
    'NodesResult',
    'AwaitableNodesResult',
    'nodes',
    'nodes_output',
]

@pulumi.output_type
class NodesResult:
    """
    A collection of values returned by Nodes.
    """
    def __init__(__self__, cluster_ids=None, create_client_token=None, id=None, ids=None, name=None, name_regex=None, node_pool_ids=None, nodes=None, output_file=None, statuses=None, total_count=None, zone_ids=None):
        if cluster_ids and not isinstance(cluster_ids, list):
            raise TypeError("Expected argument 'cluster_ids' to be a list")
        pulumi.set(__self__, "cluster_ids", cluster_ids)
        if create_client_token and not isinstance(create_client_token, str):
            raise TypeError("Expected argument 'create_client_token' to be a str")
        pulumi.set(__self__, "create_client_token", create_client_token)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if node_pool_ids and not isinstance(node_pool_ids, list):
            raise TypeError("Expected argument 'node_pool_ids' to be a list")
        pulumi.set(__self__, "node_pool_ids", node_pool_ids)
        if nodes and not isinstance(nodes, list):
            raise TypeError("Expected argument 'nodes' to be a list")
        pulumi.set(__self__, "nodes", nodes)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if statuses and not isinstance(statuses, list):
            raise TypeError("Expected argument 'statuses' to be a list")
        pulumi.set(__self__, "statuses", statuses)
        if total_count and not isinstance(total_count, int):
            raise TypeError("Expected argument 'total_count' to be a int")
        pulumi.set(__self__, "total_count", total_count)
        if zone_ids and not isinstance(zone_ids, list):
            raise TypeError("Expected argument 'zone_ids' to be a list")
        pulumi.set(__self__, "zone_ids", zone_ids)

    @property
    @pulumi.getter(name="clusterIds")
    def cluster_ids(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "cluster_ids")

    @property
    @pulumi.getter(name="createClientToken")
    def create_client_token(self) -> Optional[str]:
        """
        The create client token of node.
        """
        return pulumi.get(self, "create_client_token")

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
    def name(self) -> Optional[str]:
        """
        The name of Node.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nameRegex")
    def name_regex(self) -> Optional[str]:
        return pulumi.get(self, "name_regex")

    @property
    @pulumi.getter(name="nodePoolIds")
    def node_pool_ids(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "node_pool_ids")

    @property
    @pulumi.getter
    def nodes(self) -> Sequence['outputs.NodesNodeResult']:
        """
        The collection of Node query.
        """
        return pulumi.get(self, "nodes")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter
    def statuses(self) -> Optional[Sequence['outputs.NodesStatusResult']]:
        return pulumi.get(self, "statuses")

    @property
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        """
        The total count of Node query.
        """
        return pulumi.get(self, "total_count")

    @property
    @pulumi.getter(name="zoneIds")
    def zone_ids(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "zone_ids")


class AwaitableNodesResult(NodesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return NodesResult(
            cluster_ids=self.cluster_ids,
            create_client_token=self.create_client_token,
            id=self.id,
            ids=self.ids,
            name=self.name,
            name_regex=self.name_regex,
            node_pool_ids=self.node_pool_ids,
            nodes=self.nodes,
            output_file=self.output_file,
            statuses=self.statuses,
            total_count=self.total_count,
            zone_ids=self.zone_ids)


def nodes(cluster_ids: Optional[Sequence[str]] = None,
          create_client_token: Optional[str] = None,
          ids: Optional[Sequence[str]] = None,
          name: Optional[str] = None,
          name_regex: Optional[str] = None,
          node_pool_ids: Optional[Sequence[str]] = None,
          output_file: Optional[str] = None,
          statuses: Optional[Sequence[pulumi.InputType['NodesStatusArgs']]] = None,
          zone_ids: Optional[Sequence[str]] = None,
          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableNodesResult:
    """
    Use this data source to query detailed information of vke nodes
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.vke.nodes(cluster_ids=[
            "c123",
            "c456",
        ],
        ids=["ncaa3e5mrsferqkomi190"],
        statuses=[
            volcengine.vke.NodesStatusArgs(
                conditions_type="Progressing",
                phase="Creating",
            ),
            volcengine.vke.NodesStatusArgs(
                conditions_type="Progressing123",
                phase="Creating123",
            ),
        ])
    ```


    :param Sequence[str] cluster_ids: A list of Cluster IDs.
    :param str create_client_token: The Create Client Token.
    :param Sequence[str] ids: A list of Node IDs.
    :param str name: The Name of Node.
    :param str name_regex: A Name Regex of Node.
    :param Sequence[str] node_pool_ids: The Node Pool IDs.
    :param str output_file: File name where to save data source results.
    :param Sequence[pulumi.InputType['NodesStatusArgs']] statuses: The Status of filter.
    :param Sequence[str] zone_ids: The Zone IDs.
    """
    __args__ = dict()
    __args__['clusterIds'] = cluster_ids
    __args__['createClientToken'] = create_client_token
    __args__['ids'] = ids
    __args__['name'] = name
    __args__['nameRegex'] = name_regex
    __args__['nodePoolIds'] = node_pool_ids
    __args__['outputFile'] = output_file
    __args__['statuses'] = statuses
    __args__['zoneIds'] = zone_ids
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('volcengine:vke/nodes:Nodes', __args__, opts=opts, typ=NodesResult).value

    return AwaitableNodesResult(
        cluster_ids=pulumi.get(__ret__, 'cluster_ids'),
        create_client_token=pulumi.get(__ret__, 'create_client_token'),
        id=pulumi.get(__ret__, 'id'),
        ids=pulumi.get(__ret__, 'ids'),
        name=pulumi.get(__ret__, 'name'),
        name_regex=pulumi.get(__ret__, 'name_regex'),
        node_pool_ids=pulumi.get(__ret__, 'node_pool_ids'),
        nodes=pulumi.get(__ret__, 'nodes'),
        output_file=pulumi.get(__ret__, 'output_file'),
        statuses=pulumi.get(__ret__, 'statuses'),
        total_count=pulumi.get(__ret__, 'total_count'),
        zone_ids=pulumi.get(__ret__, 'zone_ids'))


@_utilities.lift_output_func(nodes)
def nodes_output(cluster_ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                 create_client_token: Optional[pulumi.Input[Optional[str]]] = None,
                 ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                 name: Optional[pulumi.Input[Optional[str]]] = None,
                 name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                 node_pool_ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                 output_file: Optional[pulumi.Input[Optional[str]]] = None,
                 statuses: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['NodesStatusArgs']]]]] = None,
                 zone_ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[NodesResult]:
    """
    Use this data source to query detailed information of vke nodes
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.vke.nodes(cluster_ids=[
            "c123",
            "c456",
        ],
        ids=["ncaa3e5mrsferqkomi190"],
        statuses=[
            volcengine.vke.NodesStatusArgs(
                conditions_type="Progressing",
                phase="Creating",
            ),
            volcengine.vke.NodesStatusArgs(
                conditions_type="Progressing123",
                phase="Creating123",
            ),
        ])
    ```


    :param Sequence[str] cluster_ids: A list of Cluster IDs.
    :param str create_client_token: The Create Client Token.
    :param Sequence[str] ids: A list of Node IDs.
    :param str name: The Name of Node.
    :param str name_regex: A Name Regex of Node.
    :param Sequence[str] node_pool_ids: The Node Pool IDs.
    :param str output_file: File name where to save data source results.
    :param Sequence[pulumi.InputType['NodesStatusArgs']] statuses: The Status of filter.
    :param Sequence[str] zone_ids: The Zone IDs.
    """
    ...
