import { type ClassValue, clsx } from "clsx"
import { twMerge } from "tailwind-merge"

export function cn(...inputs: ClassValue[]) {
	return twMerge(clsx(inputs))
}

export const fetcher = (url: string) =>
	fetch(url, {
		credentials: "include",
	}).then((r) => r.json())

export function getApiUrl(path: string) {
	return process.env.NODE_ENV === "development"
		? "http://localhost:3001/api" + path
		: "https://api.viewthis.app/api" + path
}
