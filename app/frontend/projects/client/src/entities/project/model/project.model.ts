export interface Project {
	id: number
	name: string
	code: string
	created_at: Date
	updated_at: Date
}

export interface ProjectPatch {
	title: string
}
