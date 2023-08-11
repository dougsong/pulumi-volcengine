// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rds_mysql

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage rds mysql allowlist
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/rds_mysql"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := rds_mysql.NewAllowlist(ctx, "foo", &rds_mysql.AllowlistArgs{
//				AllowLists: pulumi.StringArray{
//					pulumi.String("127.0.0.1"),
//				},
//				AllowListDesc: pulumi.String("terraform test zzm"),
//				AllowListName: pulumi.String("tf-test-opt"),
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
// RDS AllowList can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import volcengine:rds_mysql/allowlist:Allowlist default acl-d1fd76693bd54e658912e7337d5b****
//
// ```
type Allowlist struct {
	pulumi.CustomResourceState

	// The description of the allow list.
	AllowListDesc pulumi.StringPtrOutput `pulumi:"allowListDesc"`
	// The id of the allow list.
	AllowListId pulumi.StringOutput `pulumi:"allowListId"`
	// The name of the allow list.
	AllowListName pulumi.StringOutput `pulumi:"allowListName"`
	// The type of IP address in the whitelist. Currently only IPv4 addresses are supported.
	AllowListType pulumi.StringOutput `pulumi:"allowListType"`
	// Enter an IP address or a range of IP addresses in CIDR format.
	AllowLists pulumi.StringArrayOutput `pulumi:"allowLists"`
}

// NewAllowlist registers a new resource with the given unique name, arguments, and options.
func NewAllowlist(ctx *pulumi.Context,
	name string, args *AllowlistArgs, opts ...pulumi.ResourceOption) (*Allowlist, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AllowListName == nil {
		return nil, errors.New("invalid value for required argument 'AllowListName'")
	}
	if args.AllowLists == nil {
		return nil, errors.New("invalid value for required argument 'AllowLists'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Allowlist
	err := ctx.RegisterResource("volcengine:rds_mysql/allowlist:Allowlist", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAllowlist gets an existing Allowlist resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAllowlist(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AllowlistState, opts ...pulumi.ResourceOption) (*Allowlist, error) {
	var resource Allowlist
	err := ctx.ReadResource("volcengine:rds_mysql/allowlist:Allowlist", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Allowlist resources.
type allowlistState struct {
	// The description of the allow list.
	AllowListDesc *string `pulumi:"allowListDesc"`
	// The id of the allow list.
	AllowListId *string `pulumi:"allowListId"`
	// The name of the allow list.
	AllowListName *string `pulumi:"allowListName"`
	// The type of IP address in the whitelist. Currently only IPv4 addresses are supported.
	AllowListType *string `pulumi:"allowListType"`
	// Enter an IP address or a range of IP addresses in CIDR format.
	AllowLists []string `pulumi:"allowLists"`
}

type AllowlistState struct {
	// The description of the allow list.
	AllowListDesc pulumi.StringPtrInput
	// The id of the allow list.
	AllowListId pulumi.StringPtrInput
	// The name of the allow list.
	AllowListName pulumi.StringPtrInput
	// The type of IP address in the whitelist. Currently only IPv4 addresses are supported.
	AllowListType pulumi.StringPtrInput
	// Enter an IP address or a range of IP addresses in CIDR format.
	AllowLists pulumi.StringArrayInput
}

func (AllowlistState) ElementType() reflect.Type {
	return reflect.TypeOf((*allowlistState)(nil)).Elem()
}

type allowlistArgs struct {
	// The description of the allow list.
	AllowListDesc *string `pulumi:"allowListDesc"`
	// The name of the allow list.
	AllowListName string `pulumi:"allowListName"`
	// The type of IP address in the whitelist. Currently only IPv4 addresses are supported.
	AllowListType *string `pulumi:"allowListType"`
	// Enter an IP address or a range of IP addresses in CIDR format.
	AllowLists []string `pulumi:"allowLists"`
}

// The set of arguments for constructing a Allowlist resource.
type AllowlistArgs struct {
	// The description of the allow list.
	AllowListDesc pulumi.StringPtrInput
	// The name of the allow list.
	AllowListName pulumi.StringInput
	// The type of IP address in the whitelist. Currently only IPv4 addresses are supported.
	AllowListType pulumi.StringPtrInput
	// Enter an IP address or a range of IP addresses in CIDR format.
	AllowLists pulumi.StringArrayInput
}

func (AllowlistArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*allowlistArgs)(nil)).Elem()
}

type AllowlistInput interface {
	pulumi.Input

	ToAllowlistOutput() AllowlistOutput
	ToAllowlistOutputWithContext(ctx context.Context) AllowlistOutput
}

func (*Allowlist) ElementType() reflect.Type {
	return reflect.TypeOf((**Allowlist)(nil)).Elem()
}

func (i *Allowlist) ToAllowlistOutput() AllowlistOutput {
	return i.ToAllowlistOutputWithContext(context.Background())
}

func (i *Allowlist) ToAllowlistOutputWithContext(ctx context.Context) AllowlistOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AllowlistOutput)
}

// AllowlistArrayInput is an input type that accepts AllowlistArray and AllowlistArrayOutput values.
// You can construct a concrete instance of `AllowlistArrayInput` via:
//
//	AllowlistArray{ AllowlistArgs{...} }
type AllowlistArrayInput interface {
	pulumi.Input

	ToAllowlistArrayOutput() AllowlistArrayOutput
	ToAllowlistArrayOutputWithContext(context.Context) AllowlistArrayOutput
}

type AllowlistArray []AllowlistInput

func (AllowlistArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Allowlist)(nil)).Elem()
}

