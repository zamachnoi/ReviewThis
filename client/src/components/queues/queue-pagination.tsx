"use client"
import {
	Pagination,
	PaginationContent,
	PaginationEllipsis,
	PaginationItem,
	PaginationLink,
	PaginationNext,
	PaginationPrevious,
} from "@/components/ui/pagination"

interface MyPaginationProps {
	totalPages: number
	currentPage: number
	limit: number
	baseUrl: string
	search: string
}

export function MyPagination({
	totalPages,
	currentPage,
	limit,
	baseUrl,
	search,
}: MyPaginationProps) {
	const getPageUrl = (page: number) => {
		let url = `${baseUrl}?page=${page}&limit=${limit}`
		if (search) {
			url += `&search=${search}`
		}
		return url
	}
	const renderPaginationItems = () => {
		const paginationItems = []

		if (currentPage > 1) {
			paginationItems.push(
				<PaginationItem key={currentPage - 1}>
					<PaginationLink href={getPageUrl(currentPage - 1)}>
						{currentPage - 1}
					</PaginationLink>
				</PaginationItem>
			)
		}

		paginationItems.push(
			<PaginationItem key={currentPage}>
				<PaginationLink href={getPageUrl(currentPage)} isActive>
					{currentPage}
				</PaginationLink>
			</PaginationItem>
		)

		if (currentPage < totalPages) {
			paginationItems.push(
				<PaginationItem key={currentPage + 1}>
					<PaginationLink href={getPageUrl(currentPage + 1)}>
						{currentPage + 1}
					</PaginationLink>
				</PaginationItem>
			)
		}

		if (currentPage < totalPages - 1) {
			paginationItems.push(
				<PaginationItem key="ellipsis">
					<PaginationEllipsis />
				</PaginationItem>
			)
		}

		return paginationItems
	}

	return (
		<Pagination className="w-max">
			<PaginationContent>
				{currentPage > 1 && (
					<PaginationItem>
						<PaginationPrevious
							href={getPageUrl(currentPage - 1)}
						/>
					</PaginationItem>
				)}
				{renderPaginationItems()}
				{currentPage < totalPages && (
					<PaginationItem>
						<PaginationNext href={getPageUrl(currentPage + 1)} />
					</PaginationItem>
				)}
			</PaginationContent>
		</Pagination>
	)
}
