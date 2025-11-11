export interface UserModel {
	id: number
	code: string
	email: string
	active: boolean
	state: string
	settings: Settings
}

export interface Settings {
	theme: string
	language: string
	timezone: string
	date_format: string
	time_format: string
	week_start_day: string
}
