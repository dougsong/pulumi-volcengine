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
    'DatabasesResult',
    'AwaitableDatabasesResult',
    'databases',
    'databases_output',
]

@pulumi.output_type
class DatabasesResult:
    """
    A collection of values returned by Databases.
    """
    def __init__(__self__, databases=None, db_name=None, id=None, instance_id=None, name_regex=None, output_file=None, total_count=None):
        if databases and not isinstance(databases, list):
            raise TypeError("Expected argument 'databases' to be a list")
        pulumi.set(__self__, "databases", databases)
        if db_name and not isinstance(db_name, str):
            raise TypeError("Expected argument 'db_name' to be a str")
        pulumi.set(__self__, "db_name", db_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
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
    def databases(self) -> Sequence['outputs.DatabasesDatabaseResult']:
        """
        The collection of RDS instance account query.
        """
        return pulumi.get(self, "databases")

    @property
    @pulumi.getter(name="dbName")
    def db_name(self) -> Optional[str]:
        """
        The name of the RDS database. This field supports fuzzy queries.
        """
        return pulumi.get(self, "db_name")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> str:
        return pulumi.get(self, "instance_id")

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
        The total count of RDS database query.
        """
        return pulumi.get(self, "total_count")


class AwaitableDatabasesResult(DatabasesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return DatabasesResult(
            databases=self.databases,
            db_name=self.db_name,
            id=self.id,
            instance_id=self.instance_id,
            name_regex=self.name_regex,
            output_file=self.output_file,
            total_count=self.total_count)


def databases(db_name: Optional[str] = None,
              instance_id: Optional[str] = None,
              name_regex: Optional[str] = None,
              output_file: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableDatabasesResult:
    """
    Use this data source to query detailed information of rds mysql databases
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.rds_mysql.databases(db_name="",
        instance_id="")
    ```


    :param str db_name: The name of the RDS database.
    :param str instance_id: The id of the RDS instance.
    :param str name_regex: A Name Regex of RDS database.
    :param str output_file: File name where to save data source results.
    """
    __args__ = dict()
    __args__['dbName'] = db_name
    __args__['instanceId'] = instance_id
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('volcengine:rds_mysql/databases:Databases', __args__, opts=opts, typ=DatabasesResult).value

    return AwaitableDatabasesResult(
        databases=__ret__.databases,
        db_name=__ret__.db_name,
        id=__ret__.id,
        instance_id=__ret__.instance_id,
        name_regex=__ret__.name_regex,
        output_file=__ret__.output_file,
        total_count=__ret__.total_count)


@_utilities.lift_output_func(databases)
def databases_output(db_name: Optional[pulumi.Input[Optional[str]]] = None,
                     instance_id: Optional[pulumi.Input[str]] = None,
                     name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                     output_file: Optional[pulumi.Input[Optional[str]]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[DatabasesResult]:
    """
    Use this data source to query detailed information of rds mysql databases
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.rds_mysql.databases(db_name="",
        instance_id="")
    ```


    :param str db_name: The name of the RDS database.
    :param str instance_id: The id of the RDS instance.
    :param str name_regex: A Name Regex of RDS database.
    :param str output_file: File name where to save data source results.
    """
    ...
