import Link from "next/link"
import DynamicLogo from "../icons/dynamiclogo"

export default function NavbarLogo() {
	return (
		<div>
			<Link
				className="flex flex-row justify-center items-center gap-2"
				href="/"
			>
				<DynamicLogo color="white" width={50} />
				<h1 className="text-3xl font-semibold">
					view<span className="text-accent">this</span>
				</h1>
			</Link>
		</div>
	)
}
