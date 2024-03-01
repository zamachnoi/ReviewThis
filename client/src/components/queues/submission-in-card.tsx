import { Submission, Queue } from "@/lib/types"
import { DiscordAvatar } from "../discord-avatar"
import { User } from "@/lib/types"
import SoundCloudContentPreview from "./soundcloud-content-preview"
import Link from "next/link"
type SubmissionInCardProps = {
	submission: Submission | null
	type?: Queue["type"] | undefined
}

export default function SubmissionInCard({
	submission,
	type,
}: SubmissionInCardProps) {
	if (submission === null) {
		return <div></div>
	}
	let user: User = {
		username: submission.username,
		avatar: submission.avatar,
		discord_id: submission.discord_id,
		db_id: submission.user_id,
	}
	let content = <div>{submission.name}</div>
	if (type === "soundcloud" && submission.content) {
		content = (
			<Link className="underline" href={submission.content}>
				{submission.name}
			</Link>
		)
	}

	return (
		<div className="flex flex-row items-center border-b-2 py-2 px-4 h-[80px] max-h-[80px] justify-between">
			<div className="flex flex-row w-[40%] items-center">
				<div className="w-[33%]">
					<DiscordAvatar user={user} width={50} />
				</div>
				<div className="w-[66%]">{user.username}</div>
			</div>
			<div className="w-fit">{content}</div>
		</div>
	)
}
