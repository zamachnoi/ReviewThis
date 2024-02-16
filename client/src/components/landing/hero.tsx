import { Button } from "../ui/button"
import Link from "next/link"

export default function Hero() {
	return (
		<div className="bg-background">
			<div
				className="absolute inset-x-0 overflow-hidden top-20 -z-9 transform-gpu blur-3xl"
				aria-hidden="true"
			>
				<div
					className="relative left-[calc(50%-11rem)] aspect-[1155/678] w-[36.125rem] -translate-x-1/2 rotate-[30deg] bg-gradient-to-tr from-[#ff80b5] to-[#9089fc] opacity-30 sm:left-[calc(50%-30rem)] sm:w-[72.1875rem]"
					style={{
						clipPath:
							"polygon(74.1% 44.1%, 100% 61.6%, 97.5% 26.9%, 85.5% 0.1%, 80.7% 2%, 72.5% 32.5%, 60.2% 62.4%, 52.4% 68.1%, 47.5% 58.3%, 45.2% 34.5%, 27.5% 76.7%, 0.1% 64.9%, 17.9% 100%, 27.6% 76.8%, 76.1% 97.7%, 74.1% 44.1%)",
					}}
				/>
			</div>
			<div className="relative px-6 isolate pt-14 lg:px-8">
				<div className="flex flex-row items-center justify-around sm:py-48">
					<div className="flex flex-col items-center max-w-2xl gap-4">
						<h1 className="font-bold tracking-tight text-7xl text-foreground">
							View your{" "}
							<span className="text-accent">queues</span> in one
							place
						</h1>
						<p className="text-lg leading-8 text-secondary-foreground">
							Blah blah blah placeholder text. Blah blah blah
							placeholder text. Blah blah blah placeholder text.
							Blah blah blah placeholder text.
						</p>
						<Button asChild>
							<Link href="/queues">Get Started</Link>
						</Button>
					</div>

					<div>QUEUE PLACEHOLDER</div>
				</div>

				<div
					className="absolute inset-x-0 top-[calc(100%-13rem)] -z-10 transform-gpu overflow-hidden blur-3xl sm:top-[calc(100%-30rem)]"
					aria-hidden="true"
				>
					<div
						className="relative right-[calc(50%+3rem)] aspect-[1155/678] w-[36.125rem] -translate-x-1/2 bg-gradient-to-tr from-[#a299f0] to-[#426bfe] opacity-30 sm:left-[calc(50%+36rem)] sm:w-[72.1875rem]"
						style={{
							clipPath:
								"polygon(74.1% 44.1%, 100% 61.6%, 97.5% 26.9%, 85.5% 0.1%, 80.7% 2%, 72.5% 32.5%, 60.2% 62.4%, 52.4% 68.1%, 47.5% 58.3%, 45.2% 34.5%, 27.5% 76.7%, 0.1% 64.9%, 17.9% 100%, 27.6% 76.8%, 76.1% 97.7%, 74.1% 44.1%)",
						}}
					/>
				</div>
			</div>
		</div>
	)
}
