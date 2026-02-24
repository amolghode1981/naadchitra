import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { useState } from "react";
import "./App.css";
import App from "./App.jsx";
import Splash from "./components/landing/splash.jsx"

function Root() {
  const [started, setStarted] = useState(false);

  if (started) {
    return <App />;
  }

  return <Splash onStart={() => setStarted(true)} />;
}
createRoot(document.getElementById("root")).render(
  <StrictMode>
    {/* <App /> */}
    <Root/>
  </StrictMode>
);
