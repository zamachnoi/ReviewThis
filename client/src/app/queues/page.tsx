import { CreateQueueDialog } from "@/components/queues/create-queue"
import { CreateQueueForm } from "@/components/queues/create-queue-form"
import { getData } from "@/app/serverUtils"

export default async function Queues() {
	const data = await getData("/queues")
	console.log(data)

	return (
		<div>
			<CreateQueueDialog />
		</div>
	)
}
