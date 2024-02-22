import { getApiUrl } from "@/lib/utils"
import { Button } from "@/components/ui/button"
import { DiscordLogoIcon } from "@radix-ui/react-icons"
import Link from "next/link"
export function LoginWithDiscord() {
	return (
		<Link href={getApiUrl("/auth/discord/login")}>
			<Button className="gap-2">
				<DiscordLogoIcon />
				Login
			</Button>
		</Link>
	)
}
