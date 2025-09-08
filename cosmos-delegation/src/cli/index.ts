import { Command } from "commander";
import { default as initComputeDelegateAmount } from "./compute-delegate-cmd";

const program = new Command();

program.name("cosmos-delegation");

initComputeDelegateAmount(program);

export default program;
