"use client"
import { Input } from "../ui/input"

type QueueSearchBarProps = {
	limit: number
	page: number
	search: string
}

import React from "react" // Add the missing import statement

export function QueueSearchBar({ limit, page, search }: QueueSearchBarProps) {
	let url = `/queues?limit=${limit}&page=${1}`

	const handleSearch: React.KeyboardEventHandler<HTMLInputElement> = (
		event
	) => {
		if (event.key === "Enter") {
			event.preventDefault() // Prevent form submission
			const value = event.currentTarget.value
			const newUrl = `${url}&search=${value}`
			window.location.href = newUrl
		}
	}

	return (
		<div className="flex flex-row justify-between w-full">
			<Input
				type="search"
				placeholder="Search"
				onKeyDown={handleSearch}
			/>
		</div>
	)
}
