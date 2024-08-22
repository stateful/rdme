/* eslint-disable */
// @generated by protobuf-ts 2.9.4 with parameter output_javascript,optimize_code_size,long_type_string,add_pb_suffix,ts_nocheck,eslint_disable
// @generated from protobuf file "runme/reporter/v1/reporter.proto" (package "runme.reporter.v1", syntax proto3)
// tslint:disable
// @ts-nocheck
/* eslint-disable */
// @generated by protobuf-ts 2.9.4 with parameter output_javascript,optimize_code_size,long_type_string,add_pb_suffix,ts_nocheck,eslint_disable
// @generated from protobuf file "runme/reporter/v1/reporter.proto" (package "runme.reporter.v1", syntax proto3)
// tslint:disable
// @ts-nocheck
import { ServiceType } from "@protobuf-ts/runtime-rpc";
import { MessageType } from "@protobuf-ts/runtime";
import { Notebook } from "../../parser/v1/parser_pb";
// @generated message type with reflection information, may provide speed optimized methods
class TransformRequest$Type extends MessageType {
    constructor() {
        super("runme.reporter.v1.TransformRequest", [
            { no: 1, name: "notebook", kind: "message", T: () => Notebook },
            { no: 2, name: "auto_save", kind: "scalar", opt: true, T: 8 /*ScalarType.BOOL*/ },
            { no: 3, name: "repository", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/ },
            { no: 4, name: "branch", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/ },
            { no: 5, name: "commit", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/ },
            { no: 6, name: "file_path", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/ },
            { no: 7, name: "file_content", kind: "scalar", opt: true, T: 12 /*ScalarType.BYTES*/ },
            { no: 8, name: "plain_output", kind: "scalar", opt: true, T: 12 /*ScalarType.BYTES*/ },
            { no: 9, name: "masked_output", kind: "scalar", opt: true, T: 12 /*ScalarType.BYTES*/ },
            { no: 10, name: "mac_address", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/ },
            { no: 11, name: "hostname", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/ },
            { no: 12, name: "platform", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/ },
            { no: 13, name: "release", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/ },
            { no: 14, name: "arch", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/ },
            { no: 15, name: "vendor", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/ },
            { no: 16, name: "shell", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/ },
            { no: 17, name: "vs_app_host", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/ },
            { no: 18, name: "vs_app_name", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/ },
            { no: 19, name: "vs_app_session_id", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/ },
            { no: 20, name: "vs_machine_id", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/ },
            { no: 21, name: "vs_metadata", kind: "map", K: 9 /*ScalarType.STRING*/, V: { kind: "scalar", T: 9 /*ScalarType.STRING*/ } }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message runme.reporter.v1.TransformRequest
 */
export const TransformRequest = new TransformRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class TransformResponse$Type extends MessageType {
    constructor() {
        super("runme.reporter.v1.TransformResponse", [
            { no: 1, name: "notebook", kind: "message", T: () => Notebook },
            { no: 2, name: "extension", kind: "message", T: () => ReporterExtension }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message runme.reporter.v1.TransformResponse
 */
export const TransformResponse = new TransformResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ReporterExtension$Type extends MessageType {
    constructor() {
        super("runme.reporter.v1.ReporterExtension", [
            { no: 1, name: "auto_save", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 2, name: "git", kind: "message", T: () => ReporterGit },
            { no: 3, name: "file", kind: "message", T: () => ReporterFile },
            { no: 4, name: "session", kind: "message", T: () => ReporterSession },
            { no: 5, name: "device", kind: "message", T: () => ReporterDevice }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message runme.reporter.v1.ReporterExtension
 */
export const ReporterExtension = new ReporterExtension$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ReporterGit$Type extends MessageType {
    constructor() {
        super("runme.reporter.v1.ReporterGit", [
            { no: 1, name: "repository", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "branch", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "commit", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message runme.reporter.v1.ReporterGit
 */
export const ReporterGit = new ReporterGit$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ReporterSession$Type extends MessageType {
    constructor() {
        super("runme.reporter.v1.ReporterSession", [
            { no: 1, name: "plain_output", kind: "scalar", T: 12 /*ScalarType.BYTES*/ },
            { no: 2, name: "masked_output", kind: "scalar", T: 12 /*ScalarType.BYTES*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message runme.reporter.v1.ReporterSession
 */
export const ReporterSession = new ReporterSession$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ReporterFile$Type extends MessageType {
    constructor() {
        super("runme.reporter.v1.ReporterFile", [
            { no: 1, name: "path", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "content", kind: "scalar", T: 12 /*ScalarType.BYTES*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message runme.reporter.v1.ReporterFile
 */
export const ReporterFile = new ReporterFile$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ReporterDevice$Type extends MessageType {
    constructor() {
        super("runme.reporter.v1.ReporterDevice", [
            { no: 1, name: "mac_address", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "hostname", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "platform", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 4, name: "release", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 5, name: "arch", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 6, name: "vendor", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 7, name: "shell", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 8, name: "vs_app_host", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 9, name: "vs_app_name", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 10, name: "vs_app_session_id", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 11, name: "vs_machine_id", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 12, name: "vs_metadata", kind: "map", K: 9 /*ScalarType.STRING*/, V: { kind: "scalar", T: 9 /*ScalarType.STRING*/ } }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message runme.reporter.v1.ReporterDevice
 */
export const ReporterDevice = new ReporterDevice$Type();
/**
 * @generated ServiceType for protobuf service runme.reporter.v1.ReporterService
 */
export const ReporterService = new ServiceType("runme.reporter.v1.ReporterService", [
    { name: "Transform", options: {}, I: TransformRequest, O: TransformResponse }
]);
