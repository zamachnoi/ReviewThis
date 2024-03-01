"use client"

import { Button } from "../ui/button"
import Link from "next/link"
import { DiscordLogoIcon } from "@radix-ui/react-icons"
import { getApiUrl } from "@/lib/utils"
import { User } from "@/lib/types"

export default function AddBot() {
	return (
		<Link href={getApiUrl("/protected/auth/discord/bot/add")}>
			<Button className="gap-2">
				<DiscordLogoIcon />
				Add Bot
			</Button>
		</Link>
	)
}
