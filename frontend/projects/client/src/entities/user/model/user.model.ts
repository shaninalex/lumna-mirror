export interface UserModel {
    id: string
    settings: Settings
    code: string
}

export interface Settings {
    theme: string
    language: string
    timezone: string
    date_format: string
    time_format: string
    week_start_day: string
}
