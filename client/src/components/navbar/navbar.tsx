import { LoginWithDiscord } from "./login-with-discord"
import NavbarLogo from "./navbar-logo"
import { ModeToggle } from "../mode-toggle"
import NavMenu from "./nav-menu"
import { getServerJwt } from "@/app/serverUtils"
import { DiscordAvatar, User } from "../discord-avatar"
import { jwtDecode } from "jwt-decode"
import { decode } from "punycode"
import AvatarDropdown from "./avatar-dropdown"

export default function Navbar() {
	const jwt = getServerJwt()
	let user: User | null = null
	if (jwt) {
		user = jwtDecode<User>(jwt)
	}
	return (
		<nav className="flex flex-row justify-between my-2 items-center mx-8">
			hi
		</nav>
	)
}
