import Image from "next/image"

// fetch data from api.viewthis.app/api/test

async function getData() {
	const res = await fetch("https://api.viewthis.app/api/test", {
		credentials: "include",
	}) //with credentials: "include"
	// The return value is *not* serialized
	// You can return Date, Map, Set, etc.

	if (!res.ok) {
		console.log(res.status, res.statusText)
		// This will activate the closest `error.js` Error Boundary
		// throw new Error("Failed to fetch data")
		return "Failed to fetch data"
	}

	return res.json()
}

export default async function Home() {
	let data = await getData()
	return (
		<main className="flex min-h-screen flex-col items-center justify-between p-24">
			<p>{data}</p>
		</main>
	)
}
