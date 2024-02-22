import { getData } from "@/app/serverUtils"
import QueueCard from "./queue-card"
import useSWR, { Fetcher } from "swr"
import { getApiUrl } from "@/app/utils"
import { Queue } from "@/lib/types"

const fetcher = (url: string) => fetch(url).then((r) => r.json())
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
