"use client"

import { zodResolver } from "@hookform/resolvers/zod"
import { useForm } from "react-hook-form"
import { z } from "zod"
import { Checkbox } from "@/components/ui/checkbox"

import { Button } from "@/components/ui/button"
import {
	Form,
	FormControl,
	FormDescription,
	FormField,
	FormItem,
	FormLabel,
	FormMessage,
} from "@/components/ui/form"
import {
	Select,
	SelectContent,
	SelectItem,
	SelectTrigger,
	SelectValue,
} from "@/components/ui/select"
import { MdiSoundcloud } from "../icons/soundcloud"
import { Input } from "@/components/ui/input"
import { DialogClose } from "@/components/ui/dialog"
import { getApiUrl } from "@/app/utils"
import { Textarea } from "@/components/ui/textarea"

enum QueueType {
	SoundCloud = "soundcloud",
}

const createQueueSchema = z.object({
	name: z.string().min(1, "Must have name").max(100, "Name is too long"),
	description: z.string().max(1000, "Description is too long"),
	type: z.string(),
	private: z.boolean(),
})

export function CreateQueueForm() {
	// 1. Define your form.
	const form = useForm<z.infer<typeof createQueueSchema>>({
		resolver: zodResolver(createQueueSchema),
		defaultValues: {
			name: "",
			description: "",
			type: QueueType.SoundCloud,
			private: false,
		},
		mode: "all",
	})

	function onSubmit(data: z.infer<typeof createQueueSchema>) {
		const url = getApiUrl("/queues")
		try {
			fetch(url, {
				method: "POST",
				headers: {
					"Content-Type": "application/json",
				},
				body: JSON.stringify(data),
				credentials: "include",
			})
		} catch (error) {
			console.error("Failed to create queue", error)
		}
	}

	return (
		<Form {...form}>
			<form
				onSubmit={form.handleSubmit(onSubmit)}
				className="w-full space-y-6"
			>
				<FormField
					control={form.control}
					name="name"
					render={({ field }) => (
						<FormItem>
							<FormLabel>Name</FormLabel>
							<FormControl>
								<Input
									placeholder="Ninja's Fortnite Clips"
									{...field}
								/>
							</FormControl>
							<FormDescription>
								This is the name of your queue.
							</FormDescription>
							<FormMessage />
						</FormItem>
					)}
				/>
				<FormField
					control={form.control}
					name="type"
					render={({ field }) => (
						<FormItem>
							<FormLabel>Type</FormLabel>
							<Select
								onValueChange={field.onChange}
								defaultValue={field.value}
							>
								<FormControl>
									<SelectTrigger>
										<SelectValue placeholder="Select a type of queue to create!" />
									</SelectTrigger>
								</FormControl>
								<SelectContent>
									<SelectItem value="soundcloud">
										<div className="flex items-center gap-2">
											<MdiSoundcloud />
											SoundCloud
										</div>
									</SelectItem>
								</SelectContent>
							</Select>
							<FormDescription>
								This is the types of links users can add to your
								queue.
							</FormDescription>
							<FormMessage />
						</FormItem>
					)}
				/>
				<FormField
					name="private"
					control={form.control}
					render={({ field }) => (
						<FormItem>
							<FormLabel>Private</FormLabel>
							<div className="flex flex-row items-center gap-4">
								<FormControl>
									<Checkbox
										checked={field.value}
										onCheckedChange={(checked) =>
											field.onChange(checked)
										}
									/>
								</FormControl>
								<FormDescription>
									This queue will be private and only visible
									to you.
								</FormDescription>
							</div>
							<FormMessage />
						</FormItem>
					)}
				/>
				<FormField
					control={form.control}
					name="description"
					render={({ field }) => (
						<FormItem>
							<FormLabel>Description</FormLabel>
							<FormControl>
								<Textarea
									placeholder="Reviewing Fortnite clips stream."
									className="resize-none"
									{...field}
								/>
							</FormControl>
							<FormDescription></FormDescription>
							<FormMessage />
						</FormItem>
					)}
				/>

				<div className="flex flex-row justify-end">
					<DialogClose asChild>
						<Button
							disabled={!form.formState.isValid}
							type="submit"
						>
							Create Queue
						</Button>
					</DialogClose>
				</div>
			</form>
		</Form>
	)
}
