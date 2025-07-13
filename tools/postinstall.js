require("concurrently")(["nuxt prepare", "go mod download && go mod tidy"]);
