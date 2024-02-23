import { CreateQueueDialog } from "@/components/queues/create-queue"
import { getData } from "@/lib/serverUtils"
import QueueCard from "@/components/queues/queue-card"
import QueueCardContainer from "@/components/queues/queue-card-container"
export default async function Queues({
	searchParams,
}: {
	searchParams?: {
		search?: string
		page?: string
		limit?: string
	}
}) {
	const search = searchParams?.search || ""
	const limit = Number(searchParams?.limit) || 9
	const page = Number(searchParams?.page) || 1
	return (
		<div className="flex flex-col items-center">
			<QueueCardContainer limit={limit} page={page} search={search} />
		</div>
	)
}
