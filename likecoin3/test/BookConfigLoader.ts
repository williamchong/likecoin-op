import fs from "fs";

export const BookConfigLoader = {
  load: (path: string) => {
    const bookConfig = fs.readFileSync(path, "utf8");
    const bookConfigJson = JSON.parse(bookConfig);
    const metadata = JSON.stringify(bookConfigJson.metadata);
    bookConfigJson.metadata = metadata;
    return bookConfigJson;
  },
};

export const BookTokenConfigLoader = {
  load: (path: string) => {
    const bookTokenConfig = fs.readFileSync(path, "utf8");
    const bookTokenConfigJson = JSON.parse(bookTokenConfig);
    const tokenString = JSON.stringify(bookTokenConfigJson);
    return tokenString;
  },
};
