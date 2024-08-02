"use client";
import { usePathname, useRouter } from "next/navigation";
import React from "react";
import SuperTokensReact, { SuperTokensWrapper } from "supertokens-auth-react";
import { EmailVerificationPreBuiltUI } from "supertokens-auth-react/recipe/emailverification/prebuiltui";
import { canHandleRoute, getRoutingComponent } from "supertokens-auth-react/ui";
import { frontendConfig, setRouter } from "../config/frontend";

if (typeof window !== "undefined") {
  // we only want to call this init function on the frontend, so we check typeof window !== 'undefined'
  SuperTokensReact.init(frontendConfig());
}

export const SuperTokensProvider: React.FC<React.PropsWithChildren<{}>> = ({
  children,
}) => {
  setRouter(useRouter(), usePathname() || window.location.pathname);

  if (canHandleRoute([/* Other pre built UI */ EmailVerificationPreBuiltUI])) {
    return getRoutingComponent([
      /* Other pre built UI */ EmailVerificationPreBuiltUI,
    ]);
  }

  return <SuperTokensWrapper>{children}</SuperTokensWrapper>;
};
