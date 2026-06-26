import { Turnstile } from "@marsidev/react-turnstile";

import { useAppContext } from "../hooks/useAppContext";

type TurnstileCaptchaProps = {
  onSuccess: (token: string) => void;
  onExpire: () => void;
  onError: () => void;
};

export function TurnstileCaptcha({
  onSuccess,
  onExpire,
  onError,
}: TurnstileCaptchaProps) {
  const { theme } = useAppContext();

  const siteKey = import.meta.env.VITE_TURNSTILE_SITE_KEY;

  if (!siteKey) {
    return (
      <p className="text-sm text-red-500">
        Turnstile site key não configurada.
      </p>
    );
  }

  return (
    <div className="mt-4">
      <Turnstile
        siteKey={siteKey}
        options={{
          theme,
        }}
        onSuccess={onSuccess}
        onExpire={onExpire}
        onError={onError}
      />
    </div>
  );
}
