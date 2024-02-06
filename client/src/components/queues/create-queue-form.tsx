"use client"
import { zodResolver } from "@hookform/resolvers/zod"
import { useForm } from "react-hook-form"
import { z } from "zod"

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
import { Input } from "@/components/ui/input"

enum QueueType {
	SoundCloud = "soundcloud",
}

const createQueueSchema = z.object({
	name: z.string().min(1, "Must have name").max(100, "Name is too long"),
	description: z
		.string()
		.min(1, "Must have description")
		.max(1000, "Description is too long"),
	type: z.nativeEnum(QueueType),
	private: z.boolean(),
})

export function ProfileForm() {
	// 1. Define your form.
	const form = useForm<z.infer<typeof createQueueSchema>>({
		resolver: zodResolver(createQueueSchema),
		defaultValues: {
			name: "",
			description: "",
			type: QueueType.SoundCloud,
			private: false,
		},
	})

	function onSubmit(data: z.infer<typeof createQueueSchema>) {
		console.log(data)
	}
}
