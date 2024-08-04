export const appInfo = {
  appName: process.env.NEXT_PUBLIC_APP_NAME || "test",
  apiDomain: process.env.NEXT_PUBLIC_API_DOMAIN || "http://localhost:3003",
  websiteDomain: process.env.WEBSITE_DOMAIN || "http://localhost:8084",
  apiBasePath: process.env.API_BASE_PATH || "/auth",
  websiteBasePath: process.env.WEBSITE_BASE_PATH || "/auth",
};
