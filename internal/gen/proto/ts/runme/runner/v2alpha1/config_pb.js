/* eslint-disable */
// @generated by protobuf-ts 2.9.3 with parameter output_javascript,optimize_code_size,long_type_string,add_pb_suffix,ts_nocheck,eslint_disable
// @generated from protobuf file "runme/runner/v2alpha1/config.proto" (package "runme.runner.v2alpha1", syntax proto3)
// tslint:disable
// @ts-nocheck
/* eslint-disable */
// @generated by protobuf-ts 2.9.3 with parameter output_javascript,optimize_code_size,long_type_string,add_pb_suffix,ts_nocheck,eslint_disable
// @generated from protobuf file "runme/runner/v2alpha1/config.proto" (package "runme.runner.v2alpha1", syntax proto3)
// tslint:disable
// @ts-nocheck
import { MessageType } from "@protobuf-ts/runtime";
/**
 * @generated from protobuf enum runme.runner.v2alpha1.CommandMode
 */
export var CommandMode;
(function (CommandMode) {
    /**
     * @generated from protobuf enum value: COMMAND_MODE_UNSPECIFIED = 0;
     */
    CommandMode[CommandMode["UNSPECIFIED"] = 0] = "UNSPECIFIED";
    /**
     * @generated from protobuf enum value: COMMAND_MODE_INLINE = 1;
     */
    CommandMode[CommandMode["INLINE"] = 1] = "INLINE";
    /**
     * @generated from protobuf enum value: COMMAND_MODE_FILE = 2;
     */
    CommandMode[CommandMode["FILE"] = 2] = "FILE";
})(CommandMode || (CommandMode = {}));
// @generated message type with reflection information, may provide speed optimized methods
class ProgramConfig$Type extends MessageType {
    constructor() {
        super("runme.runner.v2alpha1.ProgramConfig", [
            { no: 1, name: "program_name", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "arguments", kind: "scalar", repeat: 2 /*RepeatType.UNPACKED*/, T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "directory", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 4, name: "env", kind: "scalar", repeat: 2 /*RepeatType.UNPACKED*/, T: 9 /*ScalarType.STRING*/ },
            { no: 5, name: "commands", kind: "message", oneof: "source", T: () => ProgramConfig_CommandList },
            { no: 6, name: "script", kind: "scalar", oneof: "source", T: 9 /*ScalarType.STRING*/ },
            { no: 7, name: "interactive", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 8, name: "mode", kind: "enum", T: () => ["runme.runner.v2alpha1.CommandMode", CommandMode, "COMMAND_MODE_"] }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message runme.runner.v2alpha1.ProgramConfig
 */
export const ProgramConfig = new ProgramConfig$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ProgramConfig_CommandList$Type extends MessageType {
    constructor() {
        super("runme.runner.v2alpha1.ProgramConfig.CommandList", [
            { no: 1, name: "items", kind: "scalar", repeat: 2 /*RepeatType.UNPACKED*/, T: 9 /*ScalarType.STRING*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message runme.runner.v2alpha1.ProgramConfig.CommandList
 */
export const ProgramConfig_CommandList = new ProgramConfig_CommandList$Type();
