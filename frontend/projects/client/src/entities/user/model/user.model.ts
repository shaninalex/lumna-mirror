import {Identity} from '@ory/kratos-client';


export interface UserModel {
    nickname: string
    identity: Identity
}

export interface Settings {
    theme: Theme
    language: Language
    timezone: string
    date_format: string
    time_format: string
    week_start_day: WeekStartDay
}

export enum Theme {
    Light = "light",
    Dark = "dark",
    Device = "device",
}

export enum Language {
    EN = "EN",
    UA = "UA",
    DE = "DE",
}

export enum WeekStartDay {
    Monday,
    Tuesday,
    Wednesday,
    Thursday,
    Friday,
    Saturday,
    Sunday,
}
