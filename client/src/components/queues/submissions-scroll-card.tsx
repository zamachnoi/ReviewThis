import { Submission } from "@/lib/types"
import { ScrollArea } from "../ui/scroll-area"

type SubmissionsScrollCardProps = {
	queueId: number
}

export default function SubmissionsScrollCard(
	queueId: SubmissionsScrollCardProps
) {
	return (
		<ScrollArea className="h-[200px] rounded-md border p-4">
			<div>{queueId.queueId}</div>
		</ScrollArea>
	)
}
