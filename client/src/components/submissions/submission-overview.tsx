type SubmissionOverviewProps = {
	submission: SubmissionOverview
}

type SubmissionOverview = {
	id: number
	name: string
	description: string
	type: "soundcloud"
	avatar: string
	username: string
	user_id: number
}

export default function SubmissionOverview() {}
