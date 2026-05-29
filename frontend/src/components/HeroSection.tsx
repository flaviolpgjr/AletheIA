import logoIcon from "../assets/logo-icon.png";

export function HeroSection() {
  return (
    <section className="flex flex-col items-center text-center">
      <img
        src={logoIcon}
        alt="AletheIA"
        className="mb-8 w-40 drop-shadow-[0_0_35px_rgba(0,212,255,0.35)]"
      />

      <h1 className="mb-4 text-5xl font-bold tracking-tight text-zinc-950 dark:text-white">
        Alethe
        <span className="bg-gradient-to-r from-cyan-400 to-violet-500 bg-clip-text text-transparent">
          IA
        </span>
      </h1>

      <p className="mb-12 max-w-xl text-lg leading-8 text-zinc-600 dark:text-zinc-400">
        Avalie a viabilidade de promessas públicas utilizando IA, dados públicos
        e indicadores reais.
      </p>
    </section>
  );
}
