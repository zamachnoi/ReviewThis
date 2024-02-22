"use client"
import {
	Card,
	CardContent,
	CardFooter,
	CardHeader,
	CardTitle,
} from "@/components/ui/card"
import { DiscordAvatar, User } from "../discord-avatar"
import { Button } from "../ui/button"
import Link from "next/link"
import { Queue } from "@/lib/types"
import SubmissionsScrollCard from "./submissions-scroll-card"

export default function QueueCard({ queue }: { queue: Queue }) {
	// get discord profile picture from avatar and id
	const user: User = {
		username: queue.username,
		avatar: queue.avatar,
		discord_id: queue.discord_id,
		db_id: queue.user_id.toString(),
	}
	return (
		<Card className="flex flex-col">
			<CardHeader className="mb-2 border-b">
				<div className="flex flex-row items-center justify-center gap-4 w-[60vw] md:w-[40vw] lg:w-[30vw]">
					<DiscordAvatar user={user} width={80} />
					<CardTitle>{queue.name}</CardTitle>
				</div>
			</CardHeader>
			<CardContent className="px-2">
				<SubmissionsScrollCard queueId={queue.ID} />
			</CardContent>
			<CardFooter>
				<div className="flex flex-row justify-between w-full">
					<Link href={`/queues/${queue.ID}`}>
						<Button variant="outline">View</Button>
					</Link>
					<Button>Submit</Button>
				</div>
			</CardFooter>
		</Card>
	)
}
