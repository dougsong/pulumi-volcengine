// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage security group rule
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/vpc"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vpc.NewSecurityGroupRule(ctx, "g1test3", &vpc.SecurityGroupRuleArgs{
//				CidrIp:          pulumi.String("10.0.0.0/8"),
//				Description:     pulumi.String("tft1234"),
//				Direction:       pulumi.String("egress"),
//				PortEnd:         pulumi.Int(9003),
//				PortStart:       pulumi.Int(8000),
//				Protocol:        pulumi.String("tcp"),
//				SecurityGroupId: pulumi.String("sg-2d6722jpp55og58ozfd1sqtdb"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = vpc.NewSecurityGroupRule(ctx, "g1test2", &vpc.SecurityGroupRuleArgs{
//				CidrIp:          pulumi.String("10.0.0.0/24"),
//				Direction:       pulumi.String("egress"),
//				PortEnd:         pulumi.Int(9003),
//				PortStart:       pulumi.Int(8000),
//				Protocol:        pulumi.String("tcp"),
//				SecurityGroupId: pulumi.String("sg-2d6722jpp55og58ozfd1sqtdb"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = vpc.NewSecurityGroupRule(ctx, "g1test1", &vpc.SecurityGroupRuleArgs{
//				CidrIp:          pulumi.String("10.0.0.0/24"),
//				Direction:       pulumi.String("egress"),
//				PortEnd:         pulumi.Int(9003),
//				PortStart:       pulumi.Int(8000),
//				Priority:        pulumi.Int(2),
//				Protocol:        pulumi.String("tcp"),
//				SecurityGroupId: pulumi.String("sg-2d6722jpp55og58ozfd1sqtdb"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = vpc.NewSecurityGroupRule(ctx, "g1test0", &vpc.SecurityGroupRuleArgs{
//				CidrIp:          pulumi.String("10.0.0.0/24"),
//				Description:     pulumi.String("tft"),
//				Direction:       pulumi.String("ingress"),
//				Policy:          pulumi.String("drop"),
//				PortEnd:         pulumi.Int(80),
//				PortStart:       pulumi.Int(80),
//				Priority:        pulumi.Int(2),
//				Protocol:        pulumi.String("tcp"),
//				SecurityGroupId: pulumi.String("sg-2d6722jpp55og58ozfd1sqtdb"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = vpc.NewSecurityGroupRule(ctx, "g1test06", &vpc.SecurityGroupRuleArgs{
//				Description:     pulumi.String("tft"),
//				Direction:       pulumi.String("ingress"),
//				Policy:          pulumi.String("drop"),
//				PortEnd:         pulumi.Int(9003),
//				PortStart:       pulumi.Int(8000),
//				Priority:        pulumi.Int(2),
//				Protocol:        pulumi.String("tcp"),
//				SecurityGroupId: pulumi.String("sg-2d6722jpp55og58ozfd1sqtdb"),
//				SourceGroupId:   pulumi.String("sg-3rfe5j4xdnklc5zsk2hcw5c6q"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// SecurityGroupRule can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import volcengine:vpc/securityGroupRule:SecurityGroupRule default ID is a string concatenated with colons(SecurityGroupId:Protocol:PortStart:PortEnd:CidrIp:SourceGroupId:Direction:Policy:Priority)
//
// ```
type SecurityGroupRule struct {
	pulumi.CustomResourceState

	// Cidr ip of egress/ingress Rule.
	CidrIp pulumi.StringPtrOutput `pulumi:"cidrIp"`
	// description of a egress rule.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Direction of rule, ingress (inbound) or egress (outbound).
	Direction pulumi.StringOutput `pulumi:"direction"`
	// Access strategy.
	Policy pulumi.StringPtrOutput `pulumi:"policy"`
	// Port end of egress/ingress Rule.
	PortEnd pulumi.IntOutput `pulumi:"portEnd"`
	// Port start of egress/ingress Rule.
	PortStart pulumi.IntOutput `pulumi:"portStart"`
	// Priority of a security group rule.
	Priority pulumi.IntPtrOutput `pulumi:"priority"`
	// Protocol of the SecurityGroup, the value can be `tcp` or `udp` or `icmp` or `all` or `icmpv6`.
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// Id of SecurityGroup.
	SecurityGroupId pulumi.StringOutput `pulumi:"securityGroupId"`
	// ID of the source security group whose access permission you want to set.
	SourceGroupId pulumi.StringPtrOutput `pulumi:"sourceGroupId"`
	// Status of SecurityGroup.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewSecurityGroupRule registers a new resource with the given unique name, arguments, and options.
func NewSecurityGroupRule(ctx *pulumi.Context,
	name string, args *SecurityGroupRuleArgs, opts ...pulumi.ResourceOption) (*SecurityGroupRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Direction == nil {
		return nil, errors.New("invalid value for required argument 'Direction'")
	}
	if args.PortEnd == nil {
		return nil, errors.New("invalid value for required argument 'PortEnd'")
	}
	if args.PortStart == nil {
		return nil, errors.New("invalid value for required argument 'PortStart'")
	}
	if args.Protocol == nil {
		return nil, errors.New("invalid value for required argument 'Protocol'")
	}
	if args.SecurityGroupId == nil {
		return nil, errors.New("invalid value for required argument 'SecurityGroupId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource SecurityGroupRule
	err := ctx.RegisterResource("volcengine:vpc/securityGroupRule:SecurityGroupRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecurityGroupRule gets an existing SecurityGroupRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecurityGroupRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecurityGroupRuleState, opts ...pulumi.ResourceOption) (*SecurityGroupRule, error) {
	var resource SecurityGroupRule
	err := ctx.ReadResource("volcengine:vpc/securityGroupRule:SecurityGroupRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecurityGroupRule resources.
type securityGroupRuleState struct {
	// Cidr ip of egress/ingress Rule.
	CidrIp *string `pulumi:"cidrIp"`
	// description of a egress rule.
	Description *string `pulumi:"description"`
	// Direction of rule, ingress (inbound) or egress (outbound).
	Direction *string `pulumi:"direction"`
	// Access strategy.
	Policy *string `pulumi:"policy"`
	// Port end of egress/ingress Rule.
	PortEnd *int `pulumi:"portEnd"`
	// Port start of egress/ingress Rule.
	PortStart *int `pulumi:"portStart"`
	// Priority of a security group rule.
	Priority *int `pulumi:"priority"`
	// Protocol of the SecurityGroup, the value can be `tcp` or `udp` or `icmp` or `all` or `icmpv6`.
	Protocol *string `pulumi:"protocol"`
	// Id of SecurityGroup.
	SecurityGroupId *string `pulumi:"securityGroupId"`
	// ID of the source security group whose access permission you want to set.
	SourceGroupId *string `pulumi:"sourceGroupId"`
	// Status of SecurityGroup.
	Status *string `pulumi:"status"`
}

type SecurityGroupRuleState struct {
	// Cidr ip of egress/ingress Rule.
	CidrIp pulumi.StringPtrInput
	// description of a egress rule.
	Description pulumi.StringPtrInput
	// Direction of rule, ingress (inbound) or egress (outbound).
	Direction pulumi.StringPtrInput
	// Access strategy.
	Policy pulumi.StringPtrInput
	// Port end of egress/ingress Rule.
	PortEnd pulumi.IntPtrInput
	// Port start of egress/ingress Rule.
	PortStart pulumi.IntPtrInput
	// Priority of a security group rule.
	Priority pulumi.IntPtrInput
	// Protocol of the SecurityGroup, the value can be `tcp` or `udp` or `icmp` or `all` or `icmpv6`.
	Protocol pulumi.StringPtrInput
	// Id of SecurityGroup.
	SecurityGroupId pulumi.StringPtrInput
	// ID of the source security group whose access permission you want to set.
	SourceGroupId pulumi.StringPtrInput
	// Status of SecurityGroup.
	Status pulumi.StringPtrInput
}

func (SecurityGroupRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*securityGroupRuleState)(nil)).Elem()
}

type securityGroupRuleArgs struct {
	// Cidr ip of egress/ingress Rule.
	CidrIp *string `pulumi:"cidrIp"`
	// description of a egress rule.
	Description *string `pulumi:"description"`
	// Direction of rule, ingress (inbound) or egress (outbound).
	Direction string `pulumi:"direction"`
	// Access strategy.
	Policy *string `pulumi:"policy"`
	// Port end of egress/ingress Rule.
	PortEnd int `pulumi:"portEnd"`
	// Port start of egress/ingress Rule.
	PortStart int `pulumi:"portStart"`
	// Priority of a security group rule.
	Priority *int `pulumi:"priority"`
	// Protocol of the SecurityGroup, the value can be `tcp` or `udp` or `icmp` or `all` or `icmpv6`.
	Protocol string `pulumi:"protocol"`
	// Id of SecurityGroup.
	SecurityGroupId string `pulumi:"securityGroupId"`
	// ID of the source security group whose access permission you want to set.
	SourceGroupId *string `pulumi:"sourceGroupId"`
}

// The set of arguments for constructing a SecurityGroupRule resource.
type SecurityGroupRuleArgs struct {
	// Cidr ip of egress/ingress Rule.
	CidrIp pulumi.StringPtrInput
	// description of a egress rule.
	Description pulumi.StringPtrInput
	// Direction of rule, ingress (inbound) or egress (outbound).
	Direction pulumi.StringInput
	// Access strategy.
	Policy pulumi.StringPtrInput
	// Port end of egress/ingress Rule.
	PortEnd pulumi.IntInput
	// Port start of egress/ingress Rule.
	PortStart pulumi.IntInput
	// Priority of a security group rule.
	Priority pulumi.IntPtrInput
	// Protocol of the SecurityGroup, the value can be `tcp` or `udp` or `icmp` or `all` or `icmpv6`.
	Protocol pulumi.StringInput
	// Id of SecurityGroup.
	SecurityGroupId pulumi.StringInput
	// ID of the source security group whose access permission you want to set.
	SourceGroupId pulumi.StringPtrInput
}

func (SecurityGroupRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*securityGroupRuleArgs)(nil)).Elem()
}

type SecurityGroupRuleInput interface {
	pulumi.Input

	ToSecurityGroupRuleOutput() SecurityGroupRuleOutput
	ToSecurityGroupRuleOutputWithContext(ctx context.Context) SecurityGroupRuleOutput
}

func (*SecurityGroupRule) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityGroupRule)(nil)).Elem()
}

