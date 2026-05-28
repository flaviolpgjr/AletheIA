import { useState } from "react";

import { analyzePromise } from "../services/analysisService";

import type {
  AnalyzePromiseRequest,
  AnalyzePromiseResponse,
} from "../types/analysis";

export function useAnalyzePromise() {
  const [result, setResult] = useState<AnalyzePromiseResponse | null>(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  async function analyze(data: AnalyzePromiseRequest) {
    try {
      setLoading(true);
      setError(null);

      const response = await analyzePromise(data);

      setResult(response);
    } catch {
      setError("Não foi possível analisar a promessa.");
    } finally {
      setLoading(false);
    }
  }

  return {
    result,
    loading,
    error,
    analyze,
  };
}