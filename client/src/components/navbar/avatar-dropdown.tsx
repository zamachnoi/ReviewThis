"use client"
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

export default function AvatarDropdown({
	children,
}: {
	children: React.ReactNode
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
			</DropdownMenuContent>
		</DropdownMenu>
	)
}
