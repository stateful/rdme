/* eslint-disable */
// @generated by protobuf-ts 2.9.1 with parameter output_javascript,optimize_code_size,long_type_string,add_pb_suffix,ts_nocheck,eslint_disable
// @generated from protobuf file "runme/parser/v1/parser.proto" (package "runme.parser.v1", syntax proto3)
// tslint:disable
// @ts-nocheck
import { MessageType } from "@protobuf-ts/runtime";
/**
 * @generated from protobuf message runme.parser.v1.Notebook
 */
export interface Notebook {
    /**
     * @generated from protobuf field: repeated runme.parser.v1.Cell cells = 1;
     */
    cells: Cell[];
    /**
     * @generated from protobuf field: map<string, string> metadata = 2;
     */
    metadata: {
        [key: string]: string;
    };
    /**
     * @generated from protobuf field: runme.parser.v1.Frontmatter frontmatter = 3;
     */
    frontmatter?: Frontmatter;
}
/**
 * @generated from protobuf message runme.parser.v1.TextRange
 */
export interface TextRange {
    /**
     * @generated from protobuf field: uint32 start = 1;
     */
    start: number;
    /**
     * @generated from protobuf field: uint32 end = 2;
     */
    end: number;
}
/**
 * @generated from protobuf message runme.parser.v1.Cell
 */
export interface Cell {
    /**
     * @generated from protobuf field: runme.parser.v1.CellKind kind = 1;
     */
    kind: CellKind;
    /**
     * @generated from protobuf field: string value = 2;
     */
    value: string;
    /**
     * @generated from protobuf field: string language_id = 3;
     */
    languageId: string;
    /**
     * @generated from protobuf field: map<string, string> metadata = 4;
     */
    metadata: {
        [key: string]: string;
    };
    /**
     * @generated from protobuf field: runme.parser.v1.TextRange text_range = 5;
     */
    textRange?: TextRange;
}
/**
 * @generated from protobuf message runme.parser.v1.Runme
 */
export interface Runme {
    /**
     * @generated from protobuf field: string version = 1;
     */
    version: string;
    /**
     * @generated from protobuf field: string id = 2;
     */
    id: string;
}
/**
 * @generated from protobuf message runme.parser.v1.Frontmatter
 */
export interface Frontmatter {
    /**
     * @generated from protobuf field: string shell = 1;
     */
    shell: string;
    /**
     * @generated from protobuf field: string cwd = 2;
     */
    cwd: string;
    /**
     * @generated from protobuf field: bool skip_prompts = 3;
     */
    skipPrompts: boolean;
    /**
     * @generated from protobuf field: runme.parser.v1.Runme runme = 4;
     */
    runme?: Runme;
}
/**
 * @generated from protobuf message runme.parser.v1.DeserializeRequest
 */
export interface DeserializeRequest {
    /**
     * @generated from protobuf field: bytes source = 1;
     */
    source: Uint8Array;
}
/**
 * @generated from protobuf message runme.parser.v1.DeserializeResponse
 */
export interface DeserializeResponse {
    /**
     * @generated from protobuf field: runme.parser.v1.Notebook notebook = 1;
     */
    notebook?: Notebook;
}
/**
 * @generated from protobuf message runme.parser.v1.SerializeRequest
 */
export interface SerializeRequest {
    /**
     * @generated from protobuf field: runme.parser.v1.Notebook notebook = 1;
     */
    notebook?: Notebook;
    /**
     * @generated from protobuf field: runme.parser.v1.RunmeIdentity identity = 2;
     */
    identity: RunmeIdentity;
}
/**
 * @generated from protobuf message runme.parser.v1.SerializeResponse
 */
export interface SerializeResponse {
    /**
     * @generated from protobuf field: bytes result = 1;
     */
    result: Uint8Array;
}
/**
 * @generated from protobuf enum runme.parser.v1.CellKind
 */
export declare enum CellKind {
    /**
     * @generated from protobuf enum value: CELL_KIND_UNSPECIFIED = 0;
     */
    UNSPECIFIED = 0,
    /**
     * @generated from protobuf enum value: CELL_KIND_MARKUP = 1;
     */
    MARKUP = 1,
    /**
     * @generated from protobuf enum value: CELL_KIND_CODE = 2;
     */
    CODE = 2
}
/**
 * @generated from protobuf enum runme.parser.v1.RunmeIdentity
 */
export declare enum RunmeIdentity {
    /**
     * aka NONE
     *
     * @generated from protobuf enum value: RUNME_IDENTITY_UNSPECIFIED = 0;
     */
    UNSPECIFIED = 0,
    /**
     * @generated from protobuf enum value: RUNME_IDENTITY_ALL = 1;
     */
    ALL = 1,
    /**
     * @generated from protobuf enum value: RUNME_IDENTITY_DOCUMENT = 2;
     */
    DOCUMENT = 2,
    /**
     * @generated from protobuf enum value: RUNME_IDENTITY_CELL = 3;
     */
    CELL = 3
}
declare class Notebook$Type extends MessageType<Notebook> {
    constructor();
}
/**
 * @generated MessageType for protobuf message runme.parser.v1.Notebook
 */
export declare const Notebook: Notebook$Type;
declare class TextRange$Type extends MessageType<TextRange> {
    constructor();
}
/**
 * @generated MessageType for protobuf message runme.parser.v1.TextRange
 */
export declare const TextRange: TextRange$Type;
declare class Cell$Type extends MessageType<Cell> {
    constructor();
}
/**
 * @generated MessageType for protobuf message runme.parser.v1.Cell
 */
export declare const Cell: Cell$Type;
declare class Runme$Type extends MessageType<Runme> {
    constructor();
}
/**
 * @generated MessageType for protobuf message runme.parser.v1.Runme
 */
export declare const Runme: Runme$Type;
declare class Frontmatter$Type extends MessageType<Frontmatter> {
    constructor();
}
/**
 * @generated MessageType for protobuf message runme.parser.v1.Frontmatter
 */
export declare const Frontmatter: Frontmatter$Type;
declare class DeserializeRequest$Type extends MessageType<DeserializeRequest> {
    constructor();
}
/**
 * @generated MessageType for protobuf message runme.parser.v1.DeserializeRequest
 */
export declare const DeserializeRequest: DeserializeRequest$Type;
declare class DeserializeResponse$Type extends MessageType<DeserializeResponse> {
    constructor();
}
/**
 * @generated MessageType for protobuf message runme.parser.v1.DeserializeResponse
 */
export declare const DeserializeResponse: DeserializeResponse$Type;
declare class SerializeRequest$Type extends MessageType<SerializeRequest> {
    constructor();
}
/**
 * @generated MessageType for protobuf message runme.parser.v1.SerializeRequest
 */
export declare const SerializeRequest: SerializeRequest$Type;
declare class SerializeResponse$Type extends MessageType<SerializeResponse> {
    constructor();
}
/**
 * @generated MessageType for protobuf message runme.parser.v1.SerializeResponse
 */
export declare const SerializeResponse: SerializeResponse$Type;
/**
 * @generated ServiceType for protobuf service runme.parser.v1.ParserService
 */
export declare const ParserService: any;
export {};
