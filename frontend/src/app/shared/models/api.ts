import { HttpErrorResponse } from "@angular/common/http";

export interface APIResponse<T> {
    messages: string[];
    status: boolean;
    data: T;
    errors: Error[];
}

export interface Error {
    // detailed message to display in ui
    message: string;

    // additional information:
    // - detailed instructions for user
    // - additional error objects
    // - helpers
    meta?: any;

    // code - uppercase underscored error code
    // for example: AUTH_USER_NOT_FOUND, PROJECT_NOT_FOUND
    // Need for error handlers
    code: string;
}

export function fromErrorResponse(resp: HttpErrorResponse): Error[] {
    const responseData: APIResponse<any> = resp.error;

    return responseData.errors;
}
