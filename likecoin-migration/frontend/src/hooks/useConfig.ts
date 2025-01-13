import { useContext } from "react";

import ConfigContext from "../contexts/ConfigContext";
import { Config } from "../models/config";

export function useConfig(): Config {
  return useContext(ConfigContext);
}
