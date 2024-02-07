import { Button } from "@/components/ui/button"
import {
	Dialog,
	DialogContent,
	DialogDescription,
	DialogFooter,
	DialogHeader,
	DialogTitle,
	DialogTrigger,
} from "@/components/ui/dialog"

import { CreateQueueForm } from "./create-queue-form"

export function CreateQueueDialog() {
	return (
		<Dialog>
			<DialogTrigger asChild>
				<Button variant="outline">Create</Button>
			</DialogTrigger>
			<DialogContent className="sm:max-w-[425px]">
				<DialogHeader>
					<DialogTitle>Create a queue</DialogTitle>
					<DialogDescription>
						Add a table name and a read-only API key from AirTable
					</DialogDescription>
				</DialogHeader>
				<div className="grid gap-4 py-4">
					<div className="grid items-center gap-4">
						<CreateQueueForm />
					</div>
				</div>
			</DialogContent>
		</Dialog>
	)
}
