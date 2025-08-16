export type ToastType = "success" | "error" | "info" | "warning"

export interface IToast {
    message: string
    header?: string
    type: ToastType
}
