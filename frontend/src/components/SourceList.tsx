import type { PublicSource } from "../types/analysis";

type SourceListProps = {
  sources: PublicSource[];
};

export function SourceList({ sources }: SourceListProps) {
  if (sources.length === 0) {
    return null;
  }

  return (
    <div className="rounded-2xl border border-zinc-200 bg-white p-5 transition-colors dark:border-zinc-800 dark:bg-zinc-900">
      <h3 className="text-sm font-semibold text-zinc-900 dark:text-white">
        Fontes públicas sugeridas
      </h3>

      <div className="mt-4 space-y-3">
        {sources.map((source) => (
          <div
            key={source.name}
            className="rounded-xl border border-zinc-200 bg-zinc-50 p-4 transition-colors dark:border-zinc-800 dark:bg-zinc-950"
          >
            <p className="font-medium text-zinc-900 dark:text-white">
              {source.name}
            </p>

            <p className="mt-2 text-sm text-zinc-600 dark:text-zinc-300">
              {source.description}
            </p>
          </div>
        ))}
      </div>
    </div>
  );
}
