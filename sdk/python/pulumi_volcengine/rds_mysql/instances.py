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
    def __init__(__self__, charge_type=None, create_time_end=None, create_time_start=None, db_engine_version=None, id=None, instance_id=None, instance_name=None, instance_status=None, name_regex=None, output_file=None, rds_mysql_instances=None, total_count=None, zone_id=None):
        if charge_type and not isinstance(charge_type, str):
            raise TypeError("Expected argument 'charge_type' to be a str")
        pulumi.set(__self__, "charge_type", charge_type)
        if create_time_end and not isinstance(create_time_end, str):
            raise TypeError("Expected argument 'create_time_end' to be a str")
        pulumi.set(__self__, "create_time_end", create_time_end)
        if create_time_start and not isinstance(create_time_start, str):
            raise TypeError("Expected argument 'create_time_start' to be a str")
        pulumi.set(__self__, "create_time_start", create_time_start)
        if db_engine_version and not isinstance(db_engine_version, str):
            raise TypeError("Expected argument 'db_engine_version' to be a str")
        pulumi.set(__self__, "db_engine_version", db_engine_version)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if instance_name and not isinstance(instance_name, str):
            raise TypeError("Expected argument 'instance_name' to be a str")
        pulumi.set(__self__, "instance_name", instance_name)
        if instance_status and not isinstance(instance_status, str):
            raise TypeError("Expected argument 'instance_status' to be a str")
        pulumi.set(__self__, "instance_status", instance_status)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if rds_mysql_instances and not isinstance(rds_mysql_instances, list):
            raise TypeError("Expected argument 'rds_mysql_instances' to be a list")
        pulumi.set(__self__, "rds_mysql_instances", rds_mysql_instances)
        if total_count and not isinstance(total_count, int):
            raise TypeError("Expected argument 'total_count' to be a int")
        pulumi.set(__self__, "total_count", total_count)
        if zone_id and not isinstance(zone_id, str):
            raise TypeError("Expected argument 'zone_id' to be a str")
        pulumi.set(__self__, "zone_id", zone_id)

    @property
    @pulumi.getter(name="chargeType")
    def charge_type(self) -> Optional[str]:
        """
        Payment type. Value:
        PostPaid - Pay-As-You-Go
        PrePaid - Yearly and monthly (default).
        """
        return pulumi.get(self, "charge_type")

    @property
    @pulumi.getter(name="createTimeEnd")
    def create_time_end(self) -> Optional[str]:
        return pulumi.get(self, "create_time_end")

    @property
    @pulumi.getter(name="createTimeStart")
    def create_time_start(self) -> Optional[str]:
        return pulumi.get(self, "create_time_start")

    @property
    @pulumi.getter(name="dbEngineVersion")
    def db_engine_version(self) -> Optional[str]:
        """
        The engine version of the RDS instance.
        """
        return pulumi.get(self, "db_engine_version")

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
        Instance ID.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="instanceName")
    def instance_name(self) -> Optional[str]:
        """
        The name of the RDS instance.
        """
        return pulumi.get(self, "instance_name")

    @property
    @pulumi.getter(name="instanceStatus")
    def instance_status(self) -> Optional[str]:
        """
        The status of the RDS instance.
        """
        return pulumi.get(self, "instance_status")

    @property
    @pulumi.getter(name="nameRegex")
    def name_regex(self) -> Optional[str]:
        return pulumi.get(self, "name_regex")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="rdsMysqlInstances")
    def rds_mysql_instances(self) -> Sequence['outputs.InstancesRdsMysqlInstanceResult']:
        """
        The collection of RDS instance query.
        """
        return pulumi.get(self, "rds_mysql_instances")

    @property
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        """
        The total count of RDS instance query.
        """
        return pulumi.get(self, "total_count")

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> Optional[str]:
        """
        The available zone of the RDS instance.
        """
        return pulumi.get(self, "zone_id")


