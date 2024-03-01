import { Submission } from "@/lib/types"
import SubmissionInCard from "../submission-in-card"

export default function SubmissionsContainer(
	submissions: Submission[],
	queueType: "soundcloud"
) {
	return (
		<div className="flex flex-col justify-center">
			{!submissions ? (
				<SubmissionInCard submission={null} />
			) : (
				submissions.map((submission) => (
					<SubmissionInCard
						key={submission.ID}
						submission={submission}
						type={queueType}
					/>
				))
			)}
		</div>
	)
}
