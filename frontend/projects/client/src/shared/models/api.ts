export interface APIResponse<T> {
    messages: string[]
    status: boolean
    data: T
    error: any // TODO: get apperror type from backend
}
