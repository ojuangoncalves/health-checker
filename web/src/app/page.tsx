"use client";

import useSWR from "swr";
import { API_ROUTES, type Site } from "../../lib/api";
import SiteCard from "@/components/SiteCard";

const fetcher = (url: string) => fetch(url).then((res) => res.json());

export default function Page() {
  const { data, error, isLoading } = useSWR<Site[]>(API_ROUTES.SITES, fetcher, {
    refreshInterval: 30000,
  });

  if (error) {
    return <main className="p-6">Erro ao carregar sites.</main>;
  }

  if (isLoading || !data) {
    return <main className="p-6">Carregando...</main>;
  }

  return (
    <main className="p-6">
      
      <h1 className="text-2xl font-bold mb-4">Monitor de Sites</h1>
      
      <section className="grid gap-4 sm:grid-cols-2 lg:grid-cols-3">
        {data.map((site) => (
          <SiteCard key={site.id} site={site} />
        ))}
      </section>
    </main>
  );
}
