import React from "react"
import Link from "next/link"
import { ModeToggle } from "./mode-toggle"
import { buttonVariants } from "@/components/ui/button"
import { LoginWithDiscordButton } from "./login-with-discord-button"

export function Navbar() {
	return (
		<>
			<div className="w-full h-20 bg-slate sticky top-0">
				<div className="container mx-auto px-4 h-full flex flex-row justify-between">
					<div className="items-center h-full">
						<ul className="hidden md:flex gap-x-6 text-white"></ul>
						<LoginWithDiscordButton /> {/* Custom component */}
						<ModeToggle /> {/* Custom component */}
					</div>
				</div>
			</div>
		</>
	)
}
