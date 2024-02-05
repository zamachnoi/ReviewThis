export function getApiUrl(path: string) {
	return process.env.NODE_ENV === "development"
		? "http://127.0.0.1:3001/api/" + path
		: "https://api.viewthis.app/api/" + path
}
