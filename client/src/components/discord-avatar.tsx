import { User } from "@/lib/types"
export function DiscordAvatar({ user, width }: { user: User; width: number }) {
	return (
		<img
			width={width}
			height={width}
			src={
				user.avatar
					? `https://cdn.discordapp.com/avatars/${user.discord_id}/${user.avatar}.png`
					: `https://cdn.discordapp.com/embed/avatars/${
							Number(user.discord_id) % 5
					  }.png`
			}
			alt={`${user.username}'s avatar`}
			className="rounded-full "
		/>
	)
}
