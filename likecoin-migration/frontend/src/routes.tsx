import { BrowserRouter, Route, Routes } from "react-router";

import Home from "./screens/HomeScreen/HomeScreen";
import MigrateDetailScreen from "./screens/MigrateDetailsScreen/MigrateDetailsScreen";
import MigrateScreen from "./screens/MigrateScreen/MigrateScreen";

export default function _Routes() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/migrate" element={<MigrateScreen />} />
        <Route
          path="/migrate/:cosmosTxHash"
          element={<MigrateDetailScreen />}
        />
      </Routes>
    </BrowserRouter>
  );
}
