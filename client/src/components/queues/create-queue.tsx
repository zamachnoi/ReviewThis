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
		<div className="w-1/12 p-4">
			<Dialog>
				<DialogTrigger asChild>
					<Button variant="outline">Create</Button>
				</DialogTrigger>
				<DialogContent className="sm:max-w-[425px]">
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
