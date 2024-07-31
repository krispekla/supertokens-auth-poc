require("dotenv").config();

const express = require("express");
const dotenv = require("dotenv");
const morgan = require("morgan");
const cors = require("cors");
const helmet = require("helmet");
const bodyParser = require("body-parser");

dotenv.config();

const app = express();
const port = process.env.PORT || 3008;

// Middleware
app.use(morgan("dev"));
app.use(cors());
app.use(helmet());
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: true }));

app.get("/", (req, res) => {
  res.send("Hello World!");
});

app.listen(port, () => {
  console.log(`Example dashboard app listening at http://localhost:${port}`);
});
