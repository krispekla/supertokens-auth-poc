import bodyParser from "body-parser";
import cors from "cors";
import dotenv from "dotenv";
import express from "express";
import helmet from "helmet";
import morgan from "morgan";

import supertokens from "supertokens-node";
import { middleware } from "supertokens-node/framework/express";
import Dashboard from "supertokens-node/recipe/dashboard";
import Session from "supertokens-node/recipe/session";
import ThirdParty from "supertokens-node/recipe/thirdparty";

dotenv.config();

const app = express();
const port = process.env.PORT || 3008;
const supertokenCoreUrl = process.env.SUPERTOKEN_CORE_SVC_URL;

supertokens.init({
  framework: "express",
  supertokens: {
    // https://try.supertokens.com is for demo purposes. Replace this with the address of your core instance (sign up on supertokens.com), or self host a core.
    connectionURI: supertokenCoreUrl,
    // apiKey: <API_KEY(if configured)>,
  },
  appInfo: {
    // learn more about this on https://supertokens.com/docs/session/appinfo
    appName: "test",
    apiDomain: "http://localhost:3005",
    websiteDomain: "http://localhost:8084",
    apiBasePath: "/auth",
    websiteBasePath: "/auth",
  },
  recipeList: [
    Session.init(),
    ThirdParty.init({
      signInAndUpFeature: {
        // We have provided you with development keys which you can use for testing.
        // IMPORTANT: Please replace them with your own OAuth keys for production use.
        providers: [
          {
            config: {
              thirdPartyId: "google",
              clients: [
                {
                  clientId:
                    "1060725074195-kmeum4crr01uirfl2op9kd5acmi9jutn.apps.googleusercontent.com",
                  clientSecret: "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",
                },
              ],
            },
          },
          {
            config: {
              thirdPartyId: "apple",
              clients: [
                {
                  clientId: "4398792-io.supertokens.example.service",
                  additionalConfig: {
                    keyId: "7M48Y4RYDL",
                    privateKey:
                      "-----BEGIN PRIVATE KEY-----\nMIGTAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBHkwdwIBAQQgu8gXs+XYkqXD6Ala9Sf/iJXzhbwcoG5dMh1OonpdJUmgCgYIKoZIzj0DAQehRANCAASfrvlFbFCYqn3I2zeknYXLwtH30JuOKestDbSfZYxZNMqhF/OzdZFTV0zc5u5s3eN+oCWbnvl0hM+9IW0UlkdA\n-----END PRIVATE KEY-----",
                    teamId: "YWQCXGJRJL",
                  },
                },
              ],
            },
          },
        ],
      },
    }),
    Dashboard.init(),
  ],
});

// Middleware
app.use(morgan("dev"));
app.use(
  cors({
    origin: "http://localhost:3005",
    allowedHeaders: ["content-type", ...supertokens.getAllCORSHeaders()],
    credentials: true,
  })
);

// Configure helmet with CSP
app.use(
  helmet({
    contentSecurityPolicy: {
      directives: {
        defaultSrc: ["'self'"],
        scriptSrc: [
          "'self'",
          "'unsafe-inline'",
          "https://google.com",
          "https://cdn.jsdelivr.net/gh/supertokens/",
        ],
        imgSrc: [
          "https://google.com",
          "https://cdn.jsdelivr.net/gh/supertokens/",
        ],
        // Add other directives as needed
      },
    },
  })
);

app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: true }));
app.use(middleware());

app.get("/", (req, res) => {
  res.send("Hello World!");
});

app.listen(port, () => {
  console.log(`Example dashboard app listening at http://localhost:${port}`);
});