func (i *SecurityGroupRule) ToSecurityGroupRuleOutput() SecurityGroupRuleOutput {
	return i.ToSecurityGroupRuleOutputWithContext(context.Background())
}

func (i *SecurityGroupRule) ToSecurityGroupRuleOutputWithContext(ctx context.Context) SecurityGroupRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityGroupRuleOutput)
}

// SecurityGroupRuleArrayInput is an input type that accepts SecurityGroupRuleArray and SecurityGroupRuleArrayOutput values.
// You can construct a concrete instance of `SecurityGroupRuleArrayInput` via:
//
//	SecurityGroupRuleArray{ SecurityGroupRuleArgs{...} }
type SecurityGroupRuleArrayInput interface {
	pulumi.Input

	ToSecurityGroupRuleArrayOutput() SecurityGroupRuleArrayOutput
	ToSecurityGroupRuleArrayOutputWithContext(context.Context) SecurityGroupRuleArrayOutput
}

type SecurityGroupRuleArray []SecurityGroupRuleInput

func (SecurityGroupRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecurityGroupRule)(nil)).Elem()
}

func (i SecurityGroupRuleArray) ToSecurityGroupRuleArrayOutput() SecurityGroupRuleArrayOutput {
	return i.ToSecurityGroupRuleArrayOutputWithContext(context.Background())
}

