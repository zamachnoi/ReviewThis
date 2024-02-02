"use client"
import { DiscordLogoIcon } from "@radix-ui/react-icons"

import { Button } from "@/components/ui/button"
import Link from "next/link"

export function LoginWithDiscordButton() {
	const handleClick = () => {}
	return (
		<Link href="https://api.viewthis.app/api/auth/discord/login">
			<Button>
				<DiscordLogoIcon className="mr-2 h-4 w-4" />
				Login
			</Button>
		</Link>
	)
}
