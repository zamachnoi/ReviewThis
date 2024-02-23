import { Button } from "@/components/ui/button"
import {
	Dialog,
	DialogContent,
	DialogDescription,
	DialogHeader,
	DialogTitle,
	DialogTrigger,
} from "@/components/ui/dialog"

import { CreateQueueForm } from "./create-queue-form"

export function CreateQueueDialog() {
	return (
		<div className="p-4">
			<Dialog>
				<DialogTrigger asChild>
					<Button className="w-40 h-12 text-lg" variant="secondary">
						Create
					</Button>
				</DialogTrigger>
				<DialogContent className="sm:max-w-[800px]">
					<DialogHeader>
						<DialogTitle>Create a queue</DialogTitle>
						<DialogDescription>
							Add a fun little discription
						</DialogDescription>
					</DialogHeader>
					<div className="grid">
						<div className="items-center grid gap-4">
							<CreateQueueForm />
						</div>
					</div>
				</DialogContent>
			</Dialog>
		</div>
	)
}
