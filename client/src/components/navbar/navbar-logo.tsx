"use client"
import Link from "next/link"
import DynamicLogo from "./dynamiclogo"
import { useEffect, useState } from "react"

export default function NavbarLogo() {
	const [color, setColor] = useState("black") // Default color assuming light mode

	useEffect(() => {
		// Function to update SVG fill color based on the .dark class on the html element
		const updateColorForTheme = () => {
			const isDarkMode =
				document.documentElement.classList.contains("dark")
			setColor(isDarkMode ? "white" : "black")
		}

		updateColorForTheme() // Check the theme on initial mount

		// Set up a MutationObserver to detect changes in class attribute on the html element
		const observer = new MutationObserver((mutations) => {
			mutations.forEach((mutation) => {
				if (
					mutation.type === "attributes" &&
					mutation.attributeName === "class"
				) {
					updateColorForTheme() // Update color if class attribute changes
				}
			})
		})

		observer.observe(document.documentElement, {
			attributes: true, // Watch for attribute changes
			attributeFilter: ["class"], // Specifically for changes in the class attribute
		})

		// Clean up the observer to prevent memory leaks
		return () => observer.disconnect()
	}, [])
	return (
		<div>
			<Link
				className="flex flex-row justify-center items-center gap-2"
				href="/"
			>
				<DynamicLogo fill={color} />
				<h1 className="text-2xl font-semibold">viewthis</h1>
			</Link>
		</div>
	)
}