func (i AllowlistArray) ToAllowlistArrayOutput() AllowlistArrayOutput {
	return i.ToAllowlistArrayOutputWithContext(context.Background())
}

func (i AllowlistArray) ToAllowlistArrayOutputWithContext(ctx context.Context) AllowlistArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AllowlistArrayOutput)
}

// AllowlistMapInput is an input type that accepts AllowlistMap and AllowlistMapOutput values.
// You can construct a concrete instance of `AllowlistMapInput` via:
//
//	AllowlistMap{ "key": AllowlistArgs{...} }
type AllowlistMapInput interface {
	pulumi.Input

	ToAllowlistMapOutput() AllowlistMapOutput
	ToAllowlistMapOutputWithContext(context.Context) AllowlistMapOutput
}

type AllowlistMap map[string]AllowlistInput

func (AllowlistMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Allowlist)(nil)).Elem()
}

func (i AllowlistMap) ToAllowlistMapOutput() AllowlistMapOutput {
	return i.ToAllowlistMapOutputWithContext(context.Background())
}

func (i AllowlistMap) ToAllowlistMapOutputWithContext(ctx context.Context) AllowlistMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AllowlistMapOutput)
}

type AllowlistOutput struct{ *pulumi.OutputState }

func (AllowlistOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Allowlist)(nil)).Elem()
}

func (o AllowlistOutput) ToAllowlistOutput() AllowlistOutput {
	return o
}

func (o AllowlistOutput) ToAllowlistOutputWithContext(ctx context.Context) AllowlistOutput {
	return o
}

// The description of the allow list.
func (o AllowlistOutput) AllowListDesc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Allowlist) pulumi.StringPtrOutput { return v.AllowListDesc }).(pulumi.StringPtrOutput)
}

// The id of the allow list.
func (o AllowlistOutput) AllowListId() pulumi.StringOutput {
	return o.ApplyT(func(v *Allowlist) pulumi.StringOutput { return v.AllowListId }).(pulumi.StringOutput)
}

// The name of the allow list.
func (o AllowlistOutput) AllowListName() pulumi.StringOutput {
	return o.ApplyT(func(v *Allowlist) pulumi.StringOutput { return v.AllowListName }).(pulumi.StringOutput)
}

// The type of IP address in the whitelist. Currently only IPv4 addresses are supported.
func (o AllowlistOutput) AllowListType() pulumi.StringOutput {
	return o.ApplyT(func(v *Allowlist) pulumi.StringOutput { return v.AllowListType }).(pulumi.StringOutput)
}

// Enter an IP address or a range of IP addresses in CIDR format.
func (o AllowlistOutput) AllowLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Allowlist) pulumi.StringArrayOutput { return v.AllowLists }).(pulumi.StringArrayOutput)
}

type AllowlistArrayOutput struct{ *pulumi.OutputState }

func (AllowlistArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Allowlist)(nil)).Elem()
}

func (o AllowlistArrayOutput) ToAllowlistArrayOutput() AllowlistArrayOutput {
	return o
}

func (o AllowlistArrayOutput) ToAllowlistArrayOutputWithContext(ctx context.Context) AllowlistArrayOutput {
	return o
}

func (o AllowlistArrayOutput) Index(i pulumi.IntInput) AllowlistOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Allowlist {
		return vs[0].([]*Allowlist)[vs[1].(int)]
	}).(AllowlistOutput)
}

type AllowlistMapOutput struct{ *pulumi.OutputState }

func (AllowlistMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Allowlist)(nil)).Elem()
}

func (o AllowlistMapOutput) ToAllowlistMapOutput() AllowlistMapOutput {
	return o
}

func (o AllowlistMapOutput) ToAllowlistMapOutputWithContext(ctx context.Context) AllowlistMapOutput {
	return o
}

func (o AllowlistMapOutput) MapIndex(k pulumi.StringInput) AllowlistOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Allowlist {
		return vs[0].(map[string]*Allowlist)[vs[1].(string)]
	}).(AllowlistOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AllowlistInput)(nil)).Elem(), &Allowlist{})
	pulumi.RegisterInputType(reflect.TypeOf((*AllowlistArrayInput)(nil)).Elem(), AllowlistArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AllowlistMapInput)(nil)).Elem(), AllowlistMap{})
	pulumi.RegisterOutputType(AllowlistOutput{})
	pulumi.RegisterOutputType(AllowlistArrayOutput{})
	pulumi.RegisterOutputType(AllowlistMapOutput{})
}
