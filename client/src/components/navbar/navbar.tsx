import { LoginWithDiscord } from "./login-with-discord"
import NavbarLogo from "./navbar-logo"
import { ModeToggle } from "../mode-toggle"
import { getServerJwt } from "@/lib/serverUtils"
import { DiscordAvatar } from "../discord-avatar"
import { jwtDecode } from "jwt-decode"
import AvatarDropdown from "./avatar-dropdown"
import NavLink from "./nav-link"
import { User } from "@/lib/types"

export default function Navbar() {
	const jwt = getServerJwt()
	let user: User | null = null
	if (jwt) {
		user = jwtDecode<User>(jwt)
	}

	return (
		<nav
			style={{ transform: "translate3d(0,0,0)" }}
			className="sticky top-0 z-50 flex items-center justify-between w-full h-16 px-5 shadow-md bg-background backdrop-filter backdrop-blur-lg bg-opacity-60"
		>
			<NavbarLogo />
			<div className="flex flex-row items-center justify-around w-1/3">
				<NavLink href="/queues">Queues</NavLink>
				<NavLink href="/about">About</NavLink>
			</div>
			<div className="flex items-center gap-4">
				<ModeToggle />
				{user ? (
					<AvatarDropdown user={user}>
						{" "}
						<DiscordAvatar user={user} width={48} />
					</AvatarDropdown>
				) : (
					<LoginWithDiscord />
				)}
			</div>
		</nav>
	)
}
