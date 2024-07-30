import axios from "axios";
import { appInfo } from "../config/appInfo";

// Deprecated
// Our SDK adds interceptors to fetch and XHR (used by axios) to save and
// add session tokens from and to the request.
// https://supertokens.com/docs/thirdpartyemailpassword/custom-ui/handling-session-tokens#for-web
// SuperTokens.addAxiosInterceptors(axios);

axios.defaults.baseURL = appInfo.apiDomain;
export default axios;
