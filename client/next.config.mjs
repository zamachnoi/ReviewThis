/** @type {import('next').NextConfig} */
const nextConfig = {
	async rewrites() {
		return [
			{
				source: "/api/:path*",
				destination: "https://api.viewthis.app/:path*", // Proxy to your API
			},
		]
	},
}

export default nextConfig
