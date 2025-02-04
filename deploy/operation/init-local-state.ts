import fundOperator from "./src/fund-operator";

async function initLocalState() {
  console.log("Initializing local state");

  await fundOperator();
}

initLocalState()
  .then(() => {
    process.exit(0);
  })
  .catch((error) => {
    console.error(error);
    process.exitCode = 1;
  });
