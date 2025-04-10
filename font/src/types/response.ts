import {ProjectStatus, RoleType} from "./project.ts";

export interface BaseResponse<T = any> {
    code: number;
    message: string;
    data: T;
}

export type BaseResponseWithoutData = {
    code: number;
    message: string;
}

export type UserInfo = {
    id: number;
    name : string;
    user_name: string;
    email: string;
    avatar: number;
    role_type: RoleType;
}

export type Project = {
    id: number;
    name: string;
    role_type: RoleType
    status: ProjectStatus;
    owner_id: number;
    owner_name: string;
}

export type ProjectList = {
    total: number;
    list: Project[];
}