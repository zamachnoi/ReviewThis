import { getData } from "@/lib/serverUtils"
import QueueCard from "./queue-card"
import { Queue } from "@/lib/types"

export default async function QueueCardContainer({}) {
	const res: Queue[] = await getData("/queues")
	return (
		<div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
			{res.map((queue: Queue) => {
				return <QueueCard key={queue.ID} queue={queue} />
			})}
		</div>
	)
}
