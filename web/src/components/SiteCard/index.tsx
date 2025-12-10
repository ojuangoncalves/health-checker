import type { Site } from "../../../lib/api";

function statusColor(status: number) {
  if (status === 200) return 'bg-green-500'
  if (status === 0) return 'bg-gray-500'
  if (status >= 400) return 'bg-red-500'
  return 'bg-yellow-500'
}

export default function SiteCard({ site }: { site: Site }) {
  return (
    <div className="rounded-lg border border-slate-700 bg-slate-900 p-4 flex flex-col gap-2">
      <div className="flex items-center justify-between">
        <h2 className="font-semibold text-slate-50">{ site.nome }</h2>
        <span className={`px-2 py-2 text-xs rounded text-white ${statusColor(site.status)}`}>
          { site.status === 0 ? "Indisponível" : site.status }
        </span>
      </div>
      <p className="text-sm text-slate-300">{ site.url }</p>
    </div>
  )
}
