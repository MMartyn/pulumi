// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

// Export members:
export { ArgFunctionArgs, ArgFunctionResult, ArgFunctionOutputArgs } from "./argFunction";
export const argFunction: typeof import("./argFunction").argFunction = null as any;
export const argFunctionOutput: typeof import("./argFunction").argFunctionOutput = null as any;
utilities.lazyLoad(exports, ["argFunction","argFunctionOutput"], () => require("./argFunction"));

export { OtherResourceArgs } from "./otherResource";
export type OtherResource = import("./otherResource").OtherResource;
export const OtherResource: typeof import("./otherResource").OtherResource = null as any;
utilities.lazyLoad(exports, ["OtherResource"], () => require("./otherResource"));

export { ProviderArgs } from "./provider";
export type Provider = import("./provider").Provider;
export const Provider: typeof import("./provider").Provider = null as any;
utilities.lazyLoad(exports, ["Provider"], () => require("./provider"));

export { ResourceArgs } from "./resource";
export type Resource = import("./resource").Resource;
export const Resource: typeof import("./resource").Resource = null as any;
utilities.lazyLoad(exports, ["Resource"], () => require("./resource"));

export { TypeUsesArgs } from "./typeUses";
export type TypeUses = import("./typeUses").TypeUses;
export const TypeUses: typeof import("./typeUses").TypeUses = null as any;
utilities.lazyLoad(exports, ["TypeUses"], () => require("./typeUses"));


// Export enums:
export * from "./types/enums";

// Export sub-modules:
import * as types from "./types";

export {
    types,
};

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "example::OtherResource":
                return new OtherResource(name, <any>undefined, { urn })
            case "example::Resource":
                return new Resource(name, <any>undefined, { urn })
            case "example::TypeUses":
                return new TypeUses(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("example", "", _module)
pulumi.runtime.registerResourcePackage("example", {
    version: utilities.getVersion(),
    constructProvider: (name: string, type: string, urn: string): pulumi.ProviderResource => {
        if (type !== "pulumi:providers:example") {
            throw new Error(`unknown provider type ${type}`);
        }
        return new Provider(name, <any>undefined, { urn });
    },
});
