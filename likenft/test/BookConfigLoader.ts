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
