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

      <div className="rounded-lg border border-slate-700 bg-slate-900 p-4">
        <h3 className="text-sm font-semibold text-slate-200">
          Confiança da análise
        </h3>

        <p className="mt-2 text-2xl font-bold text-slate-100">
          {result.confidence}%
        </p>
      </div>

      <SummaryCard summary={result.summary} />

      <div className="rounded-lg border border-slate-700 bg-slate-900 p-4">
        <h3 className="text-sm font-semibold text-slate-200">
          Critérios avaliados
        </h3>

        <div className="mt-4 space-y-3">
          {result.criteria.map((criterion) => (
            <div
              key={criterion.key}
              className="rounded-md border border-slate-700 p-3"
            >
              <div className="flex items-center justify-between gap-4">
                <p className="font-medium text-slate-100">{criterion.name}</p>

                <p className="text-sm text-slate-300">
                  {criterion.score}/{criterion.weight}
                </p>
              </div>

              <div className="mt-2">
                <span
                  className={`rounded-full px-2 py-1 text-xs font-medium ${
                    criterion.status === "yes"
                      ? "bg-emerald-500/10 text-emerald-500"
                      : criterion.status === "partial"
                        ? "bg-yellow-500/10 text-yellow-500"
                        : "bg-red-500/10 text-red-500"
                  }`}
                >
                  {translateStatus(criterion.status)}
                </span>
              </div>

              <p className="mt-2 text-sm text-slate-300">
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
