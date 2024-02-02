import Link from "next/link"

type NavItemProps = {
	text: string
	href: string
	active: boolean
}
const NavItem = ({ text, href, active }: NavItemProps) => {
	return (
		<Link href={href}>
			<a className={`nav__item ${active ? "active" : ""}`}>{text}</a>
		</Link>
	)
}

export default NavItem
