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
    username: string;
    email: string;
    avatar: number;
}

export type Project = {
    id: number;
    title: string;
    description: string;
    status: string;
    created_at: Date;
    updated_at: Date;
}

export type ProjectList = {
    total: number;
    list: Project[];
}