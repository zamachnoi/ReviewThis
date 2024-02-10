"use client"
import { getData } from "@/app/serverUtils"
import QueueCard, { Queue } from "./queue-card"
import useSWR, { Fetcher } from "swr"
import { getApiUrl } from "@/app/utils"

const fetcher = (url: string) => fetch(url).then((r) => r.json())

export default function QueueCardContainer({}) {
	const queueUrl = getApiUrl("/queues")
	const { data, isLoading } = useSWR(queueUrl, fetcher)

	if (isLoading) return <div>Loading...</div>

	return (
		<div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
			{data.map((queue: Queue) => {
				return <QueueCard key={queue.ID} queue={queue} />
			})}
		</div>
	)
}
