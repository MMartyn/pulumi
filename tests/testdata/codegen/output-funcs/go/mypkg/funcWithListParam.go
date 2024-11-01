// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mypkg

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
	"output-funcs/mypkg/internal"
)

// Check codegen of functions with a List parameter.
func FuncWithListParam(ctx *pulumi.Context, args *FuncWithListParamArgs, opts ...pulumi.InvokeOption) (*FuncWithListParamResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv FuncWithListParamResult
	err := ctx.Invoke("mypkg::funcWithListParam", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type FuncWithListParamArgs struct {
	A []string `pulumi:"a"`
	B *string  `pulumi:"b"`
}

type FuncWithListParamResult struct {
	R string `pulumi:"r"`
}

func FuncWithListParamOutput(ctx *pulumi.Context, args FuncWithListParamOutputArgs, opts ...pulumi.InvokeOption) FuncWithListParamResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (FuncWithListParamResultOutput, error) {
			args := v.(FuncWithListParamArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv FuncWithListParamResult
			secret, err := ctx.InvokePackageRaw("mypkg::funcWithListParam", args, &rv, "", opts...)
			if err != nil {
				return FuncWithListParamResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(FuncWithListParamResultOutput)
			if secret {
				return pulumi.ToSecret(output).(FuncWithListParamResultOutput), nil
			}
			return output, nil
		}).(FuncWithListParamResultOutput)
}

type FuncWithListParamOutputArgs struct {
	A pulumi.StringArrayInput `pulumi:"a"`
	B pulumi.StringPtrInput   `pulumi:"b"`
}

func (FuncWithListParamOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FuncWithListParamArgs)(nil)).Elem()
}

type FuncWithListParamResultOutput struct{ *pulumi.OutputState }

func (FuncWithListParamResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FuncWithListParamResult)(nil)).Elem()
}

func (o FuncWithListParamResultOutput) ToFuncWithListParamResultOutput() FuncWithListParamResultOutput {
	return o
}

func (o FuncWithListParamResultOutput) ToFuncWithListParamResultOutputWithContext(ctx context.Context) FuncWithListParamResultOutput {
	return o
}

func (o FuncWithListParamResultOutput) ToOutput(ctx context.Context) pulumix.Output[FuncWithListParamResult] {
	return pulumix.Output[FuncWithListParamResult]{
		OutputState: o.OutputState,
	}
}

func (o FuncWithListParamResultOutput) R() pulumi.StringOutput {
	return o.ApplyT(func(v FuncWithListParamResult) string { return v.R }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(FuncWithListParamResultOutput{})
}
