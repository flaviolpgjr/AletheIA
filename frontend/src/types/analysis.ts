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

export type PublicSource = {
  name: string;
  description: string;
};

export type Evidence = {
  source: string;
  title: string;
  description: string;
  url: string;
  indicator: string;
  value: number;
  unit: string;
  reference: string;
};

export type AnalyzePromiseResponse = {
  summary: string;
  score: number;
  confidence: number;
  target_value: number;
  target_unit: string;
  criteria: Criterion[];
  risks: string[];
  sources: PublicSource[];
  evidence: Evidence[];
};
