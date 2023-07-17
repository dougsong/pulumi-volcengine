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
    'AccountPrivilegeDbPrivilege',
    'AccountsRdsAccountResult',
    'AccountsRdsAccountDbPrivilegeResult',
    'DatabasesRdsDatabaseResult',
    'InstanceConnectionInfo',
    'InstancesRdsInstanceResult',
    'InstancesRdsInstanceConnectionInfoResult',
    'InstancesRdsInstanceInstanceSpecResult',
    'IpListsRdsIpListResult',
    'ParameterTemplateTemplateParam',
    'ParameterTemplatesRdsParameterTemplateResult',
    'ParameterTemplatesRdsParameterTemplateTemplateParamResult',
]

@pulumi.output_type
class AccountPrivilegeDbPrivilege(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "accountPrivilege":
            suggest = "account_privilege"
        elif key == "dbName":
            suggest = "db_name"
        elif key == "accountPrivilegeStr":
            suggest = "account_privilege_str"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in AccountPrivilegeDbPrivilege. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        AccountPrivilegeDbPrivilege.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        AccountPrivilegeDbPrivilege.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 account_privilege: str,
                 db_name: str,
                 account_privilege_str: Optional[str] = None):
        """
        :param str account_privilege: The privilege type of the account.
        :param str db_name: The name of database.
        :param str account_privilege_str: The privilege string of the account.
        """
        pulumi.set(__self__, "account_privilege", account_privilege)
        pulumi.set(__self__, "db_name", db_name)
        if account_privilege_str is not None:
            pulumi.set(__self__, "account_privilege_str", account_privilege_str)

    @property
    @pulumi.getter(name="accountPrivilege")
    def account_privilege(self) -> str:
        """
        The privilege type of the account.
        """
        return pulumi.get(self, "account_privilege")

    @property
    @pulumi.getter(name="dbName")
    def db_name(self) -> str:
        """
        The name of database.
        """
        return pulumi.get(self, "db_name")

    @property
    @pulumi.getter(name="accountPrivilegeStr")
    def account_privilege_str(self) -> Optional[str]:
        """
        The privilege string of the account.
        """
        return pulumi.get(self, "account_privilege_str")


