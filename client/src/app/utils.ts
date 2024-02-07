export function getApiUrl(path: string) {
	return process.env.NODE_ENV === "development"
		? "http://localhost:3001/api" + path
		: "https://api.viewthis.app/api" + path
}
