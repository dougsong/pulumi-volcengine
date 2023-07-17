// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tos

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage tos bucket policy
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"encoding/json"
//
//	"github.com/pulumi/pulumi-volcengine/sdk/go/volcengine/tos"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			tmpJSON0, err := json.Marshal(map[string]interface{}{
//				"Statement": []map[string]interface{}{
//					map[string]interface{}{
//						"Sid":    "test",
//						"Effect": "Allow",
//						"Principal": []string{
//							"AccountId/subUserName",
//						},
//						"Action": []string{
//							"tos:List*",
//						},
//						"Resource": []string{
//							"trn:tos:::bucket-20230418",
//						},
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			json0 := string(tmpJSON0)
//			_, err := tos.NewBucketPolicy(ctx, "default", &tos.BucketPolicyArgs{
//				BucketName: pulumi.String("bucket-20230418"),
//				Policy:     pulumi.String(json0),
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
// Tos Bucket can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import volcengine:tos/bucketPolicy:BucketPolicy default bucketName:policy
//
// ```
type BucketPolicy struct {
	pulumi.CustomResourceState

	// The name of the bucket.
	BucketName pulumi.StringOutput `pulumi:"bucketName"`
	// The policy document. This is a JSON formatted string. For more information about building Volcengine IAM policy
	// documents with Terraform, see the [Volcengine IAM Policy Document Guide](https://www.volcengine.com/docs/6349/102127).
	Policy pulumi.StringOutput `pulumi:"policy"`
}

// NewBucketPolicy registers a new resource with the given unique name, arguments, and options.
func NewBucketPolicy(ctx *pulumi.Context,
	name string, args *BucketPolicyArgs, opts ...pulumi.ResourceOption) (*BucketPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BucketName == nil {
		return nil, errors.New("invalid value for required argument 'BucketName'")
	}
	if args.Policy == nil {
		return nil, errors.New("invalid value for required argument 'Policy'")
	}
	var resource BucketPolicy
	err := ctx.RegisterResource("volcengine:tos/bucketPolicy:BucketPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBucketPolicy gets an existing BucketPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBucketPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BucketPolicyState, opts ...pulumi.ResourceOption) (*BucketPolicy, error) {
	var resource BucketPolicy
	err := ctx.ReadResource("volcengine:tos/bucketPolicy:BucketPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BucketPolicy resources.
type bucketPolicyState struct {
	// The name of the bucket.
	BucketName *string `pulumi:"bucketName"`
	// The policy document. This is a JSON formatted string. For more information about building Volcengine IAM policy
	// documents with Terraform, see the [Volcengine IAM Policy Document Guide](https://www.volcengine.com/docs/6349/102127).
	Policy *string `pulumi:"policy"`
}

type BucketPolicyState struct {
	// The name of the bucket.
	BucketName pulumi.StringPtrInput
	// The policy document. This is a JSON formatted string. For more information about building Volcengine IAM policy
	// documents with Terraform, see the [Volcengine IAM Policy Document Guide](https://www.volcengine.com/docs/6349/102127).
	Policy pulumi.StringPtrInput
}

func (BucketPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketPolicyState)(nil)).Elem()
}

type bucketPolicyArgs struct {
	// The name of the bucket.
	BucketName string `pulumi:"bucketName"`
	// The policy document. This is a JSON formatted string. For more information about building Volcengine IAM policy
	// documents with Terraform, see the [Volcengine IAM Policy Document Guide](https://www.volcengine.com/docs/6349/102127).
	Policy string `pulumi:"policy"`
}

// The set of arguments for constructing a BucketPolicy resource.
type BucketPolicyArgs struct {
	// The name of the bucket.
	BucketName pulumi.StringInput
	// The policy document. This is a JSON formatted string. For more information about building Volcengine IAM policy
	// documents with Terraform, see the [Volcengine IAM Policy Document Guide](https://www.volcengine.com/docs/6349/102127).
	Policy pulumi.StringInput
}

func (BucketPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketPolicyArgs)(nil)).Elem()
}

type BucketPolicyInput interface {
	pulumi.Input

	ToBucketPolicyOutput() BucketPolicyOutput
	ToBucketPolicyOutputWithContext(ctx context.Context) BucketPolicyOutput
}

func (*BucketPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketPolicy)(nil)).Elem()
}

func (i *BucketPolicy) ToBucketPolicyOutput() BucketPolicyOutput {
	return i.ToBucketPolicyOutputWithContext(context.Background())
}

func (i *BucketPolicy) ToBucketPolicyOutputWithContext(ctx context.Context) BucketPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketPolicyOutput)
}

