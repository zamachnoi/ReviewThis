import {
	Card,
	CardContent,
	CardFooter,
	CardHeader,
	CardTitle,
} from "@/components/ui/card"
import { DiscordAvatar, User } from "../discord-avatar"
import { Queue } from "@/lib/types"
import SubmissionsScrollCard from "./submissions-scroll-card"
import QueueCardFooter from "./queue-card-footer"
import { Button } from "../ui/button"
import Link from "next/link"

export default function QueueCard({ queue }: { queue: Queue }) {
	const user: User = {
		username: queue.username,
		avatar: queue.avatar,
		discord_id: queue.discord_id,
		db_id: queue.user_id,
	}

	const date = new Date(queue.CreatedAt)

	const options: Intl.DateTimeFormatOptions = {
		year: "numeric",
		month: "short",
		day: "numeric",
	}
	const formattedDate = date.toLocaleDateString(undefined, options)
	const formattedTime = date.toLocaleTimeString()

	const finalDateTime = `${formattedDate}, ${formattedTime}`
	return (
		<Card className="flex flex-col w-[30%]">
			<CardHeader className="border-b p-2 pb-4">
				<div className="flex flex-col items-center gap-1">
					<div className="flex flex-row items-center w-full justify-around">
						<h1 className="text-sm text-muted-foreground">
							{queue.username}
						</h1>
						<h1 className="text-xs text-muted-foreground">
							{finalDateTime}
						</h1>
					</div>
					<div className="flex flex-row items-center justify-center gap-4 w-full">
						<DiscordAvatar user={user} width={50} />
						<Link href={`/queues/${queue.ID}`}>
							<Button variant="link">
								<CardTitle>{queue.name}</CardTitle>{" "}
							</Button>
						</Link>
					</div>
				</div>
			</CardHeader>
			<CardContent className="px-2 py-2">
				<SubmissionsScrollCard
					queueId={queue.ID}
					queueType={queue.type}
				/>
			</CardContent>
			{/* <CardFooter className="px-4 pb-2">
				<QueueCardFooter ID={queue.ID} />
			</CardFooter> */}
		</Card>
	)
}
