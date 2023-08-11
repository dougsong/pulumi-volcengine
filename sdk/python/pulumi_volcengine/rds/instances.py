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
    'InstancesResult',
    'AwaitableInstancesResult',
    'instances',
    'instances_output',
]

@pulumi.output_type
class InstancesResult:
    """
    A collection of values returned by Instances.
    """
    def __init__(__self__, create_end_time=None, create_start_time=None, id=None, instance_id=None, instance_status=None, instance_type=None, name_regex=None, output_file=None, rds_instances=None, region=None, total_count=None, zone=None):
        if create_end_time and not isinstance(create_end_time, str):
            raise TypeError("Expected argument 'create_end_time' to be a str")
        pulumi.set(__self__, "create_end_time", create_end_time)
        if create_start_time and not isinstance(create_start_time, str):
            raise TypeError("Expected argument 'create_start_time' to be a str")
        pulumi.set(__self__, "create_start_time", create_start_time)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if instance_status and not isinstance(instance_status, str):
            raise TypeError("Expected argument 'instance_status' to be a str")
        pulumi.set(__self__, "instance_status", instance_status)
        if instance_type and not isinstance(instance_type, str):
            raise TypeError("Expected argument 'instance_type' to be a str")
        pulumi.set(__self__, "instance_type", instance_type)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if rds_instances and not isinstance(rds_instances, list):
            raise TypeError("Expected argument 'rds_instances' to be a list")
        pulumi.set(__self__, "rds_instances", rds_instances)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if total_count and not isinstance(total_count, int):
            raise TypeError("Expected argument 'total_count' to be a int")
        pulumi.set(__self__, "total_count", total_count)
        if zone and not isinstance(zone, str):
            raise TypeError("Expected argument 'zone' to be a str")
        pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="createEndTime")
    def create_end_time(self) -> Optional[str]:
        return pulumi.get(self, "create_end_time")

    @property
    @pulumi.getter(name="createStartTime")
    def create_start_time(self) -> Optional[str]:
        return pulumi.get(self, "create_start_time")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[str]:
        """
        The ID of the RDS instance.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="instanceStatus")
    def instance_status(self) -> Optional[str]:
        """
        The status of the RDS instance.
        """
        return pulumi.get(self, "instance_status")

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> Optional[str]:
        """
        The type of the RDS instance.
        """
        return pulumi.get(self, "instance_type")

    @property
    @pulumi.getter(name="nameRegex")
    def name_regex(self) -> Optional[str]:
        return pulumi.get(self, "name_regex")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="rdsInstances")
    def rds_instances(self) -> Sequence['outputs.InstancesRdsInstanceResult']:
        """
        The collection of RDS instance query.
        """
        return pulumi.get(self, "rds_instances")

    @property
    @pulumi.getter
    def region(self) -> Optional[str]:
        """
        The region of the RDS instance.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        """
        The total count of RDS instance query.
        """
        return pulumi.get(self, "total_count")

    @property
    @pulumi.getter
    def zone(self) -> Optional[str]:
        """
        The available zone of the RDS instance.
        """
        return pulumi.get(self, "zone")


class AwaitableInstancesResult(InstancesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return InstancesResult(
            create_end_time=self.create_end_time,
            create_start_time=self.create_start_time,
            id=self.id,
            instance_id=self.instance_id,
            instance_status=self.instance_status,
            instance_type=self.instance_type,
            name_regex=self.name_regex,
            output_file=self.output_file,
            rds_instances=self.rds_instances,
            region=self.region,
            total_count=self.total_count,
            zone=self.zone)


def instances(create_end_time: Optional[str] = None,
              create_start_time: Optional[str] = None,
              instance_id: Optional[str] = None,
              instance_status: Optional[str] = None,
              instance_type: Optional[str] = None,
              name_regex: Optional[str] = None,
              output_file: Optional[str] = None,
              region: Optional[str] = None,
              zone: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableInstancesResult:
    """
    (Deprecated! Recommend use volcengine_rds_mysql_*** replace) Use this data source to query detailed information of rds instances
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.rds.instances(instance_id="mysql-0fdd3bab2e7c")
    ```


    :param str create_end_time: The end time of creating RDS instance.
    :param str create_start_time: The start time of creating RDS instance.
    :param str instance_id: The id of the RDS instance.
    :param str instance_status: The status of the RDS instance.
    :param str instance_type: The type of the RDS instance.
    :param str name_regex: A Name Regex of RDS instance.
    :param str output_file: File name where to save data source results.
    :param str region: The region of the RDS instance.
    :param str zone: The available zone of the RDS instance.
    """
    __args__ = dict()
    __args__['createEndTime'] = create_end_time
    __args__['createStartTime'] = create_start_time
    __args__['instanceId'] = instance_id
    __args__['instanceStatus'] = instance_status
    __args__['instanceType'] = instance_type
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['region'] = region
    __args__['zone'] = zone
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('volcengine:rds/instances:Instances', __args__, opts=opts, typ=InstancesResult).value

    return AwaitableInstancesResult(
        create_end_time=pulumi.get(__ret__, 'create_end_time'),
        create_start_time=pulumi.get(__ret__, 'create_start_time'),
        id=pulumi.get(__ret__, 'id'),
        instance_id=pulumi.get(__ret__, 'instance_id'),
        instance_status=pulumi.get(__ret__, 'instance_status'),
        instance_type=pulumi.get(__ret__, 'instance_type'),
        name_regex=pulumi.get(__ret__, 'name_regex'),
        output_file=pulumi.get(__ret__, 'output_file'),
        rds_instances=pulumi.get(__ret__, 'rds_instances'),
        region=pulumi.get(__ret__, 'region'),
        total_count=pulumi.get(__ret__, 'total_count'),
        zone=pulumi.get(__ret__, 'zone'))


@_utilities.lift_output_func(instances)
def instances_output(create_end_time: Optional[pulumi.Input[Optional[str]]] = None,
                     create_start_time: Optional[pulumi.Input[Optional[str]]] = None,
                     instance_id: Optional[pulumi.Input[Optional[str]]] = None,
                     instance_status: Optional[pulumi.Input[Optional[str]]] = None,
                     instance_type: Optional[pulumi.Input[Optional[str]]] = None,
                     name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                     output_file: Optional[pulumi.Input[Optional[str]]] = None,
                     region: Optional[pulumi.Input[Optional[str]]] = None,
                     zone: Optional[pulumi.Input[Optional[str]]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[InstancesResult]:
    """
    (Deprecated! Recommend use volcengine_rds_mysql_*** replace) Use this data source to query detailed information of rds instances
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.rds.instances(instance_id="mysql-0fdd3bab2e7c")
    ```


    :param str create_end_time: The end time of creating RDS instance.
    :param str create_start_time: The start time of creating RDS instance.
    :param str instance_id: The id of the RDS instance.
    :param str instance_status: The status of the RDS instance.
    :param str instance_type: The type of the RDS instance.
    :param str name_regex: A Name Regex of RDS instance.
    :param str output_file: File name where to save data source results.
    :param str region: The region of the RDS instance.
    :param str zone: The available zone of the RDS instance.
    """
    ...
