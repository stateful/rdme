/* eslint-disable */
// @generated by protobuf-ts 2.9.1 with parameter output_javascript,optimize_code_size,long_type_string,add_pb_suffix,ts_nocheck,eslint_disable
// @generated from protobuf file "runme/parser/v1/parser.proto" (package "runme.parser.v1", syntax proto3)
// tslint:disable
// @ts-nocheck
/* eslint-disable */
// @generated by protobuf-ts 2.9.1 with parameter output_javascript,optimize_code_size,long_type_string,add_pb_suffix,ts_nocheck,eslint_disable
// @generated from protobuf file "runme/parser/v1/parser.proto" (package "runme.parser.v1", syntax proto3)
// tslint:disable
// @ts-nocheck
import { ServiceType } from "@protobuf-ts/runtime-rpc";
import { MessageType } from "@protobuf-ts/runtime";
/**
 * @generated from protobuf enum runme.parser.v1.CellKind
 */
export var CellKind;
(function (CellKind) {
    /**
     * @generated from protobuf enum value: CELL_KIND_UNSPECIFIED = 0;
     */
    CellKind[CellKind["UNSPECIFIED"] = 0] = "UNSPECIFIED";
    /**
     * @generated from protobuf enum value: CELL_KIND_MARKUP = 1;
     */
    CellKind[CellKind["MARKUP"] = 1] = "MARKUP";
    /**
     * @generated from protobuf enum value: CELL_KIND_CODE = 2;
     */
    CellKind[CellKind["CODE"] = 2] = "CODE";
})(CellKind || (CellKind = {}));
// @generated message type with reflection information, may provide speed optimized methods
class Notebook$Type extends MessageType {
    constructor() {
        super("runme.parser.v1.Notebook", [
            { no: 1, name: "cells", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => Cell },
            { no: 2, name: "metadata", kind: "map", K: 9 /*ScalarType.STRING*/, V: { kind: "scalar", T: 9 /*ScalarType.STRING*/ } },
            { no: 3, name: "frontmatter", kind: "message", T: () => Frontmatter }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message runme.parser.v1.Notebook
 */
export const Notebook = new Notebook$Type();
// @generated message type with reflection information, may provide speed optimized methods
class TextRange$Type extends MessageType {
    constructor() {
        super("runme.parser.v1.TextRange", [
            { no: 1, name: "start", kind: "scalar", T: 13 /*ScalarType.UINT32*/ },
            { no: 2, name: "end", kind: "scalar", T: 13 /*ScalarType.UINT32*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message runme.parser.v1.TextRange
 */
export const TextRange = new TextRange$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Cell$Type extends MessageType {
    constructor() {
        super("runme.parser.v1.Cell", [
            { no: 1, name: "kind", kind: "enum", T: () => ["runme.parser.v1.CellKind", CellKind, "CELL_KIND_"] },
            { no: 2, name: "value", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "language_id", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 4, name: "metadata", kind: "map", K: 9 /*ScalarType.STRING*/, V: { kind: "scalar", T: 9 /*ScalarType.STRING*/ } },
            { no: 5, name: "text_range", kind: "message", T: () => TextRange }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message runme.parser.v1.Cell
 */
export const Cell = new Cell$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Runme$Type extends MessageType {
    constructor() {
        super("runme.parser.v1.Runme", [
            { no: 1, name: "version", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "id", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message runme.parser.v1.Runme
 */
export const Runme = new Runme$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Frontmatter$Type extends MessageType {
    constructor() {
        super("runme.parser.v1.Frontmatter", [
            { no: 1, name: "shell", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "cwd", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "skip_prompts", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 4, name: "runme", kind: "message", T: () => Runme }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message runme.parser.v1.Frontmatter
 */
export const Frontmatter = new Frontmatter$Type();
// @generated message type with reflection information, may provide speed optimized methods
class DeserializeRequest$Type extends MessageType {
    constructor() {
        super("runme.parser.v1.DeserializeRequest", [
            { no: 1, name: "source", kind: "scalar", T: 12 /*ScalarType.BYTES*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message runme.parser.v1.DeserializeRequest
 */
export const DeserializeRequest = new DeserializeRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class DeserializeResponse$Type extends MessageType {
    constructor() {
        super("runme.parser.v1.DeserializeResponse", [
            { no: 1, name: "notebook", kind: "message", T: () => Notebook }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message runme.parser.v1.DeserializeResponse
 */
export const DeserializeResponse = new DeserializeResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class SerializeRequest$Type extends MessageType {
    constructor() {
        super("runme.parser.v1.SerializeRequest", [
            { no: 1, name: "notebook", kind: "message", T: () => Notebook }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message runme.parser.v1.SerializeRequest
 */
export const SerializeRequest = new SerializeRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class SerializeResponse$Type extends MessageType {
    constructor() {
        super("runme.parser.v1.SerializeResponse", [
            { no: 1, name: "result", kind: "scalar", T: 12 /*ScalarType.BYTES*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message runme.parser.v1.SerializeResponse
 */
export const SerializeResponse = new SerializeResponse$Type();
/**
 * @generated ServiceType for protobuf service runme.parser.v1.ParserService
 */
export const ParserService = new ServiceType("runme.parser.v1.ParserService", [
    { name: "Deserialize", options: {}, I: DeserializeRequest, O: DeserializeResponse },
    { name: "Serialize", options: {}, I: SerializeRequest, O: SerializeResponse }
]);