func (i SecurityGroupRuleArray) ToSecurityGroupRuleArrayOutputWithContext(ctx context.Context) SecurityGroupRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityGroupRuleArrayOutput)
}

// SecurityGroupRuleMapInput is an input type that accepts SecurityGroupRuleMap and SecurityGroupRuleMapOutput values.
// You can construct a concrete instance of `SecurityGroupRuleMapInput` via:
//
//	SecurityGroupRuleMap{ "key": SecurityGroupRuleArgs{...} }
type SecurityGroupRuleMapInput interface {
	pulumi.Input

	ToSecurityGroupRuleMapOutput() SecurityGroupRuleMapOutput
	ToSecurityGroupRuleMapOutputWithContext(context.Context) SecurityGroupRuleMapOutput
}

type SecurityGroupRuleMap map[string]SecurityGroupRuleInput

func (SecurityGroupRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecurityGroupRule)(nil)).Elem()
}

func (i SecurityGroupRuleMap) ToSecurityGroupRuleMapOutput() SecurityGroupRuleMapOutput {
	return i.ToSecurityGroupRuleMapOutputWithContext(context.Background())
}

func (i SecurityGroupRuleMap) ToSecurityGroupRuleMapOutputWithContext(ctx context.Context) SecurityGroupRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityGroupRuleMapOutput)
}

type SecurityGroupRuleOutput struct{ *pulumi.OutputState }

func (SecurityGroupRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityGroupRule)(nil)).Elem()
}

func (o SecurityGroupRuleOutput) ToSecurityGroupRuleOutput() SecurityGroupRuleOutput {
	return o
}