@pulumi.output_type
class AccountsRdsAccountResult(dict):
    def __init__(__self__, *,
                 account_name: str,
                 account_status: str,
                 account_type: str,
                 db_privileges: Sequence['outputs.AccountsRdsAccountDbPrivilegeResult'],
                 id: str):
        """
        :param str account_name: The name of the database account.
        :param str account_status: The status of the database account.
        :param str account_type: The type of the database account.
        :param Sequence['AccountsRdsAccountDbPrivilegeArgs'] db_privileges: The privilege detail list of RDS instance account.
        :param str id: The ID of the RDS instance account.
        """
        pulumi.set(__self__, "account_name", account_name)
        pulumi.set(__self__, "account_status", account_status)
        pulumi.set(__self__, "account_type", account_type)
        pulumi.set(__self__, "db_privileges", db_privileges)
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter(name="accountName")
    def account_name(self) -> str:
        """
        The name of the database account.
        """
        return pulumi.get(self, "account_name")

    @property
    @pulumi.getter(name="accountStatus")
    def account_status(self) -> str:
        """
        The status of the database account.
        """
        return pulumi.get(self, "account_status")

    @property
    @pulumi.getter(name="accountType")
    def account_type(self) -> str:
        """
        The type of the database account.
        """
        return pulumi.get(self, "account_type")

    @property
    @pulumi.getter(name="dbPrivileges")
    def db_privileges(self) -> Sequence['outputs.AccountsRdsAccountDbPrivilegeResult']:
        """
        The privilege detail list of RDS instance account.
        """
        return pulumi.get(self, "db_privileges")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of the RDS instance account.
        """
        return pulumi.get(self, "id")


@pulumi.output_type
class AccountsRdsAccountDbPrivilegeResult(dict):
    def __init__(__self__, *,
                 account_privilege: str,
                 account_privilege_str: str,
                 db_name: str):
        """
        :param str account_privilege: The privilege type of the account.
        :param str account_privilege_str: The privilege string of the account.
        :param str db_name: The name of database.
        """
        pulumi.set(__self__, "account_privilege", account_privilege)
        pulumi.set(__self__, "account_privilege_str", account_privilege_str)
        pulumi.set(__self__, "db_name", db_name)

    @property
    @pulumi.getter(name="accountPrivilege")
    def account_privilege(self) -> str:
        """
        The privilege type of the account.
        """
        return pulumi.get(self, "account_privilege")

    @property
    @pulumi.getter(name="accountPrivilegeStr")
    def account_privilege_str(self) -> str:
        """
        The privilege string of the account.
        """
        return pulumi.get(self, "account_privilege_str")

    @property
    @pulumi.getter(name="dbName")
    def db_name(self) -> str:
        """
        The name of database.
        """
        return pulumi.get(self, "db_name")


@pulumi.output_type
class DatabasesRdsDatabaseResult(dict):
    def __init__(__self__, *,
                 account_names: str,
                 character_set_name: str,
                 db_name: str,
                 db_status: str,
                 id: str):
        """
        :param str account_names: The account names of the RDS database.
        :param str character_set_name: The character set of the RDS database.
        :param str db_name: The name of the RDS database.
        :param str db_status: The status of the RDS database.
        :param str id: The ID of the RDS database.
        """
        pulumi.set(__self__, "account_names", account_names)
        pulumi.set(__self__, "character_set_name", character_set_name)
        pulumi.set(__self__, "db_name", db_name)
        pulumi.set(__self__, "db_status", db_status)
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter(name="accountNames")
    def account_names(self) -> str:
        """
        The account names of the RDS database.
        """
        return pulumi.get(self, "account_names")

    @property
    @pulumi.getter(name="characterSetName")
    def character_set_name(self) -> str:
        """
        The character set of the RDS database.
        """
        return pulumi.get(self, "character_set_name")

    @property
    @pulumi.getter(name="dbName")
    def db_name(self) -> str:
        """
        The name of the RDS database.
        """
        return pulumi.get(self, "db_name")

    @property
    @pulumi.getter(name="dbStatus")
    def db_status(self) -> str:
        """
        The status of the RDS database.
        """
        return pulumi.get(self, "db_status")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of the RDS database.
        """
        return pulumi.get(self, "id")


@pulumi.output_type
class InstanceConnectionInfo(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "enableReadOnly":
            suggest = "enable_read_only"
        elif key == "enableReadWriteSplitting":
            suggest = "enable_read_write_splitting"
        elif key == "internalDomain":
            suggest = "internal_domain"
        elif key == "internalPort":
            suggest = "internal_port"
        elif key == "publicDomain":
            suggest = "public_domain"
        elif key == "publicPort":
            suggest = "public_port"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in InstanceConnectionInfo. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        InstanceConnectionInfo.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        InstanceConnectionInfo.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 enable_read_only: Optional[str] = None,
                 enable_read_write_splitting: Optional[str] = None,
                 internal_domain: Optional[str] = None,
                 internal_port: Optional[str] = None,
                 public_domain: Optional[str] = None,
                 public_port: Optional[str] = None):
        """
        :param str enable_read_only: Whether global read-only is enabled.
        :param str enable_read_write_splitting: Whether read-write separation is enabled.
        :param str internal_domain: The internal domain of the RDS instance.
        :param str internal_port: The interval port of the RDS instance.
        :param str public_domain: The public domain of the RDS instance.
        :param str public_port: The public port of the RDS instance.
        """
        if enable_read_only is not None:
            pulumi.set(__self__, "enable_read_only", enable_read_only)
        if enable_read_write_splitting is not None:
            pulumi.set(__self__, "enable_read_write_splitting", enable_read_write_splitting)
        if internal_domain is not None:
            pulumi.set(__self__, "internal_domain", internal_domain)
        if internal_port is not None:
            pulumi.set(__self__, "internal_port", internal_port)
        if public_domain is not None:
            pulumi.set(__self__, "public_domain", public_domain)
        if public_port is not None:
            pulumi.set(__self__, "public_port", public_port)

    @property
    @pulumi.getter(name="enableReadOnly")
    def enable_read_only(self) -> Optional[str]:
        """
        Whether global read-only is enabled.
        """
        return pulumi.get(self, "enable_read_only")

    @property
    @pulumi.getter(name="enableReadWriteSplitting")
    def enable_read_write_splitting(self) -> Optional[str]:
        """
        Whether read-write separation is enabled.
        """
        return pulumi.get(self, "enable_read_write_splitting")

    @property
    @pulumi.getter(name="internalDomain")
    def internal_domain(self) -> Optional[str]:
        """
        The internal domain of the RDS instance.
        """
        return pulumi.get(self, "internal_domain")

    @property
    @pulumi.getter(name="internalPort")
    def internal_port(self) -> Optional[str]:
        """
        The interval port of the RDS instance.
        """
        return pulumi.get(self, "internal_port")

    @property
    @pulumi.getter(name="publicDomain")
    def public_domain(self) -> Optional[str]:
        """
        The public domain of the RDS instance.
        """
        return pulumi.get(self, "public_domain")

    @property
    @pulumi.getter(name="publicPort")
    def public_port(self) -> Optional[str]:
        """
        The public port of the RDS instance.
        """
        return pulumi.get(self, "public_port")


