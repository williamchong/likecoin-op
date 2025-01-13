import "./App.css";

import { QueryClient, QueryClientProvider } from "@tanstack/react-query";

import { ConfigProvider } from "./providers/Config";
import Routes from "./routes";

const queryClient = new QueryClient();

function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <ConfigProvider>
        <Routes />
      </ConfigProvider>
    </QueryClientProvider>
  );
}

export default App;
