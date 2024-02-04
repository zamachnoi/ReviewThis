import Image from "next/image"
import { cookies } from "next/headers"

// fetch data from api.viewthis.app/api/test

async function getData() {
	const res = await fetch("https://api.viewthis.app/api/test", {
		headers: { Cookie: cookies().toString() },
	}) //with credentials: "include"
	// The return value is *not* serialized
	// You can return Date, Map, Set, etc.

	if (!res.ok) {
		console.log(res.status, res.statusText)
		// This will activate the closest `error.js` Error Boundary
		// throw new Error("Failed to fetch data")
		return res.status + " " + res.statusText + " Failed to fetch data"
	}

	return res.json()
}

export default async function Home() {
	let data = await getData()
	console.log(data)
	return (
		<main className="flex min-h-screen flex-col items-center justify-between p-24">
			<p>{JSON.stringify(data)}</p>
		</main>
	)
}
