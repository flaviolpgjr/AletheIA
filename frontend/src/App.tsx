import { BrowserRouter, Route, Routes } from "react-router-dom";

import { AppProvider } from "./contexts/AppProvider";
import { HomePage } from "./pages/HomePage";

function App() {
  return (
    <AppProvider>
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<HomePage />} />
          <Route path="/sobre" element={<HomePage />} />
        </Routes>
      </BrowserRouter>
    </AppProvider>
  );
}

export default App;
