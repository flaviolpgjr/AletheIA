import type { AnalyzePromiseResponse } from "../types/analysis";
import { CriteriaList } from "./CriteriaList";
import { EvidenceCard } from "./EvidenceCard";
import { RiskList } from "./RiskList";
import { ScoreCard } from "./ScoreCard";
import { SourceList } from "./SourceList";
import { SummaryCard } from "./SummaryCard";

type AnalysisResultProps = {
  result: AnalyzePromiseResponse;
};

export function AnalysisResult({ result }: AnalysisResultProps) {
  const mainEvidence = result.evidence?.[0];

  const hasTarget = result.target_value > 0 && result.target_unit;

  const relationPercentage =
    mainEvidence && mainEvidence.value > 0 && hasTarget
      ? (result.target_value / mainEvidence.value) * 100
      : null;

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

      {mainEvidence && (
        <EvidenceCard
          evidence={mainEvidence}
          targetValue={result.target_value}
          targetUnit={result.target_unit}
          relationPercentage={relationPercentage}
        />
      )}

      <CriteriaList criteria={result.criteria} />

      <SourceList sources={result.sources} />

      <RiskList risks={result.risks} />
    </div>
  );
}
