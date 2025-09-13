import {AppError} from '@client/shared/common/errors';

export interface APIResponse<T> {
    messages: string[]
    status: boolean
    data: T
    errors: AppError[]
}
