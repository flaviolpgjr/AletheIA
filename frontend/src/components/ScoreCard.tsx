type ScoreCardProps = {
  score: number;
};

export function ScoreCard({ score }: ScoreCardProps) {
  if (score < 0) {
    return (
      <div className="rounded-2xl border border-zinc-200 bg-white p-5 text-center transition-colors dark:border-zinc-800 dark:bg-zinc-900">
        <h2 className="text-lg font-semibold text-zinc-900 dark:text-white">
          Score AletheIA
        </h2>

        <div className="mt-6 rounded-xl border border-zinc-200 bg-zinc-50 p-6 dark:border-zinc-800 dark:bg-zinc-950">
          <p className="text-5xl font-bold text-zinc-500">--</p>

          <p className="mt-2 text-sm text-zinc-500 dark:text-zinc-400">
            Não identificado
          </p>
        </div>

        <div className="mt-5 inline-flex rounded-full border border-zinc-300 bg-zinc-100 px-3 py-1 text-sm font-medium text-zinc-600 dark:border-zinc-700 dark:bg-zinc-800/60 dark:text-zinc-300">
          Texto insuficiente
        </div>
      </div>
    );
  }

  const status = getScoreStatus(score);
  const percentage = Math.min(Math.max(score, 0), 100);

  return (
    <div className="rounded-2xl border border-zinc-200 bg-white p-5 text-center transition-colors dark:border-zinc-800 dark:bg-zinc-900">
      <h2 className="text-lg font-semibold text-zinc-900 dark:text-white">
        Score AletheIA
      </h2>

      <p className="mt-1 text-sm text-zinc-500 dark:text-zinc-400">
        Sustentação verificável da promessa
      </p>

      <div className="relative mx-auto mt-6 h-28 w-56">
        <svg viewBox="0 0 200 100" className="h-full w-full">
          <path
            d="M 20 90 A 80 80 0 0 1 180 90"
            fill="none"
            stroke="currentColor"
            strokeWidth="16"
            strokeLinecap="round"
            className="text-zinc-200 dark:text-zinc-800"
          />

          <path
            d="M 20 90 A 80 80 0 0 1 180 90"
            fill="none"
            stroke="url(#scoreGradient)"
            strokeWidth="16"
            strokeLinecap="round"
            strokeDasharray="251"
            strokeDashoffset={251 - (251 * percentage) / 100}
          />

          <defs>
            <linearGradient id="scoreGradient" x1="0" x2="1" y1="0" y2="0">
              <stop offset="0%" stopColor="#ef4444" />
              <stop offset="50%" stopColor="#eab308" />
              <stop offset="100%" stopColor="#22c55e" />
            </linearGradient>
          </defs>
        </svg>

        <div className="absolute inset-x-0 bottom-0">
          <p className="text-5xl font-bold text-cyan-500 dark:text-cyan-400">
            {score}
          </p>

          <p className="text-xs text-zinc-500">/100</p>
        </div>
      </div>

      <div
        className={`mt-5 inline-flex rounded-full px-3 py-1 text-sm font-medium ${status.className}`}
      >
        {status.label}
      </div>
    </div>
  );
}

function getScoreStatus(score: number) {
  if (score >= 80) {
    return {
      label: "Alta sustentação",
      className:
        "border border-emerald-500/30 bg-emerald-500/10 text-emerald-600 dark:text-emerald-300",
    };
  }

  if (score >= 50) {
    return {
      label: "Sustentação moderada",
      className:
        "border border-yellow-500/30 bg-yellow-500/10 text-yellow-600 dark:text-yellow-300",
    };
  }

  return {
    label: "Baixa sustentação",
    className:
      "border border-red-500/30 bg-red-500/10 text-red-600 dark:text-red-300",
  };
}
