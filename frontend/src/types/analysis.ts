export type AnalyzePromiseRequest = {
  text: string
}

export type AnalyzePromiseResponse = {
  summary: string
  score: number
  risks: string[]
}