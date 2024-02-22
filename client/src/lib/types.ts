export interface Feedback {
	// Define the properties of the Feedback object here
}

export interface Submission {
	id: number
	createdAt: Date
	updatedAt: Date
	content: string
	userId: number
	username: string
	avatar: string
	queueId: number
	private: boolean
	feedbacks?: Feedback[]
}
export interface Queue {
	ID: number
	createdAt: string
	updatedAt: string
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
