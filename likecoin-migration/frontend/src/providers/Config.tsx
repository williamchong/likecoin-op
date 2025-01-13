import React, { useEffect, useState } from "react";

import ConfigContext from "../contexts/ConfigContext";
import { Config, ConfigSchema } from "../models/config";

export function ConfigProvider(props: React.PropsWithChildren) {
  const [config, setConfig] = useState<Config | null>(null);

  useEffect(() => {
    fetch("/config.json")
      .then((d) => d.json())
      .then((d) => ConfigSchema.parseAsync(d))
      .then(setConfig);
  }, []);

  return config != null ? (
    <ConfigContext.Provider value={config}>
      {props.children}
    </ConfigContext.Provider>
  ) : null;
}
