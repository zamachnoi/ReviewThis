export function getApiUrl(path: string) {
	return process.env.NODE_ENV === "development"
		? "https://127.0.0.1:3000/api/redirect?path=" + path
		: "https://api.viewthis.app/api/" + path
}
