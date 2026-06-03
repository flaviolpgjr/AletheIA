import type { AnalyzePromiseResponse } from "../types/analysis";
import { RiskList } from "./RiskList";
import { ScoreCard } from "./ScoreCard";
import { SummaryCard } from "./SummaryCard";

type AnalysisResultProps = {
  result: AnalyzePromiseResponse;
};

export function AnalysisResult({ result }: AnalysisResultProps) {
  return (
    <div className="mt-6 space-y-4">
      <ScoreCard score={result.score} />

      <div className="rounded-2xl border border-zinc-200 bg-white p-5 transition-colors dark:border-zinc-800 dark:bg-zinc-900">
        <h3 className="text-sm font-semibold text-zinc-900 dark:text-white">
          Confiança da análise
        </h3>

        <p className="mt-2 text-2xl font-bold text-cyan-600 dark:text-cyan-400">
          {result.confidence}%
        </p>
      </div>

      <SummaryCard summary={result.summary} />

      <div className="rounded-2xl border border-zinc-200 bg-white p-5 transition-colors dark:border-zinc-800 dark:bg-zinc-900">
        <h3 className="text-sm font-semibold text-zinc-900 dark:text-white">
          Critérios avaliados
        </h3>

        <div className="mt-4 space-y-3">
          {result.criteria.map((criterion) => (
            <div
              key={criterion.key}
              className="rounded-xl border border-zinc-200 bg-zinc-50 p-4 transition-colors dark:border-zinc-800 dark:bg-zinc-950"
            >
              <div className="flex items-start justify-between gap-4">
                <div>
                  <p className="font-medium text-zinc-900 dark:text-white">
                    {criterion.name}
                  </p>

                  <div className="mt-2">
                    <span className={getStatusClassName(criterion.status)}>
                      {translateStatus(criterion.status)}
                    </span>
                  </div>
                </div>

                <p className="shrink-0 text-sm font-medium text-zinc-700 dark:text-zinc-300">
                  {criterion.score}/{criterion.weight}
                </p>
              </div>

              <p className="mt-3 text-sm text-zinc-600 dark:text-zinc-300">
                {criterion.explanation}
              </p>
            </div>
          ))}
        </div>
      </div>

      <RiskList risks={result.risks} />
    </div>
  );
}

function translateStatus(status: string) {
  switch (status) {
    case "yes":
      return "Sim";

    case "partial":
      return "Parcial";

    case "no":
      return "Não";

    default:
      return status;
  }
}

function getStatusClassName(status: string) {
  const baseClassName = "rounded-full px-2 py-1 text-xs font-medium";

  switch (status) {
    case "yes":
      return `${baseClassName} bg-emerald-100 text-emerald-700 dark:bg-emerald-500/10 dark:text-emerald-400`;

    case "partial":
      return `${baseClassName} bg-yellow-100 text-yellow-700 dark:bg-yellow-500/10 dark:text-yellow-400`;

    case "no":
      return `${baseClassName} bg-red-100 text-red-700 dark:bg-red-500/10 dark:text-red-400`;

    default:
      return `${baseClassName} bg-zinc-100 text-zinc-700 dark:bg-zinc-800 dark:text-zinc-300`;
  }
}
