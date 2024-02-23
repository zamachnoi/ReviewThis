import { getData } from "@/lib/serverUtils"
import QueueCard from "./queue-card"
import { Queue } from "@/lib/types"
import { MyPagination } from "./queue-pagination"
import { CreateQueueDialog } from "./create-queue"

// limit and page props
type QueueCardContainerProps = {
	limit: number
	page: number
	search: string
}

export default async function QueueCardContainer({
	limit,
	page,
	search,
}: QueueCardContainerProps) {
	console.log(limit, page, search)
	let url = `/queues?limit=${limit}&page=${page}`
	if (search) {
		url += `&search=${search}`
	}
	const res = await getData(url)
	console.log(res)
	const queues: Queue[] = res.queues
	const totalPages = res.totalPages
	return (
		<div className="inline-flex flex-col justify-center gap-4">
			<div>
				<CreateQueueDialog />
			</div>
			<div className="flex flex-wrap gap-2 justify-center">
				{queues.map((queue: Queue) => {
					return <QueueCard key={queue.ID} queue={queue} />
				})}
			</div>

			<MyPagination
				totalPages={totalPages}
				currentPage={page}
				limit={limit}
				baseUrl="/queues"
				search={search}
			/>
		</div>
	)
}
