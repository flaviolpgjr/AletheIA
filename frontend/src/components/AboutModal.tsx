import { Modal } from "./Modal";

type AboutModalProps = {
  onClose: () => void;
};

export function AboutModal({ onClose }: AboutModalProps) {
  return (
    <Modal title="Sobre o AletheIA" onClose={onClose}>
      <div className="space-y-4">
        <p>
          O <strong>AletheIA</strong> é uma plataforma experimental que utiliza
          Inteligência Artificial e dados públicos para analisar promessas
          públicas de forma transparente e explicável.
        </p>

        <p>
          O objetivo é ajudar cidadãos, jornalistas, pesquisadores e
          desenvolvedores a entenderem se uma promessa é clara, mensurável e
          verificável com base em evidências públicas.
        </p>

        <p>
          O AletheIA não prevê o futuro e não afirma se uma promessa será
          cumprida. Ele organiza informações, critérios, riscos e evidências
          para apoiar uma análise mais responsável.
        </p>

        <div>
          <h3 className="mb-2 font-semibold">Princípios do projeto</h3>

          <ul className="list-disc space-y-1 pl-5">
            <li>Não inventar dados.</li>
            <li>Mostrar fontes públicas utilizadas.</li>
            <li>Explicar os critérios avaliados.</li>
            <li>Indicar riscos, dependências e limitações.</li>
            <li>Separar evidência pública de opinião gerada por IA.</li>
          </ul>
        </div>
      </div>
    </Modal>
  );
}
