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

export default function QueueCard({ queue }: { queue: Queue }) {
	const user: User = {
		username: queue.username,
		avatar: queue.avatar,
		discord_id: queue.discord_id,
		db_id: queue.user_id,
	}
	return (
		<Card className="flex flex-col">
			<CardHeader className="mb-2 border-b">
				<div className="flex flex-col items-center gap-2">
					<h1 className="text-sm text-muted-foreground">
						{queue.username}
					</h1>
					<div className="flex flex-row items-center justify-center gap-4 w-[60vw] md:w-[40vw] lg:w-[30vw]">
						<DiscordAvatar user={user} width={60} />
						<CardTitle>{queue.name}</CardTitle>
					</div>
				</div>
			</CardHeader>
			<CardContent className="px-2">
				<SubmissionsScrollCard
					queueId={queue.ID}
					queueType={queue.type}
				/>
			</CardContent>
			<CardFooter>
				<QueueCardFooter ID={queue.ID} />
			</CardFooter>
		</Card>
	)
}