@pulumi.output_type
class InstancesRdsInstanceResult(dict):
    def __init__(__self__, *,
                 charge_status: str,
                 charge_type: str,
                 connection_info: 'outputs.InstancesRdsInstanceConnectionInfoResult',
                 create_time: str,
                 db_engine: str,
                 db_engine_version: str,
                 id: str,
                 instance_id: str,
                 instance_name: str,
                 instance_spec: 'outputs.InstancesRdsInstanceInstanceSpecResult',
                 instance_status: str,
                 instance_type: str,
                 region: str,
                 storage_space_gb: int,
                 update_time: str,
                 vpc_id: str,
                 zone: str,
                 read_only_instance_ids: Optional[Sequence[str]] = None):
        """
        :param str charge_status: The charge status of the RDS instance.
        :param str charge_type: The charge type of the RDS instance.
        :param 'InstancesRdsInstanceConnectionInfoArgs' connection_info: The connection info ot the RDS instance.
        :param str create_time: The create time of the RDS instance.
        :param str db_engine: The engine of the RDS instance.
        :param str db_engine_version: The engine version of the RDS instance.
        :param str id: The ID of the RDS instance.
        :param str instance_id: The id of the RDS instance.
        :param str instance_name: The name of the RDS instance.
        :param 'InstancesRdsInstanceInstanceSpecArgs' instance_spec: The spec type detail of RDS instance.
        :param str instance_status: The status of the RDS instance.
        :param str instance_type: The type of the RDS instance.
        :param str region: The region of the RDS instance.
        :param int storage_space_gb: The total storage GB of the RDS instance.
        :param str update_time: The update time of the RDS instance.
        :param str vpc_id: The vpc ID of the RDS instance.
        :param str zone: The available zone of the RDS instance.
        """
        pulumi.set(__self__, "charge_status", charge_status)
        pulumi.set(__self__, "charge_type", charge_type)
        pulumi.set(__self__, "connection_info", connection_info)
        pulumi.set(__self__, "create_time", create_time)
        pulumi.set(__self__, "db_engine", db_engine)
        pulumi.set(__self__, "db_engine_version", db_engine_version)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "instance_name", instance_name)
        pulumi.set(__self__, "instance_spec", instance_spec)
        pulumi.set(__self__, "instance_status", instance_status)
        pulumi.set(__self__, "instance_type", instance_type)
        pulumi.set(__self__, "region", region)
        pulumi.set(__self__, "storage_space_gb", storage_space_gb)
        pulumi.set(__self__, "update_time", update_time)
        pulumi.set(__self__, "vpc_id", vpc_id)
        pulumi.set(__self__, "zone", zone)
        if read_only_instance_ids is not None:
            pulumi.set(__self__, "read_only_instance_ids", read_only_instance_ids)

    @property
    @pulumi.getter(name="chargeStatus")
    def charge_status(self) -> str:
        """
        The charge status of the RDS instance.
        """
        return pulumi.get(self, "charge_status")

    @property
    @pulumi.getter(name="chargeType")
    def charge_type(self) -> str:
        """
        The charge type of the RDS instance.
        """
        return pulumi.get(self, "charge_type")

    @property
    @pulumi.getter(name="connectionInfo")
    def connection_info(self) -> 'outputs.InstancesRdsInstanceConnectionInfoResult':
        """
        The connection info ot the RDS instance.
        """
        return pulumi.get(self, "connection_info")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The create time of the RDS instance.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="dbEngine")
    def db_engine(self) -> str:
        """
        The engine of the RDS instance.
        """
        return pulumi.get(self, "db_engine")

    @property
    @pulumi.getter(name="dbEngineVersion")
    def db_engine_version(self) -> str:
        """
        The engine version of the RDS instance.
        """
        return pulumi.get(self, "db_engine_version")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of the RDS instance.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> str:
        """
        The id of the RDS instance.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="instanceName")
    def instance_name(self) -> str:
        """
        The name of the RDS instance.
        """
        return pulumi.get(self, "instance_name")

    @property
    @pulumi.getter(name="instanceSpec")
    def instance_spec(self) -> 'outputs.InstancesRdsInstanceInstanceSpecResult':
        """
        The spec type detail of RDS instance.
        """
        return pulumi.get(self, "instance_spec")

    @property
    @pulumi.getter(name="instanceStatus")
    def instance_status(self) -> str:
        """
        The status of the RDS instance.
        """
        return pulumi.get(self, "instance_status")

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> str:
        """
        The type of the RDS instance.
        """
        return pulumi.get(self, "instance_type")

    @property
    @pulumi.getter
    def region(self) -> str:
        """
        The region of the RDS instance.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="storageSpaceGb")
    def storage_space_gb(self) -> int:
        """
        The total storage GB of the RDS instance.
        """
        return pulumi.get(self, "storage_space_gb")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        The update time of the RDS instance.
        """
        return pulumi.get(self, "update_time")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> str:
        """
        The vpc ID of the RDS instance.
        """
        return pulumi.get(self, "vpc_id")

    @property
    @pulumi.getter
    def zone(self) -> str:
        """
        The available zone of the RDS instance.
        """
        return pulumi.get(self, "zone")

    @property
    @pulumi.getter(name="readOnlyInstanceIds")
    def read_only_instance_ids(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "read_only_instance_ids")


@pulumi.output_type
class InstancesRdsInstanceConnectionInfoResult(dict):
    def __init__(__self__, *,
                 enable_read_only: str,
                 enable_read_write_splitting: str,
                 internal_domain: str,
                 internal_port: str,
                 public_domain: str,
                 public_port: str):
        """
        :param str enable_read_only: Whether global read-only is enabled.
        :param str enable_read_write_splitting: Whether read-write separation is enabled.
        :param str internal_domain: The internal domain of the RDS instance.
        :param str internal_port: The interval port of the RDS instance.
        :param str public_domain: The public domain of the RDS instance.
        :param str public_port: The public port of the RDS instance.
        """
        pulumi.set(__self__, "enable_read_only", enable_read_only)
        pulumi.set(__self__, "enable_read_write_splitting", enable_read_write_splitting)
        pulumi.set(__self__, "internal_domain", internal_domain)
        pulumi.set(__self__, "internal_port", internal_port)
        pulumi.set(__self__, "public_domain", public_domain)
        pulumi.set(__self__, "public_port", public_port)

    @property
    @pulumi.getter(name="enableReadOnly")
    def enable_read_only(self) -> str:
        """
        Whether global read-only is enabled.
        """
        return pulumi.get(self, "enable_read_only")

    @property
    @pulumi.getter(name="enableReadWriteSplitting")
    def enable_read_write_splitting(self) -> str:
        """
        Whether read-write separation is enabled.
        """
        return pulumi.get(self, "enable_read_write_splitting")

    @property
    @pulumi.getter(name="internalDomain")
    def internal_domain(self) -> str:
        """
        The internal domain of the RDS instance.
        """
        return pulumi.get(self, "internal_domain")

    @property
    @pulumi.getter(name="internalPort")
    def internal_port(self) -> str:
        """
        The interval port of the RDS instance.
        """
        return pulumi.get(self, "internal_port")

    @property
    @pulumi.getter(name="publicDomain")
    def public_domain(self) -> str:
        """
        The public domain of the RDS instance.
        """
        return pulumi.get(self, "public_domain")

    @property
    @pulumi.getter(name="publicPort")
    def public_port(self) -> str:
        """
        The public port of the RDS instance.
        """
        return pulumi.get(self, "public_port")


@pulumi.output_type
class InstancesRdsInstanceInstanceSpecResult(dict):
    def __init__(__self__, *,
                 cpu_num: int,
                 mem_in_gb: int,
                 spec_name: str):
        """
        :param int cpu_num: The cpu core count of spec type.
        :param int mem_in_gb: The memory size(GB) of spec type.
        :param str spec_name: The name of spec type.
        """
        pulumi.set(__self__, "cpu_num", cpu_num)
        pulumi.set(__self__, "mem_in_gb", mem_in_gb)
        pulumi.set(__self__, "spec_name", spec_name)

    @property
    @pulumi.getter(name="cpuNum")
    def cpu_num(self) -> int:
        """
        The cpu core count of spec type.
        """
        return pulumi.get(self, "cpu_num")

    @property
    @pulumi.getter(name="memInGb")
    def mem_in_gb(self) -> int:
        """
        The memory size(GB) of spec type.
        """
        return pulumi.get(self, "mem_in_gb")

    @property
    @pulumi.getter(name="specName")
    def spec_name(self) -> str:
        """
        The name of spec type.
        """
        return pulumi.get(self, "spec_name")


@pulumi.output_type
class IpListsRdsIpListResult(dict):
    def __init__(__self__, *,
                 group_name: str,
                 id: str,
                 ip_lists: Sequence[str]):
        """
        :param str group_name: The name of the RDS ip list.
        :param str id: The ID of the RDS ip list.
        :param Sequence[str] ip_lists: The list of IP address.
        """
        pulumi.set(__self__, "group_name", group_name)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "ip_lists", ip_lists)

    @property
    @pulumi.getter(name="groupName")
    def group_name(self) -> str:
        """
        The name of the RDS ip list.
        """
        return pulumi.get(self, "group_name")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of the RDS ip list.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="ipLists")
    def ip_lists(self) -> Sequence[str]:
        """
        The list of IP address.
        """
        return pulumi.get(self, "ip_lists")


@pulumi.output_type
class ParameterTemplateTemplateParam(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "runningValue":
            suggest = "running_value"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ParameterTemplateTemplateParam. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ParameterTemplateTemplateParam.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ParameterTemplateTemplateParam.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 name: Optional[str] = None,
                 running_value: Optional[str] = None):
        """
        :param str name: Parameter name.
        :param str running_value: Parameter running value.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if running_value is not None:
            pulumi.set(__self__, "running_value", running_value)

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Parameter name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="runningValue")
    def running_value(self) -> Optional[str]:
        """
        Parameter running value.
        """
        return pulumi.get(self, "running_value")


