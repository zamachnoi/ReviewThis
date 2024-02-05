import Link from "next/link"
import DynamicLogo from "./dynamiclogo"

export default function NavbarLogo() {
	return (
		<div>
			<Link
				className="flex flex-row justify-center items-center gap-2"
				href="/"
			>
				<DynamicLogo />
				<h1 className="text-2xl font-semibold">viewthis</h1>
			</Link>
		</div>
	)
}
