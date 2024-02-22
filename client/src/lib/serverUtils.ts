import { cookies } from "next/headers"
import { getApiUrl } from "@/lib/utils"

export function getServerJwt(): string {
	const cookieStore = cookies()
	const viewthisJwt = cookieStore.get("_viewthis_jwt")
	if (!viewthisJwt) {
		return ""
	}
	return viewthisJwt.value
}

export async function getData(path: string) {
	const apiUrl = getApiUrl(path)

	const res = await fetch(apiUrl, {
		headers: { Cookie: cookies().toString() },
	})

	if (!res.ok) {
		switch (res.status) {
			case 400:
				// Handle bad request error
				throw new Error(
					"Bad Request: The server could not understand the request due to invalid syntax."
				)
			case 401:
				// Handle unauthorized error
				throw new Error(
					"Unauthorized: The request requires user authentication."
				)
			case 403:
				// Handle forbidden error
				throw new Error(
					"Forbidden: The server understood the request, but is refusing to fulfill it."
				)
			case 404:
				// Handle not found error
				throw new Error(
					"Not Found: The server has not found anything matching the Request-URI."
				)
			case 500:
				// Handle internal server error
				throw new Error(
					"Internal Server Error: The server encountered an unexpected condition which prevented it from fulfilling the request."
				)
			default:
				// Handle other statuses
				throw new Error(
					`${res.status} ${res.statusText}: Failed to fetch data.`
				)
		}
	}

	try {
		const body = await res.json()
		if (!body) {
			throw new Error("No response body")
		}
		return body
	} catch (error) {
		console.error("Error parsing JSON:", error)
		throw new Error("Error parsing JSON")
	}
}