class AwaitableInstancesResult(InstancesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return InstancesResult(
            charge_type=self.charge_type,
            create_time_end=self.create_time_end,
            create_time_start=self.create_time_start,
            db_engine_version=self.db_engine_version,
            id=self.id,
            instance_id=self.instance_id,
            instance_name=self.instance_name,
            instance_status=self.instance_status,
            name_regex=self.name_regex,
            output_file=self.output_file,
            rds_mysql_instances=self.rds_mysql_instances,
            total_count=self.total_count,
            zone_id=self.zone_id)


def instances(charge_type: Optional[str] = None,
              create_time_end: Optional[str] = None,
              create_time_start: Optional[str] = None,
              db_engine_version: Optional[str] = None,
              instance_id: Optional[str] = None,
              instance_name: Optional[str] = None,
              instance_status: Optional[str] = None,
              name_regex: Optional[str] = None,
              output_file: Optional[str] = None,
              zone_id: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableInstancesResult:
    """
    Use this data source to query detailed information of rds mysql instances
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.rds_mysql.instances(instance_id="mysql-72da4258c2c7")
    ```


    :param str charge_type: The charge type of the RDS instance.
    :param str create_time_end: The end time of creating RDS instance.
    :param str create_time_start: The start time of creating RDS instance.
    :param str db_engine_version: The version of the RDS instance.
    :param str instance_id: The id of the RDS instance.
    :param str instance_name: The name of the RDS instance.
    :param str instance_status: The status of the RDS instance.
    :param str name_regex: A Name Regex of RDS instance.
    :param str output_file: File name where to save data source results.
    :param str zone_id: The available zone of the RDS instance.
    """
    __args__ = dict()
    __args__['chargeType'] = charge_type
    __args__['createTimeEnd'] = create_time_end
    __args__['createTimeStart'] = create_time_start
    __args__['dbEngineVersion'] = db_engine_version
    __args__['instanceId'] = instance_id
    __args__['instanceName'] = instance_name
    __args__['instanceStatus'] = instance_status
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['zoneId'] = zone_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('volcengine:rds_mysql/instances:Instances', __args__, opts=opts, typ=InstancesResult).value

    return AwaitableInstancesResult(
        charge_type=__ret__.charge_type,
        create_time_end=__ret__.create_time_end,
        create_time_start=__ret__.create_time_start,
        db_engine_version=__ret__.db_engine_version,
        id=__ret__.id,
        instance_id=__ret__.instance_id,
        instance_name=__ret__.instance_name,
        instance_status=__ret__.instance_status,
        name_regex=__ret__.name_regex,
        output_file=__ret__.output_file,
        rds_mysql_instances=__ret__.rds_mysql_instances,
        total_count=__ret__.total_count,
        zone_id=__ret__.zone_id)


@_utilities.lift_output_func(instances)
def instances_output(charge_type: Optional[pulumi.Input[Optional[str]]] = None,
                     create_time_end: Optional[pulumi.Input[Optional[str]]] = None,
                     create_time_start: Optional[pulumi.Input[Optional[str]]] = None,
                     db_engine_version: Optional[pulumi.Input[Optional[str]]] = None,
                     instance_id: Optional[pulumi.Input[Optional[str]]] = None,
                     instance_name: Optional[pulumi.Input[Optional[str]]] = None,
                     instance_status: Optional[pulumi.Input[Optional[str]]] = None,
                     name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                     output_file: Optional[pulumi.Input[Optional[str]]] = None,
                     zone_id: Optional[pulumi.Input[Optional[str]]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[InstancesResult]:
    """
    Use this data source to query detailed information of rds mysql instances
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.rds_mysql.instances(instance_id="mysql-72da4258c2c7")
    ```


    :param str charge_type: The charge type of the RDS instance.
    :param str create_time_end: The end time of creating RDS instance.
    :param str create_time_start: The start time of creating RDS instance.
    :param str db_engine_version: The version of the RDS instance.
    :param str instance_id: The id of the RDS instance.
    :param str instance_name: The name of the RDS instance.
    :param str instance_status: The status of the RDS instance.
    :param str name_regex: A Name Regex of RDS instance.
    :param str output_file: File name where to save data source results.
    :param str zone_id: The available zone of the RDS instance.
    """
    ...