func (o SecurityGroupRuleOutput) ToSecurityGroupRuleOutputWithContext(ctx context.Context) SecurityGroupRuleOutput {
	return o
}

// Cidr ip of egress/ingress Rule.
func (o SecurityGroupRuleOutput) CidrIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityGroupRule) pulumi.StringPtrOutput { return v.CidrIp }).(pulumi.StringPtrOutput)
}

// description of a egress rule.
func (o SecurityGroupRuleOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityGroupRule) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Direction of rule, ingress (inbound) or egress (outbound).
func (o SecurityGroupRuleOutput) Direction() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityGroupRule) pulumi.StringOutput { return v.Direction }).(pulumi.StringOutput)
}

// Access strategy.
func (o SecurityGroupRuleOutput) Policy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityGroupRule) pulumi.StringPtrOutput { return v.Policy }).(pulumi.StringPtrOutput)
}

// Port end of egress/ingress Rule.
func (o SecurityGroupRuleOutput) PortEnd() pulumi.IntOutput {
	return o.ApplyT(func(v *SecurityGroupRule) pulumi.IntOutput { return v.PortEnd }).(pulumi.IntOutput)
}

// Port start of egress/ingress Rule.
func (o SecurityGroupRuleOutput) PortStart() pulumi.IntOutput {
	return o.ApplyT(func(v *SecurityGroupRule) pulumi.IntOutput { return v.PortStart }).(pulumi.IntOutput)
}

// Priority of a security group rule.
func (o SecurityGroupRuleOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SecurityGroupRule) pulumi.IntPtrOutput { return v.Priority }).(pulumi.IntPtrOutput)
}

// Protocol of the SecurityGroup, the value can be `tcp` or `udp` or `icmp` or `all` or `icmpv6`.
func (o SecurityGroupRuleOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityGroupRule) pulumi.StringOutput { return v.Protocol }).(pulumi.StringOutput)
}

// Id of SecurityGroup.
func (o SecurityGroupRuleOutput) SecurityGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityGroupRule) pulumi.StringOutput { return v.SecurityGroupId }).(pulumi.StringOutput)
}

// ID of the source security group whose access permission you want to set.
func (o SecurityGroupRuleOutput) SourceGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityGroupRule) pulumi.StringPtrOutput { return v.SourceGroupId }).(pulumi.StringPtrOutput)
}

// Status of SecurityGroup.
func (o SecurityGroupRuleOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityGroupRule) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type SecurityGroupRuleArrayOutput struct{ *pulumi.OutputState }

func (SecurityGroupRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecurityGroupRule)(nil)).Elem()
}

func (o SecurityGroupRuleArrayOutput) ToSecurityGroupRuleArrayOutput() SecurityGroupRuleArrayOutput {
	return o
}

func (o SecurityGroupRuleArrayOutput) ToSecurityGroupRuleArrayOutputWithContext(ctx context.Context) SecurityGroupRuleArrayOutput {
	return o
}

func (o SecurityGroupRuleArrayOutput) Index(i pulumi.IntInput) SecurityGroupRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SecurityGroupRule {
		return vs[0].([]*SecurityGroupRule)[vs[1].(int)]
	}).(SecurityGroupRuleOutput)
}

type SecurityGroupRuleMapOutput struct{ *pulumi.OutputState }

func (SecurityGroupRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecurityGroupRule)(nil)).Elem()
}

func (o SecurityGroupRuleMapOutput) ToSecurityGroupRuleMapOutput() SecurityGroupRuleMapOutput {
	return o
}

func (o SecurityGroupRuleMapOutput) ToSecurityGroupRuleMapOutputWithContext(ctx context.Context) SecurityGroupRuleMapOutput {
	return o
}

func (o SecurityGroupRuleMapOutput) MapIndex(k pulumi.StringInput) SecurityGroupRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SecurityGroupRule {
		return vs[0].(map[string]*SecurityGroupRule)[vs[1].(string)]
	}).(SecurityGroupRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecurityGroupRuleInput)(nil)).Elem(), &SecurityGroupRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecurityGroupRuleArrayInput)(nil)).Elem(), SecurityGroupRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecurityGroupRuleMapInput)(nil)).Elem(), SecurityGroupRuleMap{})
	pulumi.RegisterOutputType(SecurityGroupRuleOutput{})
	pulumi.RegisterOutputType(SecurityGroupRuleArrayOutput{})
	pulumi.RegisterOutputType(SecurityGroupRuleMapOutput{})
}
