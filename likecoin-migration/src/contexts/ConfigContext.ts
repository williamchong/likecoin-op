import React from "react";

import { Config } from "../models/config";

// eslint-disable-next-line @typescript-eslint/no-explicit-any
const ConfigContext = React.createContext<Config>(null as any);

export default ConfigContext;
