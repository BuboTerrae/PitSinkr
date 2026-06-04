import Main from "./components/Main";
import ThemeProvider from "./utils/ThemeContext";
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import './App.css';
import LogsPage from "./pages/LogsPage";

function App() {
  return (
    <>
      <ThemeProvider>
        <Router>
          <Routes>
            <Route path="/" element={<Main />} />
            <Route path="/logs" element={<LogsPage />} />
          </Routes>
        </Router>
      </ThemeProvider>
    </>
  )
}

export default App
