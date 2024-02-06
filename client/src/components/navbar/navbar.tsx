import { LoginWithDiscord } from "./login-with-discord"
import NavbarLogo from "./navbar-logo"
import { ModeToggle } from "../mode-toggle"
import { getServerJwt } from "@/app/serverUtils"
import { DiscordAvatar, User } from "../discord-avatar"
import { jwtDecode } from "jwt-decode"
import AvatarDropdown from "./avatar-dropdown"

export default function Navbar() {
	const jwt = getServerJwt()
	let user: User | null = null
	if (jwt) {
		user = jwtDecode<User>(jwt)
	}

	return (
		<nav
			style={{ transform: "translate3d(0,0,0)" }}
			className="flex items-center px-5 justify-between h-16 bg-background"
		>
			<NavbarLogo />
			<div className="gap-4 flex items-center bg-background">
				<ModeToggle />
				{user ? (
					<AvatarDropdown>
						{" "}
						<DiscordAvatar user={user} />
					</AvatarDropdown>
				) : (
					<LoginWithDiscord />
				)}
			</div>
		</nav>
	)
}
