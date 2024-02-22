"use client"
import ReactPlayer from "react-player/lazy"

type SoundcloudContentPreviewProps = {
	content: string
}
export default function SoundCloudContentPreview({
	content,
}: SoundcloudContentPreviewProps) {
	return (
		<ReactPlayer
			url={content}
			width={400}
			height={60}
			volume={0.25}
			style={{ overflow: "hidden" }}
			config={{
				soundcloud: {
					options: {
						show_user: true,
						single_active: true,
						visual: false,
						color: "19488A",
						theme: "dark",
						show_artwork: false,
					},
				},
			}}
		/>
	)
}
