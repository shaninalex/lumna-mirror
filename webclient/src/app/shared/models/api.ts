export interface APIResponse<T> {
    messages: string[]
    status: boolean
    data: T
    errors: string[]
}
