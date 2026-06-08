import { apiFetch } from "./api";

import type {
  AnalyzePromiseRequest,
  AnalyzePromiseResponse,
} from "../types/analysis";

export async function analyzePromise(
  payload: AnalyzePromiseRequest,
): Promise<AnalyzePromiseResponse> {
  return apiFetch<AnalyzePromiseResponse>("/promises/analyze", {
    method: "POST",
    body: JSON.stringify(payload),
  });
}
