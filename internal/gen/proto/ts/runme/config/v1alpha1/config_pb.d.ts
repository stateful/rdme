/* eslint-disable */
// @generated by protobuf-ts 2.9.4 with parameter output_javascript,optimize_code_size,long_type_string,add_pb_suffix,ts_nocheck,eslint_disable
// @generated from protobuf file "runme/config/v1alpha1/config.proto" (package "runme.config.v1alpha1", syntax proto3)
// tslint:disable
// @ts-nocheck
import { MessageType } from "@protobuf-ts/runtime";
/**
 * Kernel describes system-level configuration of the runme toolchain.
 *
 * @generated from protobuf message runme.config.v1alpha1.Kernel
 */
export interface Kernel {
    /**
     * log contains the log configuration.
     *
     * @generated from protobuf field: runme.config.v1alpha1.Kernel.Log log = 1;
     */
    log?: Kernel_Log;
    /**
     * @generated from protobuf field: runme.config.v1alpha1.Kernel.Server server = 2;
     */
    server?: Kernel_Server;
}
/**
 * @generated from protobuf message runme.config.v1alpha1.Kernel.Log
 */
export interface Kernel_Log {
    /**
     * enabled indicates whether to enable logging.
     *
     * @generated from protobuf field: bool enabled = 1;
     */
    enabled: boolean;
    /**
     * path is the path to the log output file.
     *
     * @generated from protobuf field: string path = 2;
     */
    path: string;
    /**
     * verbose is the verbosity level of the log.
     *
     * @generated from protobuf field: bool verbose = 3;
     */
    verbose: boolean;
}
/**
 * @generated from protobuf message runme.config.v1alpha1.Kernel.Server
 */
export interface Kernel_Server {
    /**
     * @generated from protobuf field: string address = 1;
     */
    address: string;
    /**
     * @generated from protobuf field: runme.config.v1alpha1.Kernel.Server.TLS tls = 2;
     */
    tls?: Kernel_Server_TLS;
}
/**
 * @generated from protobuf message runme.config.v1alpha1.Kernel.Server.TLS
 */
export interface Kernel_Server_TLS {
    /**
     * @generated from protobuf field: bool enabled = 1;
     */
    enabled: boolean;
    /**
     * @generated from protobuf field: string cert_file = 2;
     */
    certFile: string;
    /**
     * @generated from protobuf field: string key_file = 3;
     */
    keyFile: string;
}
/**
 * Project describes repo-level configuration of the runme toolchain.
 *
 * @generated from protobuf message runme.config.v1alpha1.Project
 */
export interface Project {
    /**
     * @generated from protobuf oneof: source
     */
    source: {
        oneofKind: "root";
        /**
         * root indicates a dir-based source typically including multiple Markdown files.
         *
         * @generated from protobuf field: string root = 1;
         */
        root: string;
    } | {
        oneofKind: "filename";
        /**
         * filename indicates a single Markdown file.
         *
         * @generated from protobuf field: string filename = 2;
         */
        filename: string;
    } | {
        oneofKind: undefined;
    };
    /**
     * env is the environment variables configuration.
     *
     * @generated from protobuf field: runme.config.v1alpha1.Project.Env env = 3;
     */
    env?: Project_Env;
    /**
     * find_repo_upward indicates whether to find the nearest Git repository upward.
     * This is useful to, for example, recognize .gitignore files.
     *
     * @generated from protobuf field: bool find_repo_upward = 4;
     */
    findRepoUpward: boolean;
    /**
     * ignore_paths is a list of paths to ignore relative to dir.
     *
     * @generated from protobuf field: repeated string ignore_paths = 5 [json_name = "ignore"];
     */
    ignorePaths: string[];
    /**
     * disable_gitignore indicates whether to disable the .gitignore file.
     *
     * @generated from protobuf field: bool disable_gitignore = 6;
     */
    disableGitignore: boolean;
    /**
     * filters is a list of filters to apply.
     * Filters can be applied to documents or
     * individual code blocks.
     *
     * @generated from protobuf field: repeated runme.config.v1alpha1.Project.Filter filters = 7;
     */
    filters: Project_Filter[];
}
/**
 * @generated from protobuf message runme.config.v1alpha1.Project.Env
 */