// BucketPolicyArrayInput is an input type that accepts BucketPolicyArray and BucketPolicyArrayOutput values.
// You can construct a concrete instance of `BucketPolicyArrayInput` via:
//
//	BucketPolicyArray{ BucketPolicyArgs{...} }
type BucketPolicyArrayInput interface {
	pulumi.Input

	ToBucketPolicyArrayOutput() BucketPolicyArrayOutput
	ToBucketPolicyArrayOutputWithContext(context.Context) BucketPolicyArrayOutput
}

type BucketPolicyArray []BucketPolicyInput

func (BucketPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BucketPolicy)(nil)).Elem()
}

func (i BucketPolicyArray) ToBucketPolicyArrayOutput() BucketPolicyArrayOutput {
	return i.ToBucketPolicyArrayOutputWithContext(context.Background())
}

func (i BucketPolicyArray) ToBucketPolicyArrayOutputWithContext(ctx context.Context) BucketPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketPolicyArrayOutput)
}

// BucketPolicyMapInput is an input type that accepts BucketPolicyMap and BucketPolicyMapOutput values.
// You can construct a concrete instance of `BucketPolicyMapInput` via:
//
//	BucketPolicyMap{ "key": BucketPolicyArgs{...} }
type BucketPolicyMapInput interface {
	pulumi.Input

	ToBucketPolicyMapOutput() BucketPolicyMapOutput
	ToBucketPolicyMapOutputWithContext(context.Context) BucketPolicyMapOutput
}

type BucketPolicyMap map[string]BucketPolicyInput

func (BucketPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BucketPolicy)(nil)).Elem()
}

func (i BucketPolicyMap) ToBucketPolicyMapOutput() BucketPolicyMapOutput {
	return i.ToBucketPolicyMapOutputWithContext(context.Background())
}

func (i BucketPolicyMap) ToBucketPolicyMapOutputWithContext(ctx context.Context) BucketPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketPolicyMapOutput)
}

type BucketPolicyOutput struct{ *pulumi.OutputState }

func (BucketPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketPolicy)(nil)).Elem()
}

func (o BucketPolicyOutput) ToBucketPolicyOutput() BucketPolicyOutput {
	return o
}

func (o BucketPolicyOutput) ToBucketPolicyOutputWithContext(ctx context.Context) BucketPolicyOutput {
	return o
}

// The name of the bucket.
func (o BucketPolicyOutput) BucketName() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketPolicy) pulumi.StringOutput { return v.BucketName }).(pulumi.StringOutput)
}

// The policy document. This is a JSON formatted string. For more information about building Volcengine IAM policy
// documents with Terraform, see the [Volcengine IAM Policy Document Guide](https://www.volcengine.com/docs/6349/102127).
func (o BucketPolicyOutput) Policy() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketPolicy) pulumi.StringOutput { return v.Policy }).(pulumi.StringOutput)
}

type BucketPolicyArrayOutput struct{ *pulumi.OutputState }

func (BucketPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BucketPolicy)(nil)).Elem()
}

func (o BucketPolicyArrayOutput) ToBucketPolicyArrayOutput() BucketPolicyArrayOutput {
	return o
}

func (o BucketPolicyArrayOutput) ToBucketPolicyArrayOutputWithContext(ctx context.Context) BucketPolicyArrayOutput {
	return o
}

func (o BucketPolicyArrayOutput) Index(i pulumi.IntInput) BucketPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BucketPolicy {
		return vs[0].([]*BucketPolicy)[vs[1].(int)]
	}).(BucketPolicyOutput)
}

type BucketPolicyMapOutput struct{ *pulumi.OutputState }

func (BucketPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BucketPolicy)(nil)).Elem()
}

func (o BucketPolicyMapOutput) ToBucketPolicyMapOutput() BucketPolicyMapOutput {
	return o
}

func (o BucketPolicyMapOutput) ToBucketPolicyMapOutputWithContext(ctx context.Context) BucketPolicyMapOutput {
	return o
}

func (o BucketPolicyMapOutput) MapIndex(k pulumi.StringInput) BucketPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BucketPolicy {
		return vs[0].(map[string]*BucketPolicy)[vs[1].(string)]
	}).(BucketPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BucketPolicyInput)(nil)).Elem(), &BucketPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*BucketPolicyArrayInput)(nil)).Elem(), BucketPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BucketPolicyMapInput)(nil)).Elem(), BucketPolicyMap{})
	pulumi.RegisterOutputType(BucketPolicyOutput{})
	pulumi.RegisterOutputType(BucketPolicyArrayOutput{})
	pulumi.RegisterOutputType(BucketPolicyMapOutput{})
}
