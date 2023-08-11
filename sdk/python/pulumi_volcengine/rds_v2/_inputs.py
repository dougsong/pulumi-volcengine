# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'RdsInstanceV2ChargeInfoArgs',
    'RdsInstanceV2ConnectionInfoArgs',
    'RdsInstanceV2ConnectionInfoAddressArgs',
    'RdsInstanceV2ConnectionInfoReadOnlyNodeWeightArgs',
    'RdsInstanceV2NodeInfoArgs',
]

@pulumi.input_type
class RdsInstanceV2ChargeInfoArgs:
    def __init__(__self__, *,
                 charge_type: pulumi.Input[str],
                 auto_renew: Optional[pulumi.Input[bool]] = None,
                 period: Optional[pulumi.Input[int]] = None,
                 period_unit: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] charge_type: Payment type. Value:
               PostPaid - Pay-As-You-Go
               PrePaid - Yearly and monthly (default).
        :param pulumi.Input[bool] auto_renew: Whether to automatically renew in prepaid scenarios.
        :param pulumi.Input[int] period: Purchase duration in prepaid scenarios. Default: 1.
        :param pulumi.Input[str] period_unit: The purchase cycle in the prepaid scenario.
               Month - monthly subscription (default)
               Year - Package year.
        """
        pulumi.set(__self__, "charge_type", charge_type)
        if auto_renew is not None:
            pulumi.set(__self__, "auto_renew", auto_renew)
        if period is not None:
            pulumi.set(__self__, "period", period)
        if period_unit is not None:
            pulumi.set(__self__, "period_unit", period_unit)

    @property
    @pulumi.getter(name="chargeType")
    def charge_type(self) -> pulumi.Input[str]:
        """
        Payment type. Value:
        PostPaid - Pay-As-You-Go
        PrePaid - Yearly and monthly (default).
        """
        return pulumi.get(self, "charge_type")

    @charge_type.setter
    def charge_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "charge_type", value)

    @property
    @pulumi.getter(name="autoRenew")
    def auto_renew(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to automatically renew in prepaid scenarios.
        """
        return pulumi.get(self, "auto_renew")

    @auto_renew.setter
    def auto_renew(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "auto_renew", value)

    @property
    @pulumi.getter
    def period(self) -> Optional[pulumi.Input[int]]:
        """
        Purchase duration in prepaid scenarios. Default: 1.
        """
        return pulumi.get(self, "period")

    @period.setter
    def period(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "period", value)

    @property
    @pulumi.getter(name="periodUnit")
    def period_unit(self) -> Optional[pulumi.Input[str]]:
        """
        The purchase cycle in the prepaid scenario.
        Month - monthly subscription (default)
        Year - Package year.
        """
        return pulumi.get(self, "period_unit")

    @period_unit.setter
    def period_unit(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "period_unit", value)


