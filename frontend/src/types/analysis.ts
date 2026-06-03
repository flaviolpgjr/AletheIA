export type AnalyzePromiseRequest = {
  text: string;
};

export type Criterion = {
  key: string;
  name: string;
  weight: number;
  status: "yes" | "partial" | "no";
  score: number;
  explanation: string;
};

export type AnalyzePromiseResponse = {
  summary: string;
  score: number;
  confidence: number;
  criteria: Criterion[];
  risks: string[];
};
