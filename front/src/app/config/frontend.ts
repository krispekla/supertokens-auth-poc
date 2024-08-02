import { useRouter } from "next/navigation";
import { SuperTokensConfig } from "supertokens-auth-react/lib/build/types";
import EmailPasswordReact from "supertokens-auth-react/recipe/emailpassword";
import EmailVerification from "supertokens-auth-react/recipe/emailverification";
import SessionReact from "supertokens-auth-react/recipe/session";
import ThirdPartyReact from "supertokens-auth-react/recipe/thirdparty";

import { appInfo } from "./appInfo";

const routerInfo: { router?: ReturnType<typeof useRouter>; pathName?: string } =
  {};

export function setRouter(
  router: ReturnType<typeof useRouter>,
  pathName: string
) {
  routerInfo.router = router;
  routerInfo.pathName = pathName;
}

export const frontendConfig = (): SuperTokensConfig => {
  return {
    appInfo,
    recipeList: [
      ThirdPartyReact.init({
        signInAndUpFeature: {
          providers: [
            ThirdPartyReact.Google.init(),
            ThirdPartyReact.Facebook.init(),
            ThirdPartyReact.Github.init(),
            ThirdPartyReact.Apple.init(),
          ],
        },
        style: `
            [data-supertokens~="superTokensBranding"] {
              display: none !important;
          }

          [data-supertokens~="formRow"]:last-child {
              [data-supertokens~="button"] {
                  height: 41px;
              }
          }
          
          [data-supertokens~="input"], [data-supertokens~="inputWrapper"] {
            height: 38px;
          }

          [data-supertokens~="button"][data-supertokens~="providerButton"] {
          min-height: 38px;
          }
            
           [data-supertokens~=container] {
            --palette-background: 51, 51, 51 !important;
            --palette-inputBackground: 41, 41, 41 !important;
            --palette-inputBorder: 41, 41, 41 !important;
            --palette-textTitle: 255, 255, 255 !important;
            --palette-textLabel: 255, 255, 255 !important;
            --palette-textPrimary: 255, 255, 255 !important;
            --palette-primary: 51, 112, 255 !important;
            --palette-primaryBorder: 4, 51, 161 !important;
            --palette-error: 173, 46, 46 !important;
            --palette-textInput: 169, 169, 169 !important;
            --palette-textLink: 169, 169, 169 !important;
            --font-size-1: 16px !important;
        }
        `,
      }),
      EmailPasswordReact.init(),
      EmailVerification.init({
        mode: "REQUIRED", // or "OPTIONAL"
      }),
      SessionReact.init(),
    ],
    windowHandler: (original) => ({
      ...original,
      location: {
        ...original.location,
        getPathName: () => routerInfo.pathName!,
        assign: (url) => routerInfo.router!.push(url.toString()),
        setHref: (url) => routerInfo.router!.push(url.toString()),
      },
    }),
  };
};
