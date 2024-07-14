import { useEffect, useState } from 'react'
import { createRoot } from 'react-dom/client'
import moment from 'moment'

import './styles.css'

interface IPost {
	id: number
	title: string
	content: string
	pub_time: number
	link: string
}

const fetcher = <T,>(url: string, signal: AbortSignal) => {
	return fetch(url, { signal }).then((res) => res.json()) as Promise<T>
}

function Aggregator() {
	const [news, setNews] = useState<IPost[]>([])

	useEffect(() => {
		const controller = new AbortController()
		fetcher<IPost[]>(location.origin.replace(/:[0-9]+/, '') + '/news/10', controller.signal).then((news) => {
			setNews(news)
		})
		return () => {
			controller.abort()
		}
	}, [])

	return (
		<div className="flex flex-col h-full p-4 gap-6">
			<h1 className="text-6xl max-xl:text-5xl text-center text-[#b5e916] m-2">Новостной агрегатор</h1>
			<div className="grid grid-cols-4 max-sm:grid-cols-1 max-lg:grid-cols-2 max-xl:grid-cols-3 gap-6">
				{news.map((post) => (
					<div
						key={post.id}
						className="flex flex-col bg-slate-100 hover:bg-white border-4 border-[#328818] hover:border-white transition-all rounded-xl p-4 gap-4"
					>
						<div className="flex flex-col flex-1 gap-4">
							<div className="flex justify-between">
								<a className="text-lg text-blue-600 underline" href={post.link}>
									{post.title}
								</a>
								<img className="w-6 h-6 ml-2" src="/img/rss.png" alt={post.title} />
							</div>
							<p className="break-words top line-clamp-6 text-gray-700">{post.content}</p>
						</div>
						<p className="break-words top text-right text-gray-500">{moment(post.pub_time * 1000).format('HH:mm DD-MM-YYYY')}</p>
					</div>
				))}
			</div>
		</div>
	)
}

const root = createRoot(document.getElementById('root'))
root.render(<Aggregator />)
