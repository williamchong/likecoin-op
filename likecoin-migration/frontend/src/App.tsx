import "./App.css";

import { ConfigProvider } from "./providers/Config";
import Routes from "./routes";

function App() {
  return (
    <>
      <ConfigProvider>
        <Routes />
      </ConfigProvider>
    </>
  );
}

export default App;
