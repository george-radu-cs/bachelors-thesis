import http from "http";
import https from "https";

var server;
if (GetBooleanEnvVar("IN_PRODUCTION")) {
  server = https.createServer(
    {
      cert: fs.readFileSync(GetStringEnvVar("SERVER_CERT_FILE")),
      key: fs.readFileSync(GetStringEnvVar("SERVER_PRIVATE_KEY_FILE")),
    },
    app
  );
} else {
  server = http.createServer(app);
}
