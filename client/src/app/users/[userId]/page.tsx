import { getServerJwt } from "@/lib/serverUtils"
import { jwtDecode } from "jwt-decode"
import { User } from "@/lib/types"
import AddBot from "@/components/user/add-bot"

export default function Page({ params }: { params: { userId: string } }) {
	const jwt = getServerJwt()
	let user: User | null = null
	if (jwt) {
		user = jwtDecode<User>(jwt)
	}

	return (
		<div>
			<h1>User {params.userId}</h1>
			{user && user.premium && user.db_id.toString() == params.userId && (
				<AddBot />
			)}
		</div>
	)
}
