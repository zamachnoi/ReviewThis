"use client"
import { useCookies } from "next-client-cookies"
import { getApiUrl } from "@/lib/utils"
import {
	DropdownMenu,
	DropdownMenuContent,
	DropdownMenuItem,
	DropdownMenuLabel,
	DropdownMenuSeparator,
	DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"
import Link from "next/link"
import { User } from "@/lib/types"

type AvatarDropdownProps = {
	user: User
}
export default function AvatarDropdown({
	children,
	user,
}: {
	children: React.ReactNode
	user: User
}) {
	return (
		<DropdownMenu>
			<DropdownMenuTrigger className="rounded-full">
				{children}
			</DropdownMenuTrigger>
			<DropdownMenuContent>
				<DropdownMenuLabel>My Account</DropdownMenuLabel>
				<DropdownMenuSeparator />
				<Link href={getApiUrl("/auth/discord/logout")}>
					<DropdownMenuItem className="hover:cursor-pointer">
						Logout
					</DropdownMenuItem>
				</Link>
				<Link href={`/users/${user.db_id}`}>
					<DropdownMenuItem className="hover:cursor-pointer">
						Profile
					</DropdownMenuItem>
				</Link>
			</DropdownMenuContent>
		</DropdownMenu>
	)
}
