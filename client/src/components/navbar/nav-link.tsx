import Link from "next/link"

export default function NavLink({
	children,
	href,
}: {
	children: React.ReactNode
	href: string
}) {
	return (
		<Link
			className="text-xl hover:text-accent text-accent-foreground transition-colors ease-in-out duration-75"
			href={href}
		>
			{children}
		</Link>
	)
}
