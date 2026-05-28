import { apiFetch } from "./api";

import type {
  AnalyzePromiseRequest,
  AnalyzePromiseResponse,
} from "../types/analysis";

export function analyzePromise(
  data: AnalyzePromiseRequest
): Promise<AnalyzePromiseResponse> {
  return apiFetch<AnalyzePromiseResponse>("/promises/analyze", {
    method: "POST",
    body: JSON.stringify(data),
  });
}