@pulumi.output_type
class ParameterTemplatesRdsParameterTemplateResult(dict):
    def __init__(__self__, *,
                 create_time: str,
                 id: str,
                 need_restart: bool,
                 parameter_num: int,
                 template_desc: str,
                 template_id: str,
                 template_name: str,
                 template_params: Sequence['outputs.ParameterTemplatesRdsParameterTemplateTemplateParamResult'],
                 template_type: str,
                 template_type_version: str,
                 update_time: str):
        """
        :param str create_time: Creation time.
        :param str id: The ID of the RDS parameter template.
        :param bool need_restart: Whether the template contains parameters that need to be restarted.
        :param int parameter_num: The number of parameters the template contains.
        :param str template_desc: The description of the RDS parameter template.
        :param str template_id: The ID of the RDS parameter template.
        :param str template_name: The name of the RDS parameter template.
        :param Sequence['ParameterTemplatesRdsParameterTemplateTemplateParamArgs'] template_params: Parameters contained in the template.
        :param str template_type: Parameter template database type, range of values:
               MySQL - MySQL database.
        :param str template_type_version: Parameter template database version, value range:
               MySQL_Community_5_7 - MySQL 5.7
               MySQL_8_0 - MySQL 8.0.
        :param str update_time: Update time.
        """
        pulumi.set(__self__, "create_time", create_time)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "need_restart", need_restart)
        pulumi.set(__self__, "parameter_num", parameter_num)
        pulumi.set(__self__, "template_desc", template_desc)
        pulumi.set(__self__, "template_id", template_id)
        pulumi.set(__self__, "template_name", template_name)
        pulumi.set(__self__, "template_params", template_params)
        pulumi.set(__self__, "template_type", template_type)
        pulumi.set(__self__, "template_type_version", template_type_version)
        pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        Creation time.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of the RDS parameter template.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="needRestart")
    def need_restart(self) -> bool:
        """
        Whether the template contains parameters that need to be restarted.
        """
        return pulumi.get(self, "need_restart")

    @property
    @pulumi.getter(name="parameterNum")
    def parameter_num(self) -> int:
        """
        The number of parameters the template contains.
        """
        return pulumi.get(self, "parameter_num")

    @property
    @pulumi.getter(name="templateDesc")
    def template_desc(self) -> str:
        """
        The description of the RDS parameter template.
        """
        return pulumi.get(self, "template_desc")

    @property
    @pulumi.getter(name="templateId")
    def template_id(self) -> str:
        """
        The ID of the RDS parameter template.
        """
        return pulumi.get(self, "template_id")

    @property
    @pulumi.getter(name="templateName")
    def template_name(self) -> str:
        """
        The name of the RDS parameter template.
        """
        return pulumi.get(self, "template_name")

    @property
    @pulumi.getter(name="templateParams")
    def template_params(self) -> Sequence['outputs.ParameterTemplatesRdsParameterTemplateTemplateParamResult']:
        """
        Parameters contained in the template.
        """
        return pulumi.get(self, "template_params")

    @property
    @pulumi.getter(name="templateType")
    def template_type(self) -> str:
        """
        Parameter template database type, range of values:
        MySQL - MySQL database.
        """
        return pulumi.get(self, "template_type")

    @property
    @pulumi.getter(name="templateTypeVersion")
    def template_type_version(self) -> str:
        """
        Parameter template database version, value range:
        MySQL_Community_5_7 - MySQL 5.7
        MySQL_8_0 - MySQL 8.0.
        """
        return pulumi.get(self, "template_type_version")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        Update time.
        """
        return pulumi.get(self, "update_time")


@pulumi.output_type
class ParameterTemplatesRdsParameterTemplateTemplateParamResult(dict):
    def __init__(__self__, *,
                 default_value: str,
                 description: str,
                 name: str,
                 restart: bool,
                 running_value: str,
                 value_range: str):
        """
        :param str default_value: Parameter default value.
        :param str description: Parameter description.
        :param str name: Parameter name.
        :param bool restart: Whether the modified parameters need to be restarted to take effect.
        :param str running_value: Parameter running value.
        :param str value_range: Parameter value range.
        """
        pulumi.set(__self__, "default_value", default_value)
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "restart", restart)
        pulumi.set(__self__, "running_value", running_value)
        pulumi.set(__self__, "value_range", value_range)

    @property
    @pulumi.getter(name="defaultValue")
    def default_value(self) -> str:
        """
        Parameter default value.
        """
        return pulumi.get(self, "default_value")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Parameter description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Parameter name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def restart(self) -> bool:
        """
        Whether the modified parameters need to be restarted to take effect.
        """
        return pulumi.get(self, "restart")

    @property
    @pulumi.getter(name="runningValue")
    def running_value(self) -> str:
        """
        Parameter running value.
        """
        return pulumi.get(self, "running_value")

    @property
    @pulumi.getter(name="valueRange")
    def value_range(self) -> str:
        """
        Parameter value range.
        """
        return pulumi.get(self, "value_range")


