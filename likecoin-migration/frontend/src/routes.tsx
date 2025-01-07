import { BrowserRouter, Route, Routes } from "react-router";

import Home from "./screens/HomeScreen/HomeScreen";

export default function _Routes() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Home />} />
      </Routes>
    </BrowserRouter>
  );
}
