"use client"
import Link from "next/link"
import { Button } from "@/components/ui/button"

type QueueCardFooterProps = {
	ID: number
}

export default function QueueCardFooter(queue: QueueCardFooterProps) {
	return (
		<div className="flex flex-row justify-between w-full">
			<Link href={`/queues/${queue.ID}`}>
				<Button variant="outline">View</Button>
			</Link>
			<Button>Submit</Button>
		</div>
	)
}
