"use client"
import {
	Card,
	CardContent,
	CardDescription,
	CardFooter,
	CardHeader,
	CardTitle,
} from "@/components/ui/card"
import { DiscordAvatar, User } from "../discord-avatar"
import { Button } from "../ui/button"
import { ScrollArea } from "../ui/scroll-area"
export interface Queue {
	ID: number
	createdAt: string
	updatedAt: string
	name: string
	description: string
	type: "soundcloud"
	discord_id: string
	user_id: number
	username: string
	avatar: string
	private: boolean
	submissions: Submission[]
}

export interface Submission {
	id: number
	createdAt: string
	updatedAt: string
	userId: number
	username: string
	avatar: string
	content: string
}

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
			<CardContent>
				<ScrollArea className="w-full p-4 border rounded-md h-[200px]">
					<p>
						{`Jokester began sneaking into the castle in the middle
						of the night and leaving jokes all over the place: under
						the king's pillow, in his soup, even in the royal
						toilet. The king was furious, but he couldn't seem to
						stop Jokester. And then, one day, the people of the
						kingdom discovered that the jokes left by Jokester were
						so funny that they couldn't help but laugh. And once
						they started laughing, they couldn't stop. Jokester
						began sneaking into the castle in the middle of the
						night and leaving jokes all over the place: under the
						king's pillow, in his soup, even in the royal toilet.
						The king was furious, but he couldn't seem to stop
						Jokester. And then, one day, the people of the kingdom
						discovered that the jokes left by Jokester were so funny
						that they couldn't help but laugh. And once they started
						laughing, they couldn't stop. Jokester began sneaking
						into the castle in the middle of the night and leaving
						jokes all over the place: under the king's pillow, in
						his soup, even in the royal toilet. The king was
						furious, but he couldn't seem to stop Jokester. And
						then, one day, the people of the kingdom discovered that
						the jokes left by Jokester were so funny that they
						couldn't help but laugh. And once they started laughing,
						they couldn't stop.`}
					</p>
				</ScrollArea>
			</CardContent>
			<CardFooter>
				<div className="flex flex-row justify-between w-full">
					<Button variant="outline">View</Button>
					<Button>Submit</Button>
				</div>
			</CardFooter>
		</Card>
	)
}
