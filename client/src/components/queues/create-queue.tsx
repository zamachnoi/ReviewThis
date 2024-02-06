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
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import {
	Select,
	SelectContent,
	SelectItem,
	SelectTrigger,
	SelectValue,
} from "@/components/ui/select"
import { MdiSoundcloud } from "../icons/soundcloud"

export function CreateQueue() {
	return (
		<Dialog>
			<DialogTrigger asChild>
				<Button variant="outline">Create Queue</Button>
			</DialogTrigger>
			<DialogContent className="sm:max-w-[425px]">
				<DialogHeader>
					<DialogTitle>Create a Queue</DialogTitle>
					<DialogDescription>
						Create a queue for people to submit to!
					</DialogDescription>
				</DialogHeader>
				<div className="grid gap-5 py-4">
					<div className="flex flex-row justify-between items-center space-y-1.5">
						<Label htmlFor="name" className="text-left">
							Name
						</Label>
						<Input
							id="name"
							defaultValue="Pedro Duarte"
							className="w-5/6"
						/>
					</div>
					<div className="grid gap-4 py-4">
						<div className="flex flex-row items-center justify-between space-y-1.5">
							<Label htmlFor="type">Type</Label>
							<div className="w-5/6">
								<Select>
									<SelectTrigger id="type">
										<SelectValue placeholder="Select" />
									</SelectTrigger>
									<SelectContent position="popper">
										<SelectItem value="soundcloud">
											<div className="flex flex-row items-center gap-2">
												<MdiSoundcloud />
												SoundCloud
											</div>
										</SelectItem>
									</SelectContent>
								</Select>
							</div>
						</div>
					</div>
				</div>
				<DialogFooter>
					<Button type="submit">Save changes</Button>
				</DialogFooter>
			</DialogContent>
		</Dialog>
	)
}
