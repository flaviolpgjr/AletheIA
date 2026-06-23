import type { Evidence } from "../types/analysis";

type EvidenceCardProps = {
  evidence: Evidence;
  targetValue: number;
  targetUnit: string;
  relationPercentage: number | null;
};

export function EvidenceCard({
  evidence,
  targetValue,
  targetUnit,
  relationPercentage,
}: EvidenceCardProps) {
  return (
    <div className="rounded-2xl border border-zinc-200 bg-white p-5 transition-colors dark:border-zinc-800 dark:bg-zinc-900">
      <div className="flex items-start justify-between gap-4">
        <div>
          <h3 className="text-sm font-semibold text-zinc-900 dark:text-white">
            Evidência pública utilizada
          </h3>

          <p className="mt-2 text-sm text-zinc-600 dark:text-zinc-300">
            {evidence.description}
          </p>
        </div>

        <span className="rounded-full bg-zinc-100 px-2 py-1 text-xs font-medium text-zinc-700 dark:bg-zinc-800 dark:text-zinc-300">
          {evidence.source}
        </span>
      </div>

      <div className="mt-5 grid gap-3 sm:grid-cols-3">
        <EvidenceMetric
          label="Meta proposta"
          value={
            targetValue > 0
              ? `${formatNumber(targetValue)} ${targetUnit}`
              : "Não identificada"
          }
        />

        <EvidenceMetric
          label="Linha de base"
          value={`${formatNumber(evidence.value)} ${evidence.unit}`}
        />

        <EvidenceMetric
          label="Relação"
          value={
            relationPercentage !== null
              ? `${relationPercentage.toFixed(2)}%`
              : "Não calculada"
          }
        />
      </div>

      <div className="mt-4 rounded-xl border border-zinc-200 bg-zinc-50 p-4 text-sm text-zinc-700 transition-colors dark:border-zinc-800 dark:bg-zinc-950 dark:text-zinc-300">
        <p className="font-medium text-zinc-900 dark:text-white">
          {evidence.title}
        </p>

        <p className="mt-2">{evidence.reference}</p>

        {evidence.url && (
          <a
            href={evidence.url}
            target="_blank"
            rel="noreferrer"
            className="mt-3 inline-block text-sm font-medium text-zinc-900 hover:underline dark:text-white"
          >
            Ver fonte pública
          </a>
        )}
      </div>
    </div>
  );
}

type EvidenceMetricProps = {
  label: string;
  value: string;
};

function EvidenceMetric({ label, value }: EvidenceMetricProps) {
  return (
    <div className="rounded-xl border border-zinc-200 bg-zinc-50 p-4 transition-colors dark:border-zinc-800 dark:bg-zinc-950">
      <p className="text-xs font-medium uppercase tracking-wide text-zinc-500 dark:text-zinc-400">
        {label}
      </p>

      <p className="mt-2 text-lg font-bold text-zinc-900 dark:text-white">
        {value}
      </p>
    </div>
  );
}

function formatNumber(value: number) {
  return new Intl.NumberFormat("pt-BR", {
    maximumFractionDigits: 2,
  }).format(value);
}
