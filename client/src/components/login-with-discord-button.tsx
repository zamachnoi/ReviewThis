import { DiscordLogoIcon } from "@radix-ui/react-icons"
import Link from "next/link"

import { Button } from "@/components/ui/button"

export function LoginWithDiscordButton() {
	return (
		<Button>
			<DiscordLogoIcon className="mr-2 h-4 w-4" />
			<Link href="https://api.viewthis.app/api/auth/discord/login">
				Login
			</Link>
		</Button>
	)
}
