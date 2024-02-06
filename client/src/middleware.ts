const protectedRoutes = ["/queues/create", "/profile"]
import { NextResponse, NextRequest } from "next/server"

export function middleware(req: NextRequest) {
	if (
		protectedRoutes.includes(req.nextUrl.pathname) &&
		req.cookies.get("_viewthis_jwt") === undefined
	) {
		return NextResponse.redirect(new URL("/", req.url))
	}
}
