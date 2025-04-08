export interface BaseResponse<T = any> {
    code: number;
    message: string;
    data: T;
}

export type BaseResponseWithoutData = {
    code: number;
    message: string;
}
