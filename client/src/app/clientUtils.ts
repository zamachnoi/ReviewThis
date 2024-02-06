import { useCookies } from "next-client-cookies"

export function getClientApiUrl(path: string) {
	return process.env.NODE_ENV === "development"
		? "http://127.0.0.1:3001/api" + path
		: "https://api.viewthis.app/api" + path
}

export function parseJwt(token: string) {
	return JSON.parse(Buffer.from(token.split(".")[1], "base64").toString())
}

export function getClientJwt() {
	const cookies = useCookies()

	const viewthisJwt = cookies.get("_viewthis_jwt")

	return viewthisJwt ? viewthisJwt : ""
}

export function getClientJwtAndParse() {
	const jwt = getClientJwt()
	return jwt ? parseJwt(jwt) : ""
}
