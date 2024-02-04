import Image from "next/image"
import { cookies } from "next/headers"

// fetch data from api.viewthis.app/api/test

async function getData() {
	// Base API URL; adjust as needed for production environment
	let baseUrl =
		process.env.NODE_ENV === "development"
			? "https://localhost:3000"
			: "https://api.viewthis.app"
	let API_URL = `${baseUrl}/api/test`

	const res = await fetch(API_URL, {
		headers: { Cookie: cookies().toString() },
	})

	if (!res.ok) {
		console.log(res.status, res.statusText)
		return `${res.status} ${res.statusText} Failed to fetch data`
	}

	const text = await res.text() // First, convert it to text
	try {
		return text ? JSON.parse(text) : {} // Then, parse it as JSON if not empty
	} catch (error) {
		console.error("Error parsing JSON:", error)
		return { error: "Error parsing JSON" }
	}
}

export default async function Home() {
	let data = await getData()

	console.log(JSON.stringify(data))
	return (
		<main className="flex min-h-screen flex-col items-center justify-between p-24">
			<p>{JSON.stringify(data)}</p>
		</main>
	)
}
