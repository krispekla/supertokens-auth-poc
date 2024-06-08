export const appInfo = {
  appName: process.env.APP_NAME || "test",
  apiDomain: process.env.API_DOMAIN || "http://localhost:3000",
  websiteDomain: process.env.WEBSITE_DOMAIN || "http://localhost:3000",
  apiBasePath: process.env.API_BASE_PATH || "/auth",
  websiteBasePath: process.env.WEBSITE_BASE_PATH || "/auth",
};
