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

      <SummaryCard summary={result.summary} />

      <RiskList risks={result.risks} />
    </div>
  );
}
