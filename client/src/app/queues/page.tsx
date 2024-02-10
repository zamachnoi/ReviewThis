import { CreateQueueDialog } from "@/components/queues/create-queue"
import { getData } from "@/app/serverUtils"
import QueueCard, { Queue } from "@/components/queues/queue-card"
import QueueCardContainer from "@/components/queues/queue-card-container"

export default async function Queues() {
	return (
		<div className="flex flex-col items-center">
			<CreateQueueDialog />
			<QueueCardContainer />
		</div>
	)
}
