import { Queue, Submission } from "@/lib/types"
import { ScrollArea } from "../ui/scroll-area"
import { fetcher, getApiUrl } from "@/lib/utils"
import useSWR from "swr"
import SubmissionInCard from "./submission-in-card"

type SubmissionsScrollCardProps = {
	queueId: number
	queueType: Queue["type"]
}

export default function SubmissionsScrollCard(
	queueId: SubmissionsScrollCardProps
) {
	const { data } = useSWR<Submission[]>(
		getApiUrl(`/queues/${queueId.queueId}/submissions?limit=5&page=1`),
		fetcher
	)

	return (
		<ScrollArea className="h-[200px] rounded-md border px-2">
			<div className="flex flex-col justify-center">
				{!data ? (
					<SubmissionInCard submission={null} />
				) : (
					data.map((submission) => (
						<SubmissionInCard
							key={submission.ID}
							submission={submission}
							type={queueId.queueType}
						/>
					))
				)}
			</div>
		</ScrollArea>
	)
}
