export interface Comment {
	id: number
	entity_id: number
	entity_type: string // enum type
	user_id: number
	content: string
	created_at: Date
}
