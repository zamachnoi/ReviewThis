import { getServerJwt } from "@/lib/serverUtils"
import { jwtDecode } from "jwt-decode"
import { User } from "@/lib/types"
import AddBot from "@/components/user/add-bot"
import { getApiUrl } from "@/lib/utils"

export default function Page({ params }: { params: { queueId: string } }) {
	const jwt = getServerJwt()
	let user: User | null = null
	if (jwt) {
		user = jwtDecode<User>(jwt)
	}
	let submissions = fetch(
		getApiUrl(`/queues/${params.queueId}/submissions?limit=5&page=1`)
	).then((res) => res.json())
	console.log(submissions)
	return <div>Hi</div>
}