export interface Project_Env {
    /**
     * use_system_env indicates whether to use the system environment variables.
     *
     * @generated from protobuf field: bool use_system_env = 1;
     */
    useSystemEnv: boolean;
    /**
     * sources is a list of files with env.
     *
     * @generated from protobuf field: repeated string sources = 2;
     */
    sources: string[];
}
/**
 * @generated from protobuf message runme.config.v1alpha1.Project.Filter
 */
export interface Project_Filter {
    /**
     * type is the type of the filter.
     *
     * @generated from protobuf field: runme.config.v1alpha1.Project.FilterType type = 1;
     */
    type: Project_FilterType;
    /**
     * condition is the filter program to execute for each document or block,
     * depending on the filter type.
     *
     * The condition should be a valid Expr expression and it should return a boolean value.
     * You can read more about the Expr syntax on https://expr-lang.org/.
     *
     * @generated from protobuf field: string condition = 2;
     */
    condition: string;
}
/**
 * @generated from protobuf enum runme.config.v1alpha1.Project.FilterType
 */
export declare enum Project_FilterType {
    /**
     * @generated from protobuf enum value: FILTER_TYPE_UNSPECIFIED = 0;
     */
    UNSPECIFIED = 0,
    /**
     * @generated from protobuf enum value: FILTER_TYPE_BLOCK = 1;
     */
    BLOCK = 1,
    /**
     * @generated from protobuf enum value: FILTER_TYPE_DOCUMENT = 2;
     */
    DOCUMENT = 2
}
/**
 * Config describes the configuration of the runme toolchain, including CLI, server, and clients like VS Code extension.
 *
 * @generated from protobuf message runme.config.v1alpha1.Config
 */
export interface Config {
    /**
     * kernel is the system-level configuration.
     *
     * @generated from protobuf field: runme.config.v1alpha1.Kernel kernel = 1;
     */
    kernel?: Kernel;
    /**
     * project contains configuration applicable to the project inside a repo.
     *
     * @generated from protobuf field: runme.config.v1alpha1.Project project = 2;
     */
    project?: Project;
}
declare class Kernel$Type extends MessageType<Kernel> {
    constructor();
}
/**
 * @generated MessageType for protobuf message runme.config.v1alpha1.Kernel
 */
export declare const Kernel: Kernel$Type;
declare class Kernel_Log$Type extends MessageType<Kernel_Log> {
    constructor();
}
/**
 * @generated MessageType for protobuf message runme.config.v1alpha1.Kernel.Log
 */
export declare const Kernel_Log: Kernel_Log$Type;
declare class Kernel_Server$Type extends MessageType<Kernel_Server> {
    constructor();
}
/**
 * @generated MessageType for protobuf message runme.config.v1alpha1.Kernel.Server
 */
export declare const Kernel_Server: Kernel_Server$Type;
declare class Kernel_Server_TLS$Type extends MessageType<Kernel_Server_TLS> {
    constructor();
}
/**
 * @generated MessageType for protobuf message runme.config.v1alpha1.Kernel.Server.TLS
 */
export declare const Kernel_Server_TLS: Kernel_Server_TLS$Type;
declare class Project$Type extends MessageType<Project> {
    constructor();
}
/**
 * @generated MessageType for protobuf message runme.config.v1alpha1.Project
 */
export declare const Project: Project$Type;
declare class Project_Env$Type extends MessageType<Project_Env> {
    constructor();
}
/**
 * @generated MessageType for protobuf message runme.config.v1alpha1.Project.Env
 */
export declare const Project_Env: Project_Env$Type;
declare class Project_Filter$Type extends MessageType<Project_Filter> {
    constructor();
}
/**
 * @generated MessageType for protobuf message runme.config.v1alpha1.Project.Filter
 */
export declare const Project_Filter: Project_Filter$Type;
declare class Config$Type extends MessageType<Config> {
    constructor();
}
/**
 * @generated MessageType for protobuf message runme.config.v1alpha1.Config
 */
export declare const Config: Config$Type;
export {};
