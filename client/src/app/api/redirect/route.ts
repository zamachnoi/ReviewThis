import { NextFetchEvent, NextRequest, NextResponse } from "next/server"
import https from "https"
import fs from "fs"
import path from "path"
import Error from "next/error"

// Use path.resolve to get the absolute path to the .crt file.

// const crtPath =
// 	"/Users/nick/Desktop/projects/ViewThis/client/certificates/localhost.pem"
// const key =
// 	"/Users/nick/Desktop/projects/ViewThis/client/certificates/localhost-key.pem"
// const caLocalCertificate = fs.readFileSync(crtPath)
// const caLocalPrivateKey = fs.readFileSync(key)
// const localAgent = new https.Agent({
// 	ca: caLocalCertificate,
// 	key: caLocalPrivateKey,
// 	rejectUnauthorized: false, // insecure, only for testing
// })
export async function GET(req: NextRequest) {
	console.log("handleRequest")
	try {
		const cookieResponse = await fetch(
			"https://api.viewthis.app/api/auth/cookie",
			{
				method: "GET",
			}
		)

		const cookieData = await cookieResponse.json()
		console.log("cookieData", cookieData)
	} catch (e) {
		console.log(e)
		return NextResponse.json({ error: "error" })
	}
	const searchParams = req.nextUrl.searchParams
	const redirectUrl = searchParams.has("path")
		? `https://api.viewthis.app/api/${searchParams.get("path")}`
		: "https://api.viewthis.app/api"
	console.log("redirectUrl", redirectUrl)

	return NextResponse.json(["noError"])
}

export async function POST(req: NextRequest) {
	return handleRequest("POST", req)
}

export async function PATCH(req: NextRequest) {
	return handleRequest("PATCH", req)
}

export async function DELETE(req: NextRequest) {
	return handleRequest("DELETE", req)
}

async function handleRequest(method: string, req: NextRequest) {}

// const apiResponse = await fetch(redirectUrl, {
// 	method: method,
// 	headers: req.headers,
// })

// const apiData = await apiResponse.json()

// return new Response(apiData, {
// 	headers: {
// 		"Set-Cookie": `jwt=${apiData._viewthis_jwt}; sameSite=strict; httpOnly=true; expiry`,
// 	},
// })