@pulumi.input_type
class RdsInstanceV2ConnectionInfoArgs:
    def __init__(__self__, *,
                 addresses: Optional[pulumi.Input[Sequence[pulumi.Input['RdsInstanceV2ConnectionInfoAddressArgs']]]] = None,
                 auto_add_new_nodes: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enable_read_only: Optional[pulumi.Input[str]] = None,
                 enable_read_write_splitting: Optional[pulumi.Input[str]] = None,
                 endpoint_id: Optional[pulumi.Input[str]] = None,
                 endpoint_name: Optional[pulumi.Input[str]] = None,
                 endpoint_type: Optional[pulumi.Input[str]] = None,
                 read_only_node_weights: Optional[pulumi.Input[Sequence[pulumi.Input['RdsInstanceV2ConnectionInfoReadOnlyNodeWeightArgs']]]] = None,
                 read_write_mode: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[Sequence[pulumi.Input['RdsInstanceV2ConnectionInfoAddressArgs']]] addresses: Address list.
        :param pulumi.Input[str] auto_add_new_nodes: When the terminal type is read-write terminal or read-only terminal, it supports setting whether new nodes are automatically added.
        :param pulumi.Input[str] description: Address description.
        :param pulumi.Input[str] enable_read_only: Whether global read-only is enabled, value: Enable: Enable. Disable: Disabled.
        :param pulumi.Input[str] enable_read_write_splitting: Whether read-write separation is enabled, value: Enable: Enable. Disable: Disabled.
        :param pulumi.Input[str] endpoint_id: Instance connection terminal ID.
        :param pulumi.Input[str] endpoint_name: The instance connection terminal name.
        :param pulumi.Input[str] endpoint_type: Terminal type:
               Cluster: The default terminal. (created by default)
               Primary: Primary node terminal.
               Custom: Custom terminal.
               Direct: Direct connection to the terminal. (Only the operation and maintenance side)
               AllNode: All node terminals. (Only the operation and maintenance side).
        :param pulumi.Input[Sequence[pulumi.Input['RdsInstanceV2ConnectionInfoReadOnlyNodeWeightArgs']]] read_only_node_weights: The list of nodes configured by the connection terminal and the corresponding read-only weights.
        :param pulumi.Input[str] read_write_mode: Read and write mode:
               ReadWrite: read and write
               ReadOnly: read only (default).
        """
        if addresses is not None:
            pulumi.set(__self__, "addresses", addresses)
        if auto_add_new_nodes is not None:
            pulumi.set(__self__, "auto_add_new_nodes", auto_add_new_nodes)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if enable_read_only is not None:
            pulumi.set(__self__, "enable_read_only", enable_read_only)
        if enable_read_write_splitting is not None:
            pulumi.set(__self__, "enable_read_write_splitting", enable_read_write_splitting)
        if endpoint_id is not None:
            pulumi.set(__self__, "endpoint_id", endpoint_id)
        if endpoint_name is not None:
            pulumi.set(__self__, "endpoint_name", endpoint_name)
        if endpoint_type is not None:
            pulumi.set(__self__, "endpoint_type", endpoint_type)
        if read_only_node_weights is not None:
            pulumi.set(__self__, "read_only_node_weights", read_only_node_weights)
        if read_write_mode is not None:
            pulumi.set(__self__, "read_write_mode", read_write_mode)

    @property
    @pulumi.getter
    def addresses(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RdsInstanceV2ConnectionInfoAddressArgs']]]]:
        """
        Address list.
        """
        return pulumi.get(self, "addresses")

    @addresses.setter
    def addresses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RdsInstanceV2ConnectionInfoAddressArgs']]]]):
        pulumi.set(self, "addresses", value)

    @property
    @pulumi.getter(name="autoAddNewNodes")
    def auto_add_new_nodes(self) -> Optional[pulumi.Input[str]]:
        """
        When the terminal type is read-write terminal or read-only terminal, it supports setting whether new nodes are automatically added.
        """
        return pulumi.get(self, "auto_add_new_nodes")

    @auto_add_new_nodes.setter
    def auto_add_new_nodes(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auto_add_new_nodes", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Address description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="enableReadOnly")
    def enable_read_only(self) -> Optional[pulumi.Input[str]]:
        """
        Whether global read-only is enabled, value: Enable: Enable. Disable: Disabled.
        """
        return pulumi.get(self, "enable_read_only")

    @enable_read_only.setter
    def enable_read_only(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "enable_read_only", value)

    @property
    @pulumi.getter(name="enableReadWriteSplitting")
    def enable_read_write_splitting(self) -> Optional[pulumi.Input[str]]:
        """
        Whether read-write separation is enabled, value: Enable: Enable. Disable: Disabled.
        """
        return pulumi.get(self, "enable_read_write_splitting")

    @enable_read_write_splitting.setter
    def enable_read_write_splitting(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "enable_read_write_splitting", value)

    @property
    @pulumi.getter(name="endpointId")
    def endpoint_id(self) -> Optional[pulumi.Input[str]]:
        """
        Instance connection terminal ID.
        """
        return pulumi.get(self, "endpoint_id")

    @endpoint_id.setter
    def endpoint_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "endpoint_id", value)

    @property
    @pulumi.getter(name="endpointName")
    def endpoint_name(self) -> Optional[pulumi.Input[str]]:
        """
        The instance connection terminal name.
        """
        return pulumi.get(self, "endpoint_name")

    @endpoint_name.setter
    def endpoint_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "endpoint_name", value)

    @property
    @pulumi.getter(name="endpointType")
    def endpoint_type(self) -> Optional[pulumi.Input[str]]:
        """
        Terminal type:
        Cluster: The default terminal. (created by default)
        Primary: Primary node terminal.
        Custom: Custom terminal.
        Direct: Direct connection to the terminal. (Only the operation and maintenance side)
        AllNode: All node terminals. (Only the operation and maintenance side).
        """
        return pulumi.get(self, "endpoint_type")

    @endpoint_type.setter
    def endpoint_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "endpoint_type", value)

    @property
    @pulumi.getter(name="readOnlyNodeWeights")
    def read_only_node_weights(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RdsInstanceV2ConnectionInfoReadOnlyNodeWeightArgs']]]]:
        """
        The list of nodes configured by the connection terminal and the corresponding read-only weights.
        """
        return pulumi.get(self, "read_only_node_weights")

    @read_only_node_weights.setter
    def read_only_node_weights(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RdsInstanceV2ConnectionInfoReadOnlyNodeWeightArgs']]]]):
        pulumi.set(self, "read_only_node_weights", value)

    @property
    @pulumi.getter(name="readWriteMode")
    def read_write_mode(self) -> Optional[pulumi.Input[str]]:
        """
        Read and write mode:
        ReadWrite: read and write
        ReadOnly: read only (default).
        """
        return pulumi.get(self, "read_write_mode")

    @read_write_mode.setter
    def read_write_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "read_write_mode", value)


@pulumi.input_type
class RdsInstanceV2ConnectionInfoAddressArgs:
    def __init__(__self__, *,
                 domain: Optional[pulumi.Input[str]] = None,
                 eip_id: Optional[pulumi.Input[str]] = None,
                 ip_address: Optional[pulumi.Input[str]] = None,
                 network_type: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] domain: Connect domain name.
        :param pulumi.Input[str] eip_id: The ID of the EIP, only valid for Public addresses.
        :param pulumi.Input[str] ip_address: The IP Address.
        :param pulumi.Input[str] network_type: Network address type, temporarily Private, Public, PublicService.
        :param pulumi.Input[str] port: The Port.
        :param pulumi.Input[str] subnet_id: Subnet ID.
        """
        if domain is not None:
            pulumi.set(__self__, "domain", domain)
        if eip_id is not None:
            pulumi.set(__self__, "eip_id", eip_id)
        if ip_address is not None:
            pulumi.set(__self__, "ip_address", ip_address)
        if network_type is not None:
            pulumi.set(__self__, "network_type", network_type)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if subnet_id is not None:
            pulumi.set(__self__, "subnet_id", subnet_id)

    @property
    @pulumi.getter
    def domain(self) -> Optional[pulumi.Input[str]]:
        """
        Connect domain name.
        """
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain", value)

    @property
    @pulumi.getter(name="eipId")
    def eip_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the EIP, only valid for Public addresses.
        """
        return pulumi.get(self, "eip_id")

    @eip_id.setter
    def eip_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "eip_id", value)

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> Optional[pulumi.Input[str]]:
        """
        The IP Address.
        """
        return pulumi.get(self, "ip_address")

    @ip_address.setter
    def ip_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_address", value)

    @property
    @pulumi.getter(name="networkType")
    def network_type(self) -> Optional[pulumi.Input[str]]:
        """
        Network address type, temporarily Private, Public, PublicService.
        """
        return pulumi.get(self, "network_type")

    @network_type.setter
    def network_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network_type", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[str]]:
        """
        The Port.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> Optional[pulumi.Input[str]]:
        """
        Subnet ID.
        """
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subnet_id", value)


@pulumi.input_type
class RdsInstanceV2ConnectionInfoReadOnlyNodeWeightArgs:
    def __init__(__self__, *,
                 node_id: Optional[pulumi.Input[str]] = None,
                 node_type: Optional[pulumi.Input[str]] = None,
                 weight: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[str] node_id: The ID of the node.
        :param pulumi.Input[str] node_type: Node type, the value is "Primary", "Secondary", "ReadOnly".
        :param pulumi.Input[int] weight: The weight of the node.
        """
        if node_id is not None:
            pulumi.set(__self__, "node_id", node_id)
        if node_type is not None:
            pulumi.set(__self__, "node_type", node_type)
        if weight is not None:
            pulumi.set(__self__, "weight", weight)

    @property
    @pulumi.getter(name="nodeId")
    def node_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the node.
        """
        return pulumi.get(self, "node_id")

    @node_id.setter
    def node_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "node_id", value)

    @property
    @pulumi.getter(name="nodeType")
    def node_type(self) -> Optional[pulumi.Input[str]]:
        """
        Node type, the value is "Primary", "Secondary", "ReadOnly".
        """
        return pulumi.get(self, "node_type")

    @node_type.setter
    def node_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "node_type", value)

    @property
    @pulumi.getter
    def weight(self) -> Optional[pulumi.Input[int]]:
        """
        The weight of the node.
        """
        return pulumi.get(self, "weight")

    @weight.setter
    def weight(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "weight", value)


