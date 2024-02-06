"use client"
import {
	DropdownMenu,
	DropdownMenuContent,
	DropdownMenuItem,
	DropdownMenuLabel,
	DropdownMenuSeparator,
	DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"

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
				<DropdownMenuItem>Profile</DropdownMenuItem>
			</DropdownMenuContent>
		</DropdownMenu>
	)
}
