export interface Feedback {
	// Define the properties of the Feedback object here
}
export type User = {
	username: string
	avatar: string
	discord_id: string
	db_id: number
	premium?: boolean
}

export interface Submission {
	ID: number
	CreatedAt: Date
	UpdatedAt: Date
	name: string
	content: string
	user_id: number
	username: string
	avatar: string
	queue_id: number
	private: boolean
	discord_id: string
}
export interface Queue {
	ID: number
	CreatedAt: string
	UpdatedAt: string
	name: string
	description: string
	type: "soundcloud"
	discord_id: string
	user_id: number
	username: string
	avatar: string
	private: boolean
	submissions?: Submission[]
}
