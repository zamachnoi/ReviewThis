import { getData } from "@/lib/serverUtils"
import QueueCard from "./queue-card"
import { Queue } from "@/lib/types"
import { MyPagination } from "./queue-pagination"
import { CreateQueueDialog } from "./create-queue"
import { QueueSearchBar } from "./queue-search-bar"

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
	let queues: Queue[] = res.queues
	if (!queues) {
		queues = []
	}
	const totalPages = res.totalPages
	return (
		<div className="inline-flex flex-col items-center justify-center gap-4 w-full">
			<div className="w-full mt-4 flex flex-row items-center justify-evenly">
				<h1 className="text-4xl font-bold">queues</h1>
			</div>
			<div className="w-1/6">
				<QueueSearchBar limit={limit} page={page} search={search} />
			</div>
			<div className="w-[80%] flex flex-wrap gap-4 justify-center">
				{queues &&
					queues.map((queue: Queue) => {
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
			<div className="w-4/5 flex flex-row justify-end fixed bottom-4 right-4">
				<CreateQueueDialog />
			</div>
		</div>
	)
}
