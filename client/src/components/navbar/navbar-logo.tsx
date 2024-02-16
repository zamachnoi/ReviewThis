import Link from "next/link"
import DynamicLogo from "../icons/dynamiclogo"
import Image from "next/image"

export default function NavbarLogo() {
	return (
		<div>
			<Link
				className="flex flex-row items-center justify-center gap-2"
				href="/"
			>
				<Image
					src={"/viewthis.png"}
					alt="viewthis"
					width={50}
					height={50}
				/>
				<h1 className="text-3xl font-semibold">
					view<span className="text-accent">this</span>
				</h1>
			</Link>
		</div>
	)
}
