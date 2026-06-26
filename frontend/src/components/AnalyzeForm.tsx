import { useState } from "react";
import type { SyntheticEvent } from "react";

import { useAnalyzePromise } from "../hooks/useAnalyzePromise";
import { AnalysisResult } from "./AnalysisResult";
import { TurnstileCaptcha } from "./TurnstileCaptcha";

export function AnalyzeForm() {
  const [text, setText] = useState("");
  const [captchaToken, setCaptchaToken] = useState<string | null>(null);

  const { analyze, loading, error, result } = useAnalyzePromise();

  async function handleSubmit(event: SyntheticEvent<HTMLFormElement>) {
    event.preventDefault();

    if (!captchaToken) {
      return;
    }

    await analyze({
      text,
      captcha_token: captchaToken,
    });
  }

  const isSubmitDisabled = loading || text.trim().length === 0 || !captchaToken;

  return (
    <form
      onSubmit={handleSubmit}
      className="rounded-2xl border border-zinc-200 bg-white p-6 shadow-2xl shadow-zinc-200/60 transition-colors dark:border-zinc-800 dark:bg-zinc-900/70 dark:shadow-cyan-950/20"
    >
      <label className="mb-3 block text-sm font-medium text-zinc-700 dark:text-zinc-200">
        Promessa ou proposta pública
      </label>

      <textarea
        value={text}
        onChange={(event) => setText(event.currentTarget.value)}
        placeholder="Ex: Vamos zerar a fila de espera por consultas e exames no SUS em até 1 ano..."
        className="min-h-40 w-full resize-none rounded-xl border border-zinc-300 bg-zinc-50 p-4 text-zinc-900 outline-none transition placeholder:text-zinc-400 focus:border-cyan-400 focus:ring-2 focus:ring-cyan-400/20 dark:border-zinc-700 dark:bg-zinc-950 dark:text-zinc-100 dark:placeholder:text-zinc-500"
      />

      <div className="mt-5 flex flex-col gap-4 sm:flex-row sm:items-end sm:justify-between">
        <div className="space-y-3">
          <TurnstileCaptcha
            onSuccess={(token) => setCaptchaToken(token)}
            onExpire={() => setCaptchaToken(null)}
            onError={() => setCaptchaToken(null)}
          />

          <span className="block text-xs text-zinc-500 dark:text-zinc-400">
            {text.length}/1000 caracteres
          </span>
        </div>

        <button
          type="submit"
          disabled={isSubmitDisabled}
          className="rounded-xl bg-gradient-to-r from-cyan-500 to-violet-600 px-5 py-3 text-sm font-semibold text-white transition hover:opacity-90 disabled:cursor-not-allowed disabled:opacity-50"
        >
          {loading ? "Analisando..." : "Analisar promessa"}
        </button>
      </div>

      {error && (
        <p className="mt-4 rounded-xl border border-red-500/30 bg-red-500/10 p-3 text-sm text-red-300">
          {error}
        </p>
      )}

      {result && <AnalysisResult result={result} />}
    </form>
  );
}
