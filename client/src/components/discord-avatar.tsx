import Image from "next/image"

export type User = {
	username: string
	avatar: string
	discord_id: string
	db_id: string
}

export function DiscordAvatar({ user }: { user: User }) {
	return (
		<img
			width={48}
			height={48}
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
