import React from "react"

const DynamicLogo = ({ width = 50 }) => {
	// Note: Tailwind's fill classes apply directly to SVG elements.
	// Make sure `foreground` is a defined color in your Tailwind config.
	return (
		<svg
			width={width}
			height={width}
			viewBox="0 0 501 501"
			className="fill-foreground stroke-foreground" // Using Tailwind classes for fill and stroke
			xmlns="http://www.w3.org/2000/svg"
		>
			<path
				d="M176.5 250.5L250.5 177L325 250.5H176.5Z"
				className="fill-current" // Tailwind class for using the current color
			/>
			<rect
				x="225.5"
				y="239.5"
				width="49"
				height="54"
				className="fill-current"
			/>
			<path
				d="M177 410.5C172.5 410.5 16 250.5 16 250.5H81L251 420.5L420.5 250.5H484.5L325 410.5L251 485.5L177 410.5Z"
				className="fill-current"
			/>
			<rect
				x="51.5"
				y="142.5"
				width="397"
				height="24"
				className="fill-current"
			/>
			<rect
				x="96.5"
				y="110.5"
				width="307"
				height="19"
				className="fill-current"
			/>
			<rect
				x="186.5"
				y="61.5"
				width="127"
				height="9"
				className="fill-current"
			/>
			<rect
				x="231.5"
				y="44.5"
				width="37"
				height="4"
				className="fill-current"
			/>
			<rect
				x="141.5"
				y="83.5"
				width="217"
				height="14"
				className="fill-current"
			/>
		</svg>
	)
}

export default DynamicLogo