@pulumi.input_type
class RdsInstanceV2NodeInfoArgs:
    def __init__(__self__, *,
                 node_spec: pulumi.Input[str],
                 node_type: pulumi.Input[str],
                 zone_id: pulumi.Input[str],
                 node_id: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] node_spec: Masternode specs. Pass
               DescribeDBInstanceSpecs Query the instance specifications that can be sold.
        :param pulumi.Input[str] node_type: Node type, the value is "Primary", "Secondary", "ReadOnly".
        :param pulumi.Input[str] zone_id: Zone ID.
        :param pulumi.Input[str] node_id: The ID of the node.
        """
        pulumi.set(__self__, "node_spec", node_spec)
        pulumi.set(__self__, "node_type", node_type)
        pulumi.set(__self__, "zone_id", zone_id)
        if node_id is not None:
            pulumi.set(__self__, "node_id", node_id)

    @property
    @pulumi.getter(name="nodeSpec")
    def node_spec(self) -> pulumi.Input[str]:
        """
        Masternode specs. Pass
        DescribeDBInstanceSpecs Query the instance specifications that can be sold.
        """
        return pulumi.get(self, "node_spec")

    @node_spec.setter
    def node_spec(self, value: pulumi.Input[str]):
        pulumi.set(self, "node_spec", value)

    @property
    @pulumi.getter(name="nodeType")
    def node_type(self) -> pulumi.Input[str]:
        """
        Node type, the value is "Primary", "Secondary", "ReadOnly".
        """
        return pulumi.get(self, "node_type")

    @node_type.setter
    def node_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "node_type", value)

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> pulumi.Input[str]:
        """
        Zone ID.
        """
        return pulumi.get(self, "zone_id")

    @zone_id.setter
    def zone_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "zone_id", value)

    @property
    @pulumi.getter(name="nodeId")
    def node_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the node.
        """
        return pulumi.get(self, "node_id")

    @node_id.setter
    def node_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "node_id", value)


