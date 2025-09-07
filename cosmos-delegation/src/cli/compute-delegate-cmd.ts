import { z } from "zod";
import computeDelegateCmd from "../logic/compute-delegate-cmd";
import CmdLikedClient from "../liked-client";
import { Command } from "commander";
import Decimal from "decimal.js";

const InvokeArgsSchema = z.object({
  cli: z.string(),
  cliFlags: z.string(),
  delegatorAddress: z.string(),
  denom: z.string(),
  selectedValidatorAddresses: z.array(z.string()),
  thresholdPercentage: z.string().transform((val) => new Decimal(val)),
});

function init(program: Command) {
  program
    .command("compute-delegate-cmd")
    .description("Compute the command to delegate")
    .argument("<delegator-address>", "The address of the delegator")
    .option(
      "--selected-validator-address <addresses...>",
      "The addresses of the selected validators",
      [],
    )
    .action(async (delegatorAddress, options) => {
      const invokeArgs = InvokeArgsSchema.parse({
        cli: process.env.CLI,
        cliFlags: process.env.CLI_FLAGS,
        delegatorAddress,
        selectedValidatorAddresses: options.selectedValidatorAddress,
        denom: process.env.DENOM,
        thresholdPercentage: process.env.THRESHOLD_PERCENT,
      });

      const likedClient = new CmdLikedClient(
        invokeArgs.cli,
        invokeArgs.cliFlags,
      );

      const cmds = await computeDelegateCmd(
        likedClient,
        invokeArgs.delegatorAddress,
        invokeArgs.selectedValidatorAddresses,
        invokeArgs.denom,
        invokeArgs.thresholdPercentage,
      );

      console.error("Commands:");

      for (const cmd of cmds) {
        console.log(`${cmd};`);
      }
    });
}

export default init;
