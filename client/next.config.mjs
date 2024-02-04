/** @type {import('next').NextConfig} */
const nextConfig = {
	async rewrites() {
		return [
			{
				source: "/redirect/:path*",
				destination:
					"/api/redirect?path=https://api.viewthis.app/api/:path*", // Pass destination path as query parameter
			},
		]
	},
}

export default nextConfig
